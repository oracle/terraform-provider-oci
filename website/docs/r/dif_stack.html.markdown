---
subcategory: "Dif"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dif_stack"
sidebar_current: "docs-oci-resource-dif-stack"
description: |-
  Provides the Stack resource in Oracle Cloud Infrastructure Dif service
---

# oci_dif_stack
This resource provides the Stack resource in Oracle Cloud Infrastructure Dif service.

Creates a Stack.


## Example Usage

```hcl
resource "oci_dif_stack" "test_stack" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.stack_display_name
	services = var.stack_services
	stack_templates = var.stack_stack_templates

	#Optional
	adb {
		#Required
		admin_password_id = oci_dif_admin_password.test_admin_password.id
		data_storage_size_in_tbs = var.stack_adb_data_storage_size_in_tbs
		db_version = var.stack_adb_db_version
		db_workload = var.stack_adb_db_workload
		ecpu = var.stack_adb_ecpu
		instance_id = var.stack_adb_instance_id

		#Optional
		is_mtls_connection_required = var.stack_adb_is_mtls_connection_required
		is_public = var.stack_adb_is_public
		subnet_id = oci_core_subnet.test_subnet.id
		tools_public_access = var.stack_adb_tools_public_access
		
		#Deploy Artifact fields
		artifact_object_storage_path = var.stack_adb_artifact_object_storage_path
		db_credentials {
			user_name = var.stack_adb_db_credentials_user_name
			secret_id = var.stack_adb_db_credentials_secret_id
			user_type = var.stack_adb_db_credentials_user_type
		}
	}
	dataflow {
		#Required
		driver_shape = var.stack_dataflow_driver_shape
		executor_shape = var.stack_dataflow_executor_shape
		instance_id = var.stack_dataflow_instance_id
		log_bucket_instance_id = var.stack_objectstorage_instance_id
		num_executors = var.stack_dataflow_num_executors
		spark_version = var.stack_dataflow_spark_version

		#Optional
		connections {
			#Required
			connection_details {

				#Optional
				dif_dependencies {
					#Required
					service_instance_id = var.stack_dependency_instance_id
					service_type = var.stack_dataflow_connections_connection_details_dif_dependencies_service_type
				}
				domain_names = var.stack_dataflow_connections_connection_details_domain_names
			}
			subnet_id = oci_core_subnet.test_subnet.id
		}
		driver_shape_config {
			#Required
			memory_in_gbs = var.stack_dataflow_driver_shape_config_memory_in_gbs
			ocpus = var.stack_dataflow_driver_shape_config_ocpus
		}
		executor_shape_config {
			#Required
			memory_in_gbs = var.stack_dataflow_executor_shape_config_memory_in_gbs
			ocpus = var.stack_dataflow_executor_shape_config_ocpus
		}
		private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
		warehouse_bucket_instance_id = var.stack_objectstorage_instance_id

		#Deploy Artifact fields
		execute = var.stack_dataflow_execute
		archive_uri = var.stack_dataflow_archive_uri
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	genai {
		#Required
		base_model = var.stack_genai_base_model
		cluster_type = var.stack_genai_cluster_type
		instance_id = var.stack_genai_instance_id
		oci_region = var.stack_genai_oci_region
		unit_count = var.stack_genai_unit_count

		#Optional
		endpoints {
			#Required
			endpoint_name = var.stack_genai_endpoint_name
			is_content_moderation_enabled = var.stack_genai_endpoints_is_content_moderation_enabled
		}
	}
	ggcs {
		#Required
		instance_id = var.stack_ggcs_instance_id
		ocpu = var.stack_ggcs_ocpu
		password_secret_id = oci_vault_secret.test_secret.id
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		connections {
			#Required
			connection_name = oci_golden_gate_connection.test_connection.name

			#Optional
			connection_id = oci_golden_gate_connection.test_connection.id
			dif_dependencies {
				#Required
				service_instance_id = var.stack_dependency_instance_id
				service_type = var.stack_ggcs_connections_dif_dependencies_service_type
			}
			gg_admin_secret_id = oci_vault_secret.test_secret.id
		}
		ogg_version = var.stack_ggcs_ogg_version
		public_subnet_id = oci_core_subnet.test_subnet.id

		#Deploy Artifact fields
		artifact_object_storage_path = var.stack_ggcs_artifact_object_storage_path
		users {
			user_name = var.stack_ggcs_users_user_name
			secret_id = var.stack_ggcs_users_secret_id
			user_type = var.stack_ggcs_users_user_type
			action = var.stack_ggcs_users_action
		}
		sources {
			source_id = var.stack_ggcs_sources_source_id
			action = var.stack_ggcs_sources_action
			should_start_source_operations = var.stack_ggcs_sources_should_start_source_operations
			target_uri = var.stack_ggcs_sources_target_uri
			target_connection_name = var.stack_ggcs_sources_target_connection_name
		}
		targets {
			target_id = var.stack_ggcs_target_id
			action = var.stack_ggcs_action
			should_start_target_operations = var.stack_ggcs_should_start_target_operations
			source_uri = var.stack_ggcs_source_uri
			source_connection_name = var.stack_ggcs_source_connection_name
		}
	}
	notification_email = var.stack_notification_email
	objectstorage {
		#Required
		instance_id = var.stack_objectstorage_instance_id
		object_versioning = var.stack_objectstorage_object_versioning
		storage_tier = var.stack_objectstorage_storage_tier

		#Optional
		auto_tiering = var.stack_objectstorage_auto_tiering
	}
	subnet_id = var.stack_deploy_artifacts_subnet_id
}
```

