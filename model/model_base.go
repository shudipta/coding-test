package model

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// Retrieve action when record is retrieved from db
	Retrieve = Action(1)

	// Update action when record is updated in db
	Update = Action(2)

	// Delete action when record is deleted in db
	Delete = Action(3)
)

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
}
