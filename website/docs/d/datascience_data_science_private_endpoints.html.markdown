---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_private_endpoints"
sidebar_current: "docs-oci-datasource-datascience-data_science_private_endpoints"
description: |-
  Provides the list of Data Science Private Endpoints in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_private_endpoints
This data source provides the list of Data Science Private Endpoints in Oracle Cloud Infrastructure Data Science service.

Lists all Data Science private endpoints in the specified compartment. The query must include compartmentId. The query can also include one other parameter. If the query doesn't include compartmentId, or includes compartmentId with two or more other parameters, then an error is returned.


## Example Usage

```hcl
data "oci_datascience_private_endpoints" "test_data_science_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.data_science_private_endpoint_created_by
	data_science_resource_type = var.data_science_private_endpoint_data_science_resource_type
	display_name = var.data_science_private_endpoint_display_name
	state = var.data_science_private_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `data_science_resource_type` - (Optional) Resource types in the Data Science service such as notebooks. 
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `state` - (Optional) The lifecycle state of the private endpoint. 


## Attributes Reference

The following attributes are exported:

* `data_science_private_endpoints` - The list of data_science_private_endpoints.

### DataSciencePrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create private endpoint.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user that created the private endpoint.
* `data_science_resource_type` - Data Science resource type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user friendly description. Avoid entering confidential information. 
* `display_name` - A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `fqdn` - Accesing the Data Science resource using FQDN. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of a private endpoint. 
* `lifecycle_details` - Details of the state of Data Science private endpoint.
* `nsg_ids` - An array of network security group OCIDs. 
* `state` - State of the Data Science private endpoint.
* `subnet_id` - The OCID of a subnet. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the Data Science private endpoint was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time that the Data Science private endpoint was updated expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

