CREATE INDEX idx_messages_chatroom_id
ON messages(chatroom_id, created_at DESC);

CREATE INDEX idx_memberships_user_id
ON memberships(user_id);

CREATE INDEX idx_memberships_chatroom_id
ON memberships(chatroom_id);
