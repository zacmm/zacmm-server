// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"github.com/pkg/errors"

	"github.com/zacmm/zacmm-server/einterfaces"
	"github.com/zacmm/zacmm-server/model"
	"github.com/zacmm/zacmm-server/store"
)

type SqlWhitelistStore struct {
	*SqlSupplier
	metrics einterfaces.MetricsInterface
}

func newSqlWhitelistStore(sqlSupplier *SqlSupplier) store.WhitelistStore {
	s := &SqlWhitelistStore{
		SqlSupplier: sqlSupplier,
	}

	for _, db := range sqlSupplier.GetAllConns() {
		table := db.AddTableWithName(model.WhitelistItem{}, "Whitelist").SetKeys(false, "IP")
		table.ColMap("IP").SetMaxSize(39)
	}

	return s
}

func (s SqlWhitelistStore) createIndexesIfNotExists() {
}

func (s SqlWhitelistStore) Add(whitelistItem *model.WhitelistItem) error {

	if len(whitelistItem.IP) == 0 {
		return store.NewErrInvalidInput("whitelist item", "ip", whitelistItem.IP)
	}

	if err := s.GetMaster().Insert(whitelistItem); err != nil {
		return errors.Wrapf(err, "failed to save whitelist item with ip=%s", whitelistItem.IP)
	}

	return nil
}

func (s SqlWhitelistStore) Delete(whitelistItem *model.WhitelistItem) error {
	_, err := s.GetMaster().Exec("DELETE FROM Whitelist WHERE IP = :IP", map[string]interface{}{"IP": whitelistItem.IP})
	if err != nil {
		return errors.Wrapf(err, "failed to delete from Whitelist with ip=%s", whitelistItem.IP)
	}

	return nil
}

func (s SqlWhitelistStore) Get() ([]string, error) {
	var ips []string

	query := s.getQueryBuilder().
		Select("IP").
		From("Whitelist")

	queryString, args, err := query.ToSql()
	if err != nil {
		return []string{}, errors.Wrap(err, "incoming_whitelist_tosql")
	}

	if _, err := s.GetReplica().Select(&ips, queryString, args...); err != nil {
		return []string{}, errors.Wrap(err, "failed to find ips")
	}

	return ips, nil
}
