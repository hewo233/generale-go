package model

type CellType string

const (
	CellTypeEmpty    CellType = "empty"
	CellTypeMountain CellType = "mountain"
	CellTypeTroop    CellType = "troop"
	cellTypeKing     CellType = "king"
	cellTypeBarrack  CellType = "barrack"
)

type Cell struct {
	Type     CellType `json:"type"`
	OwnerID  string   `json:"owner_id,omitempty"`     // the id of a player , if the cell is a troop, null means no owner
	TroopNum int      `json:"troop_health,omitempty"` // the health of a troop, if the cell is a troop. Or the health of a barrack, if the cell is a barrack.
}

type MapModel struct {
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Seed   int      `json:"seed"`
	Cells  [][]Cell `json:"cells"`
}
