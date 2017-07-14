package data_models

type LoginAuthQuery struct {
	UserId string `bson:"userId"`
	Password string `bson:"password"`
}
