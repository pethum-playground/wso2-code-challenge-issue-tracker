package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")
	gptUrl := os.Getenv("GPT_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	req, err := json.Marshal(map[string]string{"message": "This is an issue tracker application. This issue type is an software issue give me a solution for this issue. Issue is there is an application which is run payroll in react. I cannot run it. Give answer in json"})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(gptUrl+"/v1/chat", "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	gptResp := []GptResponse{}
	json.Unmarshal(body, &gptResp)
	fmt.Println(gptResp)

	user := Application{Name: "app.analyze", Value: gptResp[0].Message.Content}
	db.Create(&user)
}
