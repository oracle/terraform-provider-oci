// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ServiceMeshAccessPolicyRequiredOnlyResource = ServiceMeshAccessPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Required, acctest.Create, ServiceMeshAccessPolicyRepresentation)

	ServiceMeshAccessPolicyResourceConfig = ServiceMeshAccessPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Optional, acctest.Update, ServiceMeshAccessPolicyRepresentation)

	ServiceMeshServiceMeshAccessPolicySingularDataSourceRepresentation = map[string]interface{}{
		"access_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_access_policy.test_access_policy.id}`},
	}

	ServiceMeshServiceMeshAccessPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_access_policy.test_access_policy.id}`},
		"mesh_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_mesh.mesh1.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshAccessPolicyDataSourceFilterRepresentation}}
	ServiceMeshAccessPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_mesh_access_policy.test_access_policy.id}`}},
	}

	ServiceMeshAccessPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"mesh_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_mesh.mesh1.id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"rules":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshAccessPolicyRulesRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ServiceMeshAccessPolicyRulesRepresentation = map[string]interface{}{
		"action":      acctest.Representation{RepType: acctest.Required, Create: `ALLOW`, Update: `ALLOW`},
		"destination": acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshAccessPolicyRulesDestinationRepresentation},
		"source":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshAccessPolicyRulesSourceRepresentation},
	}
	ServiceMeshAccessPolicyRulesDestinationRepresentation = map[string]interface{}{
		"type":               acctest.Representation{RepType: acctest.Required, Create: `ALL_VIRTUAL_SERVICES`, Update: `VIRTUAL_SERVICE`},
		"virtual_service_id": acctest.Representation{RepType: acctest.Optional, Update: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
	}
	ServiceMeshAccessPolicyRulesSourceRepresentation = map[string]interface{}{
		"type":               acctest.Representation{RepType: acctest.Required, Create: `ALL_VIRTUAL_SERVICES`, Update: `INGRESS_GATEWAY`},
		"ingress_gateway_id": acctest.Representation{RepType: acctest.Optional, Update: `${oci_service_mesh_ingress_gateway.ingress_gateway_1.id}`},
	}

	ServiceMeshAccessPolicyResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "mesh1", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
			"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "ingress_gateway_1", acctest.Required, acctest.Create, ServiceMeshIngressGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service", "virtual_service_1", acctest.Required, acctest.Create, ServiceMeshVirtualServiceRepresentation)
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshAccessPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshAccessPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	resourceName := "oci_service_mesh_access_policy.test_access_policy"
	datasourceName := "data.oci_service_mesh_access_policies.test_access_policies"
	singularDatasourceName := "data.oci_service_mesh_access_policy.test_access_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+certificateAuthorityIdVariableStr+compartmentIdVariableStr+ServiceMeshAccessPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Optional, acctest.Create, ServiceMeshAccessPolicyRepresentation), "servicemesh", "accessPolicy", t)

	acctest.ResourceTest(t, testAccCheckServiceMeshAccessPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshAccessPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Required, acctest.Create, ServiceMeshAccessPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.0.type", "ALL_VIRTUAL_SERVICES"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.0.type", "ALL_VIRTUAL_SERVICES"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshAccessPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshAccessPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshAccessPolicyRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.0.type", "ALL_VIRTUAL_SERVICES"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.0.type", "ALL_VIRTUAL_SERVICES"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceMeshAccessPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ServiceMeshAccessPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.0.type", "ALL_VIRTUAL_SERVICES"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.0.type", "ALL_VIRTUAL_SERVICES"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshAccessPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(ServiceMeshAccessPolicyRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mesh_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.destination.0.type", "VIRTUAL_SERVICE"),
				resource.TestCheckResourceAttrSet(resourceName, "rules.0.destination.0.virtual_service_id"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "rules.0.source.0.ingress_gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.source.0.type", "INGRESS_GATEWAY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_access_policies", "test_access_policies", acctest.Optional, acctest.Update, ServiceMeshServiceMeshAccessPolicyDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshAccessPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Optional, acctest.Update, ServiceMeshAccessPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "mesh_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "access_policy_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_access_policy", "test_access_policy", acctest.Required, acctest.Create, ServiceMeshServiceMeshAccessPolicySingularDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshAccessPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.destination.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.destination.0.type", "VIRTUAL_SERVICE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.source.0.type", "INGRESS_GATEWAY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ServiceMeshAccessPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceMeshAccessPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceMeshClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_mesh_access_policy" {
			noResourceFound = false
			request := oci_service_mesh.GetAccessPolicyRequest{}

			tmp := rs.Primary.ID
			request.AccessPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")

			response, err := client.GetAccessPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_mesh.AccessPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ServiceMeshAccessPolicy") {
		resource.AddTestSweepers("ServiceMeshAccessPolicy", &resource.Sweeper{
			Name:         "ServiceMeshAccessPolicy",
			Dependencies: acctest.DependencyGraph["accessPolicy"],
			F:            sweepServiceMeshAccessPolicyResource,
		})
	}
}

func sweepServiceMeshAccessPolicyResource(compartment string) error {
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()
	accessPolicyIds, err := getServiceMeshAccessPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, accessPolicyId := range accessPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[accessPolicyId]; !ok {
			deleteAccessPolicyRequest := oci_service_mesh.DeleteAccessPolicyRequest{}

			deleteAccessPolicyRequest.AccessPolicyId = &accessPolicyId

			deleteAccessPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")
			_, error := serviceMeshClient.DeleteAccessPolicy(context.Background(), deleteAccessPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting AccessPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", accessPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &accessPolicyId, ServiceMeshAccessPolicySweepWaitCondition, time.Duration(3*time.Minute),
				ServiceMeshAccessPolicySweepResponseFetchOperation, "service_mesh", true)
		}
	}
	return nil
}

func getServiceMeshAccessPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AccessPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()

	listAccessPoliciesRequest := oci_service_mesh.ListAccessPoliciesRequest{}
	listAccessPoliciesRequest.CompartmentId = &compartmentId
	listAccessPoliciesRequest.LifecycleState = oci_service_mesh.AccessPolicyLifecycleStateActive
	listAccessPoliciesResponse, err := serviceMeshClient.ListAccessPolicies(context.Background(), listAccessPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AccessPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, accessPolicy := range listAccessPoliciesResponse.Items {
		id := *accessPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AccessPolicyId", id)
	}
	return resourceIds, nil
}

func ServiceMeshAccessPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if accessPolicyResponse, ok := response.Response.(oci_service_mesh.GetAccessPolicyResponse); ok {
		return accessPolicyResponse.LifecycleState != oci_service_mesh.AccessPolicyLifecycleStateDeleted
	}
	return false
}

func ServiceMeshAccessPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceMeshClient().GetAccessPolicy(context.Background(), oci_service_mesh.GetAccessPolicyRequest{
		AccessPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
