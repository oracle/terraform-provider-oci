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
	IdentityDomainsOAuth2ClientCredentialRequiredOnlyResource = IdentityDomainsOAuth2ClientCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Required, acctest.Create, IdentityDomainsOAuth2ClientCredentialRepresentation)

	IdentityDomainsOAuth2ClientCredentialResourceConfig = IdentityDomainsOAuth2ClientCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Optional, acctest.Update, IdentityDomainsOAuth2ClientCredentialRepresentation)

	IdentityDomainsIdentityDomainsOAuth2ClientCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"o_auth2client_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_oauth2client_credential.test_oauth2client_credential.id}`},
		"attribute_sets":              acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsOAuth2ClientCredentialDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"oauth2client_credential_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"oauth2client_credential_filter": acctest.Representation{RepType: acctest.Optional, Create: `user.value eq \"${oci_identity_domains_user.test_user.id}\"`},
		"attribute_sets":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsOAuth2ClientCredentialRepresentation = map[string]interface{}{
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `name`},
		"schemas":         acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:oauth2ClientCredential`}},
		"scopes":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsOAuth2ClientCredentialScopesRepresentation},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":      acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"is_reset_secret": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"status":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsOAuth2ClientCredentialTagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionself_change_user": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsOAuth2ClientCredentialUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation},
		"user":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsOAuth2ClientCredentialUserRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsOAuth2ClientCredential},
	}

	ignoreChangeForIdentityDomainsOAuth2ClientCredential = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			// properties that are `returned:never`
			`status`,
			`urnietfparamsscimschemasoracleidcsextensionself_change_user`,
			// mutability: writeOnly
			`is_reset_secret`,
		}},
	}
	IdentityDomainsOAuth2ClientCredentialScopesRepresentation = map[string]interface{}{
		"audience": acctest.Representation{RepType: acctest.Required, Create: `audience`},
		"scope":    acctest.Representation{RepType: acctest.Required, Create: `scope`},
	}
	IdentityDomainsOAuth2ClientCredentialTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	IdentityDomainsOAuth2ClientCredentialUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation = map[string]interface{}{
		"allow_self_change": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	IdentityDomainsOAuth2ClientCredentialUserRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}

	IdentityDomainsOAuth2ClientCredentialResourceDependencies = TestDomainDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsOAuth2ClientCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsOAuth2ClientCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_oauth2client_credential.test_oauth2client_credential"
	datasourceName := "data.oci_identity_domains_oauth2client_credentials.test_oauth2client_credentials"
	singularDatasourceName := "data.oci_identity_domains_oauth2client_credential.test_oauth2client_credential"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsOAuth2ClientCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Optional, acctest.Create, IdentityDomainsOAuth2ClientCredentialRepresentation), "identitydomains", "oAuth2ClientCredential", t)

	print(config + compartmentIdVariableStr + IdentityDomainsOAuth2ClientCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Optional, acctest.Create, IdentityDomainsOAuth2ClientCredentialRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsOAuth2ClientCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsOAuth2ClientCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Required, acctest.Create, IdentityDomainsOAuth2ClientCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.scope", "scope"),
				resource.TestCheckResourceAttrSet(resourceName, "secret"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsOAuth2ClientCredentialResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsOAuth2ClientCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Optional, acctest.Create, IdentityDomainsOAuth2ClientCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.scope", "scope"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "user.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user.0.value"),
				resource.TestCheckResourceAttrSet(resourceName, "secret"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "oAuth2ClientCredentials", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oauth2client_credentials", "test_oauth2client_credentials", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsOAuth2ClientCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsOAuth2ClientCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Optional, acctest.Update, IdentityDomainsOAuth2ClientCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "oauth2client_credential_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "oauth2client_credentials.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oauth2client_credentials.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oauth2client_credential", "test_oauth2client_credential", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsOAuth2ClientCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsOAuth2ClientCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "o_auth2client_credential_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.audience", "audience"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.scope", "scope"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user.0.value"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsOAuth2ClientCredentialRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_oauth2client_credential", "oAuth2ClientCredentials"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
				"is_reset_secret",
				"secret",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsOAuth2ClientCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_oauth2client_credential" {
			noResourceFound = false
			request := oci_identity_domains.GetOAuth2ClientCredentialRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.OAuth2ClientCredentialId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetOAuth2ClientCredential(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsOAuth2ClientCredential") {
		resource.AddTestSweepers("IdentityDomainsOAuth2ClientCredential", &resource.Sweeper{
			Name:         "IdentityDomainsOAuth2ClientCredential",
			Dependencies: acctest.DependencyGraph["oAuth2ClientCredential"],
			F:            sweepIdentityDomainsOAuth2ClientCredentialResource,
		})
	}
}

func sweepIdentityDomainsOAuth2ClientCredentialResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	oAuth2ClientCredentialIds, err := getIdentityDomainsOAuth2ClientCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, oAuth2ClientCredentialId := range oAuth2ClientCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[oAuth2ClientCredentialId]; !ok {
			deleteOAuth2ClientCredentialRequest := oci_identity_domains.DeleteOAuth2ClientCredentialRequest{}

			deleteOAuth2ClientCredentialRequest.OAuth2ClientCredentialId = &oAuth2ClientCredentialId

			deleteOAuth2ClientCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteOAuth2ClientCredential(context.Background(), deleteOAuth2ClientCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting OAuth2ClientCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", oAuth2ClientCredentialId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsOAuth2ClientCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OAuth2ClientCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listOAuth2ClientCredentialsRequest := oci_identity_domains.ListOAuth2ClientCredentialsRequest{}
	listOAuth2ClientCredentialsResponse, err := identityDomainsClient.ListOAuth2ClientCredentials(context.Background(), listOAuth2ClientCredentialsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OAuth2ClientCredential list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oAuth2ClientCredential := range listOAuth2ClientCredentialsResponse.Resources {
		id := *oAuth2ClientCredential.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OAuth2ClientCredentialId", id)
	}
	return resourceIds, nil
}
