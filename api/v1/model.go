package v1

type OKResponse struct {
	Status string `json:"status" example:"OK"`
}

type ErrorResponse struct {
	Status string `json:"status" example:"An error description"`
}

type UnauthorizedResponse struct {
	Status string `json:"status" example:"Unauthorized"`
}

type SignupInfo struct {
	Username string `json:"username" example:"peterweb"`
	Password string `json:"password" example:"passw0rd"`
}

type LoginInfo struct {
	Username string `json:"username" example:"peterweb"`
	Password string `json:"password" example:"passw0rd"`
}

type TokenResp struct {
	Token string `json:"access_token" example:"abx...xyz"`
}

type LoginSuccessResp struct {
	Status string    `json:"status" example:"OK"`
	Data   TokenResp `json:"data"`
}

type ShapeInfo struct {
	Shape string   `json:"shape" example:"triangle"`
	Edges []string `json:"edges" example:"['1', '2', '3']"`
}

type EdgeInfo struct {
	Edges []string `json:"edges" example:"['1', '2', '3']"`
}

type ShapeId struct {
	ShapeId int `json:"shape_id" example:"123"`
}

type ShapeCreateResp struct {
	Status string  `json:"status" example:"OK"`
	Data   ShapeId `json:"data"`
}

type ShapeItem struct {
	ShapeId int `json:"shape_id" example:"123"`
	ShapeInfo
}

type AllShapeInfoResp struct {
	Status string      `json:"status" example:"OK"`
	Data   []ShapeItem `json:"data"`
}

type ShapeCalculationBody struct {
	Query string `json:"query" example:"{area(shape_id:1)\n perimeter(shape_id:2)}"`
}

type ShapeValue struct {
	Area      string `json:"area" example:"12.1234"`
	Perimeter string `json:"perimeter" example:"30.0000"`
}

type ShapeCalculationResp struct {
	Status string     `json:"status" example:"OK"`
	Data   ShapeValue `json:"data"`
}
