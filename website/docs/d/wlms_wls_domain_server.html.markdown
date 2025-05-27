---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_server"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_server"
description: |-
  Provides details about a specific Wls Domain Server in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_server
This data source provides details about a specific Wls Domain Server resource in Oracle Cloud Infrastructure Wlms service.

Gets information about the specified server in a WebLogic domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_server" "test_wls_domain_server" {
	#Required
	server_id = oci_wlms_server.test_server.id
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `server_id` - (Required) The unique identifier of a server.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

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

