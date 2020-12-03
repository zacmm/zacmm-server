// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
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
		table := db.AddTableWithName(model.WhitelistItem{}, "Whitelist").SetKeys(false, "UserId", "IP")
		table.ColMap("UserId").SetMaxSize(26)
		table.ColMap("IP").SetMaxSize(39)
	}

	return s
}

func (s SqlWhitelistStore) createIndexesIfNotExists() {
}

func (s SqlWhitelistStore) Add(whitelistItem *model.WhitelistItem) error {

	if len(whitelistItem.UserId) == 0 {
		return store.NewErrInvalidInput("whitelist item", "user id", whitelistItem.UserId)
	}

	if len(whitelistItem.IP) == 0 {
		return store.NewErrInvalidInput("whitelist item", "ip", whitelistItem.IP)
	}

	if err := s.GetMaster().Insert(whitelistItem); err != nil {
		return errors.Wrapf(err, "failed to save whitelist item with user_id=%s and ip=%s", whitelistItem.UserId, whitelistItem.IP)
	}

	return nil
}

func (s SqlWhitelistStore) Delete(whitelistItem *model.WhitelistItem) error {
	_, err := s.GetMaster().Exec("DELETE FROM Whitelist WHERE UserId = :UserId AND IP = :IP", map[string]interface{}{"UserId": whitelistItem.UserId, "IP": whitelistItem.IP})
	if err != nil {
		return errors.Wrapf(err, "failed to delete from Whitelist with user id=%s and ip=%s", whitelistItem.UserId, whitelistItem.IP)
	}

	return nil
}

func (s SqlWhitelistStore) GetByUserId(userId string) ([]string, error) {
	var ips []string

	query := s.getQueryBuilder().
		Select("IP").
		From("Whitelist").
		Where(sq.Eq{"UserId": userId})

	queryString, args, err := query.ToSql()
	if err != nil {
		return []string{}, errors.Wrap(err, "incoming_whitelist_tosql")
	}

	if _, err := s.GetReplica().Select(&ips, queryString, args...); err != nil {
		return []string{}, errors.Wrap(err, "failed to find ips")
	}

	return ips, nil
}
