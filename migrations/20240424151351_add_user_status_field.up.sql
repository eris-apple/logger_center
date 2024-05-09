ALTER TABLE users
ADD COLUMN IF NOT EXISTS status varchar null default 'pending';
