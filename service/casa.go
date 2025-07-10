package service

import (
	json2 "encoding/json"
	"time"
	"fmt"
	"github.com/BeesNestInc/CassetteOS/model"
	"github.com/BeesNestInc/CassetteOS/pkg/utils/httper"
	"github.com/tidwall/gjson"
)

type CasaService interface {
	GetCasaosVersion() model.Version
}

type casaService struct{}

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}
/**
 * @description: get remote version
 * @return {model.Version}
 */
func (o *casaService) GetCasaosVersion() model.Version {
	keyName := "casa_version"
	var dataStr string
	var version model.Version
	if result, ok := Cache.Get(keyName); ok {
		dataStr, ok = result.(string)
		if ok {
			if err := json2.Unmarshal([]byte(dataStr), &version); err == nil {
				return version
			} else {
				fmt.Println("⚠️ キャッシュのデコード失敗:", err)
			}
		}
	}
	v := httper.Get("https://api.github.com/repos/BeesNestInc/CassetteOS-Tools/releases/latest", nil)
	if !gjson.Valid(v) {
		fmt.Println("⚠️ 無効なJSONレスポンス:", v)
		return model.Version{}
	}

	var githubRelease GitHubRelease
	if err := json2.Unmarshal([]byte(v), &githubRelease); err != nil {
		fmt.Println("⚠️ GitHubレスポqンスのデコード失敗:", err)
		return model.Version{}
	}

	version.Version = githubRelease.TagName
	version.ChangeLog = githubRelease.Body

	if len(version.Version) > 0 {
		Cache.Set(keyName, version, time.Minute*20)
	}

	return version
}

func NewCasaService() CasaService {
	return &casaService{}
}
