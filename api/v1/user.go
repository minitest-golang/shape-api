package v1

import (
	"minitest/db"
	"minitest/utils"
	"net/http"
	"regexp"
	"strings"
)

const PasswordLenMin = 8

// GetUserName will return username string based on the Bearer token
// in the request header. This function will throw an exception
// if the token is not valid, or the tocken has expired.
func GetUserName(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if strings.Contains(authorization, "Bearer") {
		authorization = strings.Split(authorization, " ")[1]
	}
	username := utils.JwtDecode(authorization)
	return username
}

// @Summary Signup
// @Description Client uses this API to create a new user.
// @Tags UserApi
// @Accept json
// @Produce json
// @Param SignUpInfo body v1.SignupInfo true "User Signup Body"
// @Success 200 {object} v1.OKResponse
// @Failure 400 {object} v1.ErrorResponse
// @Router /app/v1/user/signup [post]
func UserSignUpApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusBadRequest
	var status string = utils.UNKNOWN_ERROR
	var respData interface{}
	utils.Block{
		Try: func() {
			requestBody := utils.RequestJson(r)
			username := requestBody["username"].(string)
			password := requestBody["password"].(string)
			passwordMd5 := utils.Md5(password)

			if username == "" {
				utils.ThrowException(utils.USER_NAME_ERROR)
			}
			if len(password) < PasswordLenMin {
				utils.ThrowException(utils.PASSWORD_TOO_SHORT)
			}

			isMatch, err := regexp.MatchString("^[a-zA-Z0-9_+-]*$", username)
			if err != nil || !isMatch {
				utils.ThrowException(utils.USER_NAME_ERROR)
			}

			err = db.UserTblInstance().Create(username, passwordMd5)
			if err != nil {
				utils.ThrowException(utils.SIGNUP_ERROR)
			}
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = e
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

// @Summary Login
// @Description User uses this API to login the system.
// @Tags UserApi
// @Accept json
// @Produce json
// @Param LoginInfo body v1.LoginInfo true "User Login Body"
// @Success 200 {object} v1.LoginSuccessResp
// @Failure 401 {object} v1.UnauthorizedResponse
// @Router /app/v1/user/login [post]
func UserLoginApiHandler(rw http.ResponseWriter, r *http.Request) {
	var respCode int = http.StatusUnauthorized
	var status string = utils.UNAUTHORIZED
	var respData = map[string]interface{}{}

	utils.Block{
		Try: func() {
			requestBody := utils.RequestJson(r)
			username := requestBody["username"].(string)
			password := requestBody["password"].(string)
			passwordMd5 := utils.Md5(password)

			err := db.UserTblInstance().Login(username, passwordMd5)
			if err != nil {
				utils.ThrowException(utils.UNAUTHORIZED)
			}
			respData["access_token"] = utils.JwtToken(username)
			respCode = http.StatusOK
		},
		Catch: func(e string) {
			status = utils.UNAUTHORIZED
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
