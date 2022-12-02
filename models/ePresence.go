package models

type EPresence struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment;"`
	IDUser    int    `json:"id_user" form:"id_user" gorm:"type: int"`
	Type      string `json:"type" form:"type" gorm:"type: varchar(255)"`
	IsApprove string `json:"isApprove" form:"isApprove" gorm:"type: varchar(255)"`
	Tanggal   string `json:"tanggal" form:"tanggal" gorm:"type: varchar(255)"`
	Waktu     string `json:"waktu" form:"waktu" gorm:"type: varchar(255)"`
}
