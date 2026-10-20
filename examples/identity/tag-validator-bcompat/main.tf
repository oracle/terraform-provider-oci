resource "oci_identity_tag_namespace" "validator_bcompat" {
  compartment_id = var.tenancy_ocid
  description    = "Temporary TERSI-5524 backward-compatibility validation namespace"
  name           = "tf-tersi-5524-validator-bcompat-2"
}

resource "oci_identity_tag" "validator_bcompat" {
  description      = "Temporary TERSI-5524 backward-compatibility validation tag"
  name             = "validator-bcompat"
  tag_namespace_id = oci_identity_tag_namespace.validator_bcompat.id

  validator {
    validator_type = "ENUM"
    values         = ["value1", "value2"]
  }
}
