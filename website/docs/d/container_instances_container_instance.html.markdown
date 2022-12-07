---
subcategory: "Container Instances"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_container_instances_container_instance"
sidebar_current: "docs-oci-datasource-container_instances-container_instance"
description: |-
  Provides details about a specific Container Instance in Oracle Cloud Infrastructure Container Instances service
---

# Data Source: oci_container_instances_container_instance
This data source provides details about a specific Container Instance resource in Oracle Cloud Infrastructure Container Instances service.

Gets a ContainerInstance by identifier

## Example Usage

```hcl
data "oci_container_instances_container_instance" "test_container_instance" {
	#Required
	container_instance_id = oci_container_instances_container_instance.test_container_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `container_instance_id` - (Required) The system-generated unique identifier for the ContainerInstance.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - Availability Domain where the ContainerInstance is running.
* `compartment_id` - Compartment Identifier
* `container_count` - The number of containers on this Instance
* `container_restart_policy` - The container restart policy is applied for all containers in container instance.
* `containers` - The Containers on this Instance
	* `container_id` - The ID of the Container on this Instance.
	* `display_name` - Display name for the Container.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Display name for the ContainerInstance. Can be renamed.
* `dns_config` - DNS settings for containers.
	* `nameservers` - Name server IP address
	* `options` - Options allows certain internal resolver variables to be modified.
	* `searches` - Search list for host-name lookup.
* `fault_domain` - Fault Domain where the ContainerInstance is running.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `graceful_shutdown_timeout_in_seconds` - Duration in seconds processes within a Container have to gracefully terminate. This applies whenever a Container must be halted, such as when the Container Instance is deleted. Processes will first be sent a termination signal. After this timeout is reached, the processes will be sent a termination signal.
* `id` - Unique identifier that is immutable on creation
* `image_pull_secrets` - The image pull secrets for accessing private registry to pull images for containers
	* `registry_endpoint` - The registry endpoint of the container image.
	* `secret_id` - The OCID of the secret for registry credentials.
	* `secret_type` - The type of ImagePullSecret.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `shape` - The shape of the Container Instance. The shape determines the resources available to the Container Instance.
* `shape_config` - The shape configuration for a Container Instance. The shape configuration determines the resources allocated to the Instance and it's containers. 
	* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
	* `networking_bandwidth_in_gbps` - The networking bandwidth available to the instance, in gigabits per second. 
	* `ocpus` - The total number of OCPUs available to the instance. 
	* `processor_description` - A short description of the instance's processor (CPU). 
* `state` - The current state of the ContainerInstance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the ContainerInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the ContainerInstance was updated. An RFC3339 formatted datetime string
* `vnics` - The virtual networks available to containers running on this Container Instance.
	* `vnic_id` - The ID of the Virtual Network Interface Card (VNIC) over which Containers accessing this network can communicate with the larger Virtual Client Network. 
* `volume_count` - The number of volumes that attached to this Instance
* `volumes` - A Volume represents a directory with data that is accessible across multiple containers in a ContainerInstance. 
	* `backing_store` - Volume type that we are using for empty dir where it could be either File Storage or Memory
	* `configs` - Contains string key value pairs which can be mounted as individual files inside the container. The value needs to be base64 encoded. It is decoded to plain text before the mount. 
		* `data` - The base64 encoded contents of the file. The contents are decoded to plain text before mounted as a file to a container inside container instance. 
		* `file_name` - The name of the file. The fileName should be unique across the volume. 
		* `path` - (Optional) Relative path for this file inside the volume mount directory. By default, the file is presented at the root of the volume mount path. 
	* `name` - The name of the volume. This has be unique cross single ContainerInstance. 
	* `volume_type` - The type of volume.

