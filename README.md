
migration:
Install Atlas from macOS or Linux by running:
curl -sSf https://atlasgo.sh | sh

run gen migration:
atlas migrate diff --env gorm

run migrate:
atlas schema apply --env gorm -u "postgres://dev:fastx123@localhost:5432/fastx?sslmode=disable"