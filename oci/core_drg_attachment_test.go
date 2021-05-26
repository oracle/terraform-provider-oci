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
	"github.com/oracle/oci-go-sdk/v41/common"
	oci_core "github.com/oracle/oci-go-sdk/v41/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DrgAttachmentRequiredOnlyResource = DrgAttachmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation)

	drgAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"attachment_type":    Representation{repType: Optional, create: `VCN`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"drg_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"drg_route_table_id": Representation{repType: Optional, create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"network_id":         Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"state":              Representation{repType: Optional, create: `ATTACHED`},
		"vcn_id":             Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.id}`},
		"filter":             RepresentationGroup{Required, drgAttachmentDataSourceFilterRepresentation}}
	drgAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_drg_attachment.test_drg_attachment.id}`}},
	}

	drgAttachmentRepresentation = map[string]interface{}{
		"drg_id":             Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"drg_route_table_id": Representation{repType: Optional, create: `${oci_core_drg_route_table.test_drg_route_table.id}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"network_details":    RepresentationGroup{Required, drgAttachmentNetworkDetailsRepresentation},
		"lifecycle":          RepresentationGroup{Required, ignoreChangesLBRepresentation},
	}
	drgAttachmentNetworkDetailsRepresentation = map[string]interface{}{
		"id":             Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"type":           Representation{repType: Required, create: `VCN`},
		"route_table_id": Representation{repType: Required, create: `${oci_core_route_table.test_route_table.id}`},
	}

	drgAttachmentTriggerRepresentation = map[string]interface{}{
		"drg_id":          Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"network_details": RepresentationGroup{Required, drgAttachmentNetworkDetailsRepresentation},
		"lifecycle":       RepresentationGroup{Required, ignoreChangesLBRepresentation},
		"display_name":    Representation{repType: Optional, create: `displayName3`, update: `displayName4`},
		"remove_export_drg_route_distribution_trigger": Representation{repType: Optional, create: `false`, update: `true`},
	}

	drgAttachmentExportDistributionUpdateRepresentation = map[string]interface{}{
		"drg_id":          Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"network_details": RepresentationGroup{Required, drgAttachmentNetworkDetailsRepresentation},
		"lifecycle":       RepresentationGroup{Required, ignoreChangesLBRepresentation},
		"display_name":    Representation{repType: Optional, create: `displayName3`, update: `displayName4`},
		"remove_export_drg_route_distribution_trigger": Representation{repType: Optional, create: `false`, update: `true`},
		"export_drg_route_distribution_id":             Representation{repType: Optional, create: `${oci_core_drg.test_drg.default_export_drg_route_distribution_id}`, update: `${oci_core_drg.test_drg.default_export_drg_route_distribution_id}`},
	}

	DrgAttachmentResourceDependencies = generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestCoreDrgAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment.test_drg_attachment"
	datasourceName := "data.oci_core_drg_attachments.test_drg_attachments"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DrgAttachmentResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Optional, Create, drgAttachmentRepresentation), "core", "drgAttachment", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDrgAttachmentDestroy,
		Steps: []resource.TestStep{
			//verify create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			//delete, before next create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies,
			},
			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Optional, Create, drgAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			//verify, updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Optional, Update, drgAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify remove export trigger
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentTriggerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "drg_id")
						return err
					},
				),
			},
			// verify updates with export trigger
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Optional, Create, drgAttachmentTriggerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "drg_id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify export drg route distribution id update
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Optional, Create, drgAttachmentExportDistributionUpdateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "export_drg_route_distribution_id"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentResourceDependencies,
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_drg_attachments", "test_drg_attachments", Optional, Update, drgAttachmentDataSourceRepresentation) +
					compartmentIdVariableStr + DrgAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Optional, Update, drgAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.drg_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.drg_route_table_id"),
					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.is_cross_tenancy"),
					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.network_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.network_details.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.network_details.0.route_table_id"),
					resource.TestCheckResourceAttr(datasourceName, "drg_attachments.0.network_details.0.type", "VCN"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.route_table_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_attachments.0.vcn_id"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreDrgAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg_attachment" {
			noResourceFound = false
			request := oci_core.GetDrgAttachmentRequest{}

			tmp := rs.Primary.ID
			request.DrgAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreDrgAttachment") {
		resource.AddTestSweepers("CoreDrgAttachment", &resource.Sweeper{
			Name:         "CoreDrgAttachment",
			Dependencies: DependencyGraph["drgAttachment"],
			F:            sweepCoreDrgAttachmentResource,
		})
	}
}

func sweepCoreDrgAttachmentResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	drgAttachmentIds, err := getDrgAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, drgAttachmentId := range drgAttachmentIds {
		if ok := SweeperDefaultResourceId[drgAttachmentId]; !ok {
			deleteDrgAttachmentRequest := oci_core.DeleteDrgAttachmentRequest{}

			deleteDrgAttachmentRequest.DrgAttachmentId = &drgAttachmentId

			deleteDrgAttachmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrgAttachment(context.Background(), deleteDrgAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting DrgAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgAttachmentId, error)
				continue
			}
			waitTillCondition(testAccProvider, &drgAttachmentId, drgAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				drgAttachmentSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getDrgAttachmentIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DrgAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "DrgAttachmentId", id)
	}
	return resourceIds, nil
}

func drgAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgAttachmentResponse, ok := response.Response.(oci_core.GetDrgAttachmentResponse); ok {
		return drgAttachmentResponse.LifecycleState != oci_core.DrgAttachmentLifecycleStateDetached
	}
	return false
}

func drgAttachmentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetDrgAttachment(context.Background(), oci_core.GetDrgAttachmentRequest{
		DrgAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
