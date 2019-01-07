// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/loadbalancer"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerBackendSetTestSuite struct {
	suite.Suite
	Providers           map[string]terraform.ResourceProvider
	Config              string
	ResourceName        string
	BackendResourceName string
}

func (s *ResourceLoadBalancerBackendSetTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.0.0/24"
		display_name        = "-tf-subnet"
	}
	
	resource "oci_load_balancer" "t" {
		shape = "100Mbps"
		compartment_id = "${var.compartment_id}"
		subnet_ids = ["${oci_core_subnet.t.id}"]
		display_name = "-tf-lb"
		is_private = true
	}
	
	resource "oci_load_balancer_certificate" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
		ca_certificate = "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"
		certificate_name = "tf_cert_name"
		private_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1vqH90VOIZHh/FL9pEH23op3t+FcMBSKQ+ijCoDZfil5m12X\nUQWUTS/ZFauGo0q1O4esHI8vusL6eyBVDOiTvyGfgXBYWq+yPkIGlqcuOdxN/bQv\nhfAJN5hcJYjosT/KhlwuuYO7V4ZWcgpacoyETh5bMinEPJ1k7e6dF4/4p3lisERP\n2zkxDtP98ZlUbo0ZF6sXPe9PJqYUCk0DXkH1GkffX6Zp+39q+ywJVP6Up54USuZ5\nmCJRyrGn0hq5+ha/UNLm6ggxa3sE2+/Rzx9sFrWLuxuEKQJJ4KPBuNreIHQk2z6K\nTH316SplmZRgTXqotTpYgWF9Ml/14ljvlDo6RwIDAQABAoIBAQCFzHXVQ1BGenpR\nRgHROrEAfuPWETAESLRpYaAgCGPVLtEeDpj/913+0FnnL9NjTDsR6vYG7GNDdNja\nyxvEJfjWy4Fv2VFUV+ey8fsRxslxf5kW3w946BWEgZJQVi6lKtPM3hDCq6ds6RJi\ndeknRCeQSzptNSuKoldP8uPY52VWLYTyy/ODwtSCFZKTm7iTD3RkpAqNMPs2V8EA\nRbjiu70q4Kk+ozHQ/0wtOSjZinR6LW7e+6bXxdVV0hdt52h2YzdQrQkHu5ATm54P\nm1S9PbiyV06BSuU37oZbyWiblP6rsJIqucEXSSmTY/5PCG+huzTseig1bot19eD0\nwEs9YAixAoGBAPgKMZ+VhA3eoAx21R62SmJJKdCYZ+qjPmvUt4tW3dJxYTweJo6A\nwyh+p3VsBwE6L8hPD1OyEvUrl64uQEMYSMmNaLpaMqNbytP8h6uiFGyRUSnkYvxF\nPcKW6UjKZNyGWT+dOnJN07DnwvZgOx2JZJkgFpaAEL6g0Hzm1bI/6dkPAoGBAN3g\nt+0sDfgYr3raGAPLQJCVgHV3MdIDHP/ebZBT1NTcxL7Wf/+0WAvnwD7DPsXj84JG\nzNMk7+EzwClAWGAQJMymC9NltfygyI0JjI+88nVk3mrpDq/zqR/vkn2R1T78No/X\nEWTnMmHlzDg4HsGTWKDg7jrmYSas0NvMHPtBbVtJAoGAKV0B03wKjomOpSV3+uwp\nUWSkDX4s7isU8MSDa0AsM7jmnzDj+yWr5efhIyrFrEW4zC2q/6kVkj8Xx1s9KjM1\niC8FxPXftfBLzbgyI8QepdBB+bt1al5do0KpWpMt6Lyay4n7wi4KXFj54T5A/Xb5\nCLQaMDThFfkZa4rPHi+cXq8CgYAfk665a06lo2W99zH5wEB1E0HP9eG6QMUsyQwQ\nwU2F6dF6U26uBo2NTDM4+3KAmVt7i/X0iso047eSZ1zsdv+1vF/sewo2ZO+F2vkN\nL9fVy0A4OOjlM6k7KU5Q3qNZrm1ZdUM9eAXclubEjYAbDoxLgReGfGkRJwEmdtsd\nCwe0OQKBgQD3m85OXSSf+xlm7tGO66bcxHifkp4XfkqWxFwpfYkxNtYZfFpN8jL5\niS0OyLldmJbCVB6EIs5ylW86aeZMH/JecPTxOnaT4qc43PMrLi4MSa65Gp/Zgs1U\nyO0hfWlpH2ncUIuQEksXEPSKQUjvdQl7pD8kghCbDYbm3zsjw3rkyA==\n-----END RSA PRIVATE KEY-----"
		public_certificate = "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"
	}`
	s.ResourceName = "oci_load_balancer_backendset.t"
	s.BackendResourceName = "oci_load_balancer_backend.t"
}

// todo: BackendSet health_checker appears to be missing 3 properties
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/HealthCheckerDetails
func (s *ResourceLoadBalancerBackendSetTestSuite) TestAccResourceLoadBalancerBackendSet_basic() {
	var res, res2 string
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "ROUND_ROBIN"
					health_checker {
						interval_ms = 30000
						port = 1234
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "30000"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "1234"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res, err = fromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test update
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 4321
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test add session persistence
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "IP_HASH"
					health_checker {
						interval_ms = 29999
						port = 4321
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					session_persistence_configuration {
						cookie_name = "lb-session1"
						disable_fallback = true
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "IP_HASH"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test add ssl
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 4321
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}
				
					session_persistence_configuration {
						cookie_name = "lb-session2"
						disable_fallback = false
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session2"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test create backend and update backendset (the two operations may happen concurrently)
			// This test is needed because CreateBackend and UpdateBackendSet both update the same backend set
			// resource. Test that both changes are applied sequentially.
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 8080
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					session_persistence_configuration {
						cookie_name = "lb-session2"
						disable_fallback = false
					}

					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.id}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				// Create a new backend
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					// The state file could show either 0 or 1 backends in backend_set; depending on the order of operations.
					// If UpdateBackendSet happens first, then you will see 0. If CreateBackend happens first, then you will see 1.
					//resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.0.backend.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "weight", "1"),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "name"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
				),
				ExpectNonEmptyPlan: true,
			},
			// test ForceNew backend and update backendset (the operations may happen concurrently)
			// This test is needed because DeleteBackend, CreateBackend and UpdateBackendSet all update the same backend set
			// resource. Test that all changes are applied sequentially.
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 80
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					session_persistence_configuration {
						cookie_name = "lb-session2"
						disable_fallback = false
					}

					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.id}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 80
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					// The state file could show either 0 or 1 backends in backend_set; depending on the order of operations.
					// If UpdateBackendSet happens first, then you will see 0. If CreateBackend happens first, then you will see 1.
					//resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.0.backend.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "80"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "port", "80"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "weight", "1"),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "name"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestResourceLoadBalancerBackendSetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendSetTestSuite))
}
