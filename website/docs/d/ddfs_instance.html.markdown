---
subcategory: "Ddfs"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ddfs_instance"
sidebar_current: "docs-oci-datasource-ddfs-instance"
description: |-
  Provides details about a specific Instance in Oracle Cloud Infrastructure Ddfs service
---

# Data Source: oci_ddfs_instance
This data source provides details about a specific Instance resource in Oracle Cloud Infrastructure Ddfs service.

Gets information about an Instance.

## Example Usage

```hcl
data "oci_ddfs_instance" "test_instance" {
	#Required
	instance_id = oci_ddfs_instance.test_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Instance.


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

