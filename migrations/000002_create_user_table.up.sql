-- Create user table
CREATE TABLE IF NOT EXISTS "user" (
    id BIGSERIAL PRIMARY KEY,
    xid VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(100) UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    age VARCHAR(10),
    avatar VARCHAR(255),
    status_id INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by JSONB,
    version BIGINT NOT NULL DEFAULT 1,
    metadata JSONB DEFAULT '{}'
);

-- Create indexes for user table
CREATE INDEX IF NOT EXISTS idx_user_xid ON "user"(xid);
CREATE INDEX IF NOT EXISTS idx_user_email ON "user"(email);
CREATE INDEX IF NOT EXISTS idx_user_phone ON "user"(phone);
CREATE INDEX IF NOT EXISTS idx_user_username ON "user"(username);

