# zhipuai_go Project
ZhipuaiGo is a Go-based application that provides an interface to interact with a Language Model API. This project is designed to be lightweight, efficient, and easy to use for handling text-based requests and generating responses.
## Features
- Integration with a Language Model API for generating responses.
- JSON-based request and response format.
- Easy configuration through an INI file.
## Requirements
- Go 1.16 or later.
- A valid API key for the Language Model API.
## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/vvxf/zhipuai_go.git
    cd zhipuai_go
    ```
2. Install dependencies:
    ```sh
    go mod tidy
    ```
3. Configure the application by creating a `config.ini` file in the `conf` directory with the following content:
    ```ini
    [api]
    url = "YOUR_API_URL"
    key = "YOUR_API_KEY"
    ```
## Usage
To run the application and test the local function:
1. Build the application:
    ```sh
    go build
    ```
2. Run the application:
    ```sh
    ./zhipuai_go
    ```
3. The application will execute a sample request to the Language Model API and print the response to the console.
   ```go
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
  }
  ```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
## Contact
For any questions or suggestions, please open an issue on GitHub.
Thank you for using zhipuai_go!