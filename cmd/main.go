package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"movies/configs"

	_ "github.com/lib/pq"
)

var (

	postgresHost =  flag.String("postgres connection",configs.Envs.DB_HOST, "postgres host")
	postgresName =  flag.String("postgres connection",configs.Envs.DBName, "postgres name")
	postgresPort =  flag.String("postgres connection",configs.Envs.DB_PORT, "postgres port")
	postgresUser =  flag.String("postgres connection",configs.Envs.DBUser, "postgres User")
	postgresPassword =  flag.String("postgres connection",configs.Envs.DBPassword, "postgres Password")

)

func main(){
	log := slog.Default()

	conn := fmt.Sprintf("host=%s port=%s user=%s " +
	    "password=%s dbname=%s sslmode=disable",
		*postgresHost, *postgresPort, *postgresUser, *postgresPassword, *postgresName )

	db, err := sql.Open("postgres",conn)
	if err != nil{
		log.Error("", "can't connect to database %s",err)
	}
	defer db.Close()











	

}











