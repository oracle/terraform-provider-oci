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

Replicate domain to a new region. This is an asynchronous call - where, at start,
{@code state} of this domain in replica region is set to ENABLING_REPLICATION.
On domain replication completion the {@code state} will be set to REPLICATION_ENABLED.

To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
the async operation's status.

If the replica region's {@code state} is already ENABLING_REPLICATION or REPLICATION_ENABLED,
returns 409 CONFLICT.
- If the domain doesn't exists, returns 404 NOT FOUND.
- If home region is same as replication region, return 400 BAD REQUEST.
- If Domain is not active or being updated, returns 400 BAD REQUEST.
- If any internal error occurs, return 500 INTERNAL SERVER ERROR.


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

* `domain_id` - (Required) The OCID of the domain
* `replica_region` - (Optional) A region for which domain replication is requested for. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Domain Replication To Region
	* `update` - (Defaults to 20 minutes), when updating the Domain Replication To Region
	* `delete` - (Defaults to 20 minutes), when destroying the Domain Replication To Region


## Import

Import is not supported for this resource.

