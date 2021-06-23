---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_certificate"
sidebar_current: "docs-oci-resource-apigateway-certificate"
description: |-
  Provides the Certificate resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_certificate
This resource provides the Certificate resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new Certificate.


## Example Usage

```hcl
resource "oci_apigateway_certificate" "test_certificate" {
	#Required
	certificate = "${var.certificate_certificate}"
	compartment_id = "${var.compartment_id}"
	private_key = "${var.certificate_private_key}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.certificate_display_name}"
	freeform_tags = {"Department"= "Finance"}
	intermediate_certificates = "${var.certificate_intermediate_certificates}"
}
```

## Argument Reference

The following arguments are supported:

* `certificate` - (Required) The data of the leaf certificate in pem format.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `intermediate_certificates` - (Optional) The intermediate certificate data associated with the certificate in pem format.
* `private_key` - (Required) The private key associated with the certificate in pem format.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `certificate` - The data of the leaf certificate in pem format.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `intermediate_certificates` - The intermediate certificate data associated with the certificate in pem format.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `state` - The current state of the certificate.
* `subject_names` - The entity to be secured by the certificate and additional host names.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_not_valid_after` - The date and time the certificate will expire.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Certificate
	* `update` - (Defaults to 20 minutes), when updating the Certificate
	* `delete` - (Defaults to 20 minutes), when destroying the Certificate


## Import

Certificates can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_certificate.test_certificate "id"
```

