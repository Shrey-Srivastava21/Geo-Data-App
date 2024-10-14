package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
}

// internal/models/geospatial.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Geospatial struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	GeoJSON     string `json:"geojson"`
	KML         string `json:"kml"`
}
