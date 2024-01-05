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
	IdentityDomainsCloudGateServerRequiredOnlyResource = IdentityDomainsCloudGateServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Required, acctest.Create, IdentityDomainsCloudGateServerRepresentation)

	IdentityDomainsCloudGateServerResourceConfig = IdentityDomainsCloudGateServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Optional, acctest.Update, IdentityDomainsCloudGateServerRepresentation)

	IdentityDomainsCloudGateServerSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_gate_server_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_cloud_gate_server.test_cloud_gate_server.id}`},
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":               acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"authorization":                acctest.Representation{RepType: acctest.Optional, Create: `authorization`, Update: `authorization2`},
		"resource_type_schema_version": acctest.Representation{RepType: acctest.Optional, Create: `resourceTypeSchemaVersion`, Update: `resourceTypeSchemaVersion2`},
	}

	IdentityDomainsCloudGateServerDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"cloud_gate_server_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"cloud_gate_server_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"authorization":            acctest.Representation{RepType: acctest.Optional, Create: `authorization`, Update: `authorization2`},
		"start_index":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	IdentityDomainsCloudGateServerRepresentation = map[string]interface{}{
		"cloud_gate":     acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsCloudGateServerCloudGateRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"host_name":      acctest.Representation{RepType: acctest.Required, Create: `hostName`, Update: `hostName2`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"port":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:CloudGateServer`}},
		"ssl":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"authorization":  acctest.Representation{RepType: acctest.Optional, Create: `authorization`, Update: `authorization2`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"nginx_settings": acctest.Representation{RepType: acctest.Optional, Create: `nginxSettings`, Update: `nginxSettings2`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsCloudGateServerTagsRepresentation},
	}
	IdentityDomainsCloudGateServerCloudGateRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_cloud_gate.test_cloud_gate.id}`},
	}
	IdentityDomainsCloudGateServerIdcsCreatedByRepresentation = map[string]interface{}{
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"ocid":    acctest.Representation{RepType: acctest.Optional, Create: `ocid`, Update: `ocid2`},
		"ref":     acctest.Representation{RepType: acctest.Optional, Create: `ref`, Update: `ref2`},
		"type":    acctest.Representation{RepType: acctest.Optional, Create: `User`, Update: `App`},
	}
	IdentityDomainsCloudGateServerIdcsLastModifiedByRepresentation = map[string]interface{}{
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"display": acctest.Representation{RepType: acctest.Optional, Create: `display`, Update: `display2`},
		"ocid":    acctest.Representation{RepType: acctest.Optional, Create: `ocid`, Update: `ocid2`},
		"ref":     acctest.Representation{RepType: acctest.Optional, Create: `ref`, Update: `ref2`},
		"type":    acctest.Representation{RepType: acctest.Optional, Create: `User`, Update: `App`},
	}
	IdentityDomainsCloudGateServerMetaRepresentation = map[string]interface{}{
		"created":       acctest.Representation{RepType: acctest.Optional, Create: `created`, Update: `created2`},
		"last_modified": acctest.Representation{RepType: acctest.Optional, Create: `lastModified`, Update: `lastModified2`},
		"location":      acctest.Representation{RepType: acctest.Optional, Create: `location`, Update: `location2`},
		"resource_type": acctest.Representation{RepType: acctest.Optional, Create: `resourceType`, Update: `resourceType2`},
		"version":       acctest.Representation{RepType: acctest.Optional, Create: `version`, Update: `version2`},
	}
	IdentityDomainsCloudGateServerTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsCloudGateServerResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate", "test_cloud_gate", acctest.Required, acctest.Create, IdentityDomainsCloudGateRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsCloudGateServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsCloudGateServerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_cloud_gate_server.test_cloud_gate_server"
	datasourceName := "data.oci_identity_domains_cloud_gate_servers.test_cloud_gate_servers"
	singularDatasourceName := "data.oci_identity_domains_cloud_gate_server.test_cloud_gate_server"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsCloudGateServerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Optional, acctest.Create, IdentityDomainsCloudGateServerRepresentation), "identitydomains", "cloudGateServer", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsCloudGateServerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Required, acctest.Create, IdentityDomainsCloudGateServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "host_name", "hostName"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateServerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Optional, acctest.Create, IdentityDomainsCloudGateServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authorization", "authorization"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "host_name", "hostName"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.0.type", "User"),
				resource.TestCheckResourceAttr(resourceName, "is_opc_service", "false"),
				resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "nginx_settings", "nginxSettings"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl", "false"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "cloudGateServers", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsCloudGateServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Optional, acctest.Update, IdentityDomainsCloudGateServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authorization", "authorization2"),
				resource.TestCheckResourceAttr(resourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "host_name", "hostName2"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_opc_service", "false"),
				resource.TestCheckResourceAttr(resourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "nginx_settings", "nginxSettings2"),
				resource.TestCheckResourceAttr(resourceName, "port", "11"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ssl", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_cloud_gate_servers", "test_cloud_gate_servers", acctest.Optional, acctest.Update, IdentityDomainsCloudGateServerDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsCloudGateServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Optional, acctest.Update, IdentityDomainsCloudGateServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "cloud_gate_server_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "authorization", "authorization2"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "10"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_cloud_gate_server", "test_cloud_gate_server", acctest.Required, acctest.Create, IdentityDomainsCloudGateServerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsCloudGateServerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_gate_server_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cloud_gate.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_name", "hostName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "meta.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nginx_settings", "nginxSettings2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ssl", "true"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsCloudGateServerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_cloud_gate_server", "cloudGateServers"),
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

func testAccCheckIdentityDomainsCloudGateServerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_cloud_gate_server" {
			noResourceFound = false
			request := oci_identity_domains.GetCloudGateServerRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			tmp := rs.Primary.ID
			request.CloudGateServerId = &tmp

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetCloudGateServer(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsCloudGateServer") {
		resource.AddTestSweepers("IdentityDomainsCloudGateServer", &resource.Sweeper{
			Name:         "IdentityDomainsCloudGateServer",
			Dependencies: acctest.DependencyGraph["cloudGateServer"],
			F:            sweepIdentityDomainsCloudGateServerResource,
		})
	}
}

func sweepIdentityDomainsCloudGateServerResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	cloudGateServerIds, err := getIdentityDomainsCloudGateServerIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudGateServerId := range cloudGateServerIds {
		if ok := acctest.SweeperDefaultResourceId[cloudGateServerId]; !ok {
			deleteCloudGateServerRequest := oci_identity_domains.DeleteCloudGateServerRequest{}

			deleteCloudGateServerRequest.CloudGateServerId = &cloudGateServerId

			deleteCloudGateServerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteCloudGateServer(context.Background(), deleteCloudGateServerRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudGateServer %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudGateServerId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsCloudGateServerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudGateServerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listCloudGateServersRequest := oci_identity_domains.ListCloudGateServersRequest{}
	listCloudGateServersResponse, err := identityDomainsClient.ListCloudGateServers(context.Background(), listCloudGateServersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudGateServer list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudGateServer := range listCloudGateServersResponse.Resources {
		id := *cloudGateServer.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudGateServerId", id)
	}
	return resourceIds, nil
}
