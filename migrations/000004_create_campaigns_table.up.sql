CREATE TABLE IF NOT EXISTS campaigns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    branch_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    bonus_type VARCHAR(50),
    bonus_value DECIMAL(5, 2),
    min_purchase DECIMAL(10, 2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::TEXT, now()) NOT NULL,
    FOREIGN KEY (branch_id) REFERENCES branches (id)
);
-- comments
COMMENT ON TABLE campaigns IS 'Campaigns table';
COMMENT ON COLUMN campaigns.id IS 'Campaign ID';
COMMENT ON COLUMN campaigns.branch_id IS 'Branch ID';
COMMENT ON COLUMN campaigns.start_date IS 'Campaign start date';
COMMENT ON COLUMN campaigns.end_date IS 'Campaign end date';
COMMENT ON COLUMN campaigns.bonus_type IS 'Campaign bonus type';
COMMENT ON COLUMN campaigns.bonus_value IS 'Campaign bonus value';
COMMENT ON COLUMN campaigns.min_purchase IS 'Campaign minimum purchase';
COMMENT ON COLUMN campaigns.created_at IS 'Campaign created at';
-- Add index to improve performance of searches by id
CREATE INDEX IF NOT EXISTS idx_campaigns_id ON campaigns (id);
-- Add index to improve performance of searches by branch_id
CREATE INDEX IF NOT EXISTS idx_campaigns_branch_id ON campaigns (branch_id);
