package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/model"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type SlackRequest struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
	// Blocks      []Blocks      `json:"blocks"`
}

type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

type Text struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

type Field struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Block struct {
	Type     string  `json:"type"`
	Text     *Text   `json:"text,omitempty"`
	Elements []Field `json:"elements,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

// mandatory, first position
func New(headerText string) *SlackRequest {

	executionTimeBlock := Block{
		Type: "context",
		Elements: []Field{
			{
				Type: "mrkdwn",
				Text: fmt.Sprintf("Executed at: %v", time.Now().Format(config.Common.TimeFormat)),
			},
		},
	}

	return &SlackRequest{
		Text: headerText,
		Attachments: []Attachment{
			{
				// Color: attachmentColor,
				Blocks: []Block{
					// headerBlock,
					// statusBlock,
					executionTimeBlock,
					// databaseBlock,
				},
			},
		},
	}
}

// mandatory, second position
func (s *SlackRequest) SetStatus(err error) *SlackRequest {
	var statusText, attachmentColor string

	switch err {
	case nil:
		statusText = ":large_green_circle: *Success*"
		attachmentColor = "#36a64f"
	default:
		statusText = ":red_circle: *Failed*"
		attachmentColor = "#db0f42"
	}

	statusBlock := Block{
		Type: "context",
		Elements: []Field{
			{
				Type: "mrkdwn",
				Text: fmt.Sprintf("Status: %v", statusText),
			},
		},
	}

	finishedTimeBlock := Block{
		Type: "context",
		Elements: []Field{
			{
				Type: "mrkdwn",
				Text: fmt.Sprintf("Finished at: %v", time.Now().Format(config.Common.TimeFormat)),
			},
		},
	}

	return &SlackRequest{
		Text: s.Text,
		Attachments: []Attachment{
			{
				Color: attachmentColor,
				Blocks: []Block{
					statusBlock,
					s.Attachments[0].Blocks[0],
					finishedTimeBlock,
				},
			},
		},
	}
}

// optional, third position
func (s *SlackRequest) ReportCheck(manifests []model.ExpiryData) *SlackRequest {
	InformationBlock := Block{
		Type: "section",
		Fields: []Field{
			{
				Type: "mrkdwn",
				Text: "*Checked Domain :*\n",
			},
		},
	}
	s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, InformationBlock)

	//add divider
	Spacer := Block{
		Type: "divider",
	}

	numbercounter := 1
	for _, v := range manifests {
		s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, Spacer)

		ReportBlock := Block{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("\n %v. https://%v | expired upon *%v* days, *(%v)*", numbercounter, v.Domainname, v.Remainingdays, v.Expireddate),
			},
		}
		numbercounter = numbercounter + 1
		s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, ReportBlock)
	}
	return s
}

// optional, third position notify ssl need to be updated
func (s *SlackRequest) ReminderSlack(manifestsunderthreshold []model.ExpiryData) *SlackRequest {
	numbercounter := 1
	//add divider
	Spacer := Block{
		Type: "divider",
	}
	for _, v := range manifestsunderthreshold {

		s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, Spacer)

		ReportBlock := Block{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("\n %v. please update the SSL of https://%v as it will expire in *%v* days", numbercounter, v.Domainname, v.Remainingdays),
			},
		}
		numbercounter = numbercounter + 1
		s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, ReportBlock)
	}
	return s
}

func (s *SlackRequest) ErrorReportSlack(manifesterror []model.ExpiryData) *SlackRequest {
	numbercounter := 1
	//add divider
	Spacer := Block{
		Type: "divider",
	}
	for _, v := range manifesterror {

		s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, Spacer)

		ReportBlock := Block{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("\n %v. please check the connection to or the manifest of https://%v | we countered an error", numbercounter, v.Domainname),
			},
		}
		numbercounter = numbercounter + 1
		s.Attachments[0].Blocks = append(s.Attachments[0].Blocks, ReportBlock)
	}
	return s
}

// mandatory, last position
func (s SlackRequest) Send() {

	body, _ := json.Marshal(s)
	log.Debug().Str("job", "notify").Str("request", string(body)).Send()

	req, err := http.NewRequest(http.MethodPost, config.Common.Slackwebhook, bytes.NewBuffer(body))
	client := &http.Client{Timeout: 10 * time.Second}
	if err != nil {
		log.Info().Str("job", "request construct").Str("error", err.Error()).Send()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Info().Str("job", "send slack request").Str("error", err.Error()).Send()
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	log.Debug().Str("job", "notify").Str("response", buf.String()).Send()

	if buf.String() != "ok" {
		log.Debug().Str("job", "notify").Str("response", buf.String()).Str("error", "non-ok response from slack").Send()
		return
	}

}
