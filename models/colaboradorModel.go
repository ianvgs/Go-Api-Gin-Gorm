package models

import "gorm.io/gorm"

type Colaborador struct {
	gorm.Model
	Id        int    `json:"id" binding:"required" gorm:"primaryKey"`
	Nome      string `json:"nome" `
	Email     string `json:"email" `
	Ativo     string `json:"ativo" `
	CreatedAt string `json:"createdAt" `
	UpdatedAt string `json:"updatedAt" `
	Noticias  []Noticia
}
