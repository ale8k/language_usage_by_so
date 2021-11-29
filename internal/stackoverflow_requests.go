package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ale8k/language_usage_by_so/pkg/errors"
)

// Base path to 2.3 API of stack exchange
var soUrl = url.URL{Scheme: "https", Host: "api.stackexchange.com", Path: "/2.3"}

type SOQuestionResp struct {
	Items    []SOQuestionItem `json:"items"`
	HasMore  bool             `json:"has_more"`
	QuotaMax int              `json:"quota_max"`
	// TODO: Handle remaing == 0
	QuotaRemaining int `json:"quota_remaining"`
}

// We're owner, don't need deets on them
type SOQuestionItem struct {
	Tags             []string `json:"tags"`
	IsAnswered       bool     `json:"is_answered"`
	ViewCount        int      `json:"view_count"`
	AnswerCount      int      `json:"answer_count"`
	Score            int      `json:"score"`
	LastActivityDate int      `json:"last_activity_date"`
	CreationDate     int      `json:"creation_date"`
	QuestionID       int      `json:"question_id"`
	ContentLicense   string   `json:"content_license"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
}

// Get SO questions created on a specific date
// Some query params are set by default for convenience
// 	- page: "the page index"
// 	- pageSize: "the pageSize relative to the page index"
// 	- frameDateInSeconds: "the date to go from in seconds (it rounds it to nearest 12AM)"
// 	- tags: "the tags to filter questions by, separated by ';' between each tag"
// Returns an error if the unmarshal fails, wrapped with the response code for debugging
func GetCreatedQuestionsSync(page int, pageSize int, fromDateInSeconds int, tags string) (*SOQuestionResp, error) {
	url := getUrlCopy(soUrl)
	url.Path += "/questions"
	itoa := strconv.Itoa

	q := url.Query()
	q.Add("page", itoa(page))
	q.Add("pagesize", itoa(pageSize))
	q.Add("fromdate", itoa(fromDateInSeconds))
	q.Add("tagged", tags)
	q.Add("order", "desc")
	q.Add("sort", "creation")
	q.Add("site", "stackoverflow")
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	errors.HandleErrFatal(err)

	defer resp.Body.Close()

	var respBody *SOQuestionResp

	respBuf, err := io.ReadAll(resp.Body)
	errors.HandleErrFatal(err)

	err = json.Unmarshal(respBuf, &respBody)

	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("error with req, respcode: %v, err: %w", resp.StatusCode, err)
		return nil, err
	}

	return respBody, nil
}

// Copies url into pointer for extension
func getUrlCopy(url url.URL) *url.URL {
	internalUrl := soUrl
	return &internalUrl
}
