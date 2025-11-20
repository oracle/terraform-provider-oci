---
subcategory: "Psa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psa_psa_services"
sidebar_current: "docs-oci-datasource-psa-psa_services"
description: |-
  Provides the list of Psa Services in Oracle Cloud Infrastructure Psa service
---

# Data Source: oci_psa_psa_services
This data source provides the list of Psa Services in Oracle Cloud Infrastructure Psa service.

List the Oracle Cloud Infrastructure services available for Private Service Access catalog in the region, sorted by service name.


## Example Usage

```hcl
data "oci_psa_psa_services" "test_psa_services" {

	#Optional
	display_name = var.psa_service_display_name
	service_id = oci_psa_psa_service.test_psa_service.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `service_id` - (Optional) The unique identifier of the Oracle Cloud Infrastructure service.


## Attributes Reference

The following attributes are exported:

* `psa_service_collection` - The list of psa_service_collection.

### PsaService Reference

The following attributes are exported:

* `items` - List of PsaServiceSummary.
	* `description` - A description of the Oracle Cloud Infrastructure service. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `fqdns` - The public facing service FQDNs, which are going to be used to access the service.  Example: `xyz.oraclecloud.com` 
	* `id` - A unique Oracle Cloud Infrastructure service identifier.  Example: `object-storage-api` 
	* `is_v6enabled` - This optional field will indicate that whether service is IPv6 enabled. 

