package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	// accountClient *account.Client
	// catalogClient *catalog.Client
	// orderClient   *order.Client
}

func NewGraphqlServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	return &Server{}, nil
	// accountClient, err := account.Client(accountUrl)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogClient, err := catalog.Client(catalogUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	return nil, err
	// }
	// orderClient, err := order.Client(orderUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	catalogClient.Close()
	// 	return nil, err
	// }

	// return &Server{
	// 	accountClient,
	// 	catalogClient,
	// 	orderClient,
	// }, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
