package model

type CategoryModelInsert struct {
	BaseModel `bson:",inline"`
	Name      string `bson:"name"`
	Type      int    `bson:"type"`
	Slug      string `bson:"slug"`
}

type CategoryModelGet struct {
	BaseModel
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Type int    `bson:"type"`
	Slug string `bson:"slug"`
}

type CategoryModelRequest struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
	Slug string `bson:"slug"`
}
