package mysql

import (
	"context"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"http3-server-poc/cmd/api/config"
	"http3-server-poc/internal/domain/models"
	"http3-server-poc/internal/infrastructure/mysql/boilermodels"
)

type GeodataRepository struct {
	database Database
}

func NewGeodataRepository(database Database) *GeodataRepository {
	return &GeodataRepository{
		database: database,
	}
}

func (r *GeodataRepository) SaveGeoshot(
	geoshot models.Geoshot,
	imagePath string,
	jsonPath string,
) error {
	timestamp, err := time.Parse("'2024-05-14 18:00:00'", geoshot.Timestamp)
	if err != nil {
		return err
	}

	boilerGeoshot := boilermodels.Geoshot{
		EventID:      null.NewInt(config.Cfg.EventId, true),
		DeviceID:     null.NewInt(geoshot.DeviceId, true),
		Imgpath:      null.NewString(imagePath, true),
		Lat:          null.NewFloat64(geoshot.Coordinates[0][0], true),
		Lon:          null.NewFloat64(geoshot.Coordinates[0][1], true),
		Timestamp:    null.NewTime(timestamp, true),
		Age:          null.NewInt(geoshot.Age, true),
		Buffered:     null.NewInt(0, true),
		Onstage:      null.NewInt(-1, true),
		Eventhos:     null.NewInt(-1, true),
		EventhosStat: null.Int{},
		Jsonpath:     null.NewString(jsonPath, false),
		Synced:       null.Int{},
	}

	err = boilerGeoshot.Insert(
		context.Background(),
		r.database,
		boil.Whitelist(
			boilermodels.GeoshotColumns.ID,
			boilermodels.GeoshotColumns.EventID,
			boilermodels.GeoshotColumns.DeviceID,
			boilermodels.GeoshotColumns.Imgpath,
			boilermodels.GeoshotColumns.Lat,
			boilermodels.GeoshotColumns.Lon,
			boilermodels.GeoshotColumns.Timestamp,
			boilermodels.GeoshotColumns.Age,
			boilermodels.GeoshotColumns.Buffered,
			boilermodels.GeoshotColumns.Onstage,
			boilermodels.GeoshotColumns.Eventhos,
			boilermodels.GeoshotColumns.EventhosStat,
			boilermodels.GeoshotColumns.Jsonpath,
			boilermodels.GeoshotColumns.Synced,
		),
	)
	if err != nil {
		return err
	}

	return nil
}