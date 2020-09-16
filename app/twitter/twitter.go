package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetTwitterApi() *anaconda.TwitterApi {
	loadEnv()
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
