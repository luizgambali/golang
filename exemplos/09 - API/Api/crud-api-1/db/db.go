package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

// Movie struct
type Movie struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// inicializa a conexão com o banco de dados
func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Movie{})
}

// adiciona um filme
func CreateMovie(movie *Movie) (*Movie, error) {
	movie.ID = uuid.New().String()
	res := db.Create(&movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie, nil
}

// retorna os dados de um filme
func GetMovie(id string) (*Movie, error) {
	var movie Movie
	res := db.First(&movie, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("movie of id %s not found", id))
	}
	return &movie, nil
}

// retorna todos os filmes
func GetMovies() ([]*Movie, error) {
	var movies []*Movie
	res := db.Find(&movies)
	if res.Error != nil {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}

// atualiza um filme
func UpdateMovie(movie *Movie) (*Movie, error) {
	var movieToUpdate Movie
	result := db.Model(&movieToUpdate).Where("id = ?", movie.ID).Updates(movie)
	if result.RowsAffected == 0 {
		return &movieToUpdate, errors.New("movie not updated")
	}
	return movie, nil
}

// exclui um filme
func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Where("id = ?", id).Delete(&deletedMovie)
	if result.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
}
