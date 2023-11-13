package main

// import package

// elaina cantik jir
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
	nama  string
	harga int
}

func main() {
	a := app.New()
	w := a.NewWindow("Belanja kuyyy")
	w.Resize(fyne.NewSize(450, 400))
	content := HomePage(w)
	w.SetContent(content)
	w.ShowAndRun()
}

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

func signIn(writer fyne.Window, back func(writer fyne.Window) *fyne.Container) {

	LoginPage(writer, back)

}

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

func IndexPage(writer fyne.Window) {
	// Isi barang
	barang := []Barang{

		{
			nama:  "Kipas",
			harga: 10000,
		},
		{
			nama:  "Kayu",
			harga: 20200,
		},
		{
			nama:  "Pensil",
			harga: 20200,
		},
		{
			nama:  "TV",
			harga: 20200,
		},
		{
			nama:  "Bantal Waifu",
			harga: 20200,
		},
		{
			nama:  "Nasi padang",
			harga: 10000,
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
		img := canvas.NewImageFromFile("./bang.png")
		a := container.NewGridWrap(fyne.NewSize(250, 250))
		a.Add(img)
		c.Add(widget.NewCard(b.nama, strconv.Itoa(b.harga), a))
	}

	mas.Add(c)
	master := container.NewVScroll(mas)
	writer.SetContent(master)
}
