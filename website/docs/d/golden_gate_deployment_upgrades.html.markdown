---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_upgrades"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_upgrades"
description: |-
  Provides the list of Deployment Upgrades in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_upgrades
This data source provides the list of Deployment Upgrades in Oracle Cloud Infrastructure Golden Gate service.

Lists the Deployment Upgrades in a compartment.


## Example Usage

```hcl
data "oci_golden_gate_deployment_upgrades" "test_deployment_upgrades" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	display_name = var.deployment_upgrade_display_name
	state = var.deployment_upgrade_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `deployment_id` - (Optional) The ID of the deployment in which to list resources. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `state` - (Optional) A filter to return only the resources that match the 'lifecycleState' given. 


## Attributes Reference

The following attributes are exported:

* `deployment_upgrade_collection` - The list of deployment_upgrade_collection.

### DeploymentUpgrade Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `deployment_upgrade_type` - The type of the deployment upgrade: MANUAL or AUTOMATIC 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment upgrade being referenced. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `lifecycle_sub_state` - Possible GGS lifecycle sub-states. 
* `ogg_version` - Version of OGG 
* `state` - Possible lifecycle states. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_finished` - The date and time the request was finished. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_started` - The date and time the request was started. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

