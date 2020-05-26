package challonge

import (
	"fmt"
	"time"
)

// Use these structs for any PUT or POST
type Tournament struct {
	TournamentKey `json:"tournament"`
}

type TournamentKey struct {
	AcceptAttachments                bool      `json:"accept_attachments,omitempty"`
	AllowParticipantMatchReporting   bool      `json:"allow_participant_match_reporting,omitempty"`
	AnonymousVoting                  bool      `json:"anonymous_voting,omitempty"`
	Category                         string    `json:"category,omitempty"`
	CheckInDuration                  int64     `json:"check_in_duration,omitempty"`
	CompletedAt                      time.Time `json:"completed_at,omitempty"`
	CreatedAt                        time.Time `json:"created_at,omitempty"`
	CreatedByApi                     bool      `json:"created_by_api,omitempty"`
	CreditCapped                     bool      `json:"credit_capped,omitempty"`
	Description                      string    `json:"description,omitempty"`
	GameId                           int64     `json:"game_id,omitempty"`
	GroupStagesEnabled               bool      `json:"group_stages_enabled,omitempty"`
	HideForum                        bool      `json:"hide_forum,omitempty"`
	HideSeeds                        bool      `json:"hide_seeds,omitempty"`
	HoldThirdPlaceMatch              bool      `json:"hold_third_place_match,omitempty"`
	Id                               int64     `json:"id,omitempty"`
	MaxPredictionsPerUser            int       `json:"max_predictions_per_user,omitempty"`
	Name                             string    `json:"name,omitempty"`
	NotifyUsersWhenMatchesOpen       bool      `json:"notify_users_when_matches_open,omitempty"`
	NotifyUsersWhenTheTournamentEnds bool      `json:"notify_users_when_the_tournament_ends,omitempty"`
	OpenSignup                       bool      `json:"open_signup,omitempty"`
	ParticipantsCount                int       `json:"participants_count,omitempty"`
	PredictionMethod                 int       `json:"prediction_method,omitempty"`
	PredictionsOpenedAt              time.Time `json:"predictions_opened_at,omitempty"`
	Private                          bool      `json:"private,omitempty"`
	ProgressMeter                    int       `json:"progress_meter,omitempty"`
	PtsForBye                        string    `json:"pts_for_bye,omitempty"`
	PtsForGameTie                    string    `json:"pts_for_game_tie,omitempty"`
	PtsForGameWin                    string    `json:"pts_for_game_win,omitempty"`
	PtsForMatchTie                   string    `json:"pts_for_match_tie,omitempty"`
	PtsForMatchWin                   string    `json:"pts_for_match_win,omitempty"`
	QuickAdvance                     bool      `json:"quick_advance,omitempty"`
	RankedBy                         string    `json:"ranked_by,omitempty"`
	RequireScoreAgreement            bool      `json:"require_score_agreement,omitempty"`
	RrPtsForGameTie                  string    `json:"rr_pts_for_game_tie,omitempty"`
	RrPtsForGameWin                  string    `json:"rr_pts_for_game_win,omitempty"`
	RrPtsForMatchTie                 string    `json:"rr_pts_for_match_tie,omitempty"`
	RrPtsForMatchWin                 string    `json:"rr_pts_for_match_win,omitempty"`
	SequentialPairings               bool      `json:"sequential_pairings,omitempty"`
	ShowRounds                       bool      `json:"show_rounds,omitempty"`
	SignupCap                        int       `json:"signup_cap,omitempty"`
	StartAt                          time.Time `json:"start_at,omitempty"`
	StartedAt                        time.Time `json:"started_at,omitempty"`
	StartedCheckingInAt              time.Time `json:"started_checking_in_at,omitempty"`
	State                            string    `json:"state,omitempty"`
	SwissRounds                      int       `json:"swiss_rounds,omitempty"`
	Teams                            bool      `json:"teams,omitempty"`
	TieBreaks                        []string  `json:"tie_breaks,omitempty"`
	TournamentType                   string    `json:"tournament_type,omitempty"`
	UpdatedAt                        time.Time `json:"updated_at,omitempty"`
	Url                              string    `json:"url,omitempty"`
	DescriptionSource                string    `json:"description_source,omitempty"`
	Subdomain                        string    `json:"subdomain,omitempty"`
	FullChallongeUrl                 string    `json:"full_challonge_url,omitempty"`
	LiveImageUrl                     string    `json:"live_image_url,omitempty"`
	SignUpUrl                        string    `json:"sign_up_url,omitempty"`
	ReviewBeforeFinalizing           bool      `json:"review_before_finalizing,omitempty"`
	AcceptingPredictions             bool      `json:"accepting_predictions,omitempty"`
	ParticipantsLocked               bool      `json:"participants_locked,omitempty"`
	GameName                         string    `json:"game_name,omitempty"`
	ParticipantsSwappable            bool      `json:"participants_swappable,omitempty"`
	TeamConvertable                  bool      `json:"team_convertable,omitempty"`
	GroupStagesWereStarted           bool      `json:"group_stages_were_started,omitempty"`
	GrandFinalsModifier              string    `json:"grand_finals_modifier,omitempty"`
}

