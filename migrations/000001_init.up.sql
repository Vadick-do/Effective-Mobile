CREATE SCHEMA efmobapp;

CREATE TABLE efmobapp.subscriptions (
    id UUID PRIMARY KEY,
    service_name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date VARCHAR(7) NOT NULL,
    end_date VARCHAR(7),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);