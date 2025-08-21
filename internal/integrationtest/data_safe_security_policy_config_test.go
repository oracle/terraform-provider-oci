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
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSecurityPolicyConfigRequiredOnlyResource = DataSafeSecurityPolicyConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Required, acctest.Create, DataSafeSecurityPolicyConfigCreateRepresentation)

	DataSafeSecurityPolicyConfigResourceConfig = DataSafeSecurityPolicyConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Optional, acctest.Update, DataSafeSecurityPolicyConfigUpdateRepresentation)

	DataSafeSecurityPolicyConfigSingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_policy_config.test_security_policy_config.id}`},
	}

	DataSafeSecurityPolicyConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"security_policy_config_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_security_policy_config.test_security_policy_config.id}`},
		"security_policy_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.security_policy_id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DataSafeSecurityPolicyConfigCreateRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"unified_audit_policy_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSecurityPolicyConfigUnifiedAuditPolicyConfigRepresentation},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSecurityPolicyConfigSystemTagsChangesRep},
	}

	DataSafeSecurityPolicyConfigUpdateRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"Department": "Accounting"}},
		"firewall_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeSecurityPolicyConfigFirewallConfigRepresentation},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSecurityPolicyConfigSystemTagsChangesRep},
	}

	DataSafeSecurityPolicyConfigFirewallConfigRepresentation = map[string]interface{}{
		"exclude_job":              acctest.Representation{RepType: acctest.Optional, Update: `INCLUDED`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Update: `ENABLED`},
		"violation_log_auto_purge": acctest.Representation{RepType: acctest.Optional, Update: `ENABLED`},
	}
	DataSafeSecurityPolicyConfigUnifiedAuditPolicyConfigRepresentation = map[string]interface{}{
		"exclude_datasafe_user": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
	}

	IgnoreSecurityPolicyConfigSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyConfigResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyConfigResource_basic(t *testing.T) {
	t.Skip("Skipping this test as security policy ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyId := utils.GetEnvSettingWithBlankDefault("security_policy_ocid")
	securityPolicyIdVariableStr := fmt.Sprintf("variable \"security_policy_id\" { default = \"%s\" }\n", securityPolicyId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_security_policy_config.test_security_policy_config"
	datasourceName := "data.oci_data_safe_security_policy_configs.test_security_policy_configs"
	singularDatasourceName := "data.oci_data_safe_security_policy_config.test_security_policy_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSecurityPolicyConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Optional, acctest.Create, DataSafeSecurityPolicyConfigCreateRepresentation), "datasafe", "securityPolicyConfig", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSecurityPolicyConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Required, acctest.Create, DataSafeSecurityPolicyConfigCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Optional, acctest.Create, DataSafeSecurityPolicyConfigCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "unified_audit_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "unified_audit_policy_config.0.exclude_datasafe_user", "ENABLED"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyConfigCreateRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "unified_audit_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "unified_audit_policy_config.0.exclude_datasafe_user", "ENABLED"),

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
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Optional, acctest.Update, DataSafeSecurityPolicyConfigUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "description"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "firewall_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "firewall_config.0.status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "unified_audit_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "unified_audit_policy_config.0.exclude_datasafe_user", "ENABLED"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_configs", "test_security_policy_configs", acctest.Optional, acctest.Update, DataSafeSecurityPolicyConfigDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Optional, acctest.Update, DataSafeSecurityPolicyConfigUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_config_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "security_policy_config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_policy_config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_config", "test_security_policy_config", acctest.Required, acctest.Create, DataSafeSecurityPolicyConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_config_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "firewall_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "firewall_config.0.status", "ENABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "unified_audit_policy_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "unified_audit_policy_config.0.exclude_datasafe_user", "ENABLED"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSecurityPolicyConfigRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSecurityPolicyConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_security_policy_config" {
			noResourceFound = false
			request := oci_data_safe.GetSecurityPolicyConfigRequest{}

			tmp := rs.Primary.ID
			request.SecurityPolicyConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSecurityPolicyConfig(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SecurityPolicyConfigLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSecurityPolicyConfig") {
		resource.AddTestSweepers("DataSafeSecurityPolicyConfig", &resource.Sweeper{
			Name:         "DataSafeSecurityPolicyConfig",
			Dependencies: acctest.DependencyGraph["securityPolicyConfig"],
			F:            sweepDataSafeSecurityPolicyConfigResource,
		})
	}
}

func sweepDataSafeSecurityPolicyConfigResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	securityPolicyConfigIds, err := getDataSafeSecurityPolicyConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, securityPolicyConfigId := range securityPolicyConfigIds {
		if ok := acctest.SweeperDefaultResourceId[securityPolicyConfigId]; !ok {
			deleteSecurityPolicyConfigRequest := oci_data_safe.DeleteSecurityPolicyConfigRequest{}

			deleteSecurityPolicyConfigRequest.SecurityPolicyConfigId = &securityPolicyConfigId

			deleteSecurityPolicyConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSecurityPolicyConfig(context.Background(), deleteSecurityPolicyConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityPolicyConfig %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityPolicyConfigId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityPolicyConfigId, DataSafeSecurityPolicyConfigSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSecurityPolicyConfigSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSecurityPolicyConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityPolicyConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSecurityPolicyConfigsRequest := oci_data_safe.ListSecurityPolicyConfigsRequest{}
	listSecurityPolicyConfigsRequest.CompartmentId = &compartmentId
	listSecurityPolicyConfigsRequest.LifecycleState = oci_data_safe.ListSecurityPolicyConfigsLifecycleStateActive
	listSecurityPolicyConfigsResponse, err := dataSafeClient.ListSecurityPolicyConfigs(context.Background(), listSecurityPolicyConfigsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityPolicyConfig list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityPolicyConfig := range listSecurityPolicyConfigsResponse.Items {
		id := *securityPolicyConfig.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityPolicyConfigId", id)
	}
	return resourceIds, nil
}

func DataSafeSecurityPolicyConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityPolicyConfigResponse, ok := response.Response.(oci_data_safe.GetSecurityPolicyConfigResponse); ok {
		return securityPolicyConfigResponse.LifecycleState != oci_data_safe.SecurityPolicyConfigLifecycleStateDeleted
	}
	return false
}

func DataSafeSecurityPolicyConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSecurityPolicyConfig(context.Background(), oci_data_safe.GetSecurityPolicyConfigRequest{
		SecurityPolicyConfigId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
