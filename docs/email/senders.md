# oci_email_sender

## Sender Resource

### Sender Reference

The following attributes are exported:

* `email_address` - The email address of the sender.
* `id` - The unique OCID of the sender.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.us-phoenix-1.oraclecloud.com/Content/Email/Concepts/emaildeliveryoverview.htm#spf). 
* `state` - The current status of the approved sender.
* `time_created` - The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 



### Create Operation
Creates a sender for a tenancy in a given compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the sender.
* `email_address` - (Required) The email address of the sender.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_email_sender" "test_sender" {
	#Required
	compartment_id = "${var.compartment_id}"
	email_address = "${var.sender_email_address}"
}
```


## Sender Singular DataSource


### Get Operation
Gets an approved sender for a given `senderId`.

The following arguments are supported:

* `sender_id` - (Required) The unique OCID of the sender.


### Example Usage

```hcl
data "oci_email_sender" "test_sender" {
	#Required
	sender_id = "${var.sender_sender_id}"
}
```
# oci_email_senders

## Sender DataSource

Gets a list of senders.

### List Operation
Gets a collection of approved sender email addresses and sender IDs.

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment.
* `email_address` - (Optional) The email address of the approved sender.
* `state` - (Optional) The current state of a sender.


The following attributes are exported:

* `senders` - The list of senders.

### Example Usage

```hcl
data "oci_email_senders" "test_senders" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	email_address = "${var.sender_email_address}"
	state = "${var.sender_state}"
}
```