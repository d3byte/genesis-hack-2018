package models

type (
	Spreadsheet struct {
		// URL token
		Token string          `json:"token"`
		URL   string          `json:"url,omitempty"`
		Data  [][]interface{} `json:"data,omitempty" swaggertype:"array"`
	}

	SpreadsheetClear struct {
		// URL token
		Token string `json:"token"`
	}

	Data struct {
		Values           []interface{} `json:"values" swaggertype:"array"`
		SearchValue      string        `json:"searchValue"`
		RangeSearchValue string        `json:"-"`
	}
)
