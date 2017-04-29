package main

import (
	"encoding/json"
	"fmt"
)

/*type Rekening struct {
	NoRekening string `json:"norekening"`
	AtasNama   string `json:"atasnama"`
	Bank       string `json:"bank"`
}

type Pengguna struct {
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
}*/

type Biasa struct {
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
}

type Berisik struct {
	Maklumi string  `json:"maklumi"`
	Haha    string  `json:"haha"`
	Hihi    []Biasa `json:"hihi"`
}

func main() {
	var berisik Berisik

	jsonhaha := []byte(`{"maklumi":"cacat","haha":"blabla","hihi":[{"nama":"william","kelas":"5A"},{"nama":"lemah","kelas":"6B"}]}`)
	//jsonblob := "{username: 'williamhanugra', pass: 'ganteng123', fotoprofil: 'blabla.jpg',nama: 'Lu William Hanugra',iddiri: '135060700111084',jenisid: 1,tgllahir: '14 April 2017',norek: [{norekening:'44444',atasnama:'William Hanugra',bank:'IPB Syariah'}],email: 'cipatonthesky@gmail.com',gender: 'L',nohp: '087873766464',alamat: 'Pondok Bu Sri'}"

	//fmt.Println(json)
	err := json.Unmarshal(jsonhaha, &berisik)
	if err != nil {
		fmt.Println("Gagal coy")
	}
	fmt.Printf("%+v", berisik)
	fmt.Println()
	fmt.Printf("%+v", berisik.Hihi[0])

	/*type ColorGroup struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}*/
	/*var b Pengguna
	err := json.NewDecoder([]byte(jsonblob)).Decode(&b)
	//var cg ColorGroup
	//a := "{status: 1, message: 'Reds'}"
	//err := json.NewDecoder(fmt.Scan()).Decode(&cg)
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)
	fmt.Println(b.Username)*/

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
}
