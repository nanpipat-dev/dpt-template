package main

import (
	"log"

	"PROJECT_GIT_PATH/cmd"
	"PROJECT_GIT_PATH/consts"

	core "gitlab.finema.co/finema/idin-core"
)

func main() {
	// Code
	switch core.NewEnv().Config().Service {
	case consts.ServiceAPI:
		cmd.APIRun()
	// case consts.ServiceMigration:
	// 	cmd.MigrationRun()
	// case consts.ServiceCronjob:
	// 	cmd.CronjobRun()
	default:
		log.Fatal("Service not support")
	}
}
