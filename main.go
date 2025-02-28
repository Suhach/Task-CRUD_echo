package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ---------PATCH----HANDLER------------------------\\
func patchNote(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр ID из URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Преобразуем ID в число
	noteID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}
	// Проверяем, существует ли заметка с таким ID
	var note Note
	result := DB.First(&note, noteID)
	if result.Error != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}
	// Декодируем JSON-тело запроса в map для частичного обновления
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}
	//Готовим поля для обновления: разрешаем обновление только "note" и "is_done"
	updates := make(map[string]interface{})
	if val, exists := payload["note"]; exists {
		updates["note"] = val
	}
	if val, exists := payload["is_done"]; exists {
		updates["is_done"] = val
	}
	// Если нет полей для обновления, возвращаем ошибку
	if len(updates) == 0 {
		http.Error(w, "Нет полей для обновления", http.StatusBadRequest)
		return
	}
	// Обновляем запись в базе данных
	if result := DB.Model(&note).Updates(updates); result.Error != nil {
		http.Error(w, "Ошибка при обновлении заметки", http.StatusInternalServerError)
		return
	}
	// Отправляем обновлённую заметку в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// -----------DELETE-----Handler----------------\\
func deleteNote(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр ID из URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}
	// Преобразуем ID в число
	noteID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}
	// Проверяем, существует ли заметка с таким ID
	var note Note
	result := DB.First(&note, noteID)
	if result.Error != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}
	result = DB.Delete(&note)
	if result.Error != nil {
		http.Error(w, "Ошибка при удалении заметки", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Заметка удалена!"})
}

// ----------GET----Handler--------------------\\
func seeNote(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	result := DB.Find(&notes)
	if result.Error != nil {
		http.Error(w, "Ошибка при получении задач.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getNoteByID(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметры URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Преобразуем ID в целое число
	noteID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	// Ищем запись по ID
	var note Note
	result := DB.First(&note, noteID)
	if result.Error != nil {
		http.Error(w, "Заметка не найдена", http.StatusNotFound)
		return
	}

	// Устанавливаем заголовок и отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// ----------POST---Handler--------------------\\
func setNote(w http.ResponseWriter, r *http.Request) {
	var req Note
	//----JSON--decode--form--Body--\\
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// -------------Обработка ошибки--\\
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	//------save--note----in global--var--------\\
	DB.Create(&req)
	//--------wH().Set("C-T","a/j")--------\\
	w.Header().Set("Content-Type", "application/json")
	//--------Status---200----\\\\\\\\
	w.WriteHeader(http.StatusCreated)
	//----------send--update------NE(w).E(map[string]string{"":""})
	json.NewEncoder(w).Encode(req)
}

// -----------Main-----FUNC---------------------\\
func main() {
	InitDB()

	DB.AutoMigrate(&Note{})

	router := mux.NewRouter()
	router.HandleFunc("/api/note/{id}", getNoteByID).Methods("GET")
	router.HandleFunc("/api/note", seeNote).Methods("GET")
	router.HandleFunc("/api/note", setNote).Methods("POST")
	router.HandleFunc("/api/note/{id}", deleteNote).Methods("DELETE")
	router.HandleFunc("/api/note/{id}", patchNote).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}
