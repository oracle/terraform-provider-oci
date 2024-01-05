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
	IdentityDomainsApiKeyRequiredOnlyResource = IdentityDomainsApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Required, acctest.Create, IdentityDomainsApiKeyRepresentation)

	IdentityDomainsApiKeyResourceConfig = IdentityDomainsApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Optional, acctest.Update, IdentityDomainsApiKeyRepresentation)

	IdentityDomainsIdentityDomainsApiKeySingularDataSourceRepresentation = map[string]interface{}{
		"api_key_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_api_key.test_api_key.id}`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsApiKeyDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"api_key_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"api_key_filter": acctest.Representation{RepType: acctest.Optional, Create: `user.value eq \"${oci_identity_domains_user.test_user.id}\"`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsApiKeyRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"key":            acctest.Representation{RepType: acctest.Required, Create: "-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\\nGwIDAQAB\\n-----END PUBLIC KEY-----"},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:apikey`}},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsApiKeyTagsRepresentation},
		"urnietfparamsscimschemasoracleidcsextensionself_change_user": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsApiKeyUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation},
		"user":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsApiKeyUserRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsApiKey},
	}

	ignoreChangeForIdentityDomainsApiKey = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			// properties that are `returned:never`
			`urnietfparamsscimschemasoracleidcsextensionself_change_user`,
		}},
	}
	IdentityDomainsApiKeyTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	IdentityDomainsApiKeyUrnietfparamsscimschemasoracleidcsextensionselfChangeUserRepresentation = map[string]interface{}{
		"allow_self_change": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	IdentityDomainsApiKeyUserRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}

	IdentityDomainsApiKeyResourceDependencies = TestDomainDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsApiKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsApiKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_api_key.test_api_key"
	datasourceName := "data.oci_identity_domains_api_keys.test_api_keys"
	singularDatasourceName := "data.oci_identity_domains_api_key.test_api_key"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsApiKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Optional, acctest.Create, IdentityDomainsApiKeyRepresentation), "identitydomains", "apiKey", t)

	print(config + compartmentIdVariableStr + IdentityDomainsApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Optional, acctest.Create, IdentityDomainsApiKeyRepresentation))

	acctest.ResourceTest(t, testAccCheckIdentityDomainsApiKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Required, acctest.Create, IdentityDomainsApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "key", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\nGwIDAQAB\n-----END PUBLIC KEY-----"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApiKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Optional, acctest.Create, IdentityDomainsApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "key", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\nGwIDAQAB\n-----END PUBLIC KEY-----"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "apiKeys", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_api_keys", "test_api_keys", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsApiKeyDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Optional, acctest.Update, IdentityDomainsApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "api_key_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "api_keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "api_keys.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_api_key", "test_api_key", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsApiKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApiKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\nGwIDAQAB\n-----END PUBLIC KEY-----"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user.0.value"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsApiKeyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_api_key", "apiKeys"),
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

func testAccCheckIdentityDomainsApiKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_api_key" {
			noResourceFound = false
			request := oci_identity_domains.GetApiKeyRequest{}

			tmp := rs.Primary.ID
			request.ApiKeyId = &tmp

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

			_, err := client.GetApiKey(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsApiKey") {
		resource.AddTestSweepers("IdentityDomainsApiKey", &resource.Sweeper{
			Name:         "IdentityDomainsApiKey",
			Dependencies: acctest.DependencyGraph["apiKey"],
			F:            sweepIdentityDomainsApiKeyResource,
		})
	}
}

func sweepIdentityDomainsApiKeyResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	apiKeyIds, err := getIdentityDomainsApiKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, apiKeyId := range apiKeyIds {
		if ok := acctest.SweeperDefaultResourceId[apiKeyId]; !ok {
			deleteApiKeyRequest := oci_identity_domains.DeleteApiKeyRequest{}

			deleteApiKeyRequest.ApiKeyId = &apiKeyId

			deleteApiKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteApiKey(context.Background(), deleteApiKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting ApiKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", apiKeyId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsApiKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApiKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listApiKeysRequest := oci_identity_domains.ListApiKeysRequest{}
	listApiKeysResponse, err := identityDomainsClient.ListApiKeys(context.Background(), listApiKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ApiKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, apiKey := range listApiKeysResponse.Resources {
		id := *apiKey.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApiKeyId", id)
	}
	return resourceIds, nil
}
