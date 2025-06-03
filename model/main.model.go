package model

type MainModel struct {
	Id    string `json:"id" gorm:"primaryKey"`
	Data1 string `json:"data_1"`
	Data2 string `json:"data_2"`
	Data3 string `json:"data_3"`
	Data4 string `json:"data_4"`
	// TestinputID string
	// Testinput   TestInput `gorm:"foreignKey:TestinputID"`
}

type MainModelFill struct {
	Data1 string `json:"data_1"`
	Data2 string `json:"data_2"`
	Data3 string `json:"data_3"`
	Data4 string `json:"data_4"`
	// Nested *TestInput
}

type MetadataResponse struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

type TestInput struct {
	// ID    string
	Data5 string `json:"data_5"`
	Data6 string `json:"data_6"`
	Data7 string `json:"data_7"`
	Data8 string `json:"data_8"`
}
