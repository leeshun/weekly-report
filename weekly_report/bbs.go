package weekly_report

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	categoryID = flag.Int("category_id", 0, "category id")
	topicID    = flag.Int("topic_id", 0, "topic id")
	postURL    = flag.String("post_url", "", "bbs post url")
)

type requestBody struct {
	Raw                       string `json:"raw"`
	UnListTopic               bool   `json:"unlist_topic"`
	Category                  int    `json:"category"`
	TopicId                   int    `json:"topic_id"`
	IsWarning                 bool   `json:"is_warning"`
	Archetype                 string `json:"archetype"`
	TypingDurationMsecs       int    `json:"typing_duration_msecs"`
	ComposerOpenDurationMsecs int    `json:"composer_open_duration_msecs"`
	SharedDraft               bool   `json:"shared_draft"`
	NestedPost                bool   `json:"nested_post"`
}

func makeRequestBody(data string) (io.Reader, error) {
	req := requestBody{
		Raw:                       data,
		UnListTopic:               false,
		Category:                  *categoryID,
		TopicId:                   *topicID,
		IsWarning:                 false,
		Archetype:                 "regular",
		TypingDurationMsecs:       1000,
		ComposerOpenDurationMsecs: 18750,
		SharedDraft:               false,
		NestedPost:                true,
	}
	d, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data into json with cause %w", err)
	}
	return bytes.NewReader(d), nil
}

func Post(data string) error {
	body, err := makeRequestBody(data)
	if err != nil {
		return fmt.Errorf("failed to make request body with cause %w", err)
	}
	_, err = http.Post(*postURL, "application/x-www-form-urlencoded; charset=UTF-8", body)
	return err
}
