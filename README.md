# Backend

# Tech Stack
- [sqlc](https://github.com/kyleconroy/sqlc) for generating golang functions that match postgres database schema
- [gin](https://github.com/gin-gonic/gin) for http server
- [viper](https://github.com/spf13/viper) for loading config / environment variables

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
- Need to implement Mock Db for testing HTTP API in GO and achieve 100% coverage
- Need to add additional tests for api module...
- Middleware need testing
- Transfers API doesn't have authorization yet