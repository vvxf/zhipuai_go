# zhipuai_go Project
zhipuai_go is a Go-based application that provides an interface to interact with glm Language Model API. This project is designed to be lightweight, efficient, and easy to use for handling text-based requests and generating responses.
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
```go
package main

import (
    "fmt"
    "github.com/vvxf/zhipuai_go/api"
)

func main() {
    apiURL := "https://open.bigmodel.cn/api/paas/v4/chat/completions"
    apiKey := ""

    // init application
    appService := api.NewLLMApplicationService(apiURL, apiKey)

    resp, _ := appService.HandleRequest("glm-4-flash", []api.Message{
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
For any questions or suggestions, please open an issue on GitHub or reach out to us at `zhipuaigo @ outlook.com`.
Thank you for using zhipuai_go!
