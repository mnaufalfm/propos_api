package projek

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"../../authorization"
	"../user"
)

type Donatur struct {
	IdDonatur    bson.ObjectId `json:"iddonatur"`
	JumlahDonasi string        `json:"jumlahdonasi"`
}

type Komentator struct {
	IdKomentator bson.ObjectId `json:"idkomentator"`
	Komen        string        `json:"komen"`
	TanggalKomen string        `json:"tanggalkomen"`
}

type Content struct {
	AlamatKonten string `json:"alamatkonten"`
	JenisKonten  int    `json:"jeniskonten"` //1: video, 2: gambar
	AsalKonten   int    `json:"asalkonten"`  //1: luar, 2: database
}

type Description struct {
	Penjelasan string    `json:"penjelasan"`
	Konten     []Content `json:"konten"`
}

type Projek struct {
	Id                bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	NamaProjek        string          `json:"namaprojek"`
	FotoProjek        []string        `json:"fotoprojek"` //simpan alamatnya saja
	LinkYoutube       string          `json:"linkyoutube"`
	PenjelasanSingkat string          `json:"penjelasansingkat"`
	LatarBelakang     Description     `json:"latarbelakang"`
	IdPemilik         bson.ObjectId   `json:"idpemilik"`
	IdAnggota         []bson.ObjectId `json:"idanggota"`
	ParaDonatur       []Donatur       `json:"paradonatur"`
	ParaKomen         []Komentator    `json:"parakomen"`
	IdLikers          []bson.ObjectId `json:"idlikers"`
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

func CheckDupProjek(s *mgo.Session, p Projek) bool {
	//Rencana pengembangan menggunakan algoritme information retrieval
	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("projek")

	d, _ := c.Find(bson.M{"namaprojek": p.NamaProjek}).Count()

	if d > 0 {
		return false
	}

	return true
}

func UploadProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request) string {
	//Belum ada upload gambar
	resBody, _ := ioutil.ReadAll(r.Body)
	token := string(resBody)
	tokenSplit := strings.Split(token, ".")
	if len(tokenSplit) < 4 {
		return ErrorReturn(w, "Format Request Salah", http.StatusBadRequest)
	}
	//fmt.Println(tokenSplit[0] + "." + tokenSplit[1] + "." + tokenSplit[2])
	if !jwt.CheckToken(tokenSplit[0] + "." + tokenSplit[1] + "." + tokenSplit[2]) {
		return ErrorReturn(w, "Token yang Dikirimkan Invalid", http.StatusBadRequest)
	}

	var proyek Projek

	ses := s.Copy()
	defer ses.Close()

	//err := json.NewDecoder(r.Body).Decode(&proyek)
	err := json.Unmarshal([]byte(tokenSplit[3]), &proyek)
	if err != nil {
		return ErrorReturn(w, "Menambahkan Projek Gagal", http.StatusBadRequest)

	}

	c := ses.DB("propos").C("projek")

	if !CheckDupProjek(ses, proyek) {
		return ErrorReturn(w, "Projek Sudah Pernah Dibuat", http.StatusBadRequest)
	}

	err = c.Insert(proyek)
	if err != nil {
		return ErrorReturn(w, "Tidak Ada Jaringan", http.StatusInternalServerError)
	}

	return SuccessReturn(w, "Projek Berhasil Dibuat", http.StatusCreated)
}

func EditProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request, path string) string {
	//format panggilan localhost:9000/projek/edit/idprojek
	var projek Projek
	var pengg user.Pengguna
	ses := s.Copy()
	defer ses.Close()

	resBody, _ := ioutil.ReadAll(r.Body)
	token := string(resBody)
	tokenSplit := strings.Split(token, ".")
	if len(tokenSplit) < 4 {
		return ErrorReturn(w, "Format Request Salah", http.StatusBadRequest)
	}
	//fmt.Println(tokenSplit[0] + "." + tokenSplit[1] + "." + tokenSplit[2])
	if !jwt.CheckToken(tokenSplit[0] + "." + tokenSplit[1] + "." + tokenSplit[2]) {
		return ErrorReturn(w, "Token yang Dikirimkan Invalid", http.StatusBadRequest)
	}
	mess := jwt.Base64ToString(tokenSplit[1])

	err := json.Unmarshal([]byte(mess), &pengg)
	if err != nil {
		panic(err)
	}

	c := ses.DB("propos").C("projek")

	err = c.Find(bson.M{"_id": path}).One(&projek)

	if projek.IdPemilik != pengg.Id {
		return ErrorReturn(w, "Anda Tidak Diperkenankan Mengedit Projek", http.StatusForbidden)
	}

	var bsonn map[string]interface{}
	err = json.Unmarshal([]byte(tokenSplit[3]), &bsonn)
	if err != nil {
		panic(err)
	}

	err = c.Update(bson.M{"_id": path}, bson.M{"$set": bsonn})
	if err != nil {
		return ErrorReturn(w, "Gagal Edit Projek", http.StatusBadRequest)
	}

	return SuccessReturn(w, "Berhasil Edit Projek", http.StatusOK)
}

func GetAllProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request) string {
	//localhost:9000/projek/
	var allProjek []Projek
	var ret []byte

	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("projek")

	err := c.Find(nil).All(&allProjek)
	if err != nil {
		return ErrorReturn(w, "Tidak Ada Projek", http.StatusBadRequest)
	}

	ret, err = json.Marshal(allProjek)

	w.WriteHeader(http.StatusOK)
	return string(ret)
}

//Digunakan untuk mengatur path dari Projek (/projek/...)
func ProjekController(url string, w http.ResponseWriter, r *http.Request) string {
	url = url[1:]
	pathe := strings.Split(url, "/")
	fmt.Println(pathe[0] + " " + pathe[1])

	ses, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer ses.Close()
	ses.SetMode(mgo.Monotonic, true)

	if len(pathe) >= 2 {
		if pathe[1] == "upload" {
			return UploadProjek(ses, w, r)
		}
	}
}
