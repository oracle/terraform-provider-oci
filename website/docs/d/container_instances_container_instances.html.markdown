---
subcategory: "Container Instances"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_container_instances_container_instances"
sidebar_current: "docs-oci-datasource-container_instances-container_instances"
description: |-
  Provides the list of Container Instances in Oracle Cloud Infrastructure Container Instances service
---

# Data Source: oci_container_instances_container_instances
This data source provides the list of Container Instances in Oracle Cloud Infrastructure Container Instances service.

Returns a list of container instances.


## Example Usage

```hcl
data "oci_container_instances_container_instances" "test_container_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.container_instance_availability_domain
	display_name = var.container_instance_display_name
	state = var.container_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `container_instance_collection` - The list of container_instance_collection.

### ContainerInstance Reference

The following attributes are exported:

* `availability_domain` - The availability domain to place the container instance.
* `compartment_id` - The OCID of the compartment.
* `container_count` - The number of containers on the container instance.
* `container_restart_policy` - The container restart policy is applied for all containers in container instance.
* `containers` - The containers on the container instance.
	* `container_id` - The OCID of the container.
	* `display_name` - Display name for the Container.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_config` - DNS settings for containers.
	* `nameservers` - IP address of the name server..
	* `options` - Options allows certain internal resolver variables to be modified.
	* `searches` - Search list for hostname lookup.
* `fault_domain` - The fault domain to place the container instance.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `graceful_shutdown_timeout_in_seconds` - The amount of time that processes in a container have to gracefully end when the container must be stopped. For example, when you delete a container instance. After the timeout is reached, the processes are sent a signal to be deleted.
* `id` - An OCID that cannot be changed.
* `image_pull_secrets` - The image pulls secrets so you can access private registry to pull container images.
	* `registry_endpoint` - The registry endpoint of the container image.
	* `secret_id` - The OCID of the secret for registry credentials.
	* `secret_type` - The type of ImagePullSecret.
* `lifecycle_details` - A message that describes the current state of the container in more detail. Can be used to provide actionable information. 
* `shape` - The shape of the container instance. The shape determines the number of OCPUs, amount of memory, and other resources that are allocated to a container instance.
* `shape_config` - The shape configuration for a container instance. The shape configuration determines the resources thats are available to the container instance and its containers. 
	* `memory_in_gbs` - The total amount of memory available to the container instance, in gigabytes. 
	* `networking_bandwidth_in_gbps` - The networking bandwidth available to the container instance, in gigabits per second. 
	* `ocpus` - The total number of OCPUs available to the container instance. 
	* `processor_description` - A short description of the container instance's processor (CPU). 
* `state` - The current state of the container instance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. 
* `time_created` - The time the container instance was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `time_updated` - The time the container instance was updated, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `vnics` - The virtual networks available to the containers in the container instance.
	* `vnic_id` - The identifier of the virtual network interface card (VNIC) over which the containers accessing this network can communicate with the larger virtual cloud network. 
* `volume_count` - The number of volumes that are attached to the container instance.
* `volumes` - A volume is a directory with data that is accessible across multiple containers in a container instance. 
	* `backing_store` - The volume type of the empty directory, can be either File Storage or Memory.
	* `configs` - Contains string key value pairs which can be mounted as individual files inside the container. The value needs to be base64 encoded. It is decoded to plain text before the mount. 
		* `data` - The base64 encoded contents of the file. The contents are decoded to plain text before mounted as a file to a container inside container instance. 
		* `file_name` - The name of the file. The fileName should be unique across the volume. 
		* `path` - (Optional) Relative path for this file inside the volume mount directory. By default, the file is presented at the root of the volume mount path. 
	* `name` - The name of the volume. This must be unique within a single container instance. 
	* `volume_type` - The type of volume.

