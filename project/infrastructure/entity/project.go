package entity

type Projects struct {
	ID          int     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      string  `gorm:"not null;column:user_id" json:"user_id;foreignkey:UserID;constraint:OnDelete:CASCADE;" binding:"max=1000;column:user_id"` // Assuming UserID is a string (VARCHAR) for consistency with foreign key
	ProjectName string  `gorm:"not null" json:"name; column:name"`
	Client      string  `gorm:"null" json:"client; column:client"`                   // Nullable in database
	Amount      float32 `gorm:"type:decimal(15,2);null" json:"amount;column:amount"` // Nullable in database
	Tracked     float32 `gorm:"type:decimal(15,2);default:0.00;null" json:"tracked;column:tracked"`
	TimeEntryID int     `gorm:"not null" json:"time_entry_id; column:time_entry_id"`
	// Default value specified
	// Default value specified
}

func (Projects) TableName() string {
	return "Projects"
}
