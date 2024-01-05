// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
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
	IdentityDomainsAuthTokenRequiredOnlyResource = IdentityDomainsAuthTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Required, acctest.Create, IdentityDomainsAuthTokenRepresentation)

	IdentityDomainsAuthTokenResourceConfig = IdentityDomainsAuthTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Optional, acctest.Update, IdentityDomainsAuthTokenRepresentation)

	IdentityDomainsIdentityDomainsAuthTokenSingularDataSourceRepresentation = map[string]interface{}{
		"auth_token_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_auth_token.test_auth_token.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsAuthTokenDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"auth_token_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"auth_token_filter": acctest.Representation{RepType: acctest.Optional, Create: `user.value eq \"${oci_identity_domains_user.test_user.id}\"`},
		"attribute_sets":    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":       acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsAuthTokenRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:authToken`}},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":     acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthTokenTagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionself_change_user": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsAuthTokenUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation},
		"user":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsAuthTokenUserRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsAuthToken},
	}

	ignoreChangeForIdentityDomainsAuthToken = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			// properties that are `returned:never`
			`status`,
			`urnietfparamsscimschemasoracleidcsextensionself_change_user`,
		}},
	}
	IdentityDomainsAuthTokenTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	IdentityDomainsAuthTokenUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation = map[string]interface{}{
		"allow_self_change": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	IdentityDomainsAuthTokenUserRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}

	IdentityDomainsAuthTokenResourceDependencies = TestDomainDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsAuthTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsAuthTokenResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_auth_token.test_auth_token"
	datasourceName := "data.oci_identity_domains_auth_tokens.test_auth_tokens"
	singularDatasourceName := "data.oci_identity_domains_auth_token.test_auth_token"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsAuthTokenResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Optional, acctest.Create, IdentityDomainsAuthTokenRepresentation), "identitydomains", "authToken", t)

	print(config + compartmentIdVariableStr + IdentityDomainsAuthTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Optional, acctest.Create, IdentityDomainsAuthTokenRepresentation))

	acctest.ResourceTest(t, testAccCheckIdentityDomainsAuthTokenDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Required, acctest.Create, IdentityDomainsAuthTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthTokenResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsAuthTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Optional, acctest.Create, IdentityDomainsAuthTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "user.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user.0.value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "authTokens", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_auth_tokens", "test_auth_tokens", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsAuthTokenDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAuthTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Optional, acctest.Update, IdentityDomainsAuthTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "auth_token_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "auth_tokens.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auth_tokens.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_auth_token", "test_auth_token", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAuthTokenSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAuthTokenResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auth_token_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user.0.value"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsAuthTokenRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_auth_token", "authTokens"),
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

func testAccCheckIdentityDomainsAuthTokenDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_auth_token" {
			noResourceFound = false
			request := oci_identity_domains.GetAuthTokenRequest{}

			tmp := rs.Primary.ID
			request.AuthTokenId = &tmp

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetAuthToken(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsAuthToken") {
		resource.AddTestSweepers("IdentityDomainsAuthToken", &resource.Sweeper{
			Name:         "IdentityDomainsAuthToken",
			Dependencies: acctest.DependencyGraph["authToken"],
			F:            sweepIdentityDomainsAuthTokenResource,
		})
	}
}

func sweepIdentityDomainsAuthTokenResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	authTokenIds, err := getIdentityDomainsAuthTokenIds(compartment)
	if err != nil {
		return err
	}
	for _, authTokenId := range authTokenIds {
		if ok := acctest.SweeperDefaultResourceId[authTokenId]; !ok {
			deleteAuthTokenRequest := oci_identity_domains.DeleteAuthTokenRequest{}

			deleteAuthTokenRequest.AuthTokenId = &authTokenId

			deleteAuthTokenRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteAuthToken(context.Background(), deleteAuthTokenRequest)
			if error != nil {
				fmt.Printf("Error deleting AuthToken %s %s, It is possible that the resource is already deleted. Please verify manually \n", authTokenId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsAuthTokenIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AuthTokenId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listAuthTokensRequest := oci_identity_domains.ListAuthTokensRequest{}
	listAuthTokensResponse, err := identityDomainsClient.ListAuthTokens(context.Background(), listAuthTokensRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AuthToken list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, authToken := range listAuthTokensResponse.Resources {
		id := *authToken.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AuthTokenId", id)
	}
	return resourceIds, nil
}
