package psa

import (
	"fmt"

	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportPsaPrivateServiceAccesHints.GetIdFn = getPsaPrivateServiceAccesId
	tf_export.RegisterCompartmentGraphs("psa", psaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getPsaPrivateServiceAccesId(resource *tf_export.OCIResource) (string, error) {

	privateServiceAccessId, ok := resource.SourceAttributes["private_service_access_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find privateServiceAccessId for Psa PrivateServiceAcces")
	}
	return GetPrivateServiceAccesCompositeId(privateServiceAccessId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportPsaPrivateServiceAccesHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_psa_private_service_access",
	DatasourceClass:        "oci_psa_private_service_accesses",
	DatasourceItemsAttr:    "private_service_access_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "private_service_access",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_psa.PrivateServiceAccessLifecycleStateActive),
	},
}

var psaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportPsaPrivateServiceAccesHints},
	},
}