// Internal structs that are normalized for null fields
type InternalTournament struct {
	InternalTournamentKey `json:"tournament"`
}

type InternalTournamentKey struct {
	AcceptAttachments                ChallongeBool     `json:"accept_attachments,omitempty"`
	AllowParticipantMatchReporting   ChallongeBool     `json:"allow_participant_match_reporting,omitempty"`
	AnonymousVoting                  ChallongeBool     `json:"anonymous_voting,omitempty"`
	Category                         ChallongeString   `json:"category,omitempty"`
	CheckInDuration                  ChallongeInt64    `json:"check_in_duration,omitempty"`
	CompletedAt                      ChallongeTime     `json:"completed_at,omitempty"`
	CreatedAt                        ChallongeTime     `json:"created_at,omitempty"`
	CreatedByApi                     ChallongeBool     `json:"created_by_api,omitempty"`
	CreditCapped                     ChallongeBool     `json:"credit_capped,omitempty"`
	Description                      ChallongeString   `json:"description,omitempty"`
	GameId                           ChallongeInt64    `json:"game_id,omitempty"`
	GroupStagesEnabled               ChallongeBool     `json:"group_stages_enabled,omitempty"`
	HideForum                        ChallongeBool     `json:"hide_forum,omitempty"`
	HideSeeds                        ChallongeBool     `json:"hide_seeds,omitempty"`
	HoldThirdPlaceMatch              ChallongeBool     `json:"hold_third_place_match,omitempty"`
	Id                               ChallongeInt64    `json:"id,omitempty"`
	MaxPredictionsPerUser            ChallongeInt      `json:"max_predictions_per_user,omitempty"`
	Name                             ChallongeString   `json:"name,omitempty"`
	NotifyUsersWhenMatchesOpen       ChallongeBool     `json:"notify_users_when_matches_open,omitempty"`
	NotifyUsersWhenTheTournamentEnds ChallongeBool     `json:"notify_users_when_the_tournament_ends,omitempty"`
	OpenSignup                       ChallongeBool     `json:"open_signup,omitempty"`
	ParticipantsCount                ChallongeInt      `json:"participants_count,omitempty"`
	PredictionMethod                 ChallongeInt      `json:"prediction_method,omitempty"`
	PredictionsOpenedAt              ChallongeTime     `json:"predictions_opened_at,omitempty"`
	Private                          ChallongeBool     `json:"private,omitempty"`
	ProgressMeter                    ChallongeInt      `json:"progress_meter,omitempty"`
	PtsForBye                        ChallongeString   `json:"pts_for_bye,omitempty"`
	PtsForGameTie                    ChallongeString   `json:"pts_for_game_tie,omitempty"`
	PtsForGameWin                    ChallongeString   `json:"pts_for_game_win,omitempty"`
	PtsForMatchTie                   ChallongeString   `json:"pts_for_match_tie,omitempty"`
	PtsForMatchWin                   ChallongeString   `json:"pts_for_match_win,omitempty"`
	QuickAdvance                     ChallongeBool     `json:"quick_advance,omitempty"`
	RankedBy                         ChallongeString   `json:"ranked_by,omitempty"`
	RequireScoreAgreement            ChallongeBool     `json:"require_score_agreement,omitempty"`
	RrPtsForGameTie                  ChallongeString   `json:"rr_pts_for_game_tie,omitempty"`
	RrPtsForGameWin                  ChallongeString   `json:"rr_pts_for_game_win,omitempty"`
	RrPtsForMatchTie                 ChallongeString   `json:"rr_pts_for_match_tie,omitempty"`
	RrPtsForMatchWin                 ChallongeString   `json:"rr_pts_for_match_win,omitempty"`
	SequentialPairings               ChallongeBool     `json:"sequential_pairings,omitempty"`
	ShowRounds                       ChallongeBool     `json:"show_rounds,omitempty"`
	SignupCap                        ChallongeInt      `json:"signup_cap,omitempty"`
	StartAt                          ChallongeTime     `json:"start_at,omitempty"`
	StartedAt                        ChallongeTime     `json:"started_at,omitempty"`
	StartedCheckingInAt              ChallongeTime     `json:"started_checking_in_at,omitempty"`
	State                            ChallongeString   `json:"state,omitempty"`
	SwissRounds                      ChallongeInt      `json:"swiss_rounds,omitempty"`
	Teams                            ChallongeBool     `json:"teams,omitempty"`
	TieBreaks                        []ChallongeString `json:"tie_breaks,omitempty"`
	TournamentType                   ChallongeString   `json:"tournament_type,omitempty"`
	UpdatedAt                        ChallongeTime     `json:"updated_at,omitempty"`
	Url                              ChallongeString   `json:"url,omitempty"`
	DescriptionSource                ChallongeString   `json:"description_source,omitempty"`
	Subdomain                        ChallongeString   `json:"subdomain,omitempty"`
	FullChallongeUrl                 ChallongeString   `json:"full_challonge_url,omitempty"`
	LiveImageUrl                     ChallongeString   `json:"live_image_url,omitempty"`
	SignUpUrl                        ChallongeString   `json:"sign_up_url,omitempty"`
	ReviewBeforeFinalizing           ChallongeBool     `json:"review_before_finalizing,omitempty"`
	AcceptingPredictions             ChallongeBool     `json:"accepting_predictions,omitempty"`
	ParticipantsLocked               ChallongeBool     `json:"participants_locked,omitempty"`
	GameName                         ChallongeString   `json:"game_name,omitempty"`
	ParticipantsSwappable            ChallongeBool     `json:"participants_swappable,omitempty"`
	TeamConvertable                  ChallongeBool     `json:"team_convertable,omitempty"`
	GroupStagesWereStarted           ChallongeBool     `json:"group_stages_were_started,omitempty"`
	GrandFinalsModifier              ChallongeString   `json:"grand_finals_modifier,omitempty"`
}

