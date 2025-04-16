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

Creates a container instance and deploys the containers on it.


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
		security_context {

			#Optional
			capabilities {

				#Optional
				add_capabilities = var.container_instance_containers_security_context_capabilities_add_capabilities
				drop_capabilities = var.container_instance_containers_security_context_capabilities_drop_capabilities
			}
			is_non_root_user_check_enabled = var.container_instance_containers_security_context_is_non_root_user_check_enabled
			is_root_file_system_readonly = var.container_instance_containers_security_context_is_root_file_system_readonly
			run_as_group = var.container_instance_containers_security_context_run_as_group
			run_as_user = var.container_instance_containers_security_context_run_as_user
			security_context_type = var.container_instance_containers_security_context_security_context_type
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

* `availability_domain` - (Required) The availability domain where the container instance runs.
* `compartment_id` - (Required) (Updatable) The compartment OCID.
* `container_restart_policy` - (Optional) Container restart policy
* `containers` - (Required) The containers to create on this container instance.
	* `arguments` - (Optional) A list of string arguments for a container's ENTRYPOINT process.

		Many containers use an ENTRYPOINT process pointing to a shell (/bin/bash). For those containers, this argument list specifies the main command in the container process.

		The total size of all arguments combined must be 64 KB or smaller. 
	* `command` - (Optional) An optional command that overrides the ENTRYPOINT process. If you do not provide a value, the existing ENTRYPOINT process defined in the image is used. 
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. 
	* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. If you don't provide a name, a name is generated automatically. 
	* `environment_variables` - (Optional) A map of additional environment variables to set in the environment of the container's ENTRYPOINT process. These variables are in addition to any variables already defined in the container's image.

		The total size of all environment variables combined, name and values, must be 64 KB or smaller. 
	* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `health_checks` - (Optional) list of container health checks to check container status and take appropriate action if container status is failed. There are two types of health checks that we currently support HTTP and TCP. 
		* `failure_action` - (Optional) The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default action is KILL. If failure action is KILL, the container will be subject to the container restart policy. 
		* `failure_threshold` - (Optional) Number of consecutive failures at which we consider the check failed.
		* `headers` - (Applicable when health_check_type=HTTP) Container health check HTTP headers.
			* `name` - (Required when health_check_type=HTTP) Container HTTP header Key.
			* `value` - (Required when health_check_type=HTTP) Container HTTP header value.
		* `health_check_type` - (Required) Container health check type.
		* `initial_delay_in_seconds` - (Optional) The initial delay in seconds before start checking container health status.
		* `interval_in_seconds` - (Optional) Number of seconds between two consecutive runs for checking container health.
		* `name` - (Optional) Health check name.
		* `path` - (Required when health_check_type=HTTP) Container health check HTTP path.
		* `port` - (Required) Container health check HTTP port.
		* `success_threshold` - (Optional) Number of consecutive successes at which we consider the check succeeded again after it was in failure state.
		* `timeout_in_seconds` - (Optional) Length of waiting time in seconds before marking health check failed.
	* `image_url` - (Required) A URL identifying the image that the container runs in, such as docker.io/library/busybox:latest. If you do not provide a tag, the tag will default to latest.

		If no registry is provided, will default the registry to public docker hub `docker.io/library`.

		The registry used for container image must be reachable over the Container Instance's VNIC. 
	* `is_resource_principal_disabled` - (Optional) Determines if the container will have access to the container instance resource principal.

		This method utilizes resource principal version 2.2. For information on how to use the exposed resource principal elements, see https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_authentication_methods.htm#sdk_authentication_methods_resource_principal. 
	* `resource_config` - (Optional) The size and amount of resources available to the container. 
		* `memory_limit_in_gbs` - (Optional) The maximum amount of memory that can be consumed by the container's process.

			If you do not set a value, then the process may use all available memory on the instance. 
		* `vcpus_limit` - (Optional) The maximum amount of CPUs that can be consumed by the container's process.

			If you do not set a value, then the process can use all available CPU resources on the instance.

			CPU usage is defined in terms of logical CPUs. This means that the maximum possible value on an E3 ContainerInstance with 1 OCPU is 2.0.

			A container with a 2.0 vcpusLimit could consume up to 100% of the CPU resources available on the container instance. Values can be fractional. A value of "1.5" means that the container can consume at most the equivalent of 1 and a half logical CPUs worth of CPU capacity. 
	* `security_context` - (Optional) Security context for container.
		* `capabilities` - (Optional) Linux Container capabilities to configure capabilities of container. 
			* `add_capabilities` - (Optional) A list of additional configurable container capabilities. 
			* `drop_capabilities` - (Optional) A list of container capabilities that can be dropped. 
		* `is_non_root_user_check_enabled` - (Optional) Indicates if the container must run as a non-root user. If true, the service validates the container image at runtime to ensure that it is not going to run with UID 0 (root) and fails the container instance creation if the validation fails. 
		* `is_root_file_system_readonly` - (Optional) Determines if the container will have a read-only root file system. Default value is false.
		* `run_as_group` - (Optional) The group ID (GID) to run the entrypoint process of the container. Uses runtime default if not provided.
		* `run_as_user` - (Optional) The user ID (UID) to run the entrypoint process of the container. Defaults to user specified UID in container image metadata if not provided. This must be provided if runAsGroup is provided. 
		* `security_context_type` - (Optional) The type of security context
	* `volume_mounts` - (Optional) List of the volume mounts. 
		* `is_read_only` - (Optional) Whether the volume was mounted in read-only mode. By default, the volume is not read-only.
		* `mount_path` - (Required) The volume access path.
		* `partition` - (Optional) If there is more than one partition in the volume, reference this number of partitions. Here is an example: Number  Start   End     Size    File system  Name                  Flags 1      1049kB  106MB   105MB   fat16        EFI System Partition  boot, esp 2      106MB   1180MB  1074MB  xfs 3      1180MB  50.0GB  48.8GB                                     lvm 
		* `sub_path` - (Optional) A subpath inside the referenced volume.
		* `volume_name` - (Required) The name of the volume. Avoid entering confidential information.
	* `working_directory` - (Optional) The working directory within the container's filesystem for the container process. If not specified, the default working directory from the image is used. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. If you don't provide a name, a name is generated automatically. 
* `dns_config` - (Optional) Allow customers to define DNS settings for containers. If this is not provided, the containers use the default DNS settings of the subnet. 
	* `nameservers` - (Optional) IP address of a name server that the resolver should query, either an IPv4 address (in dot notation), or an IPv6 address in colon (and possibly dot) notation. If null, uses nameservers from subnet dhcpDnsOptions. 
	* `options` - (Optional) Options allows certain internal resolver variables to be modified. Options are a list of objects in https://man7.org/linux/man-pages/man5/resolv.conf.5.html. Examples: ["ndots:n", "edns0"]. 
	* `searches` - (Optional) Search list for host-name lookup. If null, we will use searches from subnet dhcpDnsOptios.
* `fault_domain` - (Optional) The fault domain where the container instance runs. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `graceful_shutdown_timeout_in_seconds` - (Optional) The amount of time that processes in a container have to gracefully end when the container must be stopped. For example, when you delete a container instance. After the timeout is reached, the processes are sent a signal to be deleted.
* `image_pull_secrets` - (Optional) The image pulls secrets so you can access private registry to pull container images.
	* `password` - (Required when secret_type=BASIC) The password which should be used with the registry for authentication. The value is expected in base64 format.
	* `registry_endpoint` - (Required) The registry endpoint of the container image.
	* `secret_id` - (Required when secret_type=VAULT) The OCID of the secret for registry credentials.
	* `secret_type` - (Required) The type of ImagePullSecret.
	* `username` - (Required when secret_type=BASIC) The username which should be used with the registry for authentication. The value is expected in base64 format.
* `shape` - (Required) The shape of the container instance. The shape determines the resources available to the container instance.
* `shape_config` - (Required) The size and amount of resources available to the container instance. 
	* `memory_in_gbs` - (Optional) The total amount of memory available to the container instance (GB). 
	* `ocpus` - (Required) The total number of OCPUs available to the container instance. 
* `vnics` - (Required) The networks available to containers on this container instance.
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. 
	* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
	* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP. Used for DNS. 
	* `is_public_ip_assigned` - (Optional) Whether the VNIC should be assigned a public IP address. 
	* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. 
	* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. 
	* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. 
	* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in. 
* `volumes` - (Optional) A volume is a directory with data that is accessible across multiple containers in a container instance.

	You can attach up to 32 volumes to single container instance. 
	* `backing_store` - (Applicable when volume_type=EMPTYDIR) The volume type of the empty directory, can be either File Storage or Memory.
	* `configs` - (Applicable when volume_type=CONFIGFILE) Contains key value pairs which can be mounted as individual files inside the container. The value needs to be base64 encoded. It is decoded to plain text before the mount. 
		* `data` - (Required when volume_type=CONFIGFILE) The base64 encoded contents of the file. The contents are decoded to plain text before mounted as a file to a container inside container instance. 
		* `file_name` - (Required when volume_type=CONFIGFILE) The name of the file. The fileName should be unique across the volume. 
		* `path` - (Applicable when volume_type=CONFIGFILE) (Optional) Relative path for this file inside the volume mount directory. By default, the file is presented at the root of the volume mount path. 
	* `name` - (Required) The name of the volume. This must be unique within a single container instance. 
	* `volume_type` - (Required) The type of volume.
* `state` - (Optional) (Updatable) The target state for the Container Instance. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
	* `name` - The name of the volume. This must be unique within a single container instance. 
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

