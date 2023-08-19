package models

import (
	"time"
)

type Noticia struct {
	Id            int       `gorm:"column:id" binding:"required" gorm:"primaryKey"`
	ColaboradorId int       `gorm:"column:idColaborador" binding:"required" `
	CategoriaId   int       `gorm:"column:idCategoria" binding:"required" `
	Titulo        string    `gorm:"column:titulo" binding:"required"`
	Resumo        string    `gorm:"column:resumo" binding:"required"`
	Texto         string    `gorm:"column:texto" binding:"required"`
	ImgPath       string    `gorm:"column:imgPath" binding:"required"`
	IdCategoria   int       `gorm:"column:idCategoria" binding:"required"`
	Ativo         string    `gorm:"column:ativo" binding:"required"`
	CreatedAt     time.Time `gorm:"column:createdAt" binding:"required"`
	UpdatedAt     time.Time `gorm:"column:updatedAt" binding:"required"`
	Views         string    `gorm:"column:views" binding:"required"`
	Colaborador   Colaborador
	Categoria     Categoria
}

/* func (n Noticia) FormattedDate() string {
	return n.CreatedAt.Format("Jan 02, 2006")

} */

/* func (n *Noticia) FormatCreatedAt() {
	n.CreatedAt = n.CreatedAt.Format("Jan 02, 2006")
}
*/
