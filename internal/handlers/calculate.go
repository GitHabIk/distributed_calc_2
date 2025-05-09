package handlers

import (
	"distributed-calculator/internal/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Knetic/govaluate" // Импортируем пакет для вычислений
)

type calcRequest struct {
	Expression string `json:"expression"`
}

type calcResponse struct {
	Result string `json:"result"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем userID из заголовка
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Парсим тело запроса
	var req calcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Выполняем вычисления с использованием пакета govaluate
	expression, err := govaluate.NewEvaluableExpression(req.Expression)
	if err != nil {
		http.Error(w, "Invalid expression", http.StatusBadRequest)
		return
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		http.Error(w, "Error evaluating expression", http.StatusInternalServerError)
		return
	}

	// Преобразуем результат в строку
	resultStr := result.(float64)

	// Сохраняем результат в базе данных
	if err := db.SaveCalculation(userID, req.Expression, strconv.FormatFloat(resultStr, 'f', -1, 64)); err != nil {
		http.Error(w, "Failed to save calculation", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	json.NewEncoder(w).Encode(calcResponse{Result: strconv.FormatFloat(resultStr, 'f', -1, 64)})
}
