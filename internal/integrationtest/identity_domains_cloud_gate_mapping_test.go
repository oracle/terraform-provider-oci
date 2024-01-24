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
	IdentityDomainsCloudGateMappingRequiredOnlyResource = IdentityDomainsCloudGateMappingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Required, acctest.Create, IdentityDomainsCloudGateMappingRepresentation)

	IdentityDomainsCloudGateMappingResourceConfig = IdentityDomainsCloudGateMappingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Optional, acctest.Update, IdentityDomainsCloudGateMappingRepresentation)

	IdentityDomainsCloudGateMappingSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_gate_mapping_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_cloud_gate_mapping.test_cloud_gate_mapping.id}`},
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":        acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsCloudGateMappingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"cloud_gate_mapping_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"cloud_gate_mapping_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":            acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":               acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsCloudGateMappingRepresentation = map[string]interface{}{
		"cloud_gate":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsCloudGateMappingCloudGateRepresentation},
		"gateway_app":     acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsCloudGateMappingGatewayAppRepresentation},
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"policy_name":     acctest.Representation{RepType: acctest.Required, Create: `default`},
		"resource_prefix": acctest.Representation{RepType: acctest.Required, Create: `resourcePrefix`, Update: `resourcePrefix2`},
		"schemas":         acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:CloudGateMapping`}},
		"server":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsCloudGateMappingServerRepresentation},
		"proxy_pass":      acctest.Representation{RepType: acctest.Required, Create: `https://www.oracle.com:443`},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"nginx_settings":  acctest.Representation{RepType: acctest.Optional, Create: `nginxSettings`, Update: `nginxSettings2`},
		"tags":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsCloudGateMappingTagsRepresentation},
	}
	IdentityDomainsCloudGateMappingCloudGateRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_cloud_gate.test_cloud_gate.id}`},
	}

	IdentityDomainsCloudGateMappingsAppRepresentation = map[string]interface{}{
		"based_on_template": acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAppBasedOnTemplateRepresentation},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `appDisplayName`},
		"idcs_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":           acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:App`}},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `name`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsApp},
	}
	IdentityDomainsCloudGateMappingGatewayAppRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_app.test_app.id}`},
	}
	IdentityDomainsCloudGateMappingServerRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_cloud_gate_server.test_cloud_gate_server.id}`},
	}
	IdentityDomainsCloudGateMappingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsCloudGateMappingResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Required, acctest.Create, IdentityDomainsCloudGateRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_app", "test_app", acctest.Required, acctest.Create, IdentityDomainsCloudGateMappingsAppRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Required, acctest.Create, IdentityDomainsCloudGateServerRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsCloudGateMappingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsCloudGateMappingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_cloud_gate_mapping.test_cloud_gate_mapping"
	datasourceName := "data.oci_identity_domains_cloud_gate_mappings.test_cloud_gate_mappings"
	singularDatasourceName := "data.oci_identity_domains_cloud_gate_mapping.test_cloud_gate_mapping"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsCloudGateMappingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Optional, acctest.Create, IdentityDomainsCloudGateMappingRepresentation), "identitydomains", "cloudGateMapping", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsCloudGateMappingDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Required, acctest.Create, IdentityDomainsCloudGateMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway_app.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_name"),
				resource.TestCheckResourceAttr(resourceName, "resource_prefix", "resourcePrefix"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "server.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateMappingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Optional, acctest.Create, IdentityDomainsCloudGateMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "gateway_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "gateway_app.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_opc_service", "false"),
				resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "nginx_settings", "nginxSettings"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_pass", "https://www.oracle.com:443"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "server.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "cloudGateMappings", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Optional, acctest.Update, IdentityDomainsCloudGateMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "gateway_app.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "nginx_settings", "nginxSettings2"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_pass", "https://www.oracle.com:443"),
				resource.TestCheckResourceAttr(resourceName, "resource_prefix", "resourcePrefix2"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "server.#", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_cloud_gate_mappings", "test_cloud_gate_mappings", acctest.Optional, acctest.Update, IdentityDomainsCloudGateMappingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsCloudGateMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Optional, acctest.Update, IdentityDomainsCloudGateMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "cloud_gate_mapping_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_cloud_gate_mapping", "test_cloud_gate_mapping", acctest.Required, acctest.Create, IdentityDomainsCloudGateMappingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsCloudGateMappingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_gate_mapping_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gateway_app.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nginx_settings", "nginxSettings2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_pass", "https://www.oracle.com:443"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_prefix", "resourcePrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "server.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsCloudGateMappingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_cloud_gate_mapping", "cloudGateMappings"),
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

func testAccCheckIdentityDomainsCloudGateMappingDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_cloud_gate_mapping" {
			noResourceFound = false
			request := oci_identity_domains.GetCloudGateMappingRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.CloudGateMappingId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetCloudGateMapping(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsCloudGateMapping") {
		resource.AddTestSweepers("IdentityDomainsCloudGateMapping", &resource.Sweeper{
			Name:         "IdentityDomainsCloudGateMapping",
			Dependencies: acctest.DependencyGraph["cloudGateMapping"],
			F:            sweepIdentityDomainsCloudGateMappingResource,
		})
	}
}

func sweepIdentityDomainsCloudGateMappingResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	cloudGateMappingIds, err := getIdentityDomainsCloudGateMappingIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudGateMappingId := range cloudGateMappingIds {
		if ok := acctest.SweeperDefaultResourceId[cloudGateMappingId]; !ok {
			deleteCloudGateMappingRequest := oci_identity_domains.DeleteCloudGateMappingRequest{}

			deleteCloudGateMappingRequest.CloudGateMappingId = &cloudGateMappingId

			deleteCloudGateMappingRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteCloudGateMapping(context.Background(), deleteCloudGateMappingRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudGateMapping %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudGateMappingId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsCloudGateMappingIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudGateMappingId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listCloudGateMappingsRequest := oci_identity_domains.ListCloudGateMappingsRequest{}
	listCloudGateMappingsResponse, err := identityDomainsClient.ListCloudGateMappings(context.Background(), listCloudGateMappingsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudGateMapping list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudGateMapping := range listCloudGateMappingsResponse.Resources {
		id := *cloudGateMapping.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudGateMappingId", id)
	}
	return resourceIds, nil
}
