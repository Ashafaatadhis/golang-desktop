package main

// package
import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Barang struct {
	Nama   string	`json:"nama"`
	Harga  int		`json:"harga"`
	Gambar string	`json:"gambar"`
}

// func main
func main() {
	a := app.New()
	w := a.NewWindow("Belanja kuyyy")
	w.Resize(fyne.NewSize(650, 550))
	content := HomePage(w)
	w.SetContent(content)
	w.ShowAndRun()
}

// func hamepage
func HomePage(writer fyne.Window) *fyne.Container {

	welcomeLabel := widget.NewLabelWithStyle("selamat datang di toko belaja kuyyy", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	btn1 := widget.NewButton("Sign In", func() {
		signIn(writer, HomePage)
	})

	btn2 := widget.NewButton("Sign Up", func() {
		signUp(writer, HomePage)

	})

	content := container.New(layout.NewCenterLayout(), container.New(layout.NewGridLayoutWithRows(2), welcomeLabel, container.New(layout.NewGridLayoutWithColumns(2), btn1, btn2)))
	return content

}

// func sign in
func signIn(writer fyne.Window, back func(writer fyne.Window) *fyne.Container) {

	LoginPage(writer, back)

}

// func sign up
func signUp(writer fyne.Window, back func(writer fyne.Window) *fyne.Container) {
	var username string
	var password string

	c := new(fyne.Container)
	c.Layout = layout.NewVBoxLayout()
	usernameLabel := widget.NewLabel("username")
	usernameLabel.Alignment = fyne.TextAlignCenter

	usernameEntry := widget.NewEntry()
	usernameEntry.OnChanged = func(s string) {
		username = s
	}
	passwordLabel := widget.NewLabel("password")
	passwordLabel.Alignment = fyne.TextAlignCenter

	passwordEntry := widget.NewEntry()
	passwordEntry.Password = true
	passwordEntry.OnChanged = func(s string) {
		password = s

	}
	c.Add(usernameLabel)
	c.Add(usernameEntry)
	c.Add(passwordLabel)
	c.Add(passwordEntry)
	btnLay := new(fyne.Container)
	btnLay.Layout = layout.NewGridLayoutWithColumns(2)

	backButton := widget.NewButton("Back", func() {
		writer.SetContent(back(writer))
	})
	submitButton := widget.NewButton("Submit", func() {
		fmt.Print(username, password)
		RegistConfPage(writer, back)
	})

	btnLay.Add(submitButton)
	btnLay.Add(backButton)
	c.Add(btnLay)
	writer.SetContent(c)
}

// Label sign up
func RegistConfPage(writer fyne.Window, back func(writer fyne.Window) *fyne.Container) {
	master := new(fyne.Container)
	master.Layout = layout.NewCenterLayout()
	vb := new(fyne.Container)
	vb.Layout = layout.NewVBoxLayout()
	vb.Add(widget.NewTextGridFromString("Akun anda berhasil dibuat"))
	vb.Add(widget.NewButton("Login", func() {
		LoginPage(writer, back)
	}))
	master.Add(vb)
	writer.SetContent(master)
}

// label dan login page
func LoginPage(writer fyne.Window, back func(writer fyne.Window) *fyne.Container) {

	var username string
	var password string

	c := new(fyne.Container)
	c.Layout = layout.NewVBoxLayout()
	usernameLabel := widget.NewLabel("username")
	usernameLabel.Alignment = fyne.TextAlignCenter

	usernameEntry := widget.NewEntry()
	usernameEntry.OnChanged = func(s string) {
		username = s
	}
	passwordLabel := widget.NewLabel("password")
	passwordLabel.Alignment = fyne.TextAlignCenter

	passwordEntry := widget.NewEntry()
	passwordEntry.Password = true
	passwordEntry.OnChanged = func(s string) {
		password = s
	}
	c.Add(usernameLabel)
	c.Add(usernameEntry)
	c.Add(passwordLabel)
	c.Add(passwordEntry)
	btnLay := new(fyne.Container)
	btnLay.Layout = layout.NewGridLayoutWithColumns(2)

	backButton := widget.NewButton("Back", func() {
		writer.SetContent(back(writer))
	})
	submitButton := widget.NewButton("Submit", func() {

		if username == "admin" && password == "123" {
			IndexPage(writer)
		} else {
			writer.SetContent(back(writer))
		}
	})

	btnLay.Add(submitButton)
	btnLay.Add(backButton)
	c.Add(btnLay)
	writer.SetContent(c)

}

// BELUM KELAR
func CheckoutPage(writer fyne.Window) {
	
	master := new(fyne.Container)
	subMaster := new(fyne.Container)
	subMaster.Layout = layout.NewVBoxLayout()

	master.Layout = layout.NewCenterLayout()
	con := container.NewCenter()
	out := canvas.NewText("Terima kasih telah Berbelanja", color.Black)
	btn := widget.NewButton("Lanjut" , func ()  {
		IndexPage(writer)
	})

	out.Alignment = fyne.TextAlignCenter
	out.TextStyle = fyne.TextStyle{Bold: true}
	out.TextSize = 30
	con.Add(btn)
	
	subMaster.Add(out)
	subMaster.Add(con)
	master.Add(subMaster)
	
	writer.SetContent(master)

}

func PilihBarang(writer fyne.Window){
	
	var barang []Barang
	var listBarang []string
	var namabarang string
	var hargabarang int

	ReadJsonToStruct(&barang)

	for _,i := range barang {
		listBarang = append(listBarang, i.Nama)
	}

	c := new(fyne.Container)
	c.Layout = layout.NewVBoxLayout()
	namabarangLabel := widget.NewLabel("nama barang")
	namabarangLabel.Alignment = fyne.TextAlignCenter

	hargabarangLabel := widget.NewLabel("harga")
	hargabarangLabel.Alignment = fyne.TextAlignCenter

	hargabarangEntry := widget.NewEntry()
	hargabarangEntry.OnChanged = func(s string) {
		namabarang = s
	}


	namabarangEntry := widget.NewSelect(listBarang, func(s string) {
		// buat variabel harga
		var harga int

		// loooping struct barang
		for _, i := range barang {
			if i.Nama == s {
				harga = i.Harga
			}
		}

		hargabarangEntry.SetText(strconv.Itoa(harga))
	 
	})

	lanjut := widget.NewButton("Lanjut", func() {
		CheckoutPage(writer)
	})

	back := widget.NewButton("Back", func() {
		IndexPage(writer)
	})

	layoutKananKiri := container.NewGridWithColumns(2)
	layoutKananKiri.Add(back)
	layoutKananKiri.Add(lanjut)
	

	fmt.Printf(namabarang, hargabarang)
	c.Add(namabarangLabel)
	c.Add(namabarangEntry)
	c.Add(hargabarangLabel)
	c.Add(hargabarangEntry)
	c.Add(layoutKananKiri)
 

	writer.SetContent(c)

}



// halaman utama
func IndexPage(writer fyne.Window) {
	// Isi barang
	var barang []Barang

	 
	ReadJsonToStruct(&barang)
	

	welcome := canvas.NewText("Selamat datang di Belanja Kuyy", color.Black)
	welcome.Alignment = fyne.TextAlignCenter
	welcome.TextStyle = fyne.TextStyle{Bold: true}
	welcome.TextSize = 23
	mas := container.NewVBox()
	subMas := container.NewGridWithColumns(2)
	mas.Add(welcome)
	c := new(fyne.Container)
	c.Layout = layout.NewGridLayoutWithColumns(2)

	for _, b := range barang {
		img := canvas.NewImageFromFile(b.Gambar)
		a := container.NewGridWrap(fyne.NewSize(250, 250))
		a.Add(img)
		c.Add(widget.NewCard(b.Nama, "Rp"+strconv.Itoa(b.Harga), a))
	}

	checkout := widget.NewButton("Checkout", func() {
		PilihBarang(writer)
	})
	back := widget.NewButton("Back", func() {		
		writer.SetContent(HomePage(writer))
	})

	subMas.Add(back)
	subMas.Add(checkout)

	mas.Add(c)
	mas.Add(subMas)
	master := container.NewVScroll(mas)
	writer.SetContent(master)
}

func ReadJsonToStruct(barang *[]Barang) {
	content, err := ioutil.ReadFile("./db.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(content, barang)
}