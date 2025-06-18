---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain"
sidebar_current: "docs-oci-datasource-wlms-wls_domain"
description: |-
  Provides details about a specific Wls Domain in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain
This data source provides details about a specific Wls Domain resource in Oracle Cloud Infrastructure Wlms service.

Gets a specific WebLogic domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain" "test_wls_domain" {
	#Required
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `configuration` - The WebLogic domain configuration. 
	* `admin_server_control_mode` - Whether to manage the admin server using Node Manager or scripts.
	* `admin_server_start_script_path` - Path to admin server start script.
	* `admin_server_stop_script_path` - Path to admin server stop script.
	* `is_patch_enabled` - Whether or not the WebLogic domain is enabled for patching.
	* `is_rollback_on_failure` - Whether or not to rollback on failure during patching of WebLogic domain.
	* `managed_server_control_mode` - Whether to manage the managed server using Node Manager or scripts.
	* `managed_server_start_script_path` - Path to managed server start script.
	* `managed_server_stop_script_path` - Path to managed server stop script.
	* `servers_shutdown_timeout` - Servers shutdown timeout.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name that does not have to be unique and is changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
* `is_accepted_terms_and_conditions` - Whether or not the terms of use agreement has been accepted for the WebLogic domain.
* `lifecycle_details` - A message that describes the current state of the WebLogic domain in more detail. For example, it can be used to provide actionable information for a resource in the Failed state. 
* `middleware_type` - The middleware type on the administration server of the WebLogic domain.
* `patch_readiness_status` - The patch readiness status of the WebLogic domain.
* `state` - The current state of the WebLogic service domain.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the WebLogic domain was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the WebLogic domain was updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `weblogic_version` - The version of the WebLogic domain.

