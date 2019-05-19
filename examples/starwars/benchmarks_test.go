package starwars

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/handler"
	"github.com/samlitowitz/graphqlc-gen-gqlgen/examples/starwars/generated"
)

func BenchmarkSimpleQueryNoArgs(b *testing.B) {
	server := handler.GraphQL(generated.NewExecutableSchema(NewResolver()))

	q := `{"query":"{ search(text:\"Luke\") { ... on Human { starships { name } } } }"}`

	var body strings.Reader
	r := httptest.NewRequest("POST", "/graphql", &body)

	b.ReportAllocs()
	b.ResetTimer()

	rec := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		body.Reset(q)
		rec.Body.Reset()
		server.ServeHTTP(rec, r)
		if rec.Body.String() != `{"data":{"search":[{"starships":[{"name":"X-Wing"},{"name":"Imperial shuttle"}]}]}}` {
			b.Fatalf("Unexpected response")
		}

	}
}
