CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id int NOT NULL,
    content text NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now(),
    updated_at timestamp(0) with time zone DEFAULT now()
)
