package meetups

type Meetup struct {
	Created       int    `json:"created"`
	Duration      int    `json:"duration"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	RsvpLimit     int    `json:"rsvp_limit"`
	Status        string `json:"status"`
	Time          int    `json:"time"`
	LocalDate     string `json:"local_date"`
	LocalTime     string `json:"local_time"`
	Updated       int    `json:"updated"`
	UtcOffset     int    `json:"utc_offset"`
	WaitlistCount int    `json:"wait_list_count"`
	YesRsvpCount  int    `json:"yes_rsvp_count"`
	Venue         Venue  `json:"venue"`
	Group         Group  `json:"group"`
	Link          string `json:"link"`
	Description   string `json:"description"`
	Visibility    string `json:"visibility"`
}

type Venue struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Lat                  float64 `json:"lat"`
	Lon                  float64 `json:"lon"`
	Repinned             bool    `json:"repinned"`
	Address              string  `json:"address_1"`
	City                 string  `json:"city"`
	Country              string  `json:"country"`
	LocalizedCountryName string  `json:"localized_country_name"`
}

type Group struct {
	Created           int     `json:"created"`
	Name              string  `json:"name"`
	ID                int     `json:"id"`
	JoinMode          string  `json:"join_mode"`
	Lat               float64 `json:"lat"`
	Lon               float64 `json:"lon"`
	Urlname           string  `json:"urlname"`
	Who               string  `json:"who"`
	LocalizedLocation string  `json:"localized_location"`
	Region            string  `json:"region"`
}
