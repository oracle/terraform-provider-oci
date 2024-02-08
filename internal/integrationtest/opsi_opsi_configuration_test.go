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
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiOpsiConfigurationRequiredOnlyResource = OpsiOpsiConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Required, acctest.Create, OpsiOpsiConfigurationRepresentation)

	OpsiOpsiConfigurationResourceConfig = OpsiOpsiConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Update, OpsiOpsiConfigurationRepresentation)

	OpsiOpsiOpsiConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"opsi_configuration_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_opsi_configuration.test_opsi_configuration.id}`},
		"config_item_custom_status":       acctest.Representation{RepType: acctest.Optional, Create: []string{`customized`}},
		"config_item_field":               acctest.Representation{RepType: acctest.Optional, Create: []string{`metadata`, `name`, `value`}},
		"config_items_applicable_context": acctest.Representation{RepType: acctest.Optional, Create: []string{`DB_CAPACITY_PLANNING`}},
		"opsi_config_field":               acctest.Representation{RepType: acctest.Optional, Create: []string{`configItems`}},
	}

	OpsiOpsiOpsiConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"opsi_config_type": acctest.Representation{RepType: acctest.Required, Create: []string{`UX_CONFIGURATION`}, Update: []string{`UX_CONFIGURATION`}},
		"state":            acctest.Representation{RepType: acctest.Required, Create: []string{`ACTIVE`}, Update: []string{`ACTIVE`}},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiOpsiConfigurationDataSourceFilterRepresentation}}

	OpsiOpsiConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_opsi_configuration.test_opsi_configuration.id}`}},
	}

	OpsiOpsiConfigurationRepresentation = map[string]interface{}{
		"opsi_config_type":                acctest.Representation{RepType: acctest.Required, Create: `UX_CONFIGURATION`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_item_custom_status":       acctest.Representation{RepType: acctest.Optional, Create: []string{`customized`}},
		"config_item_field":               acctest.Representation{RepType: acctest.Optional, Create: []string{`metadata`, `name`, `value`}},
		"config_items":                    []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: OpsiOpsiConfigurationConfigItemsRepresentation1}, {RepType: acctest.Optional, Group: OpsiOpsiConfigurationConfigItemsRepresentation2}},
		"config_items_applicable_context": acctest.Representation{RepType: acctest.Optional, Create: []string{`DB_CAPACITY_PLANNING`}},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"opsi_config_field":               acctest.Representation{RepType: acctest.Optional, Create: []string{`configItems`}},
		//"system_tags":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"systemTags": "value"}, Update: map[string]string{"systemTags": "updatedValue"}},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiConfigIgnoreChangesRep},
	}

	OpsiConfigIgnoreChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`, `config_items`}},
	}

	OpsiOpsiConfigurationConfigItemsRepresentation1 = map[string]interface{}{
		"config_item_type": acctest.Representation{RepType: acctest.Required, Create: `BASIC`, Update: `BASIC`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `dbHighCpuThreshold`, Update: `dbHighCpuThreshold`},
		"value":            acctest.Representation{RepType: acctest.Required, Create: `88`, Update: `85`},
	}

	OpsiOpsiConfigurationConfigItemsRepresentation2 = map[string]interface{}{
		"config_item_type": acctest.Representation{RepType: acctest.Optional, Create: `BASIC`, Update: `BASIC`},
		"name":             acctest.Representation{RepType: acctest.Optional, Create: `dbHighMemoryThreshold`, Update: `dbHighMemoryThreshold`},
		"value":            acctest.Representation{RepType: acctest.Optional, Create: `80`, Update: `75`},
	}

	OpsiOpsiConfigurationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/default
func TestOpsiOpsiConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOpsiConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_opsi_configuration.test_opsi_configuration"
	datasourceName := "data.oci_opsi_opsi_configurations.test_opsi_configurations"
	singularDatasourceName := "data.oci_opsi_opsi_configuration.test_opsi_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpsiOpsiConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Create, OpsiOpsiConfigurationRepresentation), "opsi", "opsiConfiguration", t)

	acctest.ResourceTest(t, testAccCheckOpsiOpsiConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OpsiOpsiConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Required, acctest.Create, OpsiOpsiConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "opsi_config_type", "UX_CONFIGURATION"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiOpsiConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OpsiOpsiConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Create, OpsiOpsiConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_item_custom_status.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_item_field.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "config_items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "config_items.0.config_item_type", "BASIC"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "config_items_applicable_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "opsi_config_field.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "opsi_config_type", "UX_CONFIGURATION"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OpsiOpsiConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiOpsiConfigurationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "config_item_custom_status.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_item_field.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "config_items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "config_items.0.config_item_type", "BASIC"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "config_items_applicable_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "opsi_config_field.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "opsi_config_type", "UX_CONFIGURATION"),
				//resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),

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
			Config: config + compartmentIdVariableStr + OpsiOpsiConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Update, OpsiOpsiConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_item_custom_status.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_item_field.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "config_items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "config_items.0.config_item_type", "BASIC"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.0.name", "name2"),
				//resource.TestCheckResourceAttr(resourceName, "config_items.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "config_items_applicable_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "opsi_config_field.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "opsi_config_type", "UX_CONFIGURATION"),
				//resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_opsi_configurations", "test_opsi_configurations", acctest.Required, acctest.Update, OpsiOpsiOpsiConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiOpsiConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Update, OpsiOpsiConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "opsi_config_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "opsi_configurations_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "opsi_configurations_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_opsi_configuration", "test_opsi_configuration", acctest.Optional, acctest.Create, OpsiOpsiOpsiConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiOpsiConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "config_item_custom_status.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_item_field.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_items_applicable_context.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "opsi_config_field.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opsi_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_items.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_items.0.config_item_type", "BASIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_items.0.metadata.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "config_items.0.name", "name2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "config_items.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "opsi_config_type", "UX_CONFIGURATION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + OpsiOpsiConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"config_items",
				"config_item_custom_status",
				"config_item_field",
				"opsi_config_field",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpsiOpsiConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_opsi_configuration" {
			noResourceFound = false
			request := oci_opsi.GetOpsiConfigurationRequest{}

			tmp := rs.Primary.ID
			request.OpsiConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetOpsiConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.OpsiConfigurationLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("OpsiOpsiConfiguration") {
		resource.AddTestSweepers("OpsiOpsiConfiguration", &resource.Sweeper{
			Name:         "OpsiOpsiConfiguration",
			Dependencies: acctest.DependencyGraph["opsiConfiguration"],
			F:            sweepOpsiOpsiConfigurationResource,
		})
	}
}

func sweepOpsiOpsiConfigurationResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	opsiConfigurationIds, err := getOpsiOpsiConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, opsiConfigurationId := range opsiConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[opsiConfigurationId]; !ok {
			deleteOpsiConfigurationRequest := oci_opsi.DeleteOpsiConfigurationRequest{}

			deleteOpsiConfigurationRequest.OpsiConfigurationId = &opsiConfigurationId

			deleteOpsiConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteOpsiConfiguration(context.Background(), deleteOpsiConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting OpsiConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", opsiConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &opsiConfigurationId, OpsiOpsiConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiOpsiConfigurationSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiOpsiConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OpsiConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listOpsiConfigurationsRequest := oci_opsi.ListOpsiConfigurationsRequest{}
	listOpsiConfigurationsRequest.CompartmentId = &compartmentId
	listOpsiConfigurationsRequest.LifecycleState = []oci_opsi.OpsiConfigurationLifecycleStateEnum{oci_opsi.OpsiConfigurationLifecycleStateActive}
	listOpsiConfigurationsResponse, err := operationsInsightsClient.ListOpsiConfigurations(context.Background(), listOpsiConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OpsiConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, opsiConfiguration := range listOpsiConfigurationsResponse.Items {
		id := *opsiConfiguration.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OpsiConfigurationId", id)
	}
	return resourceIds, nil
}

func OpsiOpsiConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if opsiConfigurationResponse, ok := response.Response.(oci_opsi.GetOpsiConfigurationResponse); ok {
		return opsiConfigurationResponse.GetLifecycleState() != oci_opsi.OpsiConfigurationLifecycleStateDeleted
	}
	return false
}

func OpsiOpsiConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetOpsiConfiguration(context.Background(), oci_opsi.GetOpsiConfigurationRequest{
		OpsiConfigurationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
