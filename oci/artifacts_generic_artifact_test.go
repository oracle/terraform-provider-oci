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
	oci_artifacts "github.com/oracle/oci-go-sdk/v46/artifacts"
	"github.com/oracle/oci-go-sdk/v46/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	GenericArtifactResourceConfig = GenericArtifactResourceDependencies +
		generateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Optional, Update, genericArtifactRepresentation)

	genericArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_id": Representation{repType: Required, create: `${oci_artifacts_generic_artifact.test_generic_artifact.id}`},
	}

	genericArtifactDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"repository_id":  Representation{repType: Required, create: `${oci_artifacts_repository.test_repository.id}`},
		"artifact_path":  Representation{repType: Optional, create: `artifactPath`},
		"display_name":   Representation{repType: Optional, create: `displayName`},
		"id":             Representation{repType: Optional, create: `${oci_artifacts_generic_artifact.test_generic_artifact.id}`},
		"sha256":         Representation{repType: Optional, create: `sha256`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"version":        Representation{repType: Optional, create: `1.0`},
		"filter":         RepresentationGroup{Required, genericArtifactDataSourceFilterRepresentation}}
	genericArtifactDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_artifacts_generic_artifact.test_generic_artifact.id}`}},
	}

	genericArtifactRepresentation = map[string]interface{}{
		"artifact_id":   Representation{repType: Required, create: `${oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path.id}`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	}

	GenericArtifactResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Required, Create, repositoryRepresentation) +
		generateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation)
)

// issue-routing-tag: artifacts/default
func TestArtifactsGenericArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsGenericArtifactResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_artifacts_generic_artifact.test_generic_artifact"
	datasourceName := "data.oci_artifacts_generic_artifacts.test_generic_artifacts"
	singularDatasourceName := "data.oci_artifacts_generic_artifact.test_generic_artifact"

	var resId, resId2 string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+GenericArtifactResourceDependencies+
		generateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Required, Create, genericArtifactRepresentation), "artifacts", "genericArtifact", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckArtifactsGenericArtifactDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + GenericArtifactResourceDependencies +
					generateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Required, Create, genericArtifactRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(

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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + GenericArtifactResourceDependencies +
					generateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Optional, Update, genericArtifactRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "artifact_path"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "repository_id"),
					resource.TestCheckResourceAttrSet(resourceName, "sha256"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_artifacts_generic_artifacts", "test_generic_artifacts", Optional, Update, genericArtifactDataSourceRepresentation) +
					compartmentIdVariableStr + GenericArtifactResourceDependencies +
					generateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Optional, Update, genericArtifactRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "artifact_path", "artifactPath"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "sha256"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(datasourceName, "version", "1.0"),

					resource.TestCheckResourceAttr(datasourceName, "generic_artifact_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", Required, Create, genericArtifactSingularDataSourceRepresentation) +
					compartmentIdVariableStr + GenericArtifactResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(singularDatasourceName, "artifact_path"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "sha256"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				),
			},
		},
	})
}

func testAccCheckArtifactsGenericArtifactDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).artifactsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_artifacts_generic_artifact" {
			noResourceFound = false
			request := oci_artifacts.GetGenericArtifactRequest{}

			if value, ok := rs.Primary.Attributes["artifact_id"]; ok {
				request.ArtifactId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "artifacts")

			response, err := client.GetGenericArtifact(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_artifacts.GenericArtifactLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("ArtifactsGenericArtifact") {
		resource.AddTestSweepers("ArtifactsGenericArtifact", &resource.Sweeper{
			Name:         "ArtifactsGenericArtifact",
			Dependencies: DependencyGraph["genericArtifact"],
			F:            sweepArtifactsGenericArtifactResource,
		})
	}
}

func sweepArtifactsGenericArtifactResource(compartment string) error {
	artifactsClient := GetTestClients(&schema.ResourceData{}).artifactsClient()
	genericArtifactIds, err := getGenericArtifactIds(compartment)
	if err != nil {
		return err
	}
	for _, genericArtifactId := range genericArtifactIds {
		if ok := SweeperDefaultResourceId[genericArtifactId]; !ok {
			deleteGenericArtifactRequest := oci_artifacts.DeleteGenericArtifactRequest{}

			deleteGenericArtifactRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "artifacts")
			_, error := artifactsClient.DeleteGenericArtifact(context.Background(), deleteGenericArtifactRequest)
			if error != nil {
				fmt.Printf("Error deleting GenericArtifact %s %s, It is possible that the resource is already deleted. Please verify manually \n", genericArtifactId, error)
				continue
			}
			waitTillCondition(testAccProvider, &genericArtifactId, genericArtifactSweepWaitCondition, time.Duration(3*time.Minute),
				genericArtifactSweepResponseFetchOperation, "artifacts", true)
		}
	}
	return nil
}

func getGenericArtifactIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "GenericArtifactId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	artifactsClient := GetTestClients(&schema.ResourceData{}).artifactsClient()

	listGenericArtifactsRequest := oci_artifacts.ListGenericArtifactsRequest{}
	listGenericArtifactsRequest.CompartmentId = &compartmentId

	repositoryIds, error := getRepositoryIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting repositoryId required for GenericArtifact resource requests \n")
	}
	for _, repositoryId := range repositoryIds {
		listGenericArtifactsRequest.RepositoryId = &repositoryId

		state := oci_artifacts.GenericArtifactLifecycleStateAvailable
		listGenericArtifactsRequest.LifecycleState = (*string)(&state)
		listGenericArtifactsResponse, err := artifactsClient.ListGenericArtifacts(context.Background(), listGenericArtifactsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting GenericArtifact list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, genericArtifact := range listGenericArtifactsResponse.Items {
			id := *genericArtifact.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "GenericArtifactId", id)
		}

	}
	return resourceIds, nil
}

func genericArtifactSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if genericArtifactResponse, ok := response.Response.(oci_artifacts.GetGenericArtifactResponse); ok {
		return genericArtifactResponse.LifecycleState != oci_artifacts.GenericArtifactLifecycleStateDeleted
	}
	return false
}

func genericArtifactSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.artifactsClient().GetGenericArtifact(context.Background(), oci_artifacts.GetGenericArtifactRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
