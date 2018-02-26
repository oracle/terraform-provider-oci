# oci_load_balancer_listener

## Listener Resource

### Listener Reference

The following attributes are exported:

* `default_backend_set_name` - The name of the associated backend set.
* `load_balancer_id` - The Load Balancer Id
* `name` - A friendly name for the listener. It must be unique and it cannot be changed. Avoid entering confidential information. Example: `My listener`
* `port` - The communication port for the listener. Example: `80`
* `protocol` - The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation. Example: `HTTP`
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `My_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates.  Example: `true` 


### Create Operation
Adds a listener to a load balancer.

The following arguments are supported:

* `default_backend_set_name` - (Required) The name of the associated backend set.
* `load_balancer_id` - (Required) The Load Balancer Id
* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed. Avoid entering confidential information. Example: `My listener`
* `port` - (Required) The communication port for the listener. Example: `80`
* `protocol` - (Required) The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation. Example: `HTTP`
* `ssl_configuration` - (Optional) 
	* `certificate_name` - (Required) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `My_certificate_bundle` 
	* `verify_depth` - (Optional) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) Whether the load balancer listener should verify peer certificates.  Example: `true` 


### Update Operation
Updates a listener for a given load balancer.

The following arguments support updates:
* `default_backend_set_name` - (Required) The name of the associated backend set.
* `port` - (Required) The communication port for the listener. Example: `80`
* `protocol` - (Required) The protocol on which the listener accepts connection requests. To get a list of valid protocols, use the [ListProtocols](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ListProtocols) operation. Example: `HTTP`
* `ssl_configuration` - (Optional) 
	* `certificate_name` - (Required) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `My_certificate_bundle` 
	* `verify_depth` - (Optional) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) Whether the load balancer listener should verify peer certificates.  Example: `true` 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_load_balancer_listener" "test_listener" {
  load_balancer_id         = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
  name                     = "${var.name}"
  default_backend_set_name = "${var.default_backend_set_name}"
  port                     = "${var.port}"
  protocol                 = "${var.protocol}"

  ssl_configuration {
      certificate_name        = "${var.certificate_name}"
      verify_depth            = "${var.verify_depth}"
      verify_peer_certificate = "${var.verify_peer_certificate}"
  }
}
```