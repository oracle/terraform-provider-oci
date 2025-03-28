---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_backups"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_backups"
description: |-
  Provides the list of Deployment Backups in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_backups
This data source provides the list of Deployment Backups in Oracle Cloud Infrastructure Golden Gate service.

Lists the Backups in a compartment.


## Example Usage

```hcl
data "oci_golden_gate_deployment_backups" "test_deployment_backups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	display_name = var.deployment_backup_display_name
	state = var.deployment_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 
* `deployment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `state` - (Optional) A filter to return only the resources that match the 'lifecycleState' given. 


## Attributes Reference

The following attributes are exported:

* `deployment_backup_collection` - The list of deployment_backup_collection.

### DeploymentBackup Reference

The following attributes are exported:

* `backup_source_type` - Possible deployment backup source types. 
* `backup_type` - Possible Deployment backup types. 
* `bucket` - Name of the bucket where the object is to be uploaded in the object storage
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `deployment_type` - The type of deployment, which can be any one of the Allowed values.  NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.  Its use is discouraged in favor of 'DATABASE_ORACLE'. 
* `display_name` - An object's Display Name. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced. 
* `is_automatic` - True if this object is automatically created 
* `is_metadata_only` - Parameter to allow users to create backup without trails
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `namespace` - Name of namespace that serves as a container for all of your buckets
* `object` - Name of the object to be uploaded to object storage
* `ogg_version` - Version of OGG 
* `size_in_bytes` - The size of the backup stored in object storage (in bytes) 
* `state` - Possible lifecycle states. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_backup_finished` - The time of the resource backup finish. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_of_backup` - The time of the resource backup. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

