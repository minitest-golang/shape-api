package v1

import (
	"minitest/db"
	"minitest/shape"
	"minitest/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

// @Summary Create Shape
// @Description Client uses this API to create a shape.
// @Tags ShapeApi
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param ShapeInfo body v1.ShapeInfo true "Shape Information"
// @Success 200 {object} v1.ShapeCreateResp
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/shape/create [post]
func CreateShapeApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.BAD_REQUEST
	var respData interface{}
	utils.Block{
		Try: func() {
			username := GetUserName(r)
			requestBody := utils.RequestJson(r)
			objShape := requestBody["shape"].(string)
			edges := requestBody["edges"].([]interface{})
			edgeStrs := []string{}
			for _, e := range edges {
				edgeStrs = append(edgeStrs, e.(string))
			}

			_, err := shape.CreateShape(objShape, edgeStrs)
			if err != nil {
				utils.ErrorLog("Failed to create shape (err=%s)", err.Error())
				utils.ThrowException(err.Error())
			}
			id, err := db.ShapeTblInstance().Create(username, objShape, edgeStrs)
			if err != nil {
				utils.ErrorLog("Failed to create shape (err=%s)", err.Error())
				utils.ThrowException(err.Error())
			}
			respData = map[string]int64{
				"shape_id": id,
			}
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
			if strings.Contains(status, "Unauthorized") {
				respCode = http.StatusUnauthorized
			}
		},
		Finally: func() {
			if respCode == 200 || respCode == 204 {
				status = utils.OK
			}
			var resp = utils.JsonResponse{Status: status, Data: respData}
			resp.Response(rw, respCode)
		},
	}.Do()
}

// @Summary Get All Shapes
// @Description Client uses this API to get all created shapes.
// @Tags ShapeApi
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} v1.AllShapeInfoResp
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/shape [get]
func GetAllShapeApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.UNKNOWN_ERROR
	var respData interface{}
	utils.Block{
		Try: func() {
			username := GetUserName(r)
			shapes, err := db.ShapeTblInstance().GetAll(username)
			if err != nil {
				utils.ErrorLog("Failed to get all shapes (user=%s, err=%s)", username, err.Error())
				utils.ThrowException(err.Error())
			}
			respData = shapes
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
			if strings.Contains(status, "Unauthorized") {
				respCode = http.StatusUnauthorized
			}
		},
		Finally: func() {
			if respCode == 200 || respCode == 204 {
				status = utils.OK
			}
			var resp = utils.JsonResponse{Status: status, Data: respData}
			resp.Response(rw, respCode)
		},
	}.Do()
}

// @Summary Get Specific Shape From ID
// @Description Client uses this API to get a created shape with specific shape ID.
// @Tags ShapeApi
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} v1.ShapeInfo
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/shape/{shape_id} [get]
func GetShapeApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.UNKNOWN_ERROR
	var respData interface{}
	utils.Block{
		Try: func() {
			username := GetUserName(r)
			params := mux.Vars(r)
			shape_id, ok := params["shape_id"]
			if !ok {
				utils.ThrowException(utils.BAD_REQUEST)
			}
			sid, err := strconv.ParseInt(shape_id, 10, 64)
			if err != nil {
				utils.ThrowException(utils.BAD_REQUEST)
			}
			shapes, err := db.ShapeTblInstance().Get(sid)
			if err != nil {
				utils.ErrorLog("Failed to get shape (user=%s, shapeId=%d, err=%s)", username, sid, err.Error())
				utils.ThrowException(err.Error())
			}
			respData = shapes
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
			if strings.Contains(status, "Unauthorized") {
				respCode = http.StatusUnauthorized
			}
		},
		Finally: func() {
			if respCode == 200 || respCode == 204 {
				status = utils.OK
			}
			var resp = utils.JsonResponse{Status: status, Data: respData}
			resp.Response(rw, respCode)
		},
	}.Do()
}

