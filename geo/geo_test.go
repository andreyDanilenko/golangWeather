// black box
package geo_test

import (
	"app/weather/geo"
	"testing"
)

// package geo
// white box

func TestGetMyLocation(t *testing.T) {
	// Arrange - подгатовка результата
	city := "Lonn"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполнение
	got, err := geo.GetMyLocation(city)
	// Assert - проверка с результата
	if err != nil {
		t.Error("Error getting city")
	}
	if got.City != expected.City {
		t.Errorf("Expected %v, Received %v", expected.City, got.City)
	}
}
