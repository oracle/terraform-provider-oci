package tags

import "github.com/oracle/terraform-provider-oci/internal/acctest"

var (
	TagResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_identity_tag", "test_defined_tag", acctest.Required, acctest.Create, TagResourceRepresentation)
	TagResourceRepresentation = map[string]interface{}{
		"name":             acctest.Representation{RepType: acctest.Required, Create: `tfDefinedTag`},
		"description":      acctest.Representation{RepType: acctest.Required, Create: `BasedDB: Defined Tag for Terraform Testing`},
		"is_retired":       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tag_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.test_tag_namespace.id}`},
	}
)
