package resolver

import (
	"changelog/model"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// CreateChange :
func (r *Resolver) CreateChange(args struct{ Input *model.CreateChangeInput }) *ChangeResolver {

	var inputCode model.ChangeData
	r.db.First(&inputCode, "code = ?", args.Input.Code)
	if inputCode.Code != "" {
		inputError := model.AdviseMessageData{
			Code:        "409",
			Description: "Conflict: Code already stored",
			Level:       "ERROR",
		}
		err := []model.AdviseMessageData{}
		err = append(err, inputError)

		return &ChangeResolver{change: nil, err: &err}
	}

	time := time.Now()
	inputData := model.ChangeData{
		Code:       graphql.ID(*args.Input.Code),
		Message:    args.Input.Message,
		Project:    &args.Input.Project,
		ChangeDate: time,
		Type:       args.Input.Type,
		ReleaseID:  1,
		SourceID:   2,
	}
	r.db.Create(&inputData)
	return &ChangeResolver{change: &inputData, err: nil}
}

// UpdateChange :
func (r *Resolver) UpdateChange(args struct{ Input *model.UpdateChangeInput }) *ChangeResolver {

	var c model.ChangeData
	r.db.Where("code = ?", args.Input.Code).First(&c)
	if c.Code == "" {
		inputError := model.AdviseMessageData{
			Code:        "404",
			Description: "Not found: Code",
			Level:       "ERROR",
		}
		err := []model.AdviseMessageData{}
		err = append(err, inputError)

		return &ChangeResolver{change: nil, err: &err}
	}

	r.db.Model(&c).Updates(model.ChangeData{
		Message:    args.Input.Message,
		Project:    &args.Input.Project,
		ChangeDate: time.Now(),
		Type:       args.Input.Type})

	return &ChangeResolver{change: &c, err: nil}
}

// DeleteChange :
func (r *Resolver) DeleteChange(args struct{ Input *model.DeleteChangeInput }) *ChangeResolver {

	c := model.ChangeData{
		Code: graphql.ID(args.Input.Code),
	}
	var d model.ChangeData
	r.db.Where("code = ?", args.Input.Code).First(&d)
	if d.Code == "" {
		inputError := model.AdviseMessageData{
			Code:        "404",
			Description: "Not found: Code",
			Level:       "ERROR",
		}
		err := []model.AdviseMessageData{}
		err = append(err, inputError)
		return &ChangeResolver{change: nil, err: &err}

	}
	// Get all matched records
	r.db.Where("code = ?", args.Input.Code).Delete(c)

	return &ChangeResolver{change: &c, err: nil}
}
