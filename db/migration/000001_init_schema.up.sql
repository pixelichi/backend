-- Users -----------------------------------------------------------------------------------
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE plants (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  species VARCHAR(255),
  date_planted TIMESTAMP WITH TIME ZONE,
  last_watered TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_plants_user_id ON plants(user_id);

CREATE TABLE plant_photos (
  id BIGSERIAL PRIMARY KEY,
  plant_id BIGINT REFERENCES plants(id) ON DELETE CASCADE,
  photo_object_key VARCHAR(255) NOT NULL,
  is_main BOOLEAN DEFAULT FALSE,
  date_taken TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_plant_photos_plant_id ON plant_photos(plant_id);


-- CREATE INDEX ON "entries" ("account_id");

-- CREATE INDEX ON "transfers" ("from_account_id");

-- CREATE INDEX ON "transfers" ("to_account_id");

-- CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

-- COMMENT ON COLUMN "entries"."amount" IS 'can be neg or pos';

-- COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

-- ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");