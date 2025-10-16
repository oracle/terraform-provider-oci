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

func TestDatabaseDatabaseResource_exadbxs_assisted_patching(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseResource_exadbxs_assisted_patching")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database.test_database"
	scheduleDatabaseResourceName := "data.oci_database_database.schedule_database"

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{

		{
			Config: config + compartmentIdVariableStr + ExaDbXsOracleMangedDatabaseRepresentationDependencies + ExaDbXsScheduleMRDatabaseRepresentationDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, ExaDbXsOracleManagedDatabaseRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.managed_software_update_details.0.is_enrolled", "true"),
				resource.TestCheckResourceAttr(resourceName, "database.0.managed_software_update_details.0.preference_details.0.days_of_week.0", "FRIDAY"),
				resource.TestCheckResourceAttr(resourceName, "database.0.managed_software_update_details.0.preference_details.0.hour_of_day", "14"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				// resource.TestCheckResourceAttr(resourceName, "db_version", "19.0.0.0"), // minor version could be different from the input
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				func(s *terraform.State) error {
					resourceData := s.Modules[0].Resources[scheduleDatabaseResourceName].Primary.Attributes
					fmt.Printf("time Scheduled: %s", resourceData["managed_software_update_details.0.maintenance_details.0.time_scheduled"])
					scheduleTime = resourceData["managed_software_update_details.0.maintenance_details.0.time_scheduled"]
					return nil
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr + ExaDbXsOracleMangedDatabaseRepresentationDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, ExaDbXsOracleManagedDatabaseRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_schedule_management", "test_reschedule_db_software", acctest.Optional, acctest.Create, ScheduleManagedSoftwareManagementRepresentation) +
				ExaDbXsScheduleMRDatabaseRepresentationDependencies,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.managed_software_update_details.0.is_enrolled", "true"),
				resource.TestCheckResourceAttr(resourceName, "database.0.managed_software_update_details.0.preference_details.0.days_of_week.0", "SATURDAY"),
				resource.TestCheckResourceAttr(resourceName, "database.0.managed_software_update_details.0.preference_details.0.hour_of_day", "17"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				// resource.TestCheckResourceAttr(resourceName, "db_version", "19.0.0.0"), // minor version could be different from the input
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + ExaDbXsOracleMangedDatabaseRepresentationDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, ExaDbXsOracleManagedDatabaseRepresentation) +
				ExaDbXsScheduleMRDatabaseRepresentationDependencies,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) error {
					resourceData := s.Modules[0].Resources[scheduleDatabaseResourceName].Primary.Attributes
					actualScheduledTime := resourceData["managed_software_update_details.0.maintenance_details.0.time_scheduled"]

					fmt.Printf(
						"Validating time_scheduled: before=%s, after=%s\n",
						scheduleTime,
						actualScheduledTime,
					)

					if actualScheduledTime == scheduleTime {
						return fmt.Errorf(
							"expected time_scheduled to change, but it did not (before=%s, after=%s)",
							scheduleTime,
							actualScheduledTime,
						)
					}

					return nil
				},
			),
		},
	},
	)
}
