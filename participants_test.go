package challonge

import (
	"net/http"
	"testing"
	"time"
)

func TestGetParticipants(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/get_all_participants.json")
	server.StartTLS()
	defer server.Close()

	ps, err := c.GetParticipants("challonge_go_api_mock")
	if err != nil {
		t.Error(err)
	}
	if !IsNull(ps[0].ChallongeUsername) {
		t.Errorf("expected challonge username to be null but instead got %s", ps[0].ChallongeUsername)
	}
	if len(ps) != 4 {
		t.Errorf("expected number of players to be 4 but got %d", len(ps))
	}
}

func TestCreateParticipant(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/create_participant.json")
	server.StartTLS()
	defer server.Close()

	p := &Participant{ParticipantKey{
		Name: "Player0",
		Seed: 1,
		Misc: "Player0 is from Seoul.",
	}}

	ptc, err := c.CreateParticipant("challonge_go_api_mock", p)
	if err != nil {
		t.Error(err)
	}
	if ptc.Seed != 1 {
		t.Errorf("expected seed to be 1 but got %d", ptc.Seed)
	}
	if ptc.Id != 122370385 {
		t.Errorf("expected id to be 122370385 but got %d", ptc.Id)
	}
}

func TestGetParticipant(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/get_participant.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 122370385
	p, err := c.GetParticipant("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if p.Name != "Player0" {
		t.Errorf("expected name to be Player0 but got %s", p.Name)
	}
}

func TestDeleteParticipant(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/delete_participant.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 122370385
	ptc, err := c.DeleteParticipant("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if ptc.Name != "Player0" {
		t.Errorf("expected name to be Player0 but was %s", ptc.Name)
	}
}

func TestBulkAddParticipants(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/bulk_add_participants.json")
	server.StartTLS()
	defer server.Close()

	p := make([]*Participant, 0)
	p = append(p, &Participant{ParticipantKey{Name: "Player1", Seed: 1, Misc: "Player1 is from Nairobi."}})
	p = append(p, &Participant{ParticipantKey{Name: "Player2", Seed: 2, Misc: "Player2 is from Delhi."}})
	p = append(p, &Participant{ParticipantKey{Name: "Player3", Seed: 3, Misc: "Player3 is from Montpellier."}})
	p = append(p, &Participant{ParticipantKey{Name: "Player4", Seed: 4, Misc: "Player4 is from Shanghai."}})

	ptc, err := c.BulkAddParticipants("challonge_go_api_mock", p)
	if err != nil {
		t.Error(err)
	}
	if len(ptc) != 4 {
		t.Errorf("expected number of players to be 4 but got %d", len(ptc))
	}
	if ptc[0].Name != "Player1" {
		t.Errorf("expected name to be PLayer1 but got %s", ptc[0].Name)
	}
}

func TestUpdateParticipant(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/update_participant.json")
	server.StartTLS()
	defer server.Close()

	p := &Participant{ParticipantKey{
		Misc: "Player4 is actually from Harbing.",
	}}

	var id int64 = 122371718
	ptc, err := c.UpdateParticipant("challonge_go_api_mock", id, p)
	if err != nil {
		t.Error(err)
	}
	if ptc.Misc != "Player4 is actually from Harbing, and should be seeded first." {
		t.Errorf("Player1's updated misc did not match")
	}
}

func TestRandomizeParticipants(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/randomize_participants.json")
	server.StartTLS()
	defer server.Close()

	ps, err := c.RandomizeParticipants("challonge_go_api_mock")
	if err != nil {
		t.Error(err)
	}
	if ps[0].Name != "Player3" {
		t.Errorf("expected name to be Player3 but was %s", ps[0].Name)
	}
	if ps[0].Seed != 1 {
		t.Errorf("expected seed to be 1 but got %d", ps[0].Seed)
	}
}

func TestCheckInParticipant(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/check_in_participant.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 122296099
	p, err := c.CheckInParticipant("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if p.CheckedIn == false {
		t.Errorf("expected checked_in to be true but was false")
	}
}

func TestUndoCheckInParticipant(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/undo_check_in_participant.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 122296099
	p, err := c.UndoCheckInParticipant("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if p.CheckedIn == true {
		t.Errorf("expected checked_in to be true but was false")
	}
	if !IsNull(p.CheckedInAt) {
		t.Errorf("expected checked_in_at to be null but instead was %s", p.CheckedInAt.Format(time.RFC3339))
	}
}

func TestClearParticipants(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/participants/clear_participants.json")
	server.StartTLS()
	defer server.Close()

	err := c.ClearParticipants("challonge_go_api_mock")
	if err != nil {
		t.Error(err)
	}
}
