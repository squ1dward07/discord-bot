package bot

import (
	"discord-bot/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Clan struct {
	Tag         string `json:"tag"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Location    struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		IsCountry   bool   `json:"isCountry"`
		CountryCode string `json:"countryCode"`
	} `json:"location"`
	IsFamilyFriendly bool `json:"isFamilyFriendly"`
	BadgeUrls        struct {
		Small  string `json:"small"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
	} `json:"badgeUrls"`
	ClanLevel         int64 `json:"clanLevel"`
	ClanPoints        int64 `json:"clanPoints"`
	ClanVersusPoints  int64 `json:"clanVersusPoints"`
	ClanCapitalPoints int64 `json:"clanCapitalPoints"`
	CapitalLeague     struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"capitalLeague`
	RequiredTrophies int64  `json:"requiredTrophies"`
	WarFrequency     string `json:"warFrequency"`
	WarWinStreak     int64  `json:"warWinStreak"`
	WarWins          int64  `json:"warWins"`
	WarTies          int64  `json:"warTies"`
	WarLosses        int64  `json:"warLosses"`
	IsWarLogPublic   bool   `json:"isWarLogPublic"`
	WarLeague        struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"warLeague`
	Members    int64 `json:"members"`
	MemberList []struct {
		Tag      string `json:"tag"`
		Name     string `json:"name"`
		Role     string `json:"role"`
		Explevel int64  `json:"explevel"`
		League   struct {
			ID       int64  `json:"id"`
			Name     string `json:"name"`
			IconUrls struct {
				Small  string `json:"small"`
				Tiny   string `json:"tiny"`
				Medium string `json:"medium"`
			} `json:"iconUrls"`
		} `json:"league"`
		Trophies          int64 `json : "trophies"`
		VersusTrophies    int64 `json : "versusTrophies"`
		ClanRank          int64 `json : "clanRank"`
		PreviousClanRank  int64 `json : "previousClanRank"`
		Donations         int64 `json : "donations"`
		DonationsRecieved int64 `json : "donationsRecieved"`
		PlayerHouse       struct {
			Elements []struct {
				Type string `json:"type"`
				ID   int64  `json:"id"`
			} `json:"elements"`
		} `json:"playerHouse"`
	} `json:"memberList"`
	Labels []struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		IconUrls struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"labels"`
	RequiredVersusTrophies int64 `json:"requiredVersusTrophies"`
	RequiredTownHallLevel  int64 `json:"requiredTownHallLevel"`
	ClanCapital            struct {
		CapitalHallLevel int64 `json:"capitalHallLevel"`
		Districts        []struct {
			ID                int64  `json:"id"`
			Name              string `json:"name"`
			DistrictHallLevel int64  `json:"districtHallLevel"`
		} `json:"districts"`
	} `json:"clanCapital"`
	ChatLanguage struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		LanguageCode string `json:"languageCode"`
	} `json:"chatLanguage"`
}

func ClanInfo(tag string) *Clan {
	clan := Clan{}
	url := "https://api.clashofclans.com/v1/clans/%23" + tag
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return &clan
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", config.Conf.ClashToken))
	r, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &clan
	} else if r.StatusCode != 200 {
		fmt.Println("none 200 response from [%s]", url)
		return &clan
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &clan
	}
	err = json.Unmarshal(body, &clan)

	if err != nil {
		fmt.Println(err.Error())
		return &clan
	}
	return &clan
}
