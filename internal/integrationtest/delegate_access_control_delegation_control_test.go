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
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	vmclusterresourceId = utils.GetEnvSettingWithBlankDefault("TestVMClusterResourceId")

	DelegateAccessControlDelegationControlRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Required, acctest.Create, DelegateAccessControlDelegationControlRepresentation)

	DelegateAccessControlDelegationControlResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Optional, acctest.Update, DelegateAccessControlDelegationControlRepresentation)

	DelegateAccessControlDelegationControlSingularDataSourceRepresentation = map[string]interface{}{
		"delegation_control_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_delegate_access_control_delegation_control.test_delegation_control.id}`},
	}

	DelegateAccessControlDelegationControlDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		// Prakash cannot be combined with the other filters
		//"resource_id":    acctest.Representation{RepType: acctest.Required, Create: vmclusterresourceId},
		"resource_type": acctest.Representation{RepType: acctest.Optional, Create: `VMCLUSTER`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		//"service_name":   acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DelegateAccessControlDelegationControlDataSourceFilterRepresentation}}
	DelegateAccessControlDelegationControlDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_delegate_access_control_delegation_control.test_delegation_control.id}`}},
	}

	DelegateAccessControlDelegationControlRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"delegation_subscription_ids": acctest.Representation{RepType: acctest.Required,
			Create: []string{`${var.create_subs_id}`},
			Update: []string{`${var.update_subs_id}`}},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"notification_message_format": acctest.Representation{RepType: acctest.Required, Create: `JSON`, Update: `HTML`},
		"notification_topic_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.ons_topic_id}`},
		"resource_ids":                acctest.Representation{RepType: acctest.Required, Create: []string{vmclusterresourceId}, Update: []string{vmclusterresourceId}},
		"resource_type":               acctest.Representation{RepType: acctest.Required, Create: `VMCLUSTER`},
		//"defined_tags":                               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_approve_during_maintenance":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"num_approvals_required":                     acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"pre_approved_service_provider_action_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`DLGT_MGMT_FULL_ACCESS`, `DLGT_MGMT_LOG_ACCESS`}, Update: []string{`DLGT_MGMT_FULL_ACCESS`, `DLGT_MGMT_LOG_ACCESS`}},
		//"vault_id":                                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		//"vault_key_id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key.test_key.id}`},
	}

	/*DelegateAccessControlDelegationControlResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
	DefinedTagsDependencies +
	KeyResourceDependencyConfig +
	acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)*/
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegationControlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegationControlResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_delegate_access_control_delegation_control.test_delegation_control"
	datasourceName := "data.oci_delegate_access_control_delegation_controls.test_delegation_controls"
	singularDatasourceName := "data.oci_delegate_access_control_delegation_control.test_delegation_control"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Optional, acctest.Create, DelegateAccessControlDelegationControlRepresentation), "delegateaccesscontrol", "delegationControl", t)

	acctest.ResourceTest(t, testAccCheckDelegateAccessControlDelegationControlDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Required, acctest.Create, DelegateAccessControlDelegationControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "delegation_subscription_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "notification_message_format", "JSON"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VMCLUSTER"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Optional, acctest.Create, DelegateAccessControlDelegationControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "delegation_subscription_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_message_format", "JSON"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "num_approvals_required", "1"),
				resource.TestCheckResourceAttr(resourceName, "pre_approved_service_provider_action_names.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VMCLUSTER"),
				/*resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_key_id"),*/

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DelegateAccessControlDelegationControlRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "delegation_subscription_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_message_format", "JSON"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "num_approvals_required", "1"),
				resource.TestCheckResourceAttr(resourceName, "pre_approved_service_provider_action_names.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VMCLUSTER"),
				/*resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_key_id"),*/

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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Optional, acctest.Update, DelegateAccessControlDelegationControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "delegation_subscription_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_message_format", "HTML"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "num_approvals_required", "1"),
				resource.TestCheckResourceAttr(resourceName, "pre_approved_service_provider_action_names.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "resource_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VMCLUSTER"),
				/*resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_key_id"),*/

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegation_controls", "test_delegation_controls", acctest.Optional, acctest.Update, DelegateAccessControlDelegationControlDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Optional, acctest.Update, DelegateAccessControlDelegationControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				// Prakash, this is not set in output, it is an array
				//resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "VMCLUSTER"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "delegation_control_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "delegation_control_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Required, acctest.Create, DelegateAccessControlDelegationControlSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegationControlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delegation_control_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "delegation_subscription_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_approve_during_maintenance", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_message_format", "HTML"),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_approvals_required", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pre_approved_service_provider_action_names.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "VMCLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				// Prakash this is not deleted
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_deleted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DelegateAccessControlDelegationControlRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDelegateAccessControlDelegationControlDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DelegateAccessControlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_delegate_access_control_delegation_control" {
			noResourceFound = false
			request := oci_delegate_access_control.GetDelegationControlRequest{}

			tmp := rs.Primary.ID
			request.DelegationControlId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "delegate_access_control")

			response, err := client.GetDelegationControl(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_delegate_access_control.DelegationControlLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DelegateAccessControlDelegationControl") {
		resource.AddTestSweepers("DelegateAccessControlDelegationControl", &resource.Sweeper{
			Name:         "DelegateAccessControlDelegationControl",
			Dependencies: acctest.DependencyGraph["delegationControl"],
			F:            sweepDelegateAccessControlDelegationControlResource,
		})
	}
}

