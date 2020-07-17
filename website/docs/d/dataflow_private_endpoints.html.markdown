---
subcategory: "Dataflow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_private_endpoints"
sidebar_current: "docs-oci-datasource-dataflow-private_endpoints"
description: |-
  Provides the list of Private Endpoints in Oracle Cloud Infrastructure Dataflow service
---

# Data Source: oci_dataflow_private_endpoints
This data source provides the list of Private Endpoints in Oracle Cloud Infrastructure Dataflow service.

Lists all private endpoints in the specified compartment.


## Example Usage

```hcl
data "oci_dataflow_private_endpoints" "test_private_endpoints" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.private_endpoint_display_name}"
	display_name_starts_with = "${var.private_endpoint_display_name_starts_with}"
	owner_principal_id = "${var.owner_principal_id}"
	state = "${var.private_endpoint_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment. 
* `display_name` - (Optional) The query parameter for the Spark application name. Note: At a time only one optional filter can be used with `compartment_id` to get the list of Private Endpoint resources.
* `display_name_starts_with` - (Optional) The displayName prefix. 
* `owner_principal_id` - (Optional) The OCID of the user who created the resource. 
* `state` - (Optional) The LifecycleState of the private endpoint. 


## Attributes Reference

The following attributes are exported:

* `private_endpoint_collection` - The list of private_endpoint_collection.

### PrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - The OCID of a compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user-friendly description. Avoid entering confidential information. 
* `display_name` - A user-friendly name. It does not have to be unique. Avoid entering confidential information. 
* `dns_zones` - An array of DNS zone names. Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of a private endpoint. 
* `lifecycle_details` - The detailed messages about the lifecycle state. 
* `max_host_count` - The maximum number of hosts to be accessed through the private endpoint. This value is used to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up to 512. 
* `nsg_ids` - An array of network security group OCIDs. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `state` - The current state of this private endpoint. 
* `subnet_id` - The OCID of a subnet. 
* `time_created` - The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

