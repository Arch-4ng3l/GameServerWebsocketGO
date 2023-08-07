package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Arch-4ng3l/GoServerHololens/types"
	"github.com/Arch-4ng3l/GoServerHololens/util"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "hololens"
)

var password = os.Getenv("PSQLPW")
var user = os.Getenv("PSQLUN")

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
	query2 := `
	CREATE TABLE IF NOT EXISTS users (
		user_id serial PRIMARY KEY, 
		username TEXT UNIQUE NOT NULL, 
		password TEXT UNIQUE NOT NULL
	);
	`

	query3 := `
	INSERT INTO users (username, password)
	VALUES($1, $2)
	ON CONFLICT DO NOTHING
	`

	_, err := psql.DB.Exec(query)
	if err != nil {
		log.Panic(err)

		return err
	}

	_, err = psql.DB.Exec(query2)
	if err != nil {
		log.Panic(err)

		return err
	}
	_, err = psql.DB.Exec(query3, "admin", util.CreateHash("admin"))

	if err != nil {
		log.Panic(err)

		return err
	}
	return nil

}

func (psql *Postgres) GetUser(username string) *types.User {
	query := `
	SELECT * FROM users WHERE username=$1
	`

	res, err := psql.DB.Query(query, username)
	if err != nil {

		return nil
	}
	user := &types.User{}
	id := 0
	for res.Next() {
		res.Scan(&id, &user.Username, &user.Password)

		return user
	}

	return nil
}

func (psql *Postgres) GetObjects(out chan *types.Object, player *types.Object) error {

	query := `SELECT * FROM objects`

	rows, err := psql.DB.Query(query)
	if err != nil {

		return err
	}

	defer rows.Close()
	id := 0
	for rows.Next() {
		object := types.Object{}
		err = rows.Scan(&id, &object.Name, &object.X, &object.Y, &object.Z)
		if err != nil {
			continue
		}

		if object.Name == "" {
			continue
		}
		if util.CalcDistance(player, &object) <= types.RenderDistance {

			out <- &object
		}
	}

	close(out)

	return nil
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

func (psql *Postgres) GetObjectsWeb(startID int) ([]*types.Object, error) {
	query := `
		SELECT * FROM objects WHERE object_id >= $1 AND object_id <= $1 + 19
	`
	res, err := psql.DB.Query(query, startID)
	if err != nil {

		return nil, err
	}
	var objects []*types.Object

	for res.Next() {
		obj := types.Object{}
		res.Scan(&startID, &obj.Name, &obj.X, &obj.Y, &obj.Z)

		objects = append(objects, &obj)
	}

	return objects, nil
}

func (psql *Postgres) GetAmountOfObjects() uint {
	query := `
	SELECT COUNT(*) FROM objects
	`

	var num uint
	psql.DB.QueryRow(query).Scan(&num)
	fmt.Println(num)
	return uint(num)
}
