CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    mobile_number VARCHAR(20) NOT NULL UNIQUE,
    gender VARCHAR(10) NOT NULL,
    position VARCHAR(50), -- Or CREATE TYPE user_position AS ENUM ('President', 'VP', 'Secretary'); then use user_position
    status VARCHAR(20) NOT NULL DEFAULT 'active', -- e.g., 'active', 'inactive', 'suspended'
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);