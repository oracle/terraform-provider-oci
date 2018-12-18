// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

var (
	CertificateRequiredOnlyResource = CertificateResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation)

	certificateDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, certificateDataSourceFilterRepresentation}}
	certificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `certificate_name`},
		"values": Representation{repType: Required, create: []string{`${oci_load_balancer_certificate.test_certificate.certificate_name}`}},
	}

	certificateRepresentation = map[string]interface{}{
		"certificate_name":   Representation{repType: Required, create: `example_certificate_bundle`},
		"load_balancer_id":   Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"ca_certificate":     Representation{repType: Optional, create: `-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----`},
		"passphrase":         Representation{repType: Optional, create: `Mysecretunlockingcode42!1!`},
		"private_key":        Representation{repType: Optional, create: `-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1vqH90VOIZHh/FL9pEH23op3t+FcMBSKQ+ijCoDZfil5m12X\nUQWUTS/ZFauGo0q1O4esHI8vusL6eyBVDOiTvyGfgXBYWq+yPkIGlqcuOdxN/bQv\nhfAJN5hcJYjosT/KhlwuuYO7V4ZWcgpacoyETh5bMinEPJ1k7e6dF4/4p3lisERP\n2zkxDtP98ZlUbo0ZF6sXPe9PJqYUCk0DXkH1GkffX6Zp+39q+ywJVP6Up54USuZ5\nmCJRyrGn0hq5+ha/UNLm6ggxa3sE2+/Rzx9sFrWLuxuEKQJJ4KPBuNreIHQk2z6K\nTH316SplmZRgTXqotTpYgWF9Ml/14ljvlDo6RwIDAQABAoIBAQCFzHXVQ1BGenpR\nRgHROrEAfuPWETAESLRpYaAgCGPVLtEeDpj/913+0FnnL9NjTDsR6vYG7GNDdNja\nyxvEJfjWy4Fv2VFUV+ey8fsRxslxf5kW3w946BWEgZJQVi6lKtPM3hDCq6ds6RJi\ndeknRCeQSzptNSuKoldP8uPY52VWLYTyy/ODwtSCFZKTm7iTD3RkpAqNMPs2V8EA\nRbjiu70q4Kk+ozHQ/0wtOSjZinR6LW7e+6bXxdVV0hdt52h2YzdQrQkHu5ATm54P\nm1S9PbiyV06BSuU37oZbyWiblP6rsJIqucEXSSmTY/5PCG+huzTseig1bot19eD0\nwEs9YAixAoGBAPgKMZ+VhA3eoAx21R62SmJJKdCYZ+qjPmvUt4tW3dJxYTweJo6A\nwyh+p3VsBwE6L8hPD1OyEvUrl64uQEMYSMmNaLpaMqNbytP8h6uiFGyRUSnkYvxF\nPcKW6UjKZNyGWT+dOnJN07DnwvZgOx2JZJkgFpaAEL6g0Hzm1bI/6dkPAoGBAN3g\nt+0sDfgYr3raGAPLQJCVgHV3MdIDHP/ebZBT1NTcxL7Wf/+0WAvnwD7DPsXj84JG\nzNMk7+EzwClAWGAQJMymC9NltfygyI0JjI+88nVk3mrpDq/zqR/vkn2R1T78No/X\nEWTnMmHlzDg4HsGTWKDg7jrmYSas0NvMHPtBbVtJAoGAKV0B03wKjomOpSV3+uwp\nUWSkDX4s7isU8MSDa0AsM7jmnzDj+yWr5efhIyrFrEW4zC2q/6kVkj8Xx1s9KjM1\niC8FxPXftfBLzbgyI8QepdBB+bt1al5do0KpWpMt6Lyay4n7wi4KXFj54T5A/Xb5\nCLQaMDThFfkZa4rPHi+cXq8CgYAfk665a06lo2W99zH5wEB1E0HP9eG6QMUsyQwQ\nwU2F6dF6U26uBo2NTDM4+3KAmVt7i/X0iso047eSZ1zsdv+1vF/sewo2ZO+F2vkN\nL9fVy0A4OOjlM6k7KU5Q3qNZrm1ZdUM9eAXclubEjYAbDoxLgReGfGkRJwEmdtsd\nCwe0OQKBgQD3m85OXSSf+xlm7tGO66bcxHifkp4XfkqWxFwpfYkxNtYZfFpN8jL5\niS0OyLldmJbCVB6EIs5ylW86aeZMH/JecPTxOnaT4qc43PMrLi4MSa65Gp/Zgs1U\nyO0hfWlpH2ncUIuQEksXEPSKQUjvdQl7pD8kghCbDYbm3zsjw3rkyA==\n-----END RSA PRIVATE KEY-----`},
		"public_certificate": Representation{repType: Optional, create: `-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----`},
	}

	CertificateResourceDependencies = LoadBalancerResourceConfig
)

func TestLoadBalancerCertificateResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_certificate.test_certificate"
	datasourceName := "data.oci_load_balancer_certificates.test_certificates"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerCertificateDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CertificateResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Optional, Create, certificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1vqH90VOIZHh/FL9pEH23op3t+FcMBSKQ+ijCoDZfil5m12X\nUQWUTS/ZFauGo0q1O4esHI8vusL6eyBVDOiTvyGfgXBYWq+yPkIGlqcuOdxN/bQv\nhfAJN5hcJYjosT/KhlwuuYO7V4ZWcgpacoyETh5bMinEPJ1k7e6dF4/4p3lisERP\n2zkxDtP98ZlUbo0ZF6sXPe9PJqYUCk0DXkH1GkffX6Zp+39q+ywJVP6Up54USuZ5\nmCJRyrGn0hq5+ha/UNLm6ggxa3sE2+/Rzx9sFrWLuxuEKQJJ4KPBuNreIHQk2z6K\nTH316SplmZRgTXqotTpYgWF9Ml/14ljvlDo6RwIDAQABAoIBAQCFzHXVQ1BGenpR\nRgHROrEAfuPWETAESLRpYaAgCGPVLtEeDpj/913+0FnnL9NjTDsR6vYG7GNDdNja\nyxvEJfjWy4Fv2VFUV+ey8fsRxslxf5kW3w946BWEgZJQVi6lKtPM3hDCq6ds6RJi\ndeknRCeQSzptNSuKoldP8uPY52VWLYTyy/ODwtSCFZKTm7iTD3RkpAqNMPs2V8EA\nRbjiu70q4Kk+ozHQ/0wtOSjZinR6LW7e+6bXxdVV0hdt52h2YzdQrQkHu5ATm54P\nm1S9PbiyV06BSuU37oZbyWiblP6rsJIqucEXSSmTY/5PCG+huzTseig1bot19eD0\nwEs9YAixAoGBAPgKMZ+VhA3eoAx21R62SmJJKdCYZ+qjPmvUt4tW3dJxYTweJo6A\nwyh+p3VsBwE6L8hPD1OyEvUrl64uQEMYSMmNaLpaMqNbytP8h6uiFGyRUSnkYvxF\nPcKW6UjKZNyGWT+dOnJN07DnwvZgOx2JZJkgFpaAEL6g0Hzm1bI/6dkPAoGBAN3g\nt+0sDfgYr3raGAPLQJCVgHV3MdIDHP/ebZBT1NTcxL7Wf/+0WAvnwD7DPsXj84JG\nzNMk7+EzwClAWGAQJMymC9NltfygyI0JjI+88nVk3mrpDq/zqR/vkn2R1T78No/X\nEWTnMmHlzDg4HsGTWKDg7jrmYSas0NvMHPtBbVtJAoGAKV0B03wKjomOpSV3+uwp\nUWSkDX4s7isU8MSDa0AsM7jmnzDj+yWr5efhIyrFrEW4zC2q/6kVkj8Xx1s9KjM1\niC8FxPXftfBLzbgyI8QepdBB+bt1al5do0KpWpMt6Lyay4n7wi4KXFj54T5A/Xb5\nCLQaMDThFfkZa4rPHi+cXq8CgYAfk665a06lo2W99zH5wEB1E0HP9eG6QMUsyQwQ\nwU2F6dF6U26uBo2NTDM4+3KAmVt7i/X0iso047eSZ1zsdv+1vF/sewo2ZO+F2vkN\nL9fVy0A4OOjlM6k7KU5Q3qNZrm1ZdUM9eAXclubEjYAbDoxLgReGfGkRJwEmdtsd\nCwe0OQKBgQD3m85OXSSf+xlm7tGO66bcxHifkp4XfkqWxFwpfYkxNtYZfFpN8jL5\niS0OyLldmJbCVB6EIs5ylW86aeZMH/JecPTxOnaT4qc43PMrLi4MSa65Gp/Zgs1U\nyO0hfWlpH2ncUIuQEksXEPSKQUjvdQl7pD8kghCbDYbm3zsjw3rkyA==\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_certificates", "test_certificates", Optional, Update, certificateDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Optional, Update, certificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "certificates.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.public_certificate", "-----BEGIN CERTIFICATE-----\nMIIDIjCCAgoCCQCjzpcCmaYA6zANBgkqhkiG9w0BAQsFADBTMQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTEUMBIGA1UEAwwLY29tcGFueS5jb20wHhcNMTgxMjE4MDAwMDA0WhcNMTkwMTE3\nMDAwMDA0WjBTMQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1Nl\nYXR0bGUxDzANBgNVBAoMBk9yYWNsZTEUMBIGA1UEAwwLY29tcGFueS5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDW+of3RU4hkeH8Uv2kQfbeine3\n4VwwFIpD6KMKgNl+KXmbXZdRBZRNL9kVq4ajSrU7h6wcjy+6wvp7IFUM6JO/IZ+B\ncFhar7I+QgaWpy453E39tC+F8Ak3mFwliOixP8qGXC65g7tXhlZyClpyjIROHlsy\nKcQ8nWTt7p0Xj/ineWKwRE/bOTEO0/3xmVRujRkXqxc9708mphQKTQNeQfUaR99f\npmn7f2r7LAlU/pSnnhRK5nmYIlHKsafSGrn6Fr9Q0ubqCDFrewTb79HPH2wWtYu7\nG4QpAkngo8G42t4gdCTbPopMffXpKmWZlGBNeqi1OliBYX0yX/XiWO+UOjpHAgMB\nAAEwDQYJKoZIhvcNAQELBQADggEBAFKAGeN8m+zohW2BPmozh4GCdpH7dtHc9gTi\nPCYoj2uUJJs2KUOprzShhnpWtGhM+KC23KHM+nSRaGSsM55Z+SLbWvuYjnUbhQ/M\nBPTIAyrXluiaGt/jf6UWmh9u4xia9QipsWFgEXUNGDwwQU4M424/6xhZDUE/3Gfg\nPsLCXmQxbZvzAuhv2jnDi/xisPiYkdXhPoPMJD7S0CpwRkuVh/jKzkt9smaxQd/C\nB1yzgy8V3dN2VMH4WAIIfBMBDzpJoF+JT23KF3OuxXAn7IoF+ubb8RecAopDyLrh\nKRX6pjCTqyQTSL+9USXsDzCCxQVsMoap/Mi+sWKjbMI25Gu4r2Y=\n-----END CERTIFICATE-----"),
				),
			},
		},
	})
}

func testAccCheckLoadBalancerCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_certificate" {
			noResourceFound = false
			request := oci_load_balancer.ListCertificatesRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			response, err := client.ListCertificates(context.Background(), request)

			if err == nil {
				certificateName := rs.Primary.Attributes["certificate_name"]
				for _, item := range response.Items {
					if *item.CertificateName == certificateName {
						return fmt.Errorf("item still exists")
					}
				}
				// no error and item not found, that means item is deleted. continue checking next one
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
