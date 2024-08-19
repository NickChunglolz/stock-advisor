alias air='$(go env GOPATH)/bin/air'
alias templ='$(go env GOPATH)/bin/templ'
templ generate
go build -o tmp/main main/main.go
air