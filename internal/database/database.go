package database

import (
	"fmt"
	"portfolio/internal/types"

	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	GetUser(c echo.Context, username string) *User

	AddWork(req types.AddWorkReq) error
	UpdateWork(req types.UpdateWorkReq) error
	DeleteWork(id uint) error
	ToggleWorkPublished(id uint) error
	GetPublishedWorks() ([]Work, error)
	GetAllWorks() ([]Work, error)

	AddContact(req types.AddContactReq) error
	GetContacts() ([]Contact, error)

	AddAchievement(req types.AddAchievementReq) error
	UpdateAchievement(req types.UpdateAchievementReq) error
	DeleteAchievement(id uint) error
	ToggleAchievementPublished(id uint) error
	GetPublishedAchievements() ([]Achievement, error)
	GetAllAchievements() ([]Achievement, error)

	CreateEducation(req types.AddEducationReq) error
	UpdateEducation(req types.UpdateEducationReq) error
	DeleteEducation(id uint) error
	ToggleEducationPublished(id uint) error
	GetPublishedEducations() ([]Education, error)
	GetAllEducations() ([]Education, error)

	AddExperience(req types.AddExperienceReq) error
	UpdateExperience(req types.UpdateExperienceReq) error
	DeleteExperience(id uint) error
	ToggleExperiencePublish(id uint) error
	GetPublishedExperiences() ([]Experience, error)
	GetAllExperiences() ([]Experience, error)

	AddExtraCurricular(req types.AddExtraCurricularReq) error
	UpdateExtraCurricular(req types.UpdateExtraCurricularReq) error
	DeleteExtraCurricular(id uint) error
	ToggleExtraCurricularPublish(id uint) error
	GetPublishedExtraCurriculars() ([]ExtraCurricular, error)
	GetAllExtraCurriculars() ([]ExtraCurricular, error)

	AddResume(req types.AddResumeReq) error
	DeleteResume(id uint) error
	GetPublishedResume() (Resume, error)
	GetAllResume() ([]Resume, error)

	AddTechnology(req types.AddTechnologyReq) error
	DeleteTechnology(id uint) error
	ToggleTechnologyStatus(id uint) error
	GetPublishedTechnologies() ([]Technology, error)
}

type service struct {
	db *gorm.DB
}

var (
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DB")
)

func New() Service {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	dbInstance := &service{
		db: db,
	}

	PushSchema(dbInstance.db)

	log.Println("Connected to database")
	return dbInstance
}
