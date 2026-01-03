-- Create privilege table
CREATE TABLE IF NOT EXISTS privilege (
    id BIGSERIAL PRIMARY KEY,
    xid VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    exposed BOOLEAN NOT NULL DEFAULT false,
    sort INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by JSONB,
    version BIGINT NOT NULL DEFAULT 1,
    metadata JSONB DEFAULT '{}'
);

-- Create role table
CREATE TABLE IF NOT EXISTS role (
    id SERIAL PRIMARY KEY,
    xid VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    role_type_id INT NOT NULL DEFAULT 1,
    status_id INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by JSONB,
    version BIGINT NOT NULL DEFAULT 1,
    metadata JSONB DEFAULT '{}'
);

-- Create client_auth table
CREATE TABLE IF NOT EXISTS client_auth (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    client_id VARCHAR(255) NOT NULL UNIQUE,
    client_type_id INT NOT NULL,
    options JSONB NOT NULL DEFAULT '{"clientSecret": "", "tokenLifetime": 2592000}',
    status_id INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by JSONB,
    version BIGINT NOT NULL DEFAULT 1,
    metadata JSONB DEFAULT '{}'
);

-- Create role_privilege table (junction table)
CREATE TABLE IF NOT EXISTS role_privilege (
    id BIGSERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    privilege_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by JSONB,
    version BIGINT NOT NULL DEFAULT 1,
    metadata JSONB DEFAULT '{}',
    CONSTRAINT fk_role_privilege_role FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_privilege_privilege FOREIGN KEY (privilege_id) REFERENCES privilege(id) ON DELETE CASCADE,
    CONSTRAINT unique_role_privilege UNIQUE (role_id, privilege_id)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_privilege_xid ON privilege(xid);
CREATE INDEX IF NOT EXISTS idx_privilege_exposed ON privilege(exposed);
CREATE INDEX IF NOT EXISTS idx_role_xid ON role(xid);
CREATE INDEX IF NOT EXISTS idx_role_status ON role(status_id);
CREATE INDEX IF NOT EXISTS idx_client_auth_client_id ON client_auth(client_id);
CREATE INDEX IF NOT EXISTS idx_client_auth_status ON client_auth(status_id);
CREATE INDEX IF NOT EXISTS idx_role_privilege_role_id ON role_privilege(role_id);
CREATE INDEX IF NOT EXISTS idx_role_privilege_privilege_id ON role_privilege(privilege_id);
