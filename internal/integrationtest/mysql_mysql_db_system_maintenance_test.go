// Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	mysqlDbSystemMaintenance = map[string]interface{}{
		// standard required properties
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},

		// use an easier to track display name
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `TestDbSystemMaintenanceCvup`},

		// avoid wasting time setting up DBM when that's not what we're testing here
		"database_management": acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},

		// disable backup policy to avoid wasting even more time and resources
		"backup_policy": acctest.RepresentationGroup{RepType: acctest.Required, Group: disabledBackupPolicy},

		// create with custom maintenance details and then change the values on update
		"maintenance": acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemMaintenanceRepresentation},
	}

	mysqlDbSystemMaintenanceRepresentation = map[string]interface{}{
		"window_start_time":         acctest.Representation{RepType: acctest.Required, Create: `sun 01:00`},
		"maintenance_schedule_type": acctest.Representation{RepType: acctest.Optional, Create: `EARLY`, Update: `REGULAR`},
		"version_preference":        acctest.Representation{RepType: acctest.Optional, Create: `OLDEST`, Update: `SECOND_NEWEST`},
		"version_track_preference":  acctest.Representation{RepType: acctest.Optional, Create: `LONG_TERM_SUPPORT`, Update: `INNOVATION`},
	}
)

func TestMysqlMysqlDbSystemResource_maintenance(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_maintenance")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with custom CVUP values
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemMaintenance),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.maintenance_schedule_type", "EARLY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.version_preference", "OLDEST"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.version_track_preference", "LONG_TERM_SUPPORT"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance.0.target_version"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance.0.time_scheduled"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify it is possible to change the CVUP details
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemMaintenance),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.version_preference", "SECOND_NEWEST"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.version_track_preference", "INNOVATION"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance.0.target_version"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance.0.time_scheduled"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// clear before next Create
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies,
		},

		// verify Create with default CVUP values
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, MysqlMysqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.maintenance_schedule_type", "REGULAR"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.version_preference", "OLDEST"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.version_track_preference", "FOLLOW"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance.0.target_version"),
				resource.TestCheckResourceAttrSet(resourceName, "maintenance.0.time_scheduled"),
				resource.TestCheckResourceAttr(resourceName, "maintenance.0.window_start_time", "sun 01:00"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}
