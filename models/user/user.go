package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"crypto/sha256"

	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../../authorization"
)

type Rekening struct {
	NoRekening string `json:"norekening"`
	AtasNama   string `json:"atasnama"`
	Bank       string `json:"bank"`
}

type Pengguna struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	FotoProfil string        `json:"fotoprofil"` //simpan alamatnya saja
	Nama       string        `json:"nama"`
	IdDiri     string        `json:"iddiri"`
	JenisID    int           `json:"jenisid"` //1=KTP, 2=SIM, 3=Paspor
	TglLahir   string        `json:"tgllahir"`
	Norek      []Rekening    `json:"norek"`
	Email      string        `json:"email"`
	Gender     string        `json:"gender"`
	NoHp       string        `json:"nohp"`
	Alamat     string        `json:"alamat"`
}

func ErrorReturn(w http.ResponseWriter, pesan string, code int) string {
	w.WriteHeader(code)
	//fmt.Fprintf(w, "{error: %i, message: %q}", code, pesan)
	return "{error: " + strconv.Itoa(code) + ", message: " + pesan + "}"
}

/*func SuccessReturn(w http.ResponseWriter, json []byte, pesan string, code int) string {
	w.WriteHeader(code)
	//fmt.Fprintf(w, "{message: %q}", pesan)
	return string(json)
}*/

func SuccessReturn(w http.ResponseWriter, pesan string, code int) string {
	w.WriteHeader(code)
	return "{success: " + strconv.Itoa(code) + ", message: " + pesan + "}"
}

func CheckDupUser(s *mgo.Session, p Pengguna) string {
	var ret string

	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("user")

	d, _ := c.Find(bson.M{"username": p.Username}).Count()
	if d > 0 {
		ret = "Username"
	}

	d, _ = c.Find(bson.M{"iddiri": p.IdDiri}).Count()
	if d > 0 {
		if ret != "" {
			ret = ret + ", ID Diri"
		} else {
			ret = "ID Diri"
		}
	}

	d, _ = c.Find(bson.M{"email": p.Email}).Count()
	if d > 0 {
		if ret != "" {
			ret = ret + ", Email"
		} else {
			ret = "Email"
		}
	}

	d, _ = c.Find(bson.M{"nohp": p.NoHp}).Count()
	if d > 0 {
		if ret != "" {
			ret = ret + ", Nomor Handphone"
		} else {
			ret = "Nomor Handphone"
		}
	}

	return ret
}

func RegistrasiUser(ses *mgo.Session, w http.ResponseWriter, r *http.Request) string {
	var pengguna Pengguna

	sesTambah := ses.Copy()
	defer sesTambah.Close()

	err := json.NewDecoder(r.Body).Decode(&pengguna)
	if err != nil {
		return ErrorReturn(w, "Registrasi Gagal", http.StatusBadRequest)

	}

	c := sesTambah.DB("propos").C("user")
	encryptPass := sha256.Sum256([]byte(pengguna.Password))
	pengguna.Password = fmt.Sprintf("%x", encryptPass)

	checkdup := CheckDupUser(ses, pengguna)
	if checkdup != "" {
		return ErrorReturn(w, checkdup+" Sudah Digunakan", http.StatusBadRequest)
	}
	err = c.Insert(pengguna)
	if err != nil {
		/*if mgo.IsDup(err) {
			fmt.Println(err)
			return ErrorReturn(w, "Username Sudah Digunakan", http.StatusBadRequest)
		}*/
		return ErrorReturn(w, "Tidak Ada Jaringan", http.StatusInternalServerError)
	}

	//json, _ := json.Marshal(pengguna)
	return SuccessReturn(w, "Berhasil Registrasi", http.StatusCreated)
}

func GetUser(s *mgo.Session, w http.ResponseWriter, r *http.Request, path string) string {
	//Jika membuka profil user lain
	var user Pengguna
	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("user")

	err := c.Find(bson.M{"username": path}).One(&user)
	if err != nil {
		return ErrorReturn(w, "User Tidak Ditemukan", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	us, _ := json.Marshal(user)
	return string(us)
}

/*func EditUser(s *mgo.Session, w http.ResponseWriter, r *http.Request) string {

}*/

func LoginUser(s *mgo.Session, w http.ResponseWriter, r *http.Request) string {
	//Digunakan untuk login ke halaman user
	var log Pengguna
	ses := s.Copy()
	defer ses.Close()

	//fmt.Println("Login User")

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		//fmt.Println("Cari data")
		return ErrorReturn(w, "Login Gagal", http.StatusBadRequest)
	}

	c := ses.DB("propos").C("user")

	encryptPassLogin := fmt.Sprintf("%x", sha256.Sum256([]byte(log.Password)))

	err = c.Find(bson.M{"username": log.Username}).One(&log)
	if err != nil {
		//fmt.Println("User Hilang")
		return ErrorReturn(w, "User Tidak Ditemukan", http.StatusBadRequest)
	}

	encryptPass := log.Password
	//fmt.Println(encryptPass + " " + encryptPassLogin + " " + log.Password)
	if encryptPass == encryptPassLogin {
		w.WriteHeader(http.StatusOK)
		//fmt.Println("Tukang Hacking")
		retBody, err := json.MarshalIndent(log, "", " ")
		if err != nil {
			panic(err)
		}
		return jwt.TokenMaker(string(retBody), "anggunauranaufalwilliam")
		fmt.Println(log.Id) //nanti di-lock pake jwt
	}

	return ErrorReturn(w, "Password Salah", http.StatusForbidden)
}

/*func IndexCreating(s *mgo.Session) {
	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("user")

	index := mgo.Index{
		Key:        []string{"username", "iddiri", "email", "nohp"},
		Unique:     true,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

}*/

//Digunakan untuk mengontrol path dari user (/user/...)
func UserController(urle string, w http.ResponseWriter, r *http.Request) string {

	urle = urle[1:]
	pathe := strings.Split(urle, "/")
	fmt.Println(pathe[0] + " " + pathe[1])

	ses, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer ses.Close()
	ses.SetMode(mgo.Monotonic, true)
	//IndexCreating(ses)

	if pathe[0] == "login" {
		return LoginUser(ses, w, r)
	}

	if len(pathe) >= 2 {
		if pathe[1] == "registrasi" {
			return RegistrasiUser(ses, w, r)
		} else if pathe[1] != "" {
			return GetUser(ses, w, r, pathe[1])
		}
	}
	return ErrorReturn(w, "Path Tidak Ditemukan", http.StatusNotFound)
}
