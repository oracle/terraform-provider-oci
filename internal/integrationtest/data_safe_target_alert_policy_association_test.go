// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeTargetAlertPolicyAssociationRequiredOnlyResource = DataSafeTargetAlertPolicyAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Required, acctest.Create, targetAlertPolicyAssociationRepresentation)

	DataSafeTargetAlertPolicyAssociationResourceConfig = DataSafeTargetAlertPolicyAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Update, targetAlertPolicyAssociationRepresentation)

	DataSafetargetAlertPolicyAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"target_alert_policy_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id}`},
	}

	DataSafetargetAlertPolicyAssociationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                       acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"alert_policy_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.policy_id}`},
		"compartment_id_in_subtree":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"target_alert_policy_association_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id}`},
		"target_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"filter":                             acctest.RepresentationGroup{RepType: acctest.Required, Group: targetAlertPolicyAssociationDataSourceFilterRepresentation}}
	targetAlertPolicyAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_target_alert_policy_association.test_target_alert_policy_association.id}`}},
	}

	targetAlertPolicyAssociationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_enabled":     acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	targetAlertPolicyAssociationRepresentationComptUpdate = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_enabled":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}

	DataSafeTargetAlertPolicyAssociationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetAlertPolicyAssociationResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Alert Policy resource")
	httpreplay.SetScenario("TestDataSafeTargetAlertPolicyAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	policyId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	resourceName := "oci_data_safe_target_alert_policy_association.test_target_alert_policy_association"
	datasourceName := "data.oci_data_safe_target_alert_policy_associations.test_target_alert_policy_associations"
	singularDatasourceName := "data.oci_data_safe_target_alert_policy_association.test_target_alert_policy_association"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeTargetAlertPolicyAssociationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Create, targetAlertPolicyAssociationRepresentation), "datasafe", "targetAlertPolicyAssociation", t)

	acctest.ResourceTest(t, testAccCheckDataSafeTargetAlertPolicyAssociationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Required, acctest.Create, targetAlertPolicyAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Create, targetAlertPolicyAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(targetAlertPolicyAssociationRepresentationComptUpdate, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Update, targetAlertPolicyAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_alert_policy_associations", "test_target_alert_policy_associations", acctest.Optional, acctest.Update, DataSafetargetAlertPolicyAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Optional, acctest.Update, targetAlertPolicyAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_alert_policy_association_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),

				resource.TestCheckResourceAttr(datasourceName, "target_alert_policy_association_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_alert_policy_association", "test_target_alert_policy_association", acctest.Required, acctest.Create, DataSafetargetAlertPolicyAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_alert_policy_association_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + policyIdVariableStr + DataSafeTargetAlertPolicyAssociationResourceConfig,
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

func testAccCheckDataSafeTargetAlertPolicyAssociationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_target_alert_policy_association" {
			noResourceFound = false
			request := oci_data_safe.GetTargetAlertPolicyAssociationRequest{}

			tmp := rs.Primary.ID
			request.TargetAlertPolicyAssociationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetTargetAlertPolicyAssociation(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.AlertPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeTargetAlertPolicyAssociation") {
		resource.AddTestSweepers("DataSafeTargetAlertPolicyAssociation", &resource.Sweeper{
			Name:         "DataSafeTargetAlertPolicyAssociation",
			Dependencies: acctest.DependencyGraph["targetAlertPolicyAssociation"],
			F:            sweepDataSafeTargetAlertPolicyAssociationResource,
		})
	}
}

func sweepDataSafeTargetAlertPolicyAssociationResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	targetAlertPolicyAssociationIds, err := getDataSafeTargetAlertPolicyAssociationIds(compartment)
	if err != nil {
		return err
	}
	for _, targetAlertPolicyAssociationId := range targetAlertPolicyAssociationIds {
		if ok := acctest.SweeperDefaultResourceId[targetAlertPolicyAssociationId]; !ok {
			deleteTargetAlertPolicyAssociationRequest := oci_data_safe.DeleteTargetAlertPolicyAssociationRequest{}

			deleteTargetAlertPolicyAssociationRequest.TargetAlertPolicyAssociationId = &targetAlertPolicyAssociationId

			deleteTargetAlertPolicyAssociationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteTargetAlertPolicyAssociation(context.Background(), deleteTargetAlertPolicyAssociationRequest)
			if error != nil {
				fmt.Printf("Error deleting TargetAlertPolicyAssociation %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetAlertPolicyAssociationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &targetAlertPolicyAssociationId, DataSafetargetAlertPolicyAssociationsSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafetargetAlertPolicyAssociationsSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeTargetAlertPolicyAssociationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TargetAlertPolicyAssociationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listTargetAlertPolicyAssociationsRequest := oci_data_safe.ListTargetAlertPolicyAssociationsRequest{}
	listTargetAlertPolicyAssociationsRequest.CompartmentId = &compartmentId
	listTargetAlertPolicyAssociationsRequest.LifecycleState = oci_data_safe.ListTargetAlertPolicyAssociationsLifecycleStateActive
	listTargetAlertPolicyAssociationsResponse, err := dataSafeClient.ListTargetAlertPolicyAssociations(context.Background(), listTargetAlertPolicyAssociationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TargetAlertPolicyAssociation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, targetAlertPolicyAssociation := range listTargetAlertPolicyAssociationsResponse.Items {
		id := *targetAlertPolicyAssociation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TargetAlertPolicyAssociationId", id)
	}
	return resourceIds, nil
}

func DataSafetargetAlertPolicyAssociationsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if targetAlertPolicyAssociationResponse, ok := response.Response.(oci_data_safe.GetTargetAlertPolicyAssociationResponse); ok {
		return targetAlertPolicyAssociationResponse.LifecycleState != oci_data_safe.AlertPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafetargetAlertPolicyAssociationsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetTargetAlertPolicyAssociation(context.Background(), oci_data_safe.GetTargetAlertPolicyAssociationRequest{
		TargetAlertPolicyAssociationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
