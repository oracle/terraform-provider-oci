---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_service_provider_action"
sidebar_current: "docs-oci-datasource-delegate_access_control-service_provider_action"
description: |-
  Provides details about a specific Service Provider Action in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_service_provider_action
This data source provides details about a specific Service Provider Action resource in Oracle Cloud Infrastructure Delegate Access Control service.

Gets the Service Provider Action associated with the specified Service Provider Action ID.

## Example Usage

```hcl
data "oci_delegate_access_control_service_provider_action" "test_service_provider_action" {
	#Required
	service_provider_action_id = oci_delegate_access_control_service_provider_action.test_service_provider_action.id
}
```

## Argument Reference

The following arguments are supported:

* `service_provider_action_id` - (Required) Unique Oracle supplied identifier associated with the Service Provider Action.


## Attributes Reference

The following attributes are exported:

* `component` - Name of the infrastructure layer associated with the Service Provider Action.
* `customer_display_name` - Display Name of the Service Provider Action.
* `description` - Description of the Service Provider Action in terms of associated risk profile, and characteristics of the operating system commands made available to the support operator under this Service Provider Action. 
* `id` - Unique Oracle assigned identifier for the Service Provider Action.
* `name` - Unique name of the Service Provider Action.
* `properties` - Fine grained properties associated with the Delegation Control.
	* `name` - Name of the property
	* `value` - value of the property
* `resource_type` - resourceType for which the ServiceProviderAction is applicable
* `service_provider_service_types` - List of Service Provider Service Types that this Service Provider Action is applicable to.
* `state` - The current lifecycle state of the Service Provider Action.

