package main

import "fmt"

// ===================
// E-COMMERCE
// ===================
func CalculateOrderTotal(amount float64, customerType string) float64 {
	if amount <= 0 {
		return -1
	}
	switch customerType {
	case "regular":
		return amount
	case "premium":
		return amount * 0.9
	case "vip":
		return amount * 0.8
	default:
		return -1
	}
}

// ===================
// BANKING
// ===================
func CalculateInterest(balance float64, accountType string) float64 {
	if balance <= 0 {
		return 0
	}
	switch accountType {
	case "savings":
		return balance * 0.02
	case "fixed":
		return balance * 0.05
	case "checking":
		return balance * 0.005
	default:
		panic("unknown account type")
	}
}

// ===================
// HEALTHCARE
// ===================
func CalculateTreatmentCost(baseCost float64, insured bool, emergency bool) float64 {
	if baseCost < 0 {
		return 0
	}
	cost := baseCost
	if insured {
		cost *= 0.3
		if emergency {
			cost *= 1.5
		}
	}
	if cost < 0 {
		return 0
	}
	return cost
}

// ===================
// LOGISTICS
// ===================
func CalculateShipping(weight float64, express bool, international bool) float64 {
	if weight <= 0 {
		return -1
	}
	price := 5.0
	if weight > 10 {
		price += 10
	}
	if express {
		price += 15
	}
	if international {
		price += 25
	}
	return price
}

// ===================
// CENTRAL DISPATCH
// ===================
func Calculate(service string, params map[string]interface{}) float64 {
	if params == nil {
		panic("params missing")
	}
	switch service {
	case "ecommerce":
		amount := params["amount"].(float64)
		customer := params["customer"].(string)
		return CalculateOrderTotal(amount, customer)
	case "banking":
		balance := params["balance"].(float64)
		account := params["account"].(string)
		return CalculateInterest(balance, account)
	case "healthcare":
		cost := params["cost"].(float64)
		insured := params["insured"].(bool)
		emergency := params["emergency"].(bool)
		return CalculateTreatmentCost(cost, insured, emergency)
	case "logistics":
		weight := params["weight"].(float64)
		express := params["express"].(bool)
		international := params["international"].(bool)
		return CalculateShipping(weight, express, international)
	default:
		fmt.Println("unknown service")
		return -1
	}
}

func main() {

}
