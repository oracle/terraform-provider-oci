# oci_file_storage_mount_target

## MountTarget Resource

### MountTarget Reference

The following attributes are exported:

* `availability_domain` - The availability domain the mount target is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the mount target.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My mount target` 
* `export_set_id` - The OCID of the associated export set. Controls what file systems will be exported through Network File System (NFS) protocol on this mount target. 
* `id` - The OCID of the mount target.
* `lifecycle_details` - Additional information about the current 'lifecycleState'.
* `private_ip_ids` - The OCIDs of the private IP addresses associated with this mount target.
* `state` - The current state of the mount target.
* `subnet_id` - The OCID of the subnet the mount target is in.
* `time_created` - The date and time the mount target was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new mount target in the specified compartment and
subnet. You can associate a file system with a mount
target only when they exist in the same availability domain. Instances
can connect to mount targets in another availablity domain, but
you might see higher latency than with instances in the same
availability domain as the mount target.

Mount targets have one or more private IP addresses that you can
provide as the host portion of remote target parameters in
client mount commands. These private IP addresses are listed
in the privateIpIds property of the mount target and are highly available. Mount
targets also consume additional IP addresses in their subnet. 
Do not use /30 or smaller subnets for mount target creation because they
do not have sufficient available IP addresses. 
Allow at least three IP addresses for each mount target. 

For information about access control and compartments, see
[Overview of the IAM
Service](/Content/Identity/Concepts/overview.htm).

For information about availability domains, see [Regions and
Availability Domains](/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the
`ListAvailabilityDomains` operation in the Identity and Access
Management Service API.

All Oracle Cloud Infrastructure Services resources, including
mount targets, get an Oracle-assigned, unique ID called an
Oracle Cloud Identifier (OCID).  When you create a resource,
you can find its OCID in the response. You can also retrieve a
resource's OCID by using a List API operation on that resource
type, or by viewing the resource in the Console.


The following arguments are supported:

* `availability_domain` - (Required) The availability domain in which to create the mount target.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment in which to create the mount target.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My mount target` 
* `hostname_label` - (Optional) The hostname for the mount target's IP address, used for DNS resolution. The value is the hostname portion of the private IP address's fully qualified domain name (FQDN). For example, `files-1` in the FQDN `files-1.subnet123.vcn1.oraclevcn.com`. Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `files-1` 
* `ip_address` - (Optional) A private IP address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.  Example: `10.0.3.3` 
* `subnet_id` - (Required) The OCID of the subnet in which to create the mount target. 


### Update Operation
Updates the specified mount target's information.

The following arguments support updates:
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My mount target` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_file_storage_mount_target" "test_mount_target" {
	#Required
	availability_domain = "${var.mount_target_availability_domain}"
	compartment_id = "${var.compartment_id}"
	subnet_id = "${oci_file_storage_subnet.test_subnet.id}"

	#Optional
	display_name = "${var.mount_target_display_name}"
	hostname_label = "${var.mount_target_hostname_label}"
	ip_address = "${var.mount_target_ip_address}"
}
```

# oci_file_storage_mount_targets

## MountTarget DataSource

Gets a list of mount_targets.

### List Operation
Lists the mount target resources in the specified compartment.

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `export_set_id` - (Optional) The OCID of the export set.
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


The following attributes are exported:

* `mount_targets` - The list of mount_targets.

### Example Usage

```hcl
data "oci_file_storage_mount_targets" "test_mount_targets" {
	#Required
	availability_domain = "${var.mount_target_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.mount_target_display_name}"
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	id = "${var.mount_target_id}"
	state = "${var.mount_target_state}"
}
```