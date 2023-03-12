package responses

type GaGetResponse struct {
	Ga string `json:"ga" binding:"required"`
}

type GaPostReponse struct {
	Ga string `json:"ga" binding:"required"`
}
