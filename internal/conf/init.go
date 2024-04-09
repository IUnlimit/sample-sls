package conf

import (
	"github.com/IUnlimit/sample-sls/configs"
	global "github.com/IUnlimit/sample-sls/internal"
	log "github.com/sirupsen/logrus"
)

func Init() {
	versionCheck()
	fileFolder := global.ParentPath + "/"
	_, err := LoadConfig(configs.ConfigFileName, fileFolder, "yaml", configs.Config, &global.Config)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Infof("Current perpetua instance version: %s", Version)
}
