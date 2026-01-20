DROP TABLE IF EXISTS messages;
DROP INDEX IF EXISTS idx_messages_id;
DROP INDEX IF EXISTS idx_message_chat_id;
DROP INDEX IF EXISTS id_message_created_date;

DROP TABLE IF EXISTS chats;
DROP INDEX IF EXISTS idx_chats_id;
DROP INDEX IF EXISTS idx_chats_user_id_1;
DROP INDEX IF EXISTS idx_chats_user_id_2;
DROP INDEX IF EXISTS idx_chats_created_date;

DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS idx_users_login;
DROP INDEX IF EXISTS idx_users_created_date;

