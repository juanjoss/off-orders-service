package model

type Product struct {
	Barcode          string  `db:"barcode" json:"barcode"`
	Name             string  `db:"name" json:"name"`
	Quantity         string  `db:"quantity" json:"quantity"`
	ImageUrl         string  `db:"image_url" json:"image_url"`
	Energy100g       float32 `db:"energy_100g" json:"energy_100g"`
	EnergyServing    float32 `db:"energy_serving" json:"energy_serving"`
	NutrientLevelsId uint8   `db:"nutrient_levels_id" json:"nutrient_levels_id"`
	NovaGroup        uint8   `db:"nova_group" json:"nova_group"`
	NutriscoreScore  int8    `db:"nutriscore_score" json:"nutriscore_score"`
	NutriscoreGrade  string  `db:"nutriscore_grade" json:"nutriscore_grade"`
}
