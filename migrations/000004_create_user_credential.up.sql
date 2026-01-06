-- User credential table for flexible authentication
CREATE TABLE IF NOT EXISTS user_credential (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    auth_provider_id INT NOT NULL DEFAULT 1,
    credential_key VARCHAR(255) NOT NULL,
    credential_secret VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(auth_provider_id, credential_key)
);

CREATE INDEX IF NOT EXISTS idx_user_credential_user_id ON user_credential(user_id);
CREATE INDEX IF NOT EXISTS idx_user_credential_key ON user_credential(credential_key);

COMMENT ON COLUMN user_credential.auth_provider_id IS '1=PASSWORD, 2=GOOGLE, 3=FACEBOOK, 4=APPLE';
COMMENT ON COLUMN user_credential.credential_key IS 'email/phone/username for PASSWORD, provider_user_id for OAuth';
COMMENT ON COLUMN user_credential.credential_secret IS 'password_hash for PASSWORD, null for OAuth';
