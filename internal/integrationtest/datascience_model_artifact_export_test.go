// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	ModelArtifactExportRequiredOnlyResource = ModelArtifactExportDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_export", "test_model_artifact_export", acctest.Required, acctest.Create, modelArtifactExportRepresentation)
	modelArtifactExportRepresentation = map[string]interface{}{
		"model_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"namespace":            acctest.Representation{RepType: acctest.Required, Create: `ociodscdev`},
		"artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `ORACLE_OBJECT_STORAGE`},
		"source_bucket":        acctest.Representation{RepType: acctest.Required, Create: `ds-scratch`},
		"source_region":        acctest.Representation{RepType: acctest.Required, Create: `us-ashburn-1`},
		"source_object_name":   acctest.Representation{RepType: acctest.Required, Create: `custom-object-name`},
	}
	ModelArtifactExportDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelArtifactExportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelArtifactExportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceName := "oci_datascience_model_artifact_export.test_model_artifact_export"

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	acctest.SaveConfigContent(config+compartmentIdVariableStr+ModelArtifactExportDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_export", "test_model_artifact_export", acctest.Optional, acctest.Create, modelArtifactExportRepresentation), "datascience", "modelArtifactExport", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelExportArtifactDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ModelArtifactExportDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_artifact_export", "test_model_artifact_export", acctest.Required, acctest.Create, modelArtifactExportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "ociodscdev"),
				resource.TestCheckResourceAttr(resourceName, "artifact_source_type", "ORACLE_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "source_bucket", "ds-scratch"),
				resource.TestCheckResourceAttr(resourceName, "source_region", "us-ashburn-1"),
				resource.TestCheckResourceAttr(resourceName, "source_object_name", "custom-object-name"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}

func testAccCheckDatascienceModelExportArtifactDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model" {
			noResourceFound = false
			request := oci_datascience.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("CHECK Resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
		return fmt.Errorf("CHECK At least one resource was expected from the state file, but could not be found")
	}
	return nil
}
