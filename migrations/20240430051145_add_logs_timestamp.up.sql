ALTER TABLE logs
ADD COLUMN IF NOT EXISTS timestamp integer null;
