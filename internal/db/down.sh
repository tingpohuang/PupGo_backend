# migrate -database "$DB_RESOURCE_NAME://$DB_USERNAME:$DB_PASSWORD@$DB_URL"  -path ./migrations -verbose up 1
URL="${DB_RESOURCE_NAME}://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_URL}:${DB_PORT})/pupgo"
echo $URL
migrate -path ./migrations -database ${URL} force down
# migrate -database "mysql://tim:greenfield204@192.168.1.129" version
# ls ./migrations
# create -ext sql -dir db/migration