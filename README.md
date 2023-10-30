# Backend

# Setup

Get yourself GVM to manage your Golang installs
https://github.com/moovweb/gvm

This project needs Golang 19 so you can do 
```bash
gvm install go1.19
gmv use go1.19
```

To run the backend + db + minio locally on docker you can just use 
```bash
cp ./env/local.env .env
make local
```

The newly booted DB will not have the right schema so you then have to apply the migration to it
```
```

# Reference

https://github.com/techschool/simplebank

# Tech Stack

- [sqlc](https://github.com/kyleconroy/sqlc) for generating golang functions that match postgres database schema
- [gin](https://github.com/gin-gonic/gin) for http server
- [viper](https://github.com/spf13/viper) for loading config / environment variables
- [minio-go](https://github.com/minio/minio-go)

# Tooling

- [DB Diagram.io](https://dbdiagram.io/home) Used for creating SQL schema commands which sqlc can absorb

# Useful PSQL commands

To get information about a table:

```sql
SELECT column_name, data_type, character_maximum_length, column_default, is_nullable
FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '<name of table>';
```

# Validator Package for JSON

https://github.com/go-playground/validator

# TODO

- Check out how to create typescript classes from golang directly - https://github.com/tkrajina/typescriptify-golang-structs
- Need to implement Mock Db for testing HTTP API in GO and achieve 100% coverage
- Need to add additional tests for api module...
- Middleware need testing
- Transfers API doesn't have authorization yet
- Figure out CSRF along with access_token saved in cookie


CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "plants" (
  "id" BIGSERIAL PRIMARY KEY,
  -- "plant_name" VARCHAR(255) NOT NULL,
  -- "species" VARCHAR(255),
  -- date_planted TIMESTAMP WITH TIME ZONE,
  -- last_watered TIMESTAMP WITH TIME ZONE,
  -- created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- updated_at TIMESTAMP WITH TIME ZONE
);

-- ALTER TABLE "plants" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


-- CREATE TABLE "plants" (
--   "id" BIGSERIAL PRIMARY KEY,
--   "user_id" BIGINT REFERENCES users(id) ON DELETE CASCADE,
--   "plant_name" VARCHAR(255) NOT NULL,
--   "species" VARCHAR(255),
--   -- date_planted TIMESTAMP WITH TIME ZONE,
--   -- last_watered TIMESTAMP WITH TIME ZONE,
--   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
--   -- updated_at TIMESTAMP WITH TIME ZONE
-- );

-- CREATE INDEX idx_plants_user_id ON plants(user_id);

-- CREATE TABLE plant_photos (
--   id BIGSERIAL PRIMARY KEY,
--   plant_id BIGINT REFERENCES plants(id) ON DELETE CASCADE,
--   photo_object_key VARCHAR(255) NOT NULL,
--   is_main BOOLEAN DEFAULT FALSE,
--   date_taken TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
--   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
--   updated_at TIMESTAMP WITH TIME ZONE
-- );

-- CREATE INDEX idx_plant_photos_plant_id ON plant_photos(plant_id);


-- CREATE INDEX ON "entries" ("account_id");

-- CREATE INDEX ON "transfers" ("from_account_id");

-- CREATE INDEX ON "transfers" ("to_account_id");

-- CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

-- COMMENT ON COLUMN "entries"."amount" IS 'can be neg or pos';

-- COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

-- ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");