## Argument Reference

The following arguments are supported:

* `adb` - (Optional) (Updatable) ADB details if adb is included in the services.
	* `admin_password_id` - (Required) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID for admin password.
	* `data_storage_size_in_tbs` - (Required) (Updatable) The size, in terabytes, of the data volume that will be created and attached to the database.
	* `db_version` - (Required) (Updatable) A valid Oracle Database version for Autonomous Database.
	* `db_workload` - (Required) DB Workload to be used with ADB. Accepted values are OLTP, DW.
	* `ecpu` - (Required) (Updatable) The compute amount (ECPUs) available to the database.
	* `instance_id` - (Required) Id for the adw instance.
	* `is_mtls_connection_required` - (Optional) (Updatable) Specifies if the Autonomous Database requires mTLS connections.
	* `is_public` - (Optional) If true then subnetId should not be provided.
	* `subnet_id` - (Optional) The OCID of the subnet the Autonomous Database is associated with.
	* `tools_public_access` - (Optional) This is an array of CIDR (classless inter-domain routing) notations for a subnet or VCN OCID (virtual cloud network Oracle Cloud ID). Allowed only when subnetId is provided (private ADB).
    * `artifact_object_storage_path` - (Optional) Object storage path for the artifacts.
    * `db_credentials` - (Optional) DB credential details.
      * `user_name` - (Required) Username for ADB to be created or updated.
	  * `secret_id` - (Required) Vault secret OCID containing the corresponding user password.
	  * `user_type` - (Required) Type of the user. Allowed values are "ADMIN" or "CUSTOM" or "GGCS".
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the Stack in. 
* `dataflow` - (Optional) (Updatable) DATAFLOW details if dataflow is included in the services.
	* `connections` - (Optional) (Updatable) Details for connections to other services from Dataflow.
		* `connection_details` - (Required) (Updatable) Details of services to create private endpoint.
			* `dif_dependencies` - (Optional) (Updatable) List of DIF Service Dependency Details to create private endpoint.
				* `service_instance_id` - (Required) (Updatable) InstanceId of service which is part of the Stack.
				* `service_type` - (Required) (Updatable) Supported service name.
			* `domain_names` - (Optional) (Updatable) An array of DNS zone names.
		* `subnet_id` - (Required) OCID of the private subnet
	* `driver_shape` - (Required) (Updatable) The VM shape for the driver. Sets the driver cores and memory.
	* `driver_shape_config` - (Optional) (Updatable) This is used to configure the shape of the driver or executor if a flexible shape is used.
		* `memory_in_gbs` - (Required) (Updatable) The amount of memory used for the driver or executors.
		* `ocpus` - (Required) (Updatable) The total number of OCPUs used for the driver or executors. See here for details.
	* `executor_shape` - (Required) (Updatable) The VM shape for the executors. Sets the executor cores and memory.
	* `executor_shape_config` - (Optional) (Updatable) This is used to configure the shape of the driver or executor if a flexible shape is used.
		* `memory_in_gbs` - (Required) (Updatable) The amount of memory used for the driver or executors.
		* `ocpus` - (Required) (Updatable) The total number of OCPUs used for the driver or executors. See here for details.
	* `instance_id` - (Required) Id for dataflow instance
	* `log_bucket_instance_id` - (Required) (Updatable) InstanceId of log bucket created as part of objectstorage service in stack. Used for storing application run logs.
	* `num_executors` - (Required) (Updatable) The number of executor VMs requested.
	* `private_endpoint_id` - (Optional) (Updatable) OCID of the already provisioned dataflow private endpoint.
	* `spark_version` - (Required) (Updatable) The Spark version utilized to run the application.
	* `warehouse_bucket_instance_id` - (Optional) (Updatable) InstanceId of warehouse bucket created as part of objectstorage service in stack. Mandatory for SQL applications.
    * `execute` - (Optional) Contains the main file (py/jar) along with parameters & configuration to be passed to the DataFlow run.
    * `archive_uri` - (Optional) Contains the archive from object storage bucket which can be added as dependency to data flow application.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) A user-friendly name. Should be unique per compartment. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `genai` - (Optional) (Updatable) GenAi Details if genai is included in services.
	* `base_model` - (Required) Name of the base model.
	* `cluster_type` - (Required) The dedicated AI cluster type.
	* `endpoints` - (Optional) (Updatable) List of endpoints to provision for the GENAI cluster.
		* `endpoint_name` - (Required) (Updatable) Identifier for each endpoint.
		* `is_content_moderation_enabled` - (Required) (Updatable) Helps remove toxic and biased content from responses.
	* `instance_id` - (Required) (Updatable) Id for the GGCS instance to be provisioned.
	* `oci_region` - (Required) Region on which the cluster end endpoint will be provisioned.
	* `unit_count` - (Required) (Updatable) No of replicas of base model to be used for hosting.
