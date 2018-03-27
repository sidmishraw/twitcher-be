# Twitcher

This is the RESTful backend of `Twitcher`.

<!-- Client ID: 64co18zsfvj6plbqy5v2j7x5r6auza -->

<!-- Client Secret: savr6m36r8vh4fsc9jbmzb24v8704n -->

## Building instructions for Docker

This project is docker ready.

> Note: Tested on Docker 17.12.0-ce-mac55 (23011)

Default port this application listens on is `8080`.

Steps to build:

1 . Clone the git repo onto your machine.

```shell
git clone https://github.com/sidmishraw/twitcher-be.git
```

2 . Build the docker image for twitcher while being at twitcher's project root.

```shell
docker build . -f ./_ops/Dockerfile -t sidmishraw/twitcher-be:v0.0.1
```

3 . Run the docker image for twitcher.

```shell
docker run -it --rm -p 8080:8080 sidmishraw/twitcher-be:v0.0.1
```

## Building instructions for non Docker

1 . Clone the git repo onto your machine.

```shell
git clone https://github.com/sidmishraw/twitcher-be.git
```

2 . Build the app using `go build`.

```shell
go build -o twitcher -i
```

3 . Run the executable named `twitcher`.

```shell
./twitcher
```

## API endpoint(s)

* `getLiveCreators`: Queries Twitch v5 API and fetches the list of currently live content creators for the given search term.
  The default number of results per page is 25.

  > Note: I'm using TwitchAPIv5, although this is deprecated and will be removed on 12/31/18, the newer API still lacks the search capability.

  This API endpoint takes the following params:

  * `searchString`: The search term entered by the user. This is a `string` and it is URL encoded internally.
  * `pgNbr`: The page number of the results to fetch from Twitch. It is a positive integer `>= 1`.

Example API invocation: `http://localhost:8080/getLiveCreators?searchString=dota2&pgNbr=3`
Example response:

