---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_vm_instance"
sidebar_current: "docs-oci-resource-datacc-vm_instance"
description: |-
  Provides the Vm Instance resource in Oracle Cloud Infrastructure Datacc service
---

# oci_datacc_vm_instance
This resource provides the Vm Instance resource in Oracle Cloud Infrastructure Datacc service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datacc

Create an VM instance on Database Infrastructure using the specified details.


## Example Usage

```hcl
resource "oci_datacc_vm_instance" "test_vm_instance" {
	#Required
	compartment_id = var.compartment_id
	cpus_enabled = var.vm_instance_cpus_enabled
	infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
	ssh_public_keys = var.vm_instance_ssh_public_keys

	#Optional
	boot_storage_size_in_gbs = var.vm_instance_boot_storage_size_in_gbs
	data_storage_size_in_gb = var.vm_instance_data_storage_size_in_gb
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.vm_instance_description
	display_name = var.vm_instance_display_name
	dns_servers = var.vm_instance_dns_servers
	domain_name = oci_identity_domain.test_domain.name
	freeform_tags = {"bar-key"= "value"}
	gateway = var.vm_instance_gateway
	hostname = var.vm_instance_hostname
	image_id = oci_core_image.test_image.id
	ip_address = var.vm_instance_ip_address
	memory_size_in_gbs = var.vm_instance_memory_size_in_gbs
	metadata = var.vm_instance_metadata
	netmask = var.vm_instance_netmask
	ntp_servers = var.vm_instance_ntp_servers
	server_id = oci_datacc_server.test_server.id
	system_tags = var.vm_instance_system_tags
	time_zone = var.vm_instance_time_zone
	userdata = var.vm_instance_userdata
	vlan_id = oci_core_vlan.test_vlan.id
	vm_network_id = oci_datacc_vm_network.test_vm_network.id
}
```

## Argument Reference

The following arguments are supported:

* `boot_storage_size_in_gbs` - (Optional) Boot storage memory to be allocated in GBs.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM instance. 
* `cpus_enabled` - (Required) The number of CPU cores enabled for each VM instance.
* `data_storage_size_in_gb` - (Optional) (Updatable) Data storage to be allocated in GBs.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) VM instance description.
* `display_name` - (Optional) (Updatable) VM instance display name. This name does not have to be unique, and is changeable. 
* `dns_servers` - (Optional) The list of DNS server IP addresses. Maximum of 3 allowed.
* `domain_name` - (Optional) The domain name of the VM instance.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gateway` - (Optional) The gateway IP address of the VM instance network
* `hostname` - (Optional) The host name of the instance.
* `image_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM custom instance uploaded.
* `infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
* `ip_address` - (Optional) The IP address of the instance.
* `memory_size_in_gbs` - (Optional) The memory to be allocated in GBs.
* `metadata` - (Optional) Custom metadata key/value pairs which can be used to:
	* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
	* Provide additional information which is exposed inside the instance context and can be queried or referenced by user-data scripts for dynamic configuration. 
* `netmask` - (Optional) The netmask of the VM instance network.
* `ntp_servers` - (Optional) The list of NTP server IP addresses. Maximum of 3 allowed.
* `server_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute node on which VM instance should be launched.
* `ssh_public_keys` - (Required) List of public key used for SSH access to the VM instance.
* `system_tags` - (Optional) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_zone` - (Optional) The time zone to use for the VM instance.
* `userdata` - (Optional) Base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration.  For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html). 
* `vlan_id` - (Optional) The network VLAN ID.
* `vm_network_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Network.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `boot_storage_size_in_gbs` - Boot storage memory to be allocated in GBs.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM instance. 
* `cpus_enabled` - The number of CPU cores enabled for each VM instance.
* `data_storage_size_in_gb` - Data storage to be allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - VM instance description.
* `display_name` - VM instance display name. This name does not have to be unique, and is changeable. 
* `dns_servers` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `domain_name` - The domain name of the VM instance.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gateway` - The gateway IP address of the VM instance network.
* `hostname` - The host name of the instance.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM instance.
* `image_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM custom instance uploaded.
* `infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
* `ip_address` - The IP address of the instance.
* `lifecycle_details` - Lifecycle state details of the VM instance.
* `memory_size_in_gbs` - The memory to be allocated in GBs.
* `metadata` - Custom metadata key/value pairs which can be used to:
	* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
	* Provide additional information which is exposed inside the instance context and can be queried or referenced by user-data scripts for dynamic configuration. 
* `netmask` - The netmask of the VM instance network.
* `ntp_servers` - The list of NTP server addresses. Maximum of 3 allowed.
* `server_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute node on which VM instance should be launched.
* `ssh_public_keys` - List of public key used for SSH access to the VM instance.
* `state` - The current state of the VM instance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the VM instance was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time that the VM instance was last updated. An RFC3339 formatted datetime string. 
* `time_zone` - The time zone to use for the VM instance.
* `userdata` - Base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration.  For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html). 
* `vlan_id` - The network VLAN ID.
* `vm_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Network.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vm Instance
	* `update` - (Defaults to 20 minutes), when updating the Vm Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Vm Instance


## Import

VmInstances can be imported using the `id`, e.g.

```
$ terraform import oci_datacc_vm_instance.test_vm_instance "id"
```

