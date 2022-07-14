package v1_test

import (
	"errors"
	v1 "minitest/api/v1"
	"minitest/db"
	"minitest/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/smartystreets/goconvey/convey"
)

func TestCreateShapeApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	accessToken := utils.JwtToken("admin")

	mux := http.NewServeMux()
	mux.HandleFunc("/app/v1/shape/create", v1.CreateShapeApiHandler)

	convey.Convey("Test@CreateShapeApiHandler - Create Shape OK", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"shape": "square",
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/create", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)
		data := jsonRep["data"].(map[string]interface{})

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
		convey.So(int(data["shape_id"].(float64)), convey.ShouldEqual, 123)
	})

	convey.Convey("Test@CreateShapeApiHandler - Create Shape failed due to bad token", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"shape": "square",
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/create", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken+"?")

		mux.ServeHTTP(writer, request)
		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@CreateShapeApiHandler - Create Shape failed due to bad JSON", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"shape___": "square",
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/create", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)
		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})

	convey.Convey("Test@CreateShapeApiHandler - Create Shape failed due to bad edges", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"shape": "triangle",
			"edges": []string{
				"5", "7",
			},
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/create", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)
		jsonRep := readJson(writer)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "shape: Invalid edge value!")
	})

	convey.Convey("Test@CreateShapeApiHandler - Create Shape failed due to DB error", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = errors.New("DB error")

		s := reqReader(map[string]interface{}{
			"shape": "square",
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/create", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)
		jsonRep := readJson(writer)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "DB error")
	})
}

func TestGetAllShapeApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	accessToken := utils.JwtToken("admin")

	mux := http.NewServeMux()
	mux.HandleFunc("/app/v1/shape", v1.GetAllShapeApiHandler)

	convey.Convey("Test@GetAllShapeApiHandler - Get All Shape OK", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		request, _ := http.NewRequest("GET", "/app/v1/shape", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)
		data := jsonRep["data"].([]interface{})

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
		convey.So(len(data), convey.ShouldEqual, 1)
		convey.So(int(data[0].(map[string]interface{})["shape_id"].(float64)), convey.ShouldEqual, 123)
	})

	convey.Convey("Test@GetAllShapeApiHandler - Get All Shape failed due to bad tocken", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		request, _ := http.NewRequest("GET", "/app/v1/shape", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken+"?")

		mux.ServeHTTP(writer, request)
		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@GetAllShapeApiHandler - Get All Shape failed due to DB error", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = errors.New("DB Error")

		request, _ := http.NewRequest("GET", "/app/v1/shape", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)
		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})
}

func TestGetShapeApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	accessToken := utils.JwtToken("admin")

	mux := mux.NewRouter()
	mux.HandleFunc("/app/v1/shape/{shape_id}", v1.GetShapeApiHandler).Methods("GET")

	convey.Convey("Test@GetShapeApiHandler - Get All Shape OK", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		request, _ := http.NewRequest("GET", "/app/v1/shape/123", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)
		data := jsonRep["data"].(map[string]interface{})

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
		convey.So(int(data["shape_id"].(float64)), convey.ShouldEqual, 123)
	})

	convey.Convey("Test@GetShapeApiHandler - Get All Shape failed due to bad tocken", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		request, _ := http.NewRequest("GET", "/app/v1/shape/123", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken+"?")

		mux.ServeHTTP(writer, request)
		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@GetShapeApiHandler - Get All Shape failed due to DB error", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = errors.New("DB Error")

		request, _ := http.NewRequest("GET", "/app/v1/shape/123", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)
		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})
}

func TestCalculateShapeApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	accessToken := utils.JwtToken("admin")

	mux := http.NewServeMux()
	mux.HandleFunc("/app/v1/shape/calculate", v1.CalculateShapeApiHandler)

	convey.Convey("Test@CalculateShapeApiHandler - Calculate Shape OK", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"query": "{area(shape_id:123) \n perimeter(shape_id:123)}",
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/calculate", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)
		data := jsonRep["data"].(map[string]interface{})

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
		convey.So(data["area"].(string), convey.ShouldEqual, "9.0000")
		convey.So(data["perimeter"].(string), convey.ShouldEqual, "12.0000")
	})

	convey.Convey("Test@CalculateShapeApiHandler - Calculate Shape failed due to bad query", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"query": "{area}",
		})

		request, _ := http.NewRequest("POST", "/app/v1/shape/calculate", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)
		data := jsonRep["data"].(map[string]interface{})

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
		convey.So(data["area"].(string), convey.ShouldEqual, "NA")
	})
}

func TestUpdateShapeApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	accessToken := utils.JwtToken("admin")

	mux := mux.NewRouter()
	mux.HandleFunc("/app/v1/shape/{shape_id}", v1.UpdateShapeApiHandler).Methods("PUT")

	convey.Convey("Test@UpdateShapeApiHandler - Update Shape OK", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("PUT", "/app/v1/shape/123", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
	})

	convey.Convey("Test@UpdateShapeApiHandler - Update Shape failed due to bad tocken", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("PUT", "/app/v1/shape/123", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken+"?")

		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@UpdateShapeApiHandler - Update Shape failed due to bad JSON", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"**edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("PUT", "/app/v1/shape/123", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})

	convey.Convey("Test@UpdateShapeApiHandler - Update Shape failed due to DB error", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = errors.New("DB error")

		s := reqReader(map[string]interface{}{
			"edges": []string{
				"5",
			},
		})

		request, _ := http.NewRequest("PUT", "/app/v1/shape/123", s)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})
}

func TestDeleteShapeApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	accessToken := utils.JwtToken("admin")

	mux := mux.NewRouter()
	mux.HandleFunc("/app/v1/shape/{shape_id}", v1.DeleteShapeApiHandler).Methods("DELETE")

	convey.Convey("Test@DeleteShapeApiHandler - Delete Shape OK", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		request, _ := http.NewRequest("DELETE", "/app/v1/shape/123", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
	})

	convey.Convey("Test@DeleteShapeApiHandler - Delete Shape failed due to invalid tocken", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = nil

		request, _ := http.NewRequest("DELETE", "/app/v1/shape/123", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken+"?")

		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@DeleteShapeApiHandler - Delete Shape failed due to DB error", t, func() {
		shapeTbl := db.ShapeTblInstance()
		shapeTbl.(*db.MockShapeTbl).DummyError = errors.New("DB error")

		request, _ := http.NewRequest("DELETE", "/app/v1/shape/123", nil)
		writer := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+accessToken)

		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})
}
