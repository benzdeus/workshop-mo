package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "user:password@tcp(localhost:3306)/databaseName?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	DB = db

	dbClose, _ := db.DB()
	defer dbClose.Close()

	e := echo.New()

	e.GET("/bookings", GetBooking)
	e.Logger.Fatal(e.Start(":1323"))
}

func GetBooking(c echo.Context) error {
	bookings := new([]Bookings)
	DB.
		Select([]string{
			"bookings.id AS id",
			"first_name",
			"last_name",
			"start_date",
			"end_date",
			"maximum_person",
			`
		(CASE
			WHEN sum_grade='A' THEN 'ดีมาก'
			WHEN sum_grade='B' THEN 'ดี'
			WHEN sum_grade='C' THEN 'พอใช้'
			WHEN sum_grade='D' THEN 'ปรับปรุง'
			ELSE '-'
			END
		) AS sum_grade
		`,
		}).
		Joins("INNER JOIN users ON users.id = bookings.users_id").
		Joins("INNER JOIN rooms ON rooms.id = bookings.rooms_id").
		Find(&bookings)
	return c.JSON(http.StatusOK, bookings)
}

type Bookings struct {
	ID            int        `json:"id"`
	FirstName     string     `json:"frist_name"`
	LastName      string     `json:"last_name"`
	StartDate     *time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date"`
	MaximumPerson int        `json:"maximum_person"`
	SumGrade      string     `json:"sum_grade"`
}

func (b Bookings) TableName() string {
	return "bookings"
}
