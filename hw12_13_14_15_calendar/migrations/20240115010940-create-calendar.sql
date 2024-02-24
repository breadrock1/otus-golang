-- +goose up
-- Событие - основная сущность, содержит в себе поля:
--     ID - уникальный идентификатор события (можно воспользоваться UUID);
--     Заголовок - короткий текст;
--     Дата и время события;
--     Длительность события (или дата и время окончания);
--     Описание события - длинный текст, опционально;
--     ID пользователя, владельца события;
--     За сколько времени высылать уведомление, опционально.
--     Уведомление
CREATE TABLE IF NOT EXISTS event_calendar (
    id              SERIAL PRIMARY KEY,
    title           TEXT NOT NULL,
    datetime        TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    during_to       TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    description     TEXT NOT NULL,
    user_id         INTEGER NOT NULL,
    notification    INTEGER NOT NULL,
);

-- +goose down
DROP TABLE event_calendar;
