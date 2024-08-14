build_books:
	go build cmd/books_api_server/books_api_server.go

build_docker_books:
	docker build -t books_api_server:latest --build-arg="WHICH_SERVER=books_api_server" .

run_docker_books:
	docker run \
		--publish=3000:3000 \
		books_api_server:latest

build_users:
	go build cmd/users_api_server/users_api_server.go

build_docker_users:
	docker build -t users_api_server:latest --build-arg="WHICH_SERVER=users_api_server" .

run_docker_users:
	docker run \
		--publish=3000:3000 \
		users_api_server:latest
