---
subcategory: "Appmgmt Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_appmgmt_control_monitor_plugin_management"
sidebar_current: "docs-oci-resource-appmgmt_control-monitor_plugin_management"
description: |-
  Provides the Monitor Plugin Management resource in Oracle Cloud Infrastructure Appmgmt Control service
---

# oci_appmgmt_control_monitor_plugin_management
This resource provides the Monitor Plugin Management resource in Oracle Cloud Infrastructure Appmgmt Control service.

Activates Resource Plugin for compute instance identified by the instance ocid.
Stores monitored instances Id and its state. Tries to enable Resource Monitoring plugin by making
remote calls to Oracle Cloud Agent and Management Agent Cloud Service.


## Example Usage

```hcl
resource "oci_appmgmt_control_monitor_plugin_management" "test_monitor_plugin_management" {
	#Required
	monitored_instance_id = oci_appmgmt_control_monitored_instance.test_monitored_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `monitored_instance_id` - (Required) OCID of monitored instance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitor Plugin Management
	* `update` - (Defaults to 20 minutes), when updating the Monitor Plugin Management
	* `delete` - (Defaults to 20 minutes), when destroying the Monitor Plugin Management


## Import

Import is not supported for this resource.

