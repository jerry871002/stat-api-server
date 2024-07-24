package main

type MockStatStore struct {
	BattingData  []map[string]any
	PitchingData []map[string]any
	FieldingData []map[string]any
}

func NewMockStatStore() *MockStatStore {
	return &MockStatStore{
		BattingData: []map[string]any{
			{"Player": "Player1", "AVG": ".200"},
			{"Player": "Player2", "AVG": ".250"},
			{"Player": "Player3", "AVG": ".300"},
		},
		PitchingData: []map[string]any{
			{"Player": "Player1", "ERA": "3.00"},
			{"Player": "Player2", "ERA": "2.50"},
			{"Player": "Player3", "ERA": "2.00"},
		},
		FieldingData: []map[string]any{
			{"Player": "Player1", "FLDP": ".950"},
			{"Player": "Player2", "FLDP": ".960"},
			{"Player": "Player3", "FLDP": ".970"},
		},
	}
}

func (s *MockStatStore) GetPitching() ([]map[string]any, error) {
	return s.PitchingData, nil
}

func (s *MockStatStore) GetBatting() ([]map[string]any, error) {
	return s.BattingData, nil
}

func (s *MockStatStore) GetFielding() ([]map[string]any, error) {
	return s.FieldingData, nil
}
