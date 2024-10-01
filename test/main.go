package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"os"
	"github.com/fogleman/gg"
)

func SendCircle() {

	url := "http://localhost:8080"
    dc := gg.NewContext(1000, 1000)
    dc.DrawCircle(500, 500, 400)
    dc.SetRGB(1, 1, 1)
    dc.Fill()


	var buf bytes.Buffer

	png.Encode(&buf, dc.Image())
	resp, err := http.Post(url, "image/png", &buf)

	if err != nil {
		fmt.Println("Error sending image:", err)
		return
	}
	defer resp.Body.Close()

}


func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Check that the request is a POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Create a file to save the uploaded image
	out, err := os.Create("uploaded_image.png")
	if err != nil {
		http.Error(w, "Unable to create the file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the image data from the request body to the file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		http.Error(w, "Unable to save the image", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Respond back to the client
	fmt.Fprintf(w, "Image uploaded successfully")
}

func main() {
	// Route that handles the image upload
	http.HandleFunc("/submit", submitHandler)

	// Start the server on localhost:8080
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
	SendCircle()
}