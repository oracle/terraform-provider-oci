---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_database_registrations"
sidebar_current: "docs-oci-datasource-golden_gate-database_registrations"
description: |-
  Provides the list of Database Registrations in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_database_registrations
This data source provides the list of Database Registrations in Oracle Cloud Infrastructure Golden Gate service.

Note: Deprecated. Use the /connections API instead.
Lists the DatabaseRegistrations in the compartment.


## Example Usage

```hcl
data "oci_golden_gate_database_registrations" "test_database_registrations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.database_registration_display_name
	state = var.database_registration_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `state` - (Optional) A filter to return only the resources that match the 'lifecycleState' given. 


## Attributes Reference

The following attributes are exported:

* `database_registration_collection` - The list of database_registration_collection.

### DatabaseRegistration Reference

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

