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
	fmt.Println("1. Asana Access token")
	fmt.Println("2. Bing Maps API Key")
	fmt.Println("3. Bit.ly Access token")
	fmt.Println("4. Branch.IO Key and Secret")
	fmt.Println("5. Buildkite Access token")
	fmt.Println("6. ButterCMS API Key")
	fmt.Println("7. CircleCI Access Token")
	fmt.Println("8. DataDog API key")
	fmt.Println("9. Loqate API key")
	fmt.Println("10. Spotify Access Token")
	fmt.Println("11. Travis CI API token")
	fmt.Println("12. WakaTime API Key")

	fmt.Println("\nChoose checkers:")
	fmt.Scanln(&choose)

	if choose == 1 {
		var asana_access string
		fmt.Println("\nInput your Access Token")
		fmt.Scanln(&asana_access)
		fmt.Println(b.Asana(asana_access))
	} else if choose == 2 {
		fmt.Println("\nInput your bing API key")
		var bingapi string
		fmt.Scanln(&bingapi)
		fmt.Println(b.Bing(bingapi))
	} else if choose == 3 {
		var access_token string
		fmt.Println("\nInput your access token")
		fmt.Scanln(&access_token)
		fmt.Println(b.Bitly(access_token))
	} else if choose == 4 {
		var branch_key string
		var branch_secret string
		fmt.Println("\nInput your branch key")
		fmt.Scanln(&branch_key)
		fmt.Println("\nInput your branch secret")
		fmt.Scanln(&branch_secret)
		fmt.Println(b.Branch(branch_key, branch_secret))
	} else if choose == 5 {
		var buildkite_access string
		fmt.Println("\nInput your Access Token")
		fmt.Scanln(&buildkite_access)
		fmt.Println(b.Buildkite(buildkite_access))
	} else if choose == 6 {
		var butter_api string
		fmt.Println("\nInput your API key")
		fmt.Scanln(&butter_api)
		fmt.Println(b.ButterCMS(butter_api))
	} else if choose == 7 {
		var circleci_token string
		fmt.Println("\nInput your CircleCI Access Token")
		fmt.Scanln(&circleci_token)
		fmt.Println(b.Circleci(circleci_token))
	} else if choose == 8 {
		var dog_api string
		var dog_application string
		fmt.Println("\nInput your API key")
		fmt.Scanln(&dog_api)
		fmt.Println("\nInput your Application key")
		fmt.Scanln(&dog_application)
		fmt.Println(b.Datadog(dog_api, dog_application))
	} else if choose == 9 {
		var loqate_key string
		fmt.Println("\nInput your Loqate API key")
		fmt.Scanln(&loqate_key)
		fmt.Println(b.Loqate(loqate_key))
	} else if choose == 10 {
		var spotify_token string
		fmt.Println("\nInput your Spotify Access Token")
		fmt.Scanln(&spotify_token)
		fmt.Println(b.Spotify(spotify_token))
	} else if choose == 11 {
		var travis_token string
		fmt.Println("\nInput your Travis Token")
		fmt.Scanln(&travis_token)
		fmt.Println(b.Travis(travis_token))
	} else if choose == 12 {
		var wakatime_key string
		fmt.Println("\nInput your Wakatime API Key")
		fmt.Scanln(&wakatime_key)
		fmt.Println(b.Wakatime(wakatime_key))
	}

}
