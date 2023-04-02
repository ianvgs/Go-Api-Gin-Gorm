package models

import (
	"time"
)

type Colaborador struct {
	Id        int       `gorm:"column:id" binding:"required" gorm:"primaryKey"`
	Nome      string    `gorm:"column:nome" binding:"required"`
	Email     string    `gorm:"column:email" binding:"required"`
	Ativo     string    `gorm:"column:ativo" binding:"required"`
	CreatedAt time.Time `gorm:"column:createdAt" binding:"required"`
	UpdatedAt time.Time `gorm:"column:updatedAt" binding:"required"`
	Noticias  []Noticia `gorm:"foreignKey:ColaboradorId;references:Id"`
}

func (Colaborador) TableName() string {
	return "colaborador"
}
