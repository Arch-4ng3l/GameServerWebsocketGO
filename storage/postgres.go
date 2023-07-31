package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Arch-4ng3l/GoServerHololens/types"
	"github.com/Arch-4ng3l/GoServerHololens/util"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "moritz"
	password = "postgres"
	dbname   = "hololens"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres() *Postgres {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Panic(err)
		return nil
	}

	return &Postgres{
		DB: db,
	}
}

func (psql *Postgres) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS objects ( 
		object_id serial PRIMARY KEY,
		object_name TEXT UNIQUE NOT NULL,
		x FLOAT4, 
		y FLOAT4, 
		z FLOAT4
	);
	`

	_, err := psql.DB.Exec(query)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil

}

func (psql *Postgres) GetObjects(player *types.Object) ([]*types.Object, error) {

	query := `SELECT * FROM objects`

	rows, err := psql.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var objects []*types.Object
	id := 0
	for rows.Next() {
		object := types.Object{}
		err = rows.Scan(&id, &object.Name, &object.X, &object.Y, &object.Z)
		if err != nil {
			continue
		}
		if util.CalcDistance(player, &object) <= types.RenderDistance {
			objects = append(objects, &object)
		}
	}

	return objects, nil
}

func (psql *Postgres) UpdateObject(object *types.Object) error {

	query := `
	INSERT INTO objects ("object_name", "x", "y", "z")
	VALUES ($1, $2, $3, $4)
	ON CONFLICT ("object_name") DO UPDATE SET "x" = $2, "y" = $3, "z" = $4;
	`

	_, err := psql.DB.Exec(query, object.Name, object.X, object.Y, object.Z)

	return err
}
