-- Create user table (temporary)
CREATE TABLE IF NOT EXISTS "user" (
    id BIGSERIAL PRIMARY KEY,
    xid VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by JSONB,
    version BIGINT NOT NULL DEFAULT 1,
    metadata JSONB DEFAULT '{}'
);

-- Create index for user table
CREATE INDEX IF NOT EXISTS idx_user_xid ON "user"(xid);
