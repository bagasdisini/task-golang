package models

type Users struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment;"`
	Nama          string `json:"nama" form:"nama" gorm:"type: varchar(255)"`
	Email         string `json:"email" form:"email" gorm:"type: varchar(255)"`
	NPP           int64  `json:"npp" form:"npp" gorm:"type: bigint"`
	NPPSupervisor int64  `json:"npp_supervisor" form:"npp_supervisor" gorm:"type: bigint"`
	Password      string `json:"password" form:"password" gorm:"type: varchar(255)"`
}