type TournamentAction struct {
	IncludeParticipants int `json:"include_participants"`
	IncludeMatches      int `json:"include_matches"`
}

func (it *InternalTournament) normalize() *Tournament {
	return &Tournament{TournamentKey{
		AcceptAttachments:                bool(it.AcceptAttachments),
		AllowParticipantMatchReporting:   bool(it.AllowParticipantMatchReporting),
		AnonymousVoting:                  bool(it.AnonymousVoting),
		Category:                         string(it.Category),
		CheckInDuration:                  int64(it.CheckInDuration),
		CompletedAt:                      time.Time(it.CompletedAt),
		CreatedAt:                        time.Time(it.CreatedAt),
		CreatedByApi:                     bool(it.CreatedByApi),
		CreditCapped:                     bool(it.CreditCapped),
		Description:                      string(it.Description),
		GameId:                           int64(it.GameId),
		GroupStagesEnabled:               bool(it.GroupStagesEnabled),
		HideForum:                        bool(it.HideForum),
		HideSeeds:                        bool(it.HideSeeds),
		HoldThirdPlaceMatch:              bool(it.HoldThirdPlaceMatch),
		Id:                               int64(it.Id),
		MaxPredictionsPerUser:            int(it.MaxPredictionsPerUser),
		Name:                             string(it.Name),
		NotifyUsersWhenMatchesOpen:       bool(it.NotifyUsersWhenMatchesOpen),
		NotifyUsersWhenTheTournamentEnds: bool(it.NotifyUsersWhenTheTournamentEnds),
		OpenSignup:                       bool(it.OpenSignup),
		ParticipantsCount:                int(it.ParticipantsCount),
		PredictionMethod:                 int(it.PredictionMethod),
		PredictionsOpenedAt:              time.Time(it.PredictionsOpenedAt),
		Private:                          bool(it.Private),
		ProgressMeter:                    int(it.ProgressMeter),
		PtsForBye:                        string(it.PtsForBye),
		PtsForGameTie:                    string(it.PtsForGameTie),
		PtsForGameWin:                    string(it.PtsForGameWin),
		PtsForMatchTie:                   string(it.PtsForMatchTie),
		PtsForMatchWin:                   string(it.PtsForMatchWin),
		QuickAdvance:                     bool(it.QuickAdvance),
		RankedBy:                         string(it.RankedBy),
		RequireScoreAgreement:            bool(it.RequireScoreAgreement),
		RrPtsForGameTie:                  string(it.RrPtsForGameTie),
		RrPtsForGameWin:                  string(it.RrPtsForGameWin),
		RrPtsForMatchTie:                 string(it.RrPtsForMatchTie),
		RrPtsForMatchWin:                 string(it.RrPtsForMatchWin),
		SequentialPairings:               bool(it.SequentialPairings),
		ShowRounds:                       bool(it.ShowRounds),
		SignupCap:                        int(it.SignupCap),
		StartAt:                          time.Time(it.StartAt),
		StartedAt:                        time.Time(it.StartedAt),
		StartedCheckingInAt:              time.Time(it.StartedCheckingInAt),
		State:                            string(it.State),
		SwissRounds:                      int(it.SwissRounds),
		Teams:                            bool(it.Teams),
		TieBreaks:                        normalizeChallongeStringSlice(it.TieBreaks),
		TournamentType:                   string(it.TournamentType),
		UpdatedAt:                        time.Time(it.UpdatedAt),
		Url:                              string(it.Url),
		DescriptionSource:                string(it.DescriptionSource),
		Subdomain:                        string(it.Subdomain),
		FullChallongeUrl:                 string(it.FullChallongeUrl),
		LiveImageUrl:                     string(it.LiveImageUrl),
		SignUpUrl:                        string(it.SignUpUrl),
		ReviewBeforeFinalizing:           bool(it.ReviewBeforeFinalizing),
		AcceptingPredictions:             bool(it.AcceptingPredictions),
		ParticipantsLocked:               bool(it.ParticipantsLocked),
		GameName:                         string(it.GameName),
		ParticipantsSwappable:            bool(it.ParticipantsSwappable),
		TeamConvertable:                  bool(it.TeamConvertable),
		GroupStagesWereStarted:           bool(it.GroupStagesWereStarted),
		GrandFinalsModifier:              string(it.GrandFinalsModifier),
	}}
}

