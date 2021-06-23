---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_connect_harness"
sidebar_current: "docs-oci-resource-streaming-connect_harness"
description: |-
  Provides the Connect Harness resource in Oracle Cloud Infrastructure Streaming service
---

# oci_streaming_connect_harness
This resource provides the Connect Harness resource in Oracle Cloud Infrastructure Streaming service.

Starts the provisioning of a new connect harness.
To track the progress of the provisioning, you can periodically call [GetConnectHarness].
In the response, the `lifecycleState` parameter of the [ConnectHarness](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/ConnectHarness/) object tells you its current state.


## Example Usage

```hcl
resource "oci_streaming_connect_harness" "test_connect_harness" {
	#Required
	compartment_id = var.compartment_id
	name = var.connect_harness_name

	#Optional
	defined_tags = var.connect_harness_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the connect harness.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the connect harness. Avoid entering confidential information.  Example: `JDBCConnector` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the connect harness.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the connect harness.
* `lifecycle_state_details` - Any additional details about the current state of the connect harness.
* `name` - The name of the connect harness. Avoid entering confidential information.  Example: `JDBCConnector` 
* `state` - The current state of the connect harness.
* `time_created` - The date and time the connect harness was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Connect Harness
	* `update` - (Defaults to 20 minutes), when updating the Connect Harness
	* `delete` - (Defaults to 20 minutes), when destroying the Connect Harness


## Import

ConnectHarnesses can be imported using the `id`, e.g.

```
$ terraform import oci_streaming_connect_harness.test_connect_harness "id"
```

