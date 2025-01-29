---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_domain_governance"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-domain_governance"
description: |-
  Provides details about a specific Domain Governance in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_domain_governance
This data source provides details about a specific Domain Governance resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the domain governance entity.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_domain_governance" "test_domain_governance" {
	#Required
	domain_governance_id = oci_tenantmanagercontrolplane_domain_governance.test_domain_governance.id
}
```

## Argument Reference

The following arguments are supported:

* `domain_governance_id` - (Required) The domain governance OCID.


## Attributes Reference

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

