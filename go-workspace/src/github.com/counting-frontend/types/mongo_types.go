package types

// JSONObject is the interface for dealing with Json types
// type JSONObject interface {
// }

// User is a struct that is meant to store a user from our mongodb backend
type User struct {
	// The `json` struct tag maps between the json name
	// and actual name of the field
	ID              string `bson:"google_token"`
	Name            string `bson:"name"`
	Email           string `bson:"email"`
	DailyCount      int    `bson:"daily_count"`
	MonthlyCount    int    `bson:"monthly_count"`
	DailyRequests   int    `bson:"daily_requests"`
	MonthlyRequests int    `bson:"monthly_requests"`
}

// AccelData is a struct that is meant to store a the accelerometer values coming from our mongodb backend
type AccelData struct {
	// The `json` struct tag maps between the json name
	// and actual name of the field
	// TODO: DON'T FORGET BSON
	// TODO: Think about changing time to be unmarshalled into some sort of time object
	Time string `json:"time"`
	// This is almost certainly wrong.  Also why is id token in AccelerometerData???
	Values AccelerometerData `json:"values"`
}
