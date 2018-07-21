---
layout: "oci"
page_title: "OCI: oci_identity_compartments"
sidebar_current: "docs-oci-datasource-identity-compartments"
description: |-
Provides a list of Compartments
---
# Data Source: oci_identity_compartments
The Compartments data source allows access to the list of OCI compartments

Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
for the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_compartments" "test_compartments" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `compartments` - The list of compartments.

### Compartment Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the compartment. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the compartment.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy. Avoid entering confidential information. 
* `state` - The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the compartment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

