FROM golang:1.13.7-alpine3.11 as build
COPY . /app
WORKDIR /app
RUN mkdir -p bin

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -v -o bin/server .

### Put the binary onto Heroku image
FROM heroku/heroku:16
COPY --from=build /app /app
CMD ["/app/bin/server"]