// @Summary Calculate Area or Perimeter of a specific Shape based on ID
// @Description Client uses this API to Calculate Area or Perimeter of a specific Shape based on ID.
// @Tags ShapeApi
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param ShapeCalculationBody body v1.ShapeCalculationBody true "GaphQL Query"
// @Success 200 {object} v1.ShapeCalculationResp
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/shape/calculate [post]
func CalculateShapeApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.UNKNOWN_ERROR
	var respData interface{}
	utils.Block{
		Try: func() {
			_ = GetUserName(r) // Valida access token
			requestBody := utils.RequestJson(r)
			query := requestBody["query"]

			schema := ShapeScheme()
			if schema == nil {
				utils.ThrowException(utils.BAD_REQUEST)
			}
			result := graphql.Do(graphql.Params{
				Schema:        *schema,
				RequestString: query.(string),
			})
			respData = result.Data
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
			if strings.Contains(status, "Unauthorized") {
				respCode = http.StatusUnauthorized
			}
		},
		Finally: func() {
			if respCode == 200 || respCode == 204 {
				status = utils.OK
			}
			var resp = utils.JsonResponse{Status: status, Data: respData}
			resp.Response(rw, respCode)
		},
	}.Do()
}

// @Summary Update a specific shape
// @Description Client uses this API to update shape's edge values.
// @Tags ShapeApi
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param EdgeInfo body v1.EdgeInfo true "Shape's Edge Information"
// @Success 200 {object} v1.OKResponse
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/shape/{shape_id} [put]
func UpdateShapeApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.UNKNOWN_ERROR
	var respData interface{}
	utils.Block{
		Try: func() {
			username := GetUserName(r)
			params := mux.Vars(r)
			shape_id, ok := params["shape_id"]
			if !ok {
				utils.ThrowException(utils.BAD_REQUEST)
			}
			sid, err := strconv.ParseInt(shape_id, 10, 64)
			if err != nil {
				utils.ThrowException(utils.BAD_REQUEST)
			}

			requestBody := utils.RequestJson(r)
			edges := requestBody["edges"].([]interface{})
			edgeStrs := []string{}
			for _, e := range edges {
				edgeStrs = append(edgeStrs, e.(string))
			}

			shapeObj, err := db.ShapeTblInstance().Get(sid)
			if err != nil {
				utils.ErrorLog("Failed to get shape (user=%s, shapeId=%d, err=%s)", username, sid, err.Error())
				utils.ThrowException(err.Error())
			}

			_, err = shape.CreateShape(shapeObj.Shape, edgeStrs)
			if err != nil {
				utils.ErrorLog("Failed to update shape (err=%s)", err.Error())
				utils.ThrowException(err.Error())
			}

			err = db.ShapeTblInstance().Update(sid, edgeStrs)
			if err != nil {
				utils.ErrorLog("Failed to update shape (err=%s)", err.Error())
				utils.ThrowException(err.Error())
			}
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
			if strings.Contains(status, "Unauthorized") {
				respCode = http.StatusUnauthorized
			}
		},
		Finally: func() {
			if respCode == 200 || respCode == 204 {
				status = utils.OK
			}
			var resp = utils.JsonResponse{Status: status, Data: respData}
			resp.Response(rw, respCode)
		},
	}.Do()
}

// @Summary Delete a specific shape
// @Description Client uses this API to delete a shape specified by ID.
// @Tags ShapeApi
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} v1.OKResponse
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/shape/{shape_id} [delete]
func DeleteShapeApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.UNKNOWN_ERROR
	var respData interface{}
	utils.Block{
		Try: func() {
			_ = GetUserName(r)
			params := mux.Vars(r)
			shape_id, ok := params["shape_id"]
			if !ok {
				utils.ThrowException(utils.BAD_REQUEST)
			}
			sid, err := strconv.ParseInt(shape_id, 10, 64)
			if err != nil {
				utils.ThrowException(utils.BAD_REQUEST)
			}
			err = db.ShapeTblInstance().Delete(sid)
			if err != nil {
				utils.ErrorLog("Failed to delete shape (err=%s)", err.Error())
				utils.ThrowException(err.Error())
			}
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
			if strings.Contains(status, "Unauthorized") {
				respCode = http.StatusUnauthorized
			}
		},
		Finally: func() {
			if respCode == 200 || respCode == 204 {
				status = utils.OK
			}
			var resp = utils.JsonResponse{Status: status, Data: respData}
			resp.Response(rw, respCode)
		},
	}.Do()
}
