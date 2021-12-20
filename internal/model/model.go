package model

type Product struct {
	ProductCode string
	Price       int
	Stock       int
}

type Order struct {
	ProductCode string
	Quantity    int
}

type Campaign struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
	Status                 string
	Turnover               int
	AverageItemPrice       int
}
