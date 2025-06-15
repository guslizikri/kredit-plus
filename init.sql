CREATE TABLE consumers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nik VARCHAR(20) UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100) NOT NULL,
    birth_place VARCHAR(50),
    birth_date DATE,
    salary BIGINT NOT NULL,
    photo_ktp TEXT NOT NULL,
    photo_selfie TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    consumer_id UUID REFERENCES consumers(id),
    limit_id UUID REFERENCES limits(id),
    contract_number VARCHAR(50) UNIQUE NOT NULL,
    otr_price BIGINT NOT NULL,
    admin_fee BIGINT NOT NULL,
    installment BIGINT NOT NULL,
    interest BIGINT NOT NULL,
    asset_name VARCHAR(100) NOT NULL,
    source_channel VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_limits_consumer_id ON limits (consumer_id);
CREATE INDEX idx_transactions_consumer_id ON transactions (consumer_id);
CREATE INDEX idx_transactions_limit_id ON transactions (limit_id);