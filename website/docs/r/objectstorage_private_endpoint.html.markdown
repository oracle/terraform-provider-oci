---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_private_endpoint"
sidebar_current: "docs-oci-resource-objectstorage-private-endpoint"
description: |-
  Provides the Private Endpoint resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_private_endpoint
This resource provides the Private Endpoint resource in Oracle Cloud Infrastructure Object Storage service. 
It enables private network access from a specified subnet to Object Storage without traversing the public internet.

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/object_storage/private_endpoint

Creates an Object Storage Private Endpoint in a specified subnet, with access scoping for namespace, compartment, and bucket.

## Example Usage

```hcl
resource "oci_objectstorage_private_endpoint" "pe" {
  #Required
  compartment_id = var.compartment_ocid
  namespace = var.namespace_name
  name = var.pe_name
  subnet_id = oci_core_subnet.test_subnet_1.id
  
  #Optional
  prefix = var.dns_prefix
  access_targets  {
    namespace = "*"
    compartment_id = "*"
    bucket = "*"
  }
}
```

## Argument Reference

The following arguments are supported:

* `access_targets` - (Required) (Updatable) When you create a private endpoint, you can restrict access to certain Object Storage resources by specifying access targets (limit of 10). Each access target consists of the following required parameters: namespace, compartment_id and bucket.
  * `namespace` - (Required) (Updatable) Specifies the target namespace that's to be allowed to egress from the private endpoint.
  * `compartment_id` - (Required) (Updatable) Specifies what namespace/compartments the private endpoint can access. You can configure either a single compartment or all compartments.
  * `bucket` - (Required) (Updatable) Specifies what namespace/buckets within the allowed compartments the private endpoint can access. You can configure either a single bucket or all buckets within the allowed compartments.
* `additional_prefixes` (Optional) A list of additional prefixes that you can provide along with any other prefix. These resulting endpointFqdn's are added to the customer VCN's DNS record.
* `compartment_id` - (Required) The ID of the compartment in which to create the private endpoint.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `name` - (Required) The name of the private endpoint. Valid characters are uppercase or lowercase letters, numbers, hyphens, and periods. Private Endpoint names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-pe1
* `namespace` - (Required) The Object Storage namespace used for the request.
* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups (NSGs) to add the private endpoint's VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.oracle.com/en-us/iaas/Content/Network/Concepts/networksecuritygroups.htm).
* `prefix` - (Required) The DNS prefix value is part of the URL used to access Object Storage. The DNS prefix is a case-insensitive string using alpha-numeric characters (no special characters). It must be unique within the VCN.
* `private_endpoint_ip` - (Optional) The private IP address that is to be assigned to this private endpoint. If it's not available, an error is returned. If you do not provide a value, an available IP address in the subnet is automatically chosen. If you do not provide a value, an available IP address in the subnet is automatically chosen.
* `subnet_id` - (Required) The ID of the subnet that the private endpoint VNIC will be created and reside in. 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment ID in which the private endpoint resource exists in.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the private endpoint.
* `etag` - The entity tag for the Private Endpoint.
* `fqdns` - The object representing FQDN details formed using prefix and additionalPrefixes.
* `name` - The name of the private endpoint. Avoid entering confidential information. Example: my-pe1
* `namespace` - The Object Storage namespace in which the private endpoint resides.
* `prefix` - The DNS prefix value chosen which is the first part of the URL used to access Object Storage.
* `state` - The lifecycle state of the private endpoint resource. 
* `time_created` - The date and time the private endpoint was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
* `time_modified` - The date and time the private endpoint was updated, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).

## Import

Private endpoints can be imported using the `namespaceName` and the `name` of the private endpoint.

```
$ terraform import oci_objectstorage_private_endpoint.test_pe "n/{namespaceName}/pe/{peName}" 
```

