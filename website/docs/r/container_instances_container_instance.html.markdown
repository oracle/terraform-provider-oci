---
subcategory: "Container Instances"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_container_instances_container_instance"
sidebar_current: "docs-oci-resource-container_instances-container_instance"
description: |-
  Provides the Container Instance resource in Oracle Cloud Infrastructure Container Instances service
---

# oci_container_instances_container_instance
This resource provides the Container Instance resource in Oracle Cloud Infrastructure Container Instances service.

Creates a new ContainerInstance.


## Example Usage

```hcl
resource "oci_container_instances_container_instance" "test_container_instance" {
	#Required
	availability_domain = var.container_instance_availability_domain
	compartment_id = var.compartment_id
	containers {
		#Required
		image_url = var.container_instance_containers_image_url

		#Optional
		additional_capabilities = var.container_instance_containers_additional_capabilities
		arguments = var.container_instance_containers_arguments
		command = var.container_instance_containers_command
		defined_tags = var.container_instance_containers_defined_tags
		display_name = var.container_instance_containers_display_name
		environment_variables = var.container_instance_containers_environment_variables
		freeform_tags = var.container_instance_containers_freeform_tags
		health_checks {
			#Required
			health_check_type = var.container_instance_containers_health_checks_health_check_type

			#Optional
			command = var.container_instance_containers_health_checks_command
			failure_action = var.container_instance_containers_health_checks_failure_action
			failure_threshold = var.container_instance_containers_health_checks_failure_threshold
			headers {

				#Optional
				name = var.container_instance_containers_health_checks_headers_name
				value = var.container_instance_containers_health_checks_headers_value
			}
			initial_delay_in_seconds = var.container_instance_containers_health_checks_initial_delay_in_seconds
			interval_in_seconds = var.container_instance_containers_health_checks_interval_in_seconds
			name = var.container_instance_containers_health_checks_name
			path = var.container_instance_containers_health_checks_path
			port = var.container_instance_containers_health_checks_port
			success_threshold = var.container_instance_containers_health_checks_success_threshold
			timeout_in_seconds = var.container_instance_containers_health_checks_timeout_in_seconds
		}
		is_resource_principal_disabled = var.container_instance_containers_is_resource_principal_disabled
		resource_config {

			#Optional
			memory_limit_in_gbs = var.container_instance_containers_resource_config_memory_limit_in_gbs
			vcpus_limit = var.container_instance_containers_resource_config_vcpus_limit
		}
		volume_mounts {
			#Required
			mount_path = var.container_instance_containers_volume_mounts_mount_path
			volume_name = var.container_instance_containers_volume_mounts_volume_name

			#Optional
			is_read_only = var.container_instance_containers_volume_mounts_is_read_only
			partition = var.container_instance_containers_volume_mounts_partition
			sub_path = var.container_instance_containers_volume_mounts_sub_path
		}
		working_directory = var.container_instance_containers_working_directory
	}
	shape = var.container_instance_shape
	shape_config {
		#Required
		ocpus = var.container_instance_shape_config_ocpus

		#Optional
		memory_in_gbs = var.container_instance_shape_config_memory_in_gbs
	}
	vnics {
		#Required
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		defined_tags = var.container_instance_vnics_defined_tags
		display_name = var.container_instance_vnics_display_name
		freeform_tags = var.container_instance_vnics_freeform_tags
		hostname_label = var.container_instance_vnics_hostname_label
		is_public_ip_assigned = var.container_instance_vnics_is_public_ip_assigned
		nsg_ids = var.container_instance_vnics_nsg_ids
		private_ip = var.container_instance_vnics_private_ip
		skip_source_dest_check = var.container_instance_vnics_skip_source_dest_check
	}

	#Optional
	container_restart_policy = var.container_instance_container_restart_policy
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.container_instance_display_name
	dns_config {

		#Optional
		nameservers = var.container_instance_dns_config_nameservers
		options = var.container_instance_dns_config_options
		searches = var.container_instance_dns_config_searches
	}
	fault_domain = var.container_instance_fault_domain
	freeform_tags = {"bar-key"= "value"}
	graceful_shutdown_timeout_in_seconds = var.container_instance_graceful_shutdown_timeout_in_seconds
	image_pull_secrets {
		#Required
		registry_endpoint = var.container_instance_image_pull_secrets_registry_endpoint
		secret_type = var.container_instance_image_pull_secrets_secret_type

		#Optional
		password = var.container_instance_image_pull_secrets_password
		secret_id = oci_vault_secret.test_secret.id
		username = var.container_instance_image_pull_secrets_username
	}
	volumes {
		#Required
		name = var.container_instance_volumes_name
		volume_type = var.container_instance_volumes_volume_type

		#Optional
		backing_store = var.container_instance_volumes_backing_store
		configs {

			#Optional
			data = var.container_instance_volumes_configs_data
			file_name = var.container_instance_volumes_configs_file_name
			path = var.container_instance_volumes_configs_path
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) Availability Domain where the ContainerInstance should be created.
* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `container_restart_policy` - (Optional) Container restart policy
* `containers` - (Required) The Containers to create on this Instance.
	* `additional_capabilities` - (Optional) A list of additional capabilities for the container. 
	* `arguments` - (Optional) A list of string arguments for a container's entrypoint process.

		Many containers use an entrypoint process pointing to a shell, for example /bin/bash. For such containers, this argument list can also be used to specify the main command in the container process.

		All arguments together must be 64KB or smaller. 
	* `command` - (Optional) This command will override the container's entrypoint process.  If not specified, the existing entrypoint process defined in the image will be used. 
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - (Optional) Display name for the Container. There are no guarantees of uniqueness for this name. If none is provided, it will be calculated automatically. 
	* `environment_variables` - (Optional) A map of additional environment variables to set in the environment of the container's entrypoint process. These variables are in addition to any variables already defined in the container's image.

		All environment variables together, name and values, must be 64KB or smaller. 
	* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `health_checks` - (Optional) list of container health checks to check container status and take appropriate action if container status is failed. There are three types of health checks that we currently support HTTP, TCP, and Command. 
		* `command` - (Required when health_check_type=COMMAND) The list of strings which will be concatenated to a single command for checking container's status. 
		* `failure_action` - (Optional) The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default action is KILL. If failure action is KILL, the container will be subject to the container restart policy. 
		* `failure_threshold` - (Optional) Number of consecutive failures at which we consider the check failed.
		* `headers` - (Applicable when health_check_type=HTTP) Container health check Http's headers.
			* `name` - (Required when health_check_type=HTTP) Container Http header Key.
			* `value` - (Required when health_check_type=HTTP) Container Http header value.
		* `health_check_type` - (Required) Container health check type.
		* `initial_delay_in_seconds` - (Optional) The initial delay in seconds before start checking container health status.
		* `interval_in_seconds` - (Optional) Number of seconds between two consecutive runs for checking container health.
		* `name` - (Optional) Health check name.
		* `path` - (Required when health_check_type=HTTP) Container health check Http's path.
		* `port` - (Required when health_check_type=HTTP | TCP) Container health check Http's port.
		* `success_threshold` - (Optional) Number of consecutive successes at which we consider the check succeeded again after it was in failure state.
		* `timeout_in_seconds` - (Optional) Length of waiting time in seconds before marking health check failed.
	* `image_url` - (Required) The container image information. Currently only support public docker registry. Can be either image name, e.g `containerImage`, image name with version, e.g `containerImage:v1` or complete docker image Url e.g `docker.io/library/containerImage:latest`. If no registry is provided, will default the registry to public docker hub `docker.io/library`. The registry used for container image must be reachable over the Container Instance's VNIC. 
	* `is_resource_principal_disabled` - (Optional) Determines if the Container will have access to the Container Instance Resource Principal.  This method utilizes resource principal version 2.2. Please refer to  https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal  for detailed explanation of how to leverage the exposed resource principal elements. 
	* `resource_config` - (Optional) The size and amount of resources available to the Container. 
		* `memory_limit_in_gbs` - (Optional) The maximum amount of memory which may be consumed by the Container's process.  If no value is provided, then the process may use all available memory on the Instance. 
		* `vcpus_limit` - (Optional) The maximum amount of CPU utilization which may be consumed by the Container's process.  If no value is provided, then the process may consume all CPU resources on the Instance.  CPU usage is defined in terms of logical CPUs. This means that the maximum possible value on  an E3 ContainerInstance with 1 OCPU is 2.0.  A Container with that vcpusLimit could consume up to 100% of the CPU resources available on the Instance.  Values may be fractional. A value of "1.5" means that the Container  may consume at most the equivalent of 1 and a half logical CPUs worth of CPU capacity 
	* `volume_mounts` - (Optional) List of the volume mounts. 
		* `is_read_only` - (Optional) Whether the volume was mounted in read-only mode. Defaults to false if not specified.
		* `mount_path` - (Required) mountPath describes the volume access path.
		* `partition` - (Optional) If there is more than 1 partitions in the volume, this is the number of partition which be referenced. Here is a example: Number  Start   End     Size    File system  Name                  Flags 1      1049kB  106MB   105MB   fat16        EFI System Partition  boot, esp 2      106MB   1180MB  1074MB  xfs 3      1180MB  50.0GB  48.8GB                                     lvm 
		* `sub_path` - (Optional) specifies a sub-path inside the referenced volume instead of its root
		* `volume_name` - (Required) The name of the volume.
	* `working_directory` - (Optional) The working directory within the Container's filesystem for the Container process. If none is set, the Container will run in the working directory set by the container image. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Human-readable name for the ContainerInstance. If none is provided, Oracle Cloud Infrastructure will select one for you. 
* `dns_config` - (Optional) Allow customers to define DNS settings for containers. If this is not provided, the containers will use the default DNS settings of the subnet. 
	* `nameservers` - (Optional) IP address of a name server that the resolver should query, either an IPv4 address (in dot notation), or an IPv6 address in colon (and possibly dot) notation. If null, we will use nameservers from subnet dhcpDnsOptions. 
	* `options` - (Optional) Options allows certain internal resolver variables to be modified. Options are a list of objects in https://man7.org/linux/man-pages/man5/resolv.conf.5.html. Examples: ["ndots:n", "edns0"] 
	* `searches` - (Optional) Search list for host-name lookup. If null, we will use searches from subnet dhcpDnsOptios.
* `fault_domain` - (Optional) Fault Domain where the ContainerInstance should run. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `graceful_shutdown_timeout_in_seconds` - (Optional) Duration in seconds processes within a Container have to gracefully terminate. This applies whenever a Container must be halted, such as when the Container Instance is deleted. Processes will first be sent a termination signal. After this timeout is reached, the processes will be sent a termination signal.
* `image_pull_secrets` - (Optional) The image pull secrets for accessing private registry to pull images for containers
	* `password` - (Required when secret_type=BASIC) The password which should be used with the registry for authentication. The value is expected in base64 format.
	* `registry_endpoint` - (Required) The registry endpoint of the container image.
	* `secret_id` - (Required when secret_type=VAULT) The OCID of the secret for registry credentials.
	* `secret_type` - (Required) The type of ImagePullSecret.
	* `username` - (Required when secret_type=BASIC) The username which should be used with the registry for authentication. The value is expected in base64 format.
* `shape` - (Required) The shape of the Container Instance. The shape determines the resources available to the Container Instance.
* `shape_config` - (Required) The size and amount of resources available to the Container Instance. 
	* `memory_in_gbs` - (Optional) The total amount of memory available to the instance, in gigabytes. 
	* `ocpus` - (Required) The total number of OCPUs available to the instance. 
* `vnics` - (Required) The networks to make available to containers on this Instance.
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
	* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP. 
	* `is_public_ip_assigned` - (Optional) Whether the VNIC should be assigned a public IP address. 
	* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. 
	* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. 
	* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. 
	* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in. 
* `volumes` - (Optional) A Volume represents a directory with data that is accessible across multiple containers in a ContainerInstance. Up to 32 volumes can be attached to single container instance. 
	* `backing_store` - (Applicable when volume_type=EMPTYDIR) Volume type that we are using for empty dir where it could be either File Storage or Memory
	* `configs` - (Applicable when volume_type=CONFIGFILE) Contains key value pairs which can be mounted as individual files inside the container. The value needs to be base64 encoded. It is decoded to plain text before the mount. 
		* `data` - (Required when volume_type=CONFIGFILE) The base64 encoded contents of the file. The contents are decoded to plain text before mounted as a file to a container inside container instance. 
		* `file_name` - (Required when volume_type=CONFIGFILE) The name of the file. The fileName should be unique across the volume. 
		* `path` - (Applicable when volume_type=CONFIGFILE) (Optional) Relative path for this file inside the volume mount directory. By default, the file is presented at the root of the volume mount path. 
	* `name` - (Required) The name of the volume. This has be unique cross single ContainerInstance. 
	* `volume_type` - (Required) The type of volume.
* `state` - (Optional) (Updatable) The target state for the Container Instance. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Container Instance
	* `update` - (Defaults to 20 minutes), when updating the Container Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Container Instance


## Import

ContainerInstances can be imported using the `id`, e.g.

```
$ terraform import oci_container_instances_container_instance.test_container_instance "id"
```

