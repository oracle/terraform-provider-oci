package zpr

import (
	"testing"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func TestProcessZprConfigurationsSetsCompositeImportId(t *testing.T) {
	const compartmentId = "ocid1.tenancy.oc1..test"
	const configurationId = "ocid1.zprconfiguration.oc1..test"

	resources, err := processZprConfigurations(nil, []*tf_export.OCIResource{
		{
			CompartmentId: compartmentId,
			TerraformResource: tf_export.TerraformResource{
				Id: configurationId,
			},
		},
	})
	if err != nil {
		t.Fatalf("processZprConfigurations returned error: %v", err)
	}
	if len(resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(resources))
	}

	expectedImportId := compartmentId + "/" + configurationId
	if resources[0].ImportId != expectedImportId {
		t.Fatalf("expected ImportId %q, got %q", expectedImportId, resources[0].ImportId)
	}
}
