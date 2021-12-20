package repository

import "github.com/nurefsanmusevitoglu/e-commerce-platform/internal/model"

var campaigns map[string]model.Campaign

func CreateCampaign(name string, productCode string, duration int, limit int, targetSalesCount int) {
	// TO-DO: handle err if product already exists
	var campaign = model.Campaign {
		Name: name,
		ProductCode: productCode,
		Duration: duration,
		PriceManipulationLimit: limit,
		TargetSalesCount: targetSalesCount,
	}
	campaigns[campaign.Name] = campaign
}

func GetCampaignInfo(name string) model.Campaign {
	// TO-DO: handle err if product does not exist
	return campaigns[name]
}