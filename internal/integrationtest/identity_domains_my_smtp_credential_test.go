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
	IdentityDomainsMySmtpCredentialRequiredOnlyResource = IdentityDomainsMySmtpCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Required, acctest.Create, IdentityDomainsMySmtpCredentialRepresentation)

	IdentityDomainsMySmtpCredentialResourceConfig = IdentityDomainsMySmtpCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Optional, acctest.Update, IdentityDomainsMySmtpCredentialRepresentation)

	IdentityDomainsIdentityDomainsMySmtpCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_smtp_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_smtp_credential.test_my_smtp_credential.id}`},
	}

	IdentityDomainsIdentityDomainsMySmtpCredentialDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_smtp_credential_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_smtp_credential_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":               acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMySmtpCredentialRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":       acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:smtpCredential`}},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"expires_on":    acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMySmtpCredentialTagsRepresentation},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMySmtpCredential},
	}
	ignoreChangeForIdentityDomainsMySmtpCredential = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			// properties that are `returned:never`
			`status`,
			`tags`, // my_* resource will not return non-default attributes
		}},
	}
	IdentityDomainsMySmtpCredentialTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMySmtpCredentialResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMySmtpCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMySmtpCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_smtp_credential.test_my_smtp_credential"
	datasourceName := "data.oci_identity_domains_my_smtp_credentials.test_my_smtp_credentials"
	singularDatasourceName := "data.oci_identity_domains_my_smtp_credential.test_my_smtp_credential"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMySmtpCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Optional, acctest.Create, IdentityDomainsMySmtpCredentialRepresentation), "identitydomains", "mySmtpCredential", t)

	print(config + compartmentIdVariableStr + IdentityDomainsMySmtpCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Optional, acctest.Create, IdentityDomainsMySmtpCredentialRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsMySmtpCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMySmtpCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Required, acctest.Create, IdentityDomainsMySmtpCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMySmtpCredentialResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMySmtpCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Optional, acctest.Create, IdentityDomainsMySmtpCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "mySmtpCredentials", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_smtp_credentials", "test_my_smtp_credentials", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMySmtpCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMySmtpCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Optional, acctest.Update, IdentityDomainsMySmtpCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_smtp_credential_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_smtp_credentials.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_smtp_credentials.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_smtp_credential", "test_my_smtp_credential", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMySmtpCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMySmtpCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_smtp_credential_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMySmtpCredentialRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_smtp_credential", "mySmtpCredentials"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMySmtpCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_smtp_credential" {
			noResourceFound = false
			request := oci_identity_domains.GetMySmtpCredentialRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MySmtpCredentialId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMySmtpCredential(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMySmtpCredential") {
		resource.AddTestSweepers("IdentityDomainsMySmtpCredential", &resource.Sweeper{
			Name:         "IdentityDomainsMySmtpCredential",
			Dependencies: acctest.DependencyGraph["mySmtpCredential"],
			F:            sweepIdentityDomainsMySmtpCredentialResource,
		})
	}
}

func sweepIdentityDomainsMySmtpCredentialResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	mySmtpCredentialIds, err := getIdentityDomainsMySmtpCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, mySmtpCredentialId := range mySmtpCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[mySmtpCredentialId]; !ok {
			deleteMySmtpCredentialRequest := oci_identity_domains.DeleteMySmtpCredentialRequest{}

			deleteMySmtpCredentialRequest.MySmtpCredentialId = &mySmtpCredentialId

			deleteMySmtpCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMySmtpCredential(context.Background(), deleteMySmtpCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting MySmtpCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", mySmtpCredentialId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMySmtpCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MySmtpCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMySmtpCredentialsRequest := oci_identity_domains.ListMySmtpCredentialsRequest{}
	listMySmtpCredentialsResponse, err := identityDomainsClient.ListMySmtpCredentials(context.Background(), listMySmtpCredentialsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MySmtpCredential list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mySmtpCredential := range listMySmtpCredentialsResponse.Resources {
		id := *mySmtpCredential.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MySmtpCredentialId", id)
	}
	return resourceIds, nil
}
