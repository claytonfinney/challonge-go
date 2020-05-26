package challonge

import (
	"net/http"
	"testing"
	"time"
)

func TestGetTournaments(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/get_all_tournaments.json")
	server.StartTLS()
	defer server.Close()

	ts, err := c.GetTournaments()
	if err != nil {
		t.Error(err)
	}
	idx := len(ts) - 1
	if ts[idx].Name != "challonge go api mock" {
		t.Errorf(`expected name to be "challonge go api mock", got %s`, ts[idx].Name)
	}
}

func TestCreateTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/create_tournament.json")
	server.StartTLS()
	defer server.Close()

	tme, _ := time.Parse(time.RFC3339, "2020-05-19T17:42:07.165+08:00")
	trn := &Tournament{TournamentKey{
		Name:                "challonge go api mock",
		TournamentType:      "single elimination",
		Url:                 "challonge_go_api_mock",
		Description:         "testing a new api client for golang",
		OpenSignup:          false,
		HoldThirdPlaceMatch: true,
		Private:             false,
		StartAt:             tme,
		PredictionMethod:    2,
	}}

	ct, err := c.CreateTournament(trn)
	if err != nil {
		t.Error(err)
	}
	if ct.Name != "challonge go api mock" {
		t.Error()
	}
}

func TestGetTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/get_tournament.json")
	server.StartTLS()
	defer server.Close()

	ts, err := c.GetTournament("challonge_go_api_mock")
	if err != nil {
		t.Error(err)
	}
	if ts.Name != "challonge go api mock" {
		t.Errorf(`expected name to be "challonge go api mock", got %s`, ts.Name)
	}
	if ts.GameId != NULL_INT64 {
		t.Errorf(`key game_id was expected to be null, got %v`, ts.GameId)
	}
}

func TestUpdateTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/update_tournament.json")
	server.StartTLS()
	defer server.Close()

	trn := &Tournament{TournamentKey{
		Description:       "changed the description as part of the test",
		OpenSignup:        false,
		RankedBy:          "points_scored",
		AcceptAttachments: true,
	}}

	ct, err := c.UpdateTournament("challonge_go_api_mock", trn)
	if err != nil {
		t.Error(err)
	}
	if ct.AcceptAttachments != true {
		t.Errorf("expected accept_attachments to be %t but got %t", true, ct.AcceptAttachments)
	}
}

func TestDeleteTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/get_tournament.json")
	server.StartTLS()
	defer server.Close()

	ct, err := c.DeleteTournament("challonge_go_api_mock")
	if err != nil {
		t.Error(err)
	}
	if ct.GroupStagesWereStarted != false {
		t.Errorf("expected group_stages_were_started to be %t but got %t", true, ct.GroupStagesWereStarted)
	}
}

func TestStartTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/start_tournament.json")
	server.StartTLS()
	defer server.Close()

	ct, err := c.StartTournament("challonge_go_api_mock", false, false)
	if err != nil {
		t.Error(nil)
	}
	if ct.StartedAt.Format(time.RFC3339) == NULL_TIME {
		t.Errorf("expected started_at to be a non-null time")
	}
	if ct.State != "underway" {
		t.Errorf("expected tournament state to be 'underway', but got %s", ct.State)
	}
}

func TestFinalizeTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/finalize_tournament.json")
	server.StartTLS()
	defer server.Close()

	ct, err := c.StartTournament("challonge_go_api_mock", false, false)
	if err != nil {
		t.Error(nil)
	}
	if ct.CompletedAt.Format(time.RFC3339) == NULL_TIME {
		t.Errorf("expected completed_at to be a non-null time")
	}
	if ct.State != "complete" {
		t.Errorf("expected tournament state to be 'complete', but got %s", ct.State)
	}
}

func TestResetTournament(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/tournaments/reset_tournament.json")
	server.StartTLS()
	defer server.Close()

	ct, err := c.ResetTournament("challonge_go_api_mock", false, false)
	if err != nil {
		t.Error()
	}
	if ct.State != "pending" {
		t.Errorf("expected tournament state to be 'pending', but got %s", ct.State)
	}
	if ct.StartedAt.Format(time.RFC3339) != NULL_TIME {
		t.Errorf("expected started_at to be null, but instead was %s", ct.StartedAt.Format(time.RFC3339))
	}
}
