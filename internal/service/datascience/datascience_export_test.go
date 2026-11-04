package datascience

import (
	"testing"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func TestGetDatascienceModelArtifactSignatureIdUsesListResponseID(t *testing.T) {
	modelID := "model-id"
	artifactSignatureID := "artifact-signature-id"

	resource := &tf_export.OCIResource{
		Parent: &tf_export.OCIResource{TerraformResource: tf_export.TerraformResource{Id: modelID}},
		SourceAttributes: map[string]interface{}{
			"id": artifactSignatureID,
		},
	}

	got, err := getDatascienceModelArtifactSignatureId(resource)
	if err != nil {
		t.Fatalf("getDatascienceModelArtifactSignatureId returned an error: %v", err)
	}

	want := GetModelArtifactSignatureCompositeId(artifactSignatureID, modelID)
	if got != want {
		t.Fatalf("unexpected composite ID: got %q, want %q", got, want)
	}
}
