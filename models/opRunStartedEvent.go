package models

type OpRunStartedEvent struct {
  OpRunId       string `json:"opRunId"`
  OpRef         string `json:"opRef"`
  ParentOpRunId string `json:"parentOpRunId"`
  RootOpRunId   string `json:"rootOpRunId"`
}
