# oci\_load\_balancer\_backendsets

[Backendset Reference][0aa66c9f]

  [0aa66c9f]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/ "BackendsetReference"

Provides a list of load balancer backendsets.

## Example Usage

```
data "oci_load_balancer_backendsets" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The OCID of the load balancer.


## Attributes Reference
* `backendsets` - The list of backendsets.

## Backendset reference
* `name` - A friendly name for the backend set. It must be unique and it cannot be changed. Avoid entering confidential information.
* `policy` - The load balancer policy for the backend set. The default load balancing policy is 'ROUND_ROBIN'.
* `health_checker` - Health Checker Settings.
* `ssl_configuration` - SSL Configuration Settings.
* `session_persistence_configuration` - (Optional) Session persistence enables the Load Balancing Service to direct any number of requests that originate from a single logical client to a single backend web server.
