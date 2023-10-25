package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"
)

// Реализовать HTTP-сервер для работы с календарем. В рамках
// задания необходимо работать строго со стандартной
// HTTP-библиотекой.
// В рамках задания необходимо:
// 1. Реализовать вспомогательные функции для сериализации
// объектов доменной области в JSON.
// 2. Реализовать вспомогательные функции для парсинга и
// валидации параметров методов /create_event и
// /update_event.
// 3. Реализовать HTTP обработчики для каждого из методов API,
// используя вспомогательные функции и объекты доменной
// области.
// 4. Реализовать middleware для логирования запросов
// Методы API:
// ● POST /create_event
// ● POST /update_event
// ● POST /delete_event
// ● GET /events_for_day
// ● GET /events_for_week
// ● GET /events_for_month
// Параметры передаются в виде www-url-form-encoded (т.е.
// обычные user_id=3&date=2019-09-09). В GET методах параметры
// передаются через queryString, в POST через тело запроса.
// В результате каждого запроса должен возвращаться
// JSON-документ содержащий либо {"result": "..."} в случае
// успешного выполнения метода, либо {"error": "..."} в случае
// ошибки бизнес-логики.
// В рамках задачи необходимо:
// 1. Реализовать все методы.
// 2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
// 3. В случае ошибки бизнес-логики сервер должен возвращать
// HTTP 503. В случае ошибки входных данных (невалидный int
// например) сервер должен возвращать HTTP 400. В случае
// остальных ошибок сервер должен возвращать HTTP 500.
// Web-сервер должен запускаться на порту указанном в
// конфиге и выводить в лог каждый обработанный запрос.

var EventStorage = make(map[int64][]*Event)

type Event struct {
	EventId int64     `json:"event_id"`
	Date    time.Time `json:"date"`
	Event   string    `json:"event"`
}

func main() {
	http.Handle("/create_event", Logger(Method("POST")(ContentType(http.HandlerFunc(CreateEventHandler)))))
	http.Handle("/update_event", Logger(Method("UPDATE")(ContentType(http.HandlerFunc(UpdateEventHandler)))))
	http.Handle("/delete_event", Logger(Method("DELETE")(ContentType(http.HandlerFunc(DeleteEventHandler)))))
	http.Handle("/events_for_day", Logger(Method("GET")(http.HandlerFunc(EventsForDayHandler))))
	http.Handle("/events_for_week", Logger(Method("GET")(http.HandlerFunc(EventsForWeekHandler))))
	http.Handle("/events_for_month", Logger(Method("GET")(http.HandlerFunc(EventsForMonthHandler))))

	log.Println(http.ListenAndServe(":8080", nil))
}

// middlewares
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())
		h.ServeHTTP(w, r)
	})
}

func ContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func Method(method string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	response, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})

	writeJSON(w, statusCode, response)
}

func WriteSuccess(w http.ResponseWriter, statusCode int, result interface{}) {
	response, _ := json.Marshal(map[string]interface{}{
		"result": result,
	})

	writeJSON(w, statusCode, response)
}

func writeJSON(w http.ResponseWriter, statusCode int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

type CreateForm struct {
	UserId int64
	Date   time.Time
	Event  string
}

func parseCreateForm(r *http.Request) (CreateForm, error) {
	err := r.ParseForm()
	if err != nil {
		return CreateForm{}, err
	}
	userId, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	if err != nil {
		return CreateForm{}, err
	}
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return CreateForm{}, err
	}
	event := r.FormValue("event")

	return CreateForm{
		UserId: userId,
		Date:   date,
		Event:  event,
	}, nil
}

type UpdateForm struct {
	UserId  int64
	EventId int64
	Date    time.Time
	Event   string
}

func parseUpdateForm(r *http.Request) (UpdateForm, error) {
	err := r.ParseForm()
	if err != nil {
		return UpdateForm{}, err
	}
	eventId, err := strconv.ParseInt(r.FormValue("event_id"), 10, 64)
	if err != nil {
		return UpdateForm{}, err
	}
	userId, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	if err != nil {
		return UpdateForm{}, err
	}
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return UpdateForm{}, err
	}
	event := r.FormValue("event")

	return UpdateForm{
		UserId:  userId,
		EventId: eventId,
		Date:    date,
		Event:   event,
	}, nil
}

