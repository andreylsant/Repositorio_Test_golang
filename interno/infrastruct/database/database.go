package database

import "github.com/andreylsant/test/interno/dominio/campaign"

type SaveCampaign struct {
	Campaign []campaign.Campaign
}

func (s *SaveCampaign) Save(campaign *campaign.Campaign) error{
	s.Campaign = append(s.Campaign, *campaign)
	return nil
}