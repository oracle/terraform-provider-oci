package integrationtest

const (
	DefinedTagsDependencies = `
variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = "${var.tenancy_ocid}"
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
}
`
)
