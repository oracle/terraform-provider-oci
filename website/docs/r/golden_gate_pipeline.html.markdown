---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_pipeline"
sidebar_current: "docs-oci-resource-golden_gate-pipeline"
description: |-
  Provides the Pipeline resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_pipeline
This resource provides the Pipeline resource in Oracle Cloud Infrastructure Golden Gate service.

Creates a new Pipeline.


## Example Usage

```hcl
resource "oci_golden_gate_pipeline" "test_pipeline" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.pipeline_display_name
	license_model = var.pipeline_license_model
	recipe_type = var.pipeline_recipe_type
	source_connection_details {
		#Required
		connection_id = oci_golden_gate_connection.test_connection.id
	}
	target_connection_details {
		#Required
		connection_id = oci_golden_gate_connection.test_connection.id
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.pipeline_description
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		type = var.pipeline_locks_type

		#Optional
		message = var.pipeline_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.pipeline_locks_time_created
	}
	process_options {
		#Required
		initial_data_load {
			#Required
			is_initial_load = var.pipeline_process_options_initial_data_load_is_initial_load

			#Optional
			action_on_existing_table = var.pipeline_process_options_initial_data_load_action_on_existing_table
		}
		replicate_schema_change {
			#Required
			can_replicate_schema_change = var.pipeline_process_options_replicate_schema_change_can_replicate_schema_change

			#Optional
			action_on_ddl_error = var.pipeline_process_options_replicate_schema_change_action_on_ddl_error
			action_on_dml_error = var.pipeline_process_options_replicate_schema_change_action_on_dml_error
		}
		should_restart_on_failure = var.pipeline_process_options_should_restart_on_failure

		#Optional
		start_using_default_mapping = var.pipeline_process_options_start_using_default_mapping
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name. 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `license_model` - (Required) (Updatable) The Oracle license model that applies to a Deployment. 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `process_options` - (Optional) (Updatable) Required pipeline options to configure the replication process (Extract or Replicat). 
	* `initial_data_load` - (Required) (Updatable) Options required for the pipeline Initial Data Load. If enabled, copies existing data from source to target before replication. 
		* `action_on_existing_table` - (Optional) (Updatable) Action upon existing tables in target when initial Data Load is set i.e., isInitialLoad=true. 
		* `is_initial_load` - (Required) (Updatable) If ENABLED, then existing source data is also synchronized to the target when creating or updating the pipeline. 
	* `replicate_schema_change` - (Required) (Updatable) Options required for pipeline Initial Data Load. If enabled, copies existing data from source to target before replication. 
		* `action_on_ddl_error` - (Optional) (Updatable) Action upon DDL Error (active only if 'Replicate schema changes (DDL)' is selected) i.e canReplicateSchemaChange=true 
		* `action_on_dml_error` - (Optional) (Updatable) Action upon DML Error (active only if 'Replicate schema changes (DDL)' is selected) i.e canReplicateSchemaChange=true 
		* `can_replicate_schema_change` - (Required) (Updatable) If ENABLED, then addition or removal of schema is also replicated, apart from individual tables and records when creating or updating the pipeline. 
	* `should_restart_on_failure` - (Required) (Updatable) If ENABLED, then the replication process restarts itself upon failure. This option applies when creating or updating a pipeline. 
	* `start_using_default_mapping` - (Optional) (Updatable) If ENABLED, then the pipeline is started as part of pipeline creation. It uses default mapping. This option applies when creating or updating a pipeline. 
* `recipe_type` - (Required) (Updatable) The type of the recipe 
* `source_connection_details` - (Required) The source connection details for creating a pipeline. 
	* `connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `target_connection_details` - (Required) The target connection details for creating a pipeline. 
	* `connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `cpu_core_count` - The Minimum number of OCPUs to be made available for this Deployment. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline. This option applies when retrieving a pipeline. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Deployment's CPU core count. 
* `license_model` - The Oracle license model that applies to a Deployment. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `lifecycle_sub_state` - Possible lifecycle substates when retrieving a pipeline. 
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `mapping_rules` - Mapping for source/target schema/tables for the pipeline data replication. 
	* `mapping_type` - Defines the exclude/include rules of source and target schemas and tables when replicating from source to target. This option applies when creating and updating a pipeline. 
	* `source` - The source schema/table combination for replication to target. 
	* `target` - The target schema/table combination for replication from the source. 
* `pipeline_diagnostic_data` - Information regarding the pipeline diagnostic collection 
	* `bucket` - Name of the bucket where the object is to be uploaded in the object storage
	* `diagnostic_state` - The state of the pipeline diagnostics collection. 
	* `namespace` - Name of namespace that serves as a container for all of your buckets
	* `object` - Name of the diagnostic collected and uploaded to object storage
	* `time_last_collected` - The date and time the diagnostic data was last collected for the pipeline. The format is defined by  [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`. 
* `process_options` - Required pipeline options to configure the replication process (Extract or Replicat). 
	* `initial_data_load` - Options required for the pipeline Initial Data Load. If enabled, copies existing data from source to target before replication. 
		* `action_on_existing_table` - Action upon existing tables in target when initial Data Load is set i.e., isInitialLoad=true. 
		* `is_initial_load` - If ENABLED, then existing source data is also synchronized to the target when creating or updating the pipeline. 
	* `replicate_schema_change` - Options required for pipeline Initial Data Load. If enabled, copies existing data from source to target before replication. 
		* `action_on_ddl_error` - Action upon DDL Error (active only if 'Replicate schema changes (DDL)' is selected) i.e canReplicateSchemaChange=true 
		* `action_on_dml_error` - Action upon DML Error (active only if 'Replicate schema changes (DDL)' is selected) i.e canReplicateSchemaChange=true 
		* `can_replicate_schema_change` - If ENABLED, then addition or removal of schema is also replicated, apart from individual tables and records when creating or updating the pipeline. 
	* `should_restart_on_failure` - If ENABLED, then the replication process restarts itself upon failure. This option applies when creating or updating a pipeline. 
	* `start_using_default_mapping` - If ENABLED, then the pipeline is started as part of pipeline creation. It uses default mapping. This option applies when creating or updating a pipeline. 
* `recipe_type` - The type of the recipe 
* `source_connection_details` - The source connection details for creating a pipeline. 
	* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `state` - Lifecycle state of the pipeline. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `target_connection_details` - The target connection details for creating a pipeline. 
	* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_last_recorded` - When the resource was last updated. This option applies when retrieving a pipeline. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-07-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pipeline
	* `update` - (Defaults to 20 minutes), when updating the Pipeline
	* `delete` - (Defaults to 20 minutes), when destroying the Pipeline


## Import

Pipelines can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_pipeline.test_pipeline "id"
```