* `ggcs` - (Optional) (Updatable) GGCS details if ggcs is included in the services.
	* `connections` - (Optional) (Updatable) Connection details to be associated with the Goldengate deployment.
		* `connection_id` - (Optional) (Updatable) OCID of pre-created Oracle GoldenGate connection.
		* `connection_name` - (Required) (Updatable) Name of the connection to be created.
		* `dif_dependencies` - (Optional) (Updatable) List of Service Dependency Details for connection creation.
			* `service_instance_id` - (Required) (Updatable) InstanceId of service which is part of the Stack.
			* `service_type` - (Required) (Updatable) Supported service name.
		* `gg_admin_secret_id` - (Optional) (Updatable) Vault secret OCID containing password that Oracle GoldenGate uses to connect the associated system of the given technology.
	* `instance_id` - (Required) Id for the GGCS instance to provision.
	* `ocpu` - (Required) (Updatable) The Minimum number of OCPUs to be made available for this Deployment.
	* `ogg_version` - (Optional) Version of OGG.
	* `password_secret_id` - (Required) The OCID of the Secret where the deployment password is stored.
	* `public_subnet_id` - (Optional) (Updatable) The OCID of a public subnet in the customer tenancy. Can be provided only for public GGCS deployments.
	* `subnet_id` - (Required) The OCID of the subnet of the GGCS deployment's private endpoint.
    * `artifact_object_storage_path` - (Optional) Object storage root path containing GGCS artifacts.
    * `users` - (Optional) Ggcs user details to be created or updated.
      * `user_name` - (Required) username for the user.
	  * `secret_id` - (Required) Vault OCID containing password for existing or new user.
	  * `user_type` - (Required) Type of GoldenGate user. Allowed values are "OPERATOR".
	  * `action` - (Required) Action to be done over the user. Allowed values are "CREATE" or "UPDATE".
    * `sources` - (Optional) Source Detail to configure existing or new datasource.
      * `source_id` - (Required) Ggcs source artifact id.
	  * `action` - (Required) Action to be done over the user. Allowed values are "CREATE" or "UPDATE".
	  * `should_start_source_operations` - (Required) Boolean value that determines source operations should start or not.
	  * `target_uri` - (Optional) Target uri for the GoldenGate deployment where distribution path needs to be configured.
      * `target_connection_name` - (Optional) Name of assigned connection for the source.
    * `targets` - (Optional) Target Detail to configure existing or new datasource.
      * `target_id` - (Required) GGCS target artifact id.
	  * `action` - (Required) Action to be done over the user. Allowed values are "CREATE" or "UPDATE".
	  * `should_start_target_operations` - (Required) Boolean value that determines target operations should start or not.
	  * `source_uri` - (Optional) Source uri for the GoldenGate deployment from where the collector path needs to be configured.
      * `source_connection_name` - (Optional) Name of assigned connection for the target.
