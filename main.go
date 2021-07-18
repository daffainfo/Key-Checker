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
	fmt.Println("1. Bing Maps API Key")
	fmt.Println("2. Bit.ly Access token")
	fmt.Println("3. Branch.IO Key and Secret")
	fmt.Println("4. Buildkite Access token")
	fmt.Println("5. ButterCMS API Key")
	fmt.Println("6. CircleCI Access Token")
	fmt.Println("7. DataDog API key")
	fmt.Println("8. Facebook Access Token")
	fmt.Println("9. Loqate API key")
	fmt.Println("10. Spotify Access Token")
	fmt.Println("11. Travis CI API token")
	fmt.Println("12. Youtube API Key")
	fmt.Println("13. WakaTime API Key")

	fmt.Println("\nChoose checkers:")
	fmt.Scan(&choose)

	if choose == 1 {
		fmt.Print("\nInput your bing API key: ")
		var bingapi string
		fmt.Scan(&bingapi)
		fmt.Println(b.Bing(bingapi))
	} else if choose == 2 {
		var access_token string
		fmt.Println("\nInput your access token: ")
		fmt.Scan(&access_token)
		fmt.Println(b.Bitly(access_token))
	} else if choose == 3 {
		var branch_key string
		var branch_secret string
		fmt.Println("\nInput your branch key: ")
		fmt.Scan(&branch_key)
		fmt.Println("\nInput your branch secret: ")
		fmt.Scan(&branch_secret)
		fmt.Println(b.Branch(branch_key, branch_secret))
	} else if choose == 4 {
		var buildkite_access string
		fmt.Println("\nInput your Access Token: ")
		fmt.Scan(&buildkite_access)
		fmt.Println(b.Buildkite(buildkite_access))
	} else if choose == 5 {
		var butter_api string
		fmt.Println("\nInput your API key: ")
		fmt.Scan(&butter_api)
		fmt.Println(b.ButterCMS(butter_api))
	} else if choose == 6 {
		var circleci_token string
		fmt.Println("\nInput your CircleCI Access Token: ")
		fmt.Scan(&circleci_token)
		fmt.Println(b.Circleci(circleci_token))
	} else if choose == 7 {
		var dog_api string
		var dog_application string
		fmt.Println("\nInput your API key: ")
		fmt.Scan(&dog_api)
		fmt.Println("\nInput your Application key: ")
		fmt.Scan(&dog_application)
		fmt.Println(b.Datadog(dog_api, dog_application))
	} else if choose == 8 {
		var loqate_key string
		fmt.Println("\nInput your Loqate API key: ")
		fmt.Scan(&loqate_key)
		fmt.Println(b.Loqate(loqate_key))
	} else if choose == 9 {
		var spotify_token string
		fmt.Println("\nInput your Spotify Access Token: ")
		fmt.Scan(&spotify_token)
		fmt.Println(b.Spotify(spotify_token))
	} else if choose == 10 {
		var travis_token string
		fmt.Println("\nInput your Travis Token: ")
		fmt.Scan(&travis_token)
		fmt.Println(b.Travis(travis_token))
	} else if choose == 11 {
		var youtube_key string
		fmt.Println("\nInput your Youtube API Key: ")
		fmt.Scan(&youtube_key)
		fmt.Println(b.Youtube(youtube_key))
	} else if choose == 12 {
		var wakatime_key string
		fmt.Println("\nInput your Wakatime API Key: ")
		fmt.Scan(&wakatime_key)
		fmt.Println(b.Wakatime(wakatime_key))
	}

}
