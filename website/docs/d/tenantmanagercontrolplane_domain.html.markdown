---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_domain"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-domain"
description: |-
  Provides details about a specific Domain in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_domain
This data source provides details about a specific Domain resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the domain.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_domain" "test_domain" {
	#Required
	domain_id = oci_tenantmanagercontrolplane_domain.test_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required) The domain OCID.


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `domain_name` - The domain name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the domain.
* `owner_id` - The OCID of the tenancy that has started the registration process for this domain.
* `state` - Lifecycle state of the domain.
* `status` - Status of the domain.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date-time when this domain was created. An RFC 3339-formatted date and time string.
* `time_updated` - Date-time when this domain was last updated. An RFC 3339-formatted date and time string.
* `txt_record` - The code that the owner of the domain will need to add as a TXT record to their domain.

