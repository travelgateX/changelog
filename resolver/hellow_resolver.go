package resolver
import(
	graphql "github.com/graph-gophers/graphql-go"
)
type HelloWorldResolver struct{
	ID graphql.ID
	HelloWorldText string
}
func (r *HelloWorldResolver) HelloWorld() string {
	return r.HelloWorldText
}

/*func (r *helloWorldResolver) Hello(ctx context.Context) (string, error) {
	return "Hello world!", nil
}*/

// Code :
func (r *HelloWorldResolver) Code() graphql.ID {
	return r.ID
}
