---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace"
sidebar_current: "docs-oci-resource-log_analytics-namespace"
description: |-
  Provides the Namespace resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace
This resource provides the Namespace resource in Oracle Cloud Infrastructure Log Analytics service.

Onboards a tenancy with Log Analytics or Offboards a tenancy from Log Analytics functionality.

## Example Usage

```hcl
resource "oci_log_analytics_namespace" "test_namespace" {
	#Required
	compartment_id = var.compartment_id
	is_onboarded = var.is_onboarded
	namespace = var.namespace_namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment i.e. OCID of the tenancy.
* `is_onboarded` - (Required) Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
* `namespace` - (Required) The Log Analytics namespace used for the request. 

## Attributes Reference

The following attributes are exported:
* `compartment_id` - (Required) The OCID of the root compartment i.e. OCID of the tenancy.
* `is_onboarded` - (Required) Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
* `namespace` - (Required) The Log Analytics namespace used for the request.  `display_name` - A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 

## Import

Namespace can be imported using the `compartment_id` and `namespace`, e.g.

```
$ terraform import oci_log_analytics_namespace.test_namespace "compartmentId/{compartment_id}/namespace/{namespace}"
```

