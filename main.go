package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const smsPrice float32 = 0.034

func calculatePrice(smsSent int) float32 {
	endPrice := float32(smsSent) * smsPrice

	return endPrice
}

func main(){
	fmt.Println("Type how much SMS you've sent:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error #1")
		return
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	value, err := strconv.ParseFloat(input, 32)

	if err != nil {
		fmt.Println(err)
		return
	}

	var priceToCalculate int = int(value)
	var calculated float32 = calculatePrice(int(priceToCalculate))

	fmt.Println("Your SMS bill is: $" + fmt.Sprintf("%f", calculated))
}
