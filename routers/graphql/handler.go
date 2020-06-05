package graphql

import (
	//"code.gitea.io/gitea/modules/context"
	//gocontext "context"
	//"net/http"

	//"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/routers/graphql/generated"
	//gocontext "context"
	"gitea.com/macaron/macaron"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler"
)



func RegisterGraphql(m *macaron.Macaron) {

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))
//	handler_ := func (apiCtx *context.APIContext, w http.ResponseWriter, r *http.Request)  {
//
//		/*ctx := r.Context()
//		wrappedCtx := gocontext.WithValue(ctx, "api-context", apiCtx)
//		r = r.WithContext(wrappedCtx)
//		resolver.Exec(wrappedCtx)(wrappedCtx)
//		*/
//		server.ServeHTTP(w, r)
//	}

	m.Get ("/graphql", playground.Handler("GraphQL playground", "/graphql").ServeHTTP)
	m.Post("/graphql", server.ServeHTTP)
}
