package bot

import (
	"discord-bot/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Player struct {
	Tag                      string `json : "tag"`
	Name                     string `json : "name"`
	TownHallLevel            int64  `json : "townHallLevel"`
	TownHallWeaponLevel      int64  `json : "townHallWeaponLevel"`
	ExpLevel                 int64  `json : "expLevel"`
	Trophies                 int64  `json : "trophies"`
	BestTrophies             int64  `json : "bestTrophies"`
	WarStars                 int64  `json : "warStars"`
	AttackWins               int64  `json : "attackWins"`
	DefenceWins              int64  `json : "defenceWins"`
	BuilderHallLevel         int64  `json : "builderHallLevel"`
	VersusTrophies           int64  `json : "versusTrophies"`
	BestVersusTrophies       int64  `json : "bestVersusTrophies"`
	VersusBattleWins         int64  `json : "versusBattleWins"`
	Role                     string `json : "townHallLevel"`
	WarPreference            string `json : "warPreference"`
	Donations                int64  `json : "donations"`
	DonationsReceived        int64  `json : "donationsReceived"`
	ClanCapitalContributions int64  `json : "clanCapitalContributions"`
	VersusBattleWinCount     int64  `json:"versusBattleWinCount"`

	Clan struct {
		Tag       string `json : "tag"`
		Name      string `json : "name"`
		ClanLevel int64  `json : "clanLevel"`
		BadgeUrls struct {
			Small  string `json:"small"`
			Large  string `json:"large"`
			Medium string `json:"medium"`
		} `json:"badgeUrls"`
	} `json:"clan"`
	League struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		IconUrls struct {
			Small  string `json:"small"`
			Tiny   string `json:"tiny"`
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"league"`
	LegendStatistics struct {
		LegendTrophies   int64 `json:"legendTrophies"`
		BestVersusSeason struct {
			ID       string `json:"id"`
			Rank     int64  `json:"rank"`
			Trophies int64  `json:"trophies"`
		} `json:"bestVersusSeason"`
		BestSeason struct {
			ID       string `json:"id"`
			Rank     int64  `json:"rank"`
			Trophies int64  `json:"trophies"`
		} `json:"bestSeason"`
		CurrentSeason struct {
			Trophies int64 `json:"trophies"`
		} `json:"currentSeason"`
	} `json:"legendStatistics"`
	Achievements []struct {
		Name           string `json:"name"`
		Stars          int64  `json:"stars"`
		Value          int64  `json:"value"`
		Target         int64  `json:"target"`
		Info           string `json:"info"`
		CompletionInfo string `json:"completionInfo"`
		Village        string `json:"village"`
	} `json:"achievements"`
	Labels []struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		IconUrls struct {
			Medium string `json:"medium"`
			Small  string `json:"small"`
		} `json:"iconUrls"`
	} `json:"labels"`
	Troops []struct {
		Name     string `json:"name"`
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Village  string `json:"village"`
	} `json:"troops"`
	Heroes []struct {
		Name     string `json:"name"`
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Village  string `json:"village"`
	} `json:"heroes"`
	Spells []struct {
		Name     string `json:"name"`
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Village  string `json:"village"`
	} `json:"spells"`
}

func PlayerInfo(tag string) *Player {
	player := Player{}
	url := "https://api.clashofclans.com/v1/players/%23" + tag
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return &player
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", config.Conf.ClashToken))
	r, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &player
	} else if r.StatusCode != 200 {
		fmt.Println("none 200 response from [%s]", url)
		return &player
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &player
	}
	err = json.Unmarshal(body, &player)

	if err != nil {
		fmt.Println(err.Error())
		return &player
	}
	return &player
}
