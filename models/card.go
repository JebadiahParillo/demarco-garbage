package models

// The card model is used to represent card info

type Card struct {
	CardNumber     string
	NameOnCard     string
	CVV            string
	ExpirationDate string
	BillingAddress string
	BillingZip     string
	BillingState   string
}

// NewCard serves as a constructor for a card object. This object can be used to store input data from the DB or a frontend request
func NewCard(
	cardNumber string, nameOnCard string, cvv string,
	expirationDate string, billingAddress string, billingZip string,
	billingState string,
) *Card {

	return &Card{
		CardNumber:     cardNumber,
		NameOnCard:     nameOnCard,
		CVV:            cvv,
		ExpirationDate: expirationDate,
		BillingAddress: billingAddress,
		BillingZip:     billingZip,
		BillingState:   billingState,
	}

}
