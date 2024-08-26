include .env

inspect_schema:
	atlas schema inspect --env gorm --url "env://src"

generate_migration:
	atlas migrate diff --env gorm

apply_migrations:
	atlas migrate apply --env gorm -u $(DB_URL)
