---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_service_provider"
sidebar_current: "docs-oci-datasource-delegate_access_control-service_provider"
description: |-
  Provides details about a specific Service Provider in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_service_provider
This data source provides details about a specific Service Provider resource in Oracle Cloud Infrastructure Delegate Access Control service.

Gets a ServiceProvider by identifier

## Example Usage

```hcl
data "oci_delegate_access_control_service_provider" "test_service_provider" {
	#Required
	service_provider_id = oci_delegate_access_control_service_provider.test_service_provider.id
}
```

## Argument Reference

The following arguments are supported:

* `service_provider_id` - (Required) unique Service Provider identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the Delegation Control.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the Service Provider. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - Unique identifier for the Service Provider.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `name` - Unique name of the Service Provider.
* `service_provider_type` - Service Provider type.
* `service_types` - Types of services offered by this provider.
* `state` - The current lifecycle state of the Service Provider.
* `supported_resource_types` - Resource types for which this provider will provide service. Default to all if not specified.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Service Provider was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Service Provider was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

