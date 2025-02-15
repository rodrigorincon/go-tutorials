package main

import(
	"fmt"
)

type MarketplaceApi interface{ // publish a product in a marketplace, do all API comunication and threat possible errors
	PublishProduct(product string)(ok bool, error string)
	PullProducts()(ok bool, error string, products []string)
}

type AmazonApi struct{}
type EbayApi struct{}
func (m AmazonApi) PublishProduct(product string)(ok bool, error string){
	ok = true
	error = ""
	return
}
func (m AmazonApi) PullProducts()(ok bool, error string, products []string){
	ok = true
	error = ""
	products = []string{"amazon prod 1","amazon prod 2"}
	return
}

func (m EbayApi) PublishProduct(product string)(ok bool, error string){
	ok = true
	error = ""
	return
}
func (m EbayApi) PullProducts()(ok bool, error string, products []string){
	ok = true
	error = ""
	products = []string{"example1","example2"}
	return
}

type MarketplaceService struct{
	products []string
	api MarketplaceApi  // dependency injection starts here
}
func (m *MarketplaceService) update(){
	// send all local products to the marketplace
	for _, prod := range(m.products) {
		ok, err := m.api.PublishProduct(prod) // using depedency injection here
		if(!ok){
			fmt.Println("error publishing product:", err)
			return
		}
	}
	// gets new products in marketplace that we don't have locally
	ok, err, products := m.api.PullProducts() // using depedency injection here
	if(!ok){
		fmt.Println("error pulling products:", err)
		return
	}
	m.products = products
}

func main(){
	// creates an api communicator
	var api MarketplaceApi
	api = AmazonApi{}

	// creates a service
	service := MarketplaceService{
		api: api, // attach the interface implementation
		products: []string{"ex1", "ex2"},
	}

	service.update()
	fmt.Println("new prods: ", service.products)
}