func (t *Tournament) denormalize() *InternalTournament {
	return &InternalTournament{InternalTournamentKey{
		AcceptAttachments:                ChallongeBool(t.AcceptAttachments),
		AllowParticipantMatchReporting:   ChallongeBool(t.AllowParticipantMatchReporting),
		AnonymousVoting:                  ChallongeBool(t.AnonymousVoting),
		Category:                         ChallongeString(t.Category),
		CheckInDuration:                  ChallongeInt64(t.CheckInDuration),
		CompletedAt:                      ChallongeTime(t.CompletedAt),
		CreatedAt:                        ChallongeTime(t.CreatedAt),
		CreatedByApi:                     ChallongeBool(t.CreatedByApi),
		CreditCapped:                     ChallongeBool(t.CreditCapped),
		Description:                      ChallongeString(t.Description),
		GameId:                           ChallongeInt64(t.GameId),
		GroupStagesEnabled:               ChallongeBool(t.GroupStagesEnabled),
		HideForum:                        ChallongeBool(t.HideForum),
		HideSeeds:                        ChallongeBool(t.HideSeeds),
		HoldThirdPlaceMatch:              ChallongeBool(t.HoldThirdPlaceMatch),
		Id:                               ChallongeInt64(t.Id),
		MaxPredictionsPerUser:            ChallongeInt(t.MaxPredictionsPerUser),
		Name:                             ChallongeString(t.Name),
		NotifyUsersWhenMatchesOpen:       ChallongeBool(t.NotifyUsersWhenMatchesOpen),
		NotifyUsersWhenTheTournamentEnds: ChallongeBool(t.NotifyUsersWhenTheTournamentEnds),
		OpenSignup:                       ChallongeBool(t.OpenSignup),
		ParticipantsCount:                ChallongeInt(t.ParticipantsCount),
		PredictionMethod:                 ChallongeInt(t.PredictionMethod),
		PredictionsOpenedAt:              ChallongeTime(t.PredictionsOpenedAt),
		Private:                          ChallongeBool(t.Private),
		ProgressMeter:                    ChallongeInt(t.ProgressMeter),
		PtsForBye:                        ChallongeString(t.PtsForBye),
		PtsForGameTie:                    ChallongeString(t.PtsForGameTie),
		PtsForGameWin:                    ChallongeString(t.PtsForGameWin),
		PtsForMatchTie:                   ChallongeString(t.PtsForMatchTie),
		PtsForMatchWin:                   ChallongeString(t.PtsForMatchWin),
		QuickAdvance:                     ChallongeBool(t.QuickAdvance),
		RankedBy:                         ChallongeString(t.RankedBy),
		RequireScoreAgreement:            ChallongeBool(t.RequireScoreAgreement),
		RrPtsForGameTie:                  ChallongeString(t.RrPtsForGameTie),
		RrPtsForGameWin:                  ChallongeString(t.RrPtsForGameWin),
		RrPtsForMatchTie:                 ChallongeString(t.RrPtsForMatchTie),
		RrPtsForMatchWin:                 ChallongeString(t.RrPtsForMatchWin),
		SequentialPairings:               ChallongeBool(t.SequentialPairings),
		ShowRounds:                       ChallongeBool(t.ShowRounds),
		SignupCap:                        ChallongeInt(t.SignupCap),
		StartAt:                          ChallongeTime(t.StartAt),
		StartedAt:                        ChallongeTime(t.StartedAt),
		StartedCheckingInAt:              ChallongeTime(t.StartedCheckingInAt),
		State:                            ChallongeString(t.State),
		SwissRounds:                      ChallongeInt(t.SwissRounds),
		Teams:                            ChallongeBool(t.Teams),
		TieBreaks:                        denormalizeStringSlice(t.TieBreaks),
		TournamentType:                   ChallongeString(t.TournamentType),
		UpdatedAt:                        ChallongeTime(t.UpdatedAt),
		Url:                              ChallongeString(t.Url),
		DescriptionSource:                ChallongeString(t.DescriptionSource),
		Subdomain:                        ChallongeString(t.Subdomain),
		FullChallongeUrl:                 ChallongeString(t.FullChallongeUrl),
		LiveImageUrl:                     ChallongeString(t.LiveImageUrl),
		SignUpUrl:                        ChallongeString(t.SignUpUrl),
		ReviewBeforeFinalizing:           ChallongeBool(t.ReviewBeforeFinalizing),
		AcceptingPredictions:             ChallongeBool(t.AcceptingPredictions),
		ParticipantsLocked:               ChallongeBool(t.ParticipantsLocked),
		GameName:                         ChallongeString(t.GameName),
		ParticipantsSwappable:            ChallongeBool(t.ParticipantsSwappable),
		TeamConvertable:                  ChallongeBool(t.TeamConvertable),
		GroupStagesWereStarted:           ChallongeBool(t.GroupStagesWereStarted),
		GrandFinalsModifier:              ChallongeString(t.GrandFinalsModifier),
	}}
}

