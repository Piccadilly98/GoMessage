package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Piccadilly98/GoMessage/internal/domain"
	"github.com/Piccadilly98/GoMessage/internal/service"
)

type HTTPHandler struct {
	Service service.Service
}

func NewHTTPHandler(db *sql.DB, service service.Service) *HTTPHandler {
	return &HTTPHandler{
		Service: service,
	}
}

// Регистрация пользователя**
// **Эндпоинт:** `POST /api/v1/auth/register`
// **Тело запроса:**

func (h *HTTPHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var userLogin, userPassword string
	var regUserDto domain.RegistrationUserDomain
	if err := json.NewDecoder(r.Body).Decode(&regUserDto); err != nil {

	}

}

// Аутентификация пользователя**
// **Эндпоинт:** `POST /api/v1/auth/login`
// **Тело запроса:**

func (h *HTTPHandler) AuthUser(w http.ResponseWriter, r *http.Request) {}

//Получение списка чатов**

// **Эндпоинт:** `GET /api/v1/chats`
// **Заголовки:** `Authorization: Bearer <token>`
// **Параметры запроса (опционально):**

// - `limit`: число (default: 20, max: 100)
// - `offset`: число (default: 0)
// - `unread_only`: boolean (default: false)

func (h *HTTPHandler) GetChats(w http.ResponseWriter, r *http.Request) {}

// Создание личного чата (1:1)**

// **Эндпоинт:** `POST /api/v1/chats`
// **Заголовки:** `Authorization: Bearer <token>`
// **Тело запроса:**

// ```json
// {
//   "partner_id": "uuid-v4 (ID пользователя-собеседника)"
// }
// ```

func (h *HTTPHandler) CreatePrivateChat(w http.ResponseWriter, r *http.Request) {}

// Модуль обмена сообщениями**

// #### **3.3.1. Отправка сообщения**
// **WebSocket сообщение (клиент → сервер):**

func (h *HTTPHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {}

// Получение истории сообщений**

// **Эндпоинт:** `GET /api/v1/chats/{chat_id}/messages`
// **Заголовки:** `Authorization: Bearer <token>`

// **Параметры запроса:**
// - `limit`: число (default: 50, max: 200)
// - `before_id`: uuid (сообщение, ранее которого загружать)
// - `after_id`: uuid (сообщение, позже которого загружать)

func (h *HTTPHandler) GetMessageHistory(w http.ResponseWriter, r *http.Request) {}
