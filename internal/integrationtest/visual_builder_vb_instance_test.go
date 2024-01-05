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
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_visual_builder "github.com/oracle/oci-go-sdk/v65/visualbuilder"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	VisualBuilderVbInstanceRequiredOnlyResource = VisualBuilderVbInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Required, acctest.Create, VisualBuilderVbInstanceRepresentation)

	VisualBuilderVbInstanceResourceConfig = VisualBuilderVbInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, VisualBuilderVbInstanceRepresentation)

	VisualBuilderVisualBuilderVbInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"vb_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_visual_builder_vb_instance.test_vb_instance.id}`},
	}

	vbInstanceApplicationsSingularDataSourceRepresentation = map[string]interface{}{
		"vb_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_visual_builder_vb_instance.test_vb_instance.id}`},
		"idcs_open_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
	}

	VisualBuilderVisualBuilderVbInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: VisualBuilderVbInstanceDataSourceFilterRepresentation}}
	VisualBuilderVbInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_visual_builder_vb_instance.test_vb_instance.id}`}},
	}

	VisualBuilderVbInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"node_count":     acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		// Not supported yet
		// "alternate_custom_endpoints": RepresentationGroup{Optional, VisualBuilderVbInstanceAlternateCustomEndpointsRepresentation},
		"consumption_model":         acctest.Representation{RepType: acctest.Optional, Create: `UCM`},
		"custom_endpoint":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: VisualBuilderVbInstanceCustomEndpointRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"idcs_open_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"is_visual_builder_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	VisualBuilderVbInstanceAlternateCustomEndpointsRepresentation = map[string]interface{}{
		"hostname":              acctest.Representation{RepType: acctest.Required, Create: `hostname.com`, Update: `hostname2.com`},
		"certificate_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
	}
	VisualBuilderVbInstanceCustomEndpointRepresentation = map[string]interface{}{
		"hostname":              acctest.Representation{RepType: acctest.Required, Create: `hostname.com`, Update: `hostname2.com`},
		"certificate_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.oci_vault_secret_id}`},
	}

	VisualBuilderVbInstanceResourceDependencies = DefinedTagsDependencies
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

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("oci_vault_secret_id")
	vaultSecretIdStr := fmt.Sprintf("variable \"oci_vault_secret_id\" { default = \"%s\" }\n", vaultSecretId)

	resourceName := "oci_visual_builder_vb_instance.test_vb_instance"
	datasourceName := "data.oci_visual_builder_vb_instances.test_vb_instances"
	singularDatasourceName := "data.oci_visual_builder_vb_instance.test_vb_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VisualBuilderVbInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Create, VisualBuilderVbInstanceRepresentation), "visualbuilder", "vbInstance", t)

	acctest.ResourceTest(t, testAccCheckVisualBuilderVbInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsOpenIdVariableStr() + VisualBuilderVbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Required, acctest.Create, VisualBuilderVbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsOpenIdVariableStr(),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				idcsOpenIdVariableStr() +
				vaultSecretIdStr +
				VisualBuilderVbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Create, VisualBuilderVbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "alternate_custom_endpoints.0.certificate_secret_id"),
				//resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_open_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "idcs_info.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "attachments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "management_nat_gateway_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "management_vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_nat_gateway_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "service_vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsOpenIdVariableStr() + vaultSecretIdStr + VisualBuilderVbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(VisualBuilderVbInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "alternate_custom_endpoints.0.certificate_secret_id"),
				//resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_open_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
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
			Config: config + compartmentIdVariableStr + idcsOpenIdVariableStr() + vaultSecretIdStr + VisualBuilderVbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, VisualBuilderVbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "alternate_custom_endpoints.0.certificate_secret_id"),
				//resource.TestCheckResourceAttr(resourceName, "alternate_custom_endpoints.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_endpoint.0.certificate_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_open_id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_url"),
				resource.TestCheckResourceAttr(resourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_visual_builder_vb_instances", "test_vb_instances", acctest.Optional, acctest.Update, VisualBuilderVisualBuilderVbInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + idcsOpenIdVariableStr() + vaultSecretIdStr + VisualBuilderVbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, VisualBuilderVbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "vb_instance_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vb_instance_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Required, acctest.Create, VisualBuilderVisualBuilderVbInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsOpenIdVariableStr() + vaultSecretIdStr + VisualBuilderVbInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, VisualBuilderVbInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vb_instance_id"),

				//resource.TestCheckResourceAttr(singularDatasourceName, "alternate_custom_endpoints.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "alternate_custom_endpoints.0.certificate_secret_version"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "alternate_custom_endpoints.0.hostname", "hostname2"),
				//  `attachments` is an optional field and not always returned.
				//resource.TestCheckResourceAttr(singularDatasourceName, "attachments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_endpoint.0.certificate_secret_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_endpoint.0.hostname", "hostname2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_visual_builder_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_nat_gateway_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_vcn_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_nat_gateway_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_vcn_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_count", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		//// verify applications datasource
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_visual_builder_vb_instance_applications", "test_vb_instance_applications", acctest.Required, acctest.Create, vbInstanceApplicationsSingularDataSourceRepresentation) +
		//		compartmentIdVariableStr + idcsOpenIdVariableStr() + vaultSecretIdStr + VisualBuilderVbInstanceResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_visual_builder_vb_instance", "test_vb_instance", acctest.Optional, acctest.Update, VisualBuilderVbInstanceRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		// Don't know what to test as the data source will be empty because there will be error using this idcs token
		//		// The datasource returns {}
		//		func(s *terraform.State) (err error) {
		//			return nil
		//		},
		//	),
		//},
		// verify resource import
		{
			Config:            config + VisualBuilderVbInstanceRequiredOnlyResource,
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
	vbInstanceIds, err := getVisualBuilderVbInstanceIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &vbInstanceId, VisualBuilderVbInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				VisualBuilderVbInstanceSweepResponseFetchOperation, "visual_builder", true)
		}
	}
	return nil
}

func getVisualBuilderVbInstanceIds(compartment string) ([]string, error) {
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

func VisualBuilderVbInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vbInstanceResponse, ok := response.Response.(oci_visual_builder.GetVbInstanceResponse); ok {
		return vbInstanceResponse.LifecycleState != oci_visual_builder.VbInstanceLifecycleStateDeleted
	}
	return false
}

func VisualBuilderVbInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VbInstanceClient().GetVbInstance(context.Background(), oci_visual_builder.GetVbInstanceRequest{
		VbInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func idcsOpenIdVariableStr() string {
	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	return fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)
}
