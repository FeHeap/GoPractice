package main

import "fmt"

type Income interface {
	calculate() float64
	source() string
}

type FixedBilling struct {
	projectName string
	biddedAmount float64
}

type TimeAndMaterial struct {
	projectName string
	workHours float64
	hourlyRate float64
}

func (f FixedBilling) calculate() float64 {
	return f.biddedAmount
}

func (f FixedBilling) source() string {
	return f.projectName
}

func (t TimeAndMaterial) calculate() float64 {
	return t.workHours * t.hourlyRate
}

func (t TimeAndMaterial) source() string {
	return t.projectName
}

type Advertisement struct {
	adName string
	clickCount int
	incomePerclick float64
}

func (a Advertisement) calculate() float64 {
	return float64(a.clickCount) * a.incomePerclick
}

func (a Advertisement) source() string {
	return a.adName
}

func main() {
	p1 := FixedBilling{"項目1", 5000}
	p2 := FixedBilling{"項目2", 10000}
	p3 := TimeAndMaterial{"項目3", 100, 40}
	p4 := TimeAndMaterial{"項目4", 250, 20}
	p5 := Advertisement{"廣告1", 10000, 0.1}
	p6 := Advertisement{"廣告2", 20000, 0.05}
	ic := []Income{p1, p2, p3, p4, p5, p6}
	fmt.Println(calculateNetIncome(ic))
}

func calculateNetIncome(ic []Income) float64 {
	netIncome := 0.0
	for _, income := range ic {
		fmt.Printf("收入來源: %s, 收入金額: %.2f\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	return netIncome
}