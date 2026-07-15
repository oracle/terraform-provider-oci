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

Gets information about the specified container instance.

## Example Usage

```hcl
data "oci_container_instances_container_instance" "test_container_instance" {
	#Required
	container_instance_id = oci_container_instances_container_instance.test_container_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `container_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container instance.


## Attributes Reference

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
* `security_context` - Security context for all containers in a container instance.
	* `fs_group` - A special supplemental group that applies to all containers in the container instance. Some volume types allow the container instance to change ownership of the volume. The owning GID will be the fsGroup, the setgid bit will be set (new files will be owned by the fsGroup), and the permission bits are OR'd with rw-rw----. If unset, the container instance will not modify the ownership and permissions of volumes. 
	* `fs_group_change_policy` - Defines behavior of changing ownership and permission of the volume before being exposed inside the containers. This only applies to volumes which support fsGroup ownership and permissions, and will have no effect on ephemeral volumes. ON_ROOT_MISMATCH only changes permissions and ownership if the permission and ownership of the root directory does not match the expected permissions and ownership of the volume. This can improve container instance start times. ALWAYS  changes permission and ownership of the volume when it is mounted. If unset, ALWAYS is used. 
	* `security_context_type` - The type of security context
* `shape` - The shape of the container instance. The shape determines the number of OCPUs, amount of memory, and other resources that are allocated to a container instance.
* `shape_config` - The shape configuration for a container instance. The shape configuration determines the resources thats are available to the container instance and its containers. 
	* `memory_in_gbs` - The total amount of memory available to the container instance, in gigabytes. 
	* `networking_bandwidth_in_gbps` - The networking bandwidth available to the container instance, in gigabits per second. 
	* `ocpus` - The total number of OCPUs available to the container instance. 
	* `processor_description` - A short description of the container instance's processor (CPU). 
* `state` - The current state of the container instance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. 
* `tenant_id` - TenantId id of the container instance.
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
	* `export` - An Oracle Cloud Infrastructure File Storage Service (FSS) Export. Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/Export/ for more details. 
		* `id` - The OCID of the Oracle Cloud Infrastructure File Storage Service (FSS) Export.
		* `oci_fss_export_type` - Determines the mode for supplying the Oracle Cloud Infrastructure File Storage Service (FSS) Export details. The value must be an OCID unless your tenancy is allowed to use PATH as a value. 
	* `mount_command` - Specifications for the mount command to mount the Oracle Cloud Infrastructure File Storage Service (FSS) File System to Containers. 
		* `mount_options` - List of mount options to be used in the mount command. The order of this array will be maintained while preparing the mount command.
			* `option` - A generic (https://man7.org/linux/man-pages/man8/mount.8.html) or nfs (https://man7.org/linux/man-pages/man5/nfs.5.html) mount option.
			* `value` - The value of the mount option.
	* `mount_target` - An Oracle Cloud Infrastructure File Storage Service (FSS) Mount Target.  Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/MountTarget for more details. 
		* `id` - The OCID of the Oracle Cloud Infrastructure File Storage Service (FSS) Mount Target.
		* `oci_fss_mount_target_type` - Determines the mode for supplying the Oracle Cloud Infrastructure File Storage Service (FSS) Mount target details. The value must be an OCID unless your tenancy is allowed to use HOST as a value. 
	* `name` - The name of the volume. This must be unique within a single container instance. 
	* `security` - Security options for Oracle Cloud Infrastructure FSS File System.
		* `auth` - NFS authentication type to be used. Currently, only auth type SYS is supported.
		* `is_encrypted_in_transit` - Determines whether in-transit encryption needs to be enables.  Check https://docs.oracle.com/en-us/iaas/Content/File/Tasks/intransitencryption.htm#Using_Intransit_Encryption for more details. 
	* `subnet_id` - Specifies the network interface to be used for the Oracle Cloud Infrastructure File Storage Service (FSS) volume. This is a required parameter when a Container Instance is attached to more than one subnets. 
	* `volume_type` - The type of volume.

