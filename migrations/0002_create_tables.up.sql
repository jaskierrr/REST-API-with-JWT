BEGIN;

CREATE INDEX IF NOT EXISTS idx_user_id ON users(id);
CREATE INDEX IF NOT EXISTS idx_tasks_user_id ON tasks(user_id);
CREATE INDEX IF NOT EXISTS idx_referrers_user_id ON referrers(user_id);

COMMIT;