type DeleteForm struct {
	UserId  int64
	EventId int64
}

func parseDeleteForm(r *http.Request) (DeleteForm, error) {
	err := r.ParseForm()
	if err != nil {
		return DeleteForm{}, err
	}
	eventId, err := strconv.ParseInt(r.FormValue("event_id"), 10, 64)
	if err != nil {
		return DeleteForm{}, err
	}
	userId, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	if err != nil {
		return DeleteForm{}, err
	}

	return DeleteForm{
		UserId:  userId,
		EventId: eventId,
	}, nil
}

// handlers
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parseCreateForm(r)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	res := CreateEvent(&Event{
		Date:  form.Date,
		Event: form.Event,
	}, form.UserId)

	WriteSuccess(w, http.StatusCreated, res)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parseUpdateForm(r)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	res, err := UpdateEvent(&Event{
		EventId: form.EventId,
		Date:    form.Date,
		Event:   form.Event,
	}, form.UserId)
	if err != nil {
		WriteError(w, 503, err)
		return
	}

	WriteSuccess(w, http.StatusOK, res)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parseDeleteForm(r)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	res, err := DeleteEvent(&Event{
		EventId: form.EventId,
	}, form.UserId)
	if err != nil {
		WriteError(w, 503, err)
		return
	}

	WriteSuccess(w, http.StatusOK, res)
}

func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		WriteError(w, 503, err)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		WriteError(w, 503, err)
		return
	}

	res, err := EventsForDay(userId, date)
	if err != nil {
		WriteError(w, 503, err)
		return
	}
	WriteSuccess(w, 200, res)
}

func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		WriteError(w, 503, err)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		WriteError(w, 503, err)
		return
	}

	res, err := EventsForWeek(userId, date)
	if err != nil {
		WriteError(w, 503, err)
		return
	}
	WriteSuccess(w, 200, res)
}

func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		WriteError(w, 503, err)
		return
	}
	date, err := time.Parse("2006-01", r.URL.Query().Get("date"))
	if err != nil {
		WriteError(w, 503, err)
		return
	}

	res, err := EventsForMonth(userId, date)
	if err != nil {
		WriteError(w, 503, err)
		return
	}
	WriteSuccess(w, 200, res)
}

// logics
func CreateEvent(event *Event, userId int64) string {
	event.EventId = int64(len(EventStorage[userId]) + 1)
	EventStorage[userId] = append(EventStorage[userId], event)
	return fmt.Sprintf("%+v", *event)
}

func UpdateEvent(event *Event, userId int64) (string, error) {
	var found bool
	for _, v := range EventStorage[userId] {
		if v.EventId == event.EventId {
			v.Date = event.Date
			v.Event = event.Event
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("event %d not found", event.EventId)
	}
	return fmt.Sprintf("%+v", *event), nil
}

func DeleteEvent(event *Event, userId int64) (string, error) {
	var found bool
	for i, v := range EventStorage[userId] {
		if v.EventId == event.EventId {
			slices.Delete(EventStorage[userId], i, i+1)
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("event %d not found", event.EventId)
	}
	return fmt.Sprintf("%+v", *event), nil
}

func EventsForDay(userId int64, date time.Time) ([]Event, error) {
	result := []Event{}
	for _, v := range EventStorage[userId] {
		if v.Date == date {
			result = append(result, *v)
		}
	}
	if len(result) == 0 {
		return []Event{}, fmt.Errorf("no events on this date: %s", date)
	}
	return result, nil
}

func EventsForWeek(userId int64, date time.Time) ([]Event, error) {
	result := []Event{}
	for _, v := range EventStorage[userId] {
		y1, w1 := v.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, *v)
		}
	}
	if len(result) == 0 {
		return []Event{}, fmt.Errorf("no events on this week: %s", date)
	}
	return result, nil
}

func EventsForMonth(userId int64, date time.Time) ([]Event, error) {
	result := []Event{}
	for _, v := range EventStorage[userId] {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
			result = append(result, *v)
		}
	}
	if len(result) == 0 {
		return []Event{}, fmt.Errorf("no events on this month: %s", date)
	}
	return result, nil
}
