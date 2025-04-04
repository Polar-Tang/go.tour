package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		names        []string
		phoneNumbers []int
		expected     map[string]user
		errString    string
	}
	tests := []testCase{
		{
			[]string{"Eren", "Armin", "Mikasa"},
			[]int{1435555098, 987655509, 1826555456},
			map[string]user{"Eren": {"Eren", 1435555098}, "Armin": {"Armin", 987655509}, "Mikasa": {"Mikasa", 1826555456}},
			"",
		},
		{
			[]string{"Eren", "Armin"},
			[]int{1435555098, 987655509, 1826555456},
			nil,
			"invalid sizes",
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				[]string{"George", "Annie", "Reiner", "Sasha"},
				[]int{2095555981, 383855509, 482655545, 160455598},
				map[string]user{"George": {"George", 2095555981}, "Annie": {"Annie", 383855509}, "Reiner": {"Reiner", 482655545}, "Sasha": {"Sasha", 160455598}},
				"",
			},
			{
				[]string{"George", "Annie", "Reiner"},
				[]int{2095555981, 383855509, 482655545, 160455598},
				nil,
				"invalid sizes",
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output, err := getUserMap(test.names, test.phoneNumbers)
		if test.errString != "" && err == nil {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: %v
  actual err: none
`, test.names, test.phoneNumbers, test.errString)
		} else if test.errString == "" && err != nil {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: none
  actual err: %v
`, test.names, test.phoneNumbers, err)
		} else if test.errString != "" && err != nil && err.Error() != test.errString {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected err: %v
  actual err: %v
`, test.names, test.phoneNumbers, test.errString, err)
		} else if !compareMaps(output, test.expected) {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  names: %v
  phoneNumbers: %v
  expected:
%v
  actual:
%v
`, test.names, test.phoneNumbers, formatMap(test.expected), formatMap(output))
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  names: %v
  phoneNumbers: %v
  expected:
%v
  actual:
%v
`, test.names, test.phoneNumbers, formatMap(test.expected), formatMap(output))
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func formatMap(m map[string]user) string {
	var str string
	for key, value := range m {
		str += fmt.Sprintf("  * %s: %v\n", key, value)
	}
	return str
}

func compareMaps(map1, map2 map[string]user) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exist := map2[key]; !exist || value1 != value2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
