# oci_dns_record

## Records Resource

### Records Reference

The following attributes are exported:
* `compartment_id` - (Optional) The OCID of the compartment the resource belongs to.
* `domain` - The fully qualified domain name where the record can be located. 
* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. 
* `record_hash` - A unique identifier for the record within its zone. 
* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
* `rtype` - The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - The Time To Live for the record, in seconds.
* `zone_name_or_id` - The name or OCID of the target zone.



### Create Operation
Replaces records in the specified zone with the records specified in the
request body. If a specified record does not exist, it will be created.
If the record exists, then it will be updated to represent the record in
the body of the request. If a record in the zone does not exist in the
request body, the record will be removed from the zone.


The following arguments are supported:
 
* `compartment_id` - (Optional) The OCID of the compartment the resource belongs to. If supplied, it must match the Zone's compartment ocid. 
* `domain` - (Optional) The fully qualified domain name where the record can be located.  
* `rdata` - (Optional) The record's data, as whitespace-delimited tokens in type-specific presentation format.  
* `rtype` - (Optional) The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - (Optional) The Time To Live for the record, in seconds.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.


### Update Operation
Replaces records in the specified zone with the records specified in the
request body. If a specified record does not exist, it will be created.
If the record exists, then it will be updated to represent the record in
the body of the request. If a record in the zone does not exist in the
request body, the record will be removed from the zone.


The following arguments support updates:
* `compartment_id` - The OCID of the compartment the resource belongs to.
* `domain` - (Optional) The fully qualified domain name where the record can be located.  
* `rdata` - (Optional) The record's data, as whitespace-delimited tokens in type-specific presentation format.  
* `rtype` - (Optional) The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `ttl` - (Optional) The Time To Live for the record, in seconds. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_dns_records" "test_record" {
	#Required
	zone_name_or_id = "${oci_dns_zone_name_or.test_zone_name_or.id}"

	#Optional
	compartment_id = "${var.compartment_id}"
    #Optional
    domain = "${var.record_items_domain}"
    rdata = "${var.record_items_rdata}"
    rtype = "${var.record_items_rtype}"
    ttl = "${var.record_items_ttl}"
}
```

# oci_dns_records

## Records DataSource

Gets a list of records

### Get Operation
Gets all records in the specified zone. The results are
sorted by `domain` in alphabetical order by default. For more
information about records, please see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4).

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment the resource belongs to.
* `domain` - (Optional) Search by domain. Will match any record whose domain (case-insensitive) equals the provided value. 
* `domain_contains` - (Optional) Search by domain. Will match any record whose domain (case-insensitive) contains the provided value. 
* `rtype` - (Optional) Search by record type. Will match any record whose [type](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4) (case-insensitive) equals the provided value.
* `sort_by` - (Optional) The field by which to sort records. Allowed values are: domain|rtype|ttl
* `sort_order` - The order to sort the resources. Allowed values are: ASC|DESC 
* `zone_name_or_id` - (Required) The name or OCID of the target zone.
* `zone_version` - (Optional) The version of the zone for which data is requested. 


The following attributes are exported:

* `records` - A collection of DNS resource records.


### Example Usage

```
data "oci_dns_records" "test_records" {
	#Required
	zone_name_or_id = "${oci_dns_zone_name_or.test_zone_name_or.id}"

	#Optional
	compartment_id = "${var.compartment_id}"
	domain = "${var.record_domain}"
	domain_contains = "${var.record_domain_contains}"
	rtype = "${var.record_rtype}"
	zone_version = "${var.record_zone_version}"
}
```