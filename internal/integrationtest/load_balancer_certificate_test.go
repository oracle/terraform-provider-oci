// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CertificateRequiredOnlyResource = CertificateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation)

	certificateDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateDataSourceFilterRepresentation}}
	certificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `certificate_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_certificate.test_certificate.certificate_name}`}},
	}

	certificateRepresentation = map[string]interface{}{
		"certificate_name": acctest.Representation{RepType: acctest.Required, Create: `example_certificate_bundle`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"ca_certificate":   acctest.Representation{RepType: acctest.Optional, Create: caCertificate},
		// We don't test with private keys but if we were to do so we would have to set this.
		// "passphrase":         acctest.Representation{RepType: acctest.Optional, Create: `Mysecretunlockingcode42!1!`},
		"private_key":        acctest.Representation{RepType: acctest.Optional, Create: privateKeyData},
		"public_certificate": acctest.Representation{RepType: acctest.Optional, Create: caCertificate},
	}

	// The following assumes you set the TF_VAR_ca_certificate variable to something like the following:
	//     export TF_VAR_private_key_value="$(cat ~/certificate/example_2.com.key)"
	// which results in
	//     $ printenv TF_VAR_ca_certificate
	//     -----BEGIN CERTIFICATE-----
	//     MIIFRDCCAyygAwIBAgIUDB9s8795KLpchjLPGFI9sqdVaT4wDQYJKoZIhvcNAQEL
	//     ...
	//     eaiXT7X2gvU=
	//     -----END CERTIFICATE-----
	// We want the string
	//     "-----BEGIN CERTIFICATE-----\\nMIIFRDCC...\\neaiXT7X2gvU=\\n-----END CERTIFICATE-----"
	caCertificate            = strings.ReplaceAll(utils.GetEnvSettingWithBlankDefault("ca_certificate"), "\n", "\\n")
	caCertificateVariableStr = fmt.Sprintf("variable \"ca_certificate_value\" { default = \"%s\" }\n", caCertificate)

	privateKeyData        = strings.ReplaceAll(utils.GetEnvSettingWithBlankDefault("private_key_data"), "\n", "\\n")
	privateKeyVariableStr = fmt.Sprintf("variable \"private_key_value\" { default = \"%s\" }\n", privateKeyData)

	CertificateResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies + privateKeyVariableStr + caCertificateVariableStr
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerCertificateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_certificate.test_certificate"
	datasourceName := "data.oci_load_balancer_certificates.test_certificates"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CertificateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Optional, acctest.Create, certificateRepresentation), "loadbalancer", "certificate", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerCertificateDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CertificateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CertificateResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CertificateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Optional, acctest.Create, certificateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				// We don't test with private keys but if we were to do so we would have to set
				// resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
				// resource.TestMatchResourceAttr(resourceName, "private_key", regexp.MustCompile("-----BEGIN PRIVATE ENCRYPTED KEY.*")),
				resource.TestMatchResourceAttr(resourceName, "private_key", regexp.MustCompile("-----BEGIN PRIVATE KEY.*")),
				resource.TestMatchResourceAttr(resourceName, "public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_certificates", "test_certificates", acctest.Optional, acctest.Update, certificateDataSourceRepresentation) +
				compartmentIdVariableStr + CertificateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Optional, acctest.Update, certificateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "certificates.#", "1"),
				resource.TestMatchResourceAttr(datasourceName, "certificates.0.ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				resource.TestCheckResourceAttr(datasourceName, "certificates.0.certificate_name", "example_certificate_bundle"),
				resource.TestMatchResourceAttr(datasourceName, "certificates.0.public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
			),
		},
		// verify resource import
		{
			Config:            config + CertificateRequiredOnlyResource,
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_certificate" {
			noResourceFound = false
			request := oci_load_balancer.ListCertificatesRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("LoadBalancerCertificate") {
		resource.AddTestSweepers("LoadBalancerCertificate", &resource.Sweeper{
			Name:         "LoadBalancerCertificate",
			Dependencies: acctest.DependencyGraph["certificate"],
			F:            sweepLoadBalancerCertificateResource,
		})
	}
}

func sweepLoadBalancerCertificateResource(compartment string) error {
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
	certificateIds, err := getLBCertificateIds(compartment)
	if err != nil {
		return err
	}
	for _, certificateId := range certificateIds {
		if ok := acctest.SweeperDefaultResourceId[certificateId]; !ok {
			deleteCertificateRequest := oci_load_balancer.DeleteCertificateRequest{}

			deleteCertificateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
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
	ids := acctest.GetResourceIdsToSweep(compartment, "CertificateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()

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
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CertificateId", id)
		}

	}
	return resourceIds, nil
}
