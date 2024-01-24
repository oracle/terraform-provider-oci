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
	ignoreMeshDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	ServiceMeshMeshRequiredOnlyResource = ServiceMeshMeshResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Required, acctest.Create, ServiceMeshMeshRepresentation)

	MeshResourceConfig = ServiceMeshMeshResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Optional, acctest.Update, ServiceMeshMeshRepresentation)

	ServiceMeshServiceMeshMeshSingularDataSourceRepresentation = map[string]interface{}{
		"mesh_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_mesh.test_mesh.id}`},
	}

	ServiceMeshServiceMeshMeshDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_mesh.test_mesh.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: meshDataSourceFilterRepresentation}}
	meshDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_mesh_mesh.test_mesh.id}`}},
	}

	ServiceMeshMeshRepresentation = map[string]interface{}{
		"certificate_authorities": acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshMeshCertificateAuthoritiesRepresentation},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"mtls":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ServiceMeshMeshMtlsRepresentation},
	}
	ServiceMeshMeshCertificateAuthoritiesRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.certificate_authority_id}`},
	}
	ServiceMeshMeshMtlsRepresentation = map[string]interface{}{
		"minimum": acctest.Representation{RepType: acctest.Required, Create: `DISABLED`, Update: `PERMISSIVE`},
	}

	ServiceMeshMeshResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshMeshResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshMeshResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_service_mesh_mesh.test_mesh"
	datasourceName := "data.oci_service_mesh_meshes.test_meshes"
	singularDatasourceName := "data.oci_service_mesh_mesh.test_mesh"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ServiceMeshMeshResourceDependencies+certificateAuthorityIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Optional, acctest.Create, ServiceMeshMeshRepresentation), "servicemesh", "mesh", t)

	acctest.ResourceTest(t, testAccCheckServiceMeshMeshDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + certificateAuthorityIdVariableStr + ServiceMeshMeshResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.0.id", certificateAuthorityId),
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
			Config: config + compartmentIdVariableStr + certificateAuthorityIdVariableStr + ServiceMeshMeshResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + certificateAuthorityIdVariableStr + ServiceMeshMeshResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.0.id", certificateAuthorityId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mtls.0.minimum", "DISABLED"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceMeshMeshResourceDependencies + certificateAuthorityIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.0.id", certificateAuthorityId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mtls.0.minimum", "DISABLED"),
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
			Config: config + compartmentIdVariableStr + ServiceMeshMeshResourceDependencies + certificateAuthorityIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "certificate_authorities.0.id", certificateAuthorityId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mtls.0.minimum", "PERMISSIVE"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_meshes", "test_meshes", acctest.Optional, acctest.Update, ServiceMeshServiceMeshMeshDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceMeshMeshResourceDependencies + certificateAuthorityIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Optional, acctest.Update, ServiceMeshMeshRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "mesh_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "mesh_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_mesh", "test_mesh", acctest.Required, acctest.Create, ServiceMeshServiceMeshMeshSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MeshResourceConfig + certificateAuthorityIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mesh_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_authorities.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_authorities.0.id", certificateAuthorityId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mtls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mtls.0.minimum", "PERMISSIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ServiceMeshMeshRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceMeshMeshDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceMeshClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_mesh_mesh" {
			noResourceFound = false
			request := oci_service_mesh.GetMeshRequest{}

			tmp := rs.Primary.ID
			request.MeshId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")

			response, err := client.GetMesh(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_mesh.MeshLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ServiceMeshMesh") {
		resource.AddTestSweepers("ServiceMeshMesh", &resource.Sweeper{
			Name:         "ServiceMeshMesh",
			Dependencies: acctest.DependencyGraph["mesh"],
			F:            sweepServiceMeshMeshResource,
		})
	}
}

func sweepServiceMeshMeshResource(compartment string) error {
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()
	meshIds, err := getServiceMeshMeshIds(compartment)
	if err != nil {
		return err
	}
	for _, meshId := range meshIds {
		if ok := acctest.SweeperDefaultResourceId[meshId]; !ok {
			deleteMeshRequest := oci_service_mesh.DeleteMeshRequest{}

			deleteMeshRequest.MeshId = &meshId

			deleteMeshRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")
			_, error := serviceMeshClient.DeleteMesh(context.Background(), deleteMeshRequest)
			if error != nil {
				fmt.Printf("Error deleting Mesh %s %s, It is possible that the resource is already deleted. Please verify manually \n", meshId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &meshId, ServiceMeshMeshSweepWaitCondition, time.Duration(3*time.Minute),
				ServiceMeshMeshSweepResponseFetchOperation, "service_mesh", true)
		}
	}
	return nil
}

func getServiceMeshMeshIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MeshId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()

	listMeshesRequest := oci_service_mesh.ListMeshesRequest{}
	listMeshesRequest.CompartmentId = &compartmentId
	listMeshesRequest.LifecycleState = oci_service_mesh.MeshLifecycleStateActive
	listMeshesResponse, err := serviceMeshClient.ListMeshes(context.Background(), listMeshesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Mesh list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mesh := range listMeshesResponse.Items {
		id := *mesh.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MeshId", id)
	}
	return resourceIds, nil
}

func ServiceMeshMeshSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if meshResponse, ok := response.Response.(oci_service_mesh.GetMeshResponse); ok {
		return meshResponse.LifecycleState != oci_service_mesh.MeshLifecycleStateDeleted
	}
	return false
}

func ServiceMeshMeshSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceMeshClient().GetMesh(context.Background(), oci_service_mesh.GetMeshRequest{
		MeshId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
