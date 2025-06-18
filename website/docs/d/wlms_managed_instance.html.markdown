---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_managed_instance"
sidebar_current: "docs-oci-datasource-wlms-managed_instance"
description: |-
  Provides details about a specific Managed Instance in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_managed_instance
This data source provides details about a specific Managed Instance resource in Oracle Cloud Infrastructure Wlms service.

Gets information about the specified managed instance.


## Example Usage

```hcl
data "oci_wlms_managed_instance" "test_managed_instance" {
	#Required
	managed_instance_id = oci_wlms_managed_instance.test_managed_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.


## Attributes Reference

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

