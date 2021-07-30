package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	red   string = "\033[31m"
	green string = "\033[32m"
	white string = "\033[97m"
)

func parse_json() []string {
	listname := []string{
		"Amplitude API Key",
		"Asana Access token",
		"Bing Maps API Key",
		"Bit.ly Access token",
		"Block.io API Key",
		"Branch.IO Key and Secret",
		"BrowserStack Access Key",
		"Buildkite Access token",
		"ButterCMS API Key",
		"Calendarific API Key",
		"Calendly API Key",
		"CircleCI Access Token",
		"Coinmarketcap API Key",
		"Covalent API Key",
		"Cryptocompare API Key",
		"DataDog API key",
		"EAN-Search Access Token",
		"Facebook Access Token",
		"Fungenerator API Key",
		"Github Token",
		"Heroku API Key",
		"Holiday API Key",
		"Hubspot API Key",
		"Ipstack API Key",
		"Iterable API Key",
		"Jumpcloud API key",
		"Localytics API and Secret key",
		"Lokalise API Key",
		"Loqate API key",
		"MailGun API Key",
		"Mapbox API Key",
		"Pagerduty API token",
		"Pendo Integration Key",
		"PivotalTracker API Token",
		"Razorpay API key and secret key",
		"SauceLabs Username and access Key",
		"SendGrid API Token",
		"Spotify Access Token",
		"Stripe Live Token",
		"Travis CI API token",
		"Twilio Account_sid and Auth token",
		"Visual Studio App Center API Token",
		"WPEngine API Key",
		"WakaTime API Key",
		"Youtube API Key",
		"Zendesk Access token"}
	return listname
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
func basic_auth(url string, username string, password string, need_headers int, set_header string, value_header string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, password)
	if need_headers == 1 {
		req.Header.Set(set_header, value_header)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func Amplitude(amplitude_key string, amplitude_secret string) string {
	url := "https://amplitude.com/api/2/export?start=20200201T5&end=20210203T20"
	a := basic_auth(url, amplitude_key, amplitude_secret, 0, "", "")

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

func Block(block_api string) string {
	url := "https://block.io/api/v2/get_balance/?api_key=" + block_api
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[4], b)
}

func Branch(branch_key string, branch_secret string) string {
	url := "https://api2.branch.io/v1/app/" + branch_key + "?branch_secret=" + branch_secret
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[5], b)
}

func Browserstack(browserstack_username string, browserstack_key string) string {
	url := "https://api.browserstack.com/automate/plan.json"
	a := basic_auth(url, browserstack_username, browserstack_key, 0, "", "")

	return output_status(parse_json()[6], a)
}

func Buildkite(buildkite_access string) string {
	url := "https://api.buildkite.com/v2/user"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+buildkite_access)

	return output_status(parse_json()[7], b)
}

func ButterCMS(buttercms_token string) string {
	url := "https://api.buttercms.com/v2/posts/?auth_token=" + buttercms_token
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[8], b)
}

func Calendarific(calendarific_api string) string {
	url := "https://calendarific.com/api/v2/holidays?api_key=" + calendarific_api
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[9], b)
}

func Calendly(calendly_key string) string {
	url := "https://calendly.com/api/v1/users/me"
	_, b := normal_curl(url, 1, "X-TOKEN", calendly_key)

	return output_status(parse_json()[10], b)
}

func Circleci(circleci_token string) string {
	url := "https://circleci.com/api/v1.1/me?circle-token=" + circleci_token
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[11], b)
}

func Coinmarketcap(coinmarketcap_api string) string {
	url := "https://pro-api.coinmarketcap.com/v1/global-metrics/quotes/historical?CMC_PRO_API_KEY=" + coinmarketcap_api
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[12], b)
}

func Covalent(covalent_api string) string {
	url := "https://api.covalenthq.com/v1/chains/?key=" + covalent_api
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[13], b)
}

