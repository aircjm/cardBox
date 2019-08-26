package model

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/util"
)

type AnkiAddDeckRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Deck string `json:"deck"`
	} `json:"params"`
}

type AnkiAddNoteRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Note struct {
			DeckName  string `json:"deckName"`
			ModelName string `json:"modelName"`
			Fields    struct {
				Front string `json:"Front"`
				Back  string `json:"Back"`
			} `json:"fields"`
			Options struct {
				AllowDuplicate bool `json:"allowDuplicate"`
			} `json:"options"`
			Tags []string `json:"tags"`
			//Audio struct {
			//	URL      string   `json:"url"`
			//	Filename string   `json:"filename"`
			//	SkipHash string   `json:"skipHash"`
			//	Fields   []string `json:"fields"`
			//} `json:"audio"`
		} `json:"note"`
	} `json:"params"`
}

type AnkiUpdateNoteFieldsRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Note struct {
			ID     int64 `json:"id"`
			Fields struct {
				Front string `json:"Front"`
				Back  string `json:"Back"`
			} `json:"fields"`
		} `json:"note"`
	} `json:"params"`
}

func (AnkiAddNoteRequest) GetAnkiAddNote(card *trello.Card) (addNoteAnkiRequest *AnkiAddNoteRequest) {
	markdownHtml := util.ConvertMarkdown(card.Desc)
	request := new(AnkiAddNoteRequest)
	request.Action = "addNote"
	request.Version = 6
	request.Params.Note.Fields.Front = card.Name
	request.Params.Note.Fields.Back = markdownHtml
	request.Params.Note.ModelName = "trelloAdd"
	labels := card.Labels
	tags := []string{}
	if len(labels) > 0 {
		for e := range labels {
			label := labels[e]
			tags = append(tags, label.Name)
		}
	}
	request.Params.Note.Tags = tags
	request.Params.Note.Options.AllowDuplicate = false
	return request
}

type AnkiResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error"`
}
