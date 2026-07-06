---
subcategory: "Ddfs"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ddfs_instance"
sidebar_current: "docs-oci-resource-ddfs-instance"
description: |-
  Provides the Instance resource in Oracle Cloud Infrastructure Ddfs service
---

# oci_ddfs_instance
This resource provides the Instance resource in Oracle Cloud Infrastructure Ddfs service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/ddfs

Creates an Instance.


## Example Usage

```hcl
resource "oci_ddfs_instance" "test_instance" {
	#Required
	compartment_id = var.compartment_id
	idcs_url = var.instance_idcs_url

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the Instance in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `idcs_url` - (Required) The Oracle Identity Cloud Service (IDCS) URL for the identity domain associated with the instance. Use the IDCS tenant URL from your identity domain configuration.  Example: `https://idcs-1234567890.identity.oraclecloud.com` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `fhir_service_endpoint` - The HTTPS endpoint for the instance's Fast Healthcare Interoperability Resources (FHIR) service.  Example: `https://example.ddfs.oraclecloud.com/api/fhir` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Instance.
* `idcs_url` - The Oracle Identity Cloud Service (IDCS) URL for the identity domain associated with the instance. Use the IDCS tenant URL from your identity domain configuration.  Example: `https://idcs-1234567890.identity.oraclecloud.com` 
* `lifecycle_details` - A message that describes the current state of the Instance in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `public_ip` - The public IP address for the instance's FHIR service endpoint.
* `state` - The current state of the Instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.example-key": "example-value"}` 
* `time_created` - The date and time the Instance was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Instance was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Instance
	* `update` - (Defaults to 20 minutes), when updating the Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Instance


## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import oci_ddfs_instance.test_instance "id"
```

