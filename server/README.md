# server

Example of slog package usage with server. `middleware.go` sets new context with request fields
(unique id, method and path) on http request. Every handler can use `<Level>Context(...)` method on slog logger, which
will print request id, method and path.

This way each log has extra fields, so they can be grouped by id (unique request), method or path.

## example

- `go build && ./server`
- visit http://localhost:3000

```
{"time":"2023-09-26T12:01:15.259376+01:00","level":"INFO","msg":"starting server on port 3000","component":"server","request":{"method":"","path":"","id":""}}
{"time":"2023-09-26T12:01:27.39356+01:00","level":"INFO","msg":"received request","component":"server","request":{"method":"GET","path":"/","id":"167B32B3-2663-BAA1-1299-31D2CB1D0968"}}
{"time":"2023-09-26T12:01:27.393608+01:00","level":"INFO","msg":"returned response","component":"server","request":{"method":"GET","path":"/","id":"167B32B3-2663-BAA1-1299-31D2CB1D0968"}}
{"time":"2023-09-26T12:01:27.393637+01:00","level":"INFO","msg":"request took 0.11 milliseconds","component":"middleware","request":{"method":"GET","path":"/","id":"167B32B3-2663-BAA1-1299-31D2CB1D0968"}}
{"time":"2023-09-26T12:01:27.413206+01:00","level":"INFO","msg":"received request","component":"server","request":{"method":"GET","path":"/favicon.ico","id":"B524F734-5295-1D4A-37FE-96CCA47CBF1C"}}
{"time":"2023-09-26T12:01:27.413246+01:00","level":"INFO","msg":"returned response","component":"server","request":{"method":"GET","path":"/favicon.ico","id":"B524F734-5295-1D4A-37FE-96CCA47CBF1C"}}
{"time":"2023-09-26T12:01:27.41326+01:00","level":"INFO","msg":"request took 0.07 milliseconds","component":"middleware","request":{"method":"GET","path":"/favicon.ico","id":"B524F734-5295-1D4A-37FE-96CCA47CBF1C"}}
```
