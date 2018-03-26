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
// api.go
// @author Sidharth Mishra
// @created Tue Aug 29 2017 17:42:32 GMT-0700 (PDT)
// @last-modified Sun Mar 25 2018 23:26:17 GMT-0700 (PDT)
//

package twitch

import (
	"encoding/json"
)

// AppAuthTokenRes is the wrapper for the incoming response from Twitch's auth API.
// It has the app's access token and its lifetime.
type AppAuthTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint   `json:"expires_in"`
}

// Card is a wrapper that exposes just enough information to build cards and view them on
// the mobile app.
type Card struct {
	ID           float64 `json:"_id"`          // the [stream]._id
	Game         string  `json:"game"`         // the [stream].game
	Viewers      float64 `json:"viewers"`      // the [stream].viewers
	StreamType   string  `json:"streamType"`   // the [stream].stream_type
	ThumbnailURI string  `json:"thumbnailURI"` // the [stream].[preview].large
	Title        string  `json:"title"`        // the [stream].[channel].status
	StreamerName string  `json:"streamerName"` // the [stream].[channel].display_name
	StreamURI    string  `json:"streamURI"`    // the [stream].[channel].url
}

// BuildCards builds the cards json to be sent to the mobile app.
func BuildCards(res *map[string]interface{}) ([]byte, error) {
	cards := make([]*Card, 0)

	for _, s := range ((*res)["streams"]).([]interface{}) {
		card := new(Card)
		stream := s.(map[string]interface{})
		card.ID = (stream["_id"]).(float64)
		card.Game = stream["game"].(string)
		card.Viewers = stream["viewers"].(float64)
		card.StreamType = stream["stream_type"].(string)
		card.ThumbnailURI = stream["preview"].(map[string]interface{})["large"].(string)
		card.Title = stream["channel"].(map[string]interface{})["status"].(string)
		card.StreamerName = stream["channel"].(map[string]interface{})["display_name"].(string)
		card.StreamURI = stream["channel"].(map[string]interface{})["url"].(string)
		cards = append(cards, card)
	}

	return json.Marshal(cards)
}
