package test

import (
	"changelog/model"
	"changelog/resolver"
	"testing"

	"github.com/icrowley/fake"
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

/*func TestCommitResolver(t *testing.T) {
	var (
		id       = int32(rand.Intn(1000))
		message  = fake.Paragraph()
		user     = fake.EmailAddress()
		release  = fmt.Sprintf("%v.%v.%v", rand.Intn(3), rand.Intn(50), rand.Intn(100))
		resource = fake.Word()
		category = model.Added
		released = (rand.Float32() < 0.5)
	)

	r := resolver.{
		Commit: Resolver.Change{
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
}*/
func TestCreateChangeResolver(t *testing.T) {

	var (
		code    = "testCode"
		message = fake.Paragraph()
		project = fake.Paragraph()
		typo    = "ADDED"
	)

	r := model.CreateChangeInput{
		Code:    &code,
		Message: message,
		Project: project,
		Type:    model.ChangeType(typo),
	}

	if *r.Code != code {
		t.Errorf("ERROR: with code: %s", code)
	} else if r.Message != message {
		t.Errorf("ERROR: with message: %s", message)
	} else if r.Project != project {
		t.Errorf("ERRPR: with project: %s", project)
	} else if r.Type != model.ChangeType(typo) {
		t.Errorf("ERROR: with tpye: %s", typo)
	}
}

func TestUpdateChange(t *testing.T) {
	data := model.UpdateChangeInput{
		Code:    "ch2",
		Message: "messageTest",
		Project: "projectTest",
		Type:    model.ChangeType("FIXED"),
	}

	var r resolver.Resolver

	//resolver := resolver{}

	inputData := struct{ Input *model.UpdateChangeInput }{&data}
	cResovler := *r.UpdateChange(inputData)

	if data.Code != string(cResovler.Code()) {
		t.Errorf("ERROR: with code: %s", data.Code)
	}

}

/*func TestDeleteChange(t *testing.T) {
	data := model.DeleteChangeInput{
		Code: "ch2",
	}
	inputData := struct{ Input *model.DeleteChangeInput }{&data}

	var r resolver.Resolver
	cResolver := *r.DeleteChange(inputData)

	if data.Code == string(cResolver.Code()) {
		t.Errorf("ERROR: with code: %s", data.Code)
	}

}*/
