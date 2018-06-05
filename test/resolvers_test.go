package resolvers_test

import (
	"changelog/model"
	"changelog/resolver"

	"fmt"
	"math/rand"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestRootResolver(t *testing.T) {
	t.Skip("Skipping test root resolver without mocked database")
}
func TestQueryResolver(t *testing.T) {
	t.Skip("Skipping test root resolver without mocked database")
}

func TestMutationResolver(t *testing.T) {
	t.Skip("Skipping test root resolver without mocked database")
}

func TestCommitResolver(t *testing.T) {
	var (
		id       = int32(rand.Intn(1000))
		message  = fake.Paragraph()
		user     = fake.EmailAddress()
		release  = fmt.Sprintf("%v.%v.%v", rand.Intn(3), rand.Intn(50), rand.Intn(100))
		resource = fake.Word()
		category = model.Added
		released = (rand.Float32() < 0.5)
	)

	r := resolver.CommitResolver{
		Commit: model.Commit{
			ID:       id,
			Message:  message,
			User:     user,
			Release:  release,
			Resource: resource,
			Category: category,
			Released: released,
		},
	}

	assert.Equal(t, r.ID(), graphql.ID(fmt.Sprint(id)))
	assert.Equal(t, r.Message(), message)
	assert.Equal(t, *r.User(), user)
	assert.Equal(t, *r.Release(), release)
	assert.Equal(t, *r.Resource(), resource)
	assert.Equal(t, r.Category(), category)
	assert.Equal(t, r.Released(), released)
}
