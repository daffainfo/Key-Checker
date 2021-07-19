package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
)

//normal curl
func normal_curl(url string, need_headers int, set_header string, value_header string) (string, int) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if need_headers == 1 {
		req.Header.Set(set_header, value_header)
	} else if need_headers == 99 {
		req.Header.Set(set_header, value_header)
		req.Header.Set("Travis-API-Version", "3")
	} else if need_headers == 98 {
		req.Header.Set(set_header, value_header)
		req.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	} else if need_headers == 97 {
		req.Header.Set(set_header, value_header)
		req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()
	return string(body), req.StatusCode
}

//basic auth
func basic_auth(url string, username string, password string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func Amplitude(amplitude_key string, amplitude_secret string) string {
	url := "https://amplitude.com/api/2/export?start=20200201T5&end=20210203T20"
	a := basic_auth(url, amplitude_key, amplitude_secret)

	if a == 200 {
		return "\nAmplitude API Keys valid"
	} else {
		return "\nAmplitude API Keys invalid"
	}
}

func Asana(asana_token string) string {
	url := "https://app.asana.com/api/1.0/users/me"
	a, b := normal_curl(url, 1, "Authorization", "Bearer "+asana_token)
	_ = a

	if b == 200 {
		return "\nAsana Access token valid"
	} else {
		return "\nAsana Access token invalid"
	}
}

func Bing(bingapi string) string {
	url := "https://dev.virtualearth.net/REST/v1/Locations?CountryRegion=US&adminDistrict=WA&locality=Somewhere&postalCode=98001&addressLine=100%20Main%20St.&key=" + bingapi
	a, b := normal_curl(url, 0, "", "")
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
	a, b := normal_curl(url, 0, "", "")

	if a != error_msg && b == 200 {
		return "\nBit.ly Access token valid"
	} else {
		return "\nBit.ly Access token invalid"
	}
}

func Branch(branch_key string, branch_secret string) string {
	url := "https://api2.branch.io/v1/app/" + branch_key + "?branch_secret=" + branch_secret
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nBranch.IO Key and Secret valid"
	} else {
		return "\nBranch.IO Key and Secret invalid"
	}
}

func Browserstack(browserstack_username string, browserstack_key string) string {
	url := "https://api.browserstack.com/automate/plan.json"
	a := basic_auth(url, browserstack_username, browserstack_key)

	if a == 200 {
		return "\nBrowserStack Access Key valid"
	} else {
		return "\nBrowserStack Access Key invalid"
	}
}

func Buildkite(buildkite_access string) string {
	url := "https://api.buildkite.com/v2/user"
	a, b := normal_curl(url, 1, "Authorization", "Bearer "+buildkite_access)
	_ = a

	if b == 200 {
		return "\nBuildkite Access token valid"
	} else {
		return "\nBuildkite Access token invalid"
	}
}

func ButterCMS(branch_key string) string {
	url := "https://api.buttercms.com/v2/posts/?auth_token=" + branch_key
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nButterCMS API Key valid"
	} else {
		return "\nButterCMS API Key invalid"
	}
}

func Calendly(calendly_key string) string {
	url := "https://calendly.com/api/v1/users/me"
	a, b := normal_curl(url, 1, "X-TOKEN", calendly_key)
	_ = a

	if b == 200 {
		return "\nCalendly API Key valid"
	} else {
		return "\nCalendly API Key invalid"
	}
}

func Circleci(circleci_token string) string {
	url := "https://circleci.com/api/v1.1/me?circle-token=" + circleci_token
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nCircleCI Access Token valid"
	} else {
		return "\nCircleCI Access Token invalid"
	}
}

func Datadog(dog_api string, dog_application string) string {
	url := "https://api.datadoghq.com/api/v1/dashboard?api_key=" + dog_api + "&application_key=" + dog_application
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nDataDog API key valid"
	} else {
		return "\nDataDog API key invalid"
	}
}

func Facebook(facebook_token string) string {
	url := "https://developers.facebook.com/tools/debug/accesstoken/?access_token=" + facebook_token + "&version=v3.2"
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nFacebook Access Token valid"
	} else {
		return "\nFacebook Access Token invalid"
	}
}

func Github(github_username string, github_token string) string {
	url := "https://api.github.com/user"
	a := basic_auth(url, github_username, github_token)

	if a == 200 {
		return "\nGithub Token valid"
	} else {
		return "\nGithub Token invalid"
	}
}

