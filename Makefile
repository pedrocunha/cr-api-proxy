start:
	go run server.go

push:
	heroku container:push web

release:
	heroku container:release web
