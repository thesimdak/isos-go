package models

type Category struct {
	ID          int64   `json:"id"`
	CategoryKey string  `json:"categoryKey"`
	Label       string  `json:"label"`
	RopeLength  float32 `json:"ropeLength"`
}
