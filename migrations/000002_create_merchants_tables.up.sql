CREATE TABLE IF NOT EXISTS merchants (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL
);
-- comments
COMMENT ON TABLE merchants IS 'Merchants table';
COMMENT ON COLUMN merchants.id IS 'Merchant ID';
COMMENT ON COLUMN merchants.name IS 'Merchant name';
COMMENT ON COLUMN merchants.description IS 'Merchant description';
COMMENT ON COLUMN merchants.created_at IS 'Merchant created at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_merchants_id ON merchants (id);
