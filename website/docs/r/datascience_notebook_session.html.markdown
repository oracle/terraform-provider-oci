---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_notebook_session"
sidebar_current: "docs-oci-resource-datascience-notebook_session"
description: |-
  Provides the Notebook Session resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_notebook_session
This resource provides the Notebook Session resource in Oracle Cloud Infrastructure Data Science service.

Creates a new notebook session.

## Example Usage

```hcl
resource "oci_datascience_notebook_session" "test_notebook_session" {
	#Required
	compartment_id = var.compartment_id
	project_id = oci_datascience_project.test_project.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.notebook_session_display_name
	freeform_tags = {"Department"= "Finance"}
	notebook_session_config_details {
		#Required
		shape = var.notebook_session_notebook_session_config_details_shape

		#Optional
		block_storage_size_in_gbs = var.notebook_session_notebook_session_config_details_block_storage_size_in_gbs
		notebook_session_shape_config_details {

			#Optional
			cpu_baseline = var.notebook_session_notebook_session_config_details_notebook_session_shape_config_details_cpu_baseline
			memory_in_gbs = var.notebook_session_notebook_session_config_details_notebook_session_shape_config_details_memory_in_gbs
			ocpus = var.notebook_session_notebook_session_config_details_notebook_session_shape_config_details_ocpus
		}
		private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
		subnet_id = oci_core_subnet.test_subnet.id
	}
	notebook_session_configuration_details {
		#Required
		shape = var.notebook_session_notebook_session_configuration_details_shape
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		block_storage_size_in_gbs = var.notebook_session_notebook_session_configuration_details_block_storage_size_in_gbs
		notebook_session_shape_config_details {

			#Optional
			cpu_baseline = var.notebook_session_notebook_session_configuration_details_notebook_session_shape_config_details_cpu_baseline
			memory_in_gbs = var.notebook_session_notebook_session_configuration_details_notebook_session_shape_config_details_memory_in_gbs
			ocpus = var.notebook_session_notebook_session_configuration_details_notebook_session_shape_config_details_ocpus
		}
		private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	}
	notebook_session_runtime_config_details {

		#Optional
		custom_environment_variables = var.notebook_session_notebook_session_runtime_config_details_custom_environment_variables
		notebook_session_git_config_details {

			#Optional
			notebook_session_git_repo_config_collection {
				#Required
				url = var.notebook_session_notebook_session_runtime_config_details_notebook_session_git_config_details_notebook_session_git_repo_config_collection_url
			}
		}
	}
	notebook_session_storage_mount_configuration_details_list {
		#Required
		destination_directory_name = var.notebook_session_notebook_session_storage_mount_configuration_details_list_destination_directory_name
		storage_type = var.notebook_session_notebook_session_storage_mount_configuration_details_list_storage_type

		#Optional
		bucket = var.notebook_session_notebook_session_storage_mount_configuration_details_list_bucket
		destination_path = var.notebook_session_notebook_session_storage_mount_configuration_details_list_destination_path
		export_id = oci_file_storage_export.test_export.id
		mount_target_id = oci_file_storage_mount_target.test_mount_target.id
		namespace = var.notebook_session_notebook_session_storage_mount_configuration_details_list_namespace
		prefix = var.notebook_session_notebook_session_storage_mount_configuration_details_list_prefix
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the notebook session.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My NotebookSession` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `notebook_session_config_details` - (Optional) Details for the notebook session configuration.
	* `block_storage_size_in_gbs` - (Optional) A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `notebook_session_shape_config_details` - (Optional) Details for the notebook session shape configuration.
		* `cpu_baseline` - (Optional) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left bank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - (Optional) The total amount of memory available to the notebook session instance, in gigabytes. 
		* `ocpus` - (Optional) The total number of OCPUs available to the notebook session instance. 
	* `private_endpoint_id` - (Optional) The OCID of a Data Science private endpoint. 
	* `shape` - (Required) The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - (Optional) A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_configuration_details` - (Optional) (Updatable) Details for the notebook session configuration.
	* `block_storage_size_in_gbs` - (Optional) (Updatable) A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `notebook_session_shape_config_details` - (Optional) (Updatable) Details for the notebook session shape configuration.
		* `cpu_baseline` - (Optional) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left bank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to the notebook session instance, in gigabytes. 
		* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the notebook session instance. 
	* `private_endpoint_id` - (Optional) (Updatable) The OCID of a Data Science private endpoint. 
	* `shape` - (Required) (Updatable) The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - (Required) (Updatable) A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_runtime_config_details` - (Optional) (Updatable) Notebook Session runtime configuration details.
	* `custom_environment_variables` - (Optional) (Updatable) Custom environment variables for Notebook Session. These key-value pairs will be available for customers in Notebook Sessions.
	* `notebook_session_git_config_details` - (Optional) (Updatable) Git configuration Details.
		* `notebook_session_git_repo_config_collection` - (Optional) (Updatable) A collection of Git repository configurations.
			* `url` - (Required) (Updatable) The repository URL
