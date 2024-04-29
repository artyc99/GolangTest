CREATE TABLE IF NOT EXISTS "users"
(
    "id" text NOT NULL,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "age" int NOT NULL,
    "recording_date" int,
    PRIMARY KEY ("id")
    );