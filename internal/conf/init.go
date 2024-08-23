package conf

import (
	"github.com/IUnlimit/sample-sls/configs"
	global "github.com/IUnlimit/sample-sls/internal"
	"github.com/IUnlimit/sample-sls/internal/model"
	log "github.com/sirupsen/logrus"
)

// Config perpetua config.yml
var Config *model.Config

// Init method should be invoked manually after global.ParentPath be set
func Init() {
	versionCheck()
	fileFolder := global.ParentPath + "/"
	_, err := LoadConfig(configs.ConfigFileName, fileFolder, "yaml", configs.Config, &Config)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Debugf("Current perpetua instance version: %s", Version)
}
