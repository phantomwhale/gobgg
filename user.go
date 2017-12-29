package gobgg

import (
    "encoding/xml"
    "log"
    "net/http"
    "net/url"
    "strconv"
    "time"
    "io/ioutil"
)

type Query struct {
    username string
    domain string
    page int
    buddies bool
    guilds bool
    hot bool
    top bool
}

type User struct {
    Username    string      `xml:"name,attr"`
    Id      int     `xml:"id,attr"`
    Firstname struct {
            FirstName string        `xml:"value,attr"`
        } `xml:"firstname"`
    Lastname struct {
            LastName string         `xml:"value,attr"`
        } `xml:"lastname"`
    Avatarlink struct {
            AvatarLink string       `xml:"value,attr"`
        } `xml:"avatarlink"`
    Yearregistered struct {
            YearRegistered string   `xml:"value,attr"`
        } `xml:"yearregistered"`
    Lastlogin struct {
            LastLogin string    `xml:"value,attr"`
        } `xml:"lastlogin"`
    Stateorprovince struct {
            StateorProvince string  `xml:"value,attr"`
        } `xml:"stateorprovince"`
    Country struct {
            Country string  `xml:"value,attr"`
        } `xml:"country"`
    Webaddress struct {
            WebAddress string   `xml:"value,attr"`
        } `xml:"webaddress"`
    Xboxaccount struct {
            XboxAccount string  `xml:"value,attr"`
        } `xml:"xboxaccount"`
    Wiiaccount struct {
            WiiAcount string    `xml:"value,attr"`
        } `xml:"wiiaccount"`
    Psnaccount struct {
            PsnAccount string   `xml:"value,attr"`
        } `xml:"psnaccount"`
    Steamaccount struct {
            SteamAccount string `xml:"value,attr"`
        } `xml:"steamaccount"`
    Traderating struct {
            TradeRating string `xml:"value,attr"`
        } `xml:"traderating"`
    Marketrating struct {
            MarketRating string `xml:"value,attr"`
        } `xml:"marketrating"`
    Buddies []Buddy     `xml:"buddies>buddy"`
    Guilds  []Guild     `xml:"guilds>guild"`
    Tops    []Item      `xml:"top>item"`
    Hots    []Item      `xml:"hot>item"`
}

type Buddy struct {
    Username    string      `xml:"name,attr"`
    Id          int     `xml:"id,attr"`
}

type Guild struct {
    Guildname   string      `xml:"name,attr"`
    Id          int     `xml:"id,attr"`
}

type Item struct {
    Rank        int     `xml:"rank,attr"`
    Name        string      `xml:"name,attr"`
    Type        string      `xml:"type,attr"`
    Id          int     `xml:"id,attr"`
}

// Gets user information via GET, from query struct, and fills in default values
// Returns struct of user information from XML
func GetUser(q Query){

    user := User{}

    var Url *url.URL
    Url, err := url.Parse(BaseURL)
    if err != nil {
        log.Print("Error parsing url")
    }

    if q.user == "" {
        return nil
    }
    // Set defaults
    if q.domain != "" {
        q.domain = "boardgame"
    }
    if q.page <= 0 {
        q.page = 1
    }
    if q.buddies {
        q.bd = "1"
    }
    if q.guilds {
        q.gd = "1"
    }
    if q.hot {
        q.ht = "1"
    }
    if q.top {
        q.tp = "1"
    }

    Url.Path += UserSuffix
    parameters := url.Values{}
    parameters.Add("name", q.username)
    parameters.Add("buddies", q.bd)
    parameters.Add("guilds", q.gd)
    parameters.Add("hot", q.ht)
    parameters.Add("top", q.tp)
    parameters.Add("domain", q.domain)
    parameters.Add("page", strconv.Itoa(q.page))
    Url.RawQuery = parameters.Encode()

    log.Print(Url.String())
    // str name, bool buddies, guilds, hot, top, str domain, page
    resp, err := http.Get(Url.String())
    if err != nil{
        log.Print(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK{
        log.Printf("Status error: %v", resp.StatusCode)
        return
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Read body: %v", err)
    }
    log.Print(string(data))
    xml.Unmarshal(data, &user)

    return user
}