* `notification_email` - (Optional) email id to which the stack notifications would be sent.
* `objectstorage` - (Optional) (Updatable) Object Storage Details if object storage is included in services.
	* `auto_tiering` - (Optional) (Updatable) It sets the auto-tiering status on the bucket.Allowed values are "DISABLED" / "INFREQUENTACCESS"
	* `instance_id` - (Required) (Updatable) Id for Object Storage instance to be provisioned.
	* `object_versioning` - (Required) (Updatable) Mentions whether the object versioning to be enabled or not,Allowed values are "ENABLED" / "DISABLED"/"SUSPENDED"
	* `storage_tier` - (Required) Mentions which storage tier to use for the bucket,Allowed values are "STANDARD" / "ARCHIVE"
* `services` - (Required) (Updatable) List of services to be onboarded for the stack.
* `stack_templates` - (Required) (Updatable) List of templates to be onboarded for the stack.
* `subnet_id` - (Optional) (Updatable) Subnet id for the Private Endpoint creation for artifact deployment.
* `add_service_trigger` - (Optional) (Updatable) An optional property when incremented triggers Add Service. Could be set to any integer value.
* `deploy_artifacts_trigger` - (Optional) (Updatable) An optional property when incremented triggers Deploy Artifacts. Could be set to any integer value.


** IMPORTANT **
- Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values.
- Single-operation per apply: exactly one of the following is allowed in a single terraform apply:
  - Add service (increase add_service_trigger)
  - Deploy artifacts (increase deploy_artifacts_trigger)
  - Change compartment (update compartment_id)
  - Standard update (all other updatable fields such as tags, service block parameters, etc.)
  Mixing any two or more in the same plan/apply is not allowed and will result in a provider error. Split your changes across separate applies.
- Additions to services/templates or service blocks require increasing add_service_trigger. Deletions are not permitted.
## Attributes Reference

The following attributes are exported:

* `adb` - ADB details if adb is included in the services.
	* `admin_password_id` - The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID for admin password.
	* `data_storage_size_in_tbs` - The size, in terabytes, of the data volume that will be created and attached to the database.
	* `db_version` - A valid Oracle Database version for Autonomous Database.
	* `db_workload` - DB Workload to be used with ADB. Accepted values are OLTP, DW.
	* `ecpu` - The compute amount (ECPUs) available to the database.
	* `instance_id` - Id for the adw instance.
	* `is_mtls_connection_required` - Specifies if the Autonomous Database requires mTLS connections.
	* `is_public` - If true then subnetId should not be provided.
	* `subnet_id` - The OCID of the subnet the Autonomous Database is associated with.
	* `tools_public_access` - This is an array of CIDR (classless inter-domain routing) notations for a subnet or VCN OCID (virtual cloud network Oracle Cloud ID). Allowed only when subnetId is provided (private ADB).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `dataflow` - DATAFLOW details if dataflow is included in the services.
	* `connections` - Details for connections to other services from Dataflow.
		* `connection_details` - Details of services to create private endpoint.
			* `dif_dependencies` - List of DIF Service Dependency Details to create private endpoint.
				* `service_instance_id` - InstanceId of service which is part of the Stack.
				* `service_type` - Supported service name.
			* `domain_names` - An array of DNS zone names.
		* `subnet_id` - OCID of the private subnet
	* `driver_shape` - The VM shape for the driver. Sets the driver cores and memory.
	* `driver_shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used.
		* `memory_in_gbs` - The amount of memory used for the driver or executors.
		* `ocpus` - The total number of OCPUs used for the driver or executors. See here for details.
	* `executor_shape` - The VM shape for the executors. Sets the executor cores and memory.
	* `executor_shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used.
		* `memory_in_gbs` - The amount of memory used for the driver or executors.
		* `ocpus` - The total number of OCPUs used for the driver or executors. See here for details.
	* `instance_id` - Id for dataflow instance
	* `log_bucket_instance_id` - InstanceId of log bucket created as part of objectstorage service in stack. Used for storing application run logs.
	* `num_executors` - The number of executor VMs requested.
	* `private_endpoint_id` - OCID of the already provisioned dataflow private endpoint.
	* `spark_version` - The Spark version utilized to run the application.
	* `warehouse_bucket_instance_id` - InstanceId of warehouse bucket created as part of objectstorage service in stack. Mandatory for SQL applications.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `genai` - GenAI Details if genai is included in services.
	* `base_model` - Name of the base model.
	* `cluster_type` - The dedicated AI cluster type.
	* `endpoints` - List of endpoints to provision for the GENAI cluster.
		* `endpoint_name` - Identifier for each endpoint.
		* `is_content_moderation_enabled` - Helps remove toxic and biased content from responses.
	* `instance_id` - Id for the GGCS instance to be provisioned.
	* `oci_region` - Region on which the cluster end endpoint will be provisioned.
	* `unit_count` - No of replicas of base model to be used for hosting.
