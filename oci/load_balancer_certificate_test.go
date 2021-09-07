// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v47/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v47/loadbalancer"

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
		"ca_certificate":     Representation{repType: Optional, create: `${var.ca_certificate_value}`},
		"passphrase":         Representation{repType: Optional, create: `Mysecretunlockingcode42!1!`},
		"private_key":        Representation{repType: Optional, create: `${var.private_key_value}`},
		"public_certificate": Representation{repType: Optional, create: `${var.ca_certificate_value}`},
	}

	caCertificate            = getEnvSettingWithBlankDefault("ca_certificate")
	caCertificateVariableStr = fmt.Sprintf("variable \"ca_certificate_value\" { default = \"%s\" }\n", caCertificate)

	privateKeyData        = getEnvSettingWithBlankDefault("private_key_data")
	privateKeyVariableStr = fmt.Sprintf("variable \"private_key_value\" { default = \"%s\" }\n", privateKeyData)

	CertificateResourceDependencies = generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies + privateKeyVariableStr + caCertificateVariableStr
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerCertificateResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_certificate.test_certificate"
	datasourceName := "data.oci_load_balancer_certificates.test_certificates"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+CertificateResourceDependencies+
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Optional, Create, certificateRepresentation), "loadbalancer", "certificate", t)

	ResourceTest(t, testAccCheckLoadBalancerCertificateDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + CertificateResourceDependencies +
				generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
				resource.TestMatchResourceAttr(resourceName, "private_key", regexp.MustCompile("-----BEGIN RSA.*")),
				resource.TestMatchResourceAttr(resourceName, "public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),

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
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "certificates.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "certificates.0.ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(datasourceName, "certificates.0.certificate_name", "example_certificate_bundle"),
				resource.TestMatchResourceAttr(datasourceName, "certificates.0.public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
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
