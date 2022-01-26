// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	VnicAttachmentRequiredOnlyResource = VnicAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Required, acctest.Create, vnicAttachmentRepresentation)

	VnicAttachmentResourceConfig = VnicAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Optional, acctest.Create, vnicAttachmentRepresentation)

	vnicAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.id}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: vnicAttachmentDataSourceFilterRepresentation}}
	vnicAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_vnic_attachment.test_vnic_attachment.id}`}},
	}

	vnicAttachmentRepresentation = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: vnicAttachmentCreateVnicDetailsRepresentation},
		"instance_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}
	vnicAttachmentCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `attachvnictestinstance`},
		"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"private_ip":                acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"skip_source_dest_check":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	VnicAttachmentResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreVnicAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVnicAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vnic_attachment.test_vnic_attachment"
	datasourceName := "data.oci_core_vnic_attachments.test_vnic_attachments"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VnicAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Optional, acctest.Create, vnicAttachmentRepresentation), "core", "vnicAttachment", t)

	acctest.ResourceTest(t, testAccCheckCoreVnicAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Required, acctest.Create, vnicAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Optional, acctest.Create, vnicAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vnic_attachments", "test_vnic_attachments", acctest.Optional, acctest.Update, vnicAttachmentDataSourceRepresentation) +
				compartmentIdVariableStr + VnicAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Optional, acctest.Update, vnicAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"create_vnic_details.0.assign_private_dns_record",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreVnicAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vnic_attachment" {
			noResourceFound = false
			request := oci_core.GetVnicAttachmentRequest{}

			tmp := rs.Primary.ID
			request.VnicAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreVnicAttachment") {
		resource.AddTestSweepers("CoreVnicAttachment", &resource.Sweeper{
			Name:         "CoreVnicAttachment",
			Dependencies: acctest.DependencyGraph["vnicAttachment"],
			F:            sweepCoreVnicAttachmentResource,
		})
	}
}

func sweepCoreVnicAttachmentResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	vnicAttachmentIds, err := getVnicAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, vnicAttachmentId := range vnicAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[vnicAttachmentId]; !ok {
			detachVnicRequest := oci_core.DetachVnicRequest{}

			detachVnicRequest.VnicAttachmentId = &vnicAttachmentId

			detachVnicRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DetachVnic(context.Background(), detachVnicRequest)
			if error != nil {
				fmt.Printf("Error deleting VnicAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", vnicAttachmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vnicAttachmentId, vnicAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				vnicAttachmentSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVnicAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VnicAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listVnicAttachmentsRequest := oci_core.ListVnicAttachmentsRequest{}
	listVnicAttachmentsRequest.CompartmentId = &compartmentId
	listVnicAttachmentsResponse, err := computeClient.ListVnicAttachments(context.Background(), listVnicAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VnicAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vnicAttachment := range listVnicAttachmentsResponse.Items {
		id := *vnicAttachment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VnicAttachmentId", id)
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

func vnicAttachmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetVnicAttachment(context.Background(), oci_core.GetVnicAttachmentRequest{
		VnicAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
