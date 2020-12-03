// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package localcachelayer

import (
	"github.com/zacmm/zacmm-server/model"
	"github.com/zacmm/zacmm-server/store"
)

type LocalCacheWhitelistStore struct {
	store.WhitelistStore
	rootStore *LocalCacheStore
}

func (s LocalCacheWhitelistStore) GetByUserId(userId string) ([]string, error) {
	return s.WhitelistStore.GetByUserId(userId)
}

func (s LocalCacheWhitelistStore) Add(item *model.WhitelistItem) error {
	return s.WhitelistStore.Add(item)
}

func (s LocalCacheWhitelistStore) Delete(item *model.WhitelistItem) error {
	return s.WhitelistStore.Delete(item)
}
