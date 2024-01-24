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
	IdentityDomainsMyApiKeyRequiredOnlyResource = IdentityDomainsMyApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Required, acctest.Create, IdentityDomainsMyApiKeyRepresentation)

	IdentityDomainsMyApiKeyResourceConfig = IdentityDomainsMyApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Optional, acctest.Update, IdentityDomainsMyApiKeyRepresentation)

	IdentityDomainsIdentityDomainsMyApiKeySingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_api_key_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_my_api_key.test_my_api_key.id}`},
	}

	IdentityDomainsIdentityDomainsMyApiKeyDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_api_key_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_api_key_filter": acctest.Representation{RepType: acctest.Optional, Create: `id eq \"${oci_identity_domains_my_api_key.test_my_api_key.id}\"`},
		"start_index":       acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyApiKeyRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"key":           acctest.Representation{RepType: acctest.Required, Create: "-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\\nGwIDAQAB\\n-----END PUBLIC KEY-----"},
		"schemas":       acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:apikey`}},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"tags":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsMyApiKeyTagsRepresentation},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangeForIdentityDomainsMyApiKey},
	}

	ignoreChangeForIdentityDomainsMyApiKey = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{
			`tags`, // my_* resource will not return non-default attributes
		}},
	}
	IdentityDomainsMyApiKeyTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsMyApiKeyResourceDependencies = TestDomainDependencies + envApiKeyFingerprintVariableStr

	envApiKeyFingerprint            = acctest.GetEnvSettingWithBlankDefaultVar("fingerprint")
	envApiKeyFingerprintVariableStr = fmt.Sprintf("variable \"env_api_key_fingerprint\" { default = \"%s\" }\n", envApiKeyFingerprint)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyApiKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyApiKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_my_api_key.test_my_api_key"
	datasourceName := "data.oci_identity_domains_my_api_keys.test_my_api_keys"
	singularDatasourceName := "data.oci_identity_domains_my_api_key.test_my_api_key"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsMyApiKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Optional, acctest.Create, IdentityDomainsMyApiKeyRepresentation), "identitydomains", "myApiKey", t)

	print(config + compartmentIdVariableStr + IdentityDomainsMyApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Optional, acctest.Create, IdentityDomainsMyApiKeyRepresentation))

	print(acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_api_keys", "test_my_api_keys", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyApiKeyDataSourceRepresentation))
	print(acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyApiKeySingularDataSourceRepresentation))

	acctest.ResourceTest(t, testAccCheckIdentityDomainsMyApiKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Required, acctest.Create, IdentityDomainsMyApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "key", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\nGwIDAQAB\n-----END PUBLIC KEY-----"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyApiKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Optional, acctest.Create, IdentityDomainsMyApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "key", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\nGwIDAQAB\n-----END PUBLIC KEY-----"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "myApiKeys", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_api_keys", "test_my_api_keys", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsMyApiKeyDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Optional, acctest.Update, IdentityDomainsMyApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_api_key_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_api_keys.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_api_keys.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_api_key", "test_my_api_key", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyApiKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyApiKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_api_key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqWmqkNQKbC1hWnmf3Uov\nH0BfH3WwJRC9Jd6/fo8oo/vhKAO5UsoiLup/7vMpHZq4KN5wXMkGtXlnyOn3KJA0\nzv4Dni4/AmoODy56i5la4wFXLOIAbVl3QnWiw9ALw9YKKx4KYoixoc3h4MPHdxdR\nmqb/B8Niq2OSS2eUsCcbAcC41eUIx2yySqA77Qy7lLjzBAE5QBZY0V1BduG/Xi9u\nzXf6gHwXEszrKpyBIq4t5+g1sbiQWwJPtuVura8iH2gDkRKS/kRzLszn3vX0lePC\nuCqFeTcOdOieuZndkCIzEMw0UUfdw2+qNOAtNi+mOPKl0h7sM21x3OA+fHmnn1fV\nGwIDAQAB\n-----END PUBLIC KEY-----"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsMyApiKeyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_my_api_key", "myApiKeys"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsMyApiKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_my_api_key" {
			noResourceFound = false
			request := oci_identity_domains.GetMyApiKeyRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.MyApiKeyId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetMyApiKey(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsMyApiKey") {
		resource.AddTestSweepers("IdentityDomainsMyApiKey", &resource.Sweeper{
			Name:         "IdentityDomainsMyApiKey",
			Dependencies: acctest.DependencyGraph["myApiKey"],
			F:            sweepIdentityDomainsMyApiKeyResource,
		})
	}
}

func sweepIdentityDomainsMyApiKeyResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myApiKeyIds, err := getIdentityDomainsMyApiKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, myApiKeyId := range myApiKeyIds {
		if ok := acctest.SweeperDefaultResourceId[myApiKeyId]; !ok {
			deleteMyApiKeyRequest := oci_identity_domains.DeleteMyApiKeyRequest{}

			deleteMyApiKeyRequest.MyApiKeyId = &myApiKeyId

			deleteMyApiKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyApiKey(context.Background(), deleteMyApiKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting MyApiKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", myApiKeyId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyApiKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyApiKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyApiKeysRequest := oci_identity_domains.ListMyApiKeysRequest{}
	listMyApiKeysResponse, err := identityDomainsClient.ListMyApiKeys(context.Background(), listMyApiKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyApiKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myApiKey := range listMyApiKeysResponse.Resources {
		id := *myApiKey.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyApiKeyId", id)
	}
	return resourceIds, nil
}