func (c *ChallongeClient) executeTournamentPost(uri string, p, m bool) (*Tournament, error) {
	pi, mi := 0, 0
	if p {
		pi = 1
	}
	if m {
		mi = 1
	}
	o := &TournamentAction{IncludeParticipants: pi, IncludeMatches: mi}
	it := &InternalTournament{}
	err := c.Post(uri, o, it)
	if err != nil {
		return nil, err
	}
	return it.normalize(), nil
}

func (c *ChallongeClient) GetTournaments() ([]*Tournament, error) {
	uri := "tournaments"
	its := make([]InternalTournament, 0)
	ts := make([]*Tournament, 0)
	err := c.Get(uri, &its)
	if err != nil {
		return ts, err
	}
	for _, it := range its {
		ts = append(ts, it.normalize())
	}
	return ts, nil
}

func (c *ChallongeClient) GetTournament(id string) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s", id)
	t := InternalTournament{}
	err := c.Get(uri, &t)
	if err != nil {
		return nil, err
	}
	return t.normalize(), nil
}

func (c *ChallongeClient) CreateTournament(t *Tournament) (*Tournament, error) {
	uri := "tournaments"
	it := t.denormalize()
	r := &InternalTournament{}
	err := c.Post(uri, it, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) UpdateTournament(id string, t *Tournament) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s", id)
	it := t.denormalize()
	r := &InternalTournament{}
	err := c.Put(uri, it, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) DeleteTournament(id string) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s", id)
	it := &InternalTournament{}
	err := c.Delete(uri, it)
	if err != nil {
		return nil, err
	}
	return it.normalize(), nil
}

func (c *ChallongeClient) StartTournament(id string, p, m bool) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s/start", id)
	t, err := c.executeTournamentPost(uri, p, m)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *ChallongeClient) ProcessTournamentCheckIns(id string, p, m bool) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s/process_check_ins", id)
	t, err := c.executeTournamentPost(uri, p, m)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *ChallongeClient) AbortTournamentCheckIn(id string, p, m bool) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s/abort_check_in", id)
	t, err := c.executeTournamentPost(uri, p, m)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *ChallongeClient) FinalizeTournament(id string, p, m bool) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s/finalize", id)
	t, err := c.executeTournamentPost(uri, p, m)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *ChallongeClient) ResetTournament(id string, p, m bool) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s/reset", id)
	t, err := c.executeTournamentPost(uri, p, m)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *ChallongeClient) OpenTournamentForPredictions(id string, p, m bool) (*Tournament, error) {
	uri := fmt.Sprintf("tournaments/%s/open_for_predictions", id)
	t, err := c.executeTournamentPost(uri, p, m)
	if err != nil {
		return nil, err
	}
	return t, nil
}
