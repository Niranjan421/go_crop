# Go Crop Prediction Application

This is a Go implementation of the Crop Prediction Application originally written in Python using Streamlit. The Go version uses the Gin web framework to create a web application with similar functionality.

## Project Structure

```
go_crop_prediction/
├── main.go              # Main application code
├── static/              # Static assets
│   ├── css/
│   │   └── styles.css   # CSS styling
│   └── img/             # Images (created at runtime)
└── templates/           # HTML templates
    ├── home.html        # Home page
    ├── prediction.html  # Prediction form
    ├── result.html      # Results page
    └── error.html       # Error page
```

## Features

- Multi-page web application with navigation
- Home page with project information
- Prediction form for entering environmental conditions
- Result page showing the predicted crop
- Error handling
- Responsive design

## Technical Details

- Uses Gin web framework for routing and HTML rendering
- Simulates ML model prediction (in a real application, would connect to a Python microservice)
- Handles form submission and validation
- Serves static files (CSS, images)
- Creates placeholder images for testing

## How to Run

1. Install Go (if not already installed)
2. Install dependencies:
   ```
   go get -u github.com/gin-gonic/gin
   ```
3. Navigate to the project directory:
   ```
   cd go_crop_prediction
   ```
4. Run the application:
   ```
   go run main.go
   ```
5. Open a web browser and go to:
   ```
   http://localhost:8080
   ```

## Implementation Notes

- The machine learning model prediction is simulated since the original uses a joblib-loaded model
- In a production environment, you would either:
  - Create a Python microservice to handle predictions
  - Convert the model to a format usable in Go
  - Use ONNX for cross-language model compatibility
- The application creates placeholder images at runtime for testing purposes

## Differences from the Python Version

- Uses a traditional web application architecture instead of Streamlit's reactive approach
- Requires explicit routing and HTML templates
- Form handling is done via HTTP POST instead of Streamlit's state management
- Styling is implemented with CSS instead of Streamlit's built-in styling
