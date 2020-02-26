package src

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

//RegisterBill registeres a bill from the user with price and product
func RegisterBill(userID, userName, price, product string) (string, error) {

	re := regexp.MustCompile(`clothes$|electronics$|other$|food$`)
	re2 := regexp.MustCompile(`[0-9]+$`)

	//if the parameters match the regex
	if re.MatchString(product) && re2.MatchString(price) {

		bill := Bill{}
		bill.Price = price
		bill.Type = product
		err := Save(&bill, userID)
		if err != nil {
			return ":exclamation: `I was not able to register your bill!`", err
		}
		return ":white_check_mark: `I registered " + product +
			" for " + price + "kr in " + userName + " database!`", nil

	}
	var err error
	return ":exclamation: `You wrote the price or type wrong!`", err
}

//SumAllHandler sums all bills for a user
func SumAllHandler(userID, userName string) string {
	bills, err := ReturnAllBills(userID)
	if err != nil {
		return ":exclamation: `I was not able to get your bills from database!`"
	}
	sum, text := sumAllBills(bills)
	if text != "" {
		return text
	}
	return ":dollar: `" + userName + "'s total is: " + strconv.Itoa(sum) + "kr.`"
}

//SumTypeHandler sums all bills for user based on product type
func SumTypeHandler(userID, userName, productType string) string {
	bills, err := ReturnAllBills(userID)
	if err != nil {
		return ":exclamation: `I was not able to get your bills from database!`"
	}
	sum, text := sumBillType(bills, productType)
	if text != "" {
		return text
	}
	return ":dollar: `" + userName + "'s total for " + productType + " is: " + strconv.Itoa(sum) + "kr.`"
}

//DiagramHandler creates diagram and returns it
func DiagramHandler(userID string) (*os.File, string) {
	bills, err := ReturnAllBills(userID)
	if err != nil {
		return nil, ":exclamation: `I was not able to get your bills from database!`"
	}

	s, f, c, e, o, text := sum(bills)
	if text != "" {
		return nil, text
	}

	foodPercent := Percent(s, f)
	clothesPercent := Percent(s, c)
	electronicsPercent := Percent(s, e)
	otherPercent := Percent(s, o)

	file, text := pieChart(foodPercent, clothesPercent, electronicsPercent, otherPercent)

	return file, text
}

//TotalHandler returns an overview of spendings
func TotalHandler(userID string) string {
	bills, err := ReturnAllBills(userID)
	if err != nil {
		return ":exclamation: `I was not able to get your bills from database!`"
	}

	s, f, c, e, o, text := sum(bills)
	if text != "" {
		return text
	}
	text = totalPrint(strconv.Itoa(s), strconv.Itoa(f), strconv.Itoa(c), strconv.Itoa(e), strconv.Itoa(o))
	return text
}

//sum sums prices of all categories in a single bill
func sum(bills []Bill) (int, int, int, int, int, string) {
	sum, text := sumAllBills(bills)
	if text != "" {
		return 0, 0, 0, 0, 0, text
	}
	food, text := sumBillType(bills, "food")
	if text != "" {
		return 0, 0, 0, 0, 0, text
	}
	clothes, text := sumBillType(bills, "clothes")
	if text != "" {
		return 0, 0, 0, 0, 0, text
	}
	electronics, text := sumBillType(bills, "electronics")
	if text != "" {
		return 0, 0, 0, 0, 0, text
	}
	other, text := sumBillType(bills, "other")
	if text != "" {
		return 0, 0, 0, 0, 0, text
	}
	return sum, food, clothes, electronics, other, ""
}

//sumBillType sums prices of a single category in a single bill
func sumBillType(bills []Bill, productType string) (int, string) {
	var sum int = 0
	//ranging over all the bills
	for v := range bills {
		if bills[v].Type == productType {
			price, err := strconv.Atoi(bills[v].Price)
			if err != nil {
				return 0, ":exclamation: `I failed to sum your bills!`"
			}
			sum += price
		}
	}
	return sum, ""
}

//sumAllBills sums prices of all bills
func sumAllBills(bills []Bill) (int, string) {
	var sum int = 0
	for v := range bills {
		price, err := strconv.Atoi(bills[v].Price)
		if err != nil {
			return 0, ":exclamation: `I failed to sum your bills!`"
		}
		sum += price
	}
	return sum, ""
}

//pieChart creates pie chart
func pieChart(f, c, e, o float64) (*os.File, string) {

	//values for the diagram
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: c, Label: "Clothes"},
			{Value: f, Label: "Food"},
			{Value: e, Label: "Electronics"},
			{Value: o, Label: "Other"},
		},
	}

	//creates the diagram file
	file, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Failed to create diagram png!", err)
		return nil, ""
	}
	err = pie.Render(chart.PNG, file)
	if err != nil {
		fmt.Println("Filed rendering diagram!")
		return nil, ""
	}
	defer file.Close()

	open, _ := os.Open("output.png")
	if err != nil {
		fmt.Println("Could not open png")
		return nil, ""
	}

	//returns the open file
	return open, "output.png"
}

//totalPrint returns the overview of all categories and spendings
func totalPrint(sum, f, c, e, o string) string {

	budget := "**Budget:** \n" +
		">>> :womans_clothes: **Clothes:**  \t\t\t\t " + c + "kr \n" +
		":shallow_pan_of_food: **Food:**  \t\t\t\t\t  " + f + "kr \n" +
		":electric_plug: **Electronics:**  \t\t  " + e + "kr \n" +
		":grey_question: **Other:**  \t\t\t\t\t" + o + "kr \n" +
		" --------------------------------" + "\n" +
		":moneybag: **Sum:**  \t\t\t\t\t  " + sum + "kr \n"

	return budget
}
