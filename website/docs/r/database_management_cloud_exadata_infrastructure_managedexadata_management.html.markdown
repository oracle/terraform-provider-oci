---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_exadata_infrastructure_managedexadata_management"
sidebar_current: "docs-oci-resource-database_management-cloud_exadata_infrastructure_managedexadata_management"
description: |-
  Provides the Cloud Exadata Infrastructure Managedexadata Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_exadata_infrastructure_managedexadata_management
This resource provides the Cloud Exadata Infrastructure Managedexadata Management resource in Oracle Cloud Infrastructure Database Management service.

Enables Database Management for the Exadata infrastructure specified by exadataInfrastructureId. It covers the following 
components:

- Exadata infrastructure
- Exadata storage grid
- Exadata storage server


## Example Usage

```hcl
resource "oci_database_management_cloud_exadata_infrastructure_managedexadata_management" "test_cloud_exadata_infrastructure_managedexadata_management" {
	#Required
	cloud_exadata_infrastructure_id = oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
	enable_managedexadata = var.enable_managedexadata

	#Optional
	license_model = var.cloud_exadata_infrastructure_managedexadata_management_license_model
}
```

## Argument Reference

The following arguments are supported:

* `cloud_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `license_model` - (Optional) The Oracle license model that applies to the database management resources. 
* `enable_managedexadata` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Exadata Infrastructure Managedexadata Management
	* `update` - (Defaults to 20 minutes), when updating the Cloud Exadata Infrastructure Managedexadata Management
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Exadata Infrastructure Managedexadata Management
