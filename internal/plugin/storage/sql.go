package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/johann-vu/iot-scenario/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

func NewSQLDatasetRepository(connectionString string) (domain.DatasetRepository, error) {

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("opening connection to database %v", err)
	}

	err = db.AutoMigrate(&datasetModel{})
	if err != nil {
		return nil, fmt.Errorf("auto migrating dataset model %v", err)
	}

	return &sqlRepo{
		db: db,
	}, nil
}

func (r *sqlRepo) Add(ctx context.Context, d domain.Dataset) error {

	err := r.db.Create(newDatasetModel(d)).Error
	if err != nil {
		return err
	}
	return nil
}

func (dr *sqlRepo) Get(ctx context.Context, from time.Time, to time.Time) ([]domain.Dataset, error) {

	var results []datasetModel
	err := dr.db.Where("timestamp BETWEEN ? AND ?", from, to).Order("timestamp ASC").Find(&results).Error
	if err != nil {
		return nil, err
	}

	datasets := make([]domain.Dataset, len(results))
	for i, r := range results {
		datasets[i] = r.ToDomain()
	}

	return datasets, nil
}

type datasetModel struct {
	gorm.Model
	SensorID  string
	Timestamp time.Time
	Value     float64
}

func newDatasetModel(d domain.Dataset) *datasetModel {

	return &datasetModel{
		SensorID:  d.SensorID,
		Timestamp: d.Timestamp,
		Value:     d.Value,
	}
}

func (d datasetModel) ToDomain() domain.Dataset {

	return domain.Dataset{
		SensorID:  d.SensorID,
		Timestamp: d.Timestamp,
		Value:     d.Value,
	}
}
