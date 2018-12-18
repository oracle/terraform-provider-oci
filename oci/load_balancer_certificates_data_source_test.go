// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerCertificates_basic(t *testing.T) {
	providers := testAccProviders
	config := legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block = "10.0.0.0/24"
		display_name = "-tf-subnet"
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
	}
	
	data "oci_load_balancer_certificates" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
	}`

	resourceName := "data.oci_load_balancer_certificates.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "certificates.#", "1"),
					resource.TestMatchResourceAttr(resourceName, "certificates.0.ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
					resource.TestCheckResourceAttr(resourceName, "certificates.0.certificate_name", "tf_cert_name"),
					resource.TestMatchResourceAttr(resourceName, "certificates.0.public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				),
			},
		},
	})
}
