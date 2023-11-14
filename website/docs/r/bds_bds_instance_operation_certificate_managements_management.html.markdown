---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_operation_certificate_managements_management"
sidebar_current: "docs-oci-resource-bds-bds_instance_operation_certificate_managements_management"
description: |-
  Provides the Bds Instance Operation Certificate Managements Management resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_operation_certificate_managements_management
This resource provides the Bds Instance Operation Certificate Managements Management resource in Oracle Cloud Infrastructure Big Data Service service.

Configuring TLS/SSL for various ODH services running on the BDS cluster.


## Example Usage

```hcl
resource "oci_bds_bds_instance_operation_certificate_managements_management" "test_bds_instance_operation_certificate_managements_management" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_operation_certificate_managements_management_cluster_admin_password
	services = var.bds_instance_operation_certificate_managements_management_services
	enable_operation_certificate_management = var.enable_operation_certificate_management
	renew_operation_certificate_management = var.renew_operation_certificate_management

	#Optional
	host_cert_details {

		#Optional
		certificate = var.bds_instance_operation_certificate_managements_management_host_cert_details_certificate
		host_name = var.bds_instance_operation_certificate_managements_management_host_cert_details_host_name
		private_key = var.bds_instance_operation_certificate_managements_management_host_cert_details_private_key
	}
	root_certificate = var.bds_instance_operation_certificate_managements_management_root_certificate
	server_key_password = var.bds_instance_operation_certificate_managements_management_server_key_password
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster admin user.
* `host_cert_details` - (Optional) List of leaf certificates to use for services on each host. If custom host certificate is provided the root certificate becomes required.
	* `certificate` - (Optional) Certificate value in string format
	* `host_name` - (Optional) Fully qualified domain name (FQDN) of the host
	* `private_key` - (Optional) Private key of the provided certificate
* `root_certificate` - (Optional) Plain text certificate/s in order, separated by new line character. If not provided in request a self-signed root certificate is generated inside the cluster. In case hostCertDetails is provided, root certificate is mandatory.
* `server_key_password` - (Optional) Base-64 encoded password for CA certificate's private key. This value can be empty.
* `services` - (Required) List of services for which certificate needs to be enabled.
* `enable_operation_certificate_management` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.
* `renew_operation_certificate_management` - (Required) (Updatable) A required field when set to `true` calls renew action and when set to `false` defaults to enable_operation_certificate_management's value action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Operation Certificate Managements Management
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Operation Certificate Managements Management
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Operation Certificate Managements Management
