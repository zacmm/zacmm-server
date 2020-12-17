// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type PostInfo struct {
	ChannelName string `json:"channel_name,omitempty"`
	TeamName    string `json:"team_name,omitempty"`
	Members     string `json:"members,omitempty"`
}
