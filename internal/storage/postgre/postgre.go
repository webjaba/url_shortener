package postgre

import (
	"errors"
	"fmt"
	"url-shortener/internal/storage/models"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDatabase struct {
	Port     string `env:"PORT" env-default:"5432"`
	Host     string `env:"HOST" env-default:"localhost"`
	Name     string `env:"NAME" env-default:"url_shortener"`
	User     string `env:"USER" env-default:"postgres"`
	Password string `env:"PASSWORD"`
}

type PostgreStorage struct {
	db *gorm.DB
}

func ConnectDB() *gorm.DB {

	DBConfig := ConfigDatabase{}

	err := cleanenv.ReadConfig(".env", &DBConfig)

	if err != nil {
		panic(fmt.Sprintf("Unable to load db_config: %v", err.Error()))
	}

	if DBConfig.Password == "" {
		panic("password must be specified")
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		DBConfig.Host,
		DBConfig.User,
		DBConfig.Password,
		DBConfig.Name,
		DBConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitStorage(db *gorm.DB) *PostgreStorage {
	return &PostgreStorage{db: db}
}

func (s *PostgreStorage) GetURL(alias string) (string, error) {
	url := models.Url{}

	res := s.db.First(&url, "alias = ?", alias)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("url not found")
		} else {
			return "", fmt.Errorf("error of getting url: %v", res.Error)
		}
	}

	return url.Url, nil
}

func (s *PostgreStorage) AddURL(url, alias string) error {

	urlObj := models.Url{}

	res := s.db.First(&urlObj, "alias = ?", alias)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			res := s.db.Create(&models.Url{Url: url, Alias: alias})

			if res.Error != nil {
				return fmt.Errorf("error of creating url: %v", res.Error)
			}

			return nil
		} else {
			return fmt.Errorf("error of creating url: %v", res.Error)
		}
	}

	return errors.New("url already exists")
}
