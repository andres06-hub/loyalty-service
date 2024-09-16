CREATE TABLE IF NOT EXISTS redemption_transactions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  branch_id UUID,
  reward_type VARCHAR(50),
  reward_value DECIMAL(10, 2),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (branch_id) REFERENCES branches (id)
);
-- comments
COMMENT ON TABLE redemption_transactions IS 'Redemption transaction table';
COMMENT ON COLUMN redemption_transactions.id IS 'Redemption transaction ID';
COMMENT ON COLUMN redemption_transactions.user_id IS 'User ID';
COMMENT ON COLUMN redemption_transactions.branch_id IS 'Branch ID';
COMMENT ON COLUMN redemption_transactions.reward_type IS 'Reward type';
COMMENT ON COLUMN redemption_transactions.reward_value IS 'Reward value';
COMMENT ON COLUMN redemption_transactions.created_at IS 'Redemption transaction created at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_redemption_transactions_id ON redemption_transactions (id);
-- Add index to improve performance of searches by user_id
CREATE INDEX IF NOT EXISTS idx_redemption_transactions_user_id ON redemption_transactions (user_id);
-- Add index to improve performance of searches by branch_id
CREATE INDEX IF NOT EXISTS idx_redemption_transactions_branch_id ON redemption_transactions (branch_id);
