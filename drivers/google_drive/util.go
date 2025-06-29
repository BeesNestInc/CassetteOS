package google_drive

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/BeesNestInc/CassetteOS-Common/utils/logger"
	"github.com/BeesNestInc/CassetteOS/drivers/base"
	"github.com/BeesNestInc/CassetteOS/model"
	"github.com/BeesNestInc/CassetteOS/pkg/utils"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	client_id     = "private build"
	client_secret = "private build"
)

// do others that not defined in Driver interface

func (d *GoogleDrive) getRefreshToken() error {
	url := "https://www.googleapis.com/oauth2/v4/token"
	var resp base.TokenResp
	var e TokenError
	res, err := base.RestyClient.R().SetResult(&resp).SetError(&e).
		SetFormData(map[string]string{
			"client_id":     d.ClientID,
			"client_secret": d.ClientSecret,
			"code":          d.Code,
			"grant_type":    "authorization_code",
			"redirect_uri":  "https://cloudoauth.files.casaos.app",
		}).Post(url)
	if err != nil {
		return err
	}
	logger.Info("get refresh token", zap.String("res", res.String()))
	if e.Error != "" {
		return fmt.Errorf(e.Error)
	}
	d.RefreshToken = resp.RefreshToken
	return nil
}

func (d *GoogleDrive) refreshToken() error {
	url := "https://www.googleapis.com/oauth2/v4/token"
	var resp base.TokenResp
	var e TokenError
	res, err := base.RestyClient.R().SetResult(&resp).SetError(&e).
		SetFormData(map[string]string{
			"client_id":     d.ClientID,
			"client_secret": d.ClientSecret,
			"refresh_token": d.RefreshToken,
			"grant_type":    "refresh_token",
		}).Post(url)
	if err != nil {
		return err
	}
	log.Debug(res.String())
	if e.Error != "" {
		return fmt.Errorf(e.Error)
	}
	d.AccessToken = resp.AccessToken
	return nil
}

func (d *GoogleDrive) request(url string, method string, callback base.ReqCallback, resp interface{}) ([]byte, error) {
	req := base.RestyClient.R()
	req.SetHeader("Authorization", "Bearer "+d.AccessToken)
	req.SetQueryParam("includeItemsFromAllDrives", "true")
	req.SetQueryParam("supportsAllDrives", "true")
	if callback != nil {
		callback(req)
	}
	if resp != nil {
		req.SetResult(resp)
	}
	var e Error
	req.SetError(&e)
	res, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}
	if e.Error.Code != 0 {
		if e.Error.Code == 401 {
			err = d.refreshToken()
			if err != nil {
				return nil, err
			}
			return d.request(url, method, callback, resp)
		}
		return nil, fmt.Errorf("%s: %v", e.Error.Message, e.Error.Errors)
	}
	return res.Body(), nil
}

func (d *GoogleDrive) getFiles(id string) ([]File, error) {
	pageToken := "first"
	res := make([]File, 0)
	for pageToken != "" {
		if pageToken == "first" {
			pageToken = ""
		}
		var resp Files
		orderBy := "folder,name,modifiedTime desc"
		if d.OrderBy != "" {
			orderBy = d.OrderBy + " " + d.OrderDirection
		}
		query := map[string]string{
			"orderBy":  orderBy,
			"fields":   "files(id,name,mimeType,size,modifiedTime,thumbnailLink,shortcutDetails),nextPageToken",
			"pageSize": "1000",
			"q":        fmt.Sprintf("'%s' in parents and trashed = false", id),
			//"includeItemsFromAllDrives": "true",
			//"supportsAllDrives":         "true",
			"pageToken": pageToken,
		}
		_, err := d.request("https://www.googleapis.com/drive/v3/files", http.MethodGet, func(req *resty.Request) {
			req.SetQueryParams(query)
		}, &resp)
		if err != nil {
			return nil, err
		}
		pageToken = resp.NextPageToken
		res = append(res, resp.Files...)
	}
	return res, nil
}

func (d *GoogleDrive) chunkUpload(ctx context.Context, stream model.FileStreamer, url string) error {
	var defaultChunkSize = d.ChunkSize * 1024 * 1024
	var finish int64 = 0
	for finish < stream.GetSize() {
		if utils.IsCanceled(ctx) {
			return ctx.Err()
		}
		chunkSize := stream.GetSize() - finish
		if chunkSize > defaultChunkSize {
			chunkSize = defaultChunkSize
		}
		_, err := d.request(url, http.MethodPut, func(req *resty.Request) {
			req.SetHeaders(map[string]string{
				"Content-Length": strconv.FormatInt(chunkSize, 10),
				"Content-Range":  fmt.Sprintf("bytes %d-%d/%d", finish, finish+chunkSize-1, stream.GetSize()),
			}).SetBody(io.LimitReader(stream.GetReadCloser(), chunkSize)).SetContext(ctx)
		}, nil)
		if err != nil {
			return err
		}
		finish += chunkSize
	}
	return nil
}
func GetConfig() GoogleDrive {
	config := GoogleDrive{}
	config.ClientID = client_id
	config.ClientSecret = client_secret
	config.RootFolderID = "root"
	config.AuthUrl = "https://accounts.google.com/o/oauth2/auth/oauthchooseaccount?response_type=code&client_id=" + client_id + "&redirect_uri=https%3A%2F%2Fcloudoauth.files.casaos.app&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fdrive&access_type=offline&approval_prompt=force&state=${HOST}%2Fv1%2Frecover%2FGoogleDrive&service=lso&o2v=1&flowName=GeneralOAuthFlow"
	config.Icon = "./img/driver/GoogleDrive.svg"

	return config
}
