package models

import (
	"gorm.io/gorm"
)

type Noticia struct {
	gorm.Model
	Id            int    `json:"id" binding:"required" gorm:"primaryKey"`
	ColaboradorId int    `json:"colaborador_id" binding:"required"`
	Titulo        string `json:"titulo" binding:"required"`
	Resumo        string `json:"resumo" binding:"required"`
	Texto         string `json:"texto" binding:"required"`
	IdCategoria   int    `json:"categoria_id" binding:"required"`
	Ativo         string `json:"ativo" binding:"required"`
	CreatedAt     string `json:"createdAt" binding:"required"`
	UpdatedAt     string `json:"updatedAt" binding:"required"`
	Views         string `json:"views" binding:"required"`
	Colaborador   Colaborador
}
