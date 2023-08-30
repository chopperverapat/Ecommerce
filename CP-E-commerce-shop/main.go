package main

import (
	_ "fmt"
	"os"

	"cpshop/config"
	"cpshop/modules/servers"
	"cpshop/pkg/databases"
)

// test
// func main() {
// 	cfg := config.LoadConfig(envPath())
// 	// fmt.Printf("APP : %v\n", cfg.App())
// 	fmt.Printf("DB : %v\n", cfg.Db())
// 	// fmt.Printf("JWT : %v\n", cfg.Jwt())
// 	// fmt.Printf("APP URL : %v\n", cfg.App().Url())
// 	// fmt.Printf("DB URL : %v\n", cfg.Db().Url())
// 	// fmt.Printf("JWT Secret key : %v\n", cfg.Jwt().SecretKey())
// 	db := databases.ConnectDB(cfg.Db())
// 	fmt.Printf("DB connected %v", db)
// }

func main() {
	// init config
	cfg := config.LoadConfig(envPath())
	// init db
	db := databases.ConnectDB(cfg.Db())
	// fmt.Printf("DB connected %v", db)
	defer db.Close()
	// init server
	servers.Newservers(cfg, db).Start()

}

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
