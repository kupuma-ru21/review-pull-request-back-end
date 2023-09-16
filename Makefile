run:
	go run server.go

gql-gen:
	gqlgen generate

sqlboiler:
	sqlboiler --wipe psql
