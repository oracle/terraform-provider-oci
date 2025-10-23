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
	JmsFleetLibraryManagedInstanceDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
		"library_key":         acctest.Representation{RepType: acctest.Required, Create: `libraryKey`},
		"application_id":      acctest.Representation{RepType: acctest.Optional, Create: `dummy-application-id`},
		"host_name":           acctest.Representation{RepType: acctest.Optional, Create: `hostName`},
		"hostname_contains":   acctest.Representation{RepType: acctest.Optional, Create: `hostnameContains`},
		"managed_instance_id": acctest.Representation{RepType: acctest.Optional, Create: JmsManagedInstanceId},
		"time_end":            acctest.Representation{RepType: acctest.Optional, Create: `2025-07-10T15:15:15.000Z`},
		"time_start":          acctest.Representation{RepType: acctest.Optional, Create: `2025-07-10T15:15:15.000Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetLibraryManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetLibraryManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_library_managed_instances.test_fleet_library_managed_instances"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_library_managed_instances",
					"test_fleet_library_managed_instances",
					acctest.Optional,
					acctest.Create,
					JmsFleetLibraryManagedInstanceDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "host_name", "hostName"),
				resource.TestCheckResourceAttr(datasourceName, "hostname_contains", "hostnameContains"),
				resource.TestCheckResourceAttr(datasourceName, "library_key", "libraryKey"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_start"),

				resource.TestCheckResourceAttrSet(datasourceName, "library_managed_instance_usage_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "library_managed_instance_usage_collection.0.items.#", "0"),
			),
		},
	})
}
