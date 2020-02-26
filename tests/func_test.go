package tests

import (
	"net/http"
	"project/src"
	"testing"
)

// TestConvertToFahrenheit tests for expected results when converting
func TestConvertToFahrenheit(t *testing.T) {

	var tests = []struct {
		input    float32
		expected string
	}{
		{10.0, "50.0"},
		{-16.0, "3.2"},
		{5.5, "41.9"},
		{-6.6, "20.1"},
	}

	for _, test := range tests {
		if output := src.ConvertToFahrenheit(test.input); output != test.expected {
			t.Error("Test failed: {} inputed, {} expected, received: {}", test.input, test.expected, output)
		}
	}

}

//TestPercent tests for expected output
func TestPercent(t *testing.T) {

	var tests = []struct {
		input1   int
		input2   int
		expected float64
	}{
		{1000, 300, 30},
		{3546, 784, 22.109419063733785},
		{48957, 25, 0.051065220499622116},
	}

	for _, test := range tests {
		if output := src.Percent(test.input1, test.input2); output != test.expected {
			t.Error("Test failed: {} and {} inputed, {} expected, received: {}", test.input1, test.input2, test.expected, output)
		}
	}

}

//TestStatusEmoji checks for a valid emoji output
func TestSetStatusEmoji(t *testing.T) {
	url := "https://randomfox.ca/floof/"
	client := http.DefaultClient
	emoji, _, err := src.SetStatusEmoji(url, client)

	if err != nil {
		t.Error("Error: err returned not equal to nil in TestStatusEmoli")
	}
	length := len(emoji)
	if string(emoji[0]) != ":" || string(emoji[length-1]) != ":" {
		t.Error("Error: invalid emoji format in TestStatusEmoli")
	}
}
