// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
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
	LoadBalancerSslCipherSuiteResourceConfig = SslCipherSuiteResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", acctest.Optional, acctest.Update, LoadBalancerSslCipherSuiteRepresentation)

	LoadBalancerLoadBalancerSslCipherSuiteSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_cipher_suite`},
	}

	LoadBalancerLoadBalancerSslCipherSuiteDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: LoadBalancerSslCipherSuiteDataSourceFilterRepresentation},
	}
	LoadBalancerSslCipherSuiteDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite.name}`}},
	}

	LoadBalancerSslCipherSuiteRepresentation = map[string]interface{}{
		"ciphers":          acctest.Representation{RepType: acctest.Required, Create: []string{`AES128-SHA`, `AES256-SHA`}},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_cipher_suite`},
	}

	SslCipherSuiteResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerSslCipherSuiteResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerSslCipherSuiteResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite"
	datasourceName := "data.oci_load_balancer_ssl_cipher_suites.test_ssl_cipher_suites"
	singularDatasourceName := "data.oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SslCipherSuiteResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", acctest.Required, acctest.Create, LoadBalancerSslCipherSuiteRepresentation), "loadbalancer", "sslCipherSuite", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerSslCipherSuiteDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SslCipherSuiteResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", acctest.Optional, acctest.Create, LoadBalancerSslCipherSuiteRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ciphers.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_cipher_suite"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SslCipherSuiteResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SslCipherSuiteResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", acctest.Optional, acctest.Create, LoadBalancerSslCipherSuiteRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_cipher_suite"),

				func(s *terraform.State) (err error) {
					var resId string
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suites", "test_ssl_cipher_suites", acctest.Optional, acctest.Update, LoadBalancerLoadBalancerSslCipherSuiteDataSourceRepresentation) +
				compartmentIdVariableStr + SslCipherSuiteResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", acctest.Optional, acctest.Update, LoadBalancerSslCipherSuiteRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "ssl_cipher_suites.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ssl_cipher_suites.0.name", "example_cipher_suite"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", acctest.Optional, acctest.Create, LoadBalancerLoadBalancerSslCipherSuiteSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LoadBalancerSslCipherSuiteResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_cipher_suite"),

				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_cipher_suite"),
			),
		},
		// verify resource import
		{
			Config:            config + LoadBalancerSslCipherSuiteResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"state",
				"ciphers",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLoadBalancerSslCipherSuiteDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_ssl_cipher_suite" {
			noResourceFound = false
			request := oci_load_balancer.GetSSLCipherSuiteRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.Name = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

			_, err := client.GetSSLCipherSuite(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("LoadBalancerSslCipherSuite") {
		resource.AddTestSweepers("LoadBalancerSslCipherSuite", &resource.Sweeper{
			Name:         "LoadBalancerSslCipherSuite",
			Dependencies: acctest.DependencyGraph["sslCipherSuite"],
			F:            sweepLoadBalancerSslCipherSuiteResource,
		})
	}
}

func sweepLoadBalancerSslCipherSuiteResource(compartment string) error {
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
	sslCipherSuiteIds, err := getSslCipherSuiteIds(compartment)
	if err != nil {
		return err
	}
	for _, sslCipherSuiteId := range sslCipherSuiteIds {
		if ok := acctest.SweeperDefaultResourceId[sslCipherSuiteId]; !ok {
			deleteSSLCipherSuiteRequest := oci_load_balancer.DeleteSSLCipherSuiteRequest{}

			deleteSSLCipherSuiteRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteSSLCipherSuite(context.Background(), deleteSSLCipherSuiteRequest)
			if error != nil {
				fmt.Printf("Error deleting SslCipherSuite %s %s, It is possible that the resource is already deleted. Please verify manually \n", sslCipherSuiteId, error)
				continue
			}
		}
	}
	return nil
}

func getSslCipherSuiteIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SslCipherSuiteId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()

	listSSLCipherSuitesRequest := oci_load_balancer.ListSSLCipherSuitesRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for SslCipherSuite resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listSSLCipherSuitesRequest.LoadBalancerId = &loadBalancerId

		listSSLCipherSuitesResponse, err := loadBalancerClient.ListSSLCipherSuites(context.Background(), listSSLCipherSuitesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SslCipherSuite list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, sslCipherSuite := range listSSLCipherSuitesResponse.Items {
			id := *sslCipherSuite.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SslCipherSuiteId", id)
		}

	}
	return resourceIds, nil
}
