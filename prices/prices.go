package prices

import (
	"fmt"

	"example.com/practice-project/conversion"
	"example.com/practice-project/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_rate"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}
	job.InputPrices = prices

	return nil
}

func (job TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}
	result := map[string]string{}

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}
