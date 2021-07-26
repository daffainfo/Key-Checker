package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	b "github.com/daffainfo/Key-Checker/controllers"
)

const (
	red    string = "\033[31m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	blue   string = "\033[34m"
	purple string = "\033[35m"
	cyan   string = "\033[36m"
	white  string = "\033[97m"
)

var (
	input_key  = ""
	input_key2 = ""
	arr        []string
)

type JsonType struct {
	Array []string
}

func show_banner() {
	fmt.Printf(cyan + `
    __ __                 ________              __            
   / //_/__  __  __      / ____/ /_  ___  _____/ /_____  _____
  / ,< / _ \/ / / /_____/ /   / __ \/ _ \/ ___/ //_/ _ \/ ___/
 / /| /  __/ /_/ /_____/ /___/ / / /  __/ /__/ ,< /  __/ /    
/_/ |_\___/\__, /      \____/_/ /_/\___/\___/_/|_|\___/_/     
          /____/                                              
Author = Muhammad Daffa
Version = 1.0.0
` + white)

}

func show_list() {
	file, _ := ioutil.ReadFile("database/data.json")
	_ = json.Unmarshal([]byte(file), &arr)

	for i, checkers := range arr {
		fmt.Printf(green+"%d. %s%s\n", i+1, checkers, white)
	}
}

func ask_singlekey(choose_arr string) string {
	fmt.Printf(blue+"\nYou choose %s checker", choose_arr)
	fmt.Printf("\nInput your %s:%s\n", choose_arr, white)
	fmt.Scan(&input_key)
	return input_key
}

func ask_secondkey(choose_arr string, question1 string, question2 string) (string, string) {
	fmt.Printf(blue+"\nYou choose %s checker", choose_arr)
	fmt.Printf("\nInput your %s:%s\n", question1, white)
	fmt.Scan(&input_key)
	fmt.Printf("\nInput your %s:%s\n", question2, white)
	fmt.Scan(&input_key2)
	return input_key, input_key2
}

func main() {
	var choose int

	show_banner()
	show_list()

	fmt.Printf("%s\nChoose checkers: ", green)
	fmt.Scan(&choose)

	if choose == 1 {
		q1, q2 := ask_secondkey(arr[0], "Amplitude Key", "Amplitude Secret")
		fmt.Println(b.Amplitude(q1, q2))
	} else if choose == 2 {
		fmt.Println(b.Asana(ask_singlekey(arr[1])))
	} else if choose == 3 {
		fmt.Println(b.Bing(ask_singlekey(arr[2])))
	} else if choose == 4 {
		fmt.Println(b.Bitly(ask_singlekey(arr[3])))
	} else if choose == 5 {
		q1, q2 := ask_secondkey(arr[4], "Branch Key", "Branch Secret")
		fmt.Println(b.Branch(q1, q2))
	} else if choose == 6 {
		q1, q2 := ask_secondkey(arr[5], "BrowserStack Username", "BrowserStack Key")
		fmt.Println(b.Browserstack(q1, q2))
	} else if choose == 7 {
		fmt.Println(b.Buildkite(ask_singlekey(arr[6])))
	} else if choose == 8 {
		fmt.Println(b.Buildkite(ask_singlekey(arr[7])))
	} else if choose == 9 {
		fmt.Println(b.Calendly(ask_singlekey(arr[8])))
	} else if choose == 10 {
		fmt.Println(b.Circleci(ask_singlekey(arr[9])))
	} else if choose == 11 {
		q1, q2 := ask_secondkey(arr[10], "Datadog API key", "Datadog Application key")
		fmt.Println(b.Datadog(q1, q2))
	} else if choose == 12 {
		fmt.Println(b.Facebook(ask_singlekey(arr[11])))
	} else if choose == 13 {
		q1, q2 := ask_secondkey(arr[12], "Github Username", "Github Access Token")
		fmt.Println(b.Github(q1, q2))
	} else if choose == 14 {
		fmt.Println(b.Heroku(ask_singlekey(arr[13])))
	} else if choose == 15 {
		fmt.Println(b.Hubspot(ask_singlekey(arr[14])))
	} else if choose == 16 {
		fmt.Println(b.Ipstack(ask_singlekey(arr[15])))
	} else if choose == 17 {
		fmt.Println(b.Iterable(ask_singlekey(arr[16])))
	} else if choose == 18 {
		fmt.Println(b.Jumpcloud(ask_singlekey(arr[17])))
	} else if choose == 19 {
		fmt.Println(b.Lokalise(ask_singlekey(arr[18])))
	} else if choose == 20 {
		fmt.Println(b.Loqate(ask_singlekey(arr[19])))
	} else if choose == 21 {
		fmt.Println(b.Mailgun(ask_singlekey(arr[20])))
	} else if choose == 22 {
		fmt.Println(b.Mapbox(ask_singlekey(arr[21])))
	} else if choose == 23 {
		fmt.Println(b.Pagerduty(ask_singlekey(arr[22])))
	} else if choose == 24 {
		fmt.Println(b.Pendo(ask_singlekey(arr[23])))
	} else if choose == 25 {
		fmt.Println(b.Pivotaltracker(ask_singlekey(arr[24])))
	} else if choose == 26 {
		q1, q2 := ask_secondkey(arr[25], "Razor Key", "Github Razor Secret")
		fmt.Println(b.Razorpay(q1, q2))
	} else if choose == 27 {
		q1, q2 := ask_secondkey(arr[26], "SauceLabs Username", "Access Token")
		fmt.Println(b.Sauce_labs(q1, q2))
	} else if choose == 28 {
		fmt.Println(b.Sendgrid(ask_singlekey(arr[27])))
	} else if choose == 29 {
		fmt.Println(b.Spotify(ask_singlekey(arr[28])))
	} else if choose == 30 {
		fmt.Println(b.Stripe(ask_singlekey(arr[29])))
	} else if choose == 31 {
		fmt.Println(b.Travis(ask_singlekey(arr[30])))
	} else if choose == 32 {
		q1, q2 := ask_secondkey(arr[31], "Twilio SID", "Twilio Token")
		fmt.Println(b.Twilio(q1, q2))
	} else if choose == 33 {
		fmt.Println(b.Appcenter(ask_singlekey(arr[32])))
	} else if choose == 34 {
		fmt.Println(b.Youtube(ask_singlekey(arr[33])))
	} else if choose == 35 {
		fmt.Println(b.Wakatime(ask_singlekey(arr[34])))
	} else if choose == 36 {
		q1, q2 := ask_secondkey(arr[35], "WPEngine Account", "WPEngine Key")
		fmt.Println(b.WPEngine(q1, q2))
	} else if choose == 37 {
		q1, q2 := ask_secondkey(arr[36], "Zendesk Subdomain", "Zendesk Token")
		fmt.Println(b.Zendesk(q1, q2))
	}
}
