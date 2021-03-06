# Go rest backend

A simple rest backend to serve preconfigured static responses.

## Usage

You need a working Go tool chain. Install Go [https://golang.org/doc/install]

The application can be started with:

	go run server.go

When the application has started, you can test that the already preconfigured routes work by opening any of the links in a browser:

[http:localhost:8080/todo/1]
[http:localhost:8080/todo/2]
[http:localhost:8080/todo/3]

You should see a json response.

## Configuration

You can create your own configuration by modifying the __config.json__

    {
        "route":"/todo/1",
        "httpStatus":200,
        "responseFilePath":"todo1.json",
        "delay":0
    }

1. __route__ defines the url relative to http://localhost:8080
2. __httpStatus__ defines the actual http response status code.
3. __responseFilePath__ defines the path from where to read the response payload.
4. __delay__ can be used to add additional delay to response. The delay is random and __delay__ is the max value for it.

Once you have the routes configured, add the response json files and you are ready to go! :)

## License

Distributed under the Eclipse Public License, the same as Clojure.