func Heroku(heroku_key string) string {
	url := "https://api.heroku.com/apps"
	a, b := normal_curl(url, 98, "Authorization", "Bearer "+heroku_key)
	_ = a

	if b == 200 {
		return "\nHeroku API Key valid"
	} else {
		return "\nHeroku API Key invalid"
	}
}

func Hubspot(hubspot_key string) string {
	url := "https://api.hubapi.com/owners/v2/owners?hapikey=" + hubspot_key
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nHubspot API Key valid"
	} else {
		return "\nHubspot API Key invalid"
	}
}

func Ipstack(ipstack_key string) string {
	error_msg := "{\"success\":false,\"error\":{\"code\":101,\"type\":\"invalid_access_key\",\"info\":\"You have not supplied a valid API Access Key. [Technical Support: support@apilayer.com]\"}}"
	url := "https://api.ipstack.com/134.201.250.155?access_key=" + ipstack_key

	a, b := normal_curl(url, 0, "", "")
	_ = b

	if a != error_msg {
		return "\nIpstack API Key valid"
	} else {
		return "\nIpstack API Key invalid"
	}
}

func Iterable(iterable_key string) string {
	url := "https://api.iterable.com/api/export/data.json?dataTypeName=emailSend&range=Today&onlyFields=List.empty"
	a, b := normal_curl(url, 1, "Api_Key", iterable_key)
	_ = a

	if b == 200 {
		return "\nIterable API Key valid"
	} else {
		return "\nIterable API Key invalid"
	}
}

func Jumpcloud(jumpcloud_key string) string {
	url := "https://console.jumpcloud.com/api/v2/systemgroups"
	a, b := normal_curl(url, 1, "x-api-key", jumpcloud_key)
	_ = a

	if b == 200 {
		return "\nJumpCloud API Key valid"
	} else {
		return "\nJumpCloud API Key invalid"
	}
}

func Lokalise(lokalise_key string) string {
	url := "https://api.lokalise.com/api2/projects/"
	a, b := normal_curl(url, 1, "X-Api-Token", lokalise_key)
	_ = a

	if b == 200 {
		return "\nLokalise API key valid"
	} else {
		return "\nLokalise API key invalid"
	}
}

func Loqate(loqate_key string) string {
	error_msg := "{\"Items\":[{\"Error\":\"2\",\"Description\":\"Unknown key\",\"Cause\":\"The key you are using to access the service was not found.\",\"Resolution\":\"Please check that the key is correct. It should be in the form AA11-AA11-AA11-AA11.\"}]}"
	url := "http://api.addressy.com/Capture/Interactive/Find/v1.00/json3.ws?Key=" + loqate_key + "&Countries=US,CA&Language=en&Limit=5&Text=BHAR"
	a, b := normal_curl(url, 0, "", "")
	_ = b

	if a != error_msg {
		return "\nLoqate API key valid"
	} else {
		return "\nLoqate API key invalid"
	}
}
func Mailgun(mailgun_key string) string {
	url := "https://api.mailgun.net/v3/domains"
	a := basic_auth(url, "api", mailgun_key)

	if a == 200 {
		return "\nMailgun API key valid"
	} else {
		return "\nMailgun API key invalid"
	}
}

func Mapbox(mapbox_key string) string {
	url := "https://api.mapbox.com/geocoding/v5/mapbox.places/Los%20Angeles.json?access_token=" + mapbox_key
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nMapbox API key valid"
	} else {
		return "\nMapbox API key invalid"
	}
}

func Pagerduty(pagerduty_key string) string {
	url := "https://api.pagerduty.com/schedules"
	a, b := normal_curl(url, 97, "Authorization", "Token token="+pagerduty_key)
	_ = a

	if b == 200 {
		return "\nPagerduty API Key valid"
	} else {
		return "\nPagerduty API Key invalid"
	}
}

func Pendo(pendo_key string) string {
	url := "https://app.pendo.io/api/v1/feature"
	a, b := normal_curl(url, 1, "x-pendo-integration-key", pendo_key)
	_ = a

	if b == 200 {
		return "\nPendo Integration Key valid"
	} else {
		return "\nPendo Integration Key invalid"
	}
}

