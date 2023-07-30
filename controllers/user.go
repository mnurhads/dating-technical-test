package controllers

import (
	"net/http"

	"datingapp/structs"
    "datingapp/models"
    Auth "datingapp/jwt"
	"github.com/gin-gonic/gin"
  	"os"
    "path/filepath"
    "log"
    "encoding/json"
    "io/ioutil"
    "math/rand"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
)

func goDotEnvVariable(key string) string {

    // load .env file
        cur := os.Getenv("GIN_ENV")
        folder := ""
      if cur == "production" {
            folder = "/var/www/go/datingapp/"
        }
        err := godotenv.Load(filepath.Join(folder, ".env"))
      //fmt.Println(filepath.Join(path,folder, ".env"))
  
        if err != nil {
          log.Fatalf("Error loading .env file")
        }
  
    return os.Getenv(key)
  }

func (idb *InDB) RegisterService(c *gin.Context) {
    stringClientKey 	:= c.Request.Header.Get("secret-key")
    secretKey           := goDotEnvVariable("SECRET_KEY")
    var (
        request     structs.RegisterRequest
        response    structs.RegisterResponse
        errors      structs.ErrorResponse
        user        models.User
    )

    jsonData,_          := ioutil.ReadAll(c.Request.Body)
    defer c.Request.Body.Close()
    json.Unmarshal(jsonData, &request) 

    if (stringClientKey != secretKey) {
        errors.ResponseCode   = 211
        errors.ResponseMsg    = "Invalid Secret Key"

        c.JSON(http.StatusOK, errors)
        return
    }

    if(request.Password != request.Confirm) {
        errors.ResponseCode   = 211
        errors.ResponseMsg    = "Password tidak sama"

        c.JSON(http.StatusOK, errors)
        return
    }

    tx := idb.DB.Begin()

    tx.Raw("SELECT * FROM user WHERE email = ?", request.Email).Scan(&user)

    if user.Id > 0 {
        tx.Rollback()
        errors.ResponseCode   = 211
        errors.ResponseMsg        = "Email Sudah terdaftar"

        c.JSON(http.StatusOK, errors)
        return
    }

    tx.Raw("SELECT * FROM user WHERE username = ?", request.Username).Scan(&user)

    if user.Id > 0 {
        tx.Rollback()
        errors.ResponseCode   = 211
        errors.ResponseMsg        = "Username Sudah terdaftar"

        c.JSON(http.StatusOK, errors)
        return
    }

    tx.Raw("SELECT * FROM user WHERE notelp = ?", request.Notelp).Scan(&user)

    if user.Id > 0 {
        tx.Rollback()
        errors.ResponseCode   = 211
        errors.ResponseMsg        = "No Telpon Sudah terdaftar"

        c.JSON(http.StatusOK, errors)
        return
    }

    pwdHash   := []byte(request.Password)
    pwd, erro := bcrypt.GenerateFromPassword(pwdHash, bcrypt.DefaultCost)
    
    if(erro != nil) {
        panic(erro)
    }

    user.Username = request.Username
    user.Fullname = request.Fullname
    user.Email    = request.Email
    user.Notelp   = request.Notelp
    user.Password = string(pwd)
    user.Status   = "AKTIF"

    err := tx.Table("user").Create(&user).Error
    if err != nil {
        tx.Rollback()
        errors.ResponseCode   = 500
        errors.ResponseMsg    = "Internal error"

        c.JSON(http.StatusOK, errors)
        return
    }

    tx.Commit()

    response.ResponseCode   = 200
    response.ResponseMsg    = "Register Successfully"

    c.JSON(http.StatusOK, response)
    return;

}

func (idb *InDB) LoginService(c *gin.Context) {
    stringClientKey 	:= c.Request.Header.Get("secret-key")
    secretKey           := goDotEnvVariable("SECRET_KEY")
    var (
        request     structs.LoginRequest
        response    structs.LoginResponse
        errors      structs.ErrorResponse
        user        models.User
    )

    jsonData,_  := ioutil.ReadAll(c.Request.Body)
    defer c.Request.Body.Close()
    json.Unmarshal(jsonData, &request)

    if (stringClientKey != secretKey) {
        errors.ResponseCode   = 211
        errors.ResponseMsg    = "Invalid Secret Key"

        c.JSON(http.StatusOK, errors)
        return
    }

    tx := idb.DB.Begin()

    tx.Raw("SELECT * FROM user WHERE username = ?", request.Username).Scan(&user)

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		errors.ResponseCode = 206
		errors.ResponseMsg = "Failed to login, please check you email password"
        c.JSON(http.StatusOK, errors)
		return
	}

    tokenString, err:= Auth.GenerateJWT(user.Username, user.Password)
	if err != nil {
		errors.ResponseCode = 207
		errors.ResponseMsg = "Invalid generate token"

        c.JSON(http.StatusOK, errors)
		return
	}

    tx.Model(&user).Updates(map[string]interface{}{"token": tokenString})
    tx.Commit()

    response.ResponseCode           = 200
    response.ResponseMsg            = "Login Successfully"
    response.TokenData.AccesToken   = tokenString
    response.TokenData.TokenType    = "Bearer"
    response.TokenData.ExpiresIn    = "900"
    response.UserData.Username      = user.Username
    response.UserData.Fullname      = user.Fullname
    response.UserData.Email         = user.Email
    response.UserData.Notelp        = user.Notelp
    response.UserData.Status        = user.Status 

    c.JSON(http.StatusOK, response)
    return;
}

func randomString(length int) string {
	var letters = []rune("123456789098765432101112131415161718192021785627362")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

