package service

import (
	"fmt"
	"go-sql-project/internal/model"
)

func GetContributionFundDetail() ([]model.FundContributionReport, error) {
	var fundContributions []model.FundContributionReport
	if err := DB.Find(&fundContributions).Error; err != nil {
		return nil, fmt.Errorf("error fetching contribution fund details: %w", err)
	}

	return fundContributions, nil
}

func AddContributionFund(id int, amount int) (model.FundContributionReport, error) {
	var existingFundContributionReport model.FundContributionReport
	if err := DB.First(&existingFundContributionReport, id).Error; err != nil {
		return model.FundContributionReport{}, fmt.Errorf("Failed when find existingFundContributionReport")
	}

	existingFundContributionReport.Amount += amount

	DB.Save(&existingFundContributionReport)
	return existingFundContributionReport, nil
}
