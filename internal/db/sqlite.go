package db

import (
	"database/sql"

	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(path string) error {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	schema := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        login TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS calculations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        expression TEXT NOT NULL,
        result TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`
	_, err = DB.Exec(schema)
	return err
}

func CreateUser(login, hash string) error {
	_, err := DB.Exec("INSERT INTO users(login, password_hash) VALUES(?, ?)", login, hash)
	return err
}

func GetUserByLogin(login string) (int, string, error) {
	row := DB.QueryRow("SELECT id, password_hash FROM users WHERE login = ?", login)
	var id int
	var hash string
	err := row.Scan(&id, &hash)
	return id, hash, err
}

func SaveCalculation(userID int, expression, result string) error {
	_, err := DB.Exec("INSERT INTO calculations(user_id, expression, result) VALUES (?, ?, ?)", userID, expression, result)
	return err
}

func GetUserIDFromToken(tokenString string) (string, error) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", errors.New("Invalid token format")
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Пытаемся декодировать токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Здесь должен быть секретный ключ для верификации токена
		return []byte("secret_key"), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("Invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("User ID not found in token")
	}

	return userID, nil
}

// Эта функция будет сохранять задачу в базе данных
func SaveCalculationTask(userID string, expression string) (string, error) {
	// Реализуйте сохранение в базу данных и возвращение taskID
	taskID := "some_generated_task_id"
	return taskID, nil
}

// Эта функция будет обновлять статус задачи
func UpdateTaskStatus(taskID, status, result string) error {
	// Реализуйте обновление статуса задачи в базе данных
	if taskID == "" {
		return errors.New("taskID cannot be empty")
	}
	return nil
}
