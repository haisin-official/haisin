package responses

type GaGetResponse struct {
	// Ga is nillable
	Ga *string `json:"ga" binding:"required"`
}

type GaPostReponse struct {
	// Ga is required when post
	Ga string `json:"ga" binding:"required"`
}

type GaDeleteResponse struct {
	// Nothing is returned when it's success
}
