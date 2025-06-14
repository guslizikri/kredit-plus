CREATE TABLE limits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    consumer_id UUID REFERENCES consumers(id),
    tenor_month INT NOT NULL CHECK (tenor_month IN (1, 2, 3, 4)),
    limit_amount BIGINT NOT NULL,
    used_amount BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (consumer_id, tenor_month)
);

CREATE INDEX idx_limits_consumer_id ON limits (consumer_id);