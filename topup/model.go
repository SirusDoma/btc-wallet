package topup

import "time"

type Topup struct {
	ID        uint64
	Amount    float64   `gorm:"index:idx_history,unique"`
	CreatedAt time.Time `gorm:"index:idx_history,unique"` // Unique to avoid duplicate entry on retry due to network failure
	UpdatedAt time.Time
}

type History struct {
	Amount float64
	Period time.Time `gorm:"column:datetime"`
}

type ResultSet struct {
	Offset uint
	Limit  uint
	Total  uint
}

type HistoryResultSet struct {
	ResultSet
	Histories []History
}
