---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_enable_disable_archiving"
sidebar_current: "docs-oci-resource-log_analytics-namespace_storage_enable_disable_archiving"
description: |-
  Provides the Namespace Storage Enable Disable Archiving resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_storage_enable_disable_archiving
This resource provides the Namespace Storage Enable Disable Archiving resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/NamespaceStorageEnableDisableArchiving

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics
THis API enables archiving.


## Example Usage

```hcl
resource "oci_log_analytics_namespace_storage_enable_disable_archiving" "test_namespace_storage_enable_disable_archiving" {
	#Required
	namespace = var.namespace_storage_enable_disable_archiving_namespace
	enable_archiving_tenant = var.enable_archiving_tenant
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `enable_archiving_tenant` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `message` - A human-readable success string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Storage Enable Disable Archiving
	* `update` - (Defaults to 20 minutes), when updating the Namespace Storage Enable Disable Archiving
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Storage Enable Disable Archiving
