package app

type BingoMatrix struct {
	Fields [][]string `json:"fields"`
}

type Bingo struct {
	UUID   string       `json:"uuid"`
	Title  string       `json:"title"`
	Fields [4][4]string `json:"fields"`
}

type Stats struct {
	Count int `json:"count"`
}

func TestData() [][]string {
	return [][]string{
		{"test", "test", "test", "test"},
		{"rest", "rest", "rest", "rest"},
		{"rest", "rest", "rest", "rest"},
		{"rest", "rest", "rest", "rest"},
	}
}

func ApiData() Bingo {
	return Bingo{
		Title: "API Test",
		Fields: [4][4]string{
			{"1", "2", "3", "4"},
			{"5", "6", "7", "8"},
			{"9", "10", "11", "12"},
			{"13", "14", "15", "16"},
		},
	}
}
