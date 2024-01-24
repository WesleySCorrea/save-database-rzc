package models

type Groups struct {
	ID        int    `db:"id" json:"id"`
	ClientID  int    `db:"client_id" json:"clientId"`
	NameGroup string `db:"name_group" json:"nameGroup"`
}
