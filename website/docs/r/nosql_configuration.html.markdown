---
subcategory: "NoSQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_configuration"
sidebar_current: "docs-oci-resource-nosql-configuration"
description: |-
  Provides the Configuration resource in Oracle Cloud Infrastructure NoSQL Database service
---

# oci_nosql_configuration
This resource provides the Configuration resource in Oracle Cloud Infrastructure NoSQL Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/nosql-database/latest/Configuration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/object_storage

Updates the service-level configuration.  The discriminator value
`UpdateConfigurationDetails.environment` must match the service's
environment type.

A configuration serves as a centralized repository for global parameters that
affect the NoSQL service. Currently, there is only one such parameter: a
customer-provided key for encrypting NoSQL data at rest.

The Customer-Managed Encryption Keys (CMEK) feature is exclusively available
in private NoSQL environments dedicated to a single tenancy, where the CMEK
option has been enabled. Updating the configuration of the default, regional,
multi-tenancy NoSQL service is not supported.

To specify the dedicated environment, set the environment variable
CLIENT_HOST_OVERRIDES=oci_nosql.NosqlClient=$ENDPOINT
Where $ENDPOINT is the endpoint of the dedicated NoSQL environment.
For example:
$ export CLIENT_HOST_OVERRIDES=oci_nosql.NosqlClient=https://acme-widgets.nosql.oci.oraclecloud.com

## Example Usage

```hcl
resource "oci_nosql_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
	environment = "HOSTED"

	#Optional
	is_opc_dry_run = var.configuration_is_opc_dry_run
	kms_key {

		#Optional
		id = var.configuration_kms_key_id
		kms_vault_id = oci_kms_vault.test_vault.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The tenancy's OCID
* `environment` - (Required) (Updatable) The service environment type.
* `is_opc_dry_run` - (Optional) (Updatable) If true, indicates that the request is a dry run. A dry run request does not modify the configuration item details and is used only to perform validation on the submitted data. 
* `kms_key` - (Required when environment=HOSTED) (Updatable) Information about the state of the service's encryption key management. The following properties are read-only and ignored when this object is used in UpdateConfiguration: kmsKeyState, timeCreated, timeUpdated. 
	* `id` - (Applicable when environment=HOSTED) (Updatable) The OCID of the KMS encryption key assigned to this Hosted Environment. If the Hosted Environment is using an Oracle-managed Key (i.e., not a CMEK), then the id will be a null string.
	* `kms_key_state` - (Applicable when environment=HOSTED) (Updatable) The current state of the encryption key assigned to this Hosted Environment. Oracle-managed keys will always report an ACTIVE state. 
	* `kms_vault_id` - (Applicable when environment=HOSTED) (Updatable) The OCID of the vault containing the encryption key assigned to this Hosted Environment. If the Hosted Environment is using an Oracle-managed Key (i.e., not a CMEK), then the kmsVaultId will be a null string. 
	* `time_created` - (Applicable when environment=HOSTED) (Updatable) The timestamp when encryption key was first enabled for this Hosted Environment. RFC3339 formatted. 
	* `time_updated` - (Applicable when environment=HOSTED) (Updatable) The timestamp of the last update to the encryption key status. RFC3339 formatted. 


** IMPORTANT **
The configuration cannot be deleted.

## Attributes Reference

The following attributes are exported:

* `environment` - The service environment type.
* `kms_key` - Information about the state of the service's encryption key management. The following properties are read-only and ignored when this object is used in UpdateConfiguration: kmsKeyState, timeCreated, timeUpdated. 
	* `id` - The OCID of the KMS encryption key assigned to this Hosted Environment. If the Hosted Environment is using an Oracle-managed Key (i.e., not a CMEK), then the id will be a null string. 
	* `kms_key_state` - The current state of the encryption key assigned to this Hosted Environment. Oracle-managed keys will always report an ACTIVE state. 
	* `kms_vault_id` - The OCID of the vault containing the encryption key assigned to this Hosted Environment. If the Hosted Environment is using an Oracle-managed Key (i.e., not a CMEK), then the kmsVaultId will be a null string. 
	* `time_created` - The timestamp when encryption key was first enabled for this Hosted Environment. RFC3339 formatted. 
	* `time_updated` - The timestamp of the last update to the encryption key status. RFC3339 formatted. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Configuration
	* `update` - (Defaults to 20 minutes), when updating the Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Configuration


## Import

Configurations can be imported using the `id`, e.g.

```
$ terraform import oci_nosql_configuration.test_configuration "configuration/compartmentId/{compartmentId}" 
```

