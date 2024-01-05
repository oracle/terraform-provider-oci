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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreDrgAttachmentRequiredOnlyResource = CoreDrgAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, CoreDrgAttachmentRepresentation)

	CoreCoreDrgAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"attachment_type":    acctest.Representation{RepType: acctest.Optional, Create: `VCN`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"drg_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"drg_route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`, Update: `${oci_core_drg_route_table.test_drg_route_table_2.id}`},
		"network_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ATTACHED`},
		"vcn_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgAttachmentDataSourceFilterRepresentation}}
	CoreDrgAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_drg_attachment.test_drg_attachment.id}`}},
	}

	CoreDrgAttachmentRepresentation = map[string]interface{}{
		"drg_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"drg_route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`, Update: `${oci_core_drg_route_table.test_drg_route_table_2.id}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"network_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgAttachmentNetworkDetailsRepresentation},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}
	CoreDrgAttachmentNetworkDetailsRepresentation = map[string]interface{}{
		"id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `VCN`},
		"route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`, Update: `${oci_core_route_table.test_route_table_2.id}`},
		"vcn_route_type": acctest.Representation{RepType: acctest.Optional, Create: `VCN_CIDRS`, Update: `SUBNET_CIDRS`},
	}

	CoreDrgAttachmentRepresentationNoRouteTable = map[string]interface{}{
		"drg_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `NameNoTable`},
		"drg_route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"network_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgAttachmentNetworkDetailsRepresentationNoRouteTable},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	CoreDrgAttachmentNetworkDetailsRepresentationNoRouteTable = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `VCN`},
	}

	CoreDrgAttachmentTriggerRepresentation = map[string]interface{}{
		"drg_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"network_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgAttachmentNetworkDetailsRepresentation},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName3`, Update: `displayName4`},
		"remove_export_drg_route_distribution_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	CoreDrgAttachmentExportDistributionUpdateRepresentation = map[string]interface{}{
		"drg_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"network_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreDrgAttachmentNetworkDetailsRepresentation},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName3`, Update: `displayName4`},
		"remove_export_drg_route_distribution_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"export_drg_route_distribution_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.default_export_drg_route_distribution_id}`, Update: `${oci_core_drg.test_drg.default_export_drg_route_distribution_id}`},
	}

	CoreDrgAttachmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, CoreDrgRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table_2", acctest.Required, acctest.Create, CoreDrgRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table_2", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreDrgAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment.test_drg_attachment"
	datasourceName := "data.oci_core_drg_attachments.test_drg_attachments"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentRepresentation), "core", "drgAttachment", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgAttachmentDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, CoreDrgAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//delete, before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.type", "VCN"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.vcn_route_type", "VCN_CIDRS"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

		//verify, updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Update, CoreDrgAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.type", "VCN"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.vcn_route_type", "SUBNET_CIDRS"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify remove export trigger
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, CoreDrgAttachmentTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "drg_id")
					return err
				},
			),
		},
		// verify updates with export trigger
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentTriggerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "drg_id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify export drg route distribution id Update
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentExportDistributionUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "export_drg_route_distribution_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies,
		},
		//verify Create for network details with no route table
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentRepresentationNoRouteTable),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "NameNoTable"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//delete, before next Create
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_attachments", "test_drg_attachments", acctest.Optional, acctest.Update, CoreCoreDrgAttachmentDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Update, CoreDrgAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attachment_type", "VCN"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ATTACHED"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "drg_attachments.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.drg_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.drg_route_table_id"),
				resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.is_cross_tenancy"),
				resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.network_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.network_details.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.network_details.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.network_details.0.transport_only_mode"),
				resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.network_details.0.type", "VCN"),
				resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.network_details.0.vcn_route_type", "SUBNET_CIDRS"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.vcn_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreDrgAttachmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreDrgAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg_attachment" {
			noResourceFound = false
			request := oci_core.GetDrgAttachmentRequest{}

			tmp := rs.Primary.ID
			request.DrgAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetDrgAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgAttachmentLifecycleStateDetached): true,
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
	if !acctest.InSweeperExcludeList("CoreDrgAttachment") {
		resource.AddTestSweepers("CoreDrgAttachment", &resource.Sweeper{
			Name:         "CoreDrgAttachment",
			Dependencies: acctest.DependencyGraph["drgAttachment"],
			F:            sweepCoreDrgAttachmentResource,
		})
	}
}

func sweepCoreDrgAttachmentResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	drgAttachmentIds, err := getCoreDrgAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, drgAttachmentId := range drgAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[drgAttachmentId]; !ok {
			deleteDrgAttachmentRequest := oci_core.DeleteDrgAttachmentRequest{}

			deleteDrgAttachmentRequest.DrgAttachmentId = &drgAttachmentId

			deleteDrgAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrgAttachment(context.Background(), deleteDrgAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting DrgAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgAttachmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drgAttachmentId, CoreDrgAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				CoreDrgAttachmentSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreDrgAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrgAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listDrgAttachmentsRequest := oci_core.ListDrgAttachmentsRequest{}
	listDrgAttachmentsRequest.CompartmentId = &compartmentId
	listDrgAttachmentsRequest.LifecycleState = oci_core.DrgAttachmentLifecycleStateAttached
	listDrgAttachmentsResponse, err := virtualNetworkClient.ListDrgAttachments(context.Background(), listDrgAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DrgAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, drgAttachment := range listDrgAttachmentsResponse.Items {
		id := *drgAttachment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrgAttachmentId", id)
	}
	return resourceIds, nil
}

func CoreDrgAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgAttachmentResponse, ok := response.Response.(oci_core.GetDrgAttachmentResponse); ok {
		return drgAttachmentResponse.LifecycleState != oci_core.DrgAttachmentLifecycleStateDetached
	}
	return false
}

func CoreDrgAttachmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetDrgAttachment(context.Background(), oci_core.GetDrgAttachmentRequest{
		DrgAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// Adding a test case for testing the Update request. Updating both drg_route_table_id from and route_table_id simultaneously.
// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreDrgAttachmentUpdateRequest_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentResource_basic")
	defer httpreplay.SaveScenario()
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment.test_drg_attachment"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreDrgAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentRepresentation), "core", "drgAttachment", t)

	acctest.ResourceTest(t, testAccCheckCoreDrgAttachmentDestroy, []resource.TestStep{
		//verify create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Create, CoreDrgAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify, updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Update, CoreDrgAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.route_table_id"),
				resource.TestCheckResourceAttr(resourceName, "network_details.0.type", "VCN"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//delete, before next create
		{
			Config: config + compartmentIdVariableStr + CoreDrgAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Optional, acctest.Update, CoreDrgAttachmentRepresentation),
		},
	})
}
