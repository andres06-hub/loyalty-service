CREATE TABLE IF NOT EXISTS branches (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  merchant_id UUID,
  name VARCHAR(255),
  location VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL,
  FOREIGN KEY (merchant_id) REFERENCES merchants (id)
);
-- comments
COMMENT ON TABLE branches IS 'Branches table';
COMMENT ON COLUMN branches.id IS 'Branch ID';
COMMENT ON COLUMN branches.merchant_id IS 'Merchant ID';
COMMENT ON COLUMN branches.name IS 'Branch name';
COMMENT ON COLUMN branches.location IS 'Branch location';
COMMENT ON COLUMN branches.created_at IS 'Branch created at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_branches_id ON branches (id);
-- Add index to improve performance of searches by merchant_id
CREATE INDEX IF NOT EXISTS idx_branches_merchant_id ON branches (merchant_id);