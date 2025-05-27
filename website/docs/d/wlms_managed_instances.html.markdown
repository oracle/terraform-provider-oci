---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_managed_instances"
sidebar_current: "docs-oci-datasource-wlms-managed_instances"
description: |-
  Provides the list of Managed Instances in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_managed_instances
This data source provides the list of Managed Instances in Oracle Cloud Infrastructure Wlms service.

Lists managed instances that match the specified compartment or managed instance OCID. Filter the list against a variety of criteria including but not limited to its name, status and compartment.


## Example Usage

```hcl
data "oci_wlms_managed_instances" "test_managed_instances" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_display_name
	id = var.managed_instance_id
	plugin_status = var.managed_instance_plugin_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns  only resources contained within the specified compartment. 
* `display_name` - (Optional) The display name.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `plugin_status` - (Optional) The plugin status of the managed instance. 


## Attributes Reference

The following attributes are exported:

* `managed_instance_collection` - The list of managed_instance_collection.

### ManagedInstance Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `configuration` - The configuration for a managed instance. 
	* `discovery_interval` - Frequency of domain discovery to be run on the managed instance. The unit is in hours.
	* `domain_search_paths` - The whitelisted paths which domain discovery are run against.
* `display_name` - A user-friendly name that does not have to be unique and is changeable.
* `host_name` - The FQDN of the managed instance.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `os_arch` - The operating system architecture on the managed instance.
* `os_name` - The operating system name on the managed instance.
* `plugin_status` - The plugin status of the managed instance.
* `server_count` - The number of servers running in the managed instance.
* `time_created` - The date and time the managed instance was first reported (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the managed instance was last report (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 

