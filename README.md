# go-urlbuilder
A small library for safely building URLs

## Motivation

URLs, similar to SQL queries, are vulnerable to injection if unsafe data
is used in string interpolation. This is especially dangerous in microservice
architecture where security data may be passed in parameters.

In the following (contrived) example a service receives a request for a document, determines whether the user is an 
admin from the request context, then makes a request to a domain service.

```go
func Handle(r http.Request, w http.Response) error {
	documentID := r.URL.Query().Get("documentId")
	
	isAdmin := isAdmin(r.Context())
	
	http.Get(fmt.Sprintf("https://domainservice/document/%s?isAdmin=%v", documentID, isAdmin)
	
	// ...
}
``` 

An attacker could abuse this in an attack class known as 
[Server-Side Request Forgery (SSRF)](https://owasp.org/www-community/attacks/Server_Side_Request_Forgery). They would
construct a request something like this:

```
http://app.example.com/document/1234%3FisAdmin%3Dtrue%23
```

The document ID is decoded and `documentID` becomes `1234?isAdmin=true#` meaning the URL to the backend service called
is:
``` 
https://domainservice/document/1234?isAdmin=true#?isAdmin=false
```
The `#` fragment means that the `isAdmin=false` is ignored.

By using strong typing for URL construction we can avoid this, however the Go standard library does not provide fluent
ways to handle path variables.

Using this library, the above code could be rewritten like so

```go
func Handle(r http.Request, w http.Response) error {
	documentID := r.URL.Query().Get("documentId")

	isAdmin := isAdmin(r.Context())
	
	http.Get(urlbuilder.New().HTTPS().WithHost("domainservice").
		MustWithPathWithParameters("/document/?", documentID).
		MustWithQuery("isAdmin", strconv.FormatBool(isAdmin)).String())

	// ...
}
```

## Usage

For examples of usage, see the [urlbuilder_test.go](urlbuilder_test.go) file

### URL templates

Each call to the builder is idempotent (does not modify the underlying object), so you can create templates
to reduce repeated code. For example:

```go
urlBase := urlbuilder.New().HTTPS().WithHost("example.org")

docURL := urlBase.MustWithPathWithParameters("/document/?", documentID)
userURL := urlBase.MustWithPathWithParameters("/user/?", userID)
```

