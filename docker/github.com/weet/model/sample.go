package model

import "github.com/weet/service"

// sampleTableにInsertする
func CreateSample(sample Sample) (Sample, error) {
	err := db.Create(&sample).Error
	if err != nil {
		return Sample{}, err
	}
	return sample, nil
}

func GetSampleById(sampleId uint) (service.Sample, error) {
	sample := service.Sample{}
	err := db.Where("id = ?", sampleId).First(&sample).Error
	return sample, err
}
