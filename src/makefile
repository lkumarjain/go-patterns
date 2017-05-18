
TARGETPATH = "$(GOPATH)/bin/$(TARGET)"


localbuild:
	go build -o $(TARGETPATH)

test:
	go test -coverprofile ./service/cover.out -covermode=count ./service
	go test -coverprofile ./dal/cover.out -covermode=count ./dal

cover:
	echo "mode: count" > cover-all.out
	tail -n +2 ./service/cover.out >> cover-all.out
	tail -n +2 ./dal/cover.out >> cover-all.out
	gocover-cobertura < cover-all.out > cover-cobertura.xml

code-quality:
	gometalinter --vendor --tests --skip=mock --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml

mockgen:
	mockgen -package mock github.com/lkumarjain/go-patterns/src/model Service,Dal > model/mock/mocks.go
