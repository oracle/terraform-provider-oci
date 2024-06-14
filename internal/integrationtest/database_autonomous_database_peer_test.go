// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseAutonomousDatabasePeerDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	DatabaseAutonomousDatabaseRepresentation11 = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("compartment_ocid")},
		"cpu_core_count":                       acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_storage_size_in_tbs":             acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                              acctest.Representation{RepType: acctest.Required, Create: adbName},
		"admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"db_workload":                          acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"character_set":                        acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"display_name":                         acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"is_auto_scaling_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_auto_scaling_for_storage_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":                         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_mtls_connection_required":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"autonomous_maintenance_schedule_type": acctest.Representation{RepType: acctest.Optional, Create: `REGULAR`},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"customer_contacts": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousDatabaseCustomerContactsRepresentation},
		"license_model":     acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		//"scheduled_operations":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousDatabaseScheduledOperationsRepresentation},
		"whitelisted_ips":            acctest.Representation{RepType: acctest.Optional, Create: []string{`1.1.1.1/28`}},
		"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `ENABLED`},
		"timeouts":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseTimeoutsRepresentation},
		"ncharacter_set":             acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}
)

func TestDatabaseAutonomousDatabasePeerResource_basic(t *testing.T) {
	datasourceName := "data.oci_database_autonomous_database_peers.test_autonomous_database_peers"

	httpreplay.SetScenario("TestDatabaseCrossRegionDisasterRecovery_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepresentation11) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_peers", "test_autonomous_database_peers", acctest.Required, acctest.Create, DatabaseAutonomousDatabasePeerDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_peer_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_database_peer_collection.0.items.#", "0"),
				),
			},
		},
	})
}
