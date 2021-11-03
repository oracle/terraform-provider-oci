// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v54/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagementAgentInstallKeyRequiredOnlyResource = ManagementAgentInstallKeyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Required, Create, managementAgentInstallKeyRepresentation)

	ManagementAgentInstallKeyResourceConfig = ManagementAgentInstallKeyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Optional, Update, managementAgentInstallKeyRepresentation)

	managementAgentInstallKeySingularDataSourceRepresentation = map[string]interface{}{
		"management_agent_install_key_id": Representation{RepType: Required, Create: `${oci_management_agent_management_agent_install_key.test_management_agent_install_key.id}`},
	}

	managementAgentInstallKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"access_level":              Representation{RepType: Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`},
		"display_name":              Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":                    RepresentationGroup{Required, managementAgentInstallKeyDataSourceFilterRepresentation}}
	managementAgentInstallKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_management_agent_management_agent_install_key.test_management_agent_install_key.id}`}},
	}

	expirationTimeForManagementAgentInstallKey = time.Now().UTC().AddDate(0, 0, 7).Truncate(time.Millisecond)

	managementAgentInstallKeyRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":              Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"allowed_key_install_count": Representation{RepType: Optional, Create: `10`},
		"time_expires":              Representation{RepType: Optional, Create: expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)},
	}

	ManagementAgentInstallKeyResourceDependencies = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentInstallKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentInstallKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_management_agent_management_agent_install_key.test_management_agent_install_key"
	datasourceName := "data.oci_management_agent_management_agent_install_keys.test_management_agent_install_keys"
	singularDatasourceName := "data.oci_management_agent_management_agent_install_key.test_management_agent_install_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ManagementAgentInstallKeyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Optional, Create, managementAgentInstallKeyRepresentation), "managementagent", "managementAgentInstallKey", t)

	ResourceTest(t, testAccCheckManagementAgentManagementAgentInstallKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Required, Create, managementAgentInstallKeyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Optional, Create, managementAgentInstallKeyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Optional, Update, managementAgentInstallKeyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_install_keys", "test_management_agent_install_keys", Optional, Update, managementAgentInstallKeyDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentInstallKeyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Optional, Update, managementAgentInstallKeyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_install_keys.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_install_keys.0.time_expires", expirationTimeForManagementAgentInstallKey.Format(time.RFC3339Nano)),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_install_key", "test_management_agent_install_key", Required, Create, managementAgentInstallKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentInstallKeyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_install_key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_key_install_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by_principal_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_key_install_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ManagementAgentInstallKeyResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagementAgentManagementAgentInstallKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).managementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_management_agent_install_key" {
			noResourceFound = false
			request := oci_management_agent.GetManagementAgentInstallKeyRequest{}

			tmp := rs.Primary.ID
			request.ManagementAgentInstallKeyId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "management_agent")

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
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("ManagementAgentManagementAgentInstallKey") {
		resource.AddTestSweepers("ManagementAgentManagementAgentInstallKey", &resource.Sweeper{
			Name:         "ManagementAgentManagementAgentInstallKey",
			Dependencies: DependencyGraph["managementAgentInstallKey"],
			F:            sweepManagementAgentManagementAgentInstallKeyResource,
		})
	}
}

func sweepManagementAgentManagementAgentInstallKeyResource(compartment string) error {
	managementAgentClient := GetTestClients(&schema.ResourceData{}).managementAgentClient()
	managementAgentInstallKeyIds, err := getManagementAgentInstallKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, managementAgentInstallKeyId := range managementAgentInstallKeyIds {
		if ok := SweeperDefaultResourceId[managementAgentInstallKeyId]; !ok {
			deleteManagementAgentInstallKeyRequest := oci_management_agent.DeleteManagementAgentInstallKeyRequest{}

			deleteManagementAgentInstallKeyRequest.ManagementAgentInstallKeyId = &managementAgentInstallKeyId

			deleteManagementAgentInstallKeyRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteManagementAgentInstallKey(context.Background(), deleteManagementAgentInstallKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAgentInstallKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementAgentInstallKeyId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &managementAgentInstallKeyId, managementAgentInstallKeySweepWaitCondition, time.Duration(3*time.Minute),
				managementAgentInstallKeySweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}

func getManagementAgentInstallKeyIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ManagementAgentInstallKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementAgentClient := GetTestClients(&schema.ResourceData{}).managementAgentClient()

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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentInstallKeyId", id)
	}
	return resourceIds, nil
}

func managementAgentInstallKeySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementAgentInstallKeyResponse, ok := response.Response.(oci_management_agent.GetManagementAgentInstallKeyResponse); ok {
		return managementAgentInstallKeyResponse.LifecycleState != oci_management_agent.LifecycleStatesTerminated
	}
	return false
}

func managementAgentInstallKeySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.managementAgentClient().GetManagementAgentInstallKey(context.Background(), oci_management_agent.GetManagementAgentInstallKeyRequest{
		ManagementAgentInstallKeyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
