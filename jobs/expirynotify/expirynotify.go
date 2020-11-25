// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package expirynotify

import (
	"github.com/zacmm/zacmm-server/app"
	tjobs "github.com/zacmm/zacmm-server/jobs/interfaces"
)

type ExpiryNotifyJobInterfaceImpl struct {
	App *app.App
}

func init() {
	app.RegisterJobsExpiryNotifyJobInterface(func(a *app.App) tjobs.ExpiryNotifyJobInterface {
		return &ExpiryNotifyJobInterfaceImpl{a}
	})
}
