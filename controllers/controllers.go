package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
)

//normal curl
func normal_curl(url string, need_headers string, value_header string) (string, int) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if need_headers == "bearer" {
		req.Header.Set("Authorization", "Bearer "+value_header)
	} else if need_headers == "travis" {
		req.Header.Set("Authorization", "Bearer "+value_header)
		req.Header.Set("Authorization", "Bearer "+value_header)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()
	return string(body), req.StatusCode
}

func Bing(bingapi string) string {
	url := "https://dev.virtualearth.net/REST/v1/Locations?CountryRegion=US&adminDistrict=WA&locality=Somewhere&postalCode=98001&addressLine=100%20Main%20St.&key=" + bingapi
	a, b := normal_curl(url, "", "")
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
	a, b := normal_curl(url, "", "")

	if a != error_msg && b == 200 {
		return "\nBit.ly Access token valid"
	} else {
		return "\nBit.ly Access token invalid"
	}
}

func Branch(branch_key string, branch_secret string) string {
	url := "https://api2.branch.io/v1/app/" + branch_key + "?branch_secret=" + branch_secret

	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nBranch.IO Key and Secret valid"
	} else {
		return "\nBranch.IO Key and Secret invalid"
	}
}

func Buildkite(buildkite_access string) string {
	url := "https://api.buildkite.com/v2/user"

	a, b := normal_curl(url, "bearer", buildkite_access)
	_ = a

	if b == 200 {
		return "\nAsana Access token valid"
	} else {
		return "\nAsana Access token invalid"
	}
}

func ButterCMS(branch_key string) string {
	url := "https://api.buttercms.com/v2/posts/?auth_token=" + branch_key

	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nBranch.IO Key and Secret valid"
	} else {
		return "\nBranch.IO Key and Secret invalid"
	}
}

func Circleci(circleci_token string) string {
	url := "https://circleci.com/api/v1.1/me?circle-token=" + circleci_token

	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nCircleCI Access Token valid"
	} else {
		return "\nCircleCI Access Token invalid"
	}
}

func Datadog(dog_api string, dog_application string) string {
	url := "https://api.datadoghq.com/api/v1/dashboard?api_key=" + dog_api + "&application_key=" + dog_application

	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nDataDog API key valid"
	} else {
		return "\nDataDog API key invalid"
	}
}

func Facebook(facebook_token string) string {
	url := "https://developers.facebook.com/tools/debug/accesstoken/?access_token=" + facebook_token + "&version=v3.2"

	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nFacebook Access Token valid"
	} else {
		return "\nFacebook Access Token invalid"
	}
}

func Loqate(loqate_key string) string {
	error_msg := "{\"Items\":[{\"Error\":\"2\",\"Description\":\"Unknown key\",\"Cause\":\"The key you are using to access the service was not found.\",\"Resolution\":\"Please check that the key is correct. It should be in the form AA11-AA11-AA11-AA11.\"}]}"
	url := "http://api.addressy.com/Capture/Interactive/Find/v1.00/json3.ws?Key=" + loqate_key + "&Countries=US,CA&Language=en&Limit=5&Text=BHAR"

	a, b := normal_curl(url, "", "")
	_ = b

	if a != error_msg {
		return "\nDataDog API key valid"
	} else {
		return "\nDataDog API key invalid"
	}
}

func Spotify(spotify_token string) string {
	url := "https://api.spotify.com/v1/me"
	a, b := normal_curl(url, "bearer", spotify_token)
	_ = a

	if b == 200 {
		return "\nSpotify Access Token valid"
	} else {
		return "\nSpotify Access Token invalid"
	}
}

func Travis(travis_token string) string {
	url := "https://api.travis-ci.org/repos"

	a, b := normal_curl(url, "bearer", travis_token)
	_ = a

	if b == 200 {
		return "\nSpotify Access Token valid"
	} else {
		return "\nSpotify Access Token invalid"
	}
}

func Wakatime(wakatime_key string) string {
	url := "https://wakatime.com/api/v1/users/current/projects/?api_key=" + wakatime_key
	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nWakaTime API Key valid"
	} else {
		return "\nWakaTime API Key invalid"
	}
}

func Youtube(youtube_key string) string {
	url := "https://www.googleapis.com/youtube/v3/activities?part=contentDetails&maxResults=25&channelId=UC-lHJZR3Gqxm24_Vd_AJ5Yw&key=" + youtube_key

	a, b := normal_curl(url, "", "")
	_ = a

	if b == 200 {
		return "\nYoutube API Key valid"
	} else {
		return "\nYoutube API Key invalid"
	}
}
