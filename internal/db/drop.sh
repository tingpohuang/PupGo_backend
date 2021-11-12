# migrate -database "$DB_RESOURCE_NAME://$DB_USERNAME:$DB_PASSWORD@$DB_URL"  -path ./migrations -verbose up 1
URL="${DB_RESOURCE_NAME}://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_URL}:${DB_PORT})/pupgo"
migrate -database=${URL} drop
