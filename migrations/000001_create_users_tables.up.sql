CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255),
  identification_number VARCHAR(20) UNIQUE,
  email VARCHAR(255) UNIQUE,
  phone VARCHAR(20) DEFAULT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
-- comments
COMMENT ON TABLE users IS 'Users table';
COMMENT ON COLUMN users.id IS 'User ID';
COMMENT ON COLUMN users.name IS 'User name';
COMMENT ON COLUMN users.identification_number IS 'User identification number';
COMMENT ON COLUMN users.email IS 'User email';
COMMENT ON COLUMN users.phone IS 'User phone number';
COMMENT ON COLUMN users.created_at IS 'User created at';
COMMENT ON COLUMN users.updated_at IS 'User updated at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_users_id ON users (id);
-- Add index to improve performance of searches by identification_number
CREATE INDEX IF NOT EXISTS idx_users_identification_number ON users (identification_number);
-- Add index to improve performance of searches by email
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);