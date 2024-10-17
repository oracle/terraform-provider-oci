---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_database_registration"
sidebar_current: "docs-oci-resource-golden_gate-database_registration"
description: |-
  Provides the Database Registration resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_database_registration
This resource provides the Database Registration resource in Oracle Cloud Infrastructure Golden Gate service.

Note: Deprecated. Use the /connections API instead.
Creates a new DatabaseRegistration.


## Example Usage

```hcl
resource "oci_golden_gate_database_registration" "test_database_registration" {
	#Required
	alias_name = var.database_registration_alias_name
	compartment_id = var.compartment_id
	display_name = var.database_registration_display_name
	fqdn = var.database_registration_fqdn
	password = var.database_registration_password
	username = var.database_registration_username

	#Optional
	connection_string = var.database_registration_connection_string
	database_id = oci_database_database.test_database.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.database_registration_description
	freeform_tags = {"bar-key"= "value"}
	ip_address = var.database_registration_ip_address
	key_id = oci_kms_key.test_key.id
	secret_compartment_id = oci_identity_compartment.test_compartment.id
	session_mode = var.database_registration_session_mode
	subnet_id = oci_core_subnet.test_subnet.id
	vault_id = oci_kms_vault.test_vault.id
	wallet = var.database_registration_wallet
}
```

## Argument Reference

The following arguments are supported:

* `alias_name` - (Required) (Updatable) Credential store alias. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_string` - (Optional) (Updatable) Connect descriptor or Easy Connect Naming method used to connect to a database. 
* `database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name. 
* `fqdn` - (Required) (Updatable) A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `ip_address` - (Optional) The private IP address in the customer's VCN of the customer's endpoint, typically a database. 
* `key_id` - (Optional) Refers to the customer's master key OCID.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `password` - (Required) (Updatable) The password Oracle GoldenGate uses to connect the associated system of the given technology. It must conform to the specific security requirements including length, case sensitivity, and so on. 
* `secret_compartment_id` - (Optional) The OCID of the compartment where the GoldenGate Secret will be created.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `session_mode` - (Optional) (Updatable) The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection. 
* `username` - (Required) (Updatable) The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it. 
* `vault_id` - (Optional) Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault. 
* `wallet` - (Optional) (Updatable) The wallet contents Oracle GoldenGate uses to make connections to a database. This attribute is expected to be base64 encoded. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alias_name` - Credential store alias. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_string` - Connect descriptor or Easy Connect Naming method used to connect to a database. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `fqdn` - A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.  Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the databaseRegistration being referenced. 
* `ip_address` - The private IP address in the customer's VCN of the customer's endpoint, typically a database. 
* `key_id` - Refers to the customer's master key OCID.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `rce_private_ip` - A Private Endpoint IP address created in the customer's subnet.  A customer database can expect network traffic initiated by GoldenGate Service from this IP address.  It can also send network traffic to this IP address, typically in response to requests from GoldenGate Service.  The customer may use this IP address in Security Lists or Network Security Groups (NSG) as needed. 
* `secret_compartment_id` - The OCID of the compartment where the GoldenGate Secret will be created.  If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key. 
* `secret_id` - The OCID of the customer's GoldenGate Service Secret.  If provided, it references a key that customers will be required to ensure the policies are established  to permit GoldenGate to use this Secret. 
* `session_mode` - The mode of the database connection session to be established by the data client. 'REDIRECT' - for a RAC database, 'DIRECT' - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT. 
* `state` - Possible lifecycle states. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target subnet of the dedicated connection. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `username` - The username Oracle GoldenGate uses to connect the associated system of the given technology. This username must already exist and be available by the system/application to be connected to and must conform to the case sensitivty requirments defined in it. 
* `vault_id` - Refers to the customer's vault OCID.  If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate to manage secrets contained within this vault. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Registration
	* `update` - (Defaults to 20 minutes), when updating the Database Registration
	* `delete` - (Defaults to 20 minutes), when destroying the Database Registration


## Import

DatabaseRegistrations can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_database_registration.test_database_registration "id"
```

