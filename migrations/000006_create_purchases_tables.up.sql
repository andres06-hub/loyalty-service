CREATE TABLE IF NOT EXISTS purchases (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  branch_id UUID,
  purchase_amount DECIMAL(10, 2),
  reward_earned DECIMAL(10, 2),
  reward_type VARCHAR(50), -- 'points' or 'cashback'
  campaign_id UUID DEFAULT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (branch_id) REFERENCES branches (id),
  FOREIGN KEY (campaign_id) REFERENCES campaigns (id)
);
-- comments
COMMENT ON TABLE purchases IS 'Purchases table';
COMMENT ON COLUMN purchases.id IS 'Purchase ID';
COMMENT ON COLUMN purchases.user_id IS 'User ID';
COMMENT ON COLUMN purchases.branch_id IS 'Branch ID';
COMMENT ON COLUMN purchases.purchase_amount IS 'Purchase amount';
COMMENT ON COLUMN purchases.reward_earned IS 'Reward earned';
COMMENT ON COLUMN purchases.reward_type IS 'Reward type';
COMMENT ON COLUMN purchases.campaign_id IS 'Campaign ID';
COMMENT ON COLUMN purchases.created_at IS 'Purchase created at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_purchases_id ON purchases (id);
-- Add index to improve performance of searches by user_id
CREATE INDEX IF NOT EXISTS idx_purchases_user_id ON purchases (user_id);
-- Add index to improve performance of searches by branch_id
CREATE INDEX IF NOT EXISTS idx_purchases_branch_id ON purchases (branch_id);
-- Add index to improve performance of searches by campaign_id
CREATE INDEX IF NOT EXISTS idx_purchases_campaign_id ON purchases (campaign_id);

