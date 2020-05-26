package challonge

import (
	"net/http"
	"testing"
)

func TestGetMatchAttachments(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/match_attachments/get_all_match_attachments.json")
	server.StartTLS()
	defer server.Close()

	var id int64 = 201537681
	mas, err := c.GetMatchAttachments("challonge_go_api_mock", id)
	if err != nil {
		t.Error(err)
	}
	if len(mas) != 1 {
		t.Errorf("expected number of match attachments to be 1 but was instead %d", len(mas))
	}
	if mas[0].Url != "https://google.com" {
		t.Errorf("expected url to be 'https://google.com' but was instead %s", mas[0].Url)
	}
}

func TestGetMatchAttachment(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/match_attachments/get_match_attachment.json")
	server.StartTLS()
	defer server.Close()

	var mid int64 = 201537681
	var maid int64 = 513834
	ma, err := c.GetMatchAttachment("challonge_go_api_mock", mid, maid)
	if err != nil {
		t.Error(err)
	}
	if !IsNull(ma.AssetFileName) {
		t.Errorf("expected asset_file_name to be null but got non-null")
	}
}

func TestCreateMatchAttachment(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/match_attachments/create_match_attachment.json")
	server.StartTLS()
	defer server.Close()

	mak := &MatchAttachment{MatchAttachmentKey{
		Url:         "https://google.com",
		Description: "this is a match attachment!",
	}}

	var mid int64 = 201537681
	ma, err := c.CreateMatchAttachment("challonge_go_api_mock", mid, mak)
	if err != nil {
		t.Error(err)
	}
	if !IsNull(ma.AssetFileSize) {
		t.Errorf("expected asset_file_size to be null but got non-null")
	}
}

func TestUpdateMatchAttachment(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/match_attachments/update_match_attachment.json")
	server.StartTLS()
	defer server.Close()

	mak := &MatchAttachment{MatchAttachmentKey{
		Url:         "https://twitter.com",
		Description: "this is a match attachment update!",
	}}

	var mid int64 = 201537681
	var maid int64 = 513834
	ma, err := c.UpdateMatchAttachment("challonge_go_api_mock", mid, maid, mak)
	if err != nil {
		t.Error(err)
	}
	if ma.Url != "https://twitter.com" {
		t.Errorf("expected url to be 'https://twitter.com' but got %s", ma.Url)
	}
}

func TestDeleteMatchAttachment(t *testing.T) {
	c, server := testClientFile(http.StatusOK, "test_data/match_attachments/delete_match_attachment.json")
	server.StartTLS()
	defer server.Close()

	var mid int64 = 201537681
	var maid int64 = 513834
	ma, err := c.DeleteMatchAttachment("challonge_go_api_mock", mid, maid)
	if err != nil {
		t.Error(err)
	}
	if !IsNull(ma.AssetContentType) {
		t.Errorf("expected asset_content_type to be null but got non-null")
	}
}
