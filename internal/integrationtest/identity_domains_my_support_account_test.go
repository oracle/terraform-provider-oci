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
	IdentityDomainsMySupportAccountRequiredOnlyResource = IdentityDomainsMySupportAccountResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Required, acctest.Create, IdentityDomainsMySupportAccountRepresentation)

	IdentityDomainsMySupportAccountResourceConfig = IdentityDomainsMySupportAccountResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Optional, acctest.Update, IdentityDomainsMySupportAccountRepresentation)

	IdentityDomainsIdentityDomainsMySupportAccountSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_support_account_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_support_account.test_my_support_account.id}`},
	}

	IdentityDomainsIdentityDomainsMySupportAccountDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_support_account_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_support_account_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":               acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMySupportAccountRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":       acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:supportAccount`}},
		"token":         acctest.Representation{RepType: acctest.Required, Create: `${var.my_support_account_token}`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMySupportAccountTagsRepresentation},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMySupportAccount},
	}

	ignoreChangeForIdentityDomainsMySupportAccount = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			`tags`, // my_* resource will not return non-default attributes
		}},
	}
	IdentityDomainsMySupportAccountTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	/**
	Need to provide token in env.
	Go to tfadmin user My profile page, click "Link support account" and get the token, and remember to "Unlink" afterwards.
	It will expire after a while, needs to redo the above and update in env if so.
	*/
	mySupportAccountToken            = utils.GetEnvSettingWithBlankDefault("my_support_account_token")
	mySupportAccountTokenVariableStr = fmt.Sprintf("variable \"my_support_account_token\" { default = \"%s\" }\n", mySupportAccountToken)

	IdentityDomainsMySupportAccountResourceDependencies = TestDomainDependencies + mySupportAccountTokenVariableStr
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMySupportAccountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMySupportAccountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_support_account.test_my_support_account"
	datasourceName := "data.oci_identity_domains_my_support_accounts.test_my_support_accounts"
	singularDatasourceName := "data.oci_identity_domains_my_support_account.test_my_support_account"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMySupportAccountResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Optional, acctest.Create, IdentityDomainsMySupportAccountRepresentation), "identitydomains", "mySupportAccount", t)

	print(acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Required, acctest.Create, IdentityDomainsMySupportAccountRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsMySupportAccountDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMySupportAccountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Required, acctest.Create, IdentityDomainsMySupportAccountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "token"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMySupportAccountResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMySupportAccountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Optional, acctest.Create, IdentityDomainsMySupportAccountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "token"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "mySupportAccounts", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_support_accounts", "test_my_support_accounts", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMySupportAccountDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMySupportAccountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Optional, acctest.Update, IdentityDomainsMySupportAccountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_support_account_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_support_accounts.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_support_accounts.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_support_account", "test_my_support_account", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMySupportAccountSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMySupportAccountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_support_account_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMySupportAccountRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_support_account", "mySupportAccounts"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"token",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMySupportAccountDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_support_account" {
			noResourceFound = false
			request := oci_identity_domains.GetMySupportAccountRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MySupportAccountId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMySupportAccount(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMySupportAccount") {
		resource.AddTestSweepers("IdentityDomainsMySupportAccount", &resource.Sweeper{
			Name:         "IdentityDomainsMySupportAccount",
			Dependencies: acctest.DependencyGraph["mySupportAccount"],
			F:            sweepIdentityDomainsMySupportAccountResource,
		})
	}
}

func sweepIdentityDomainsMySupportAccountResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	mySupportAccountIds, err := getIdentityDomainsMySupportAccountIds(compartment)
	if err != nil {
		return err
	}
	for _, mySupportAccountId := range mySupportAccountIds {
		if ok := acctest.SweeperDefaultResourceId[mySupportAccountId]; !ok {
			deleteMySupportAccountRequest := oci_identity_domains.DeleteMySupportAccountRequest{}

			deleteMySupportAccountRequest.MySupportAccountId = &mySupportAccountId

			deleteMySupportAccountRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMySupportAccount(context.Background(), deleteMySupportAccountRequest)
			if error != nil {
				fmt.Printf("Error deleting MySupportAccount %s %s, It is possible that the resource is already deleted. Please verify manually \n", mySupportAccountId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMySupportAccountIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MySupportAccountId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMySupportAccountsRequest := oci_identity_domains.ListMySupportAccountsRequest{}
	listMySupportAccountsResponse, err := identityDomainsClient.ListMySupportAccounts(context.Background(), listMySupportAccountsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MySupportAccount list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mySupportAccount := range listMySupportAccountsResponse.Resources {
		id := *mySupportAccount.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MySupportAccountId", id)
	}
	return resourceIds, nil
}
