package tags

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	TagNamespaceResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", acctest.Required, acctest.Create, TagNamespaceResourceRepresentation)
	TagNamespaceResourceRepresentation = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `${var.tag_namespace_name}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `BasedDB: Defined Tag Namespace for Terraform Testing`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_retired":     acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
)
