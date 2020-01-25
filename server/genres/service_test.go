package genres

import "testing"

func TestFlatArrayToMap(t *testing.T) {
	r := Response{
		[]Genre{
			Genre{
				ID:   22,
				Name: "Adventure",
			},
		},
	}

	col := flatArrayToMap(r)

	if len(col) != len(r.Genres) {
		t.Errorf("expected map to have lenth=%d but got %d", len(r.Genres), len(col))
	}

	if col[22] != "Adventure" {
		t.Error("expected map to be filled with Adventure at key 22")
	}
}

func TestFlatEmptyArrayToMap(t *testing.T) {
	col := flatArrayToMap(Response{})

	if len(col) != 0 {
		t.Errorf("expected map to have lenth=0 but got %d", len(col))
	}
}
