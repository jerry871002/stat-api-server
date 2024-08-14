curl -X POST http://localhost:8080/simulate -H "Content-Type: application/json" -d '[
    {"name": "Mike Trout", "at_bat": 100, "hit": 30, "double": 5, "triple": 1, "home_run": 4, "ball_on_base": 10, "hit_by_pitch": 2},
    {"name": "Mookie Betts", "at_bat": 100, "hit": 25, "double": 6, "triple": 0, "home_run": 3, "ball_on_base": 12, "hit_by_pitch": 1},
    {"name": "Aaron Judge", "at_bat": 100, "hit": 28, "double": 7, "triple": 2, "home_run": 5, "ball_on_base": 8, "hit_by_pitch": 3},
    {"name": "Freddie Freeman", "at_bat": 100, "hit": 32, "double": 8, "triple": 1, "home_run": 6, "ball_on_base": 9, "hit_by_pitch": 2},
    {"name": "Juan Soto", "at_bat": 100, "hit": 27, "double": 4, "triple": 2, "home_run": 3, "ball_on_base": 11, "hit_by_pitch": 1},
    {"name": "Fernando Tatis Jr.", "at_bat": 100, "hit": 29, "double": 5, "triple": 1, "home_run": 4, "ball_on_base": 10, "hit_by_pitch": 2},
    {"name": "Bryce Harper", "at_bat": 100, "hit": 26, "double": 6, "triple": 0, "home_run": 3, "ball_on_base": 12, "hit_by_pitch": 1},
    {"name": "Ronald Acuña Jr.", "at_bat": 100, "hit": 31, "double": 7, "triple": 2, "home_run": 5, "ball_on_base": 8, "hit_by_pitch": 3},
    {"name": "Shohei Ohtani", "at_bat": 100, "hit": 33, "double": 8, "triple": 1, "home_run": 6, "ball_on_base": 9, "hit_by_pitch": 2}
]'

# bad input
curl -X POST http://localhost:8080/simulate -H "Content-Type: application/json" -d '[
    {"name": "Mike Trout", "at_bat": 100, "hit": 30, "double": 5, "invalid": 1, "homme_run": 4, "balll_on_base": 10, "hitt_by_pitch": 2},
    {"name": "Mookie Betts", "at_bat": 100, "hit": 25, "double": 6, "triple": 0, "home_run": 3, "ball_on_base": 12, "hit_by_pitch": 1},
    {"name": "Aaron Judge", "at_bat": 100, "hit": 28, "double": 7, "triple": 2, "home_run": 5, "ball_on_base": 8, "hit_by_pitch": 3},
    {"name": "Freddie Freeman", "at_bat": 100, "hit": 32, "double": 8, "triple": 1, "home_run": 6, "ball_on_base": 9, "hit_by_pitch": 2},
    {"name": "Juan Soto", "at_bat": 100, "hit": 27, "double": 4, "triple": 2, "home_run": 3, "ball_on_base": 11, "hit_by_pitch": 1},
    {"name": "Fernando Tatis Jr.", "at_bat": 100, "hit": 29, "double": 5, "triple": 1, "home_run": 4, "ball_on_base": 10, "hit_by_pitch": 2},
    {"name": "Bryce Harper", "at_bat": 100, "hit": 26, "double": 6, "triple": 0, "home_run": 3, "ball_on_base": 12, "hit_by_pitch": 1},
    {"name": "Ronald Acuña Jr.", "at_bat": 100, "hit": 31, "double": 7, "triple": 2, "home_run": 5, "ball_on_base": 8, "hit_by_pitch": 3},
    {"name": "Shohei Ohtani", "at_bat": 100, "hit": 33, "double": 8, "triple": 1, "home_run": 6, "ball_on_base": 9, "hit_by_pitch": 2}
]'