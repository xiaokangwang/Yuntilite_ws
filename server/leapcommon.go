package main

type YuntiLeapAction struct {
	ActionType string `json:"ActionType"`
	StreamId   string `json:"StreamId"`
	Data       string `json:"Data"`
	seqid      uint64 `json:"seqid"`
}

type YuntiLeapActionGroup struct {
	ActionGroup [24]YuntiLeapAction
	InstToken   string
	InstId      string
}
