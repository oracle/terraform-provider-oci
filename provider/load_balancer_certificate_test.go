// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CertificateRequiredOnlyResource = CertificateResourceDependencies + `
resource "oci_load_balancer_certificate" "test_certificate" {
	#Required
	certificate_name = "${var.certificate_certificate_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
`

	CertificateResourceConfig = CertificateResourceDependencies + `
resource "oci_load_balancer_certificate" "test_certificate" {
	#Required
	certificate_name = "${var.certificate_certificate_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

	#Optional
	ca_certificate = "${var.certificate_ca_certificate}"
	passphrase = "${var.certificate_passphrase}"
	private_key = "${var.certificate_private_key}"
	public_certificate = "${var.certificate_public_certificate}"
}
`
	CertificatePropertyVariables = `
variable "certificate_ca_certificate" { default = "BEGIN CERTIFICATEMIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD...AkGA1UEBhMCR0IxEND CERTIFICATE" }
variable "certificate_certificate_name" { default = "My_certificate_bundle" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY" }
variable "certificate_public_certificate" { default = "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE" }

`
	CertificateResourceDependencies = LoadBalancerPropertyVariables + LoadBalancerResourceConfig
)

func TestLoadBalancerCertificateResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_load_balancer_certificate.test_certificate"
	datasourceName := "data.oci_load_balancer_certificates.test_certificates"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + CertificatePropertyVariables + compartmentIdVariableStr + CertificateRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CertificateResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + CertificatePropertyVariables + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "BEGIN CERTIFICATEMIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD...AkGA1UEBhMCR0IxEND CERTIFICATE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "passphrase2" }
variable "certificate_private_key" { default = "privateKey2" }
variable "certificate_public_certificate" { default = "publicCertificate2" }

                ` + compartmentIdVariableStr2 + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "passphrase2"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "privateKey2"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "publicCertificate2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "passphrase2" }
variable "certificate_private_key" { default = "privateKey2" }
variable "certificate_public_certificate" { default = "publicCertificate2" }

data "oci_load_balancer_certificates" "test_certificates" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

    filter {
    	name = "id"
    	values = ["${oci_load_balancer_certificate.test_certificate.id}"]
    }
}
                ` + compartmentIdVariableStr2 + CertificateResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "certificates.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.load_balancer_id"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.public_certificate", "publicCertificate2"),
				),
			},
		},
	})
}

func TestLoadBalancerCertificateResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_certificate.test_certificate"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + CertificatePropertyVariables + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "BEGIN CERTIFICATEMIIEczCCA1ugAwIBAgIBADANBgkqhkiG9w0BAQQFAD...AkGA1UEBhMCR0IxEND CERTIFICATE"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "My_certificate_bundle" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY" }
variable "certificate_public_certificate" { default = "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "My_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CaCertificate but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY" }
variable "certificate_public_certificate" { default = "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CertificateName but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY" }
variable "certificate_public_certificate" { default = "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter LoadBalancerId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "passphrase2" }
variable "certificate_private_key" { default = "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY" }
variable "certificate_public_certificate" { default = "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "passphrase2"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "BEGIN RSA PRIVATE KEYJO1O1v2ftXMsawM90tnXwc6xhOAT1gDBC9S8DKeca...JZNUgYYwNS0dP2UKEND RSA PRIVATE KEY"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter Passphrase but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "passphrase2" }
variable "certificate_private_key" { default = "privateKey2" }
variable "certificate_public_certificate" { default = "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "passphrase2"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "privateKey2"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "BEGIN CERTIFICATEMIIC2jCCAkMCAg38MA0GCSqGSIb3DQEBBQUAMIGbM...QswCQYDVQQGEwJKUDEOMAwGEND CERTIFICATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter PrivateKey but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "certificate_ca_certificate" { default = "caCertificate2" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "passphrase2" }
variable "certificate_private_key" { default = "privateKey2" }
variable "certificate_public_certificate" { default = "publicCertificate2" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "caCertificate2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "passphrase2"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "privateKey2"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "publicCertificate2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter PublicCertificate but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
