package main

import (
	"log"
	"os"

	"github.com/brunomvsouza/ynab.go"
	"github.com/joho/godotenv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("YNAB_ACCESS_TOKEN")
	hotDogBudgetId := os.Getenv("YNAB_BUDGET_ID")

	y := ynab.NewClient(accessToken)
	b, err := y.Budget().GetBudget(hotDogBudgetId, nil)
	if err != nil {
		panic(err)
	}
	c, err := y.Category().GetCategories(hotDogBudgetId, nil)
	if err != nil {
		panic(err)
	}

	for _, account := range b.Budget.Accounts {
		if !account.Closed {
			b := float64(account.Balance)
			p := message.NewPrinter(language.English)
			p.Printf("%v: $%v\n", account.Name, b/1000)
		}
	}

	for _, category := range c.GroupWithCategories {
		if category.Name == "Spending" || category.Name == "True Expenses" {
			for _, topic := range category.Categories {
				b := float64(topic.Balance)
				p := message.NewPrinter(language.English)
				p.Printf("%v: $%v\n", topic.Name, b/1000)
			}
		}
	}
}
