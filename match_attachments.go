package challonge

import (
	"fmt"
	"time"
)

// Use these structs for any PUT or POST
type MatchAttachment struct {
	MatchAttachmentKey `json:"match_attachment"`
}

type MatchAttachmentKey struct {
	Id               int64     `json:"id,omitempty"`
	MatchId          int64     `json:"match_id,omitempty"`
	UserId           int64     `json:"user_id,omitempty"`
	Description      string    `json:"description,omitempty"`
	Url              string    `json:"url,omitempty"`
	OriginalFileName int       `json:"original_file_name,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	AssetFileName    string    `json:"asset_file_name,omitempty"`
	AssetContentType string    `json:"asset_content_type,omitempty"`
	AssetFileSize    int       `json:"asset_file_size,omitempty"`
	AssetUrl         string    `json:"asset_url,omitempty"`
}

// Internal structs that are normalized for null fields
type InternalMatchAttachment struct {
	InternalMatchAttachmentKey `json:"match_attachment"`
}

type InternalMatchAttachmentKey struct {
	Id               ChallongeInt64  `json:"id"`
	MatchId          ChallongeInt64  `json:"match_id"`
	UserId           ChallongeInt64  `json:"user_id"`
	Description      ChallongeString `json:"description"`
	Url              ChallongeString `json:"url"`
	OriginalFileName ChallongeInt    `json:"original_file_name"`
	CreatedAt        ChallongeTime   `json:"created_at"`
	UpdatedAt        ChallongeTime   `json:"updated_at"`
	AssetFileName    ChallongeString `json:"asset_file_name"`
	AssetContentType ChallongeString `json:"asset_content_type"`
	AssetFileSize    ChallongeInt    `json:"asset_file_size"`
	AssetUrl         ChallongeString `json:"asset_url"`
}

func (ima *InternalMatchAttachment) normalize() *MatchAttachment {
	return &MatchAttachment{MatchAttachmentKey{
		Id:               int64(ima.Id),
		MatchId:          int64(ima.MatchId),
		UserId:           int64(ima.UserId),
		Description:      string(ima.Description),
		Url:              string(ima.Url),
		OriginalFileName: int(ima.OriginalFileName),
		CreatedAt:        time.Time(ima.CreatedAt),
		UpdatedAt:        time.Time(ima.UpdatedAt),
		AssetFileName:    string(ima.AssetFileName),
		AssetContentType: string(ima.AssetContentType),
		AssetFileSize:    int(ima.AssetFileSize),
		AssetUrl:         string(ima.AssetUrl),
	}}
}

func (ma *MatchAttachment) denormalize() *InternalMatchAttachment {
	return &InternalMatchAttachment{InternalMatchAttachmentKey{
		Id:               ChallongeInt64(ma.Id),
		MatchId:          ChallongeInt64(ma.MatchId),
		UserId:           ChallongeInt64(ma.UserId),
		Description:      ChallongeString(ma.Description),
		Url:              ChallongeString(ma.Url),
		OriginalFileName: ChallongeInt(ma.OriginalFileName),
		CreatedAt:        ChallongeTime(ma.CreatedAt),
		UpdatedAt:        ChallongeTime(ma.UpdatedAt),
		AssetFileName:    ChallongeString(ma.AssetFileName),
		AssetContentType: ChallongeString(ma.AssetContentType),
		AssetFileSize:    ChallongeInt(ma.AssetFileSize),
		AssetUrl:         ChallongeString(ma.AssetUrl),
	}}
}

func (c *ChallongeClient) GetMatchAttachments(tid string, mid int64) ([]*MatchAttachment, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/match_attachments", tid, mid)
	imas := make([]InternalMatchAttachment, 0)
	mas := make([]*MatchAttachment, 0)
	err := c.Get(uri, &imas)
	if err != nil {
		return nil, err
	}
	for _, ima := range imas {
		mas = append(mas, ima.normalize())
	}
	return mas, nil
}

func (c *ChallongeClient) GetMatchAttachment(tid string, mid, aid int64) (*MatchAttachment, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/match_attachments/%d", tid, mid, aid)
	ma := &InternalMatchAttachment{}
	err := c.Get(uri, ma)
	if err != nil {
		return nil, err
	}
	return ma.normalize(), nil
}

func (c *ChallongeClient) CreateMatchAttachment(tid string, mid int64, ma *MatchAttachment) (*MatchAttachment, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/match_attachments", tid, mid)
	ima := ma.denormalize()
	r := &InternalMatchAttachment{}
	err := c.Post(uri, ima, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) UpdateMatchAttachment(tid string, mid, aid int64, ma *MatchAttachment) (*MatchAttachment, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/match_attachments/%d", tid, mid, aid)
	ima := ma.denormalize()
	r := &InternalMatchAttachment{}
	err := c.Put(uri, ima, r)
	if err != nil {
		return nil, err
	}
	return r.normalize(), nil
}

func (c *ChallongeClient) DeleteMatchAttachment(tid string, mid, aid int64) (*MatchAttachment, error) {
	uri := fmt.Sprintf("tournaments/%s/matches/%d/match_attachments/%d", tid, mid, aid)
	ma := &InternalMatchAttachment{}
	err := c.Delete(uri, ma)
	if err != nil {
		return nil, err
	}
	return ma.normalize(), nil
}
