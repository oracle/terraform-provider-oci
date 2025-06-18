---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_managed_instance_servers"
sidebar_current: "docs-oci-datasource-wlms-managed_instance_servers"
description: |-
  Provides the list of Managed Instance Servers in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_managed_instance_servers
This data source provides the list of Managed Instance Servers in Oracle Cloud Infrastructure Wlms service.

Gets list of servers in a specific managed instance.


## Example Usage

```hcl
data "oci_wlms_managed_instance_servers" "test_managed_instance_servers" {
	#Required
	managed_instance_id = oci_wlms_managed_instance.test_managed_instance.id

	#Optional
	name = var.managed_instance_server_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `name` - (Optional) The name of the resource.


## Attributes Reference

The following attributes are exported:

* `server_collection` - The list of server_collection.

### ManagedInstanceServer Reference

The following attributes are exported:

* `host_name` - The name of the server.
* `id` - The unique identifier of the server.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `is_admin` - Whether or not the server is an admin node.
* `jdk_path` - The JDK path on the server.
* `jdk_version` - The JDK version on the server.
* `latest_patches_status` - Whether or not the server has installed the latest patches.
* `managed_instance_id` - The managed instance ID of the server.
* `middleware_path` - The middleware path on the server.
* `middleware_type` - The middleware type on the server.
* `name` - The name of the server.
* `patch_readiness_status` - The patch readiness status of the server.
* `restart_order` - The restart order assigned to the server.
* `status` - The status of the server.
* `time_created` - The date and time the server was first reported (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the server was last reported (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `weblogic_version` - The version of the WebLogic domain of the server
* `wls_domain_id` - The ID of the WebLogic domain to which the server belongs.
* `wls_domain_name` - The name of the WebLogic domain to which the server belongs.
* `wls_domain_path` - The path of the WebLogic domain to which the server belongs.

