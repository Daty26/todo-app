CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    phone_number VARCHAR(15)  CHECK (
            phone_number ~ '^\+[0-9]+$'
            AND
            char_length(phone_number) between  10 and 15
        )
);


CREATE TABLE todoapp.tasks(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title varchar(100) NOT NULL check ( char_length(title) between 1 and 100),
    description varchar(1000) check ( char_length(description) between 1 and 1000),
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL ,
    completed_at TIMESTAMPTZ,

    user_id INTEGER NOT NULL REFERENCES todoapp.users(id)

    check (
        (completed = false and completed_at is null)
        or
        (completed = true and completed_at is not null and completed_at >= created_at)
    )

);

