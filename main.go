package main

// package
import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Barang struct {
	nama   string
	harga  int
	gambar string
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

	// var namabarang string
	// var hargabarang int

	out := canvas.NewText("Terima kasih telah Berbelanja", color.Black)
	out.Alignment = fyne.TextAlignCenter
	out.TextStyle = fyne.TextStyle{Bold: true}
	out.TextSize = 30
	writer.SetContent(out)

	// c := new(fyne.Container)
	// c.Layout = layout.NewVBoxLayout()
	// namabarangLabel := widget.NewLabel("nama barang")
	// namabarangLabel.Alignment = fyne.TextAlignCenter

	// namabarangEntry := widget.NewEntry()
	// namabarangEntry.OnChanged = func(s string) {
	// 	namabarang = s
	// }
	// hargabarangLabel := widget.NewLabel("harga")
	// hargabarangLabel.Alignment = fyne.TextAlignCenter

	// hargabarangEntry := widget.NewEntry()
	// hargabarangEntry.OnChanged = func(s string) {
	// 	namabarang = s
	// }
	// c.Add(namabarangLabel)
	// c.Add(namabarangEntry)
	// c.Add(hargabarangLabel)
	// c.Add(namabarangEntry)

}

// halaman utama
func IndexPage(writer fyne.Window) {
	// Isi barang
	barang := []Barang{

		{
			nama:   "Led TV",
			harga:  10000000,
			gambar: "./tv.jpg",
		},
		{
			nama:   "Monitor",
			harga:  1500000,
			gambar: "./monitor.png",
		},
		{
			nama:   "Laptop",
			harga:  7500000,
			gambar: "./laptop.jpeg",
		},
		{
			nama:   "Meja Gaming",
			harga:  2000000,
			gambar: "./meja.jpg",
		},
		{
			nama:   "Mouse Gaming",
			harga:  98000,
			gambar: "./mouse.png",
		},
		{
			nama:   "mechanical keyboard",
			harga:  1000000,
			gambar: "./keyboard.png",
		},
		{
			nama:   "AC",
			harga:  15000000,
			gambar: "./AC.png",
		},
		{
			nama:   "Panci",
			harga:  150000,
			gambar: "./panci.jpeg",
		},
		{
			nama:   "Masker",
			harga:  30000,
			gambar: "./masker.jpg",
		},
		{
			nama:   "Senter",
			harga:  80000,
			gambar: "./senter.jpg",
		},
		{
			nama:   "Kulkas",
			harga:  2000000,
			gambar: "./kulkas.jpeg",
		},
		{
			nama:   "Kipas Angin",
			harga:  250000,
			gambar: "./kipas.jpg",
		},
	}
	welcome := canvas.NewText("Selamat datang di Belanja Kuyy", color.Black)
	welcome.Alignment = fyne.TextAlignCenter
	welcome.TextStyle = fyne.TextStyle{Bold: true}
	welcome.TextSize = 23
	mas := container.NewVBox()
	mas.Add(welcome)
	c := new(fyne.Container)
	c.Layout = layout.NewGridLayoutWithColumns(2)

	for _, b := range barang {
		img := canvas.NewImageFromFile(b.gambar)
		a := container.NewGridWrap(fyne.NewSize(250, 250))
		a.Add(img)
		c.Add(widget.NewCard(b.nama, "Rp"+strconv.Itoa(b.harga), a))
	}

	checkout := widget.NewButton("Checkout", func() {
		CheckoutPage(writer)
	})

	mas.Add(c)
	mas.Add(checkout)
	master := container.NewVScroll(mas)
	writer.SetContent(master)
}
