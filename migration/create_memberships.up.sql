CREATE TABLE memberships (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chatroom_id INT NOT NULL REFERENCES chatrooms(id) ON DELETE CASCADE,
    role TEXT DEFAULT 'member',
    joined_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (user_id, chatroom_id)
);
