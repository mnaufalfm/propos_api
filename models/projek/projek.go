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

	"../../auth"
	"../user"
)

type Donatur struct {
	IdDonatur    string `json:"iddonatur"`
	JumlahDonasi string `json:"jumlahdonasi"`
}

type Komentator struct {
	IdKomentator string `json:"idkomentator"`
	Komen        string `json:"komen"`
	TanggalKomen string `json:"tanggalkomen"`
}

type PaketDonasi struct {
	Id           int    `json:"id"`
	NamaPaket    string `json:"namapaket"`
	JumlahDonasi string `json"jumlahdonasi"`
	Apresiasi    string `json:"apresiasi"`
}

type Projek struct {
	Id                string        `json:"id" bson:"_id,omitempty"`
	NamaProjek        string        `json:"namaprojek"`
	FotoProjek        []string      `json:"fotoprojek"` //simpan alamatnya saja
	LinkYoutube       string        `json:"linkyoutube"`
	Deadline          string        `json:"deadline"`
	Donasi            []PaketDonasi `json:"donasi"`
	Tagline           []string      `json:"tagline"`
	Kategori          []string      `json:"kategori"`
	PenjelasanSingkat string        `json:"penjelasansingkat"`
	IdPemilik         string        `json:"idpemilik"`
	IdAnggota         []string      `json:"idanggota"`
	ParaDonatur       []Donatur     `json:"paradonatur"`
	ParaKomen         []Komentator  `json:"parakomen"`
	IdLikers          []string      `json:"idlikers"`
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

func GetProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request, path string) string {
	//linknya:9000/projek/idprojek
	var ret Projek
	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("projek")

	err := c.Find(bson.M{"_id": path}).One(&ret)
	if err != nil {
		return ErrorReturn(w, "Projek Tidak Ada", http.StatusBadRequest)
	}

	rett, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	return string(rett)
}

func GetAllProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request) string {
	//linknya:9000/projek/
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

func LikeProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request, pathidprojek string) string {
	//linknya:9000/like/idprojek
	//var like []string
	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("projek")

	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ErrorReturn(w, "Token Tidak Ditemukan", http.StatusBadRequest)
	}
	token := string(resBody)

	if !jwt.CheckToken(token) {
		return ErrorReturn(w, "Token yang Dikirimkan Invalid", http.StatusBadRequest)
	}

	//err = c.Find(bson.M{"_id": bson.ObjectIdHex(pathidprojek)}).One(&like)

	iduser := strings.Split(token, ".")[1]
	iduserr := jwt.Base64ToString(iduser)
	err = c.Update(bson.M{"_id": bson.IsObjectIdHex(pathidprojek)}, bson.M{"$push": bson.M{"": iduserr}})
	if err != nil {
		return ErrorReturn(w, "Gagal Melakukan Proses Like", http.StatusInternalServerError)
	}

	return SuccessReturn(w, "Anda Berhasil Menge-Like Projek Ini", http.StatusOK)
}

func CommentProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request, idprojek string) string {
	//linknya:9000/comment/idprojek
	ses := s.Copy()
	defer ses.Close()

	c := ses.DB("propos").C("projek")

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
