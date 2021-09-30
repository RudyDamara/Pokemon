package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"crud/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	myPokemon struct {
		Name string `json:"name" gorm:"name"`
		URL  string `json:"url" gorm:"url"`
	}

	getMyPokemon struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		URLPokemon  string `json:"url_pokemon"`
		NamePokemon string `json:"name_pokemon"`
	}

	countData struct {
		CountData string `json:"count_data"`
	}

	dbConnection struct {
		Db *gorm.DB
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func catchPokemon(c echo.Context) error {

	dbConn := connectionDb()

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	t := myPokemon{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	if t.Name != "" {
		query := `with select_data as (select count(*) as count_data from my_catch 
	where name_pokemon = ?) INSERT INTO my_catch ("name", url_pokemon, name_pokemon)
	values (concat(?::text, '-', (select count_data::text from select_data)), ?, ?)`
		err = dbConn.Db.Exec(query, t.Name, t.Name, t.URL, t.Name).Error
		if err != nil {
			return models.ToJSON(c).BadRequest("Bad Request")
		}

	}
	return models.ToJSON(c).Ok(err, "Successfully")
}

func releasePokemon(c echo.Context) error {
	id := c.Param("id")

	dbConn := connectionDb()

	query := `Delete from my_catch where id = ?`
	err := dbConn.Db.Exec(query, id).Error
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	return models.ToJSON(c).Ok(err, "Successfully")
}

func editPokemon(c echo.Context) error {
	dbConn := connectionDb()

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	t := myPokemon{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	query := `Update my_catch set "name" = ? where id = ?`
	errq := dbConn.Db.Exec(query, t.Name, c.Param("id")).Error
	if errq != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	return models.ToJSON(c).Ok(err, "Successfully")
}

func getPokemon(c echo.Context) error {

	dbConn := connectionDb()
	var pokemon []getMyPokemon
	query := `select * from my_catch`
	err := dbConn.Db.Raw(query).Scan(&pokemon).Error
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}
	return models.ToJSON(c).Ok(pokemon, "Successfully")
}

func connectionDb() *dbConnection {
	dsn := "host=localhost user=postgres password=newpostgres dbname=pokemon port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dbname, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if err != nil {
			log.Error(err.Error())
			log.Info("failed to connect database")
		}
	}
	return &dbConnection{dbname}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	// Routes
	e.GET("/pokemon", getPokemon)
	e.POST("/pokemon/catch", catchPokemon)
	e.PUT("/pokemon/:id", editPokemon)
	e.DELETE("/pokemon/:id", releasePokemon)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	err := e.Start(":1323")
	if err != nil {
		log.Error("error start")
	}
}
