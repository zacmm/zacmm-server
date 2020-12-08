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

type SqlInviteStore struct {
	*SqlSupplier
	metrics einterfaces.MetricsInterface
}

func newSqlInviteStore(sqlSupplier *SqlSupplier) store.InviteStore {
	s := &SqlInviteStore{
		SqlSupplier: sqlSupplier,
	}

	for _, db := range sqlSupplier.GetAllConns() {
		table := db.AddTableWithName(model.InviteItem{}, "Invites").SetKeys(false, "InviteId")
		table.ColMap("InviteId").SetMaxSize(26)
		table.ColMap("TeamId").SetMaxSize(26)
	}

	return s
}

func (s SqlInviteStore) createIndexesIfNotExists() {
}

func (s SqlInviteStore) Add(inviteItem *model.InviteItem) error {
	if err := s.GetMaster().Insert(inviteItem); err != nil {
		return errors.Wrapf(err, "failed to save invite item with invtre_id=%s and team_id=%s", inviteItem.InviteId, inviteItem.TeamId)
	}

	return nil
}

func (s SqlInviteStore) Delete(inviteId string) error {
	_, err := s.GetMaster().Exec("DELETE FROM Invites WHERE InviteId = :InviteId", map[string]interface{}{"InviteId": inviteId})
	if err != nil {
		return errors.Wrapf(err, "failed to delete from Invites with invite id=%s", inviteId)
	}

	return nil
}

func (s SqlInviteStore) GetTeamId(inviteId string) (string, error) {
	var teamId string

	query := s.getQueryBuilder().
		Select("TeamId").
		From("Invites").
		Where(sq.Eq{"InviteId": inviteId})

	queryString, args, err := query.ToSql()
	if err != nil {
		return "", errors.Wrap(err, "incoming_invite_tosql")
	}

	if _, err := s.GetReplica().Select(&teamId, queryString, args...); err != nil {
		return "", errors.Wrap(err, "failed to find ips")
	}

	return teamId, nil
}
