---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_vm_instances"
sidebar_current: "docs-oci-datasource-datacc-vm_instances"
description: |-
  Provides the list of Vm Instances in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_vm_instances
This data source provides the list of Vm Instances in Oracle Cloud Infrastructure Datacc service.

Obtain a list of VM instances.


## Example Usage

```hcl
data "oci_datacc_vm_instances" "test_vm_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	base_server_id = oci_datacc_base_server.test_base_server.id
	display_name = var.vm_instance_display_name
	infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
	state = var.vm_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `base_server_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure Server Id.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided, it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied). 
* `display_name` - (Optional) A filter to return resources that match the entire display name given. The match is case sensitive.
* `infrastructure_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
* `state` - (Optional) A filter to return resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `vm_instance_collection` - The list of vm_instance_collection.

### VmInstance Reference

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