func Pivotaltracker(pivotaltracker_token string) string {
	url := "https://www.pivotaltracker.com/services/v5/stories/555"
	a, b := normal_curl(url, 1, "X-TrackerToken", pivotaltracker_token)
	_ = a

	if b == 200 {
		return "\nPivotalTracker API Token valid"
	} else {
		return "\nPivotalTracker API Token invalid"
	}
}

func Razorpay(razor_key string, razor_secret string) string {
	url := "https://api.razorpay.com/v1/payments"
	a := basic_auth(url, razor_key, razor_secret)

	if a == 200 {
		return "\nRazorpay API key and Secret key valid"
	} else {
		return "\nRazorpay API key and Secret key invalid"
	}
}

func Sauce_labs(sauce_username string, sauce_key string) string {
	url := "https://saucelabs.com/rest/v1/users/" + sauce_username
	a := basic_auth(url, sauce_username, sauce_key)

	if a == 200 {
		return "\nSauceLabs Username and access Key valid"
	} else {
		return "\nSauceLabs Username and access Key invalid"
	}
}

func Sendgrid(sendgrid_token string) string {
	url := "https://api.sendgrid.com/v3/scopes"
	a, b := normal_curl(url, 1, "Authorization", "Bearer "+sendgrid_token)
	_ = a

	if b == 200 {
		return "\nSendGrid API Token valid"
	} else {
		return "\nSendGrid API Token invalid"
	}
}

func Spotify(spotify_token string) string {
	url := "https://api.spotify.com/v1/me"
	a, b := normal_curl(url, 1, "Authorization", "Bearer "+spotify_token)
	_ = a

	if b == 200 {
		return "\nSpotify Access Token valid"
	} else {
		return "\nSpotify Access Token invalid"
	}
}

func Stripe(stripe_token string) string {
	url := "https://api.stripe.com/v1/charges"
	a := basic_auth(url, stripe_token, "")

	if a == 200 {
		return "\nStripe API key valid"
	} else {
		return "\nStripe API key invalid"
	}
}

func Travis(travis_token string) string {
	url := "https://api.travis-ci.org/repos"
	a, b := normal_curl(url, 99, "Authorization", "Bearer "+travis_token)
	_ = a

	if b == 200 {
		return "\nTravis CI API token valid"
	} else {
		return "\nTravis CI API token invalid"
	}
}

func Twilio(twilio_sid string, twilio_token string) string {
	url := "https://api.twilio.com/2010-04-01/Accounts.json"
	a := basic_auth(url, twilio_sid, twilio_token)

	if a == 200 {
		return "\nTwilio Account_sid and Auth token valid"
	} else {
		return "\nTwilio Account_sid and Auth token invalid"
	}
}

func Appcenter(appcenter_token string) string {
	url := "https://api.appcenter.ms/v0.1/apps"
	a, b := normal_curl(url, 1, "X-Api-Token", appcenter_token)
	_ = a

	if b == 200 {
		return "\nVisual Studio App Center API Token valid"
	} else {
		return "\nVisual Studio App Center API Token invalid"
	}
}

func Youtube(youtube_key string) string {
	url := "https://www.googleapis.com/youtube/v3/activities?part=contentDetails&maxResults=25&channelId=UC-lHJZR3Gqxm24_Vd_AJ5Yw&key=" + youtube_key
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nYoutube API Key valid"
	} else {
		return "\nYoutube API Key invalid"
	}
}

func Wakatime(wakatime_key string) string {
	url := "https://wakatime.com/api/v1/users/current/projects/?api_key=" + wakatime_key
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nWakaTime API Key valid"
	} else {
		return "\nWakaTime API Key invalid"
	}
}

func WPEngine(wp_account string, wp_key string) string {
	url := "https://api.wpengine.com/1.2/?method=site&account_name=" + wp_account + "&wpe_apikey=" + wp_key
	a, b := normal_curl(url, 0, "", "")
	_ = a

	if b == 200 {
		return "\nWPEngine API Key valid"
	} else {
		return "\nWPEngine API Key invalid"
	}
}

func Zendesk(zendesk_subdomain string, zendesk_token string) string {
	url := "https://" + zendesk_subdomain + ".zendesk.com/api/v2/tickets.json"
	a, b := normal_curl(url, 1, "Authorization", "Bearer "+zendesk_token)
	_ = a

	if b == 200 {
		return "\nZendesk Acces Token valid"
	} else {
		return "\nZendesk Acces Token invalid"
	}
}
