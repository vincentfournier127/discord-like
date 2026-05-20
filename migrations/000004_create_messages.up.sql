CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    chatroom_id INT NOT NULL REFERENCES chatrooms(id) ON DELETE CASCADE,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
