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
