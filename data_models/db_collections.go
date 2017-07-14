package data_models

type UsersCollection struct {
	UserId string `bson:"userId"`
	Password string `bson:"password"`
	FirstName string `bson:"firstName"`
	LastName string `bson:"lastName"`
	AccessToken string `bson:"accessToken"`
}