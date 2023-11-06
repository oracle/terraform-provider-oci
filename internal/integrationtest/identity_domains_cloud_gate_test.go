// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsCloudGateRequiredOnlyResource = IdentityDomainsCloudGateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Required, acctest.Create, IdentityDomainsCloudGateRepresentation)

	IdentityDomainsCloudGateResourceConfig = IdentityDomainsCloudGateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Optional, acctest.Update, IdentityDomainsCloudGateRepresentation)

	IdentityDomainsCloudGateSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_gate_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_cloud_gate.test_cloud_gate.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsCloudGateDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"cloud_gate_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"cloud_gate_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":       acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	IdentityDomainsCloudGateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_domains_cloud_gate.test_cloud_gate.id}`}},
	}

	IdentityDomainsCloudGateRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"idcs_endpoint":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":            acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:CloudGate`}},
		"active":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"attribute_sets":     acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"last_modified_time": acctest.Representation{RepType: acctest.Optional, Create: `2000-01-01T12:00:00Z`, Update: `2001-01-01T12:00:00Z`},
		"tags":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsCloudGateTagsRepresentation},
		"type":               acctest.Representation{RepType: acctest.Optional, Create: `lbaas`},
	}
	IdentityDomainsCloudGateRepresentationTypeGateway = acctest.RepresentationCopyWithNewProperties(IdentityDomainsCloudGateRepresentation, map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Optional, Create: `gateway`},
	})
	IdentityDomainsCloudGateTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsCloudGateResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsCloudGateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsCloudGateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_cloud_gate.test_cloud_gate"
	datasourceName := "data.oci_identity_domains_cloud_gates.test_cloud_gates"
	singularDatasourceName := "data.oci_identity_domains_cloud_gate.test_cloud_gate"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsCloudGateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Optional, acctest.Create, IdentityDomainsCloudGateRepresentation), "identitydomains", "cloudGate", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsCloudGateDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Required, acctest.Create, IdentityDomainsCloudGateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies,
		},
		// verify Create with gateway type
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Optional, acctest.Create, IdentityDomainsCloudGateRepresentationTypeGateway),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "last_modified_time", "2000-01-01T12:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "ocid"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "type", "gateway"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Optional, acctest.Create, IdentityDomainsCloudGateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "last_modified_time", "2000-01-01T12:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "type", "lbaas"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "cloudGates", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Optional, acctest.Update, IdentityDomainsCloudGateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "last_modified_time", "2001-01-01T12:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "type", "lbaas"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					print("resId: " + resId + ", resId2: " + resId2)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_cloud_gates", "test_cloud_gates", acctest.Optional, acctest.Update, IdentityDomainsCloudGateDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsCloudGateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Optional, acctest.Update, IdentityDomainsCloudGateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "cloud_gate_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "cloud_gates.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "cloud_gates.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Required, acctest.Create, IdentityDomainsCloudGateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsCloudGateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_gate_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_modified_time", "2001-01-01T12:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocid"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "lbaas"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsCloudGateRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_cloud_gate", "cloudGates"),
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

func testAccCheckIdentityDomainsCloudGateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_cloud_gate" {
			noResourceFound = false
			request := oci_identity_domains.GetCloudGateRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.CloudGateId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetCloudGate(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsCloudGate") {
		resource.AddTestSweepers("IdentityDomainsCloudGate", &resource.Sweeper{
			Name:         "IdentityDomainsCloudGate",
			Dependencies: acctest.DependencyGraph["cloudGate"],
			F:            sweepIdentityDomainsCloudGateResource,
		})
	}
}

func sweepIdentityDomainsCloudGateResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	cloudGateIds, err := getIdentityDomainsCloudGateIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudGateId := range cloudGateIds {
		if ok := acctest.SweeperDefaultResourceId[cloudGateId]; !ok {
			deleteCloudGateRequest := oci_identity_domains.DeleteCloudGateRequest{}

			deleteCloudGateRequest.CloudGateId = &cloudGateId

			deleteCloudGateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteCloudGate(context.Background(), deleteCloudGateRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudGate %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudGateId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsCloudGateIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudGateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listCloudGatesRequest := oci_identity_domains.ListCloudGatesRequest{}
	listCloudGatesResponse, err := identityDomainsClient.ListCloudGates(context.Background(), listCloudGatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudGate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudGate := range listCloudGatesResponse.Resources {
		id := *cloudGate.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudGateId", id)
	}
	return resourceIds, nil
}
