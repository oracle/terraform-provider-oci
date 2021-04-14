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
	"github.com/oracle/oci-go-sdk/v39/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v39/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SslCipherSuiteResourceConfig = SslCipherSuiteResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", Optional, Update, sslCipherSuiteRepresentation)

	sslCipherSuiteSingularDataSourceRepresentation = map[string]interface{}{
		"name":             Representation{repType: Required, create: `example_cipher_suite`},
		"load_balancer_id": Representation{repType: Optional, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	sslCipherSuiteDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Optional, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, sslCipherSuiteDataSourceFilterRepresentation}}
	sslCipherSuiteDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite.name}`}},
	}

	sslCipherSuiteRepresentation = map[string]interface{}{
		"name":             Representation{repType: Required, create: `example_cipher_suite`},
		"ciphers":          Representation{repType: Required, create: []string{`AES128-SHA`, `AES256-SHA`}},
		"load_balancer_id": Representation{repType: Optional, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	SslCipherSuiteResourceDependencies = generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

func TestLoadBalancerSslCipherSuiteResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerSslCipherSuiteResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite"
	datasourceName := "data.oci_load_balancer_ssl_cipher_suites.test_ssl_cipher_suites"
	singularDatasourceName := "data.oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+SslCipherSuiteResourceDependencies+
		generateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", Optional, Create, sslCipherSuiteRepresentation), "loadbalancer", "sslCipherSuite", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerSslCipherSuiteDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SslCipherSuiteResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", Optional, Create, sslCipherSuiteRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "example_cipher_suite"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + SslCipherSuiteResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + SslCipherSuiteResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", Optional, Create, sslCipherSuiteRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_cipher_suite"),

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
					generateDataSourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suites", "test_ssl_cipher_suites", Optional, Update, sslCipherSuiteDataSourceRepresentation) +
					compartmentIdVariableStr + SslCipherSuiteResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", Optional, Update, sslCipherSuiteRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "ssl_cipher_suites.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ssl_cipher_suites.0.name", "example_cipher_suite"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_ssl_cipher_suite", "test_ssl_cipher_suite", Optional, Create, sslCipherSuiteSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SslCipherSuiteResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_cipher_suite"),

					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_cipher_suite"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + SslCipherSuiteResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"state",
					"ciphers",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckLoadBalancerSslCipherSuiteDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("LoadBalancerSslCipherSuite") {
		resource.AddTestSweepers("LoadBalancerSslCipherSuite", &resource.Sweeper{
			Name:         "LoadBalancerSslCipherSuite",
			Dependencies: DependencyGraph["sslCipherSuite"],
			F:            sweepLoadBalancerSslCipherSuiteResource,
		})
	}
}

func sweepLoadBalancerSslCipherSuiteResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	sslCipherSuiteIds, err := getSslCipherSuiteIds(compartment)
	if err != nil {
		return err
	}
	for _, sslCipherSuiteId := range sslCipherSuiteIds {
		if ok := SweeperDefaultResourceId[sslCipherSuiteId]; !ok {
			deleteSSLCipherSuiteRequest := oci_load_balancer.DeleteSSLCipherSuiteRequest{}

			deleteSSLCipherSuiteRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")
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
	ids := getResourceIdsToSweep(compartment, "SslCipherSuiteId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listSSLCipherSuitesRequest := oci_load_balancer.ListSSLCipherSuitesRequest{}
	listSSLCipherSuitesResponse, err := loadBalancerClient.ListSSLCipherSuites(context.Background(), listSSLCipherSuitesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SslCipherSuite list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sslCipherSuite := range listSSLCipherSuitesResponse.Items {
		id := *sslCipherSuite.Name
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "SslCipherSuiteId", id)
	}
	return resourceIds, nil
}
