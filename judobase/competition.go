package judobase

type Contest struct {
	LastNameWhite  *string `json:"family_name_white"`
	FirstNameWhite *string `json:"given_name_white"`
	CountryWhite   *string `json:"country_short_white"`
	IpponWhite     *string `json:"ippon_w"`
	WazaWhite      *string `json:"waza_w"`

	LastNameBlue  *string `json:"family_name_blue"`
	FirstNameBlue *string `json:"given_name_blue"`
	CountryBlue   *string `json:"country_short_blue"`
	IpponBlue     *string `json:"ippon_b"`
	WazaBlue      *string `json:"waza_b"`
}

type Competition struct {
	Contests []Contest `json:"contests"`
}
