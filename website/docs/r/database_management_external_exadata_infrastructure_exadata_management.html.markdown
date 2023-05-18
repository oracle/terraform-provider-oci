---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_infrastructure_exadata_management"
sidebar_current: "docs-oci-resource-database_management-external_exadata_infrastructure_exadata_management"
description: |-
  Provides the External Exadata Infrastructure Exadata Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_exadata_infrastructure_exadata_management
This resource provides the External Exadata Infrastructure Exadata Management resource in Oracle Cloud Infrastructure Database Management service.

Enables Database Management for the Exadata infrastructure specified by externalExadataInfrastructureId. It covers the following 
components:

- Exadata infrastructure
- Exadata storage grid
- Exadata storage server


## Example Usage

```hcl
resource "oci_database_management_external_exadata_infrastructure_exadata_management" "test_external_exadata_infrastructure_exadata_management" {
	#Required
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
	enable_exadata = var.enable_exadata

	#Optional
	license_model = var.external_exadata_infrastructure_exadata_management_license_model
}
```

## Argument Reference

The following arguments are supported:

* `external_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `license_model` - (Optional) The Oracle license model. 
* `enable_exadata` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Exadata Infrastructure Exadata Management
	* `update` - (Defaults to 20 minutes), when updating the External Exadata Infrastructure Exadata Management
	* `delete` - (Defaults to 20 minutes), when destroying the External Exadata Infrastructure Exadata Management
