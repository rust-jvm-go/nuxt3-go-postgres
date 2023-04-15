package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"nuxt3-go-postgres/db"
	"nuxt3-go-postgres/web"
	"os"
)

func main() {
	d, err := sql.Open("postgres", dataSource())
	if err != nil {
		log.Fatal(err)
	}
	defer func(d *sql.DB) {
		err := d.Close()
		if err != nil {
			log.Println("Error", err)
		}
	}(d)
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(db.NewDB(d), cors)
	err = app.Serve()
	log.Println("Error", err)
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "postgresql://" + host + ":5432/nuxt3go" +
		"?user=nuxt3go&sslmode=disable&password=" + pass
}
