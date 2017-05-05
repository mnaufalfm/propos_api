package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"encoding/base64"
	"encoding/hex"
	"reflect"
)

type Rekening struct {
	NoRekening string `json:"norekening"`
	AtasNama   string `json:"atasnama"`
	Bank       string `json:"bank"`
}

type Pengguna struct {
	Id         string     `json:"id" bson:"_id,omitempty"`
	Username   string     `json:"username"`
	Password   string     `json:"pass"`
	FotoProfil string     `json:"fotoprofil"` //simpan alamatnya saja
	Nama       string     `json:"nama"`
	IdDiri     string     `json:"iddiri"`
	JenisID    int        `json:"jenisid"` //1=KTP, 2=SIM, 3=Paspor
	TglLahir   string     `json:"tgllahir"`
	Norek      []Rekening `json:"norek"`
	Email      string     `json:"email"`
	Gender     string     `json:"gender"`
	NoHp       string     `json:"nohp"`
	Alamat     string     `json:"alamat"`
}

/*type Rekening struct {
	NoRekening string `json:"norekening,omitempty"`
	AtasNama   string `json:"atasnama,omitempty"`
	Bank       string `json:"bank,omitempty"`
}

type Pengguna struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username   string        `json:"username,omitempty"`
	Password   string        `json:"password,omitempty"`
	FotoProfil string        `json:"fotoprofil,omitempty"` //simpan alamatnya saja
	Nama       string        `json:"nama,omitempty"`
	IdDiri     string        `json:"iddiri,omitempty"`
	JenisID    int           `json:"jenisid,omitempty"` //1=KTP, 2=SIM, 3=Paspor
	TglLahir   string        `json:"tgllahir,omitempty"`
	Norek      []Rekening    `json:"norek,omitempty"`
	Email      string        `json:"email,omitempty"`
	Gender     string        `json:"gender,omitempty"`
	NoHp       string        `json:"nohp,omitempty"`
	Alamat     string        `json:"alamat,omitempty"`
}*/

type Biasa struct {
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
}

type Berisik struct {
	Maklumi int     `json:"maklumi"`
	Haha    string  `json:"haha"`
	Hihi    []Biasa `json:"hihi"`
}

func main() {
	//var berisik Berisik

	//jsonhaha := []byte(`{"maklumi":5,"haha":"blabla"}`)
	//jsonblob := "{username: 'williamhanugra', pass: 'ganteng123', fotoprofil: 'blabla.jpg',nama: 'Lu William Hanugra',iddiri: '135060700111084',jenisid: 1,tgllahir: '14 April 2017',norek: [{norekening:'44444',atasnama:'William Hanugra',bank:'IPB Syariah'}],email: 'cipatonthesky@gmail.com',gender: 'L',nohp: '087873766464',alamat: 'Pondok Bu Sri'}"

	//fmt.Println(json)
	//err := json.Unmarshal(jsonhaha, &berisik)
	//if err != nil {
	//	fmt.Println("Gagal coy")
	//}
	//fmt.Printf("%+v", berisik)

	//bb := reflect.ValueOf(&berisik).Elem()
	//for i := 0; i < bb.NumField(); i++ {
	//	fmt.Printf("%s %s %v\n", bb.Type().Field(i).Name, bb.Type(), bb.Field(i).Interface())
	//}
	//fmt.Println()
	//fmt.Printf("%+v", berisik.Hihi[0])

	/*type ColorGroup struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	//var b Pengguna
	//err := json.NewDecoder([]byte(jsonblob)).Decode(&b)
	var cg ColorGroup
	a := []byte(`{"status": 1, "message": "Reds"}`)
	err := json.Unmarshal(a, cg)
	if err != nil {
		fmt.Println("error:", err)
	}
	bb := reflect.ValueOf(&cg).Elem()
	for i := 0; i < bb.NumField(); i++ {
		fmt.Println(bb.Field(i).Interface())
	}*/
	//os.Stdout.Write(b)

	/*s := "/guru/edit"
	s = s[1:]
	fmt.Println(s)*/

	/*pesan := 5
	haha := "Walah to"
	fmt.Printf("{message: %d, haha: %q}", pesan, haha)*/

	//sum := sha256.Sum256([]byte("hello world\n"))
	//fmt.Println(fmt.Sprintf("%x", sum))
	//fmt.Printf("%x", sum)

	/*var sum string
	sum = fmt.Sprintf("%x", sha256.Sum256([]byte("hello world\n")))
	fmt.Println(sum)*/

	var user Pengguna
	//http.ListenAndServe(":9000", nil)
	ses, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	defer ses.Close()
	ses.SetMode(mgo.Monotonic, true)

	c := ses.DB("propos").C("user")

	err = c.Find(bson.M{"username": "willywilly"}).One(&user)
	if err != nil {
		panic(err)
	}
	//gan, erro := json.Marshal(user.Id)
	//if erro != nil {
	//	panic(err)
	//}
	//fmt.Println(string(gan))
	//a, _ := hex.DecodeString(user.Id.Hex())
	/*b := base64.StdEncoding.EncodeToString([]byte(user.Id))
	fmt.Println(b)
	d, _ := base64.StdEncoding.DecodeString("WQl3aW8xBav4+UsI")
	e := string(d)
	f := hex.EncodeToString([]byte(e))
	fmt.Println(f)
	//fmt.Println(jwt.TokenMaker(string([]byte(user.Id)), "anggunauranaufalwilliam"))
	//fmt.Printf("%+v", user)
	//if user.Id.Hex() == f {
	//		fmt.Println("Bacot")
	//	}

	fmt.Printf("%s %s", reflect.TypeOf(user.Id), hex.EncodeToString([]byte(user.Id)))*/
	//	fmt.Println(reflect.TypeOf(user.Id.Hex()))

	//data := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9")
	//a := make([]byte, base64.RawStdEncoding.EncodedLen(len(data)))
	//str := base64.StdEncoding.EncodeToString(data)
	//base64.RawStdEncoding.Encode(a, data)
	//fmt.Println(hex.EncodeToString(a))
	//fmt.Printf("%x", a)
	//fmt.Println(str)

	//kun := []byte("berisikamatlu")
	//pes := []byte(mes)
	//h := hmac.New(sha256.New, kun)
	//h.Write(data)
	//fmt.Println(base64.RawStdEncoding.EncodeToString(h.Sum(nil)))
	//pass := "ganteng123"

	//fmt.Println(fmt.Sprintf("%x", sha256.Sum256([]byte(pass))))

	//fmt.Println(string("184054b78da172c42e37015fb66dd6968b582846f4226c9edfad9da80dc2bf22"))
	//s := append([]string{"1", "2"}, []string{"3", "4"}...)
	//fmt.Println(s)

	//var bsonn map[string]interface{}
	//err = json.Unmarshal([]byte(s), &bsonn)
	//fmt.Println(bsonn)
}