func Cryptocompare(cryptocompare_api string) string {
	error_msg := "{\"Response\":\"Error\",\"Message\":\"You need a valid auth key or api key to access this endpoint\",\"HasWarning\":false,\"Type\":1,\"RateLimit\":{},\"Data\":{}}"
	url := "https://min-api.cryptocompare.com/data/blockchain/latest?fsym=BTC&api_key=" + cryptocompare_api

	a, _ := normal_curl(url, 0, "", "")

	return output_error_msg(parse_json()[14], a, error_msg) //ini 12
}

func Datadog(dog_api string, dog_application string) string {
	url := "https://api.datadoghq.com/api/v1/dashboard?api_key=" + dog_api + "&application_key=" + dog_application
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[15], b)
}

func Ean_search(ean_token string) string {
	url := "https://api.ean-search.org/api?token=" + ean_token + "&op=product-search&format=json&name=Bananaboat"
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[16], b)
}

func Facebook(facebook_token string) string {
	url := "https://developers.facebook.com/tools/debug/accesstoken/?access_token=" + facebook_token + "&version=v3.2"
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[17], b)
}

func Fungenerator(fun_api string) string {
	url := "https://api.fungenerators.com/qrcode/text?text=cok&format=png"
	_, b := normal_curl(url, 0, "X-Fungenerators-Api-Secret", fun_api)

	return output_status(parse_json()[18], b)
}

func Github(github_username string, github_token string) string {
	url := "https://api.github.com/user"
	a := basic_auth(url, github_username, github_token, 0, "", "")

	return output_status(parse_json()[19], a)
}

func Heroku(heroku_key string) string {
	url := "https://api.heroku.com/apps"
	_, b := normal_curl(url, 98, "Authorization", "Bearer "+heroku_key)

	return output_status(parse_json()[20], b)
}

func Holiday(holiday_api string) string {
	url := "https://holidayapi.com/v1/holidays?pretty&key=" + holiday_api + "&country=ID&year=2020"
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[21], b)
}

func Hubspot(hubspot_key string) string {
	url := "https://api.hubapi.com/owners/v2/owners?hapikey=" + hubspot_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[22], b)
}

func Ipstack(ipstack_key string) string {
	error_msg := "{\"success\":false,\"error\":{\"code\":101,\"type\":\"invalid_access_key\",\"info\":\"You have not supplied a valid API Access Key. [Technical Support: support@apilayer.com]\"}}"
	url := "https://api.ipstack.com/134.201.250.155?access_key=" + ipstack_key

	a, _ := normal_curl(url, 0, "", "")

	return output_error_msg(parse_json()[23], a, error_msg)
}

func Iterable(iterable_key string) string {
	url := "https://api.iterable.com/api/export/data.json?dataTypeName=emailSend&range=Today&onlyFields=List.empty"
	_, b := normal_curl(url, 1, "Api_Key", iterable_key)

	return output_status(parse_json()[24], b)
}

func Jumpcloud(jumpcloud_key string) string {
	url := "https://console.jumpcloud.com/api/v2/systemgroups"
	_, b := normal_curl(url, 1, "x-api-key", jumpcloud_key)

	return output_status(parse_json()[25], b)
}

func Localytics(localytics_api string, localytics_secret string) string {
	url := "https://api.localytics.com/v1/apps/"
	a := basic_auth(url, localytics_api, localytics_secret, 1, "Accept", "application/vnd.localytics.v1+hal+json")

	return output_status(parse_json()[26], a)
}

func Lokalise(lokalise_key string) string {
	url := "https://api.lokalise.com/api2/projects/"
	_, b := normal_curl(url, 1, "X-Api-Token", lokalise_key)

	return output_status(parse_json()[27], b)
}

func Loqate(loqate_key string) string {
	error_msg := "{\"Items\":[{\"Error\":\"2\",\"Description\":\"Unknown key\",\"Cause\":\"The key you are using to access the service was not found.\",\"Resolution\":\"Please check that the key is correct. It should be in the form AA11-AA11-AA11-AA11.\"}]}"
	url := "http://api.addressy.com/Capture/Interactive/Find/v1.00/json3.ws?Key=" + loqate_key + "&Countries=US,CA&Language=en&Limit=5&Text=BHAR"
	a, _ := normal_curl(url, 0, "", "")

	return output_error_msg(parse_json()[28], a, error_msg)
}
func Mailgun(mailgun_key string) string {
	url := "https://api.mailgun.net/v3/domains"
	a := basic_auth(url, "api", mailgun_key, 0, "", "")

	return output_status(parse_json()[29], a)
}

