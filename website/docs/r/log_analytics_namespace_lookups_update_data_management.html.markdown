---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_lookups_update_data_management"
sidebar_current: "docs-oci-resource-log_analytics-namespace_lookups_update_data_management"
description: |-
  Provides the Namespace Lookups Update Data Management resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_lookups_update_data_management
This resource provides the Namespace Lookups Update Data Management resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/NamespaceLookupsUpdateDataManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics

Updates the lookup content. The csv file containing the content to be updated is passed in as binary data in the request.


## Example Usage

```hcl
resource "oci_log_analytics_namespace_lookups_update_data_management" "test_namespace_lookups_update_data_management" {
	#Required
	update_lookup_file_body = var.namespace_lookups_update_data_management_update_lookup_file_body
	lookup_name = var.namespace_lookups_update_data_management_lookup_name
	namespace = var.namespace_lookups_update_data_management_namespace

	#Optional
	char_encoding = var.namespace_lookups_update_data_management_char_encoding
	expect = var.namespace_lookups_update_data_management_expect
	is_force = var.namespace_lookups_update_data_management_is_force
}
```

## Argument Reference

The following arguments are supported:

* `update_lookup_file_body` - (Required) The file to use for the lookup update.
* `char_encoding` - (Optional) The character encoding of the uploaded file.
* `expect` - (Optional) A value of `100-continue` requests preliminary verification of the request method, path, and headers before the request body is sent. If no error results from such verification, the server will send a 100 (Continue) interim response to indicate readiness for the request body. The only allowed value for this parameter is "100-Continue" (case-insensitive). 
* `is_force` - (Optional) is force
* `lookup_name` - (Required) The name of the lookup to operate on.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Lookups Update Data Management
	* `update` - (Defaults to 20 minutes), when updating the Namespace Lookups Update Data Management
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Lookups Update Data Management


## Import

Import is not supported for NamespaceLookupsUpdateDataManagement