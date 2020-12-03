// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

// WebSocketRequest represents a request made to the server through a websocket.
type WhitelistItem struct {
	UserId string `json:"user_id"` // User Id
	IP     string `json:"ip"`      // Ip from the whitelist
}

func (o *WhitelistItem) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func WhitelistItemFromJson(data io.Reader) *WhitelistItem {
	var o *WhitelistItem
	json.NewDecoder(data).Decode(&o)
	return o
}
