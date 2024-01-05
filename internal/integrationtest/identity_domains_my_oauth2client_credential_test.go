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
	IdentityDomainsMyOAuth2ClientCredentialRequiredOnlyResource = IdentityDomainsMyOAuth2ClientCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Required, acctest.Create, IdentityDomainsMyOAuth2ClientCredentialRepresentation)

	IdentityDomainsMyOAuth2ClientCredentialResourceConfig = IdentityDomainsMyOAuth2ClientCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Optional, acctest.Update, IdentityDomainsMyOAuth2ClientCredentialRepresentation)

	IdentityDomainsIdentityDomainsMyOAuth2ClientCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_oauth2client_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_oauth2client_credential.test_my_oauth2client_credential.id}`},
	}

	IdentityDomainsIdentityDomainsMyOAuth2ClientCredentialDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_oauth2client_credential_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_oauth2client_credential_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":                       acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyOAuth2ClientCredentialRepresentation = map[string]interface{}{
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `name`},
		"schemas":         acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:oauth2ClientCredential`}},
		"scopes":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsMyOAuth2ClientCredentialScopesRepresentation},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":      acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"is_reset_secret": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"status":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMyOAuth2ClientCredentialTagsRepresentation},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMyOAuth2ClientCredential},
	}
	ignoreChangeForIdentityDomainsMyOAuth2ClientCredential = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			// properties that are `returned:never`
			`status`,
			`tags`, // my_* resource will not return non-default attributes
		}},
	}
	IdentityDomainsMyOAuth2ClientCredentialScopesRepresentation = map[string]interface{}{
		"audience": acctest.Representation{RepType: acctest.Required, Create: `audience`},
		"scope":    acctest.Representation{RepType: acctest.Required, Create: `scope`},
	}
	IdentityDomainsMyOAuth2ClientCredentialTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMyOAuth2ClientCredentialResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyOAuth2ClientCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyOAuth2ClientCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_oauth2client_credential.test_my_oauth2client_credential"
	datasourceName := "data.oci_identity_domains_my_oauth2client_credentials.test_my_oauth2client_credentials"
	singularDatasourceName := "data.oci_identity_domains_my_oauth2client_credential.test_my_oauth2client_credential"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMyOAuth2ClientCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Optional, acctest.Create, IdentityDomainsMyOAuth2ClientCredentialRepresentation), "identitydomains", "myOAuth2ClientCredential", t)

	print(config + compartmentIdVariableStr + IdentityDomainsMyOAuth2ClientCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Optional, acctest.Create, IdentityDomainsMyOAuth2ClientCredentialRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsMyOAuth2ClientCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyOAuth2ClientCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Required, acctest.Create, IdentityDomainsMyOAuth2ClientCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.scope", "scope"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyOAuth2ClientCredentialResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyOAuth2ClientCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Optional, acctest.Create, IdentityDomainsMyOAuth2ClientCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "is_reset_secret", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.audience", "audience"),
				resource.TestCheckResourceAttr(resourceName, "scopes.0.scope", "scope"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "myOAuth2ClientCredentials", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credentials", "test_my_oauth2client_credentials", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyOAuth2ClientCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyOAuth2ClientCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Optional, acctest.Update, IdentityDomainsMyOAuth2ClientCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_oauth2client_credential_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_oauth2client_credentials.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_oauth2client_credentials.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_oauth2client_credential", "test_my_oauth2client_credential", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyOAuth2ClientCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyOAuth2ClientCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_oauth2client_credential_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.audience", "audience"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scopes.0.scope", "scope"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMyOAuth2ClientCredentialRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_oauth2client_credential", "myOAuth2ClientCredentials"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"is_reset_secret", //writeOnly
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMyOAuth2ClientCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_oauth2client_credential" {
			noResourceFound = false
			request := oci_identity_domains.GetMyOAuth2ClientCredentialRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MyOAuth2ClientCredentialId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMyOAuth2ClientCredential(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMyOAuth2ClientCredential") {
		resource.AddTestSweepers("IdentityDomainsMyOAuth2ClientCredential", &resource.Sweeper{
			Name:         "IdentityDomainsMyOAuth2ClientCredential",
			Dependencies: acctest.DependencyGraph["myOAuth2ClientCredential"],
			F:            sweepIdentityDomainsMyOAuth2ClientCredentialResource,
		})
	}
}

func sweepIdentityDomainsMyOAuth2ClientCredentialResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myOAuth2ClientCredentialIds, err := getIdentityDomainsMyOAuth2ClientCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, myOAuth2ClientCredentialId := range myOAuth2ClientCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[myOAuth2ClientCredentialId]; !ok {
			deleteMyOAuth2ClientCredentialRequest := oci_identity_domains.DeleteMyOAuth2ClientCredentialRequest{}

			deleteMyOAuth2ClientCredentialRequest.MyOAuth2ClientCredentialId = &myOAuth2ClientCredentialId

			deleteMyOAuth2ClientCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyOAuth2ClientCredential(context.Background(), deleteMyOAuth2ClientCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting MyOAuth2ClientCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", myOAuth2ClientCredentialId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyOAuth2ClientCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyOAuth2ClientCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyOAuth2ClientCredentialsRequest := oci_identity_domains.ListMyOAuth2ClientCredentialsRequest{}
	listMyOAuth2ClientCredentialsResponse, err := identityDomainsClient.ListMyOAuth2ClientCredentials(context.Background(), listMyOAuth2ClientCredentialsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyOAuth2ClientCredential list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myOAuth2ClientCredential := range listMyOAuth2ClientCredentialsResponse.Resources {
		id := *myOAuth2ClientCredential.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyOAuth2ClientCredentialId", id)
	}
	return resourceIds, nil
}
