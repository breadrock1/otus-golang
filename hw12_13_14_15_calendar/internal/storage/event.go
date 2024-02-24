package storage

import (
	"github.com/google/uuid"
	"time"
)

var eventsQueue map[string]*Event

// Уведомление - временная сущность, в БД не хранится, складывается в очередь для рассыльщика, содержит поля:
//
//	ID события;
//	Заголовок события;
//	Дата события;
//	Пользователь, которому отправлять.
type Event struct {
	ID       string
	Title    string
	Datetime time.Time
	User     int
}

//Создать (событие);
//Обновить (ID события, событие);
//Удалить (ID события);
//СписокСобытийНаДень (дата);
//СписокСобытийНаНеделю (дата начала недели);
//СписокСобытийНaМесяц (дата начала месяца).

func Create(title string, dateTime time.Time, user int) {
	uuidVal := uuid.New().String()
	eventsQueue[uuidVal] = &Event{
		ID:       uuidVal,
		Title:    title,
		Datetime: dateTime,
		User:     user,
	}
}

func Update(id string, event *Event) {
	eventsQueue[id] = event
}

func Delete(id string) {
	delete(eventsQueue, id)
}
