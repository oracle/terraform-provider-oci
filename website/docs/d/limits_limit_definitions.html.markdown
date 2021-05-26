---
subcategory: "Limits"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_limit_definitions"
sidebar_current: "docs-oci-datasource-limits-limit_definitions"
description: |-
  Provides the list of Limit Definitions in Oracle Cloud Infrastructure Limits service
---

# Data Source: oci_limits_limit_definitions
This data source provides the list of Limit Definitions in Oracle Cloud Infrastructure Limits service.

Includes a list of resource limits that are currently supported.
If the 'areQuotasSupported' property is true, you can create quota policies on top of this limit at the
compartment level.


## Example Usage

```hcl
data "oci_limits_limit_definitions" "test_limit_definitions" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	name = var.limit_definition_name
	service_name = oci_limits_service.test_service.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the parent compartment (remember that the tenancy is simply the root compartment). 
* `name` - (Optional) Optional field, filter for a specific resource limit.
* `service_name` - (Optional) The target service name.


## Attributes Reference

The following attributes are exported:

* `limit_definitions` - The list of limit_definitions.

### LimitDefinition Reference

The following attributes are exported:

* `are_quotas_supported` - If true, quota policies can be created on top of this resource limit. 
* `description` - The limit description.
* `is_deprecated` - Indicates if the limit has been deprecated. 
* `is_dynamic` - The limit for this resource has a dynamic value that is based on consumption across all Oracle Cloud Infrastructure services. 
* `is_eligible_for_limit_increase` - Indicates if the customer can request a limit increase for this resource. 
* `is_resource_availability_supported` - Reflects whether or not the GetResourceAvailability API is supported for this limit. If not, the API returns an empty JSON response. 
* `name` - The resource limit name. To be used for writing policies (in case of quotas) or other programmatic calls. 
* `scope_type` - Reflects the scope of the resource limit, whether Global (across all regions), regional, or availability domain-specific. 
* `service_name` - The service name of the limit.

