package backend

import (
	"app-bazaar/constants"
	"fmt"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/price"
	"github.com/stripe/stripe-go/v78/product"
)

func CreateProductWithPrice(appTitle string, appDescription string, appPrice int64) (productID, priceID string, err error){
	stripe.Key = constants.STRIPE_API_KEY
	product_params := &stripe.ProductParams{
		Name: stripe.String(appTitle),
		Description: stripe.String(appDescription),
	}
	
	newProduct, err := product.New(product_params)
	if err != nil{
		fmt.Println("Failted to create product" + err.Error())
		return "", "", err
	}

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product: stripe.String(newProduct.ID),
		UnitAmount: stripe.Int64(appPrice),
	}
	
	newPrice, err := price.New(price_params)
	if err != nil{
		fmt.Println("Failed to create price" + err.Error())
		return "", "", err
	}
	
	return newProduct.ID, newPrice.ID, nil
} 

func CreateCheckoutSession(domain string, priceID string) (*stripe.CheckoutSession, error){
	stripe.Key = constants.STRIPE_API_KEY

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price: &priceID,
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL: stripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)

	if err != nil{
		fmt.Printf("session.New: %v", err)
		return nil, err
	}
	return s, nil
}