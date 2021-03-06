//
//  BSD 3-Clause License
//
// Copyright (c) 2018, Sidharth Mishra
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
// * Neither the name of the copyright holder nor the names of its
//  contributors may be used to endorse or promote products derived from
//  this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// main.go
// @author Sidharth Mishra Twitcher
// @created Tue Aug 29 2017 17:45:04 GMT-0700 (PDT)
// @last-modified Mon Mar 26 2018 21:00:11 GMT-0700 (PDT)
//

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sidmishraw/twitcher/twitch"
)

// appName is the name of this app - known to Twitch - twitcher was taken :/
const appName = "twitcher-be"

// TWITCH_BASE_URL is the base API URL for Twitch API v5.
const TWITCH_BASE_URL = "https://api.twitch.tv/kraken"

// RES_LIMIT is the total number of results allowed per page.
const RES_LIMIT = 25

// The client ID generated for this app
var clientID = "64co18zsfvj6plbqy5v2j7x5r6auza"

// The client secret key generated for this app
var clientSecret = "savr6m36r8vh4fsc9jbmzb24v8704n"

func main() {
	fmt.Println("Booting Twitcher...")
	http.HandleFunc("/getLiveCreators", getLiveCreators)
	http.ListenAndServe(":8080", nil)
}

// `fetchAppAccessToken` fetches the app access token from Twitch.
func fetchAppAccessToken() string {
	url := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials", clientID, clientSecret)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Wew!")
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to parse the app access token!")
		return ""
	}

	appAuthToken := new(twitch.AppAuthTokenRes)
	json.Unmarshal(body, appAuthToken)

	fmt.Println("AT = ", appAuthToken.AccessToken)

	return appAuthToken.AccessToken
}

// getLiveCreators is a route handler for the route `/getLiveCreators`. It queries Twitch API and
// fetches the list of currently live content creators for the given search term.
//
// Note: I'm using TwitchAPIv5, although this is deprecated and will be removed on 12/31/18,
// the newer API still lacks the search capability.
func getLiveCreators(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": "Bad request"}`))
		return
	}

	params := r.URL.Query()                    // query the URL for GET request params
	searchString := params.Get("searchString") // the search string entered by the consumer
	if nil == &searchString || len(searchString) == 0 {
		log.Println("WARNING :: searchString is `nil`  or `empty` - setting it to `+` to match all streams!!")
		searchString = " " // space matches with all games in Twitch
	}
	pageNbr, err := strconv.Atoi(params.Get("pgNbr")) // the current page nbr, offset = pfNbr - 1
	if err != nil || pageNbr < 1 {
		log.Println("Unable to parse the desired page nbr!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Unable to parse the desired page nbr! Please check the page number entered."}`))
		return
	}

	appAccessToken := fetchAppAccessToken() // the app access token generated by Twitch

	log.Printf("HARMLESS :: access token = %s, searchString = %s", appAccessToken, searchString)

	apiURL := fmt.Sprintf("%s/search/streams?query=%s&hls=true&limit=%d&offset=%d", TWITCH_BASE_URL, url.QueryEscape(searchString), RES_LIMIT, (pageNbr-1)*RES_LIMIT) // Fetch list of active streams based on the search string
	log.Println("HARMLESS :: API URL generated = ", apiURL)
	req, _ := http.NewRequest("GET", apiURL, nil)

	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Couldn't get any response from Twitch!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Couldn't get any response from Twitch!"}`))
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	jsonResponse := make(map[string]interface{})
	json.Unmarshal(body, &jsonResponse)

	cards, err := twitch.BuildCards(&jsonResponse)
	if nil != err {
		log.Println("Couldn't build Twitch stream cards!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Couldn't build Twitch stream cards!"}`))
		return
	}

	log.Printf("HARMLESS :: cards = %s", string(cards))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cards)
}
