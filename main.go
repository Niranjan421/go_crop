package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CropDict maps crop IDs to crop names
var CropDict = map[int]string{
	1: "Rice", 2: "Maize", 3: "Jute", 4: "Cotton", 5: "Coconut", 6: "Papaya",
	7: "Orange", 8: "Apple", 9: "Muskmelon", 10: "Watermelon", 11: "Grapes",
	12: "Mango", 13: "Banana", 14: "Pomegranate", 15: "Lentil", 16: "Blackgram",
	17: "Mungbean", 18: "Mothbeans", 19: "Pigeonpeas", 20: "Kidneybeans",
	21: "Chickpea", 22: "Coffee",
}

// For this implementation, we'll simulate the ML model prediction
// In a real application, you would call a Python microservice or use a Go ML library
func predictCrop(features []float64) (int, error) {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	// This is a placeholder for the actual ML prediction
	// In a real implementation, you would:
	// 1. Call a Python API that loads the joblib model
	// 2. Or use a Go ML library if the model is compatible
	
	// For demonstration, we'll return a mock prediction based on input values
	// This is just for simulation purposes
	sum := 0.0
	for _, v := range features {
		sum += v
	}
	
	// Simple logic to determine crop based on sum of features
	// This is NOT a real prediction, just a placeholder
	cropID := int(sum) % 22
	if cropID == 0 {
		cropID = 1
	}
	
	return cropID, nil
}

// imageToBase64 converts an image file to base64 encoding
func imageToBase64(imagePath string) (string, error) {
	// Read the image file
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}
	
	// Encode to base64
	base64Encoding := base64.StdEncoding.EncodeToString(imageBytes)
	return base64Encoding, nil
}

func main() {
	// Create the Gin router
	r := gin.Default()
	
	// Load HTML templates
	r.LoadHTMLGlob("templates/*")
	
	// Serve static files
	r.Static("/static", "./static")
	
	// Create the img directory if it doesn't exist
	if err := os.MkdirAll("static/img", os.ModePerm); err != nil {
		log.Fatal("Failed to create image directory:", err)
	}
	
	// Create placeholder images for testing
	createPlaceholderImages()
	
	// Home page route
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Crop Prediction Application",
		})
	})
	
	// Prediction page route
	r.GET("/prediction", func(c *gin.Context) {
		c.HTML(http.StatusOK, "prediction.html", gin.H{
			"title": "Crop Prediction Application",
		})
	})
	
	// Handle prediction form submission
	r.POST("/predict", func(c *gin.Context) {
		// Parse form values
		nitrogenStr := c.PostForm("nitrogen")
		phosphorusStr := c.PostForm("phosphorus")
		potassiumStr := c.PostForm("potassium")
		temperatureStr := c.PostForm("temperature")
		humidityStr := c.PostForm("humidity")
		phStr := c.PostForm("ph")
		rainfallStr := c.PostForm("rainfall")
		
		// Convert strings to float64
		nitrogen, _ := strconv.ParseFloat(nitrogenStr, 64)
		phosphorus, _ := strconv.ParseFloat(phosphorusStr, 64)
		potassium, _ := strconv.ParseFloat(potassiumStr, 64)
		temperature, _ := strconv.ParseFloat(temperatureStr, 64)
		humidity, _ := strconv.ParseFloat(humidityStr, 64)
		ph, _ := strconv.ParseFloat(phStr, 64)
		rainfall, _ := strconv.ParseFloat(rainfallStr, 64)
		
		// Create features array
		features := []float64{nitrogen, phosphorus, potassium, temperature, humidity, ph, rainfall}
		
		// Get prediction
		cropID, err := predictCrop(features)
		if err != nil {
			c.HTML(http.StatusOK, "error.html", gin.H{
				"title": "Prediction Error",
				"error": "Failed to make prediction",
			})
			return
		}
		
		// Get crop name
		cropName, exists := CropDict[cropID]
		if !exists {
			c.HTML(http.StatusOK, "error.html", gin.H{
				"title": "Prediction Error",
				"error": "No suitable crop found for these conditions",
			})
			return
		}
		
		// Return the prediction result
		c.HTML(http.StatusOK, "result.html", gin.H{
			"title":     "Crop Prediction Result",
			"cropName":  cropName,
			"imagePath": "static\\img\\23.jpg",
		})
	})
	
	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}

// createPlaceholderImages creates placeholder images for testing
func createPlaceholderImages() {
	// Create placeholder images with simple content
	homeImagePath := filepath.Join("static", "img", "3.jpg")
	if _, err := os.Stat(homeImagePath); os.IsNotExist(err) {
		err := ioutil.WriteFile(homeImagePath, []byte("Placeholder for home image"), 0644)
		if err != nil {
			log.Println("Failed to create home image:", err)
		}
	}
	
	successImagePath := filepath.Join("static", "img", "23.jpg")
	if _, err := os.Stat(successImagePath); os.IsNotExist(err) {
		err := ioutil.WriteFile(successImagePath, []byte("Placeholder for success image"), 0644)
		if err != nil {
			log.Println("Failed to create success image:", err)
		}
	}
	
	noResultImagePath := filepath.Join("static", "img", "45.jpg")
	if _, err := os.Stat(noResultImagePath); os.IsNotExist(err) {
		err := ioutil.WriteFile(noResultImagePath, []byte("Placeholder for no result image"), 0644)
		if err != nil {
			log.Println("Failed to create no result image:", err)
		}
	}
}
