// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package commands

import (
	"flag"
	"os"
	"testing"

	"github.com/zacmm/zacmm-server/api4"
	"github.com/zacmm/zacmm-server/mlog"
	"github.com/zacmm/zacmm-server/testlib"
)

func TestMain(m *testing.M) {
	// Command tests are run by re-invoking the test binary in question, so avoid creating
	// another container when we detect same.
	flag.Parse()
	if filter := flag.Lookup("test.run").Value.String(); filter == "ExecCommand" {
		status := m.Run()
		os.Exit(status)
		return
	}

	var options = testlib.HelperOptions{
		EnableStore:     true,
		EnableResources: true,
	}

	mlog.DisableZap()

	mainHelper = testlib.NewMainHelperWithOptions(&options)
	defer mainHelper.Close()
	api4.SetMainHelper(mainHelper)

	mainHelper.Main(m)
}
