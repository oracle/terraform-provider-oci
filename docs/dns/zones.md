# oci_dns_zone

## Zone Resource

### Zone Reference

The following attributes are exported:
* `compartment_id` - The OCID of the compartment containing the zone.
* `external_masters` - External master servers for the zone.
	* `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port.
	* `tsig` - A TSIG key
		* `algorithm` - TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. For a full list of TSIG algorithms, see [Secret Key Transaction Authentication for DNS (TSIG) Algorithm Names](http://www.iana.org/assignments/tsig-algorithm-names/tsig-algorithm-names.xhtml#tsig-algorithm-names-1) 
		* `name` - A domain name identifying the key for a given pair of hosts.
		* `secret` - A base64 string encoding the binary shared secret.
* `name` - The name of the zone.
* `self` - The canonical absolute URL of the resource.
* `serial` - The current serial of the zone. As seen in the zone's SOA record.
* `state` - The Zone's current state.
* `time_created` - The date and time the Zone was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
* `version` - Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived.
* `zone_type` - The type of the zone. Must be either `PRIMARY` or `SECONDARY`. 



### Create Operation
Creates a new zone in the specified compartment.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the zone.
* `external_masters` - (Optional) External master servers for the zone.
	* `address` - (Required) The server's IP address (IPv4 or IPv6).
	* `port` - (Optional) The server's port.
	* `tsig` - (Optional) 
		* `algorithm` - (Required) TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. For a full list of TSIG algorithms, see [Secret Key Transaction Authentication for DNS (TSIG) Algorithm Names](http://www.iana.org/assignments/tsig-algorithm-names/tsig-algorithm-names.xhtml#tsig-algorithm-names-1) 
		* `name` - (Required) A domain name identifying the key for a given pair of hosts.
		* `secret` - (Required) A base64 string encoding the binary shared secret.
* `name` - (Required) The name of the zone.
* `zone_type` - (Required) The type of the zone. Must be either `PRIMARY` or `SECONDARY`. 


### Update Operation
Updates the specified secondary zone with your new external master
server information. For more information about secondary zone, see
[Manage DNS Service Zone](https://docs.us-phoenix-1.oraclecloud.com/Content/DNS/Tasks/managingdnszones.htm).


The following arguments support updates:
* `external_masters` - External master servers for the zone.
    * `address` - The server's IP address (IPv4 or IPv6).
	* `port` - The server's port.
	* `tsig` - 
		* `algorithm` - (Required) TSIG Algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. For a full list of TSIG algorithms, see [Secret Key Transaction Authentication for DNS (TSIG) Algorithm Names](http://www.iana.org/assignments/tsig-algorithm-names/tsig-algorithm-names.xhtml#tsig-algorithm-names-1) 
		* `name` - (Required) A domain name identifying the key for a given pair of hosts.
		* `secret` - (Required) A base64 string encoding the binary shared secret.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.zone_name}"
	zone_type = "${var.zone_zone_type}"

	#Optional
	external_masters {
		#Required
		address = "${var.zone_external_masters_address}"

		#Optional
		port = "${var.zone_external_masters_port}"
		tsig {
			#Required
			algorithm = "${var.zone_external_masters_tsig_algorithm}"
			name = "${var.zone_external_masters_tsig_name}"
			secret = "${var.zone_external_masters_tsig_secret}"
		}
	}
}
```

# oci_dns_zones

## Zones DataSource

Gets a list of zones

### Get Operation
Gets a list of all zones in the specified compartment. The collection
can be filtered by name, time created, and zone type.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `name` - (Optional) A case-sensitive filter for zone names. Will match any zone with a name that equals the provided value. 
* `name_contains` - (Optional) Search by zone name. Will match any zone whose name (case-insensitive) contains the provided value.
* `sort_by` - (Optional) The field by which to sort zones. Allowed values are: name|zoneType|timeCreated
* `sort_order` - The order to sort the resources. Allowed values are: ASC|DESC  
* `state` - (Optional) The state of a resource.
* `time_created_greater_than_or_equal_to` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created on or after the indicated time. 
* `time_created_less_than` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created before the indicated time. 
* `zone_type` - (Optional) Search by zone type, `PRIMARY` or `SECONDARY`. Will match any zone whose type equals the provided value. 


The following attributes are exported:

* `zones` - A list of DNS zones.

### Example Usage

```
data "oci_dns_zones" "test_zones" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	name = "${var.zone_name}"
	name_contains = "${var.zone_name_contains}"
	state = "${var.zone_state}"
	time_created_greater_than_or_equal_to = "${var.zone_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.zone_time_created_less_than}"
	zone_type = "${var.zone_zone_type}"
}
```