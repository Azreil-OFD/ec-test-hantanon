-- Active: 1731967266078@@195.133.50.12@5432@default_db
CREATE TABLE IF NOT EXISTS friends (
    user_id UUID NOT NULL,          -- ID пользователя, который отправил запрос
    friend_id UUID NOT NULL,        -- ID друга, который получил запрос
    status VARCHAR(20) NOT NULL,    -- Статус запроса: "pending", "accepted", "declined"
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Время создания запроса
    PRIMARY KEY (user_id, friend_id),  -- Уникальность пары user_id и friend_id
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,  -- Если пользователь удален, все его связи в таблице удаляются
    FOREIGN KEY (friend_id) REFERENCES users(id) ON DELETE CASCADE  -- Если друг удален, все его связи в таблице удаляются
);