```json
[
  {
    "_id": 28089597632,
    "game": "Dota 2",
    "viewers": 14,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_maxyf4-640x360.jpg",
    "title": "MaxyF4 - subiendo puntitos",
    "streamerName": "maxyf4",
    "streamURI": "https://www.twitch.tv/maxyf4"
  },
  {
    "_id": 28087349264,
    "game": "Dota 2",
    "viewers": 13,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_seipk-640x360.jpg",
    "title": "i may or may not have reinstalled",
    "streamerName": "seipk",
    "streamURI": "https://www.twitch.tv/seipk"
  },
  {
    "_id": 28086963264,
    "game": "Dota 2",
    "viewers": 11,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_momoisan1-640x360.jpg",
    "title": "Mode trapito, AMANECIDA",
    "streamerName": "momoisan1",
    "streamURI": "https://www.twitch.tv/momoisan1"
  },
  {
    "_id": 28088285632,
    "game": "Dota 2",
    "viewers": 11,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_redditdota2league-640x360.jpg",
    "title": "RD2L Playoffs - Mr Matieu vs Ripley - Casted by Pablo",
    "streamerName": "redditdota2league",
    "streamURI": "https://www.twitch.tv/redditdota2league"
  },
  {
    "_id": 28089233728,
    "game": "Dota 2",
    "viewers": 10,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_ddgodd-640x360.jpg",
    "title": " dD ツ  RANK 4 DOTABUUFF100% GOODMANNER",
    "streamerName": "dDGodd",
    "streamURI": "https://www.twitch.tv/ddgodd"
  },
  {
    "_id": 28088377200,
    "game": "Dota 2",
    "viewers": 10,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_badgerpablo-640x360.jpg",
    "title": "RD2L Playoffs Matieu v Ripley",
    "streamerName": "BadgerPablo",
    "streamURI": "https://www.twitch.tv/badgerpablo"
  },
  {
    "_id": 28087112768,
    "game": "Dota 2",
    "viewers": 9,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_risenagain34-640x360.jpg",
    "title": "[ENG] Drunk DOTA with friends",
    "streamerName": "Risenagain34",
    "streamURI": "https://www.twitch.tv/risenagain34"
  },
  {
    "_id": 28087021472,
    "game": "Dota 2",
    "viewers": 9,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_weedmandota-640x360.jpg",
    "title": "Dota With Weedman",
    "streamerName": "weedmandota",
    "streamURI": "https://www.twitch.tv/weedmandota"
  },
  {
    "_id": 28088148576,
    "game": "Dota 2",
    "viewers": 9,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_ravenswings-640x360.jpg",
    "title": "Green hero strim?",
    "streamerName": "Ravenswings",
    "streamURI": "https://www.twitch.tv/ravenswings"
  },
  {
    "_id": 28085942928,
    "game": "Dota 2",
    "viewers": 9,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_geimer95151-640x360.jpg",
    "title": "Восприятие Dota 2 с позитивной стороны",
    "streamerName": "Geimer95151",
    "streamURI": "https://www.twitch.tv/geimer95151"
  },
  {
    "_id": 28089109200,
    "game": "Dota 2",
    "viewers": 9,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_neverseemyface-640x360.jpg",
    "title": ":v",
    "streamerName": "neverseemyface",
    "streamURI": "https://www.twitch.tv/neverseemyface"
  },
  {
    "_id": 28085601344,
    "game": "Dota 2",
    "viewers": 9,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_theekstranghero-640x360.jpg",
    "title": "BAGONG CHUPIT GAMING [FIL/ENG]",
    "streamerName": "TheEkstrangHero",
    "streamURI": "https://www.twitch.tv/theekstranghero"
  },
  {
    "_id": 28087850016,
    "game": "dota 2",
    "viewers": 8,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_sakii89-640x360.jpg",
    "title": "(ENG/ESP/ARG) ~ adivinen quien sigue enferma :3",
    "streamerName": "Sakii89",
    "streamURI": "https://www.twitch.tv/sakii89"
  },
  {
    "_id": 28088397248,
    "game": "Dota 2",
    "viewers": 8,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_htlannie-640x360.jpg",
    "title": "Aggro af",
    "streamerName": "htlannie",
    "streamURI": "https://www.twitch.tv/htlannie"
  },
  {
    "_id": 28088378208,
    "game": "Dota 2",
    "viewers": 8,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_miredia-640x360.jpg",
    "title": "[ENG/US:E 1080p/60fps] Budget Miredia - Dota 2 Solo matches",
    "streamerName": "Miredia",
    "streamURI": "https://www.twitch.tv/miredia"
  },
  {
    "_id": 28087136560,
    "game": "Dota 2",
    "viewers": 8,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_tinirso-640x360.jpg",
    "title":
      "[PT-BR] O follow é só pros humildes!! Rumo aos 200 follows, chega na live!!",
    "streamerName": "tinirso",
    "streamURI": "https://www.twitch.tv/tinirso"
  },
  {
    "_id": 28084832192,
    "game": "Dota 2",
    "viewers": 8,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_pnpsangmin-640x360.jpg",
    "title": "[KR/EN] Actor ShankS 상민 도타 사실월클 DOTA2 World CLASS",
    "streamerName": "쉥크스",
    "streamURI": "https://www.twitch.tv/pnpsangmin"
  },
  {
    "_id": 28089961056,
    "game": "Dota 2",
    "viewers": 8,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_chilling-640x360.jpg",
    "title": "Divine 5 Support (5) SOLO RANKED / chilling / amanecida",
    "streamerName": "Chilling",
    "streamURI": "https://www.twitch.tv/chilling"
  },
  {
    "_id": 28083273664,
    "game": "Dota 2",
    "viewers": 7,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_lunaticim-640x360.jpg",
    "title": "Twitch Kompozisyonu'da Nedir?",
    "streamerName": "lunaticim",
    "streamURI": "https://www.twitch.tv/lunaticim"
  },
  {
    "_id": 28089297984,
    "game": "Dota 2",
    "viewers": 7,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_chicoxxx-640x360.jpg",
    "title": "일본서버 지박령",
    "streamerName": "치코",
    "streamURI": "https://www.twitch.tv/chicoxxx"
  },
  {
    "_id": 28089044288,
    "game": "Dota 2",
    "viewers": 7,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_saharnizza-640x360.jpg",
    "title": "Ламповая ночь *)",
    "streamerName": "Saharnizza",
    "streamURI": "https://www.twitch.tv/saharnizza"
  },
  {
    "_id": 28086327472,
    "game": "Dota 2",
    "viewers": 7,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_blackimustv-640x360.jpg",
    "title": "{Divine 4} Ranked Queue. ( No cam today )",
    "streamerName": "BlackimusTV",
    "streamURI": "https://www.twitch.tv/blackimustv"
  },
  {
    "_id": 28089952064,
    "game": "Dota 2",
    "viewers": 6,
    "streamType": "live",
    "thumbnailURI":
      "https://static-cdn.jtvnw.net/previews-ttv/live_user_immunepwnz-640x360.jpg",
    "title": "privetuli",
    "streamerName": "immunepwnz",
    "streamURI": "https://www.twitch.tv/immunepwnz"
  }
]
```
