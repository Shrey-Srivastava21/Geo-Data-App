package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"geo-data-app/internal/models"
	"geo-data-app/internal/database"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your_jwt_secret") // Store securely

// Register handler
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user to database
	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var dbUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": dbUser.Username,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// UploadGeoJSON handler
func UploadGeoJSON(w http.ResponseWriter, r *http.Request) {
	// Assume token is passed as Authorization header
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var geospatial models.Geospatial
	err = json.NewDecoder(r.Body).Decode(&geospatial)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Store the geoJSON
	if err := database.DB.Create(&geospatial).Error; err != nil {
		http.Error(w, "Could not save geospatial data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(geospatial)
}

// RetrieveGeoJSON handler
func RetrieveGeoJSON(w http.ResponseWriter, r *http.Request) {
	// Assume token is passed as Authorization header
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var geospatial []models.Geospatial
	if err := database.DB.Find(&geospatial).Error; err != nil {
		http.Error(w, "Could not retrieve geospatial data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(geospatial)
}
