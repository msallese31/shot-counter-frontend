package types

type AccelerometerData struct {

	// The `json` struct tag maps between the json name
	// and actual name of the field
	X []float32 `json:"X_Acc"`
	Y []float32 `json:"Y_Acc"`
	Z []float32 `json:"Z_Acc"`
}
