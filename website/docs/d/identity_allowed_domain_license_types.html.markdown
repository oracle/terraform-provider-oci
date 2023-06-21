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

List the allowed domain license types supported by Oracle Cloud Infrastructure If {@code currentLicenseTypeName} provided, returns allowed license types a domain with the specified license type name can migrate to.
If {@code name} is provided, it filters and returns resources that match the given license type name.
Otherwise, returns all valid license types that are currently supported.

- If license type details are retrieved sucessfully, return 202 ACCEPTED.
- If any internal error occurs, return 500 INTERNAL SERVER ERROR.


## Example Usage

```hcl
data "oci_identity_allowed_domain_license_types" "test_allowed_domain_license_types" {

	#Optional
	current_license_type_name = var.allowed_domain_license_type_current_license_type_name
}
```

## Argument Reference

The following arguments are supported:

* `current_license_type_name` - (Optional) The domain license type


## Attributes Reference

The following attributes are exported:

* `allowed_domain_license_types` - The list of allowed_domain_license_types.

### AllowedDomainLicenseType Reference

The following attributes are exported:

* `description` - The license type description.
* `license_type` - The license type identifier.  Example: "oracle-apps-premium" 
* `name` - The license type name.  Example: "Oracle Apps Premium" 

