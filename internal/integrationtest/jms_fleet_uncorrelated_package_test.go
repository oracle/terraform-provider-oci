// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsFleetUncorrelatedPackageDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
		"application_id":      acctest.Representation{RepType: acctest.Optional, Create: `dummy-application-id`},
		"managed_instance_id": acctest.Representation{RepType: acctest.Optional, Create: JmsManagedInstanceId},
		"package_name":        acctest.Representation{RepType: acctest.Optional, Create: `packageName`},
		"time_end":            acctest.Representation{RepType: acctest.Optional, Create: `2025-07-10T15:15:15.000Z`},
		"time_start":          acctest.Representation{RepType: acctest.Optional, Create: `2025-07-10T15:15:15.000Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetUncorrelatedPackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetUncorrelatedPackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_uncorrelated_packages.test_fleet_uncorrelated_packages"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_uncorrelated_packages",
					"test_fleet_uncorrelated_packages",
					acctest.Optional,
					acctest.Create,
					JmsFleetUncorrelatedPackageDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "package_name", "packageName"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_start"),

				resource.TestCheckResourceAttrSet(datasourceName, "uncorrelated_package_usage_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "uncorrelated_package_usage_collection.0.items.#", "0"),
			),
		},
	})
}
