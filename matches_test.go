package challonge

import (
	"net/http"
	"testing"
	"time"
)

func TestGetMatches(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/matches/get_all_matches.json")
	server.StartTLS()
	defer server.Close()

	m, err := c.GetMatches("challonge_go_api_mock")
	if err != nil {
		t.Error(err)
	}
	if len(m) != 4 {
		t.Errorf("expected number of matches to be 4 but got %d", len(m))
	}
}

func TestGetMatch(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/matches/get_match.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 201537681
	m, err := c.GetMatch("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if !IsNull(m.CompletedAt) {
		t.Errorf("expected completed_at to be null but was instead %s", m.CompletedAt.Format(time.RFC3339))
	}
}

func TestUpdateMatch(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/matches/update_match.json")
	server.StartTLS()
	defer server.Close()

	mk := &Match{MatchKey{
		ScoresCsv: "3-1",
		WinnerId:  122371717,
	}}

	var id int64 = 201537681
	m, err := c.UpdateMatch("challonge_go_api_mock", id, mk)
	if err != nil {
		t.Error(err)
	}
	if m.ScoresCsv != "3-1" {
		t.Errorf("expected scores_csv to be '3-1' but was instead %s", m.ScoresCsv)
	}
	if IsNull(m.CompletedAt) {
		t.Errorf("expected completed_at to be non-null but was instead null")
	}
}

func TestReopenMatch(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/matches/reopen_match.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 201537681
	m, err := c.ReopenMatch("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if m.State != "open" {
		t.Errorf("expected state to be 'open' but was instead %s", m.State)
	}
	if !IsNull(m.CompletedAt) {
		t.Errorf("expected completed_at to be null but was instead %s", m.CompletedAt.Format(time.RFC3339))
	}
}

func TestMarkMatchAsUnderway(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/matches/mark_match_as_underway.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 201537681
	m, err := c.MarkMatchAsUnderway("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if IsNull(m.UnderwayAt) {
		t.Errorf("expected started_at to be non-null but was instead null")
	}
}

func TestUnmarkMatchAsUnderway(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/matches/unmark_match_as_underway.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 201537681
	m, err := c.UnmarkMatchAsUnderway("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if !IsNull(m.UnderwayAt) {
		t.Errorf("expected underway_at to be null but was instead %s", m.UnderwayAt.Format(time.RFC3339))
	}
}
