// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package localcachelayer

import (
	"github.com/zacmm/zacmm-server/model"
	"github.com/zacmm/zacmm-server/store"
)

type LocalCacheInviteStore struct {
	store.InviteStore
	rootStore *LocalCacheStore
}

func (s LocalCacheInviteStore) GetTeamId(inviteId string) (string, error) {
	return s.InviteStore.GetTeamId(inviteId)
}

func (s LocalCacheInviteStore) Add(item *model.InviteItem) error {
	return s.InviteStore.Add(item)
}

func (s LocalCacheInviteStore) Delete(inviteId string) error {
	return s.InviteStore.Delete(item)
}
