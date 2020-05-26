package challonge

import (
	"fmt"
	"time"
)

// Use these structs for any PUT or POST
type Participant struct {
	ParticipantKey `json:"participant"`
}

type ParticipantKey struct {
	Active                                bool      `json:"active,omitempty"`
	CheckedInAt                           time.Time `json:"checked_in_at,omitempty"`
	CreatedAt                             time.Time `json:"created_at,omitempty"`
	FinalRank                             int       `json:"final_rank,omitempty"`
	GroupId                               int64     `json:"group_id,omitempty"`
	Icon                                  string    `json:"icon,omitempty"`
	Id                                    int64     `json:"id,omitempty"`
	InvitationId                          int64     `json:"invitation_id,omitempty"`
	InviteEmail                           string    `json:"invite_email,omitempty"`
	Misc                                  string    `json:"misc,omitempty"`
	Name                                  string    `json:"name,omitempty"`
	OnWaitingList                         bool      `json:"on_waiting_list,omitempty"`
	Seed                                  int       `json:"seed,omitempty"`
	ParticipantId                         int       `json:"tournament_id,omitempty"`
	UpdatedAt                             time.Time `json:"updated_at,omitempty"`
	ChallongeUsername                     string    `json:"challonge_username,omitempty"`
	ChallongeEmailAddressVerified         bool      `json:"challonge_email_address_verified,omitempty"`
	Removable                             bool      `json:"removable,omitempty"`
	ParticipatableOrInvitationAttached    bool      `json:"participatable_or_invitation_attached,omitempty"`
	ConfirmRemove                         bool      `json:"confirm_remove,omitempty"`
	InvitationPending                     bool      `json:"invitation_pending,omitempty"`
	DisplayNameWithInvitationEmailAddress string    `json:"display_name_with_invitation_email_address,omitempty"`
	EmailHash                             string    `json:"email_hash,omitempty"`
	Username                              string    `json:"username,omitempty"`
	AttachedParticipatablePortraitURL     string    `json:"attached_participatable_portrait_url,omitempty"`
	CanCheckIn                            bool      `json:"can_check_in,omitempty"`
	CheckedIn                             bool      `json:"checked_in,omitempty"`
	Reactivatable                         bool      `json:"reactivatable,omitempty"`
}

type BulkParticipants struct {
	Participants []BulkParticipant `json:"participants"`
}

type BulkParticipant struct {
	Name              string `json:"name,omitempty"`
	Seed              int    `json:"seed,omitempty"`
	Misc              string `json:"misc,omitempty"`
	InviteNameOrEmail string `json:"invite_name_or_email,omitempty"`
}

// Internal structs that are normalized for null fields
type InternalParticipant struct {
	InternalParticipantKey `json:"participant"`
}

type InternalParticipantKey struct {
	Active                                ChallongeBool   `json:"active,omitempty"`
	CheckedInAt                           ChallongeTime   `json:"checked_in_at,omitempty"`
	CreatedAt                             ChallongeTime   `json:"created_at,omitempty"`
	FinalRank                             ChallongeInt    `json:"final_rank,omitempty"`
	GroupId                               ChallongeInt64  `json:"group_id,omitempty"`
	Icon                                  ChallongeString `json:"icon,omitempty"`
	Id                                    ChallongeInt64  `json:"id,omitempty"`
	InvitationId                          ChallongeInt64  `json:"invitation_id,omitempty"`
	InviteEmail                           ChallongeString `json:"invite_email,omitempty"`
	Misc                                  ChallongeString `json:"misc,omitempty"`
	Name                                  ChallongeString `json:"name,omitempty"`
	OnWaitingList                         ChallongeBool   `json:"on_waiting_list,omitempty"`
	Seed                                  ChallongeInt    `json:"seed,omitempty"`
	ParticipantId                         ChallongeInt    `json:"tournament_id,omitempty"`
	UpdatedAt                             ChallongeTime   `json:"updated_at,omitempty"`
	ChallongeUsername                     ChallongeString `json:"challonge_username,omitempty"`
	ChallongeEmailAddressVerified         ChallongeBool   `json:"challonge_email_address_verified,omitempty"`
	Removable                             ChallongeBool   `json:"removable,omitempty"`
	ParticipatableOrInvitationAttached    ChallongeBool   `json:"participatable_or_invitation_attached,omitempty"`
	ConfirmRemove                         ChallongeBool   `json:"confirm_remove,omitempty"`
	InvitationPending                     ChallongeBool   `json:"invitation_pending,omitempty"`
	DisplayNameWithInvitationEmailAddress ChallongeString `json:"display_name_with_invitation_email_address,omitempty"`
	EmailHash                             ChallongeString `json:"email_hash,omitempty"`
	Username                              ChallongeString `json:"username,omitempty"`
	AttachedParticipatablePortraitURL     ChallongeString `json:"attached_participatable_portrait_url,omitempty"`
	CanCheckIn                            ChallongeBool   `json:"can_check_in,omitempty"`
	CheckedIn                             ChallongeBool   `json:"checked_in,omitempty"`
	Reactivatable                         ChallongeBool   `json:"reactivatable,omitempty"`
}

