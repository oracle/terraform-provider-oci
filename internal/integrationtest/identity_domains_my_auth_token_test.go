// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

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
	IdentityDomainsMyAuthTokenRequiredOnlyResource = IdentityDomainsMyAuthTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Required, acctest.Create, IdentityDomainsMyAuthTokenRepresentation)

	IdentityDomainsMyAuthTokenResourceConfig = IdentityDomainsMyAuthTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Optional, acctest.Update, IdentityDomainsMyAuthTokenRepresentation)

	IdentityDomainsIdentityDomainsMyAuthTokenSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_auth_token_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_auth_token.test_my_auth_token.id}`},
	}

	IdentityDomainsIdentityDomainsMyAuthTokenDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_auth_token_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_auth_token_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":          acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyAuthTokenRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":       acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:authToken`}},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":    acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMyAuthTokenTagsRepresentation},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMyAuthToken},
	}

	ignoreChangeForIdentityDomainsMyAuthToken = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			`tags`, // my_* resource will not return non-default attributes
			`status`,
		}},
	}
	IdentityDomainsMyAuthTokenTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMyAuthTokenResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyAuthTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyAuthTokenResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_auth_token.test_my_auth_token"
	datasourceName := "data.oci_identity_domains_my_auth_tokens.test_my_auth_tokens"
	singularDatasourceName := "data.oci_identity_domains_my_auth_token.test_my_auth_token"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMyAuthTokenResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Optional, acctest.Create, IdentityDomainsMyAuthTokenRepresentation), "identitydomains", "myAuthToken", t)

	print(config + compartmentIdVariableStr + IdentityDomainsMyAuthTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Optional, acctest.Create, IdentityDomainsMyAuthTokenRepresentation))

	acctest.ResourceTest(t, testAccCheckIdentityDomainsMyAuthTokenDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyAuthTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Required, acctest.Create, IdentityDomainsMyAuthTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyAuthTokenResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyAuthTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Optional, acctest.Create, IdentityDomainsMyAuthTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "myAuthTokens", resId)
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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_auth_tokens", "test_my_auth_tokens", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyAuthTokenDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyAuthTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Optional, acctest.Update, IdentityDomainsMyAuthTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_auth_token_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_auth_tokens.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_auth_tokens.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_auth_token", "test_my_auth_token", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyAuthTokenSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyAuthTokenResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_auth_token_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMyAuthTokenRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_auth_token", "myAuthTokens"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMyAuthTokenDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_auth_token" {
			noResourceFound = false
			request := oci_identity_domains.GetMyAuthTokenRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MyAuthTokenId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMyAuthToken(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMyAuthToken") {
		resource.AddTestSweepers("IdentityDomainsMyAuthToken", &resource.Sweeper{
			Name:         "IdentityDomainsMyAuthToken",
			Dependencies: acctest.DependencyGraph["myAuthToken"],
			F:            sweepIdentityDomainsMyAuthTokenResource,
		})
	}
}

func sweepIdentityDomainsMyAuthTokenResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myAuthTokenIds, err := getIdentityDomainsMyAuthTokenIds(compartment)
	if err != nil {
		return err
	}
	for _, myAuthTokenId := range myAuthTokenIds {
		if ok := acctest.SweeperDefaultResourceId[myAuthTokenId]; !ok {
			deleteMyAuthTokenRequest := oci_identity_domains.DeleteMyAuthTokenRequest{}

			deleteMyAuthTokenRequest.MyAuthTokenId = &myAuthTokenId

			deleteMyAuthTokenRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyAuthToken(context.Background(), deleteMyAuthTokenRequest)
			if error != nil {
				fmt.Printf("Error deleting MyAuthToken %s %s, It is possible that the resource is already deleted. Please verify manually \n", myAuthTokenId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyAuthTokenIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyAuthTokenId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyAuthTokensRequest := oci_identity_domains.ListMyAuthTokensRequest{}
	listMyAuthTokensResponse, err := identityDomainsClient.ListMyAuthTokens(context.Background(), listMyAuthTokensRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyAuthToken list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myAuthToken := range listMyAuthTokensResponse.Resources {
		id := *myAuthToken.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyAuthTokenId", id)
	}
	return resourceIds, nil
}
