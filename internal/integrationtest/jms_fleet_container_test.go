// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsFleetContainerFleetId           = utils.GetEnvSettingWithBlankDefault("fleet_ocid")
	JmsFleetContainerManagedInstanceId = utils.GetEnvSettingWithBlankDefault("managed_instance_ocid")

	JmsFleetContainerDataSourceRepresentation = map[string]interface{}{
		"fleet_id":                              acctest.Representation{RepType: acctest.Required, Create: JmsFleetContainerFleetId},
		"application_name":                      acctest.Representation{RepType: acctest.Optional, Create: `dummy-application-name`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"jre_security_status":                   acctest.Representation{RepType: acctest.Optional, Create: `EARLY_ACCESS`},
		"jre_version":                           acctest.Representation{RepType: acctest.Optional, Create: `jreVersion`},
		"managed_instance_id":                   acctest.Representation{RepType: acctest.Optional, Create: JmsFleetContainerManagedInstanceId},
		"time_started_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2025-07-10T15:15:15.000Z`},
		"time_started_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `2025-07-10T15:15:15.000Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetContainerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetContainerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_containers.test_fleet_containers"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_containers",
					"test_fleet_containers",
					acctest.Optional,
					acctest.Create,
					JmsFleetContainerDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_name"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "jre_security_status", "EARLY_ACCESS"),
				resource.TestCheckResourceAttr(datasourceName, "jre_version", "jreVersion"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_started_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_started_less_than_or_equal_to"),

				resource.TestCheckResourceAttrSet(datasourceName, "container_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "container_collection.0.items.#", "0"),
			),
		},
	})
}
