---
subcategory: "Audit"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_audit_configuration"
sidebar_current: "docs-oci-resource-audit-configuration"
description: |-
  Provides the Configuration resource in Oracle Cloud Infrastructure Audit service
---

# oci_audit_configuration
This resource provides the Configuration resource in Oracle Cloud Infrastructure Audit service.


## Example Usage

```hcl
resource "oci_audit_configuration" "test_configuration" {
	#Required
	compartment_id = var.tenancy_ocid
	retention_period_days = var.configuration_retention_period_days
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) ID of the root compartment (tenancy)
* `retention_period_days` - (Required) (Updatable) The retention period setting, specified in days. The minimum is 90, the maximum 365.  Example: `90` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `retention_period_days` - The retention period setting, specified in days. The minimum is 90, the maximum 365.  Example: `90` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Configuration
	* `update` - (Defaults to 20 minutes), when updating the Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Configuration


## Import

Import is not supported for this resource.

