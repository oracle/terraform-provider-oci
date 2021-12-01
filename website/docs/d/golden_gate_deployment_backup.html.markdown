---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_backup"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_backup"
description: |-
  Provides details about a specific Deployment Backup in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_backup
This data source provides details about a specific Deployment Backup resource in Oracle Cloud Infrastructure Golden Gate service.

Retrieves a DeploymentBackup.


## Example Usage

```hcl
data "oci_golden_gate_deployment_backup" "test_deployment_backup" {
	#Required
	deployment_backup_id = oci_golden_gate_deployment_backup.test_deployment_backup.id
}
```

## Argument Reference

The following arguments are supported:

* `deployment_backup_id` - (Required) A unique DeploymentBackup identifier. 


## Attributes Reference

The following attributes are exported:

* `backup_type` - Possible Deployment backup types. 
* `bucket` - Name of the bucket where the object is to be uploaded in the object storage
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `display_name` - An object's Display Name. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced. 
* `is_automatic` - True if this object is automatically created 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `namespace` - Name of namespace that serves as a container for all of your buckets
* `object` - Name of the object to be uploaded to object storage
* `ogg_version` - Version of OGG 
* `size_in_bytes` - The size of the backup stored in object storage (in bytes) 
* `state` - Possible lifecycle states. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_backup_finished` - The time of the resource backup finish. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_of_backup` - The time of the resource backup. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

