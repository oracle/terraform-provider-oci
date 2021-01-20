// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v33/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v33/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
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
		"ca_certificate":     Representation{repType: Optional, create: `-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----`},
		"passphrase":         Representation{repType: Optional, create: `Mysecretunlockingcode42!1!`},
		"private_key":        Representation{repType: Optional, create: `-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAo83kaUQXpCcSoEuRVFX3jztWDNKtWpjNG240f0RpERI1NnZt\nHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19geIXR6TeavT+W5iRh4goK+N7gubYk\nSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCtCgd4MDlsvLv/YHCLvJL4JgRxKyev\njlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc+Je9EC3MWWxd5jBwXu3vgIYRuGR4\nDPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJKN0NXp5obaQToYqMsvAZyHoEyfCB\nDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t58QIDAQABAoIBADIyHuOPJTt9abzL\nS26vpVw0D6uAR/UyS/Ay9k1ltliv3rSg19DaHlwLjPwqnvCx7jBgTeVCYZhAkvgx\nkSsGDDcCsw+npXiG6wP9dC1jbHdVPUJLqZTPqB6sZCu8bM9RIE4Z/DcUY+HRN3qh\nmoh5wn0HSvJkNokjhx+TfY687uQfDMu0de4V2UPScZ7mboCu9HqK9qu0/krdTMH1\nrtnnFGEnx/Pe38YJl0fWxo8BHKHprwEvWX0MQzQeklnUtxREMuofSAOBe/I2DJGh\n1I94b6I66ypxuX0qAozT1MPbJGuaR+puyKawLNAQmZa9pgrrFK7e8PQUzrGVpVCp\nFtwx420CgYEA0uX/G0ycia0UTdkxkIsKIiLjs12LC0XmYjiWgkoL0PjiZzcPITn6\nvqqqGSz44HwtbrttZPm3Mo79yJ5xFiHCX0vFJykgy6cfS94imMgm8qIOS0bXjX7w\nxH2BOgp0H32LP/Zt7owcWJLEIQCjj0/4+Nvu0GskGVHlE8EYrXWf1E8CgYEAxtWk\nxBo52uNXL712VGDPNxprVGUpWSbkXpE+7wtRso5LnAnAj6dpmZsGe2zaYjRIt3Ls\nGnno4HUmwpQ5yXlHFpDUJvb2soXq3afnuAh5aVu6aKRQoG/5o3cD4pOupNbjDDNs\nTVLtTLIAIYDbph/j7pV/JnJ2WHcdk6WiVJoW/b8CgYAopLZzNmJ8jeR51D+fEYyU\nY5DqQj7Hn2L0zt8CoO6CCVToe03pI1lVYWKCk44rBQNkca51ZUKO9cum3BIDJ+Jj\npyCJmX1+geigIGEefIQ1AlIq464q0Knp1B4RZ25Vm0Y4v28UJ+BWmYI+sfbTaaAb\npZbyh5NfZc717aKp2x9ANQKBgHQpvOkUqVhIGVe6yLbjGCyJMstLjqyXHDRjhvEB\nG+nFWEcBK47Br+Adwdu57JwTD6ida3LMZlE8IDjtgBVE1VNJqahaACasNlrpDWdn\nDAeRn4Yi+TfCM4Zcsdhdj1qecGdgY5WJLTnxhEIOlkSnvPJWRMKhfKKSdKUdz4i9\nvVDhAoGAEHxfhFLVwdTa0RMUq3KYSXa5WqLANRn2e62Cc3eUWekcUjbIATRF5AIo\nm0WS+rURZWy1Fd6fGg8sRHock0+vxwqeP6OlyW4tJMhL33NrNbgyvkXlMMIX6bC4\nUq8aAew0B3j61UUsTqhHMhYwIS3GOIHx/O10wwINPnUMIVER3Wg=\n-----END RSA PRIVATE KEY-----`},
		"public_certificate": Representation{repType: Optional, create: `-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----`},
	}

	CertificateResourceDependencies = generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

func TestLoadBalancerCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerCertificateResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_certificate.test_certificate"
	datasourceName := "data.oci_load_balancer_certificates.test_certificates"

	var resId string

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
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAo83kaUQXpCcSoEuRVFX3jztWDNKtWpjNG240f0RpERI1NnZt\nHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19geIXR6TeavT+W5iRh4goK+N7gubYk\nSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCtCgd4MDlsvLv/YHCLvJL4JgRxKyev\njlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc+Je9EC3MWWxd5jBwXu3vgIYRuGR4\nDPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJKN0NXp5obaQToYqMsvAZyHoEyfCB\nDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t58QIDAQABAoIBADIyHuOPJTt9abzL\nS26vpVw0D6uAR/UyS/Ay9k1ltliv3rSg19DaHlwLjPwqnvCx7jBgTeVCYZhAkvgx\nkSsGDDcCsw+npXiG6wP9dC1jbHdVPUJLqZTPqB6sZCu8bM9RIE4Z/DcUY+HRN3qh\nmoh5wn0HSvJkNokjhx+TfY687uQfDMu0de4V2UPScZ7mboCu9HqK9qu0/krdTMH1\nrtnnFGEnx/Pe38YJl0fWxo8BHKHprwEvWX0MQzQeklnUtxREMuofSAOBe/I2DJGh\n1I94b6I66ypxuX0qAozT1MPbJGuaR+puyKawLNAQmZa9pgrrFK7e8PQUzrGVpVCp\nFtwx420CgYEA0uX/G0ycia0UTdkxkIsKIiLjs12LC0XmYjiWgkoL0PjiZzcPITn6\nvqqqGSz44HwtbrttZPm3Mo79yJ5xFiHCX0vFJykgy6cfS94imMgm8qIOS0bXjX7w\nxH2BOgp0H32LP/Zt7owcWJLEIQCjj0/4+Nvu0GskGVHlE8EYrXWf1E8CgYEAxtWk\nxBo52uNXL712VGDPNxprVGUpWSbkXpE+7wtRso5LnAnAj6dpmZsGe2zaYjRIt3Ls\nGnno4HUmwpQ5yXlHFpDUJvb2soXq3afnuAh5aVu6aKRQoG/5o3cD4pOupNbjDDNs\nTVLtTLIAIYDbph/j7pV/JnJ2WHcdk6WiVJoW/b8CgYAopLZzNmJ8jeR51D+fEYyU\nY5DqQj7Hn2L0zt8CoO6CCVToe03pI1lVYWKCk44rBQNkca51ZUKO9cum3BIDJ+Jj\npyCJmX1+geigIGEefIQ1AlIq464q0Knp1B4RZ25Vm0Y4v28UJ+BWmYI+sfbTaaAb\npZbyh5NfZc717aKp2x9ANQKBgHQpvOkUqVhIGVe6yLbjGCyJMstLjqyXHDRjhvEB\nG+nFWEcBK47Br+Adwdu57JwTD6ida3LMZlE8IDjtgBVE1VNJqahaACasNlrpDWdn\nDAeRn4Yi+TfCM4Zcsdhdj1qecGdgY5WJLTnxhEIOlkSnvPJWRMKhfKKSdKUdz4i9\nvVDhAoGAEHxfhFLVwdTa0RMUq3KYSXa5WqLANRn2e62Cc3eUWekcUjbIATRF5AIo\nm0WS+rURZWy1Fd6fGg8sRHock0+vxwqeP6OlyW4tJMhL33NrNbgyvkXlMMIX6bC4\nUq8aAew0B3j61UUsTqhHMhYwIS3GOIHx/O10wwINPnUMIVER3Wg=\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
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
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.ca_certificate", "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.public_certificate", "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"passphrase",
					"private_key",
					"state",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckLoadBalancerCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_certificate" {
			noResourceFound = false
			request := oci_load_balancer.ListCertificatesRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("LoadBalancerCertificate") {
		resource.AddTestSweepers("LoadBalancerCertificate", &resource.Sweeper{
			Name:         "LoadBalancerCertificate",
			Dependencies: DependencyGraph["certificate"],
			F:            sweepLoadBalancerCertificateResource,
		})
	}
}

func sweepLoadBalancerCertificateResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	certificateIds, err := getLBCertificateIds(compartment)
	if err != nil {
		return err
	}
	for _, certificateId := range certificateIds {
		if ok := SweeperDefaultResourceId[certificateId]; !ok {
			deleteCertificateRequest := oci_load_balancer.DeleteCertificateRequest{}

			deleteCertificateRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteCertificate(context.Background(), deleteCertificateRequest)
			if error != nil {
				fmt.Printf("Error deleting Certificate %s %s, It is possible that the resource is already deleted. Please verify manually \n", certificateId, error)
				continue
			}
		}
	}
	return nil
}

func getLBCertificateIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CertificateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listCertificatesRequest := oci_load_balancer.ListCertificatesRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for Certificate resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listCertificatesRequest.LoadBalancerId = &loadBalancerId

		listCertificatesResponse, err := loadBalancerClient.ListCertificates(context.Background(), listCertificatesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Certificate list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, certificate := range listCertificatesResponse.Items {
			id := *certificate.CertificateName
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "CertificateId", id)
		}

	}
	return resourceIds, nil
}
