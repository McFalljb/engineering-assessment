package models

type FoodTruck struct {
	ID                      int64   `json:"id" gorm:"primary_key;autoIncrement"`
	LocationID              int     `json:"locationid" gorm:"type:int;column:locationid"`
	Applicant               string  `json:"applicant"`
	FacilityType            string  `json:"facilitytype" gorm:"column:facilitytype"`
	CNN                     int64   `json:"cnn"`
	LocationDescription     string  `json:"locationdescription" gorm:"column:locationdescription"`
	Address                 string  `json:"address"`
	Blocklot                string  `json:"blocklot"`
	Block                   string  `json:"block"`
	Lot                     string  `json:"lot"`
	Permit                  string  `json:"permit"`
	Status                  string  `json:"status"`
	FoodItems               string  `json:"fooditems" gorm:"column:fooditems"`
	X                       float64 `json:"x"`
	Y                       float64 `json:"y"`
	Latitude                float64 `json:"latitude"`
	Longitude               float64 `json:"longitude"`
	Schedule                string  `json:"schedule"`
	DaysHours               string  `json:"dayshours,omitempty"`
	NOISent                 string  `json:"noisent,omitempty"`
	Approved                string  `json:"approved"`
	Received                int     `json:"received"`
	PriorPermit             bool    `json:"priorpermit"`
	ExpirationDate          string  `json:"expirationdate"`
	Location                string  `json:"location"`
	FirePreventionDistricts int     `json:"fire_prevention_districts"`
	PoliceDistricts         int     `json:"police_districts"`
	SupervisorDistricts     int     `json:"supervisor_districts"`
	ZipCodes                int     `json:"zip_codes"`
	NeighborhoodsOld        int     `json:"neighborhoods_old"`
}
