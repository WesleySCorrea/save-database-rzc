package models

import (
	"encoding/json"
	"time"
)

type Geral struct {
	ClientID                 int             `db:"client_id" json:"clientId"`
	ID                       int64           `db:"id" json:"id"`
	CreationDate             time.Time       `db:"creation_date" json:"creationDate"`
	UpdateDate               time.Time       `db:"update_date" json:"updateDate"`
	Phone                    string          `db:"phone" json:"phone"`
	ParticipantPhone         string          `db:"participant_phone" json:"participantPhone"`
	SenderName               string          `db:"sender_name" json:"senderName"`
	MessageID                string          `db:"message_id" json:"messageId"`
	ReferenceMessageID       string          `db:"reference_message_id" json:"referenceMessageId"`
	Momment                  int64           `db:"momment" json:"momment"`
	IsStatusReply            bool            `db:"is_status_reply" json:"isStatusReply"`
	ChatLid                  string          `db:"chat_lid" json:"chatLid"`
	ConnectedPhone           string          `db:"connected_phone" json:"connectedPhone"`
	WaitingMessage           bool            `db:"waiting_message" json:"waitingMessage"`
	IsEdit                   bool            `db:"is_edit" json:"isEdit"`
	IsGroup                  bool            `db:"is_group" json:"isGroup"`
	IsNewsletter             bool            `db:"is_newsletter" json:"isNewsletter"`
	InstanceID               string          `db:"instance_id" json:"instanceId"`
	FromMe                   bool            `db:"from_me" json:"fromMe"`
	FromApi                  bool            `db:"from_api" json:"fromApi"`
	Status                   string          `db:"status" json:"status"`
	ChatName                 string          `db:"chat_name" json:"chatName"`
	SenderPhoto              string          `db:"sender_photo" json:"senderPhoto"`
	PhotoStatus              string          `db:"photo_status" json:"photoStatus"`
	Broadcast                bool            `db:"broadcast" json:"broadcast"`
	ParticipantLid           string          `db:"participant_lid" json:"participantLid"`
	MessageExpirationSeconds int             `db:"message_expiration_seconds" json:"messageExpirationSeconds"`
	Forwarded                bool            `db:"forwarded" json:"forwarded"`
	Type                     string          `db:"type" json:"type"`
	MessageDateTime          time.Time       `db:"message_date_time" json:"messageDateTime"`
	SenderLid                string          `db:"sender_lid" json:"senderLid"`
	TextMessage              json.RawMessage `db:"text_message" json:"text"`
	Reaction                 json.RawMessage `db:"reaction" json:"reaction"`
	ExternalAdReply          json.RawMessage `db:"external_ad_reply" json:"externalAdReply"`
	Image                    json.RawMessage `db:"image" json:"image"`
	Audio                    json.RawMessage `db:"audio" json:"audio"`
	Video                    json.RawMessage `db:"video" json:"video"`
	PTV                      json.RawMessage `db:"ptv" json:"ptv"`
	Contact                  json.RawMessage `db:"contact" json:"contact"`
	Documento                json.RawMessage `db:"document" json:"document"`
	Location                 json.RawMessage `db:"location" json:"location"`
	Sticker                  json.RawMessage `db:"sticker" json:"sticker"`
	GIF                      json.RawMessage `db:"gif" json:"gif"`
	Poll                     json.RawMessage `db:"poll" json:"poll"`
	PollVote                 json.RawMessage `db:"poll_vote" json:"pollVote"`
}
