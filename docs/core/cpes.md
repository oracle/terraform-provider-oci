# oci_core_cpe

## Cpe Resource

### Cpe Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the CPE.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premises router.
* `time_created` - The date and time the CPE was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new virtual Customer-Premises Equipment (CPE) object in the specified compartment. For
more information, see [IPSec VPNs](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPsec.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the CPE to reside. Notice that the CPE doesn't have to be in the same compartment as the IPSec
connection or other Networking Service components. If you're not sure which compartment to
use, put the CPE in the same compartment as the DRG. For more information about
compartments and access control, see [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must provide the public IP address of your on-premises router. See
[Configuring Your On-Premises Router for an IPSec VPN](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/configuringCPE.htm).

You may optionally specify a *display name* for the CPE, otherwise a default is provided. It does not have to
be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the CPE.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `ip_address` - (Required) The public IP address of the on-premises router.  Example: `143.19.23.16` 


### Update Operation
Updates the specified CPE's display name.
Avoid entering confidential information.


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_cpe" "test_cpe" {
	#Required
	compartment_id = "${var.compartment_id}"
	ip_address = "${var.cpe_ip_address}"

	#Optional
	display_name = "${var.cpe_display_name}"
}
```

# oci_core_cpes

## Cpe DataSource

Gets a list of cpes.

### List Operation
Lists the Customer-Premises Equipment objects (CPEs) in the specified compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


The following attributes are exported:

* `cpes` - The list of cpes.

### Example Usage

```hcl
data "oci_core_cpes" "test_cpes" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```