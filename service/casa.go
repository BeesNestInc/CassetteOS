package service

import (
	json2 "encoding/json"
	"time"

	"github.com/BeesNestInc/CassetteOS/model"
	"github.com/BeesNestInc/CassetteOS/pkg/config"
	"github.com/BeesNestInc/CassetteOS/pkg/utils/httper"
	"github.com/tidwall/gjson"
)

type CasaService interface {
	GetCasaosVersion() model.Version
}

type casaService struct{}

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
			data := gjson.Get(dataStr, "data")
			json2.Unmarshal([]byte(data.String()), &version)
			return version
		}
	}

	v := httper.Get(config.ServerInfo.ServerApi + "/v1/sys/version", nil)
	data := gjson.Get(v, "data")
	json2.Unmarshal([]byte(data.String()), &version)
	if len(version.Version) > 0 {
		Cache.Set(keyName, v, time.Minute*20)
	}

	return version
}

func NewCasaService() CasaService {
	return &casaService{}
}
