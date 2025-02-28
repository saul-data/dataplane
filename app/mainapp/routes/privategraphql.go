package routes

import (
	"net/http"

	generated "github.com/saul-data/dataplane/app/mainapp/graphql/private"
	privategraphql "github.com/saul-data/dataplane/app/mainapp/graphql/private/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

/* GraphQL Handlers */
func PrivateGraphqlHandler() fiber.Handler {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &privategraphql.Resolver{}}))

	return phttpHandler(h)
}

func phttpHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("fiberCtx", c)
		handler := fasthttpadaptor.NewFastHTTPHandler(h)
		handler(c.Context())
		return nil
	}
}
