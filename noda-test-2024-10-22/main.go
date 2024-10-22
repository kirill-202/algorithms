package main


//"https://docs.noda.live/docs/noda-card-api/d693c62fcfbfb-checking-the-payment-status"


import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"github.com/joho/godotenv"
)

type Payment struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	BillingAddress *BillingAddress
	Card *Card
}


type BillingAddress struct {
	AddressLine string
	City string
	State string
	Country string
	Zip string
}

func NewAddress(addr, city, state, country, zip string) *BillingAddress {
	return &BillingAddress{
		AddressLine: addr,
		City: city, 
		State: state,
		Country: country,
		Zip: zip,
	}
}


type Card struct {
	CardNumber string
	Expiration string
	CVV  string
	CardHolder string
	CardID  *string
}


func NewCard(cardNumber, expiration, cvv, cardHolder string, cardID *string) *Card {
	return &Card{
		CardNumber: cardNumber,
		Expiration: expiration,
		CVV:        cvv,
		CardHolder: cardHolder,
		CardID:     cardID,
	}
}

func (p *Payment) CreatePayment(card *Card, billingAddress *BillingAddress, client *http.Client) error {
		p.BillingAddress = billingAddress
		p.Card = card
		p.ID = ""

		jsonData, err := json.Marshal(p); if err != nil {
			return fmt.Errorf("can't marshal json: %v", err)
		}
		fmt.Println("Request Body:", string(jsonData))
		body := bytes.NewBuffer(jsonData)

		req, err := http.NewRequest("POST", "https://api.stage.noda.live/api/payments", body); if err != nil {
			return fmt.Errorf("can't create new request: %v", err)
		}
	
		apiKey := os.Getenv("API_KEY")
		req.Header.Add("x-api-key", apiKey)
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req); if err != nil {
			return fmt.Errorf("issue with getting response: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error response from the server: %v", resp.StatusCode)
		}
		

		var tempP Payment
		byteBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response Body: %s\n", string(byteBody))
		
		err = json.Unmarshal(byteBody, &tempP); if err != nil {
			return fmt.Errorf("can't unmarshal %v", err)
		}

		fmt.Printf("%+v\n", tempP)
		return nil

	}



func main() {
	fmt.Printf("Program has started")
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	paymentId := "763480e5-c2de-4dec-acb0-baee89b8d6b7"

	client := http.DefaultClient

	req, err := http.NewRequest("GET", "https://api.stage.noda.live/api/payments/"+paymentId, nil); if err != nil {
		log.Fatal("Can't create request:", err)
	}
	req.Header.Add("x-api-key", apiKey)

	resp, err := client.Do(req); if err != nil {
		log.Fatal("cant process the request", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("The request failed with status", resp.StatusCode)
	}



	body, err := io.ReadAll(resp.Body); if err != nil {
		log.Fatal("Cant't read response body", err)
	}

	var payment Payment
	err = json.Unmarshal(body, &payment); if err != nil {
		fmt.Println("Error parsing body", err)
	}

	fmt.Printf("%+v\n", payment)

	addr := NewAddress("Golang", "Bangkok", "Thai", "Thail", "101")
	card := NewCard("4253258107133225", "1527", "123", "Boris Bro", nil)

	err = payment.CreatePayment(card, addr, client); if err != nil {
		log.Fatalln("Issue with creating payment: ", err)
	}

}

