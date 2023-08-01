---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_vm_cluster_ords_certificate_management"
sidebar_current: "docs-oci-resource-database-autonomous_vm_cluster_ords_certificate_management"
description: |-
  Provides the Autonomous Vm Cluster Ords Certificate Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_vm_cluster_ords_certificate_management
This resource provides the Autonomous Vm Cluster Ords Certificate Management resource in Oracle Cloud Infrastructure Database service.

Rotates the Oracle REST Data Services (ORDS) certificates for Autonomous Exadata VM cluster.


## Example Usage

```hcl
resource "oci_database_autonomous_vm_cluster_ords_certificate_management" "test_autonomous_vm_cluster_ords_certificate_management" {
	#Required
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	certificate_generation_type = var.autonomous_vm_cluster_ords_certificate_management_certificate_generation_type

	#Optional
	ca_bundle_id = oci_certificates_management_ca_bundle.test_ca_bundle.id
	certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
	certificate_id = oci_apigateway_certificate.test_certificate.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_vm_cluster_id` - (Required) The autonomous VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `ca_bundle_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate bundle. 
* `certificate_authority_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate authority. 
* `certificate_generation_type` - (Required) Specify SYSTEM for using Oracle managed certificates. Specify BYOC when you want to bring your own certificate.
* `certificate_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate to use. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Vm Cluster Ords Certificate Management
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Vm Cluster Ords Certificate Management
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Vm Cluster Ords Certificate Management


## Import

Import is not supported for this resource.

