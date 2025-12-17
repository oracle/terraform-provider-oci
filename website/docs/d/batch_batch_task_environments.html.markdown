---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_task_environments"
sidebar_current: "docs-oci-datasource-batch-batch_task_environments"
description: |-
  Provides the list of Batch Task Environments in Oracle Cloud Infrastructure Batch service
---

# Data Source: oci_batch_batch_task_environments
This data source provides the list of Batch Task Environments in Oracle Cloud Infrastructure Batch service.

Lists the task environments by compartment or environment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). You can filter and sort them by various properties like lifecycle state, display name and also ocid. All properties require an exact match. List operation only provides a summary information, use GetBatchTaskEnvironment to get the full details on a specific context

## Example Usage

```hcl
data "oci_batch_batch_task_environments" "test_batch_task_environments" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.batch_task_environment_display_name
	id = var.batch_task_environment_id
	state = var.batch_task_environment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task environment.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `batch_task_environment_collection` - The list of batch_task_environment_collection.

### BatchTaskEnvironment Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The batch task environment description.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. If not specified or provided as null or empty string, it be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchtaskenvironment20250914115623. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task environment.
* `image_url` - The URL of the ocir image.
* `security_context` - Security context for container runtime configuration.

	See also [docs](https://docs.oracle.com/en-us/iaas/api/#/en/container-instances/20210415/datatypes/LinuxSecurityContext). 
	* `fs_group` - A special supplemental group ID that applies to all containers in a pod.
	* `run_as_group` - Group ID for running processes inside the container.
	* `run_as_user` - User ID for running processes inside the container.
* `state` - The current state of the batch task environment. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the batch task environment was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the batch task environment was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `volumes` - List of volumes attached to the image. The use cases of the volumes are but not limited to: read the input of the task and write the output. 
	* `local_mount_directory_path` - The local path to mount the NFS share to.
	* `mount_target_export_path` - The path to the directory on the NFS server to be mounted.
	* `mount_target_fqdn` - The FQDN of the NFS server to connect to.
	* `name` - The name of the NfsVolume.
	* `type` - Discriminator for sub-entities.
* `working_directory` - Container's working directory.

