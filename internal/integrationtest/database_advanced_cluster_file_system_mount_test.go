package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAdvancedClusterFileSystemMountRepresentation = map[string]interface{}{
		"advanced_cluster_file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system.id}`},
	}

	DatabaseAdvancedClusterFileSystemMountResourceDependencies = DatabaseAdvancedClusterFileSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Create, DatabaseAdvancedClusterFileSystemRepresentation)
)

func TestDatabaseAdvancedClusterFileSystemMountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdvancedClusterFileSystemMountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_advanced_cluster_file_system_mount.test_advanced_cluster_file_system_mount"
	var resId string
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAdvancedClusterFileSystemMountResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system_mount", "test_advanced_cluster_file_system_mount", acctest.Required, acctest.Create, DatabaseAdvancedClusterFileSystemMountRepresentation), "database", "advancedClusterFileSystemMount", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		{
			Config: config + compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemMountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system_mount", "test_advanced_cluster_file_system_mount", acctest.Required, acctest.Create, DatabaseAdvancedClusterFileSystemMountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "advanced_cluster_file_system_id"),
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
