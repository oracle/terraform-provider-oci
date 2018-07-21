---
layout: "oci"
page_title: "OCI: oci_audit_configuration"
sidebar_current: "docs-oci-resource-configuration"
description: |-
Creates and manages an OCI Configuration
---

# oci_audit_configuration
The `oci_audit_configuration` resource creates and manages an OCI Configuration



## Example Usage

```hcl
resource "oci_audit_configuration" "test_configuration" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  
  #Optional
  retention_period_days = "${var.configuration_retention_period_days}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) ID of the root compartment (tenancy)
* `retention_period_days` - (Optional) (Updatable) The retention period days

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `retention_period_days` - The retention period days
