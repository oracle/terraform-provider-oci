---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_domains"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-domains"
description: |-
  Provides the list of Domains in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_domains
This data source provides the list of Domains in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Return a (paginated) list of domains.


## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_domains" "test_domains" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	domain_id = oci_tenantmanagercontrolplane_domain.test_domain.id
	name = var.domain_name
	state = var.domain_state
	status = var.domain_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `domain_id` - (Optional) The domain OCID.
* `name` - (Optional) A filter to return only resources that exactly match the name given.
* `state` - (Optional) The lifecycle state of the resource.
* `status` - (Optional) The status of the domain.


## Attributes Reference

The following attributes are exported:

* `domain_collection` - The list of domain_collection.

### Domain Reference

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

