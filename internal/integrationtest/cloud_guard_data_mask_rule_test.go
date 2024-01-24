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
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudGuardDataMaskRuleRequiredOnlyResource = CloudGuardDataMaskRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Required, acctest.Create, CloudGuardDataMaskRuleRepresentation)

	CloudGuardDataMaskRuleResourceConfig = CloudGuardDataMaskRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Optional, acctest.Update, CloudGuardDataMaskRuleRepresentation)

	CloudGuardCloudGuardDataMaskRuleSingularDataSourceRepresentation = map[string]interface{}{
		"data_mask_rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_data_mask_rule.test_data_mask_rule.id}`},
	}

	CloudGuardCloudGuardDataMaskRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"access_level":          acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"data_mask_rule_status": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"iam_group_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"target_type":           acctest.Representation{RepType: acctest.Optional, Create: `targetType`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardDataMaskRuleDataSourceFilterRepresentation}}
	CloudGuardDataMaskRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_data_mask_rule.test_data_mask_rule.id}`}},
	}

	CloudGuardDataMaskRuleRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"data_mask_categories":  acctest.Representation{RepType: acctest.Required, Create: []string{`PII`}, Update: []string{`PHI`}},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"iam_group_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_group.test_group.id}`},
		"target_selected":       acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardDataMaskRuleTargetSelectedRepresentation},
		"data_mask_rule_status": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	CloudGuardDataMaskRuleTargetSelectedRepresentation = map[string]interface{}{
		"kind":   acctest.Representation{RepType: acctest.Required, Create: `ALL`, Update: `ALL`},
		"values": acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{}},
	}

	CloudGuardDataMaskRuleResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, IdentityGroupRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardDataMaskRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardDataMaskRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_cloud_guard_data_mask_rule.test_data_mask_rule"
	datasourceName := "data.oci_cloud_guard_data_mask_rules.test_data_mask_rules"
	singularDatasourceName := "data.oci_cloud_guard_data_mask_rule.test_data_mask_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardDataMaskRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Optional, acctest.Create, CloudGuardDataMaskRuleRepresentation), "cloudguard", "dataMaskRule", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardDataMaskRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardDataMaskRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Required, acctest.Create, CloudGuardDataMaskRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "data_mask_categories.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "iam_group_id"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.0.kind", "ALL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardDataMaskRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardDataMaskRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Optional, acctest.Create, CloudGuardDataMaskRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "data_mask_categories.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_mask_rule_status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "iam_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.0.kind", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.0.values.#", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudGuardDataMaskRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Optional, acctest.Update, CloudGuardDataMaskRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "data_mask_categories.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_mask_rule_status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "iam_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.0.kind", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "target_selected.0.values.#", "0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_data_mask_rules", "test_data_mask_rules", acctest.Optional, acctest.Update, CloudGuardCloudGuardDataMaskRuleDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardDataMaskRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Optional, acctest.Update, CloudGuardDataMaskRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "data_mask_rule_status", "DISABLED"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttr(datasourceName, "target_type", "targetType"),
				resource.TestCheckResourceAttr(datasourceName, "data_mask_rule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_mask_rule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_data_mask_rule", "test_data_mask_rule", acctest.Required, acctest.Create, CloudGuardCloudGuardDataMaskRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardDataMaskRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_mask_rule_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_mask_categories.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_mask_rule_status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_selected.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_selected.0.kind", "ALL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_selected.0.values.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudGuardDataMaskRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardDataMaskRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_data_mask_rule" {
			noResourceFound = false
			request := oci_cloud_guard.GetDataMaskRuleRequest{}

			tmp := rs.Primary.ID
			request.DataMaskRuleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetDataMaskRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudGuardDataMaskRule") {
		resource.AddTestSweepers("CloudGuardDataMaskRule", &resource.Sweeper{
			Name:         "CloudGuardDataMaskRule",
			Dependencies: acctest.DependencyGraph["dataMaskRule"],
			F:            sweepCloudGuardDataMaskRuleResource,
		})
	}
}

func sweepCloudGuardDataMaskRuleResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	dataMaskRuleIds, err := getCloudGuardDataMaskRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, dataMaskRuleId := range dataMaskRuleIds {
		if ok := acctest.SweeperDefaultResourceId[dataMaskRuleId]; !ok {
			deleteDataMaskRuleRequest := oci_cloud_guard.DeleteDataMaskRuleRequest{}

			deleteDataMaskRuleRequest.DataMaskRuleId = &dataMaskRuleId

			deleteDataMaskRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteDataMaskRule(context.Background(), deleteDataMaskRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting DataMaskRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataMaskRuleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dataMaskRuleId, CloudGuardDataMaskRuleSweepWaitCondition, time.Duration(3*time.Minute),
				CloudGuardDataMaskRuleSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getCloudGuardDataMaskRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataMaskRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listDataMaskRulesRequest := oci_cloud_guard.ListDataMaskRulesRequest{}
	listDataMaskRulesRequest.CompartmentId = &compartmentId
	listDataMaskRulesRequest.LifecycleState = oci_cloud_guard.ListDataMaskRulesLifecycleStateActive
	listDataMaskRulesResponse, err := cloudGuardClient.ListDataMaskRules(context.Background(), listDataMaskRulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataMaskRule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataMaskRule := range listDataMaskRulesResponse.Items {
		id := *dataMaskRule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataMaskRuleId", id)
	}
	return resourceIds, nil
}

func CloudGuardDataMaskRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataMaskRuleResponse, ok := response.Response.(oci_cloud_guard.GetDataMaskRuleResponse); ok {
		return dataMaskRuleResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func CloudGuardDataMaskRuleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetDataMaskRule(context.Background(), oci_cloud_guard.GetDataMaskRuleRequest{
		DataMaskRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
