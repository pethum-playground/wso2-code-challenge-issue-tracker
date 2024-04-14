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
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")
	gptUrl := os.Getenv("GPT_URL")
	message := os.Getenv("MESSAGE")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&Application{})
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := json.Marshal(map[string]string{"message": message})
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

	id, err := strconv.Atoi(os.Getenv("ANALYZE_ID"))
	lastDateId, err := strconv.Atoi(os.Getenv("ANALYZE_LAST_DATE_ID"))
	analyze := Application{Id: id, Name: "app.analyze", Value: gptResp[0].Message.Content}
	analyzeDate := Application{Id: lastDateId, Name: "app.last_analyze_date", Value: time.Now().Local().Format("2006-01-02")}
	db.Save(&analyze)
	db.Save(&analyzeDate)
}
