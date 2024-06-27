---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegation_control_resources"
sidebar_current: "docs-oci-datasource-delegate_access_control-delegation_control_resources"
description: |-
  Provides the list of Delegation Control Resources in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_delegation_control_resources
This data source provides the list of Delegation Control Resources in Oracle Cloud Infrastructure Delegate Access Control service.

Returns a list of resources associated with the Delegation Control.


## Example Usage

```hcl
data "oci_delegate_access_control_delegation_control_resources" "test_delegation_control_resources" {
	#Required
	delegation_control_id = oci_delegate_access_control_delegation_control.test_delegation_control.id
}
```

## Argument Reference

The following arguments are supported:

* `delegation_control_id` - (Required) unique Delegation Control identifier


## Attributes Reference

The following attributes are exported:

* `delegation_control_resource_collection` - The list of delegation_control_resource_collection.

### DelegationControlResource Reference

The following attributes are exported:

* `items` - List of DelegationControlResourceSummary objects.
	* `id` - OCID of the resource.
	* `resource_status` - The current status of the resource in Delegation Control.

