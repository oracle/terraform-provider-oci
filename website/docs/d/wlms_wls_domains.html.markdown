---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domains"
sidebar_current: "docs-oci-datasource-wlms-wls_domains"
description: |-
  Provides the list of Wls Domains in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domains
This data source provides the list of Wls Domains in Oracle Cloud Infrastructure Wlms service.

Gets all WebLogic domains in a given compartment.


## Example Usage

```hcl
data "oci_wlms_wls_domains" "test_wls_domains" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.wls_domain_display_name
	id = var.wls_domain_id
	middleware_type = var.wls_domain_middleware_type
	patch_readiness_status = var.wls_domain_patch_readiness_status
	state = var.wls_domain_state
	weblogic_version = var.wls_domain_weblogic_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns  only resources contained within the specified compartment. 
* `display_name` - (Optional) The display name.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
* `middleware_type` - (Optional) A filter to return WebLogic domains based on the type of middleware of the WebLogic domain. 
* `patch_readiness_status` - (Optional) A filter to return domains based on the patch readiness status. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 
* `weblogic_version` - (Optional) A filter to return WebLogic domains based on the WebLogic version. 


## Attributes Reference

The following attributes are exported:

* `wls_domain_collection` - The list of wls_domain_collection.

### WlsDomain Reference

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

