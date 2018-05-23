package context

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Change data struct
type Change struct {
	gorm.Model

	Body string
	Date time.Time
}
