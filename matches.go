package challonge

import (
	"fmt"
	"time"
)

// Use these structs for any PUT or POST
type Match struct {
	MatchKey `json:"match"`
}

type MatchKey struct {
	AttachmentCount           int         `json:"attachment_count,omitempty"`
	CompletedAt               time.Time   `json:"completed_at,omitempty"`
	CreatedAt                 time.Time   `json:"created_at,omitempty"`
	Forfeited                 bool        `json:"forfeited,omitempty"`
	GroupId                   int64       `json:"group_id,omitempty"`
	HasAttachment             bool        `json:"has_attachment,omitempty"`
	Id                        int64       `json:"id,omitempty"`
	Identifier                string      `json:"identifier,omitempty"`
	Location                  interface{} `json:"location,omitempty"`
	LoserId                   int64       `json:"loser_id,omitempty"`
	OgImageUrl                string      `json:"og_image_url,omitempty"`
	Optional                  bool        `json:"optional,omitempty"`
	Player1Id                 int64       `json:"player1_id,omitempty"`
	Player1IsPrereqMatchLoser bool        `json:"player1_is_prereq_match_loser,omitempty"`
	Player1PrereqMatchId      int64       `json:"player1_prereq_match_id,omitempty"`
	Player1Votes              int         `json:"player1_votes,omitempty"`
	Player2Id                 int64       `json:"player2_id,omitempty"`
	Player2IsPrereqMatchLoser bool        `json:"player2_is_prereq_match_loser,omitempty"`
	Player2PrereqMatchId      int64       `json:"player2_prereq_match_id,omitempty"`
	Player2Votes              int         `json:"player2_votes,omitempty"`
	Round                     int         `json:"round,omitempty"`
	RushBId                   int64       `json:"rushb_id,omitempty"`
	ScheduledTime             time.Time   `json:"scheduled_time,omitempty"`
	StartedAt                 time.Time   `json:"started_at,omitempty"`
	State                     string      `json:"state,omitempty"`
	TournamentId              int64       `json:"tournament_id,omitempty"`
	UnderwayAt                time.Time   `json:"underway_at,omitempty"`
	UpdatedAt                 time.Time   `json:"updated_at,omitempty"`
	WinnerId                  int64       `json:"winner_id,omitempty"`
	PrerequisiteMatchIdsCsv   string      `json:"prerequisite_match_ids_csv,omitempty"`
	ScoresCsv                 string      `json:"scores_csv,omitempty"`
}

// Internal structs that are normalized for null fields
type InternalMatch struct {
	InternalMatchKey `json:"match"`
}

type InternalMatchKey struct {
	AttachmentCount           ChallongeInt    `json:"attachment_count,omitempty"`
	CompletedAt               ChallongeTime   `json:"completed_at,omitempty"`
	CreatedAt                 ChallongeTime   `json:"created_at,omitempty"`
	Forfeited                 ChallongeBool   `json:"forfeited,omitempty"`
	GroupId                   ChallongeInt64  `json:"group_id,omitempty"`
	HasAttachment             ChallongeBool   `json:"has_attachment,omitempty"`
	Id                        ChallongeInt64  `json:"id,omitempty"`
	Identifier                ChallongeString `json:"identifier,omitempty"`
	Location                  interface{}     `json:"location,omitempty"`
	LoserId                   ChallongeInt64  `json:"loser_id,omitempty"`
	OgImageUrl                ChallongeString `json:"og_image_url,omitempty"`
	Optional                  ChallongeBool   `json:"optional,omitempty"`
	Player1Id                 ChallongeInt64  `json:"player1_id,omitempty"`
	Player1IsPrereqMatchLoser ChallongeBool   `json:"player1_is_prereq_match_loser,omitempty"`
	Player1PrereqMatchId      ChallongeInt64  `json:"player1_prereq_match_id,omitempty"`
	Player1Votes              ChallongeInt    `json:"player1_votes,omitempty"`
	Player2Id                 ChallongeInt64  `json:"player2_id,omitempty"`
	Player2IsPrereqMatchLoser ChallongeBool   `json:"player2_is_prereq_match_loser,omitempty"`
	Player2PrereqMatchId      ChallongeInt64  `json:"player2_prereq_match_id,omitempty"`
	Player2Votes              ChallongeInt    `json:"player2_votes,omitempty"`
	Round                     ChallongeInt    `json:"round,omitempty"`
	RushBId                   ChallongeInt64  `json:"rushb_id,omitempty"`
	ScheduledTime             ChallongeTime   `json:"scheduled_time,omitempty"`
	StartedAt                 ChallongeTime   `json:"started_at,omitempty"`
	State                     ChallongeString `json:"state,omitempty"`
	TournamentId              ChallongeInt64  `json:"tournament_id,omitempty"`
	UnderwayAt                ChallongeTime   `json:"underway_at,omitempty"`
	UpdatedAt                 ChallongeTime   `json:"updated_at,omitempty"`
	WinnerId                  ChallongeInt64  `json:"winner_id,omitempty"`
	PrerequisiteMatchIdsCsv   ChallongeString `json:"prerequisite_match_ids_csv,omitempty"`
	ScoresCsv                 ChallongeString `json:"scores_csv,omitempty"`
}

