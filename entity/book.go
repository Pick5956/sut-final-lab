package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title string  `valid:"length(3|100)~Title must be between 50 and 5000"`
	Price float64 `valid:"range(50|5000)~Price must be between 50 and 5000,required~Price must be between 50 and 5000"`
	Code  string  `valid:"matches(BK\\w+$)~Code must start with BK followed by 6 digits (0-9),uppercase~Code must start with BK followed by 6 digits (0-9),required~Code must start with BK followed by 6 digits (0-9),length(8|8)~Code must start with BK followed by 6 digits (0-9)"`
}
