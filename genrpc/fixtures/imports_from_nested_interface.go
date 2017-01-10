package test

import (
	"github.com/chris-skud/genrpc/genrpc/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
