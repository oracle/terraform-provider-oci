
# oci\_core\_instances

[Instance Reference][fa44b1ae]

  [fa44b1ae]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/ "InstanceReference"

Gets a list of instances.

## Example Usage

```
resource "oci_core_instance" "testInstance" {
	#Required
	availability_domain = "${var.availability_domain}"
	compartment_id = "${var.compartment_id}"
	image = "${var.image_id}"
	shape = "${var.shape}"

	# Optional
	create_vnic_details {
		#Required (subnet_id may also be specified at the root level)
		subnet_id = "${var.create_vnic_details_subnet_id}"

		#Optional
		assign_public_ip = "${var.create_vnic_details_assign_public_ip}"
		display_name = "${var.create_vnic_details_display_name}"
		hostname_label = "${var.create_vnic_details_hostname_label}"
		private_ip = "${var.create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.skip_source_dest_check}"
	}
	display_name = "${var.display_name}"
	hostname_label = "${var.hostname_label}"
	subnet_id = "${var.create_vnic_details_subnet_id}"
	ipxe_script = "${var.ipxe_script}"
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}
	extended_metadata {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	}
}

```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `create_vnic_details` - (Optional) Details for creating a new VNIC. See [Create Vnic Details](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/CreateVnicDetails).
* `shape` - (Required) The shape of an instance.
* `subnet_id` - (Optional) The OCID of the subnet. This must be specified either here or in `create_vnic_details`.
* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP. This may be specified either here or in `create_vnic_details`.
* `availability_domain` - (Optional) The name of the Availability Domain.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `image` - (Required) The OCID of the image used to boot the instance.
* `ipxe_script` - (Optional) This is an advanced option. See the [instance API reference](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/) for details.
* `metadata` - (Optional) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.
* `extended_metadata` - (Optional) Like metadata but allows nested metadata if you pass a valid JSON string as a value

## Create VNIC Details Argument Reference

* `assign_public_ip` - (Optional) Whether the VNIC should be assigned a public IP address.
* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information.
* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP.
* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC.
* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in.
* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip).

## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

## Instance Reference
* `availability_domain` - The Availability Domain the instance is running in.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the instance.
* `image_id` - The image used to boot the instance. You can enumerate all available images by calling `ListImages`.
* `state` - The current state of the instance: [PROVISIONING, RUNNING, STARTING, STOPPING, STOPPED, CREATING_IMAGE, TERMINATING, TERMINATED]
* `metadata` - Custom metadata that you provide.
* `extended_metadata` - Custom nested metadata that you provide. If you pass in a valid JSON string as a value then it will be converted to a JSON object; otherwise we will take the string value.
* `region` - The region that contains the Availability Domain the instance is running in.
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance.
* `time_created` - The date and time the instance was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `public_ip` - The public IP address of instance vnic (if enabled).
* `private_ip` - The private IP address of instance vnic. To set the private IP address, use the `private_ip` argument in create_vnic_details.
