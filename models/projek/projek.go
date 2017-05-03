package projek

import (
	"fmt"
	"net/http"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type Projek struct {
	Id                bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	NamaProjek        string          `json:"namaprojek"`
	FotoProjek        []string        `json:"fotoprojek"` //simpan alamatnya saja
	PenjelasanSingkat string          `json:"penjelasansingkat"`
	IdPemilik         bson.ObjectId   `json:"idpemilik"`
	IdAnggota         []bson.ObjectId `json:"idanggota"`
	ParaDonatur       []Donatur       `json:"paradonatur"`
	ParaKomen         []Komentator    `json:"parakomen"`
	IdLikers          []bson.ObjectId `json:"idlikers"`
}

func UploadProjek(s *mgo.Session, w http.ResponseWriter, r *http.Request) string {
	
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
