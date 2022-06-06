package model

type NutrientLevels struct {
	Fat          string `db:"fat"`
	SaturatedFat string `db:"saturated_fat"`
	Sugar        string `db:"sugar"`
	Salt         string `db:"salt"`
}

type Nutriments struct {
	Energy100g    float32
	EnergyServing float32
	NOVA          uint8 `db:"nova_group"`
}

type NutriscoreData struct {
	Score int8   `db:"nutriscore"`
	Grade string `db:"nutriscore_grade"`
}
