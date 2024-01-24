// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsNetworkPerimeterRequiredOnlyResource = IdentityDomainsNetworkPerimeterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Required, acctest.Create, IdentityDomainsNetworkPerimeterRepresentation)

	IdentityDomainsNetworkPerimeterResourceConfig = IdentityDomainsNetworkPerimeterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Optional, acctest.Update, IdentityDomainsNetworkPerimeterRepresentation)

	IdentityDomainsNetworkPerimeterSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"network_perimeter_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_network_perimeter.test_network_perimeter.id}`},
		"attribute_sets":       acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsNetworkPerimeterDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"network_perimeter_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"network_perimeter_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsNetworkPerimeterRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"ip_addresses":   acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsNetworkPerimeterIpAddressesRepresentation},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:NetworkPerimeter`}},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_id":    acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsNetworkPerimeterTagsRepresentation},
	}
	IdentityDomainsNetworkPerimeterIpAddressesRepresentation = map[string]interface{}{
		"value":   acctest.Representation{RepType: acctest.Required, Create: `192.0.2.0-192.0.2.255`, Update: `FE80:0000:0000:0000:0202:B3FF:FE1E:8329`},
		"type":    acctest.Representation{RepType: acctest.Optional, Create: `RANGE`, Update: `EXACT`},
		"version": acctest.Representation{RepType: acctest.Optional, Create: `IPV4`, Update: `IPV6`},
	}
	IdentityDomainsNetworkPerimeterTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsNetworkPerimeterResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsNetworkPerimeterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsNetworkPerimeterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_network_perimeter.test_network_perimeter"
	datasourceName := "data.oci_identity_domains_network_perimeters.test_network_perimeters"
	singularDatasourceName := "data.oci_identity_domains_network_perimeter.test_network_perimeter"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsNetworkPerimeterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Optional, acctest.Create, IdentityDomainsNetworkPerimeterRepresentation), "identitydomains", "networkPerimeter", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsNetworkPerimeterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNetworkPerimeterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Required, acctest.Create, IdentityDomainsNetworkPerimeterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.value", "192.0.2.0-192.0.2.255"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNetworkPerimeterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNetworkPerimeterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Optional, acctest.Create, IdentityDomainsNetworkPerimeterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.type", "RANGE"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.value", "192.0.2.0-192.0.2.255"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "networkPerimeters", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNetworkPerimeterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Optional, acctest.Update, IdentityDomainsNetworkPerimeterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.type", "EXACT"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.value", "FE80:0000:0000:0000:0202:B3FF:FE1E:8329"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.version", "IPV6"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_network_perimeters", "test_network_perimeters", acctest.Optional, acctest.Update, IdentityDomainsNetworkPerimeterDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsNetworkPerimeterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Optional, acctest.Update, IdentityDomainsNetworkPerimeterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "network_perimeter_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "network_perimeters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_perimeters.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_network_perimeter", "test_network_perimeter", acctest.Required, acctest.Create, IdentityDomainsNetworkPerimeterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsNetworkPerimeterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_perimeter_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.0.type", "EXACT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.0.value", "FE80:0000:0000:0000:0202:B3FF:FE1E:8329"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.0.version", "IPV6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsNetworkPerimeterRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_network_perimeter", "networkPerimeters"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsNetworkPerimeterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_network_perimeter" {
			noResourceFound = false
			request := oci_identity_domains.GetNetworkPerimeterRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.NetworkPerimeterId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetNetworkPerimeter(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsNetworkPerimeter") {
		resource.AddTestSweepers("IdentityDomainsNetworkPerimeter", &resource.Sweeper{
			Name:         "IdentityDomainsNetworkPerimeter",
			Dependencies: acctest.DependencyGraph["networkPerimeter"],
			F:            sweepIdentityDomainsNetworkPerimeterResource,
		})
	}
}

func sweepIdentityDomainsNetworkPerimeterResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	networkPerimeterIds, err := getIdentityDomainsNetworkPerimeterIds(compartment)
	if err != nil {
		return err
	}
	for _, networkPerimeterId := range networkPerimeterIds {
		if ok := acctest.SweeperDefaultResourceId[networkPerimeterId]; !ok {
			deleteNetworkPerimeterRequest := oci_identity_domains.DeleteNetworkPerimeterRequest{}

			deleteNetworkPerimeterRequest.NetworkPerimeterId = &networkPerimeterId

			deleteNetworkPerimeterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteNetworkPerimeter(context.Background(), deleteNetworkPerimeterRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkPerimeter %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkPerimeterId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsNetworkPerimeterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkPerimeterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listNetworkPerimetersRequest := oci_identity_domains.ListNetworkPerimetersRequest{}
	listNetworkPerimetersResponse, err := identityDomainsClient.ListNetworkPerimeters(context.Background(), listNetworkPerimetersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkPerimeter list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkPerimeter := range listNetworkPerimetersResponse.Resources {
		id := *networkPerimeter.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkPerimeterId", id)
	}
	return resourceIds, nil
}
