package lytics

import (
  "time"
  "strconv"
)

const (
  userRecommendEndpoint = "content/recommend/user/:fieldName/:fieldVal"
)

type Document struct {
  Url             string             `json:"url"`
  Title           string             `json:"title"`
  Description     string             `json:"description"`
  Topics          []string           `json:"topics"`
  TopicRelevances map[string]float64 `json:"topic_relevances"`
  PrimaryImage    string             `json:"primary_image"`
  Author          string             `json:"author"`
  Created         time.Time          `json:"created"`
  Hash            string             `json:"id"`
  Sitename        string             `json:"sitename,omitempty"`
  Stream          string             `json:"stream"`
  Path            []string           `json:"path,omitempty"`
  Aspects         []string           `json:"aspects,omitempty"`
  Language        string             `json:"language,omitempty"`
  Updated         time.Time          `json:"updated,omitempty"`
  Fetched         time.Time          `json:"fetched,omitempty"`
  Meta            []string           `json:"meta,omitempty"`
  Id              string             `json:"hashedurl,omitempty"`
}

type Recommendation struct {
  *Document
  Confidence      float64            `json:"confidence"`
  Visited         bool               `json:"visited"`
  VisitRate       float64            `json:"visitrate,omitempty"`
}

// GetUserContentRecommendation returns
// a list of documents to recommend the user
// based on their content affinity
func (l *Client) GetUserContentRecommendation(fieldName, fieldVal, ql string, limit int, shuffle bool) ([]Recommendation, error) {
  res := ApiResp{}
  data := []Recommendation{}
  params := map[string]string{}

  if limit > -1 {
    params["limit"] = strconv.Itoa(limit)
  }

  if shuffle {
    params["shuffle"] = "true"
  }

  if ql != "" {
    params["ql"] = ql
  }

  // make the request
  err := l.Get(parseLyticsURL(userRecommendEndpoint, map[string]string{"fieldName": fieldName, "fieldVal": fieldVal}), params, nil, &res, &data)
  if err != nil {
    return data, err
  }

  return data, nil
}
