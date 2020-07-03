// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VnicAttachmentRequiredOnlyResource = VnicAttachmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", Required, Create, vnicAttachmentRepresentation)

	VnicAttachmentResourceConfig = VnicAttachmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", Optional, Create, vnicAttachmentRepresentation)

	vnicAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_id":         Representation{repType: Optional, create: `${oci_core_instance.test_instance.id}`},
		"filter":              RepresentationGroup{Required, vnicAttachmentDataSourceFilterRepresentation}}
	vnicAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_vnic_attachment.test_vnic_attachment.id}`}},
	}

	vnicAttachmentRepresentation = map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Required, vnicAttachmentCreateVnicDetailsRepresentation},
		"instance_id":         Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
		"nic_index":           Representation{repType: Optional, create: `0`},
	}
	vnicAttachmentCreateVnicDetailsRepresentation = map[string]interface{}{
		"subnet_id":              Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"assign_public_ip":       Representation{repType: Optional, create: `false`},
		"defined_tags":           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{repType: Optional, create: `displayName`},
		"freeform_tags":          Representation{repType: Optional, create: map[string]string{"Department": "Accounting"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":         Representation{repType: Optional, create: `attachvnictestinstance`},
		"nsg_ids":                Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{}},
		"private_ip":             Representation{repType: Optional, create: `10.0.0.5`},
		"skip_source_dest_check": Representation{repType: Optional, create: `false`},
	}

	VnicAttachmentResourceDependencies = OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestCoreVnicAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVnicAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vnic_attachment.test_vnic_attachment"
	datasourceName := "data.oci_core_vnic_attachments.test_vnic_attachments"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVnicAttachmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", Required, Create, vnicAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", Optional, Create, vnicAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "attachvnictestinstance"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.nsg_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_vnic_attachments", "test_vnic_attachments", Optional, Update, vnicAttachmentDataSourceRepresentation) +
					compartmentIdVariableStr + VnicAttachmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", Optional, Update, vnicAttachmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),

					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.nic_index", "0"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.vlan_tag"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.vnic_id"),
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

func testAccCheckCoreVnicAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vnic_attachment" {
			noResourceFound = false
			request := oci_core.GetVnicAttachmentRequest{}

			tmp := rs.Primary.ID
			request.VnicAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetVnicAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VnicAttachmentLifecycleStateDetached): true,
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
	if !inSweeperExcludeList("CoreVnicAttachment") {
		resource.AddTestSweepers("CoreVnicAttachment", &resource.Sweeper{
			Name:         "CoreVnicAttachment",
			Dependencies: DependencyGraph["vnicAttachment"],
			F:            sweepCoreVnicAttachmentResource,
		})
	}
}

func sweepCoreVnicAttachmentResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	vnicAttachmentIds, err := getVnicAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, vnicAttachmentId := range vnicAttachmentIds {
		if ok := SweeperDefaultResourceId[vnicAttachmentId]; !ok {
			detachVnicRequest := oci_core.DetachVnicRequest{}

			detachVnicRequest.VnicAttachmentId = &vnicAttachmentId

			detachVnicRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeClient.DetachVnic(context.Background(), detachVnicRequest)
			if error != nil {
				fmt.Printf("Error deleting VnicAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", vnicAttachmentId, error)
				continue
			}
			waitTillCondition(testAccProvider, &vnicAttachmentId, vnicAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				vnicAttachmentSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVnicAttachmentIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VnicAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

	listVnicAttachmentsRequest := oci_core.ListVnicAttachmentsRequest{}
	listVnicAttachmentsRequest.CompartmentId = &compartmentId
	listVnicAttachmentsResponse, err := computeClient.ListVnicAttachments(context.Background(), listVnicAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VnicAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vnicAttachment := range listVnicAttachmentsResponse.Items {
		id := *vnicAttachment.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VnicAttachmentId", id)
	}
	return resourceIds, nil
}

func vnicAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vnicAttachmentResponse, ok := response.Response.(oci_core.GetVnicAttachmentResponse); ok {
		return vnicAttachmentResponse.LifecycleState != oci_core.VnicAttachmentLifecycleStateDetached
	}
	return false
}

func vnicAttachmentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.computeClient().GetVnicAttachment(context.Background(), oci_core.GetVnicAttachmentRequest{
		VnicAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
