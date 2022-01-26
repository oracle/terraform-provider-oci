// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_visual_builder "github.com/oracle/oci-go-sdk/v56/visualbuilder"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VbInstanceRequiredOnlyResource = VbInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Required, acctest.Create, vbInstanceRepresentation)

	VbInstanceResourceConfig = VbInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, vbInstanceRepresentation)

	vbInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"vb_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_visual_builder_vb_instance.test_vb_instance.id}`},
	}

	vbInstanceApplicationsSingularDataSourceRepresentation = map[string]interface{}{
		"vb_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_visual_builder_vb_instance.test_vb_instance.id}`},
		"idcs_open_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
	}

	vbInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: vbInstanceDataSourceFilterRepresentation}}
	vbInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_visual_builder_vb_instance.test_vb_instance.id}`}},
	}

	vbInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"node_count":     acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		// Not supported yet
		// "alternate_custom_endpoints": RepresentationGroup{Optional, vbInstanceAlternateCustomEndpointsRepresentation},
		"consumption_model":         acctest.Representation{RepType: acctest.Optional, Create: `UCM`},
		"custom_endpoint":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: vbInstanceCustomEndpointRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"idcs_open_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"is_visual_builder_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	vbInstanceAlternateCustomEndpointsRepresentation = map[string]interface{}{
		"hostname":              acctest.Representation{RepType: acctest.Required, Create: `hostname`, Update: `hostname2`},
		"certificate_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
	}
	vbInstanceCustomEndpointRepresentation = map[string]interface{}{
		"hostname":              acctest.Representation{RepType: acctest.Required, Create: `hostname.com`, Update: `hostname2.com`},
		"certificate_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.oci_vault_secret_id}`},
	}

	VbInstanceResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", acctest.Required, acctest.Create, secretDataSourceRepresentation)
)

// issue-routing-tag: visual_builder/default
func TestVisualBuilderVbInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestVisualBuilderVbInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_visual_builder_vb_instance.test_vb_instance"
	datasourceName := "data.oci_visual_builder_vb_instances.test_vb_instances"
	singularDatasourceName := "data.oci_visual_builder_vb_instance.test_vb_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VbInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Create, vbInstanceRepresentation), "visualbuilder", "vbInstance", t)

	acctest.ResourceTest(t, testAccCheckVisualBuilderVbInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Required, acctest.Create, vbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VbInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Create, vbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "alternate_custom_endpoints.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_open_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vbInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "alternate_custom_endpoints.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_open_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + VbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, vbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "alternate_custom_endpoints.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_open_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_visual_builder_vb_instances", "test_vb_instances", acctest.Optional, acctest.Update, vbInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + VbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, vbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "vb_instance_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vb_instance_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Required, acctest.Create, vbInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VbInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vb_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "alternate_custom_endpoints.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alternate_custom_endpoints.0.certificate_secret_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "alternate_custom_endpoints.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_endpoint.0.certificate_secret_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_count", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VbInstanceResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"idcs_open_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckVisualBuilderVbInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VbInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_visual_builder_vb_instance" {
			noResourceFound = false
			request := oci_visual_builder.GetVbInstanceRequest{}

			tmp := rs.Primary.ID
			request.VbInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "visual_builder")

			response, err := client.GetVbInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_visual_builder.VbInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("VisualBuilderVbInstance") {
		resource.AddTestSweepers("VisualBuilderVbInstance", &resource.Sweeper{
			Name:         "VisualBuilderVbInstance",
			Dependencies: acctest.DependencyGraph["vbInstance"],
			F:            sweepVisualBuilderVbInstanceResource,
		})
	}
}

func sweepVisualBuilderVbInstanceResource(compartment string) error {
	vbInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).VbInstanceClient()
	vbInstanceIds, err := getVbInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, vbInstanceId := range vbInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[vbInstanceId]; !ok {
			deleteVbInstanceRequest := oci_visual_builder.DeleteVbInstanceRequest{}

			deleteVbInstanceRequest.VbInstanceId = &vbInstanceId

			deleteVbInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "visual_builder")
			_, error := vbInstanceClient.DeleteVbInstance(context.Background(), deleteVbInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting VbInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", vbInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vbInstanceId, vbInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				vbInstanceSweepResponseFetchOperation, "visual_builder", true)
		}
	}
	return nil
}

func getVbInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VbInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	vbInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).VbInstanceClient()

	listVbInstancesRequest := oci_visual_builder.ListVbInstancesRequest{}
	listVbInstancesRequest.CompartmentId = &compartmentId
	listVbInstancesRequest.LifecycleState = oci_visual_builder.ListVbInstancesLifecycleStateActive
	listVbInstancesResponse, err := vbInstanceClient.ListVbInstances(context.Background(), listVbInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VbInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vbInstance := range listVbInstancesResponse.Items {
		id := *vbInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VbInstanceId", id)
	}
	return resourceIds, nil
}

func vbInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vbInstanceResponse, ok := response.Response.(oci_visual_builder.GetVbInstanceResponse); ok {
		return vbInstanceResponse.LifecycleState != oci_visual_builder.VbInstanceLifecycleStateDeleted
	}
	return false
}

func vbInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VbInstanceClient().GetVbInstance(context.Background(), oci_visual_builder.GetVbInstanceRequest{
		VbInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
