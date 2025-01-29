---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_domain_governances"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-domain_governances"
description: |-
  Provides the list of Domain Governances in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_domain_governances
This data source provides the list of Domain Governances in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Return a (paginated) list of domain governance entities.


## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_domain_governances" "test_domain_governances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	domain_governance_id = oci_tenantmanagercontrolplane_domain_governance.test_domain_governance.id
	domain_id = oci_tenantmanagercontrolplane_domain.test_domain.id
	name = var.domain_governance_name
	state = var.domain_governance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `domain_governance_id` - (Optional) The domain governance OCID.
* `domain_id` - (Optional) The domain OCID.
* `name` - (Optional) A filter to return only resources that exactly match the name given.
* `state` - (Optional) The lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `domain_governance_collection` - The list of domain_governance_collection.

### DomainGovernance Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `domain_id` - The OCID of the domain associated with this domain governance entity.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the domain governance entity.
* `is_governance_enabled` - Indicates whether governance is enabled for this domain.
* `ons_subscription_id` - The ONS subscription associated with this domain governance entity.
* `ons_topic_id` - The ONS topic associated with this domain governance entity.
* `owner_id` - The OCID of the tenancy that owns this domain governance entity.
* `state` - Lifecycle state of the domain governance entity.
* `subscription_email` - Email address to be used to notify the user, and that the ONS subscription will be created with.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date-time when this domain governance was created. An RFC 3339-formatted date and time string.
* `time_updated` - Date-time when this domain governance was last updated. An RFC 3339-formatted date and time string.

