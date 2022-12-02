package postgres

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"work_view/internal/adapter/spi/repository/postgres/models"
	"work_view/internal/domain"
)

type DB struct {
	db *gorm.DB
}

func New(dsn string) (*DB, error) {
	d := &DB{}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
	return d, nil
}
func Migrate() {
	//db, err := gorm.Open(postgres.Open("postgresql://root@localhost:26257?sslmode=disable"), &gorm.Config{})
	db, err := gorm.Open(postgres.Open("postgresql://root@db:26257?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{}, &models.Project{})
}

func (d *DB) ProjectCreate(title string) (domain.Project, error) {
	project := models.Project{
		Title: title,
	}
	result := d.db.Create(&project)

	if result.Error != nil {
		return domain.Project{}, result.Error
	}

	return project.MainTo(), nil
}

func (d *DB) ProjectGet(id uuid.UUID) (domain.Project, error) {
	var project = models.Project{ModelBase: models.ModelBase{ID: id}}
	d.db.Take(&project)

	return project.MainTo(), nil
}
func (d *DB) ProjectDeleteById(id uuid.UUID) error {

	return nil
}
