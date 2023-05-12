package goodrussians

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"example.com/m/v2/internal/auditlog"
	"fmt"
	"time"
)

const (
	fullScaleInvasionDate = "2022-02-24"
	YYYYMMDD              = "2006-01-02"
)

type goodRussians struct {
	PersonnellUnits         uint `json:"personnel_units"`
	Tanks                   uint `json:"tanks"`
	ArtillerySystems        uint `json:"artillery_systems"`
	ArmoredFighningVehicles uint `json:"armoured_fighting_vehicles"`
	Mlrs                    uint `json:"mlrs"`
	AAWarfareSystems        uint `json:"aa_warfare_systems"`
	Planes                  uint `json:"planes"`
	Helicopters             uint `json:"helicopters"`
	VechiclesFuelTanks      uint `json:"vehicles_fuel_tanks"`
	WarshipCutters          uint `json:"warships_cutters"`
	CruiseMissiles          uint `json:"cruise_missiles"`
	UavSystems              uint `json:"uav_systems"`
	AtgmSrbmSystems         uint `json:"atgm_srbm_systems"`
}

type statsRussiansResponse struct {
	Stats goodRussians `json:"stats"`
}

type goodRussiansResponse struct {
	Data statsRussiansResponse `json:"data"`
}

type service struct {
	auditLog auditlog.AuditLogService
}

func NewGoodRussiansCounter(auditLog auditlog.AuditLogService) *service {
	return &service{
		auditLog: auditLog,
	}
}

func (s *service) Count(ctx context.Context) {
	var goodRussiansNew goodRussians
	var goodRussiansOld goodRussians
	var err error
	currentDay := fullScaleInvasionDate

	for s.ifDateInThePast(currentDay) {
		goodRussiansOld = goodRussiansNew

		goodRussiansNew, err = s.getGoodRussians(currentDay)
		if err != nil {
			fmt.Printf("cannot get good russinans for date: %s err: %w", currentDay, err)
		}

		fmt.Print(currentDay + ": ")
		err = s.auditLog.CreateSnapshot(ctx, "goodRussians", goodRussiansOld, goodRussiansNew)
		if err != nil {
			fmt.Printf("cannot create snapshot for date: %s err: %w", currentDay, err)
		}

		currentDay = s.getNextDay(currentDay)
	}

	fmt.Println("To be continued..")
}

func (s *service) getNextDay(date string) string {
	t, _ := time.Parse(YYYYMMDD, date)
	t1 := t.Add(24 * time.Hour)
	return t1.Format(YYYYMMDD)
}

func (s *service) ifDateInThePast(date string) bool {
	t, _ := time.Parse(YYYYMMDD, date)

	return time.Now().After(t)
}

func (s *service) getGoodRussians(t string) (goodRussians, error) {
	client := http.DefaultClient
	resp, err := client.Get(fmt.Sprintf("https://russianwarship.rip/api/v1/statistics/%s", t))
	if err != nil {
		return goodRussians{}, fmt.Errorf("cannot get good russians on date")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return goodRussians{}, fmt.Errorf("cannot get response body")
	}

	var deserializeResp goodRussiansResponse

	err = json.Unmarshal(responseBody, &deserializeResp)
	if err != nil {
		return goodRussians{}, fmt.Errorf("cannot deserialize response body")
	}

	return deserializeResp.Data.Stats, nil
}