func (ip *InternalParticipant) normalize() *Participant {
	return &Participant{ParticipantKey{
		Active:                                bool(ip.Active),
		CheckedInAt:                           time.Time(ip.CheckedInAt),
		CreatedAt:                             time.Time(ip.CreatedAt),
		FinalRank:                             int(ip.FinalRank),
		GroupId:                               int64(ip.GroupId),
		Icon:                                  string(ip.Icon),
		Id:                                    int64(ip.Id),
		InvitationId:                          int64(ip.InvitationId),
		InviteEmail:                           string(ip.InviteEmail),
		Misc:                                  string(ip.Misc),
		Name:                                  string(ip.Name),
		OnWaitingList:                         bool(ip.OnWaitingList),
		Seed:                                  int(ip.Seed),
		ParticipantId:                         int(ip.ParticipantId),
		UpdatedAt:                             time.Time(ip.UpdatedAt),
		ChallongeUsername:                     string(ip.ChallongeUsername),
		ChallongeEmailAddressVerified:         bool(ip.ChallongeEmailAddressVerified),
		Removable:                             bool(ip.Removable),
		ParticipatableOrInvitationAttached:    bool(ip.ParticipatableOrInvitationAttached),
		ConfirmRemove:                         bool(ip.ConfirmRemove),
		InvitationPending:                     bool(ip.InvitationPending),
		DisplayNameWithInvitationEmailAddress: string(ip.DisplayNameWithInvitationEmailAddress),
		EmailHash:                             string(ip.EmailHash),
		Username:                              string(ip.Username),
		AttachedParticipatablePortraitURL:     string(ip.AttachedParticipatablePortraitURL),
		CanCheckIn:                            bool(ip.CanCheckIn),
		CheckedIn:                             bool(ip.CheckedIn),
		Reactivatable:                         bool(ip.Reactivatable),
	}}
}

func (p *Participant) denormalize() *InternalParticipant {
	return &InternalParticipant{InternalParticipantKey{
		Active:                                ChallongeBool(p.Active),
		CheckedInAt:                           ChallongeTime(p.CheckedInAt),
		CreatedAt:                             ChallongeTime(p.CreatedAt),
		FinalRank:                             ChallongeInt(p.FinalRank),
		GroupId:                               ChallongeInt64(p.GroupId),
		Icon:                                  ChallongeString(p.Icon),
		Id:                                    ChallongeInt64(p.Id),
		InvitationId:                          ChallongeInt64(p.InvitationId),
		InviteEmail:                           ChallongeString(p.InviteEmail),
		Misc:                                  ChallongeString(p.Misc),
		Name:                                  ChallongeString(p.Name),
		OnWaitingList:                         ChallongeBool(p.OnWaitingList),
		Seed:                                  ChallongeInt(p.Seed),
		ParticipantId:                         ChallongeInt(p.ParticipantId),
		UpdatedAt:                             ChallongeTime(p.UpdatedAt),
		ChallongeUsername:                     ChallongeString(p.ChallongeUsername),
		ChallongeEmailAddressVerified:         ChallongeBool(p.ChallongeEmailAddressVerified),
		Removable:                             ChallongeBool(p.Removable),
		ParticipatableOrInvitationAttached:    ChallongeBool(p.ParticipatableOrInvitationAttached),
		ConfirmRemove:                         ChallongeBool(p.ConfirmRemove),
		InvitationPending:                     ChallongeBool(p.InvitationPending),
		DisplayNameWithInvitationEmailAddress: ChallongeString(p.DisplayNameWithInvitationEmailAddress),
		EmailHash:                             ChallongeString(p.EmailHash),
		Username:                              ChallongeString(p.Username),
		AttachedParticipatablePortraitURL:     ChallongeString(p.AttachedParticipatablePortraitURL),
		CanCheckIn:                            ChallongeBool(p.CanCheckIn),
		CheckedIn:                             ChallongeBool(p.CheckedIn),
		Reactivatable:                         ChallongeBool(p.Reactivatable),
	}}
}

