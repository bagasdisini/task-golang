package e_presencedto

type CreatePresenceRequest struct {
	Type    string `json:"type" validate:"required"`
	Tanggal string `json:"tanggal" validate:"required"`
	Waktu   string `json:"waktu" validate:"required"`
}

type UpdatePresenceRequest struct {
	IsApprove string `json:"isApprove"`
}

type PresenceResponse struct {
	IDUser       int    `json:"id_user"`
	Nama         string `json:"nama_user"`
	Tanggal      string `json:"tanggal"`
	WaktuMasuk   string `json:"waktu_masuk"`
	WaktuPulang  string `json:"waktu_pulang"`
	StatusMasuk  string `json:"status_masuk"`
	StatusPulang string `json:"status_pulang"`
}
