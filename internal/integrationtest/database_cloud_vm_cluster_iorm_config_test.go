package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	CloudVmClusterIormConfigRequiredOnlyResource = CloudVmClusterIormConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Required, acctest.Create, cloudVmClusterIormConfigRepresentation)

	CloudVmClusterIormConfigResourceConfig = CloudVmClusterIormConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Optional, acctest.Update, cloudVmClusterIormConfigRepresentation)

	cloudVmClusterIormConfigSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
	}

	cloudVmClusterIormConfigRepresentation = map[string]interface{}{
		"cloud_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
		"objective":           acctest.Representation{RepType: acctest.Optional, Create: `AUTO`, Update: `BALANCED`},
		"db_plans":            acctest.RepresentationGroup{RepType: acctest.Required, Group: dbPlanRepresentation},
	}

	CloudVmClusterIormConfigResourceDependencies = AvailabilityDomainConfig + CloudVmClusterRequiredOnlyResource
)

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudVmClusterIormConfigResource_basic(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseCloudVmClusterIormConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_cloud_vm_cluster_iorm_config.test_cloud_vm_cluster_iorm_config"

	singularDatasourceName := "data.oci_database_cloud_vm_cluster_iorm_config.test_cloud_vm_cluster_iorm_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudVmClusterIormConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Optional, acctest.Create, cloudVmClusterIormConfigRepresentation), "database", "exadataIormConfig", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterIormConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Required, acctest.Create, cloudVmClusterIormConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "db_plans.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterIormConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterIormConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Optional, acctest.Create, cloudVmClusterIormConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "objective", "AUTO"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudVmClusterIormConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Optional, acctest.Update, cloudVmClusterIormConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "objective", "BALANCED"),
				resource.TestCheckResourceAttr(resourceName, "db_plans.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_cloud_vm_cluster_iorm_config", "test_cloud_vm_cluster_iorm_config", acctest.Required, acctest.Create, cloudVmClusterIormConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudVmClusterIormConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_vm_cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "db_plans.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "objective", "BALANCED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
