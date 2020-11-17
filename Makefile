.new-migration
# migrate create -ext=sql -dir=migrations/ -seq roles

.execute-migration
# migrate -database postgres://localhost/arco?sslmode=disable -path migrations/ up