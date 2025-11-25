migration_up:
	 migrate -path ./migrations -database "postgresql://postgres:password@localhost:5434/mydb?sslmode=disable" -verbose up
