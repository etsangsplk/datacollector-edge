package dpm

import "github.com/streamsets/datacollector-edge/container/common"

const (
	VALIDATE_PIPELINE                 = 1000
	SAVE_PIPELINE                     = 1001
	SAVE_RULES_PIPELINE               = 1002
	START_PIPELINE                    = 1003
	STOP_PIPELINE                     = 1004
	RESET_OFFSET_PIPELINE             = 1005
	DELETE_PIPELINE                   = 1006
	DELETE_HISTORY_PIPELINE           = 1007
	PING_FREQUENCY_ADJUSTMENT         = 1008
	STOP_DELETE_PIPELINE              = 1009
	SSO_DISCONNECTED_MODE_CREDENTIALS = 1010
	SYNC_ACL                          = 1011
	STATUS_PIPELINE                   = 2000
	SDC_INFO_EVENT                    = 2001
	STATUS_MULTIPLE_PIPELINES         = 2002
	ACK_EVENT                         = 5000

	ACK_EVENT_SUCCESS = "SUCCESS"
	ACK_EVENT_ERROR   = "ERROR"
	ACK_EVENT_IGNORE  = "IGNORE"
)

type ClientEvent struct {
	EventId      string   `json:"eventId"`
	Destinations []string `json:"destinations"`
	RequiresAck  bool     `json:"requiresAck"`
	IsAckEvent   bool     `json:"ackEvent"`
	EventTypeId  int      `json:"eventTypeId"`
	Payload      string   `json:"payload"`
	OrgId        string   `json:"orgId"`
}

type ServerEvent struct {
	EventId      string `json:"eventId"`
	From         string `json:"from"`
	RequiresAck  bool   `json:"requiresAck"`
	IsAckEvent   bool   `json:"isAckEvent"`
	EventTypeId  int    `json:"eventTypeId"`
	Payload      string `json:"payload"`
	ReceivedTime int64  `json:"receivedTime"`
	OrgId        string `json:"orgId"`
}

type AckEvent struct {
	AckEventStatus string `json:"ackEventStatus"`
	Message        string `json:"message"`
}

type SDCInfoEvent struct {
	EdgeId        string            `json:"sdcId"`
	HttpUrl       string            `json:"httpUrl"`
	GoVersion     string            `json:"javaVersion"`
	EdgeBuildInfo *common.BuildInfo `json:"sdcBuildInfo"`
	Labels        []string          `json:"labels"`
	Edge          bool              `json:"edge"`
}

type PipelineBaseEvent struct {
	Name string `json:"name"`
	Rev  string `json:"rev"`
	User string `json:"user"`
}

type PipelineSaveEvent struct {
	Name                          string                        `json:"name"`
	Rev                           string                        `json:"rev"`
	User                          string                        `json:"user"`
	PipelineConfigurationAndRules PipelineConfigurationAndRules `json:"pipelineConfigurationAndRules"`
	Description                   string                        `json:"description"`
	Offset                        string                        `json:"offset"`
	OffsetProtocolVersion         float64                       `json:"offsetProtocolVersion"`
	Acl                           interface{}                   `json:"acl"`
}

type PipelineConfigurationAndRules struct {
	PipelineConfig string `json:"pipelineConfig"`
	PipelineRules  string `json:"pipelineRules"`
}