* `notebook_session_storage_mount_configuration_details_list` - (Optional) (Updatable) Collection of NotebookSessionStorageMountConfigurationDetails.
	* `bucket` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage bucket
	* `destination_directory_name` - (Required) (Updatable) The local directory name to be mounted
	* `destination_path` - (Optional) (Updatable) The local path of the mounted directory, excluding directory name.
	* `export_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the export
	* `mount_target_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the mount target
	* `namespace` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage namespace
	* `prefix` - (Applicable when storage_type=OBJECT_STORAGE) (Updatable) Prefix in the bucket to mount
	* `storage_type` - (Required) (Updatable) The type of storage.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the notebook session.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the notebook session.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My NotebookSession` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session.
* `lifecycle_details` - Details about the state of the notebook session.
* `notebook_session_config_details` - Details for the notebook session configuration.
	* `block_storage_size_in_gbs` - A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `notebook_session_shape_config_details` - Details for the notebook session shape configuration.
		* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left bank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory available to the notebook session instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the notebook session instance. 
	* `private_endpoint_id` - The OCID of a Data Science private endpoint. 
	* `shape` - The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_configuration_details` - Details for the notebook session configuration.
	* `block_storage_size_in_gbs` - A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `notebook_session_shape_config_details` - Details for the notebook session shape configuration.
		* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left bank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory available to the notebook session instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the notebook session instance. 
	* `private_endpoint_id` - The OCID of a Data Science private endpoint. 
	* `shape` - The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_runtime_config_details` - Notebook Session runtime configuration details.
	* `custom_environment_variables` - Custom environment variables for Notebook Session. These key-value pairs will be available for customers in Notebook Sessions.
	* `notebook_session_git_config_details` - Git configuration Details.
		* `notebook_session_git_repo_config_collection` - A collection of Git repository configurations.
			* `url` - The repository URL
* `notebook_session_storage_mount_configuration_details_list` - Collection of NotebookSessionStorageMountConfigurationDetails.
	* `bucket` - The object storage bucket
	* `destination_directory_name` - The local directory name to be mounted
	* `destination_path` - The local path of the mounted directory, excluding directory name.
	* `export_id` - OCID of the export
	* `mount_target_id` - OCID of the mount target
	* `namespace` - The object storage namespace
	* `prefix` - Prefix in the bucket to mount
	* `storage_type` - The type of storage.
* `notebook_session_url` - The URL to interact with the notebook session.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the notebook session.
* `state` - The state of the notebook session.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Notebook Session
	* `update` - (Defaults to 20 minutes), when updating the Notebook Session
	* `delete` - (Defaults to 20 minutes), when destroying the Notebook Session


## Import

NotebookSessions can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_notebook_session.test_notebook_session "id"
```

