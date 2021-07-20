package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	arr []string
)

const (
	red   string = "\033[31m"
	green string = "\033[32m"
	white string = "\033[97m"
)

func parse_json() []string {
	file, _ := ioutil.ReadFile("database/data.json")
	_ = json.Unmarshal([]byte(file), &arr)
	return arr
}

func output_status(array_list string, code_status int) string {
	if code_status == 200 {
		return "\n" + green + array_list + " valid" + white + "\n"
	} else {
		return "\n" + red + array_list + " invalid" + white + "\n"
	}
}

func output_error_msg(array_list string, error_msg string, msg string) string {
	if error_msg != msg {
		return "\n" + green + array_list + " valid" + white + "\n"
	} else {
		return "\n" + red + array_list + " invalid" + white + "\n"
	}
}

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

	return output_status(parse_json()[0], a)
}

func Asana(asana_token string) string {
	url := "https://app.asana.com/api/1.0/users/me"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+asana_token)

	return output_status(parse_json()[1], b)
}

func Bing(bingapi string) string {
	url := "https://dev.virtualearth.net/REST/v1/Locations?CountryRegion=US&adminDistrict=WA&locality=Somewhere&postalCode=98001&addressLine=100%20Main%20St.&key=" + bingapi
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[2], b)
}

func Bitly(bitly_token string) string {
	error_msg := "{\"status_code\":500,\"status_txt\":\"INVALID_ARG_ACCESS_TOKEN\",\"data\":[]}"
	url := "https://api-ssl.bitly.com/v3/shorten?access_token=" + bitly_token + "&longUrl=https://www.google.com"
	a, _ := normal_curl(url, 0, "", "")

	return output_error_msg(parse_json()[3], a, error_msg)
}

func Branch(branch_key string, branch_secret string) string {
	url := "https://api2.branch.io/v1/app/" + branch_key + "?branch_secret=" + branch_secret
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[4], b)
}

func Browserstack(browserstack_username string, browserstack_key string) string {
	url := "https://api.browserstack.com/automate/plan.json"
	a := basic_auth(url, browserstack_username, browserstack_key)

	return output_status(parse_json()[5], a)
}

func Buildkite(buildkite_access string) string {
	url := "https://api.buildkite.com/v2/user"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+buildkite_access)

	return output_status(parse_json()[6], b)
}

func ButterCMS(branch_key string) string {
	url := "https://api.buttercms.com/v2/posts/?auth_token=" + branch_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[7], b)
}

func Calendly(calendly_key string) string {
	url := "https://calendly.com/api/v1/users/me"
	_, b := normal_curl(url, 1, "X-TOKEN", calendly_key)

	return output_status(parse_json()[8], b)
}

func Circleci(circleci_token string) string {
	url := "https://circleci.com/api/v1.1/me?circle-token=" + circleci_token
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[9], b)
}

func Datadog(dog_api string, dog_application string) string {
	url := "https://api.datadoghq.com/api/v1/dashboard?api_key=" + dog_api + "&application_key=" + dog_application
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[10], b)
}

func Facebook(facebook_token string) string {
	url := "https://developers.facebook.com/tools/debug/accesstoken/?access_token=" + facebook_token + "&version=v3.2"
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[11], b)
}

func Github(github_username string, github_token string) string {
	url := "https://api.github.com/user"
	a := basic_auth(url, github_username, github_token)

	return output_status(parse_json()[12], a)
}

func Heroku(heroku_key string) string {
	url := "https://api.heroku.com/apps"
	_, b := normal_curl(url, 98, "Authorization", "Bearer "+heroku_key)

	return output_status(parse_json()[13], b)
}

func Hubspot(hubspot_key string) string {
	url := "https://api.hubapi.com/owners/v2/owners?hapikey=" + hubspot_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[14], b)
}

func Ipstack(ipstack_key string) string {
	error_msg := "{\"success\":false,\"error\":{\"code\":101,\"type\":\"invalid_access_key\",\"info\":\"You have not supplied a valid API Access Key. [Technical Support: support@apilayer.com]\"}}"
	url := "https://api.ipstack.com/134.201.250.155?access_key=" + ipstack_key

	a, _ := normal_curl(url, 0, "", "")

	return output_error_msg(parse_json()[15], a, error_msg)
}

func Iterable(iterable_key string) string {
	url := "https://api.iterable.com/api/export/data.json?dataTypeName=emailSend&range=Today&onlyFields=List.empty"
	_, b := normal_curl(url, 1, "Api_Key", iterable_key)

	return output_status(parse_json()[16], b)
}

func Jumpcloud(jumpcloud_key string) string {
	url := "https://console.jumpcloud.com/api/v2/systemgroups"
	_, b := normal_curl(url, 1, "x-api-key", jumpcloud_key)

	return output_status(parse_json()[17], b)
}

