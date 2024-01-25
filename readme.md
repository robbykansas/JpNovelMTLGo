go test -cover -coverpkg=./internal/controller ./test -coverprofile=c.out

go tool cover -html="c.out"

mockery --dir=./internal/repository --name=TranslateRepository