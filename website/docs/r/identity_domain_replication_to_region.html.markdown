---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domain_replication_to_region"
sidebar_current: "docs-oci-resource-identity-domain_replication_to_region"
description: |-
  Provides the Domain Replication To Region resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_domain_replication_to_region
This resource provides the Domain Replication To Region resource in Oracle Cloud Infrastructure Identity service.

(For tenancies that support identity domains) Replicates the identity domain to a new region (provided that the region is the 
tenancy home region or other region that the tenancy subscribes to). You can only replicate identity domains that are in an ACTIVE 
`lifecycleState` and not currently updating or already replicating. You also can only trigger the replication of secondary identity domains. 
The default identity domain is automatically replicated to all regions that the tenancy subscribes to.

After you send the request, the `state` of the identity domain in the replica region is set to ENABLING_REPLICATION. When the operation 
completes, the `state` is set to REPLICATION_ENABLED.

To track the progress of the request, submitting an HTTP GET on the /iamWorkRequests/{iamWorkRequestsId} endpoint retrieves
the operation's status.


## Example Usage

```hcl
resource "oci_identity_domain_replication_to_region" "test_domain_replication_to_region" {
	#Required
	domain_id = oci_identity_domain.test_domain.id

	#Optional
	replica_region = var.domain_replication_to_region_replica_region
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required) The OCID of the identity domain.
* `replica_region` - (Optional) A region to which you want identity domain replication to occur. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Domain Replication To Region
	* `update` - (Defaults to 20 minutes), when updating the Domain Replication To Region
	* `delete` - (Defaults to 20 minutes), when destroying the Domain Replication To Region


## Import

Import is not supported for this resource.

