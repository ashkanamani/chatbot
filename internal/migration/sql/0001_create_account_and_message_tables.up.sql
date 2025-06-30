CREATE TABLE IF NOT EXISTS accounts (
    id BIGINT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT,
    username TEXT,
    phone_number TEXT,
    joined_at TIMESTAMPTZ NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    blocked BOOLEAN DEFAULT FALSE,
    link_token TEXT
);
CREATE TABLE IF NOT EXISTS messages (
    message_id BIGINT PRIMARY KEY,
    from_user_id BIGINT NOT NULL,
    to_user_id BIGINT NOT NULL,
    text TEXT NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    delivered BOOLEAN DEFAULT FALSE,
    reply_to_message_id BIGINT,
    CONSTRAINT fk_from_user FOREIGN KEY (from_user_id) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_to_user FOREIGN KEY (to_user_id) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_reply_to FOREIGN KEY (reply_to_message_id) REFERENCES messages(message_id) ON DELETE SET NULL
);