package domain

type Block struct {
	Hash string `json:"hash"`
	Data string `json:"data" binding:"required"`
}
