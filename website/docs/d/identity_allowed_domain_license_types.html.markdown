---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_allowed_domain_license_types"
sidebar_current: "docs-oci-datasource-identity-allowed_domain_license_types"
description: |-
  Provides the list of Allowed Domain License Types in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_allowed_domain_license_types
This data source provides the list of Allowed Domain License Types in Oracle Cloud Infrastructure Identity service.

(For tenancies that support identity domains) Lists the license types for identity domains supported by Oracle Cloud Infrastructure. 
(License types are also referred to as domain types.)

If `currentLicenseTypeName` is provided, then the request returns license types that the identity domain with the specified license 
type name can change to. Otherwise, the request returns all valid license types currently supported.


## Example Usage

```hcl
data "oci_identity_allowed_domain_license_types" "test_allowed_domain_license_types" {

	#Optional
	current_license_type_name = var.allowed_domain_license_type_current_license_type_name
}
```

## Argument Reference

The following arguments are supported:

* `current_license_type_name` - (Optional) The license type of the identity domain.


## Attributes Reference

The following attributes are exported:

* `allowed_domain_license_types` - The list of allowed_domain_license_types.

### AllowedDomainLicenseType Reference

The following attributes are exported:

* `description` - The license type description.
* `license_type` - The license type identifier.  Example: "oracle-apps-premium" 
* `name` - The license type name.  Example: "Oracle Apps Premium" 

