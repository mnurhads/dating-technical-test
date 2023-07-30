package controllers

import (
	"net/http"
    "time"
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
        photo       models.Photo
        profil      models.Profil
        persen      models.Persentase
        userNow     models.UserLink
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

    // tampilkan id user terbaru
    tx.Raw("SELECT id FROM user WHERE username = ?", request.Username).Scan(&userNow)

    photo.UserId = userNow.Id
    photos := tx.Table("photo").Create(&photo).Error
    
    if(photos != nil) {
        panic("photo not insert")
    }

    profil.UserId = userNow.Id
    profils := tx.Table("profil").Create(&profil).Error
    
    if(profils != nil) {
        panic("profil not insert")
    }

    persen.UserId = userNow.Id
    persens := tx.Table("persentase").Create(&persen).Error
    
    if(persens != nil) {
        panic("photo not insert")
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
    response.Data.TokenData.AccesToken   = tokenString
    response.Data.TokenData.TokenType    = "Bearer"
    response.Data.TokenData.ExpiresIn    = "900"
    response.Data.UserData.Username      = user.Username
    response.Data.UserData.Fullname      = user.Fullname
    response.Data.UserData.Email         = user.Email
    response.Data.UserData.Notelp        = user.Notelp
    response.Data.UserData.Status        = user.Status 

    c.JSON(http.StatusOK, response)
    return;
}

func (idb *InDB) ProfilService(c *gin.Context) {
    stringClientKey 	:= c.Request.Header.Get("secret-key")
    secretKey           := goDotEnvVariable("SECRET_KEY")
    var (
        request     structs.RequestProfilList
        response    structs.ResponseProfilLists
        errors      structs.ErrorResponse
        user        models.User
        profilNew   models.ProfilLink
        premi       models.Premium
        status      string
        gender      string
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

    // tampilkan data user
    tx.Raw("SELECT * FROM user WHERE id = ?", request.UserId).Scan(&user)
    if(user.Id == 0) {
        errors.ResponseCode   = 404
        errors.ResponseMsg    = "User Tidak diketahui"

        c.JSON(http.StatusOK, errors)
        return
    }
    // select premium
    tx.Raw("SELECT * FROM premium WHERE user_id = ?", request.UserId).Scan(&premi)
    
    if(premi.Id < 1) {
        tx.Raw("SELECT profil.gender, profil.age, profil.birthdate, profil.birth_info, profil.bio, photo.image FROM profil JOIN user ON profil.user_id = user.id JOIN photo ON profil.user_id = photo.user_id WHERE profil.user_id = ? AND profil.lokasi = ? LIMIT 10", request.Lokasi, request.UserId).Scan(&profilNew)
        status = "NOT PREMIUM"
    } else {
        tx.Raw("SELECT profil.gender, profil.age, profil.birthdate, profil.birth_info, profil.bio, photo.image FROM profil JOIN user ON profil.user_id = user.id JOIN photo ON profil.user_id = photo.user_id WHERE profil.user_id = ? AND profil.lokasi = ?", request.Lokasi, request.UserId).Scan(&profilNew)
        status = "PREMIUM MEMBER"
    }

    if(profilNew.Gender == 1) {
        gender = "Laki-Laki"
    } else {
        gender = "Perempuan"
    }

    tx.Commit()

    response.ResponseCode = 200
    response.ResponseMsg  = "List Profil"
    response.Premium      = status
    response.Profil.Age               = profilNew.Age
    response.Profil.Birthdate         = profilNew.Birthdate
    response.Profil.BirthInfo         = profilNew.BirthInfo
    response.Profil.Bio               = profilNew.Bio
    response.Profil.Gender            = gender
    response.Profil.Image             = profilNew.Image
    response.Profil.Lokasi            = request.Lokasi

    c.JSON(http.StatusOK, response)
    return;
}

func (idb *InDB) ProfilUpdateService(c *gin.Context) {
    stringClientKey 	:= c.Request.Header.Get("secret-key")
    secretKey           := goDotEnvVariable("SECRET_KEY")
    var (
        request     structs.RequestUpdateProfil
        response    structs.ResponseUpdateProfil
        errors      structs.ErrorResponse
        user        models.User
        profil      models.Profil
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

    birthDate := request.Birthdate
    now, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		panic(err)
	}

    tx := idb.DB.Begin()

    // tampilkan data user
    tx.Raw("SELECT * FROM user WHERE id = ?", request.UserId).Scan(&user)
    if(user.Id == 0) {
        errors.ResponseCode   = 404
        errors.ResponseMsg    = "User Tidak diketahui"

        c.JSON(http.StatusOK, errors)
        return
    }

    // tampilkan data profil
    tx.Raw("SELECT * FROM profil WHERE user_id = ?", request.UserId).Scan(&profil)

    // update profil
    tx.Exec("UPDATE profil SET gender = ?, age = ?, birthdate = ?, birth_info = ?, bio = ?, lokasi = ? WHERE user_id = ?", request.Gender.Value, request.Age, string(now.Format("2006-01-02")), request.BirthInfo, request.Bio, request.Lokasi, request.UserId)
    // update photo
    tx.Exec("UPDATE photo SET image = ? WHERE user_id = ?", request.Image, request.UserId)

    tx.Commit()

    // response
    response.ResponseCode = 200
    response.ResponseMsg  = "Updated Profil Successfully"
    response.ProfilData.UserData.Username = user.Username
    response.ProfilData.UserData.Fullname = user.Fullname
    response.ProfilData.UserData.Notelp   = user.Notelp
    response.ProfilData.UserData.Email    = user.Email
    response.ProfilData.Age               = profil.Age
    response.ProfilData.Birthdate         = profil.Birthdate
    response.ProfilData.BirthInfo         = profil.BirthInfo
    response.ProfilData.Bio               = profil.Bio
    response.ProfilData.Gender.Kode       = request.Gender.Kode
    response.ProfilData.Gender.Value      = request.Gender.Value
    response.ProfilData.Image.Gambar      = request.Image

    c.JSON(http.StatusOK, response)
    return;
}

func (idb *InDB) PremiumService(c *gin.Context) {
    stringClientKey 	:= c.Request.Header.Get("secret-key")
    secretKey           := goDotEnvVariable("SECRET_KEY")
    var (
        request     structs.PremiumRequest
        response    structs.PremiumResponse
        errors      structs.ErrorResponse
        user        models.User
        premium     models.Premium
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

    exp := request.Expired
    now, err := time.Parse("2006-01-02", exp)
	if err != nil {
		panic(err)
	}

    tx := idb.DB.Begin()

    // tampilkan data user
    tx.Raw("SELECT * FROM user WHERE id = ?", request.UserId).Scan(&user)
    if(user.Id == 0) {
        errors.ResponseCode   = 404
        errors.ResponseMsg    = "User Tidak diketahui"

        c.JSON(http.StatusOK, errors)
        return
    }

    premium.UserId    = request.UserId
    premium.Expired   = string(now.Format("2006-01-02"))
    premi := tx.Table("premium").Create(&premium).Error

    if(premi != nil) {
        errors.ResponseCode   = 404
        errors.ResponseMsg    = "Upgrade Premium Failed"

        c.JSON(http.StatusOK, errors)
        return
    }

    tx.Commit()

    response.ResponseCode = 200
    response.ResponseMsg  = "Upgrade Premium Profil Successfully"

    c.JSON(http.StatusOK, response)
    return
}

func (idb *InDB) LikeService(c *gin.Context) {

}

func (idb *InDB) DislikeService(c *gin.Context) {

}

func (idb *InDB) MatchService(c *gin.Context) {

}

func randomString(length int) string {
	var letters = []rune("123456789098765432101112131415161718192021785627362")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

