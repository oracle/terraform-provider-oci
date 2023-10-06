// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterAttachmentRequiredOnlyResource = ContainerengineClusterAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Required, acctest.Create, ContainerengineClusterAttachmentRepresentation)

	ContainerengineClusterAttachmentResourceConfig = ContainerengineClusterAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Optional, acctest.Update, ContainerengineClusterAttachmentRepresentation)

	ContainerengineClusterAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_attachment.test_cluster_attachment.id}`},
	}

	ContainerengineClusterAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster_attachment.test_cluster_attachment.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterAttachmentDataSourceFilterRepresentation}}
	ContainerengineClusterAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_cluster_attachment.test_cluster_attachment.id}`}},
	}

	ContainerengineClusterAttachmentRepresentation = map[string]interface{}{
		"cluster_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"cluster_namespace_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"depends_on":                   acctest.Representation{RepType: acctest.Required, Create: []string{`oci_containerengine_cluster.test_cluster`}},
	}

	ContainerengineClusterAttachmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(ContainerengineClusterRepresentation, map[string]interface{}{
			"cluster_pod_network_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
				"cni_type": acctest.Representation{RepType: acctest.Required, Create: `OCI_VCN_IP_NATIVE`},
			}},
			"endpoint_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
				"is_public_ip_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
				"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
			}},
			"type": acctest.Representation{RepType: acctest.Required, Create: `ENHANCED_CLUSTER`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
				"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
				"route_table_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_route_table.test_route_table.id}`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
			"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: map[string]interface{}{
				"protocol": acctest.Representation{RepType: acctest.Required, Create: `6`},
				"source":   acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
				"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
					"max": acctest.Representation{RepType: acctest.Required, Create: `6443`},
					"min": acctest.Representation{RepType: acctest.Required, Create: `6443`},
				}},
			}}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreRouteTableRepresentation, map[string]interface{}{
			"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
				"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
				"destination":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
				"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `CIDR_BLOCK`},
			}},
		})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_containerengine_cluster_attachment.test_cluster_attachment"
	datasourceName := "data.oci_containerengine_cluster_attachments.test_cluster_attachments"
	singularDatasourceName := "data.oci_containerengine_cluster_attachment.test_cluster_attachment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Optional, acctest.Create, ContainerengineClusterAttachmentRepresentation), "containerengine", "clusterAttachment", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Required, acctest.Create, ContainerengineClusterAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
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
			Config: config + compartmentIdVariableStr + ContainerengineClusterAttachmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Optional, acctest.Create, ContainerengineClusterAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ContainerengineClusterAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerengineClusterAttachmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ContainerengineClusterAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Optional, acctest.Update, ContainerengineClusterAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_attachments", "test_cluster_attachments", acctest.Optional, acctest.Update, ContainerengineClusterAttachmentDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterAttachmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Optional, acctest.Update, ContainerengineClusterAttachmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_attachment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_attachment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Required, acctest.Create, ContainerengineClusterAttachmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_attachment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineClusterAttachmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineClusterAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster_attachment" {
			noResourceFound = false
			request := oci_containerengine.GetClusterAttachmentRequest{}

			tmp := rs.Primary.ID
			request.ClusterAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetClusterAttachment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.ClusterAttachmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineClusterAttachment") {
		resource.AddTestSweepers("ContainerengineClusterAttachment", &resource.Sweeper{
			Name:         "ContainerengineClusterAttachment",
			Dependencies: acctest.DependencyGraph["clusterAttachment"],
			F:            sweepContainerengineClusterAttachmentResource,
		})
	}
}

func sweepContainerengineClusterAttachmentResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterAttachmentIds, err := getContainerengineClusterAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterAttachmentId := range clusterAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[clusterAttachmentId]; !ok {
			deleteClusterAttachmentRequest := oci_containerengine.DeleteClusterAttachmentRequest{}

			deleteClusterAttachmentRequest.ClusterAttachmentId = &clusterAttachmentId

			deleteClusterAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteClusterAttachment(context.Background(), deleteClusterAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterAttachmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterAttachmentId, ContainerengineClusterAttachmentSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineClusterAttachmentSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineClusterAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClusterAttachmentsRequest := oci_containerengine.ListClusterAttachmentsRequest{}
	listClusterAttachmentsRequest.CompartmentId = &compartmentId
	listClusterAttachmentsRequest.LifecycleState = oci_containerengine.ClusterAttachmentLifecycleStateActive
	listClusterAttachmentsResponse, err := containerEngineClient.ListClusterAttachments(context.Background(), listClusterAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ClusterAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, clusterAttachment := range listClusterAttachmentsResponse.Items {
		id := *clusterAttachment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterAttachmentId", id)
	}
	return resourceIds, nil
}

func ContainerengineClusterAttachmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterAttachmentResponse, ok := response.Response.(oci_containerengine.GetClusterAttachmentResponse); ok {
		return clusterAttachmentResponse.LifecycleState != oci_containerengine.ClusterAttachmentLifecycleStateDeleted
	}
	return false
}

func ContainerengineClusterAttachmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetClusterAttachment(context.Background(), oci_containerengine.GetClusterAttachmentRequest{
		ClusterAttachmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
