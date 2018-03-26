# oci_core_public_ip

## PublicIp Resource

### PublicIp Reference

The following attributes are exported:

* `availability_domain` - The public IP's Availability Domain. This property is set only for ephemeral public IPs (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the Availability Domain of the assigned private IP.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment containing the public IP. For an ephemeral public IP, this is the same compartment as the private IP's. For a reserved public IP that is currently assigned, this can be a different compartment than the assigned private IP's. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The public IP's Oracle ID (OCID).
* `ip_address` - The public IP address of the `publicIp` object.  Example: `129.146.2.1` 
* `lifetime` - Defines when the public IP is deleted and released back to Oracle's public IP pool.
  * `EPHEMERAL`: The lifetime is tied to the lifetime of its assigned private IP. The ephemeral public IP is automatically deleted when its private IP is deleted, when the VNIC is terminated, or when the instance is terminated. An ephemeral public IP must always be assigned to a private IP.  
  * `RESERVED`: You control the public IP's lifetime. You can delete a reserved public IP whenever you like. It does not need to be assigned to a private IP at all times.  
  For more information and comparison of the two types, see [Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm). 
* `private_ip_id` - The OCID of the private IP that the public IP is currently assigned to, or in the process of being assigned to. 
* `scope` - Whether the public IP is regional or specific to a particular Availability Domain.  
  * `REGION`: The public IP exists within a region and can be assigned to a private IP in any Availability Domain in the region. Reserved public IPs have `scope` = `REGION`.  
  * `AVAILABILITY_DOMAIN`: The public IP exists within the Availability Domain of the private IP it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs have `scope` = `AVAILABILITY_DOMAIN`. 
* `state` - The public IP's current state.
* `time_created` - The date and time the public IP was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a public IP. Use the `lifetime` property to specify whether it's an ephemeral or
reserved public IP. For information about limits on how many you can create, see
[Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm).

* **For an ephemeral public IP:** You must also specify a `privateIpId` with the OCID of
the primary private IP you want to assign the public IP to. The public IP is created in
the same Availability Domain as the private IP. An ephemeral public IP must always be
assigned to a private IP, and only to the *primary* private IP on a VNIC, not a secondary
private IP.

* **For a reserved public IP:** You may also optionally assign the public IP to a private
IP by specifying `privateIpId`. Or you can later assign the public IP with
[UpdatePublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/UpdatePublicIp).

**Note:** When assigning a public IP to a private IP, the private IP must not already have
a public IP with `lifecycleState` = ASSIGNING or ASSIGNED. If it does, an error is returned.

Also, for reserved public IPs, the optional assignment part of this operation is
asynchronous. Poll the public IP's `lifecycleState` to determine if the assignment
succeeded.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID. 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `lifetime` - (Required) Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm). 
* `private_ip_id` - (Optional) The OCID of the private IP to assign the public IP to.  Required for an ephemeral public IP because it must always be assigned to a private IP (specifically a *primary* private IP).  Optional for a reserved public IP. If you don't provide it, the public IP is created but not assigned to a private IP. You can later assign the public IP with [UpdatePublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/UpdatePublicIp). 


### Update Operation
Updates the specified public IP. You must specify the object's OCID. Use this operation if you want to:

* Assign a reserved public IP in your pool to a private IP.
* Move a reserved public IP to a different private IP.
* Unassign a reserved public IP from a private IP (which returns it to your pool
of reserved public IPs).
* Change the display name for a public IP (either ephemeral or reserved).

Assigning, moving, and unassigning a reserved public IP are asynchronous
operations. Poll the public IP's `lifecycleState` to determine if the operation
succeeded.

**Note:** When moving a reserved public IP, the target private IP
must not already have a public IP with `lifecycleState` = ASSIGNING or ASSIGNED. If it
does, an error is returned. Also, the initial unassignment from the original
private IP always succeeds, but the assignment to the target private IP is asynchronous and
could fail silently (for example, if the target private IP is deleted or has a different public IP
assigned to it in the interim). If that occurs, the public IP remains unassigned and its
`lifecycleState` switches to AVAILABLE (it is not reassigned to its original private IP).
You must poll the public IP's `lifecycleState` to determine if the move succeeded.

Regarding ephemeral public IPs:

* If you want to assign an ephemeral public IP to a primary private IP, use
[CreatePublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/CreatePublicIp).
* You can't move an ephemeral public IP to a different private IP.
* If you want to unassign an ephemeral public IP from its private IP, use
[DeletePublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/DeletePublicIp), which
unassigns and deletes the ephemeral public IP.

**Note:** If a public IP (either ephemeral or reserved) is assigned to a secondary private
IP (see [PrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp)), and you move that secondary
private IP to another VNIC, the public IP moves with it.

**Note:** There's a limit to the number of [public IPs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/)
a VNIC or instance can have. If you try to move a reserved public IP
to a VNIC or instance that has already reached its public IP limit, an error is
returned. For information about the public IP limits, see
[Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm).


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `private_ip_id` - The OCID of the private IP to assign the public IP to.  Required for an ephemeral public IP because it must always be assigned to a private IP (specifically a *primary* private IP).  Optional for a reserved public IP. If you don't provide it, the public IP is created but not assigned to a private IP. You can later assign the public IP with [UpdatePublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/UpdatePublicIp). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = "${var.compartment_id}"
	lifetime = "${var.public_ip_lifetime}"

	#Optional
	display_name = "${var.public_ip_display_name}"
	private_ip_id = "${oci_core_private_ip.test_private_ip.id}"
}
```

# oci_core_public_ips

## PublicIp DataSource

Gets a list of public_ips.

### List Operation
Lists either the ephemeral or reserved [PublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/) objects
in the specified compartment.

To list your reserved public IPs, set `scope` = `REGION`, and leave the
`availabilityDomain` parameter empty.

To list your ephemeral public IPs, set `scope` = `AVAILABILITY_DOMAIN`, and set the
`availabilityDomain` parameter to the desired Availability Domain. An ephemeral public IP
is always in the same Availability Domain and compartment as the private IP it's assigned to.

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `scope` - (Required) Whether the public IP is regional or specific to a particular Availability Domain.  
  * `REGION`: The public IP exists within a region and can be assigned to a private IP in any Availability Domain in the region. Reserved public IPs have `scope` = `REGION`.  
  * `AVAILABILITY_DOMAIN`: The public IP exists within the Availability Domain of the private IP it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs have `scope` = `AVAILABILITY_DOMAIN`. 


The following attributes are exported:

* `public_ips` - The list of public_ips.

### Example Usage

```
data "oci_core_public_ips" "test_public_ips" {
	#Required
	compartment_id = "${var.compartment_id}"
	scope = "${var.public_ip_scope}"

	#Optional
	availability_domain = "${var.public_ip_availability_domain}"
}
```