package stats

import "testing"

func TestCalculateDrops(t *testing.T) {
	tests := []struct {
		name          string
		hands         float64
		vpip          float64
		pfr           float64
		targetLimit   float64
		expectedLen   int
		expectedFirst float64
	}{
		{"Standard down to 20", 2450, 26.8, 14.5, 20.0, 7, 76},
		{"Nit down to 15", 2450, 18.2, 12.0, 15.0, 4, 301}, 
		{"Already below target", 1000, 20.0, 10.0, 22.0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := CalculateDrops(tt.hands, tt.vpip, tt.pfr, tt.targetLimit)
			
			if len(results) != tt.expectedLen {
				t.Errorf("Expected %d results, got %d", tt.expectedLen, len(results))
			}

			if len(results) > 0 && results[0].FoldsNeeded != tt.expectedFirst {
				t.Errorf("Expected first fold requirement to be %v, got %v", tt.expectedFirst, results[0].FoldsNeeded)
			}
		})
	}
}
