---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_certificate"
sidebar_current: "docs-oci-datasource-apigateway-certificate"
description: |-
  Provides details about a specific Certificate in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_certificate
This data source provides details about a specific Certificate resource in Oracle Cloud Infrastructure API Gateway service.

Gets a certificate by identifier.

## Example Usage

```hcl
data "oci_apigateway_certificate" "test_certificate" {
	#Required
	certificate_id = "${oci_apigateway_certificate.test_certificate.id}"
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Required) The ocid of the certificate.


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

