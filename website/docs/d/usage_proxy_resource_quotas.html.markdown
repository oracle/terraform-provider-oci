---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_resource_quotas"
sidebar_current: "docs-oci-datasource-usage_proxy-resource_quotas"
description: |-
  Provides the list of Resource Quotas in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_resource_quotas
This data source provides the list of Resource Quotas in Oracle Cloud Infrastructure Usage Proxy service.

Returns the resource quota details under a tenancy
> **Important**: Calls to this API will only succeed against the endpoint in the home region.


## Example Usage

```hcl
data "oci_usage_proxy_resource_quotas" "test_resource_quotas" {
	#Required
	compartment_id = var.compartment_id
	service_name = oci_core_service.test_service.name

	#Optional
	service_entitlement = var.resource_quota_service_entitlement
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `service_entitlement` - (Optional) Service entitlement Id.
* `service_name` - (Required) Service Name.


## Attributes Reference

The following attributes are exported:

* `resource_quotum_collection` - The list of resource_quotum_collection.

### ResourceQuota Reference

The following attributes are exported:

* `is_allowed` - Used to indicate if further quota consumption isAllowed.
* `items` - The list of resource quota details.
	* `affected_resource` - The affected resource name.
	* `balance` - The quota balance.
	* `is_allowed` - Used to indicate if further quota consumption isAllowed.
	* `is_dependency` - Used to indicate any resource dependencies.
	* `is_overage` - Used to indicate if overages are incurred.
	* `name` - The resource name.
	* `purchased_limit` - The purchased quota limit.
	* `service` - The service name.

