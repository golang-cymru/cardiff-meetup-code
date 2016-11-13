package main

import (
	"os"
	"net/http"
	"log"
	"github.com/skratchdot/open-golang/open" /// to open the web browser to the url when running the program
	"github.com/nfnt/resize"
	"image"
	"bytes"
	"reflect"
	"image/color"
	_ "image/png"
	_ "image/jpeg"
	"encoding/json"
)

const asciiWidth = 120
const ASCIISTR = "MND8OZ$7I?+=~:,.."
var responseObj     map[string]interface{}

type GistFile struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string              `json:"description"`
	Public      bool                `json:"public"`
	Files       map[string]GistFile `json:"files"`
}

func main() {
	// Configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	bindIp := os.Getenv("BINDIP")
	if bindIp == "" {
		bindIp = "127.0.0.1"
	}
	serverAddr := bindIp+":"+port
	server := http.Server{
		Addr: serverAddr,
	}
	http.HandleFunc("/", catGist)

	log.Println("Starting web server on ", serverAddr)

	open.Start("http://"+serverAddr)
	server.ListenAndServe()
}

func catGist(w http.ResponseWriter, r *http.Request) {

	log.Println("Request for endpoint")

	response, err := http.Get("http://thecatapi.com/api/images/get?format=src&results_per_page=1&type=jpg")
	if err != nil {
		log.Println("ERROR Retrieving Cat Image", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer response.Body.Close()
	img, _, err := image.Decode(response.Body)
	if err != nil {
		log.Println("ERROR Decoding Cat Image", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	asciidImage := Convert2Ascii(ScaleImage(img, asciiWidth))

	//fmt.Fprintf(w, string(asciidImage)) //if you just want to see the ascii in the browser

	catGist := map[string]GistFile{}
	catGist["cat"] = GistFile{string(asciidImage)}

	gist := Gist{
		"Cat ascii gist",
		true,
		catGist,
	}
	b, err := json.Marshal(gist)
	if err != nil {
		log.Println("ERROR Marshalling gist JSON", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	br := bytes.NewBuffer(b)
	resp, err := http.Post("https://api.github.com/gists", "application/json", br)
	if err != nil {
		log.Println("ERROR Posting Ascii to Github", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&responseObj)
	if err != nil {
		log.Println("ERROR Github JSON response error", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	log.Println("CAT GIST URL GENERATED AT:", responseObj["html_url"])
	http.Redirect(w, r, responseObj["html_url"].(string), http.StatusFound)
	return

}
func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}
func Convert2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}