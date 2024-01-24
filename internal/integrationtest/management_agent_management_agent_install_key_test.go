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
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagementAgentInstallKeyRequiredOnlyResource = ManagementAgentInstallKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Required, acctest.Create, managementAgentInstallKeyRepresentation)

	ManagementAgentInstallKeyResourceConfig = ManagementAgentInstallKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Optional, acctest.Update, managementAgentInstallKeyRepresentation)

	managementAgentInstallKeySingularDataSourceRepresentation = map[string]interface{}{
		"management_agent_install_key_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_management_agent_management_agent_install_key.test_management_agent_install_key.id}`},
	}

	managementAgentInstallKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: managementAgentInstallKeyDataSourceFilterRepresentation}}
	managementAgentInstallKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_management_agent_management_agent_install_key.test_management_agent_install_key.id}`}},
	}

	expirationTimeForManagementAgentInstallKey = time.Now().UTC().AddDate(0, 0, 7).Truncate(time.Millisecond)

	managementAgentInstallKeyRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"allowed_key_install_count": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_unlimited":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"time_expires":              acctest.Representation{RepType: acctest.Optional, Create: expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)},
	}

	ManagementAgentInstallKeyResourceDependencies = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentInstallKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentInstallKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_management_agent_management_agent_install_key.test_management_agent_install_key"
	datasourceName := "data.oci_management_agent_management_agent_install_keys.test_management_agent_install_keys"
	singularDatasourceName := "data.oci_management_agent_management_agent_install_key.test_management_agent_install_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagementAgentInstallKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Optional, acctest.Create, managementAgentInstallKeyRepresentation), "managementagent", "managementAgentInstallKey", t)

	acctest.ResourceTest(t, testAccCheckManagementAgentManagementAgentInstallKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Required, acctest.Create, managementAgentInstallKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Optional, acctest.Create, managementAgentInstallKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)),

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
			Config: config + compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Optional, acctest.Update, managementAgentInstallKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_install_keys", "test_management_agent_install_keys", acctest.Optional, acctest.Update, managementAgentInstallKeyDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Optional, acctest.Update, managementAgentInstallKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.0.allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.created_by_principal_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.current_key_install_count"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.0.is_unlimited", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.0.time_expires", expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", acctest.Required, acctest.Create, managementAgentInstallKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentInstallKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_install_key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by_principal_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_install_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ManagementAgentInstallKeyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagementAgentManagementAgentInstallKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_management_agent_install_key" {
			noResourceFound = false
			request := oci_management_agent.GetManagementAgentInstallKeyRequest{}

			tmp := rs.Primary.ID
			request.ManagementAgentInstallKeyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")

			response, err := client.GetManagementAgentInstallKey(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_management_agent.LifecycleStatesTerminated): true, string(oci_management_agent.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("ManagementAgentManagementAgentInstallKey") {
		resource.AddTestSweepers("ManagementAgentManagementAgentInstallKey", &resource.Sweeper{
			Name:         "ManagementAgentManagementAgentInstallKey",
			Dependencies: acctest.DependencyGraph["managementAgentInstallKey"],
			F:            sweepManagementAgentManagementAgentInstallKeyResource,
		})
	}
}

func sweepManagementAgentManagementAgentInstallKeyResource(compartment string) error {
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()
	managementAgentInstallKeyIds, err := getManagementAgentInstallKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, managementAgentInstallKeyId := range managementAgentInstallKeyIds {
		if ok := acctest.SweeperDefaultResourceId[managementAgentInstallKeyId]; !ok {
			deleteManagementAgentInstallKeyRequest := oci_management_agent.DeleteManagementAgentInstallKeyRequest{}

			deleteManagementAgentInstallKeyRequest.ManagementAgentInstallKeyId = &managementAgentInstallKeyId

			deleteManagementAgentInstallKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteManagementAgentInstallKey(context.Background(), deleteManagementAgentInstallKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAgentInstallKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementAgentInstallKeyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managementAgentInstallKeyId, managementAgentInstallKeySweepWaitCondition, time.Duration(3*time.Minute),
				managementAgentInstallKeySweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}

func getManagementAgentInstallKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementAgentInstallKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listManagementAgentInstallKeysRequest := oci_management_agent.ListManagementAgentInstallKeysRequest{}
	listManagementAgentInstallKeysRequest.CompartmentId = &compartmentId
	listManagementAgentInstallKeysRequest.LifecycleState = oci_management_agent.ListManagementAgentInstallKeysLifecycleStateActive
	listManagementAgentInstallKeysResponse, err := managementAgentClient.ListManagementAgentInstallKeys(context.Background(), listManagementAgentInstallKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAgentInstallKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAgentInstallKey := range listManagementAgentInstallKeysResponse.Items {
		id := *managementAgentInstallKey.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentInstallKeyId", id)
	}
	return resourceIds, nil
}

func managementAgentInstallKeySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementAgentInstallKeyResponse, ok := response.Response.(oci_management_agent.GetManagementAgentInstallKeyResponse); ok {
		return managementAgentInstallKeyResponse.LifecycleState != oci_management_agent.LifecycleStatesTerminated && managementAgentInstallKeyResponse.LifecycleState != oci_management_agent.LifecycleStatesDeleted
	}
	return false
}

func managementAgentInstallKeySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementAgentClient().GetManagementAgentInstallKey(context.Background(), oci_management_agent.GetManagementAgentInstallKeyRequest{
		ManagementAgentInstallKeyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
