package model

type DictTypeRes struct {
	DictName string `json:"name"`
	Remark   string `json:"remark"`
}


type DictDataRes struct {
	DictValue string `json:"key"`
	DictLabel string `json:"value"`
	IsDefault int    `json:"isDefault"`
	Remark    string `json:"remark"`
}
