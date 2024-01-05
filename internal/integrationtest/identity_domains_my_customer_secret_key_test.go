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
	IdentityDomainsMyCustomerSecretKeyRequiredOnlyResource = IdentityDomainsMyCustomerSecretKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Required, acctest.Create, IdentityDomainsMyCustomerSecretKeyRepresentation)

	IdentityDomainsMyCustomerSecretKeyResourceConfig = IdentityDomainsMyCustomerSecretKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Optional, acctest.Update, IdentityDomainsMyCustomerSecretKeyRepresentation)

	IdentityDomainsIdentityDomainsMyCustomerSecretKeySingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_customer_secret_key_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_customer_secret_key.test_my_customer_secret_key.id}`},
	}

	IdentityDomainsIdentityDomainsMyCustomerSecretKeyDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_customer_secret_key_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_customer_secret_key_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":                   acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyCustomerSecretKeyRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":       acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:customerSecretKey`}},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"expires_on":    acctest.Representation{RepType: acctest.Optional, Create: `2030-01-01T00:00:00Z`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMyCustomerSecretKeyTagsRepresentation},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMyCustomerSecretKey},
	}
	ignoreChangeForIdentityDomainsMyCustomerSecretKey = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			// properties that are `returned:never`
			`status`,
			`tags`, // my_* resource will not return non-default attributes
		}},
	}
	IdentityDomainsMyCustomerSecretKeyTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMyCustomerSecretKeyResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyCustomerSecretKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyCustomerSecretKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_customer_secret_key.test_my_customer_secret_key"
	datasourceName := "data.oci_identity_domains_my_customer_secret_keys.test_my_customer_secret_keys"
	singularDatasourceName := "data.oci_identity_domains_my_customer_secret_key.test_my_customer_secret_key"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMyCustomerSecretKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Optional, acctest.Create, IdentityDomainsMyCustomerSecretKeyRepresentation), "identitydomains", "myCustomerSecretKey", t)

	print(config + compartmentIdVariableStr + IdentityDomainsMyCustomerSecretKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Optional, acctest.Create, IdentityDomainsMyCustomerSecretKeyRepresentation))
	acctest.ResourceTest(t, testAccCheckIdentityDomainsMyCustomerSecretKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyCustomerSecretKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Required, acctest.Create, IdentityDomainsMyCustomerSecretKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyCustomerSecretKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyCustomerSecretKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Optional, acctest.Create, IdentityDomainsMyCustomerSecretKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "myCustomerSecretKeys", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_customer_secret_keys", "test_my_customer_secret_keys", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyCustomerSecretKeyDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyCustomerSecretKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Optional, acctest.Update, IdentityDomainsMyCustomerSecretKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_customer_secret_key_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_customer_secret_keys.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_customer_secret_keys.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_customer_secret_key", "test_my_customer_secret_key", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyCustomerSecretKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyCustomerSecretKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_customer_secret_key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expires_on", "2030-01-01T00:00:00Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMyCustomerSecretKeyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_customer_secret_key", "myCustomerSecretKeys"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMyCustomerSecretKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_customer_secret_key" {
			noResourceFound = false
			request := oci_identity_domains.GetMyCustomerSecretKeyRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MyCustomerSecretKeyId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMyCustomerSecretKey(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMyCustomerSecretKey") {
		resource.AddTestSweepers("IdentityDomainsMyCustomerSecretKey", &resource.Sweeper{
			Name:         "IdentityDomainsMyCustomerSecretKey",
			Dependencies: acctest.DependencyGraph["myCustomerSecretKey"],
			F:            sweepIdentityDomainsMyCustomerSecretKeyResource,
		})
	}
}

func sweepIdentityDomainsMyCustomerSecretKeyResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myCustomerSecretKeyIds, err := getIdentityDomainsMyCustomerSecretKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, myCustomerSecretKeyId := range myCustomerSecretKeyIds {
		if ok := acctest.SweeperDefaultResourceId[myCustomerSecretKeyId]; !ok {
			deleteMyCustomerSecretKeyRequest := oci_identity_domains.DeleteMyCustomerSecretKeyRequest{}

			deleteMyCustomerSecretKeyRequest.MyCustomerSecretKeyId = &myCustomerSecretKeyId

			deleteMyCustomerSecretKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyCustomerSecretKey(context.Background(), deleteMyCustomerSecretKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting MyCustomerSecretKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", myCustomerSecretKeyId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyCustomerSecretKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyCustomerSecretKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyCustomerSecretKeysRequest := oci_identity_domains.ListMyCustomerSecretKeysRequest{}
	listMyCustomerSecretKeysResponse, err := identityDomainsClient.ListMyCustomerSecretKeys(context.Background(), listMyCustomerSecretKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyCustomerSecretKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myCustomerSecretKey := range listMyCustomerSecretKeysResponse.Resources {
		id := *myCustomerSecretKey.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyCustomerSecretKeyId", id)
	}
	return resourceIds, nil
}
