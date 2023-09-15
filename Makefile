run:
	go run server.go

gql-gen:
	go run github.com/99designs/gqlgen generate

sqlboiler:
	sqlboiler --wipe psql
