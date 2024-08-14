package main

type MockStatStore struct {
	TeamData     []Team
	BattingData  []map[string]any
}

func NewMockStatStore() *MockStatStore {
	return &MockStatStore{
		TeamData: []Team{
			{"Team1", 2024},
			{"Team2", 2024},
			{"Team3", 2024},
		},
		BattingData: []map[string]any{
			{"name": "Player1", "at_bat": "50", "hit:": "10"},
			{"name": "Player2", "at_bat": "100", "hit:": "20"},
			{"name": "Player3", "at_bat": "150", "hit:": "30"},
		},
	}
}

func (s *MockStatStore) GetTeams() ([]Team, error) {
	return s.TeamData, nil
}

func (s *MockStatStore) GetBattingStat(team string, year int) ([]map[string]any, error) {
	if year == 2024 {
		return s.BattingData, nil
	}
	return nil, nil
}
