---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_vm_instance"
sidebar_current: "docs-oci-datasource-datacc-vm_instance"
description: |-
  Provides details about a specific Vm Instance in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_vm_instance
This data source provides details about a specific Vm Instance resource in Oracle Cloud Infrastructure Datacc service.

Obtain the VM instance on Database Infrastructure that has the specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_datacc_vm_instance" "test_vm_instance" {
	#Required
	vm_instance_id = oci_datacc_vm_instance.test_vm_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `vm_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM instance.


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

