package consolemanager

import "fmt"

type ConsoleManager struct {
}

func New() *ConsoleManager {
	return &ConsoleManager{}
}

func (cmd ConsoleManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices.")
	var prices []string

	for {
		var price string
		fmt.Println("Price: ")
		_, err := fmt.Scanln(&price)
		if err != nil {
			return nil, err
		}
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd ConsoleManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
