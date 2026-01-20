CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL UNIQUE,
    created_date  TIMESTAMP DEFAULT NOW(),
    updated_date  TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_login ON users(login);
CREATE INDEX IF NOT EXISTS idx_users_created_date ON users(created_date DESC);

CREATE TABLE IF NOT EXISTS chats(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id_1 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_id_2 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_date  TIMESTAMP DEFAULT NOW(),

    CONSTRAINT different_users 
        CHECK (user_id_1 != user_id_2)
);

CREATE INDEX IF NOT EXISTS idx_chats_id ON chats(id);
CREATE INDEX IF NOT EXISTS idx_chats_user_id_1 ON chats(user_id_1);
CREATE INDEX IF NOT EXISTS idx_chats_user_id_2 ON chats(user_id_2);
CREATE INDEX IF NOT EXISTS idx_chats_created_date ON chats(created_date DESC);

CREATE TABLE IF NOT EXISTS messages(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id_sender UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_id_recipient UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chat_id UUID NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    created_date TIMESTAMP DEFAULT NOW(),
    is_received	BOOLEAN DEFAULT false,
    text TEXT  NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_messages_id ON messages(id);
CREATE INDEX IF NOT EXISTS idx_message_chat_id ON messages(chat_id);
CREATE INDEX IF NOT EXISTS id_message_created_date ON messages(created_date DESC);