// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// We need a few resources such as KMS vault and container image pre-created
	// Therefore, before running the signature tests below, please first set the following env var:
	// TF_VAR_container_image_ocid, TF_VAR_kms_vault_ocid, TF_VAR_management_endpoint, and TF_VAR_container_image_signing_signature

	message             = utils.GetEnvSettingWithBlankDefault("container_image_signing_signature")
	signingAlgorithm    = "SHA_224_RSA_PKCS_PSS"
	signingAlgorithmStr = fmt.Sprintf("variable \"signingAlgorithm\" { default = \"%s\" }\n", signingAlgorithm)
	description         = "Image built by TC"
	descriptionStr      = fmt.Sprintf("variable \"description\" { default = \"%s\" }\n", description)
	metadata            = "{\\\"buildNumber\\\":\\\"123\\\"}"
	metadataStr         = fmt.Sprintf("variable \"metadata\" { default = \"%s\" }\n", metadata)

	ArtifactsArtifactscontainerImageSignatureRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Required, acctest.Create, ArtifactsArtifactscontainerContainerImageSignatureRepresentation)

	ArtifactsArtifactscontainerImageResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_image", "test_container_image", acctest.Required, acctest.Create, ArtifactsArtifactscontainerImageSingularDataSourceRepresentation)

	ArtifactsArtifactscontainerContainerImageSignatureKmsSignResourceDependencies = KmsSignResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_sign", "test_container_image_signature_kms_sign", acctest.Required, acctest.Create, ArtifactsArtifactscontainerContainerImageSignatureKmsSignRepresentation)

	ArtifactsArtifactscontainerContainerImageSignatureKmsSignRepresentation = map[string]interface{}{
		"crypto_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":            acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"message":           acctest.Representation{RepType: acctest.Required, Create: message},
		"signing_algorithm": acctest.Representation{RepType: acctest.Required, Create: signingAlgorithm},
		"message_type":      acctest.Representation{RepType: acctest.Optional, Create: `RAW`},
	}

	ArtifactsArtifactscontainerContainerImageSignatureRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_artifacts_container_image.test_container_image.compartment_id}`},
		"image_id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_artifacts_container_image.test_container_image.image_id}`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_sign.test_container_image_signature_kms_sign.key_version_id}`},
		"message":            acctest.Representation{RepType: acctest.Required, Create: message},
		"signature":          acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_sign.test_container_image_signature_kms_sign.signature}`},
		"signing_algorithm":  acctest.Representation{RepType: acctest.Required, Create: signingAlgorithm},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ArtifactsArtifactscontainerContainerImageSignatureSingularDataSourceRepresentation = map[string]interface{}{
		"image_signature_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_artifacts_container_image_signature.test_container_image_signature.id}`},
	}

	ArtifactsArtifactscontainerContainerImageSignatureResourceDependencies = descriptionStr + metadataStr + signingAlgorithmStr +
		DefinedTagsDependencies +
		ArtifactsArtifactscontainerImageResourceConfig +
		ArtifactsArtifactscontainerContainerImageSignatureKmsSignResourceDependencies

	ArtifactsArtifactscontainerContainerImageSignatureDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_artifacts_container_image.test_container_image.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_artifacts_container_image_signature.test_container_image_signature.display_name}`},
		"image_digest":              acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_artifacts_container_image.test_container_image.digest}`},
		"image_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_artifacts_container_image.test_container_image.image_id}`},
		"kms_key_id":                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"kms_key_version_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_sign.test_container_image_signature_kms_sign.key_version_id}`},
		"repository_id":             acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_artifacts_container_image.test_container_image.repository_id}`},
		"repository_name":           acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_artifacts_container_image.test_container_image.repository_name}`},
		"signing_algorithm":         acctest.Representation{RepType: acctest.Optional, Create: signingAlgorithm},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ArtifactscontainerImageSignatureDataSourceFilterRepresentation},
	}

	ArtifactscontainerImageSignatureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_artifacts_container_image_signature.test_container_image_signature.id}`}},
	}
)

// issue-routing-tag: artifacts/default
func TestArtifactsContainerImageSignatureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsContainerImageSignatureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	resourceName := "oci_artifacts_container_image_signature.test_container_image_signature"
	datasourceName := "data.oci_artifacts_container_image_signatures.test_container_image_signatures"
	singularDatasourceName := "data.oci_artifacts_container_image_signature.test_container_image_signature"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckArtifactsContainerImageSignatureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + ArtifactsArtifactscontainerContainerImageSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Required, acctest.Create, ArtifactsArtifactscontainerContainerImageSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "message"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", signingAlgorithm),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + ArtifactsArtifactscontainerImageResourceConfig,
		},
		// verify Create with optionals
		{
			Config: config + ArtifactsArtifactscontainerContainerImageSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Optional, acctest.Create, ArtifactsArtifactscontainerContainerImageSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttr(resourceName, "message", message),
				resource.TestCheckResourceAttrPair(resourceName, "signature", "oci_kms_sign.test_container_image_signature_kms_sign", "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", signingAlgorithm),
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

		// verify updates to updatable parameters
		{
			Config: config + ArtifactsArtifactscontainerContainerImageSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Optional, acctest.Update, ArtifactsArtifactscontainerContainerImageSignatureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttr(resourceName, "message", message),
				resource.TestCheckResourceAttrPair(resourceName, "signature", "oci_kms_sign.test_container_image_signature_kms_sign", "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", signingAlgorithm),
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
			Config: config + ArtifactsArtifactscontainerContainerImageSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Optional, acctest.Update, ArtifactsArtifactscontainerContainerImageSignatureRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_image_signatures", "test_container_image_signatures", acctest.Optional, acctest.Update, ArtifactsArtifactscontainerContainerImageSignatureDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_digest"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "repository_name"),
				resource.TestCheckResourceAttr(datasourceName, "signing_algorithm", signingAlgorithm),

				resource.TestCheckResourceAttr(datasourceName, "container_image_signature_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_signature_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_signature_collection.0.items.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_signature_collection.0.items.0.freeform_tags.%", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config + ArtifactsArtifactscontainerContainerImageSignatureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Optional, acctest.Update, ArtifactsArtifactscontainerContainerImageSignatureRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_image_signature", "test_container_image_signature", acctest.Required, acctest.Create, ArtifactsArtifactscontainerContainerImageSignatureSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_signature_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "message", encodedMessage),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "signature"),
				resource.TestCheckResourceAttr(singularDatasourceName, "signing_algorithm", signingAlgorithm),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + ArtifactsArtifactscontainerImageSignatureRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckArtifactsContainerImageSignatureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ArtifactsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_artifacts_container_image_signature" {
			noResourceFound = false
			request := oci_artifacts.GetContainerImageSignatureRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.ImageSignatureId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "artifacts")

			response, err := client.GetContainerImageSignature(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_artifacts.ContainerImageSignatureLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ArtifactsContainerImageSignature") {
		resource.AddTestSweepers("ArtifactsContainerImageSignature", &resource.Sweeper{
			Name:         "ArtifactsContainerImageSignature",
			Dependencies: acctest.DependencyGraph["containerImageSignature"],
			F:            sweepArtifactsContainerImageSignatureResource,
		})
	}
}

func sweepArtifactsContainerImageSignatureResource(compartment string) error {
	artifactsClient := acctest.GetTestClients(&schema.ResourceData{}).ArtifactsClient()
	containerImageSignatureIds, err := getContainerImageSignatureIds(compartment)
	if err != nil {
		return err
	}
	for _, containerImageSignatureId := range containerImageSignatureIds {
		if ok := acctest.SweeperDefaultResourceId[containerImageSignatureId]; !ok {
			deleteContainerImageSignatureRequest := oci_artifacts.DeleteContainerImageSignatureRequest{}

			deleteContainerImageSignatureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "artifacts")
			_, error := artifactsClient.DeleteContainerImageSignature(context.Background(), deleteContainerImageSignatureRequest)
			if error != nil {
				fmt.Printf("Error deleting ContainerImageSignature %s %s, It is possible that the resource is already deleted. Please verify manually \n", containerImageSignatureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &containerImageSignatureId, ArtifactsContainerImageSignatureSweepWaitCondition, time.Duration(3*time.Minute),
				ArtifactsContainerImageSignatureSweepResponseFetchOperation, "artifacts", true)
		}
	}
	return nil
}

func getContainerImageSignatureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ContainerImageSignatureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	artifactsClient := acctest.GetTestClients(&schema.ResourceData{}).ArtifactsClient()

	listContainerImageSignaturesRequest := oci_artifacts.ListContainerImageSignaturesRequest{}
	listContainerImageSignaturesRequest.CompartmentId = &compartmentId
	listContainerImageSignaturesResponse, err := artifactsClient.ListContainerImageSignatures(context.Background(), listContainerImageSignaturesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ContainerImageSignature list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, containerImageSignature := range listContainerImageSignaturesResponse.Items {
		id := *containerImageSignature.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ContainerImageSignatureId", id)
	}
	return resourceIds, nil
}

func ArtifactsContainerImageSignatureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if containerImageSignatureResponse, ok := response.Response.(oci_artifacts.GetContainerImageSignatureResponse); ok {
		return containerImageSignatureResponse.LifecycleState != oci_artifacts.ContainerImageSignatureLifecycleStateDeleted
	}
	return false
}

func ArtifactsContainerImageSignatureSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ArtifactsClient().GetContainerImageSignature(context.Background(), oci_artifacts.GetContainerImageSignatureRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
