package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtSecret = []byte("your-secret-key-change-in-production")

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Attendance struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	CheckIn   time.Time `json:"check_in"`
	CheckOut  *time.Time `json:"check_out"`
	Date      string    `json:"date"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func initDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=pstgres dbname=absensi_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createTables()
	fmt.Println("Database connected successfully!")
}

func createTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	attendanceTable := `
	CREATE TABLE IF NOT EXISTS attendances (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		check_in TIMESTAMP NOT NULL,
		check_out TIMESTAMP,
		date DATE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(attendanceTable)
	if err != nil {
		log.Fatal(err)
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("user_id", int(claims["user_id"].(float64)))
		c.Next()
	}
}

func register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var userID int
	err = db.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		req.Name, req.Email, hashedPassword,
	).Scan(&userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	token, _ := generateToken(userID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
		"user": gin.H{
			"id":    userID,
			"name":  req.Name,
			"email": req.Email,
		},
	})
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	err := db.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email = $1",
		req.Email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !checkPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := generateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func checkIn(c *gin.Context) {
	userID := c.GetInt("user_id")
	today := time.Now().Format("2006-01-02")

	var existingID int
	err := db.QueryRow(
		"SELECT id FROM attendances WHERE user_id = $1 AND date = $2",
		userID, today,
	).Scan(&existingID)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already checked in today"})
		return
	}

	var attendanceID int
	err = db.QueryRow(
		"INSERT INTO attendances (user_id, check_in, date) VALUES ($1, $2, $3) RETURNING id",
		userID, time.Now(), today,
	).Scan(&attendanceID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Check-in successful",
		"id":      attendanceID,
	})
}

func checkOut(c *gin.Context) {
	userID := c.GetInt("user_id")
	today := time.Now().Format("2006-01-02")

	result, err := db.Exec(
		"UPDATE attendances SET check_out = $1 WHERE user_id = $2 AND date = $3 AND check_out IS NULL",
		time.Now(), userID, today,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check out"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No active check-in found or already checked out"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Check-out successful"})
}

func getTodayStatus(c *gin.Context) {
	userID := c.GetInt("user_id")
	today := time.Now().Format("2006-01-02")

	var attendance Attendance
	err := db.QueryRow(
		"SELECT id, user_id, check_in, check_out, date FROM attendances WHERE user_id = $1 AND date = $2",
		userID, today,
	).Scan(&attendance.ID, &attendance.UserID, &attendance.CheckIn, &attendance.CheckOut, &attendance.Date)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, gin.H{"checked_in": false})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"checked_in":  true,
		"checked_out": attendance.CheckOut != nil,
		"attendance":  attendance,
	})
}

func getAttendances(c *gin.Context) {
	userID := c.GetInt("user_id")

	rows, err := db.Query(
		`SELECT a.id, a.user_id, u.name, a.check_in, a.check_out, a.date 
		FROM attendances a 
		JOIN users u ON a.user_id = u.id 
		WHERE a.user_id = $1 
		ORDER BY a.date DESC 
		LIMIT 30`,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendances"})
		return
	}
	defer rows.Close()

	var attendances []Attendance
	for rows.Next() {
		var att Attendance
		if err := rows.Scan(&att.ID, &att.UserID, &att.UserName, &att.CheckIn, &att.CheckOut, &att.Date); err != nil {
			continue
		}
		attendances = append(attendances, att)
	}

	c.JSON(http.StatusOK, attendances)
}

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	router.POST("/api/register", register)
	router.POST("/api/login", login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.POST("/check-in", checkIn)
		protected.POST("/check-out", checkOut)
		protected.GET("/today-status", getTodayStatus)
		protected.GET("/attendances", getAttendances)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	router.Run(":" + port)
}
