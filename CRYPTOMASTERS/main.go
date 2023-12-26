package main

import (
	"cryptomasters/api"
	"fmt"
	"sync"
)

func main(){
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies{
		wg.Add(1)//tell the wait to add one rountine
		//go getCurrencyData(currency)
		//fmt.Println("T")
		//wg.Done() //Remove from the count from wait routine when routine completes the execution
		//wg.Done get executed without waiting for the routine to be done so not a good idea to do it here
		//Solution is a lambda function
		//go func(){
		//	getCurrencyData(currency)//closures work like this, we are getting reference to the currency variable here for some reason like wtf!!! Its getting the currency the last loop because the currency variable changes in the loop while it is being executed
			//To solve this we need to pass currency as the parameter to the function instead of using the argument of loop locally in inside function
		//	wg.Done()
	//	}()

		//Here is the solution

		go func(currency string){
			getCurrencyData(currency)
			wg.Done()
		}(currency)

	}
	wg.Wait()
	//go getCurrencyData("BTC")
	//go getCurrencyData("ETH")
	//go getCurrencyData("BCH")
	//time.Sleep(time.Second*2)//Not a good idea to use sleep to wait better to use something like waitgroup

}

func getCurrencyData(currency string){
	//fmt.Println("O")
	rate, err := api.GetRate(currency)
	if err == nil{
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}
