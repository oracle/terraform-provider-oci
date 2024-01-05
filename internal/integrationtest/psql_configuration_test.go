// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsqlConfigurationRequiredOnlyResource = PsqlConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Required, acctest.Create, PsqlConfigurationRepresentation)

	PsqlConfigurationResourceConfig = PsqlConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Update, PsqlConfigurationRepresentation)

	PsqlConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_psql_configuration.test_configuration.id}`},
	}

	PsqlConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_version":     acctest.Representation{RepType: acctest.Optional, Create: `14`},
		"shape":          acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E4.Flex.4.64GB`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlConfigurationDataSourceFilterRepresentation},
	}
	PsqlConfigurationDisplayNameDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `terraform-test-config`, Update: `terraform-test-config-2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlConfigurationDataSourceFilterRepresentation},
	}
	PsqlConfigurationIDDataSourceRepresentation = map[string]interface{}{
		"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_psql_configuration.test_configuration.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlConfigurationDataSourceFilterRepresentation},
	}

	PsqlConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_psql_configuration.test_configuration.id}`}},
	}

	PsqlConfigurationRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_configuration_overrides":  acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlConfigurationDbConfigurationOverridesRepresentation},
		"db_version":                  acctest.Representation{RepType: acctest.Required, Create: `14`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `terraform-test-config`, Update: `terraform-test-config-2`},
		"instance_memory_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `64`},
		"instance_ocpu_count":         acctest.Representation{RepType: acctest.Required, Create: `4`},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description1`, Update: `description2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"system_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sys-namespace.tag-key": "value"}},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesRep},
	}

	ignoreChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`, `freeform_tags`}},
	}

	PsqlConfigurationDbConfigurationOverridesRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlConfigurationDbConfigurationOverridesItemsRepresentation},
	}
	PsqlConfigurationDbConfigurationOverridesItemsRepresentation = map[string]interface{}{
		"config_key":             acctest.Representation{RepType: acctest.Required, Create: `effective_io_concurrency`},
		"overriden_config_value": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	PsqlConfigurationResourceDependencies = AvailabilityDomainConfig + DefinedTagsDependencies
)

// issue-routing-tag: psql/default
func TestPsqlConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_psql_configuration.test_configuration"
	datasourceName := "data.oci_psql_configurations.test_configurations"
	singularDatasourceName := "data.oci_psql_configuration.test_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PsqlConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Create, PsqlConfigurationRepresentation), "psql", "configuration", t)

	acctest.ResourceTest(t, testAccCheckPsqlConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Required, acctest.Create, PsqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.config_key", "effective_io_concurrency"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.overriden_config_value", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraform-test-config"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "64"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// // delete before next Create
		{
			Config: config + compartmentIdVariableStr + PsqlConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Create, PsqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.config_key", "effective_io_concurrency"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.overriden_config_value", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraform-test-config"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "64"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlConfigurationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.config_key", "effective_io_concurrency"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.overriden_config_value", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraform-test-config"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "64"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Update, PsqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.config_key", "effective_io_concurrency"),
				resource.TestCheckResourceAttr(resourceName, "db_configuration_overrides.0.items.0.overriden_config_value", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraform-test-config-2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "instance_memory_size_in_gbs", "64"),
				resource.TestCheckResourceAttr(resourceName, "instance_ocpu_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_configurations", "test_configurations", acctest.Optional, acctest.Update, PsqlConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Update, PsqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "shape", "VM.Standard.E4.Flex.4.64GB"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "configuration_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.display_name", "terraform-test-config-2"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.shape", "VM.Standard.E4.Flex.4.64GB"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.#", "1"),
			),
		},
		// verify datasource with display name
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_configurations", "test_configurations", acctest.Optional, acctest.Update, PsqlConfigurationDisplayNameDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Update, PsqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "terraform-test-config-2"),

				resource.TestCheckResourceAttr(datasourceName, "display_name", "terraform-test-config-2"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.display_name", "terraform-test-config-2"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.shape", "VM.Standard.E4.Flex.4.64GB"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.#", "1"),
			),
		},
		// verify datasource with configuration id
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_configurations", "test_configurations", acctest.Optional, acctest.Update, PsqlConfigurationIDDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Optional, acctest.Update, PsqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "configuration_id"),

				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.display_name", "terraform-test-config-2"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.shape", "VM.Standard.E4.Flex.4.64GB"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "configuration_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_configuration", "test_configuration", acctest.Required, acctest.Create, PsqlConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "terraform-test-config-2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_memory_size_in_gbs", "64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_ocpu_count", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "VM.Standard.E4.Flex.4.64GB"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + PsqlConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"db_configuration_overrides",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckPsqlConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PostgresqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_psql_configuration" {
			noResourceFound = false
			request := oci_psql.GetConfigurationRequest{}

			tmp := rs.Primary.ID
			request.ConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")

			response, err := client.GetConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_psql.ConfigurationLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("PsqlConfiguration") {
		resource.AddTestSweepers("PsqlConfiguration", &resource.Sweeper{
			Name:         "PsqlConfiguration",
			Dependencies: acctest.DependencyGraph["configuration"],
			F:            sweepPsqlConfigurationResource,
		})
	}
}

func sweepPsqlConfigurationResource(compartment string) error {
	postgresqlClient := acctest.GetTestClients(&schema.ResourceData{}).PostgresqlClient()
	configurationIds, err := getPsqlConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, configurationId := range configurationIds {
		if ok := acctest.SweeperDefaultResourceId[configurationId]; !ok {
			deleteConfigurationRequest := oci_psql.DeleteConfigurationRequest{}

			deleteConfigurationRequest.ConfigurationId = &configurationId

			deleteConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")
			_, error := postgresqlClient.DeleteConfiguration(context.Background(), deleteConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting Configuration %s %s, It is possible that the resource is already deleted. Please verify manually \n", configurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &configurationId, PsqlConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				PsqlConfigurationSweepResponseFetchOperation, "psql", true)
		}
	}
	return nil
}

func getPsqlConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	postgresqlClient := acctest.GetTestClients(&schema.ResourceData{}).PostgresqlClient()

	listConfigurationsRequest := oci_psql.ListConfigurationsRequest{}
	listConfigurationsRequest.CompartmentId = &compartmentId
	listConfigurationsRequest.LifecycleState = oci_psql.ConfigurationLifecycleStateActive
	listConfigurationsResponse, err := postgresqlClient.ListConfigurations(context.Background(), listConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Configuration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, configuration := range listConfigurationsResponse.Items {
		id := *configuration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConfigurationId", id)
	}
	return resourceIds, nil
}

func PsqlConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if configurationResponse, ok := response.Response.(oci_psql.GetConfigurationResponse); ok {
		return configurationResponse.LifecycleState != oci_psql.ConfigurationLifecycleStateDeleted
	}
	return false
}

func PsqlConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PostgresqlClient().GetConfiguration(context.Background(), oci_psql.GetConfigurationRequest{
		ConfigurationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
