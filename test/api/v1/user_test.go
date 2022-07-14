package v1_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	v1 "minitest/api/v1"
	"minitest/db"
	"minitest/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func reqReader(req map[string]interface{}) *strings.Reader {
	b, err := json.Marshal(req)
	if err != nil {
		panic("*** TEST SETUP ERROR! ***")
	}
	return strings.NewReader(string(b))
}

func readJson(r *httptest.ResponseRecorder) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ThrowException(utils.BAD_REQUEST)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		utils.ThrowException(utils.BAD_JSON)
	}
	return data
}

func TestUserSignUpApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	mux := http.NewServeMux()
	mux.HandleFunc("/app/v1/user/signup", v1.UserSignUpApiHandler)

	convey.Convey("Test@UserSignUpApiHandler - Signup OK", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"username": "admin",
			"password": "admin@123",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/signup", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
	})

	convey.Convey("Test@UserSignUpApiHandler - Empty username", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"username": "",
			"password": "admin",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/signup", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})

	convey.Convey("Test@UserSignUpApiHandler - Username has special character", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"username": "admin@1",
			"password": "admin",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/signup", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})

	convey.Convey("Test@UserSignUpApiHandler - Password too short", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"username": "admin",
			"password": "admin",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/signup", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})

	convey.Convey("Test@UserSignUpApiHandler - DB error", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = errors.New("DB Error")

		s := reqReader(map[string]interface{}{
			"username": "admin",
			"password": "admin@123",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/signup", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusBadRequest)
	})
}

func TestUserLoginApiHandler(t *testing.T) {
	// Use mock DB for unit test
	db.DB = db.NewDbDriver("mockdb")

	mux := http.NewServeMux()
	mux.HandleFunc("/app/v1/user/login", v1.UserLoginApiHandler)

	convey.Convey("Test@UserLoginApiHandler - Login OK", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"username": "admin",
			"password": "admin@123",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/login", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		jsonRep := readJson(writer)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(jsonRep["status"].(string), convey.ShouldEqual, "OK")
	})

	convey.Convey("Test@UserLoginApiHandler - Login failed due to invalid JSON", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"user_name": "admin",
			"password":  "admin@123",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/login", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@UserLoginApiHandler - Login failed due to invalid JSON", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = nil

		s := reqReader(map[string]interface{}{
			"username":  "admin",
			"pass_word": "admin@123",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/login", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})

	convey.Convey("Test@UserLoginApiHandler - Login failed due to DB error", t, func() {
		userTbl := db.UserTblInstance()
		userTbl.(*db.MockDBUserTbl).DummyError = errors.New("DB error")

		s := reqReader(map[string]interface{}{
			"username": "admin",
			"password": "admin@123",
		})

		request, _ := http.NewRequest("POST", "/app/v1/user/login", s)
		writer := httptest.NewRecorder()
		mux.ServeHTTP(writer, request)

		convey.So(writer.Code, convey.ShouldEqual, http.StatusUnauthorized)
	})
}