func sweepDelegateAccessControlDelegationControlResource(compartment string) error {
	delegateAccessControlClient := acctest.GetTestClients(&schema.ResourceData{}).DelegateAccessControlClient()
	delegationControlIds, err := getDelegateAccessControlDelegationControlIds(compartment)
	if err != nil {
		return err
	}
	for _, delegationControlId := range delegationControlIds {
		if ok := acctest.SweeperDefaultResourceId[delegationControlId]; !ok {
			deleteDelegationControlRequest := oci_delegate_access_control.DeleteDelegationControlRequest{}

			deleteDelegationControlRequest.DelegationControlId = &delegationControlId

			deleteDelegationControlRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "delegate_access_control")
			_, error := delegateAccessControlClient.DeleteDelegationControl(context.Background(), deleteDelegationControlRequest)
			if error != nil {
				fmt.Printf("Error deleting DelegationControl %s %s, It is possible that the resource is already deleted. Please verify manually \n", delegationControlId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &delegationControlId, DelegateAccessControlDelegationControlSweepWaitCondition, time.Duration(3*time.Minute),
				DelegateAccessControlDelegationControlSweepResponseFetchOperation, "delegate_access_control", true)
		}
	}
	return nil
}

func getDelegateAccessControlDelegationControlIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DelegationControlId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	delegateAccessControlClient := acctest.GetTestClients(&schema.ResourceData{}).DelegateAccessControlClient()

	listDelegationControlsRequest := oci_delegate_access_control.ListDelegationControlsRequest{}
	listDelegationControlsRequest.CompartmentId = &compartmentId
	listDelegationControlsRequest.LifecycleState = oci_delegate_access_control.DelegationControlLifecycleStateActive
	listDelegationControlsResponse, err := delegateAccessControlClient.ListDelegationControls(context.Background(), listDelegationControlsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DelegationControl list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, delegationControl := range listDelegationControlsResponse.Items {
		id := *delegationControl.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DelegationControlId", id)
	}
	return resourceIds, nil
}

func DelegateAccessControlDelegationControlSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if delegationControlResponse, ok := response.Response.(oci_delegate_access_control.GetDelegationControlResponse); ok {
		return delegationControlResponse.LifecycleState != oci_delegate_access_control.DelegationControlLifecycleStateDeleted
	}
	return false
}

func DelegateAccessControlDelegationControlSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DelegateAccessControlClient().GetDelegationControl(context.Background(), oci_delegate_access_control.GetDelegationControlRequest{
		DelegationControlId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
