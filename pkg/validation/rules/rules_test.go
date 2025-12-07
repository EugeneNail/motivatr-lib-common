package rules

import (
	"testing"
	"time"
)

const noError = ""

func TestRequired(t *testing.T) {
	tableTests := []struct {
		name string
		data map[string]any
		want string
	}{
		{"Missing field", map[string]any{}, "The test field is required"},

		{"Empty string", map[string]any{"test": ""}, "The test field is required"},
		{"Non-empty string", map[string]any{"test": "Merlin"}, noError},

		{"Zero number", map[string]any{"test": 0}, "The test field is required"},
		{"Greater than zero number", map[string]any{"test": 45}, noError},
		{"Less than zero number", map[string]any{"test": -45}, noError},

		{"Empty array", map[string]any{"test": []string{}}, "The test field is required"},
		{"Non-empty array", map[string]any{"test": []string{"1984", "Crime and punishment"}}, noError},

		{"Empty map", map[string]any{"test": map[string]string{}}, "The test field is required"},
		{"Non-empty map", map[string]any{"test": map[int]string{1: "1984", 2: "Crime and punishment"}}, noError},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Required()(tt.data, "test")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	ruleFunc := Max(10)
	tableTests := []struct {
		name string
		data map[string]any
		want string
	}{
		{"Latin string exceeds the limit", map[string]any{"test": "Nostradamus"}, "The test field must not be greater than 10 characters"},
		{"Latin string equals the limit", map[string]any{"test": "Ereshkigal"}, noError},
		{"Latin string does not exceed the limit", map[string]any{"test": "Sebastian"}, noError},

		{"Cyrillic string exceeds the limit", map[string]any{"test": "Максимилиан"}, "The test field must not be greater than 10 characters"},
		{"Cyrillic string equals the limit", map[string]any{"test": "Александра"}, noError},
		{"Cyrillic string does not exceed the limit", map[string]any{"test": "Сергей"}, noError},

		{"Logographic string exceeds the limit", map[string]any{"test": "騒音騒音騒音騒音騒音騒"}, "The test field must not be greater than 10 characters"},
		{"Logographic string equals the limit", map[string]any{"test": "騒音騒音騒音騒音騒音"}, noError},
		{"Logographic string does not exceed the limit", map[string]any{"test": "騒音騒音騒音"}, noError},

		{"Number exceeds the limit", map[string]any{"test": 15}, "The test field must not be greater than 10"},
		{"Number equals the limit", map[string]any{"test": 10}, noError},
		{"Number does not exceed the limit", map[string]any{"test": 5}, noError},

		{"Array size exceeds the limit", map[string]any{"test": []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}}, "The test field must not have more than 10 items"},
		{"Array size equals the limit", map[string]any{"test": []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}}, noError},
		{"Array size does not exceed the limit", map[string]any{"test": []string{"1", "2", "3"}}, noError},

		{"Map size exceeds the limit", map[string]any{"test": map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "11"}}, "The test field must not have more than 10 items"},
		{"Map size equals the limit", map[string]any{"test": map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10"}}, noError},
		{"Map size does not exceed the limit", map[string]any{"test": map[int]string{1: "1", 2: "2", 3: "3"}}, noError},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ruleFunc(tt.data, "test")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	ruleFunc := Min(3)
	tableTests := []struct {
		name string
		data map[string]any
		want string
	}{
		{"Latin string falls below the limit", map[string]any{"test": "Al"}, "The test field must not be less than 3 characters"},
		{"Latin string equals the limit", map[string]any{"test": "Tom"}, noError},
		{"Latin string does not fall below the limit", map[string]any{"test": "Sebastian"}, noError},

		{"Cyrillic string falls below the limit", map[string]any{"test": "Ян"}, "The test field must not be less than 3 characters"},
		{"Cyrillic string equals the limit", map[string]any{"test": "Ада"}, noError},
		{"Cyrillic string does not fall below the limit", map[string]any{"test": "Сергей"}, noError},

		{"Logographic string falls below the limit", map[string]any{"test": "音騒"}, "The test field must not be less than 3 characters"},
		{"Logographic string equals the limit", map[string]any{"test": "騒音騒"}, noError},
		{"Logographic string does not fall below the limit", map[string]any{"test": "騒音騒音騒音"}, noError},

		{"Number falls below the limit", map[string]any{"test": 2}, "The test field must not be less than 3"},
		{"Number equals the limit", map[string]any{"test": 3}, noError},
		{"Number does not fall below the limit", map[string]any{"test": 5}, noError},

		{"Array size falls below the limit", map[string]any{"test": []string{"1", "2"}}, "The test field must not have less than 3 items"},
		{"Array size equals the limit", map[string]any{"test": []string{"1", "2", "3"}}, noError},
		{"Array size does not exceed the limit", map[string]any{"test": []string{"1", "2", "3", "4", "5"}}, noError},

		{"Map size falls below the limit", map[string]any{"test": map[int]string{1: "1", 2: "2"}}, "The test field must not have less than 3 items"},
		{"Map size equals the limit", map[string]any{"test": map[int]string{1: "1", 2: "2", 3: "3"}}, noError},
		{"Map size does not fall below the limit", map[string]any{"test": map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5"}}, noError},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ruleFunc(tt.data, "test")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDate(t *testing.T) {
	type tableTest struct {
		name string
		data map[string]any
		want string
	}

	tableTests := []tableTest{
		{"Accepts only 2006-01-02 format", map[string]any{"test": "2025-01-01 12:34:56"}, "The test field format is invalid"},
		{"Date exceeds long month capacity", map[string]any{"test": "2025-01-32"}, "The test field format is invalid"},
		{"Date exceeds short month capacity", map[string]any{"test": "2025-09-31"}, "The test field format is invalid"},
		{"Date exceeds February capacity", map[string]any{"test": "2025-02-29"}, "The test field format is invalid"},
		{"Date fits February capacity", map[string]any{"test": "2025-02-28"}, noError},
		{"Date exceeds leap February capacity", map[string]any{"test": "2025-02-30"}, "The test field format is invalid"},
		{"Date fits leap February capacity", map[string]any{"test": "2024-02-29"}, noError},
		{"Previous century", map[string]any{"test": "1999-12-31"}, "The test field format is invalid"},
		{"Next century", map[string]any{"test": "2100-01-01"}, "The test field format is invalid"},
	}

	firstDate, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		t.Fatalf("cannot parse date 2024-01-01: %v", err)
	}

	for i := 0; i <= 365; i++ {
		date := firstDate.Add(time.Hour * 24 * time.Duration(i)).Format("2006-01-02")
		test := tableTest{"Valid date " + date, map[string]any{"test": date}, noError}

		tableTests = append(tableTests, test)
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Date()(tt.data, "test"); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRegex(t *testing.T) {
	tableTests := []struct {
		name    string
		pattern string
		data    map[string]any
		want    string
	}{
		{"Alpha one word", Alpha, map[string]any{"test": "Success"}, noError},
		{"Alpha merged words", Alpha, map[string]any{"test": "SaySomething"}, noError},
		{"Alpha separate words", Alpha, map[string]any{"test": "Say something"}, "The test field format is invalid"},
		{"Alpha dot", Alpha, map[string]any{"test": "Say something."}, "The test field format is invalid"},
		{"Alpha dash", Alpha, map[string]any{"test": "Block-one"}, "The test field format is invalid"},

		{"AlphaNum one word", AlphaNum, map[string]any{"test": "Success"}, noError},
		{"AlphaNum merged words", AlphaNum, map[string]any{"test": "SaySomething"}, noError},
		{"AlphaNum separate words", AlphaNum, map[string]any{"test": "Say something"}, "The test field format is invalid"},
		{"AlphaNum one word with a number", AlphaNum, map[string]any{"test": "4chan"}, noError},
		{"AlphaNum numbers only", AlphaNum, map[string]any{"test": "12345"}, noError},
		{"AlphaNum two words with a number", AlphaNum, map[string]any{"test": "There are 3 words"}, "The test field format is invalid"},
		{"AlphaNum dot", AlphaNum, map[string]any{"test": "2025.12.31"}, "The test field format is invalid"},
		{"AlphaNum dash", AlphaNum, map[string]any{"test": "2025-12-31"}, "The test field format is invalid"},

		{"San one word", San, map[string]any{"test": "Success"}, noError},
		{"San merged words", San, map[string]any{"test": "SaySomething"}, noError},
		{"San separate words", San, map[string]any{"test": "Say something"}, noError},
		{"San one word with a number", San, map[string]any{"test": "4chan"}, noError},
		{"San numbers only", San, map[string]any{"test": "12345"}, noError},
		{"San two words with a number", San, map[string]any{"test": "There are 3 words"}, noError},
		{"San dash", San, map[string]any{"test": "2025-12-31"}, "The test field format is invalid"},
		{"San dot", San, map[string]any{"test": "2025.12.31"}, "The test field format is invalid"},

		{"Sand one word", Sand, map[string]any{"test": "Success"}, noError},
		{"Sand merged words", Sand, map[string]any{"test": "SaySomething"}, noError},
		{"Sand separate words", Sand, map[string]any{"test": "Say something"}, noError},
		{"Sand one word with a number", Sand, map[string]any{"test": "4chan"}, noError},
		{"Sand numbers only", Sand, map[string]any{"test": "12345"}, noError},
		{"Sand two words with a number", Sand, map[string]any{"test": "There are 3 words"}, noError},
		{"Sand dash", Sand, map[string]any{"test": "2025-12-31"}, noError},
		{"Sand dot", Sand, map[string]any{"test": "2025.12.31"}, "The test field format is invalid"},
	}

	for _, tt := range tableTests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Regex(tt.pattern)(tt.data, "test"); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
