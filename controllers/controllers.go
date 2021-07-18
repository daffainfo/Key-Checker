package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//normal curl
func normal_curl(url string) (string, int) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nStatus Code:\n" + strconv.Itoa(req.StatusCode))
	fmt.Println("\nResponse Body:\n" + string(body))

	defer req.Body.Close()

	return string(body), req.StatusCode
}

func Asana(asana_access string) string {
	req, err := http.NewRequest("GET", "https://app.asana.com/api/1.0/users/me", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+asana_access)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.Status == "401 Unauthorized" {
		return "Invalid API key"
	} else {
		return "API key valid"
	}
}
func Bing(bingapi string) string {
	url := "https://dev.virtualearth.net/REST/v1/Locations?CountryRegion=US&adminDistrict=WA&locality=Somewhere&postalCode=98001&addressLine=100%20Main%20St.&key=" + bingapi
	a, b := normal_curl(url)
	_ = a

	if b == 200 {
		return "\nBing Maps API Key valid"
	} else {
		return "\nBing Maps API Key invalid"
	}

}

func Bitly(bitly_token string) string {
	error_msg := "{\"status_code\":500,\"status_txt\":\"INVALID_ARG_ACCESS_TOKEN\",\"data\":[]}"
	url := "https://api-ssl.bitly.com/v3/shorten?access_token=" + bitly_token + "&longUrl=https://www.google.com"
	a, b := normal_curl(url)

	if a != error_msg && b == 200 {
		return "\nBit.ly Access token valid"
	} else {
		return "\nBit.ly Access token invalid"
	}
}

func Branch(branch_key string, branch_secret string) string {
	url := "https://api2.branch.io/v1/app/" + branch_key + "?branch_secret=" + branch_secret

	a, b := normal_curl(url)
	_ = a

	if b == 200 {
		return "\nBranch.IO Key and Secret valid"
	} else {
		return "\nBranch.IO Key and Secret invalid"
	}
}

func Buildkite(buildkite_access string) string {
	req, err := http.NewRequest("GET", "https://api.buildkite.com/v2/user", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+buildkite_access)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.Status == "401 Unauthorized" {
		return "Invalid API key"
	} else {
		return "API key valid"
	}
}

func ButterCMS(branch_key string) string {
	url := "https://api.buttercms.com/v2/posts/?auth_token=" + branch_key

	a, b := normal_curl(url)
	_ = a

	if b == 200 {
		return "\nBranch.IO Key and Secret valid"
	} else {
		return "\nBranch.IO Key and Secret invalid"
	}
}

func Circleci(circleci_token string) string {
	url := "https://circleci.com/api/v1.1/me?circle-token=" + circleci_token

	a, b := normal_curl(url)
	_ = a

	if b == 200 {
		return "\nCircleCI Access Token valid"
	} else {
		return "\nCircleCI Access Token invalid"
	}
}

func Datadog(dog_api string, dog_application string) string {
	url := "https://api.datadoghq.com/api/v1/dashboard?api_key=" + dog_api + "&application_key=" + dog_application

	a, b := normal_curl(url)
	_ = a

	if b == 200 {
		return "\nDataDog API key valid"
	} else {
		return "\nDataDog API key invalid"
	}
}

func Loqate(loqate_key string) string {
	error_msg := "{\"Items\":[{\"Error\":\"2\",\"Description\":\"Unknown key\",\"Cause\":\"The key you are using to access the service was not found.\",\"Resolution\":\"Please check that the key is correct. It should be in the form AA11-AA11-AA11-AA11.\"}]}"
	url := "http://api.addressy.com/Capture/Interactive/Find/v1.00/json3.ws?Key=" + loqate_key + "&Countries=US,CA&Language=en&Limit=5&Text=BHAR"

	a, b := normal_curl(url)
	_ = b

	if a != error_msg {
		return "\nDataDog API key valid"
	} else {
		return "\nDataDog API key invalid"
	}
}

func Spotify(spotify_token string) string {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+spotify_token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.Status == "401 Unauthorized" {
		return "Invalid API key"
	} else {
		return "API key valid"
	}
}
func Travis(travis_token string) string {
	req, err := http.NewRequest("GET", "https://api.travis-ci.org/repos", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Travis-API-Version", "3")
	req.Header.Set("Authorization", "token "+travis_token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.Status == "403 Forbidden" {
		return "Invalid API key"
	} else {
		return "API key valid"
	}
}

func Wakatime(wakatime_key string) string {
	url := "https://wakatime.com/api/v1/users/current/projects/?api_key=" + wakatime_key
	a, b := normal_curl(url)
	_ = a

	if b == 200 {
		return "\nWakaTime API Key valid"
	} else {
		return "\nWakaTime API Key invalid"
	}
}
