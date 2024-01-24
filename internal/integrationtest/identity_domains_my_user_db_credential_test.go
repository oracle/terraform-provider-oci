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
	IdentityDomainsMyUserDbCredentialRequiredOnlyResource = IdentityDomainsMyUserDbCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Required, acctest.Create, IdentityDomainsMyUserDbCredentialRepresentation)

	IdentityDomainsMyUserDbCredentialResourceConfig = IdentityDomainsMyUserDbCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Optional, acctest.Update, IdentityDomainsMyUserDbCredentialRepresentation)

	IdentityDomainsIdentityDomainsMyUserDbCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_user_db_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_user_db_credential.test_my_user_db_credential.id}`},
	}

	IdentityDomainsIdentityDomainsMyUserDbCredentialDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_user_db_credential_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_user_db_credential_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":                  acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyUserDbCredentialRepresentation = map[string]interface{}{
		"db_password":   acctest.Representation{RepType: acctest.Required, Create: `dbPassword123456`},
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":       acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:UserDbCredentials`}},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":    acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMyUserDbCredentialTagsRepresentation},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsMyUserDbCredential},
	}
	ignoreChangeForIdentityDomainsMyUserDbCredential = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			// properties that are `returned:never`
			`status`,
			`tags`, // my_* resource will not return non-default attributes
			`db_password`,
		}},
	}
	IdentityDomainsMyUserDbCredentialTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMyUserDbCredentialResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyUserDbCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyUserDbCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_user_db_credential.test_my_user_db_credential"
	datasourceName := "data.oci_identity_domains_my_user_db_credentials.test_my_user_db_credentials"
	singularDatasourceName := "data.oci_identity_domains_my_user_db_credential.test_my_user_db_credential"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMyUserDbCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Optional, acctest.Create, IdentityDomainsMyUserDbCredentialRepresentation), "identitydomains", "myUserDbCredential", t)

	print(config + compartmentIdVariableStr + IdentityDomainsMyUserDbCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Optional, acctest.Create, IdentityDomainsMyUserDbCredentialRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsMyUserDbCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyUserDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Required, acctest.Create, IdentityDomainsMyUserDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_password"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyUserDbCredentialResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyUserDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Optional, acctest.Create, IdentityDomainsMyUserDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_password"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "myUserDbCredentials", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_user_db_credentials", "test_my_user_db_credentials", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyUserDbCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyUserDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Optional, acctest.Update, IdentityDomainsMyUserDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_user_db_credential_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_user_db_credentials.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_user_db_credentials.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_user_db_credential", "test_my_user_db_credential", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyUserDbCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyUserDbCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_user_db_credential_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_password"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMyUserDbCredentialRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_user_db_credential", "myUserDbCredentials"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMyUserDbCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_user_db_credential" {
			noResourceFound = false
			request := oci_identity_domains.GetMyUserDbCredentialRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MyUserDbCredentialId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMyUserDbCredential(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMyUserDbCredential") {
		resource.AddTestSweepers("IdentityDomainsMyUserDbCredential", &resource.Sweeper{
			Name:         "IdentityDomainsMyUserDbCredential",
			Dependencies: acctest.DependencyGraph["myUserDbCredential"],
			F:            sweepIdentityDomainsMyUserDbCredentialResource,
		})
	}
}

func sweepIdentityDomainsMyUserDbCredentialResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myUserDbCredentialIds, err := getIdentityDomainsMyUserDbCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, myUserDbCredentialId := range myUserDbCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[myUserDbCredentialId]; !ok {
			deleteMyUserDbCredentialRequest := oci_identity_domains.DeleteMyUserDbCredentialRequest{}

			deleteMyUserDbCredentialRequest.MyUserDbCredentialId = &myUserDbCredentialId

			deleteMyUserDbCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyUserDbCredential(context.Background(), deleteMyUserDbCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting MyUserDbCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", myUserDbCredentialId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyUserDbCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyUserDbCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyUserDbCredentialsRequest := oci_identity_domains.ListMyUserDbCredentialsRequest{}
	listMyUserDbCredentialsResponse, err := identityDomainsClient.ListMyUserDbCredentials(context.Background(), listMyUserDbCredentialsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyUserDbCredential list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myUserDbCredential := range listMyUserDbCredentialsResponse.Resources {
		id := *myUserDbCredential.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyUserDbCredentialId", id)
	}
	return resourceIds, nil
}