func Mapbox(mapbox_key string) string {
	url := "https://api.mapbox.com/geocoding/v5/mapbox.places/Los%20Angeles.json?access_token=" + mapbox_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[30], b)
}

func Pagerduty(pagerduty_key string) string {
	url := "https://api.pagerduty.com/schedules"
	_, b := normal_curl(url, 97, "Authorization", "Token token="+pagerduty_key)

	return output_status(parse_json()[31], b)
}

func Pendo(pendo_key string) string {
	url := "https://app.pendo.io/api/v1/feature"
	_, b := normal_curl(url, 1, "x-pendo-integration-key", pendo_key)

	return output_status(parse_json()[32], b)
}

func Pivotaltracker(pivotaltracker_token string) string {
	url := "https://www.pivotaltracker.com/services/v5/stories/555"
	_, b := normal_curl(url, 1, "X-TrackerToken", pivotaltracker_token)

	return output_status(parse_json()[33], b)
}

func Razorpay(razor_key string, razor_secret string) string {
	url := "https://api.razorpay.com/v1/payments"
	a := basic_auth(url, razor_key, razor_secret, 0, "", "")

	return output_status(parse_json()[34], a)
}

func Sauce_labs(sauce_username string, sauce_key string) string {
	url := "https://saucelabs.com/rest/v1/users/" + sauce_username
	a := basic_auth(url, sauce_username, sauce_key, 0, "", "")

	return output_status(parse_json()[35], a)
}

func Sendgrid(sendgrid_token string) string {
	url := "https://api.sendgrid.com/v3/scopes"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+sendgrid_token)

	return output_status(parse_json()[36], b)
}

func Spotify(spotify_token string) string {
	url := "https://api.spotify.com/v1/me"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+spotify_token)

	return output_status(parse_json()[37], b)
}

func Stripe(stripe_token string) string {
	url := "https://api.stripe.com/v1/charges"
	a := basic_auth(url, stripe_token, "", 0, "", "")

	return output_status(parse_json()[38], a)
}

func Travis(travis_token string) string {
	url := "https://api.travis-ci.org/repos"
	_, b := normal_curl(url, 99, "Authorization", "Bearer "+travis_token)

	return output_status(parse_json()[39], b)
}

func Twilio(twilio_sid string, twilio_token string) string {
	url := "https://api.twilio.com/2010-04-01/Accounts.json"
	a := basic_auth(url, twilio_sid, twilio_token, 0, "", "")

	return output_status(parse_json()[40], a)
}

func Appcenter(appcenter_token string) string {
	url := "https://api.appcenter.ms/v0.1/apps"
	_, b := normal_curl(url, 1, "X-Api-Token", appcenter_token)

	return output_status(parse_json()[41], b)
}

func Youtube(youtube_key string) string {
	url := "https://www.googleapis.com/youtube/v3/activities?part=contentDetails&maxResults=25&channelId=UC-lHJZR3Gqxm24_Vd_AJ5Yw&key=" + youtube_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[42], b)
}

func Wakatime(wakatime_key string) string {
	url := "https://wakatime.com/api/v1/users/current/projects/?api_key=" + wakatime_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[43], b)
}

func WPEngine(wp_account string, wp_key string) string {
	url := "https://api.wpengine.com/1.2/?method=site&account_name=" + wp_account + "&wpe_apikey=" + wp_key
	_, b := normal_curl(url, 0, "", "")

	return output_status(parse_json()[44], b)
}

func Zendesk(zendesk_subdomain string, zendesk_token string) string {
	url := "https://" + zendesk_subdomain + ".zendesk.com/api/v2/tickets.json"
	_, b := normal_curl(url, 1, "Authorization", "Bearer "+zendesk_token)

	return output_status(parse_json()[45], b)
}
