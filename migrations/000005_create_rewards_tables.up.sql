CREATE TABLE IF NOT EXISTS rewards (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  branch_id UUID DEFAULT NULL,
  reward_type VARCHAR(50), -- 'points' or 'cashback'
  reward_value DECIMAL(10, 2),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (branch_id) REFERENCES branches (id)
);
-- comments
COMMENT ON TABLE rewards IS 'Rewards table';
COMMENT ON COLUMN rewards.id IS 'Reward ID';
COMMENT ON COLUMN rewards.user_id IS 'User ID';
COMMENT ON COLUMN rewards.branch_id IS 'Branch ID';
COMMENT ON COLUMN rewards.reward_type IS 'Reward type';
COMMENT ON COLUMN rewards.reward_value IS 'Reward value';
COMMENT ON COLUMN rewards.created_at IS 'Reward created at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_rewards_id ON rewards (id);
-- Add index to improve performance of searches by user_id
CREATE INDEX IF NOT EXISTS idx_rewards_user_id ON rewards (user_id);
-- Add index to improve performance of searches by branch_id
CREATE INDEX IF NOT EXISTS idx_rewards_branch_id ON rewards (branch_id);
