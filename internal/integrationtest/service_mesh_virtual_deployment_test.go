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
	ServiceMeshVirtualDeploymentRequiredOnlyResource = ServiceMeshVirtualDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Required, acctest.Create, ServiceMeshVirtualDeploymentRepresentation)

	ServiceMeshVirtualDeploymentResourceConfig = ServiceMeshVirtualDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Optional, acctest.Update, ServiceMeshVirtualDeploymentRepresentation)

	ServiceMeshServiceMeshVirtualDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"virtual_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_deployment.test_virtual_deployment.id}`},
	}

	ServiceMeshServiceMeshVirtualDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_virtual_deployment.test_virtual_deployment.id}`},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"virtual_service_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshVirtualDeploymentDataSourceFilterRepresentation}}
	ServiceMeshVirtualDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_mesh_virtual_deployment.test_virtual_deployment.id}`}},
	}

	ServiceMeshVirtualDeploymentRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
		"virtual_service_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
		"listeners":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: ServiceMeshVirtualDeploymentListenersRepresentation},
		"service_discovery":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ServiceMeshVirtualDeploymentServiceDiscoveryRepresentation},
		"access_logging":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ServiceMeshVirtualDeploymentAccessLoggingRepresentation},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ServiceMeshVirtualDeploymentRepresentationWithDisabledServiceDiscovery = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name_disabled`},
		"service_discovery":  acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshVirtualDeploymentDisabledServiceDiscoveryRepresentation},
		"virtual_service_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
	}

	ServiceMeshVirtualDeploymentAccessLoggingRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	ServiceMeshVirtualDeploymentListenersRepresentation = map[string]interface{}{
		"idle_timeout_in_ms":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"request_timeout_in_ms": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"port":                  acctest.Representation{RepType: acctest.Required, Create: `8080`, Update: `8081`},
		"protocol":              acctest.Representation{RepType: acctest.Required, Create: `HTTP`, Update: `GRPC`},
	}
	ServiceMeshVirtualDeploymentServiceDiscoveryRepresentation = map[string]interface{}{
		"type":     acctest.Representation{RepType: acctest.Required, Create: `DNS`, Update: `DNS`},
		"hostname": acctest.Representation{RepType: acctest.Optional, Create: `hostname`, Update: `hostname2`},
	}

	ServiceMeshVirtualDeploymentDisabledServiceDiscoveryRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
	}

	ServiceMeshVirtualDeploymentResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "mesh1", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
			"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service", "virtual_service_1", acctest.Required, acctest.Create, ServiceMeshVirtualServiceRepresentation)
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshVirtualDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshVirtualDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	resourceName := "oci_service_mesh_virtual_deployment.test_virtual_deployment"
	datasourceName := "data.oci_service_mesh_virtual_deployments.test_virtual_deployments"
	singularDatasourceName := "data.oci_service_mesh_virtual_deployment.test_virtual_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+certificateAuthorityIdVariableStr+compartmentIdVariableStr+ServiceMeshVirtualDeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Optional, acctest.Create, ServiceMeshVirtualDeploymentRepresentation), "servicemesh", "virtualDeployment", t)

	acctest.ResourceTest(t, testAccCheckServiceMeshVirtualDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Required, acctest.Create, ServiceMeshVirtualDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.type", "DISABLED"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies,
		},

		// verify Create with Disabled Mode
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Required, acctest.Create, ServiceMeshVirtualDeploymentRepresentationWithDisabledServiceDiscovery),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name_disabled"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.type", "DISABLED"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshVirtualDeploymentRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_logging.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.idle_timeout_in_ms", "10"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.request_timeout_in_ms", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.type", "DNS"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

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
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceMeshVirtualDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ServiceMeshVirtualDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_logging.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.idle_timeout_in_ms", "10"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.request_timeout_in_ms", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.type", "DNS"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

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
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Optional, acctest.Update, ServiceMeshVirtualDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "access_logging.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listeners.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.idle_timeout_in_ms", "11"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.port", "8081"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.protocol", "GRPC"),
				resource.TestCheckResourceAttr(resourceName, "listeners.0.request_timeout_in_ms", "11"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(resourceName, "service_discovery.0.type", "DNS"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_virtual_deployments", "test_virtual_deployments", acctest.Optional, acctest.Update, ServiceMeshServiceMeshVirtualDeploymentDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Optional, acctest.Update, ServiceMeshVirtualDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_service_id"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_deployment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "test_virtual_deployment", acctest.Required, acctest.Create, ServiceMeshServiceMeshVirtualDeploymentSingularDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_logging.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "access_logging.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listeners.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listeners.0.port", "8081"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listeners.0.idle_timeout_in_ms", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listeners.0.protocol", "GRPC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listeners.0.request_timeout_in_ms", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_discovery.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_discovery.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_discovery.0.type", "DNS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ServiceMeshVirtualDeploymentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceMeshVirtualDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceMeshClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_mesh_virtual_deployment" {
			noResourceFound = false
			request := oci_service_mesh.GetVirtualDeploymentRequest{}

			tmp := rs.Primary.ID
			request.VirtualDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")

			response, err := client.GetVirtualDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_mesh.VirtualDeploymentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ServiceMeshVirtualDeployment") {
		resource.AddTestSweepers("ServiceMeshVirtualDeployment", &resource.Sweeper{
			Name:         "ServiceMeshVirtualDeployment",
			Dependencies: acctest.DependencyGraph["virtualDeployment"],
			F:            sweepServiceMeshVirtualDeploymentResource,
		})
	}
}

func sweepServiceMeshVirtualDeploymentResource(compartment string) error {
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()
	virtualDeploymentIds, err := getServiceMeshVirtualDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, virtualDeploymentId := range virtualDeploymentIds {
		if ok := acctest.SweeperDefaultResourceId[virtualDeploymentId]; !ok {
			deleteVirtualDeploymentRequest := oci_service_mesh.DeleteVirtualDeploymentRequest{}

			deleteVirtualDeploymentRequest.VirtualDeploymentId = &virtualDeploymentId

			deleteVirtualDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")
			_, error := serviceMeshClient.DeleteVirtualDeployment(context.Background(), deleteVirtualDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting VirtualDeployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", virtualDeploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &virtualDeploymentId, ServiceMeshVirtualDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				ServiceMeshVirtualDeploymentSweepResponseFetchOperation, "service_mesh", true)
		}
	}
	return nil
}

func getServiceMeshVirtualDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VirtualDeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()

	listVirtualDeploymentsRequest := oci_service_mesh.ListVirtualDeploymentsRequest{}
	listVirtualDeploymentsRequest.CompartmentId = &compartmentId
	listVirtualDeploymentsRequest.LifecycleState = oci_service_mesh.VirtualDeploymentLifecycleStateActive
	listVirtualDeploymentsResponse, err := serviceMeshClient.ListVirtualDeployments(context.Background(), listVirtualDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VirtualDeployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, virtualDeployment := range listVirtualDeploymentsResponse.Items {
		id := *virtualDeployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VirtualDeploymentId", id)
	}
	return resourceIds, nil
}

func ServiceMeshVirtualDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if virtualDeploymentResponse, ok := response.Response.(oci_service_mesh.GetVirtualDeploymentResponse); ok {
		return virtualDeploymentResponse.LifecycleState != oci_service_mesh.VirtualDeploymentLifecycleStateDeleted
	}
	return false
}

func ServiceMeshVirtualDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceMeshClient().GetVirtualDeployment(context.Background(), oci_service_mesh.GetVirtualDeploymentRequest{
		VirtualDeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
