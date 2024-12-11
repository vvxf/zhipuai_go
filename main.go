package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"zhipuai_go/application"
	"zhipuai_go/domain"

	"github.com/spf13/viper"
)

func initConfig() {
	// Set the name of the config file (without extension)
	viper.SetConfigName("config") // If the filename is config.conf, just specify "config"

	// Add configuration file path (optional, if the config file is in the same directory as the program, this is not needed)
	viper.AddConfigPath("./conf")

	// Set the configuration file type to "ini"
	viper.SetConfigType("ini")

	// Enable automatic environment variable parsing
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func getAPIURL() string {
	return viper.GetString("api.url")
}

// func getServerPort() int {
// 	return viper.GetInt("server.port")
// }

func getAPIKey() string {
	return viper.GetString("api.key")
}

// RequestHandler handles incoming HTTP requests.
func RequestHandler(appService *application.LLMApplicationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			Model    string           `json:"model"`
			Messages []domain.Message `json:"messages"`
		}

		// Parse the incoming JSON data
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Call the application service to handle the business logic
		response, err := appService.HandleRequest(request.Model, request.Messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set response headers and write the response body
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	initConfig()

	apiURL := getAPIURL()
	apiKey := getAPIKey()

	// 初始化应用层服务
	appService := application.NewLLMApplicationService(apiURL, apiKey)

	resp, _ := appService.HandleRequest("glm-4-flash", []domain.Message{
		{
			Role:    "user",
			Content: "Hello, llm!",
		},
	})

	fmt.Println(resp.Choices[0].Message.Content)

	// Create an HTTP server
	// serverPort := getServerPort()
	// fmt.Printf("Server will run on port: %d\n", serverPort)

	// http.HandleFunc("/llm/generate", RequestHandler(appService))

	// fmt.Printf("Server is running on port %d...\n", serverPort)
	// err := http.ListenAndServe(":"+fmt.Sprint(serverPort), nil)
	// if err != nil {
	// 	fmt.Printf("Error starting server: %v\n", err)
	// }
}