func normalizeParticipantSlice(ips []InternalParticipant) []*Participant {
	ps := make([]*Participant, 0)
	for _, ip := range ips {
		ps = append(ps, ip.normalize())
	}
	return ps
}

func denormalizeParticipantSlice(ps []*Participant) []InternalParticipant {
	ips := make([]InternalParticipant, 0)
	for _, p := range ps {
		ips = append(ips, *p.denormalize())
	}
	return ips
}

func (c *ChallongeClient) executeParticipantPost(uri string) ([]*Participant, error) {
	ips := make([]InternalParticipant, 0)
	err := c.Post(uri, nil, &ips)
	if err != nil {
		return nil, err
	}
	ps := normalizeParticipantSlice(ips)
	return ps, nil
}

func (c *ChallongeClient) GetParticipants(id string) ([]*Participant, error) {
	uri := fmt.Sprintf("tournament/%s/participants", id)
	ips := make([]InternalParticipant, 0)
	err := c.Get(uri, &ips)
	if err != nil {
		return nil, err
	}
	ps := normalizeParticipantSlice(ips)
	return ps, nil
}

func (c *ChallongeClient) GetParticipant(tid string, pid int64) (*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/%d", tid, pid)
	p := Participant{}
	err := c.Get(uri, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *ChallongeClient) CreateParticipant(id string, p *Participant) (*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants", id)
	np := p.denormalize()
	r := &InternalParticipant{}
	err := c.Post(uri, np, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) BulkAddParticipants(id string, bps *BulkParticipants) ([]*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/bulk_add", id)
	r := make([]InternalParticipant, 0)
	err := c.Post(uri, bps, &r)
	if err != nil {
		return nil, err
	}
	nps := normalizeParticipantSlice(r)
	return nps, nil
}

func (c *ChallongeClient) UpdateParticipant(tid string, pid int64, p *Participant) (*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/%d", tid, pid)
	r := &InternalParticipant{}
	ip := p.denormalize()
	err := c.Put(uri, ip, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) DeleteParticipant(tid string, pid int64) (*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/%d", tid, pid)
	ip := &InternalParticipant{}
	err := c.Delete(uri, ip)
	if err != nil {
		return nil, err
	}
	return ip.normalize(), nil
}

func (c *ChallongeClient) RandomizeParticipants(id string) ([]*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/randomize", id)
	ps, err := c.executeParticipantPost(uri)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *ChallongeClient) CheckInParticipant(tid string, pid int64) (*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/%d/check_in", tid, pid)
	p := &InternalParticipant{}
	err := c.Post(uri, nil, p)
	if err != nil {
		return nil, err
	}
	return p.normalize(), nil
}

func (c *ChallongeClient) UndoCheckInParticipant(tid string, pid int64) (*Participant, error) {
	uri := fmt.Sprintf("tournaments/%s/participants/%d/undo_check_in", tid, pid)
	p := &InternalParticipant{}
	err := c.Post(uri, nil, p)
	if err != nil {
		return nil, err
	}
	return p.normalize(), nil
}

func (c *ChallongeClient) ClearParticipants(id string) error {
	uri := fmt.Sprintf("tournaments/%s/participants/clear", id)
	m := make(map[string]string)
	err := c.Delete(uri, &m)
	if err != nil {
		return err
	}
	return nil
}
