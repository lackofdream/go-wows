package wows

type PlayersResponse struct {
	Status string `json:"status"`
	Meta   struct {
		Count int `json:"count"`
	} `json:"meta"`
	Data []Players `json:"data"`
}

type Players struct {
	Nickname  string `json:"nickname"`
	AccountID int    `json:"account_id"`
}

type PlayerPersonalDataResponse struct {
	Status string `json:"status"`
	Meta   struct {
		Count  int         `json:"count"`
		Hidden interface{} `json:"hidden"`
	} `json:"meta"`
	Data map[int]PlayerPersonalData `json:"data"`
}

type PlayerPersonalData struct {
	LastBattleTime int         `json:"last_battle_time"`
	AccountID      int         `json:"account_id"`
	LevelingTier   int         `json:"leveling_tier"`
	CreatedAt      int         `json:"created_at"`
	LevelingPoints int         `json:"leveling_points"`
	UpdatedAt      int         `json:"updated_at"`
	Private        interface{} `json:"private"`
	HiddenProfile  bool        `json:"hidden_profile"`
	LogoutAt       int         `json:"logout_at"`
	Karma          interface{} `json:"karma"`
	Statistics     struct {
		Distance int `json:"distance"`
		Battles  int `json:"battles"`
		Pvp      struct {
			MaxXp             int `json:"max_xp"`
			DamageToBuildings int `json:"damage_to_buildings"`
			MainBattery       struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				Hits           int   `json:"hits"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
				Shots          int   `json:"shots"`
			} `json:"main_battery"`
			MaxShipsSpottedShipID int64 `json:"max_ships_spotted_ship_id"`
			MaxDamageScouting     int   `json:"max_damage_scouting"`
			ArtAgro               int64 `json:"art_agro"`
			MaxXpShipID           int64 `json:"max_xp_ship_id"`
			ShipsSpotted          int   `json:"ships_spotted"`
			SecondBattery         struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				Hits           int   `json:"hits"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
				Shots          int   `json:"shots"`
			} `json:"second_battery"`
			MaxFragsShipID            int64       `json:"max_frags_ship_id"`
			Xp                        int         `json:"xp"`
			SurvivedBattles           int         `json:"survived_battles"`
			DroppedCapturePoints      int         `json:"dropped_capture_points"`
			MaxDamageDealtToBuildings int         `json:"max_damage_dealt_to_buildings"`
			TorpedoAgro               int         `json:"torpedo_agro"`
			Draws                     int         `json:"draws"`
			ControlCapturedPoints     int         `json:"control_captured_points"`
			BattlesSince510           int         `json:"battles_since_510"`
			MaxTotalAgroShipID        int64       `json:"max_total_agro_ship_id"`
			PlanesKilled              int         `json:"planes_killed"`
			Battles                   int         `json:"battles"`
			MaxShipsSpotted           int         `json:"max_ships_spotted"`
			MaxSuppressionsShipID     interface{} `json:"max_suppressions_ship_id"`
			SurvivedWins              int         `json:"survived_wins"`
			Frags                     int         `json:"frags"`
			DamageScouting            int         `json:"damage_scouting"`
			MaxTotalAgro              int         `json:"max_total_agro"`
			MaxFragsBattle            int         `json:"max_frags_battle"`
			CapturePoints             int         `json:"capture_points"`
			Ramming                   struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
			} `json:"ramming"`
			SuppressionsCount    int `json:"suppressions_count"`
			MaxSuppressionsCount int `json:"max_suppressions_count"`
			Torpedoes            struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				Hits           int   `json:"hits"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
				Shots          int   `json:"shots"`
			} `json:"torpedoes"`
			MaxPlanesKilledShipID int64 `json:"max_planes_killed_ship_id"`
			Aircraft              struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
			} `json:"aircraft"`
			TeamCapturePoints               int   `json:"team_capture_points"`
			ControlDroppedPoints            int   `json:"control_dropped_points"`
			MaxDamageDealt                  int   `json:"max_damage_dealt"`
			MaxDamageDealtToBuildingsShipID int64 `json:"max_damage_dealt_to_buildings_ship_id"`
			MaxDamageDealtShipID            int64 `json:"max_damage_dealt_ship_id"`
			Wins                            int   `json:"wins"`
			Losses                          int   `json:"losses"`
			DamageDealt                     int   `json:"damage_dealt"`
			MaxPlanesKilled                 int   `json:"max_planes_killed"`
			MaxScoutingDamageShipID         int64 `json:"max_scouting_damage_ship_id"`
			TeamDroppedCapturePoints        int   `json:"team_dropped_capture_points"`
			BattlesSince512                 int   `json:"battles_since_512"`
		} `json:"pvp"`
	} `json:"statistics"`
	Nickname       string `json:"nickname"`
	StatsUpdatedAt int    `json:"stats_updated_at"`
}

type PlayerAchievementsResponse struct {
	Status string `json:"status"`
	Meta   struct {
		Count  int         `json:"count"`
		Hidden interface{} `json:"hidden"`
	} `json:"meta"`
	Data map[int]PlayerAchievements `json:"data"`
}

type PlayerAchievements struct {
	LastBattleTime int         `json:"last_battle_time"`
	AccountID      int         `json:"account_id"`
	LevelingTier   int         `json:"leveling_tier"`
	CreatedAt      int         `json:"created_at"`
	LevelingPoints int         `json:"leveling_points"`
	UpdatedAt      int         `json:"updated_at"`
	Private        interface{} `json:"private"`
	HiddenProfile  bool        `json:"hidden_profile"`
	LogoutAt       int         `json:"logout_at"`
	Karma          interface{} `json:"karma"`
	Statistics     struct {
		Distance int `json:"distance"`
		Battles  int `json:"battles"`
		Pvp      struct {
			MaxXp             int `json:"max_xp"`
			DamageToBuildings int `json:"damage_to_buildings"`
			MainBattery       struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				Hits           int   `json:"hits"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
				Shots          int   `json:"shots"`
			} `json:"main_battery"`
			MaxShipsSpottedShipID int64 `json:"max_ships_spotted_ship_id"`
			MaxDamageScouting     int   `json:"max_damage_scouting"`
			ArtAgro               int   `json:"art_agro"`
			MaxXpShipID           int64 `json:"max_xp_ship_id"`
			ShipsSpotted          int   `json:"ships_spotted"`
			SecondBattery         struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				Hits           int   `json:"hits"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
				Shots          int   `json:"shots"`
			} `json:"second_battery"`
			MaxFragsShipID            int64       `json:"max_frags_ship_id"`
			Xp                        int         `json:"xp"`
			SurvivedBattles           int         `json:"survived_battles"`
			DroppedCapturePoints      int         `json:"dropped_capture_points"`
			MaxDamageDealtToBuildings int         `json:"max_damage_dealt_to_buildings"`
			TorpedoAgro               int         `json:"torpedo_agro"`
			Draws                     int         `json:"draws"`
			ControlCapturedPoints     int         `json:"control_captured_points"`
			BattlesSince510           int         `json:"battles_since_510"`
			MaxTotalAgroShipID        int64       `json:"max_total_agro_ship_id"`
			PlanesKilled              int         `json:"planes_killed"`
			Battles                   int         `json:"battles"`
			MaxShipsSpotted           int         `json:"max_ships_spotted"`
			MaxSuppressionsShipID     interface{} `json:"max_suppressions_ship_id"`
			SurvivedWins              int         `json:"survived_wins"`
			Frags                     int         `json:"frags"`
			DamageScouting            int         `json:"damage_scouting"`
			MaxTotalAgro              int         `json:"max_total_agro"`
			MaxFragsBattle            int         `json:"max_frags_battle"`
			CapturePoints             int         `json:"capture_points"`
			Ramming                   struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
			} `json:"ramming"`
			SuppressionsCount    int `json:"suppressions_count"`
			MaxSuppressionsCount int `json:"max_suppressions_count"`
			Torpedoes            struct {
				MaxFragsBattle int   `json:"max_frags_battle"`
				Frags          int   `json:"frags"`
				Hits           int   `json:"hits"`
				MaxFragsShipID int64 `json:"max_frags_ship_id"`
				Shots          int   `json:"shots"`
			} `json:"torpedoes"`
			MaxPlanesKilledShipID int64 `json:"max_planes_killed_ship_id"`
			Aircraft              struct {
				MaxFragsBattle int         `json:"max_frags_battle"`
				Frags          int         `json:"frags"`
				MaxFragsShipID interface{} `json:"max_frags_ship_id"`
			} `json:"aircraft"`
			TeamCapturePoints               int         `json:"team_capture_points"`
			ControlDroppedPoints            int         `json:"control_dropped_points"`
			MaxDamageDealt                  int         `json:"max_damage_dealt"`
			MaxDamageDealtToBuildingsShipID interface{} `json:"max_damage_dealt_to_buildings_ship_id"`
			MaxDamageDealtShipID            int64       `json:"max_damage_dealt_ship_id"`
			Wins                            int         `json:"wins"`
			Losses                          int         `json:"losses"`
			DamageDealt                     int         `json:"damage_dealt"`
			MaxPlanesKilled                 int         `json:"max_planes_killed"`
			MaxScoutingDamageShipID         int64       `json:"max_scouting_damage_ship_id"`
			TeamDroppedCapturePoints        int         `json:"team_dropped_capture_points"`
			BattlesSince512                 int         `json:"battles_since_512"`
		} `json:"pvp"`
	} `json:"statistics"`
	Nickname       string `json:"nickname"`
	StatsUpdatedAt int    `json:"stats_updated_at"`
}
