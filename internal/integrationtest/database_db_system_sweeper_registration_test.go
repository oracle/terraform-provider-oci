// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"sync"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var databaseDbSystemSweeperRegisterOnce sync.Once

func registerDatabaseDbSystemSweeper() {
	databaseDbSystemSweeperRegisterOnce.Do(func() {
		if acctest.DependencyGraph == nil {
			acctest.InitDependencyGraph()
		}
		resource.AddTestSweepers("DatabaseDbSystem", &resource.Sweeper{
			Name:         "DatabaseDbSystem",
			Dependencies: acctest.DependencyGraph["dbSystem"],
			F:            sweepDatabaseDbSystemResource,
		})
	})
}
