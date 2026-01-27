// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpManagementApplianceRequiredOnlyResource = OcvpManagementApplianceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Required, acctest.Create, OcvpManagementApplianceRepresentation)

	OcvpManagementApplianceResourceConfig = OcvpManagementApplianceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Optional, acctest.Update, OcvpManagementApplianceRepresentation)

	OcvpManagementApplianceSingularDataSourceRepresentation = map[string]interface{}{
		"management_appliance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_management_appliance.test_management_appliance.id}`},
	}

	OcvpManagementApplianceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"management_appliance_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_management_appliance.test_management_appliance.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `NEEDS_ATTENTION`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpManagementApplianceDataSourceFilterRepresentation},
	}
	OcvpManagementApplianceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_management_appliance.test_management_appliance.id}`}},
	}

	OcvpManagementApplianceRepresentation = map[string]interface{}{
		"configuration":   acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpManagementApplianceConfigurationRepresentation},
		"connections":     acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpManagementApplianceConnectionsRepresentation},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"sddc_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"public_ssh_keys": acctest.Representation{RepType: acctest.Optional, Create: `${var.public_ssh_keys}`},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}
	OcvpManagementApplianceConfigurationRepresentation = map[string]interface{}{
		"is_log_ingestion_enabled":      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_metrics_collection_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"metrics":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`metrics`}, Update: []string{`metrics2`}},
		"support_bundle_bucket_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_bucket.test_bucket.bucket_id}`},
	}
	OcvpManagementApplianceConnectionsRepresentation = map[string]interface{}{
		"credentials_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.credentials_secret_id}`},
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `VCENTER`, Update: `ADMIN_VCENTER`},
	}

	OcvpManagementApplianceResourceDependencies = OcvpSddcResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Create, OcvpSddcRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(VaultVaultSecretSingularDataSourceRepresentation, map[string]interface{}{
				"secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.credentials_secret_id}`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectStorageBucketRepresentation, map[string]interface{}{
			"namespace": acctest.Representation{RepType: acctest.Required, Create: `${var.object_storage_namespace}`},
		}))
)

// issue-routing-tag: ocvp/default
func TestOcvpManagementApplianceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpManagementApplianceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	credentialsSecretId := utils.GetEnvSettingWithBlankDefault("credentials_secret_id")
	credentialsSecretIdVariableStr := fmt.Sprintf("variable \"credentials_secret_id\" { default = \"%s\" }\n", credentialsSecretId)

	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	getNamespacesReq := oci_object_storage.GetNamespaceRequest{
		CompartmentId: &compartmentId,
	}
	getNamespacesResp, _ := objectStorageClient.GetNamespace(context.Background(), getNamespacesReq)
	namespaceVariableStr := fmt.Sprintf("variable \"object_storage_namespace\" { default = \"%s\" }\n", *getNamespacesResp.Value)

	publicSshKeys := utils.GetEnvSettingWithBlankDefault("public_ssh_keys")
	publicSshKeysVariableStr := fmt.Sprintf("variable \"public_ssh_keys\" { default = \"%s\" }\n", publicSshKeys)

	resourceName := "oci_ocvp_management_appliance.test_management_appliance"
	datasourceName := "data.oci_ocvp_management_appliances.test_management_appliances"
	singularDatasourceName := "data.oci_ocvp_management_appliance.test_management_appliance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+namespaceVariableStr+credentialsSecretId+publicSshKeysVariableStr+OcvpManagementApplianceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Optional, acctest.Create, OcvpManagementApplianceRepresentation), "ocvp", "managementAppliance", t)

	acctest.ResourceTest(t, testAccCheckOcvpManagementApplianceDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + namespaceVariableStr + credentialsSecretIdVariableStr + publicSshKeysVariableStr + OcvpManagementApplianceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Required, acctest.Create, OcvpManagementApplianceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_log_ingestion_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_metrics_collection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "connections.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "connections.0.credentials_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "connections.0.type", "VCENTER"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + namespaceVariableStr + credentialsSecretIdVariableStr + publicSshKeysVariableStr + OcvpManagementApplianceResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + namespaceVariableStr + credentialsSecretIdVariableStr + publicSshKeysVariableStr + OcvpManagementApplianceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Optional, acctest.Create, OcvpManagementApplianceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_log_ingestion_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_metrics_collection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.metrics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "configuration.0.support_bundle_bucket_id"),
				resource.TestCheckResourceAttr(resourceName, "connections.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "connections.0.credentials_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "connections.0.type", "VCENTER"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + namespaceVariableStr + credentialsSecretIdVariableStr + publicSshKeysVariableStr + OcvpManagementApplianceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Optional, acctest.Update, OcvpManagementApplianceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_log_ingestion_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_metrics_collection_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.metrics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "configuration.0.support_bundle_bucket_id"),
				resource.TestCheckResourceAttr(resourceName, "connections.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "connections.0.credentials_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "connections.0.type", "ADMIN_VCENTER"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_management_appliances", "test_management_appliances", acctest.Optional, acctest.Update, OcvpManagementApplianceDataSourceRepresentation) +
				compartmentIdVariableStr + namespaceVariableStr + credentialsSecretIdVariableStr + publicSshKeysVariableStr + OcvpManagementApplianceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Optional, acctest.Update, OcvpManagementApplianceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_appliance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NEEDS_ATTENTION"),

				resource.TestCheckResourceAttr(datasourceName, "management_appliance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "management_appliance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_management_appliance", "test_management_appliance", acctest.Required, acctest.Create, OcvpManagementApplianceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + namespaceVariableStr + credentialsSecretIdVariableStr + publicSshKeysVariableStr + OcvpManagementApplianceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_appliance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_log_ingestion_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_metrics_collection_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.metrics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connections.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connections.0.type", "ADMIN_VCENTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "heartbeat_connection_states.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_configuration_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + OcvpManagementApplianceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"public_ssh_keys",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOcvpManagementApplianceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementApplianceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_management_appliance" {
			noResourceFound = false
			request := oci_ocvp.GetManagementApplianceRequest{}

			tmp := rs.Primary.ID
			request.ManagementApplianceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetManagementAppliance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.ManagementApplianceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpManagementAppliance") {
		resource.AddTestSweepers("OcvpManagementAppliance", &resource.Sweeper{
			Name:         "OcvpManagementAppliance",
			Dependencies: acctest.DependencyGraph["managementAppliance"],
			F:            sweepOcvpManagementApplianceResource,
		})
	}
}