func Lokalise(lokalise_key string) string {
	url := "https://api.lokalise.com/api2/projects/"
	_, b := normal_curl(url, 1, "X-Api-Token", lokalise_key)

	return output_status(parse_json()[18], b)
}

func Loqate(loqate_key string) string {
	error_msg := "{\"Items\":[{\"Error\":\"2\",\"Description\":\"Unknown key\",\"Cause\":\"The key you are using to access the service was not found.\",\"Resolution\":\"Please check that the key is correct. It should be in the form AA11-AA11-AA11-AA11.\"}]}"
	url := "http://api.addressy.com/Capture/Interactive/Find/v1.00/json3.ws?Key=" + loqate_key + "&Countries=US,CA&Language=en&Limit=5&Text=BHAR"
	a, _ := normal_curl(url, 0, "", "")

	return output_error_msg(parse_json()[19], a, error_msg)
}
func Mailgun(mailgun_key string) string {
	url := "https://api.mailgun.net/v3/domains"
	a := basic_auth(url, "api", mailgun_key)

	return output_status(parse_json()[20], a)
}

func Mapbox(mapbox_key string) string {
	url := "https://api.mapbox.com/geocoding/v5/mapbox.places/Los%20Angeles.json?access_token=" + mapbox_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[21], b)
}

func Pagerduty(pagerduty_key string) string {
	url := "https://api.pagerduty.com/schedules"
	_, b := normal_curl(url, 97, "Authorization", "Token token="+pagerduty_key)

	return output_status(parse_json()[22], b)
}

func Pendo(pendo_key string) string {
	url := "https://app.pendo.io/api/v1/feature"
	_, b := normal_curl(url, 1, "x-pendo-integration-key", pendo_key)

	return output_status(parse_json()[23], b)
}

func Pivotaltracker(pivotaltracker_token string) string {
	url := "https://www.pivotaltracker.com/services/v5/stories/555"
	_, b := normal_curl(url, 1, "X-TrackerToken", pivotaltracker_token)

	return output_status(parse_json()[24], b)
}

func Razorpay(razor_key string, razor_secret string) string {
	url := "https://api.razorpay.com/v1/payments"
	a := basic_auth(url, razor_key, razor_secret)

	return output_status(parse_json()[25], a)
}

func Sauce_labs(sauce_username string, sauce_key string) string {
	url := "https://saucelabs.com/rest/v1/users/" + sauce_username
	a := basic_auth(url, sauce_username, sauce_key)

	return output_status(parse_json()[26], a)
}

func Sendgrid(sendgrid_token string) string {
	url := "https://api.sendgrid.com/v3/scopes"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+sendgrid_token)

	return output_status(parse_json()[27], b)
}

func Spotify(spotify_token string) string {
	url := "https://api.spotify.com/v1/me"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+spotify_token)

	return output_status(parse_json()[28], b)
}

func Stripe(stripe_token string) string {
	url := "https://api.stripe.com/v1/charges"
	a := basic_auth(url, stripe_token, "")

	return output_status(parse_json()[29], a)
}

func Travis(travis_token string) string {
	url := "https://api.travis-ci.org/repos"
	_, b := normal_curl(url, 99, "Authorization", "Bearer "+travis_token)

	return output_status(parse_json()[30], b)
}

func Twilio(twilio_sid string, twilio_token string) string {
	url := "https://api.twilio.com/2010-04-01/Accounts.json"
	a := basic_auth(url, twilio_sid, twilio_token)

	return output_status(parse_json()[31], a)
}

func Appcenter(appcenter_token string) string {
	url := "https://api.appcenter.ms/v0.1/apps"
	_, b := normal_curl(url, 1, "X-Api-Token", appcenter_token)

	return output_status(parse_json()[32], b)
}

func Youtube(youtube_key string) string {
	url := "https://www.googleapis.com/youtube/v3/activities?part=contentDetails&maxResults=25&channelId=UC-lHJZR3Gqxm24_Vd_AJ5Yw&key=" + youtube_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[33], b)
}

func Wakatime(wakatime_key string) string {
	url := "https://wakatime.com/api/v1/users/current/projects/?api_key=" + wakatime_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[34], b)
}

func WPEngine(wp_account string, wp_key string) string {
	url := "https://api.wpengine.com/1.2/?method=site&account_name=" + wp_account + "&wpe_apikey=" + wp_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[35], b)
}

func Zendesk(zendesk_subdomain string, zendesk_token string) string {
	url := "https://" + zendesk_subdomain + ".zendesk.com/api/v2/tickets.json"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+zendesk_token)

	return output_status(parse_json()[36], b)
}