* `ggcs` - GGCS details if ggcs is included in the services.
	* `connections` - Connection details to be associated with the Goldengate deployment.
		* `connection_id` - OCID of pre-created Oracle GoldenGate connection.
		* `connection_name` - Name of the connection to be created.
		* `dif_dependencies` - List of Service Dependency Details for connection creation.
			* `service_instance_id` - InstanceId of service which is part of the Stack.
			* `service_type` - Supported service name.
		* `gg_admin_secret_id` - Vault secret OCID containing password that Oracle GoldenGate uses to connect the associated system of the given technology.
	* `instance_id` - Id for the GGCS instance to provision.
	* `ocpu` - The Minimum number of OCPUs to be made available for this Deployment.
	* `ogg_version` - Version of OGG.
	* `password_secret_id` - The OCID of the Secret where the deployment password is stored.
	* `public_subnet_id` - The OCID of a public subnet in the customer tenancy. Can be provided only for public GGCS deployments.
	* `subnet_id` - The OCID of the subnet of the GGCS deployment's private endpoint.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Stack.
* `lifecycle_details` - A message that describes the current state of the Stack in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `notification_email` - email id to which the stack notifications would be sent.
* `objectstorage` - Object Storage Details if object storage is included in services.
	* `auto_tiering` - It sets the auto-tiering status on the bucket.Allowed values are "DISABLED" / "INFREQUENTACCESS"
	* `instance_id` - Id for Object Storage instance to be provisioned.
	* `object_versioning` - Mentions whether the object versioning to be enabled or not,Allowed values are "ENABLED" / "DISABLED"/"SUSPENDED"
	* `storage_tier` - Mentions which storage tier to use for the bucket,Allowed values are "STANDARD" / "ARCHIVE"
* `service_details` - Details of the service onboarded for the data intelligence stack.
	* `additional_details` - Additional details about the provisioned services
		* `assigned_connections` - connections assigned to Golden Gate deployment
			* `connection_id` - OCID of the connection.
			* `connection_name` - Name of the connection.
			* `requested_by` - Specifies who has made this connection.
		* `endpoint_details` - details of all endpoints assigned to cluster
			* `endpoint_id` - OCID of the endpoint.
			* `endpoint_name` - Identifier for each endpoint.
		* `model_id` - OCID of model
		* `model_version` - version of model
		* `oci_region` - region of cluster
		* `private_endpoint_id` - OCID of model
	* `current_artifact_path` - name of the service
	* `display_name` - name of the service
	* `instance_id` - ID for the service instance.
	* `service_id` - ID for the service
	* `service_type` - name of the cloud service
	* `service_url` - url for the service
	* `status` - state of the service
* `services` - List of services to be onboarded for the stack.
* `stack_templates` - List of templates to be onboarded for the stack.
* `state` - The current state of the Stack.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Stack was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Stack was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Stack
	* `update` - (Defaults to 12 hours), when updating the Stack
	* `delete` - (Defaults to 12 hours), when destroying the Stack
