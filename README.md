# cr-api-proxy

This is a simple Go server that proxies requests to `https://api.clashroyale.com`. Its Dockerfile is currently configured to be deployed on Heroku. The following environment variables can be set:

- `PORT` - Port to run the server
- `TOKEN`- ClashRoyale API generated token
- `FIXIE_URL` - To proxy requests through Fixie (Heroku plugin - not needed if you have a static IP)
- `PASSWORD` - Basic Auth required to use the proxy


Example how to start the server locally:
```
PORT=8080 TOKEN=XXXX go run server.go
```

Then you can do something like:
```
curl -v "127.0.0.1:8080/v1/players/%232GJ82VQ99/upcomingchests"
```
