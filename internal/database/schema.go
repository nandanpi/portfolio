package database

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null; size:255"`
	Role     string `json:"role"`
}

type Work struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	DemoLink    string `json:"demo_link"`
	Image       string `json:"image"`

	Published bool `json:"published" gorm:"default:false"`

	TechStack []*Technology `json:"tech_stack" gorm:"many2many:work_technologies;"`
}

type Resume struct {
	gorm.Model

	Title string `json:"title"`
	Link  string `json:"link"`

	Published bool `json:"published" gorm:"default:false"`
}

type Contact struct {
	gorm.Model

	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Technology struct {
	gorm.Model

	Name string `json:"name"`
	Logo string `json:"logo"`

	Published bool `json:"published" gorm:"default:false"`

	Works []*Work `json:"works" gorm:"many2many:work_technologies;"`
}

type Experience struct {
	gorm.Model

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Logo        string    `json:"logo"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`

	Published bool `json:"published" gorm:"default:false"`
}

type Education struct {
	gorm.Model

	School    string    `json:"school"`
	Degree    string    `json:"degree"`
	Aggregate string    `json:"aggregate"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Published bool      `json:"published" gorm:"default:false"`
}

type Achievement struct {
	gorm.Model

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
	Published   bool      `json:"published" gorm:"default:false"`
}

type ExtraCurricular struct {
	gorm.Model

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
	Published   bool      `json:"published" gorm:"default:false"`
}

func PushSchema(db *gorm.DB) {
	err := db.AutoMigrate(&Users{}, &Work{}, &Resume{}, &Contact{}, &Technology{}, &Experience{}, &Education{}, &Achievement{}, &ExtraCurricular{})
	if err != nil {
		panic("Failed to migrate database schema")
	}
	log.Println("Migrated DB schema")
}
