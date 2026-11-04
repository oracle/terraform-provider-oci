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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceModelArtifactSignatureRequiredOnlyResource = DatascienceModelArtifactSignatureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Required, acctest.Create, DatascienceModelArtifactSignatureRepresentation)

	DatascienceModelArtifactSignatureResourceConfig = DatascienceModelArtifactSignatureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Optional, acctest.Update, DatascienceModelArtifactSignatureRepresentation)

	DatascienceModelArtifactSignatureSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_signature_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"model_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
	}

	DatascienceModelArtifactSignatureDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"digest":         acctest.Representation{RepType: acctest.Optional, Create: `digest`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelArtifactSignatureDataSourceFilterRepresentation}}
	DatascienceModelArtifactSignatureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model_artifact_signature.test_model_artifact_signature.id}`}},
	}

	DatascienceModelArtifactSignatureRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key_version.test_key_version.id}`},
		"model_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"signing_algorithm":  acctest.Representation{RepType: acctest.Required, Create: `SHA_224_RSA_PKCS_PSS`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"verify_trigger":     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DatascienceModelArtifactSignatureResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: datascience/default
// This test covers the artifact-signature resource lifecycle introduced for Data Science models.
func TestDatascienceModelArtifactSignatureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelArtifactSignatureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model_artifact_signature.test_model_artifact_signature"
	datasourceName := "data.oci_datascience_model_artifact_signatures.test_model_artifact_signatures"
	singularDatasourceName := "data.oci_datascience_model_artifact_signature.test_model_artifact_signature"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelArtifactSignatureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Optional, acctest.Create, DatascienceModelArtifactSignatureRepresentation), "datascience", "modelArtifactSignature", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelArtifactSignatureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelArtifactSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Required, acctest.Create, DatascienceModelArtifactSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS_PSS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelArtifactSignatureResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelArtifactSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Optional, acctest.Create, DatascienceModelArtifactSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "digest"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS_PSS"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceModelArtifactSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelArtifactSignatureRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "digest"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS_PSS"),
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
			Config: config + compartmentIdVariableStr + DatascienceModelArtifactSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Optional, acctest.Update, DatascienceModelArtifactSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "digest"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS_PSS"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_artifact_signatures", "test_model_artifact_signatures", acctest.Optional, acctest.Update, DatascienceModelArtifactSignatureDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelArtifactSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Optional, acctest.Update, DatascienceModelArtifactSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "digest", "digest"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "model_artifact_signatures.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_artifact_signatures.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.created_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.digest"),
				resource.TestCheckResourceAttr(datasourceName, "model_artifact_signatures.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "model_artifact_signatures.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_artifact_signatures.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_artifact_signature", "test_model_artifact_signature", acctest.Required, acctest.Create, DatascienceModelArtifactSignatureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelArtifactSignatureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "artifact_signature_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digest"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "signature"),
				resource.TestCheckResourceAttr(singularDatasourceName, "signing_algorithm", "SHA_224_RSA_PKCS_PSS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceModelArtifactSignatureRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceModelArtifactSignatureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model_artifact_signature" {
			noResourceFound = false
			request := oci_datascience.GetModelArtifactSignatureRequest{}

			if value, ok := rs.Primary.Attributes["artifact_signature_id"]; ok {
				request.ArtifactSignatureId = &value
			}

			if value, ok := rs.Primary.Attributes["model_id"]; ok {
				request.ModelId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModelArtifactSignature(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ArtifactSignatureLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceModelArtifactSignature") {
		resource.AddTestSweepers("DatascienceModelArtifactSignature", &resource.Sweeper{
			Name:         "DatascienceModelArtifactSignature",
			Dependencies: acctest.DependencyGraph["modelArtifactSignature"],
			F:            sweepDatascienceModelArtifactSignatureResource,
		})
	}
}

func sweepDatascienceModelArtifactSignatureResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelArtifactSignatureIds, err := getDatascienceModelArtifactSignatureIds(compartment)
	if err != nil {
		return err
	}
	for _, modelArtifactSignatureId := range modelArtifactSignatureIds {
		if ok := acctest.SweeperDefaultResourceId[modelArtifactSignatureId]; !ok {
			deleteModelArtifactSignatureRequest := oci_datascience.DeleteModelArtifactSignatureRequest{}

			deleteModelArtifactSignatureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModelArtifactSignature(context.Background(), deleteModelArtifactSignatureRequest)
			if error != nil {
				fmt.Printf("Error deleting ModelArtifactSignature %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelArtifactSignatureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelArtifactSignatureId, DatascienceModelArtifactSignatureSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceModelArtifactSignatureSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceModelArtifactSignatureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelArtifactSignatureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listModelArtifactSignaturesRequest := oci_datascience.ListModelArtifactSignaturesRequest{}
	listModelArtifactSignaturesRequest.CompartmentId = &compartmentId

	modelIds, error := getDatascienceModelIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting modelId required for ModelArtifactSignature resource requests \n")
	}
	for _, modelId := range modelIds {
		listModelArtifactSignaturesRequest.ModelId = &modelId

		listModelArtifactSignaturesRequest.LifecycleState = oci_datascience.ListModelArtifactSignaturesLifecycleStateActive
		listModelArtifactSignaturesResponse, err := dataScienceClient.ListModelArtifactSignatures(context.Background(), listModelArtifactSignaturesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ModelArtifactSignature list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, modelArtifactSignature := range listModelArtifactSignaturesResponse.Items {
			id := *modelArtifactSignature.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelArtifactSignatureId", id)
		}

	}
	return resourceIds, nil
}

func DatascienceModelArtifactSignatureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelArtifactSignatureResponse, ok := response.Response.(oci_datascience.GetModelArtifactSignatureResponse); ok {
		return modelArtifactSignatureResponse.LifecycleState != oci_datascience.ArtifactSignatureLifecycleStateDeleted
	}
	return false
}

func DatascienceModelArtifactSignatureSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModelArtifactSignature(context.Background(), oci_datascience.GetModelArtifactSignatureRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
