package models

import "gorm.io/gorm"

type Worker struct{
	gorm.Model
	Fullname string `json:"name" validate:"min=4,alpha"`
	Email string `json:"email" validate:"(required,email)`//`gorm:"unique"`
	Password string `json:"pass" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Gender string `json:"gender"`
	Rating float32 `gorm:"default:0.0"`
	TaskCompleted uint `gorm:"default:0"`
	WalletId string
}