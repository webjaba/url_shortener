package postgre

import (
	"errors"
	"fmt"
	apperrors "url-shortener/internal/app_errors"
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
			return "", apperrors.ErrURLNotFound
		} else {
			return "", fmt.Errorf("error of getting url: %v", res.Error)
		}
	}

	return url.Url, nil
}

func (s *PostgreStorage) AddURL(url, alias string) (string, error) {

	urlObj := models.Url{}

	res := s.db.Where("alias = ? OR url = ?", alias, url).First(&urlObj)

	if res.Error == nil {
		if urlObj.Alias == alias {
			return "", apperrors.ErrAliasAlreadyOccupied
		} else {
			return urlObj.Alias, apperrors.ErrURLAlreadyExists
		}
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res := s.db.Create(&models.Url{Url: url, Alias: alias})

		if res.Error != nil {
			return "", fmt.Errorf("error of creating url: %v", res.Error)
		}

		return alias, nil
	} else {
		return "", fmt.Errorf("error of creating url: %v", res.Error)
	}
}
