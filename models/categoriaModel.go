package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model

	Id             int    `json:"id" binding:"required" gorm:"primaryKey"`
	Nome           string `json:"nome" `
	Descricao      string `json:"descricao" `
	Img            string `json:"img" `
	Bloqueio       string `json:"bloqueio"`
	Ativo          string `json:"ativo" `
	DataVencimento string `json:"dataVEncimento" `
	CreatedAt      string `json:"createdAt" `
	UpdatedAt      string `json:"updatedAt" `
}
