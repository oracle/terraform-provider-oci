variable "identity_provider_defined_tags_value" {
  default = "value"
}

variable "identity_provider_description" {
  default = "created by terraform"
}

variable "identity_provider_freeform_tags" {
  type = "map"

  default = {
    Department = "Finance"
  }
}

variable "identity_provider_metadata" {
  default = ""
}

variable "identity_provider_metadata_file" {
  default = "sampleFederationMetadata.xml"
}

variable "identity_provider_metadata_url" {
  default = "metadataUrl"
}

variable "identity_provider_name" {
  default = "test-idp-saml2-adfs"
}

variable "identity_provider_product_type" {
  default = "ADFS"
}

variable "identity_provider_protocol" {
  default = "SAML2"
}

variable "idp_group_mapping_idp_group_name" {
  default = "test-idp-group-name"
}
