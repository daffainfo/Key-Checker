package main

import (
	b "Key-Checker/controllers"
	"fmt"
)

func main() {

	var choose int

	fmt.Println(`
    __ __                 ________              __            
   / //_/__  __  __      / ____/ /_  ___  _____/ /_____  _____
  / ,< / _ \/ / / /_____/ /   / __ \/ _ \/ ___/ //_/ _ \/ ___/
 / /| /  __/ /_/ /_____/ /___/ / / /  __/ /__/ ,< /  __/ /    
/_/ |_\___/\__, /      \____/_/ /_/\___/\___/_/|_|\___/_/     
          /____/                                              

  `)
	fmt.Println("List of API Checker")
	fmt.Println("1. Amplitude API Key")
	fmt.Println("2. Asana Access token")
	fmt.Println("3. Bing Maps API Key")
	fmt.Println("4. Bit.ly Access token")
	fmt.Println("5. Branch.IO Key and Secret")
	fmt.Println("6. BrowserStack Access Key")
	fmt.Println("7. Buildkite Access token")
	fmt.Println("8. ButterCMS API Key")
	fmt.Println("9. Calendly API Key")
	fmt.Println("10. CircleCI Access Token")
	fmt.Println("11. DataDog API key")
	fmt.Println("12. Facebook Access Token")
	fmt.Println("13. Github Token")
	fmt.Println("14. Heroku API Key")
	fmt.Println("15. Hubspot API Key")
	fmt.Println("16. Ipstack API Key")
	fmt.Println("17. Iterable API Key")
	fmt.Println("18. Jumpcloud API key")
	fmt.Println("19. Lokalise API Key")
	fmt.Println("20. Loqate API key")
	fmt.Println("21. MailGun API Key")
	fmt.Println("22. Mapbox API Key")
	fmt.Println("23. Pagerduty API token")
	fmt.Println("24. Pendo Integration Key")
	fmt.Println("25. PivotalTracker API Token")
	fmt.Println("26. Razorpay API key and secret key")
	fmt.Println("27. SauceLabs Username and access Key")
	fmt.Println("28. SendGrid API Token")
	fmt.Println("29. Spotify Access Token")
	fmt.Println("30. Stripe Live Token")
	fmt.Println("31. Travis CI API token")
	fmt.Println("32. Twilio Account_sid and Auth token")
	fmt.Println("33. Visual Studio App Center API Token")
	fmt.Println("34. Youtube API Key")
	fmt.Println("35. WakaTime API Key")
	fmt.Println("36. WPEngine API Key")
	fmt.Println("37. Zendesk Access token")

	fmt.Println("\nChoose checkers:")
	fmt.Scan(&choose)

	if choose == 1 {
		var amplitude_key string
		var amplitude_secret string
		fmt.Print("\nInput your Amplitude Key: ")
		fmt.Scan(&amplitude_key)
		fmt.Print("\nInput your Amplitude Secret: ")
		fmt.Scan(&amplitude_secret)
		fmt.Println(b.Amplitude(amplitude_key, amplitude_secret))
	} else if choose == 2 {
		fmt.Print("\nInput your Asana Access token: ")
		var asana_token string
		fmt.Scan(&asana_token)
		fmt.Println(b.Asana(asana_token))
	} else if choose == 3 {
		fmt.Print("\nInput your bing API key: ")
		var bingapi string
		fmt.Scan(&bingapi)
		fmt.Println(b.Bing(bingapi))
	} else if choose == 4 {
		var access_token string
		fmt.Print("\nInput your Bit.ly Access token: ")
		fmt.Scan(&access_token)
		fmt.Println(b.Bitly(access_token))
	} else if choose == 5 {
		var branch_key string
		var branch_secret string
		fmt.Print("\nInput your branch key: ")
		fmt.Scan(&branch_key)
		fmt.Print("\nInput your branch secret: ")
		fmt.Scan(&branch_secret)
		fmt.Println(b.Branch(branch_key, branch_secret))
	} else if choose == 6 {
		var browserstack_username string
		var browserstack_key string
		fmt.Print("\nInput your BrowserStack Username: ")
		fmt.Scan(&browserstack_username)
		fmt.Print("\nInput your BrowserStack Key: ")
		fmt.Scan(&browserstack_key)
		fmt.Println(b.Browserstack(browserstack_username, browserstack_key))
	} else if choose == 7 {
		var buildkite_access string
		fmt.Print("\nInput your Buildkite Access token: ")
		fmt.Scan(&buildkite_access)
		fmt.Println(b.Buildkite(buildkite_access))
	} else if choose == 8 {
		var butter_api string
		fmt.Print("\nInput your ButterCMS API Key: ")
		fmt.Scan(&butter_api)
		fmt.Println(b.ButterCMS(butter_api))
	} else if choose == 9 {
		var calendly_key string
		fmt.Print("\nInput your Calendly API Key: ")
		fmt.Scan(&calendly_key)
		fmt.Println(b.Calendly(calendly_key))
	} else if choose == 10 {
		var circleci_token string
		fmt.Print("\nInput your CircleCI Access Token: ")
		fmt.Scan(&circleci_token)
		fmt.Println(b.Circleci(circleci_token))
	} else if choose == 11 {
		var dog_api string
		var dog_application string
		fmt.Print("\nInput your Datadog API key: ")
		fmt.Scan(&dog_api)
		fmt.Print("\nInput your Datadog Application key: ")
		fmt.Scan(&dog_application)
		fmt.Println(b.Datadog(dog_api, dog_application))
	} else if choose == 12 {
		var facebook_token string
		fmt.Println("\nInput your Facebook Access Token: ")
		fmt.Scan(&facebook_token)
		fmt.Println(b.Facebook(facebook_token))
	} else if choose == 13 {
		var github_username string
		var github_token string
		fmt.Print("\nInput your Github Username: ")
		fmt.Scan(&github_username)
		fmt.Print("\nInput your Github Access Token: ")
		fmt.Scan(&github_token)
		fmt.Println(b.Github(github_username, github_token))
	} else if choose == 14 {
		var heroku_key string
		fmt.Print("\nInput your Heroku API Key: ")
		fmt.Scan(&heroku_key)
		fmt.Println(b.Heroku(heroku_key))
	} else if choose == 15 {
		var hubspot_key string
		fmt.Print("\nInput your Hubspot API Key: ")
		fmt.Scan(&hubspot_key)
		fmt.Println(b.Hubspot(hubspot_key))
	} else if choose == 16 {
		var ipstack_key string
		fmt.Print("\nInput your Ipstack API Key: ")
		fmt.Scan(&ipstack_key)
		fmt.Println(b.Ipstack(ipstack_key))
	} else if choose == 17 {
		var iterable_key string
		fmt.Print("\nInput your Iterable API Key: ")
		fmt.Scan(&iterable_key)
		fmt.Println(b.Iterable(iterable_key))
	} else if choose == 18 {
		var jumpcloud_key string
		fmt.Print("\nInput your JumpCloud API Key: ")
		fmt.Scan(&jumpcloud_key)
		fmt.Println(b.Jumpcloud(jumpcloud_key))
	} else if choose == 19 {
		var lokalise_key string
		fmt.Print("\nInput your Lokalise API key: ")
		fmt.Scan(&lokalise_key)
		fmt.Println(b.Lokalise(lokalise_key))
	} else if choose == 20 {
		var loqate_key string
		fmt.Print("\nInput your Loqate API key: ")
		fmt.Scan(&loqate_key)
		fmt.Println(b.Loqate(loqate_key))
	} else if choose == 21 {
		var mailgun_key string
		fmt.Print("\nInput your Mailgun API key: ")
		fmt.Scan(&mailgun_key)
		fmt.Println(b.Mailgun(mailgun_key))
	} else if choose == 22 {
		var mapbox_key string
		fmt.Print("\nInput your Mapbox API key: ")
		fmt.Scan(&mapbox_key)
		fmt.Println(b.Mapbox(mapbox_key))
	} else if choose == 23 {
		var pagerduty_key string
		fmt.Print("\nInput your Pagerduty API token: ")
		fmt.Scan(&pagerduty_key)
		fmt.Println(b.Pagerduty(pagerduty_key))
	} else if choose == 24 {
		var pendo_key string
		fmt.Print("\nInput your Pendo Integration Key: ")
		fmt.Scan(&pendo_key)
		fmt.Println(b.Pendo(pendo_key))
	} else if choose == 25 {
		var pivotaltracker_token string
		fmt.Print("\nInput your PivotalTracker API Token: ")
		fmt.Scan(&pivotaltracker_token)
		fmt.Println(b.Pivotaltracker(pivotaltracker_token))
	} else if choose == 26 {
		var razor_key string
		var razor_secret string
		fmt.Print("\nInput your Razor Key: ")
		fmt.Scan(&razor_key)
		fmt.Print("\nInput your Razor Secret: ")
		fmt.Scan(&razor_secret)
		fmt.Println(b.Razorpay(razor_key, razor_secret))
	} else if choose == 27 {
		var sauce_username string
		var sauce_key string
		fmt.Print("\nInput your SauceLabs Username: ")
		fmt.Scan(&sauce_username)
		fmt.Print("\nInput your SauceLabs Access Token: ")
		fmt.Scan(&sauce_key)
		fmt.Println(b.Sauce_labs(sauce_username, sauce_key))
	} else if choose == 28 {
		var sendgrid_token string
		fmt.Print("\nInput your SendGrid API Token: ")
		fmt.Scan(&sendgrid_token)
		fmt.Println(b.Sendgrid(sendgrid_token))
	} else if choose == 29 {
		var spotify_token string
		fmt.Print("\nInput your Spotify Access Token: ")
		fmt.Scan(&spotify_token)
		fmt.Println(b.Spotify(spotify_token))
	} else if choose == 30 {
		var stripe_token string
		fmt.Print("\nInput your Stripe Live Token: ")
		fmt.Scan(&stripe_token)
		fmt.Println(b.Stripe(stripe_token))
	} else if choose == 31 {
		var travis_token string
		fmt.Print("\nInput your Travis Token: ")
		fmt.Scan(&travis_token)
		fmt.Println(b.Travis(travis_token))
	} else if choose == 32 {
		var twilio_sid string
		var twilio_token string
		fmt.Print("\nInput your Twilio SID: ")
		fmt.Scan(&twilio_sid)
		fmt.Print("\nInput your Twilio Token: ")
		fmt.Scan(&twilio_token)
		fmt.Println(b.Twilio(twilio_sid, twilio_token))
	} else if choose == 33 {
		var appcenter_token string
		fmt.Print("\nInput your Visual Studio App Center API Token: ")
		fmt.Scan(&appcenter_token)
		fmt.Println(b.Appcenter(appcenter_token))
	} else if choose == 34 {
		var youtube_key string
		fmt.Print("\nInput your Youtube API Key: ")
		fmt.Scan(&youtube_key)
		fmt.Println(b.Youtube(youtube_key))
	} else if choose == 35 {
		var wakatime_key string
		fmt.Print("\nInput your Wakatime API Key: ")
		fmt.Scan(&wakatime_key)
		fmt.Println(b.Wakatime(wakatime_key))
	} else if choose == 36 {
		var wp_account string
		var wp_key string
		fmt.Print("\nInput your WPEngine Account: ")
		fmt.Scan(&wp_account)
		fmt.Print("\nInput your WPEngine Key: ")
		fmt.Scan(&wp_key)
		fmt.Println(b.WPEngine(wp_account, wp_key))
	} else if choose == 37 {
		var zendesk_subdomain string
		var zendesk_token string
		fmt.Print("\nInput Zendesk Subdomain: ")
		fmt.Scan(&zendesk_subdomain)
		fmt.Print("\nInput Zendesk Token: ")
		fmt.Scan(&zendesk_token)
		fmt.Println(b.Zendesk(zendesk_subdomain, zendesk_token))
	}
}
