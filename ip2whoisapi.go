package ip2whois

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/idna"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// The IP2WhoisResult struct stores all of the available
// whois info found in the IP2WHOIS API.
type IP2WhoisResult struct {
	Domain      string    `json:"domain"`
	DomainID    string    `json:"domain_id"`
	Status      string    `json:"status"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
	ExpireDate  time.Time `json:"expire_date"`
	DomainAge   int       `json:"domain_age"`
	WhoisServer string    `json:"whois_server"`
	Registrar   struct {
		IanaID string `json:"iana_id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
	} `json:"registrar"`
	Registrant struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"registrant"`
	Admin struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"admin"`
	Tech struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"tech"`
	Billing struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"billing"`
	Nameservers []string `json:"nameservers"`
}

// The Api struct is the main object used to query the IP2WHOIS API.
type Api struct {
	apiKey string
}

const baseURL = "api.ip2whois.com/v2"

// OpenApi initializes with the API key
func OpenApi(apikey string) *Api {
	var api = &Api{}
	api.apiKey = apikey
	return api
}

// LookUp will return all whois fields based on the queried domain
func (a *Api) LookUp(domain string) (IP2WhoisResult, error) {
	var res IP2WhoisResult

	myUrl := "https://" + baseURL + "?key=" + url.QueryEscape(a.apiKey) + "&domain=" + url.QueryEscape(domain)

	resp, err := http.Get(myUrl)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return res, err
		}

		err = json.Unmarshal(bodyBytes, &res)

		if err != nil {
			return res, err
		}

		return res, nil
	}

	return res, errors.New("Error HTTP " + strconv.Itoa(int(resp.StatusCode)))
}

// GetPunycode will convert normal text to Punycode
func (a *Api) GetPunycode(domain string) (string, error) {
	puny, err := url.Parse("https://" + domain)
	if err != nil {
		return "", err
	}
	str, err := idna.ToASCII(puny.Host)
	if err != nil {
		return "", err
	}
	return str, nil
}

// GetNormalText will convert Punycode to normal text
func (a *Api) GetNormalText(domain string) (string, error) {
	puny, err := url.Parse("https://" + domain)
	if err != nil {
		return "", err
	}
	str, err := idna.ToUnicode(puny.Host)
	if err != nil {
		return "", err
	}
	return str, nil
}
