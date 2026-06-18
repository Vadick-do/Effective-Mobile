CREATE SCHEMA efmobapp;

CREATE TABLE efmobapp.subscriptions (
    id UUID PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    service_name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);