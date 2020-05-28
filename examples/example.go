package main

import (
	"fmt"
	"os"
	"time"

	"github.com/claytonfinney/challonge-go"
)

func main() {
	c := challonge.NewChallongeClient(os.Getenv("CHALLONGE_USERNAME"), os.Getenv("CHALLONGE_API_KEY"))
	tme, _ := time.Parse(time.RFC3339, "2020-05-13T19:45:07.000Z")
	trn := challonge.Tournament{challonge.TournamentKey{
		Name:                "challonge go api mock",
		TournamentType:      "single elimination",
		Url:                 "challonge_go_api_mock",
		Description:         "testing a new api client for golang",
		OpenSignup:          false,
		HoldThirdPlaceMatch: true,
		Private:             false,
		StartAt:             tme,
		AcceptAttachments:   true,
	}}

	ct, err := c.CreateTournament(&trn)
	if err != nil {
		panic(err)
	}

	p := challonge.BulkParticipants{}
	p.Participants = append(p.Participants, challonge.BulkParticipant{Name: "Player1", Seed: 1, Misc: "Player1 is from Nairobi."})
	p.Participants = append(p.Participants, challonge.BulkParticipant{Name: "Player2", Seed: 2, Misc: "Player2 is from Delhi."})
	p.Participants = append(p.Participants, challonge.BulkParticipant{Name: "Player3", Seed: 3, Misc: "Player3 is from Montpellier."})
	p.Participants = append(p.Participants, challonge.BulkParticipant{Name: "Player4", Seed: 4, Misc: "Player4 is from Shanghai."})

	_, err = c.BulkAddParticipants(ct.Url, &p)
	if err != nil {
		panic(err)
	}

	_, err = c.StartTournament(ct.Url, false, false)
	if err != nil {
		panic(err)
	}

	ms, err := c.GetMatches(ct.Url)
	if err != nil {
		panic(err)
	}

	mid := ms[0].Id
	p1id := ms[0].Player1Id

	mk := challonge.Match{challonge.MatchKey{
		ScoresCsv: "3-1",
		WinnerId:  p1id,
	}}

	_, err = c.UpdateMatch(ct.Url, mid, &mk)
	if err != nil {
		panic(err)
	}

	ma := challonge.MatchAttachment{challonge.MatchAttachmentKey{
		Url:         "https://google.com",
		Description: "this is a match attachment!",
	}}

	_, err = c.CreateMatchAttachment(ct.Url, mid, &ma)
	if err != nil {
		panic(err)
	}

	_, err = c.DeleteTournament(ct.Url)
	if err != nil {
		panic(err)
	}

	if challonge.IsNull(ct.Url) {
		fmt.Println("url is null")
	}

	fmt.Println(ct.Name)
}
