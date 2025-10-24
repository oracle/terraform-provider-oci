---
subcategory: "Dif"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dif_stack"
sidebar_current: "docs-oci-datasource-dif-stack"
description: |-
  Provides details about a specific Stack in Oracle Cloud Infrastructure Dif service
---

# Data Source: oci_dif_stack
This data source provides details about a specific Stack resource in Oracle Cloud Infrastructure Dif service.

Gets information about a Stack.

## Example Usage

```hcl
data "oci_dif_stack" "test_stack" {
	#Required
	stack_id = oci_dif_stack.test_stack.id
}
```

## Argument Reference

The following arguments are supported:

* `stack_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Stack.


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

