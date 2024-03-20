package cmd

import (
	"fmt"
	"net/http"
	"os"

	core "gitlab.finema.co/finema/idin-core"
)

func APIRun() {
	env := core.NewEnv()
	db, err := core.NewDatabase(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "MySQL: %v", err)
		os.Exit(1)
	}

	// if db != nil {
	// 	db.Exec("SET GLOBAL sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));")
	// }

	e := core.NewHTTPServer(&core.HTTPContextOptions{
		AllowOrigins: []string{"*"},
		ContextOptions: &core.ContextOptions{
			DB:  db,
			ENV: env,
		},
	})

	e.GET("/healthz", core.WithHTTPContext(func(c core.IHTTPContext) error {
		return c.String(http.StatusOK, "OK!")
	}))

	core.StartHTTPServer(e, env)

}