func sweepOcvpManagementApplianceResource(compartment string) error {
	managementApplianceClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementApplianceClient()
	managementApplianceIds, err := getOcvpManagementApplianceIds(compartment)
	if err != nil {
		return err
	}
	for _, managementApplianceId := range managementApplianceIds {
		if ok := acctest.SweeperDefaultResourceId[managementApplianceId]; !ok {
			deleteManagementApplianceRequest := oci_ocvp.DeleteManagementApplianceRequest{}

			deleteManagementApplianceRequest.ManagementApplianceId = &managementApplianceId

			deleteManagementApplianceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := managementApplianceClient.DeleteManagementAppliance(context.Background(), deleteManagementApplianceRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAppliance %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementApplianceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managementApplianceId, OcvpManagementApplianceSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpManagementApplianceSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpManagementApplianceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementApplianceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementApplianceClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementApplianceClient()
	activeLifecycleStates := []oci_ocvp.ListManagementAppliancesLifecycleStateEnum{
		oci_ocvp.ListManagementAppliancesLifecycleStateActive,
		oci_ocvp.ListManagementAppliancesLifecycleStateNeedsAttention,
	}
	for _, lifecycleState := range activeLifecycleStates {
		fetchedIds, err := getOcvpManagementApplianceIdsForLifecycleState(*managementApplianceClient, compartmentId, lifecycleState)
		if err != nil {
			return resourceIds, fmt.Errorf("Error getting %s ManagementAppliance list for compartment id : %s , %s \n", lifecycleState, compartmentId, err)
		}
		resourceIds = append(resourceIds, fetchedIds...)
	}

	for _, id := range resourceIds {
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementApplianceId", id)
	}
	return resourceIds, nil
}

func getOcvpManagementApplianceIdsForLifecycleState(
	client oci_ocvp.ManagementApplianceClient,
	compartmentId string,
	lifecycleState oci_ocvp.ListManagementAppliancesLifecycleStateEnum) ([]string, error) {

	var resourceIds []string

	listManagementAppliancesRequest := oci_ocvp.ListManagementAppliancesRequest{}
	listManagementAppliancesRequest.CompartmentId = &compartmentId
	listManagementAppliancesRequest.LifecycleState = lifecycleState
	listManagementAppliancesResponse, err := client.ListManagementAppliances(context.Background(), listManagementAppliancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAppliance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAppliance := range listManagementAppliancesResponse.Items {
		id := *managementAppliance.Id
		resourceIds = append(resourceIds, id)
	}
	return resourceIds, nil
}

func OcvpManagementApplianceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementApplianceResponse, ok := response.Response.(oci_ocvp.GetManagementApplianceResponse); ok {
		return managementApplianceResponse.LifecycleState != oci_ocvp.ManagementApplianceLifecycleStateDeleted
	}
	return false
}

func OcvpManagementApplianceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementApplianceClient().GetManagementAppliance(context.Background(), oci_ocvp.GetManagementApplianceRequest{
		ManagementApplianceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