func (im *InternalMatch) normalize() *Match {
	return &Match{MatchKey{
		AttachmentCount:           int(im.AttachmentCount),
		CompletedAt:               time.Time(im.CompletedAt),
		CreatedAt:                 time.Time(im.CreatedAt),
		Forfeited:                 bool(im.Forfeited),
		GroupId:                   int64(im.GroupId),
		HasAttachment:             bool(im.HasAttachment),
		Id:                        int64(im.Id),
		Identifier:                string(im.Identifier),
		Location:                  im.Location,
		LoserId:                   int64(im.LoserId),
		OgImageUrl:                string(im.OgImageUrl),
		Optional:                  bool(im.Optional),
		Player1Id:                 int64(im.Player1Id),
		Player1IsPrereqMatchLoser: bool(im.Player1IsPrereqMatchLoser),
		Player1PrereqMatchId:      int64(im.Player1PrereqMatchId),
		Player1Votes:              int(im.Player1Votes),
		Player2Id:                 int64(im.Player2Id),
		Player2IsPrereqMatchLoser: bool(im.Player2IsPrereqMatchLoser),
		Player2PrereqMatchId:      int64(im.Player2PrereqMatchId),
		Player2Votes:              int(im.Player2Votes),
		Round:                     int(im.Round),
		RushBId:                   int64(im.RushBId),
		ScheduledTime:             time.Time(im.ScheduledTime),
		StartedAt:                 time.Time(im.StartedAt),
		State:                     string(im.State),
		TournamentId:              int64(im.TournamentId),
		UnderwayAt:                time.Time(im.UnderwayAt),
		UpdatedAt:                 time.Time(im.UpdatedAt),
		WinnerId:                  int64(im.WinnerId),
		PrerequisiteMatchIdsCsv:   string(im.PrerequisiteMatchIdsCsv),
		ScoresCsv:                 string(im.ScoresCsv),
	}}
}

func (m *Match) denormalize() *InternalMatch {
	return &InternalMatch{InternalMatchKey{
		AttachmentCount:           ChallongeInt(m.AttachmentCount),
		CompletedAt:               ChallongeTime(m.CompletedAt),
		CreatedAt:                 ChallongeTime(m.CreatedAt),
		Forfeited:                 ChallongeBool(m.Forfeited),
		GroupId:                   ChallongeInt64(m.GroupId),
		HasAttachment:             ChallongeBool(m.HasAttachment),
		Id:                        ChallongeInt64(m.Id),
		Identifier:                ChallongeString(m.Identifier),
		Location:                  m.Location,
		LoserId:                   ChallongeInt64(m.LoserId),
		OgImageUrl:                ChallongeString(m.OgImageUrl),
		Optional:                  ChallongeBool(m.Optional),
		Player1Id:                 ChallongeInt64(m.Player1Id),
		Player1IsPrereqMatchLoser: ChallongeBool(m.Player1IsPrereqMatchLoser),
		Player1PrereqMatchId:      ChallongeInt64(m.Player1PrereqMatchId),
		Player1Votes:              ChallongeInt(m.Player1Votes),
		Player2Id:                 ChallongeInt64(m.Player2Id),
		Player2IsPrereqMatchLoser: ChallongeBool(m.Player2IsPrereqMatchLoser),
		Player2PrereqMatchId:      ChallongeInt64(m.Player2PrereqMatchId),
		Player2Votes:              ChallongeInt(m.Player2Votes),
		Round:                     ChallongeInt(m.Round),
		RushBId:                   ChallongeInt64(m.RushBId),
		ScheduledTime:             ChallongeTime(m.ScheduledTime),
		StartedAt:                 ChallongeTime(m.ScheduledTime),
		State:                     ChallongeString(m.State),
		TournamentId:              ChallongeInt64(m.TournamentId),
		UnderwayAt:                ChallongeTime(m.UnderwayAt),
		UpdatedAt:                 ChallongeTime(m.UpdatedAt),
		WinnerId:                  ChallongeInt64(m.WinnerId),
		PrerequisiteMatchIdsCsv:   ChallongeString(m.PrerequisiteMatchIdsCsv),
		ScoresCsv:                 ChallongeString(m.ScoresCsv),
	}}
}

func (c *ChallongeClient) GetMatches(id string) ([]*Match, error) {
	uri := fmt.Sprintf("tournaments/%s/matches", id)
	ims := make([]InternalMatch, 0)
	ms := make([]*Match, 0)
	err := c.Get(uri, &ims)
	if err != nil {
		return ms, err
	}
	for _, im := range ims {
		ms = append(ms, im.normalize())
	}
	return ms, nil
}

func (c *ChallongeClient) GetMatch(tid string, mid int64) (*Match, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d", tid, mid)
	im := &InternalMatch{}
	err := c.Get(uri, &im)
	if err != nil {
		return nil, err
	}
	return im.normalize(), nil
}

func (c *ChallongeClient) UpdateMatch(tid string, mid int64, m *Match) (*Match, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d", tid, mid)
	im := m.denormalize()
	r := &InternalMatch{}
	err := c.Put(uri, im, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) executeMatchPost(uri string) (*Match, error) {
	im := &InternalMatch{}
	err := c.Post(uri, nil, im)
	if err != nil {
		return nil, err
	}
	return im.normalize(), nil
}

func (c *ChallongeClient) ReopenMatch(tid string, mid int64) (*Match, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/reopen", tid, mid)
	m, err := c.executeMatchPost(uri)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ChallongeClient) MarkMatchAsUnderway(tid string, mid int64) (*Match, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/mark_as_underway", tid, mid)
	m, err := c.executeMatchPost(uri)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ChallongeClient) UnmarkMatchAsUnderway(tid string, mid int64) (*Match, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/unmark_as_underway", tid, mid)
	m, err := c.executeMatchPost(uri)
	if err != nil {
		return nil, err
	}
	return m, nil
}
