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

CREATE INDEX idx_transactions_consumer_id ON transactions (consumer_id);
CREATE INDEX idx_transactions_limit_id ON transactions (limit_id);