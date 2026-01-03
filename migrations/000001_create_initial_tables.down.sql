-- Drop tables in reverse order to respect foreign key constraints
DROP TABLE IF EXISTS role_privilege;
DROP TABLE IF EXISTS client_auth;
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS privilege;
