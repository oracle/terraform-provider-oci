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
	IdentityDomainsUserDbCredentialRequiredOnlyResource = IdentityDomainsUserDbCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Required, acctest.Create, IdentityDomainsUserDbCredentialRepresentation)

	IdentityDomainsUserDbCredentialResourceConfig = IdentityDomainsUserDbCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Optional, acctest.Update, IdentityDomainsUserDbCredentialRepresentation)

	IdentityDomainsIdentityDomainsUserDbCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"user_db_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user_db_credential.test_user_db_credential.id}`},
		"attribute_sets":        acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsUserDbCredentialDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"user_db_credential_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"user_db_credential_filter": acctest.Representation{RepType: acctest.Optional, Create: `user.value eq \"${oci_identity_domains_user.test_user.id}\"`},
		"attribute_sets":            acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":               acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsUserDbCredentialRepresentation = map[string]interface{}{
		"db_password":    acctest.Representation{RepType: acctest.Required, Create: `dbPassword123456`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:UserDbCredentials`}},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":     acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserDbCredentialTagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionself_change_user": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsUserDbCredentialUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation},
		"user":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsUserDbCredentialUserRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForIdentityDomainsUserDbCredential},
	}

	ignoreChangeForIdentityDomainsUserDbCredential = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			// properties that are `returned:never`
			`status`,
			`urnietfparamsscimschemasoracleidcsextensionself_change_user`,
			`db_password`,
		}},
	}
	IdentityDomainsUserDbCredentialTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	IdentityDomainsUserDbCredentialUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation = map[string]interface{}{
		"allow_self_change": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	IdentityDomainsUserDbCredentialUserRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}

	IdentityDomainsUserDbCredentialResourceDependencies = TestDomainDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(IdentityDomainsUserRepresentation, map[string]interface{}{
				"urnietfparamsscimschemasoracleidcsextensiondb_credentials_user": acctest.RepresentationGroup{RepType: acctest.Required, Group: TestDbUserName},
				"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangeForTestDbUserNameUser},
			}))

	TestDbUserName                    = map[string]interface{}{"db_user_name": acctest.Representation{RepType: acctest.Required, Create: `dbUserName`}}
	ignoreChangeForTestDbUserNameUser = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`emails`, `schemas`, `urnietfparamsscimschemasoracleidcsextensiondb_credentials_user`}},
	}
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsUserDbCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsUserDbCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_user_db_credential.test_user_db_credential"
	datasourceName := "data.oci_identity_domains_user_db_credentials.test_user_db_credentials"
	singularDatasourceName := "data.oci_identity_domains_user_db_credential.test_user_db_credential"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsUserDbCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Optional, acctest.Create, IdentityDomainsUserDbCredentialRepresentation), "identitydomains", "userDbCredential", t)

	print(config + compartmentIdVariableStr + IdentityDomainsUserDbCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Optional, acctest.Create, IdentityDomainsUserDbCredentialRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsUserDbCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Required, acctest.Create, IdentityDomainsUserDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_password"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserDbCredentialResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsUserDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Optional, acctest.Create, IdentityDomainsUserDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_password"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "userDbCredentials", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_user_db_credentials", "test_user_db_credentials", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsUserDbCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsUserDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Optional, acctest.Update, IdentityDomainsUserDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "user_db_credential_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "user_db_credentials.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "user_db_credentials.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_user_db_credential", "test_user_db_credential", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsUserDbCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsUserDbCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_db_credential_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_password"),
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
			Config:            config + IdentityDomainsUserDbCredentialRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_user_db_credential", "userDbCredentials"),
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

func testAccCheckIdentityDomainsUserDbCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_user_db_credential" {
			noResourceFound = false
			request := oci_identity_domains.GetUserDbCredentialRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			tmp := rs.Primary.ID
			request.UserDbCredentialId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetUserDbCredential(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsUserDbCredential") {
		resource.AddTestSweepers("IdentityDomainsUserDbCredential", &resource.Sweeper{
			Name:         "IdentityDomainsUserDbCredential",
			Dependencies: acctest.DependencyGraph["userDbCredential"],
			F:            sweepIdentityDomainsUserDbCredentialResource,
		})
	}
}

func sweepIdentityDomainsUserDbCredentialResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	userDbCredentialIds, err := getIdentityDomainsUserDbCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, userDbCredentialId := range userDbCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[userDbCredentialId]; !ok {
			deleteUserDbCredentialRequest := oci_identity_domains.DeleteUserDbCredentialRequest{}

			deleteUserDbCredentialRequest.UserDbCredentialId = &userDbCredentialId

			deleteUserDbCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteUserDbCredential(context.Background(), deleteUserDbCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting UserDbCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", userDbCredentialId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsUserDbCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UserDbCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listUserDbCredentialsRequest := oci_identity_domains.ListUserDbCredentialsRequest{}
	listUserDbCredentialsResponse, err := identityDomainsClient.ListUserDbCredentials(context.Background(), listUserDbCredentialsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UserDbCredential list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, userDbCredential := range listUserDbCredentialsResponse.Resources {
		id := *userDbCredential.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UserDbCredentialId", id)
	}
	return resourceIds, nil
}
