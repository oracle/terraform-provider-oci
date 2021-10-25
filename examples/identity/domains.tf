/*
 * This example file shows how to create a domain. Change compartment of domain and toggle states
 */

variable "domain_defined_tags_value" {
  default = "value"
}

#Change the description to trigger an update operation and to call PUT /domains/{domainId}
variable "domain_description" {
  default = "descriptionCreate10"
}

variable "domain_email" {
  default = "email"
}

variable "domain_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "domain_name" {
  default = "displayNameActualCreate10"
}

#Change the licenseType to call POST /domains/{domainId}/actions/changeLicenseType, {example "premium, free"}
variable "domain_licenseType" {
  default = "premium"
}

variable "domain_homeRegion" {
default = "ca-toronto-1"
}

#Toggle between ACTIVE and INACTIVE to call POST /domains/{domainId}/actions/activate and  POST /domains/{domainId}/actions/deactivate respectively
variable "domain_state" {
  default = "ACTIVE"
}

resource "oci_identity_domain" "test_domain4" {
#Required
# Replace var.compartment_id with var.move_compartment_id to test POST /domains/{domainId}/actions/changeCompartment
  compartment_id = var.compartment_id
  description    = var.domain_description
  display_name   = var.domain_name
  license_type   = var.domain_licenseType
  home_region =  var.domain_homeRegion
  state =  var.domain_state
}

#data source for GET/domain/{domainid}
data "oci_identity_domain" "get_domain" {
#Required
domain_id = oci_identity_domain.test_domain4.id
}

#data source for allowedLicenseType call
data "oci_identity_allowed_domain_license_types" "test_allowed_domain_license_types" {
}

#data source for list domains call
data "oci_identity_domains" "list_domain" {
#Required
compartment_id = var.compartment_id
}

#output of GET /domains/{domainId}
output "domains" {
value       = oci_identity_domain.test_domain4.id
description = "Associated OCID of domain"
}

#output for allowed license_type
output "domain_license_type" {
value       = data.oci_identity_allowed_domain_license_types.test_allowed_domain_license_types.allowed_domain_license_types
description = "domain license type"
}

#output for LIST/domains
output "list_domains" {
value       = data.oci_identity_domains.list_domain.domains
description = "domains ocid"
}

