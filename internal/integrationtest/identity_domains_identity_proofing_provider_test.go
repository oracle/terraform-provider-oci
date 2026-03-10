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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsIdentityProofingProviderRequiredOnlyResource = IdentityDomainsIdentityProofingProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderRepresentation)

	IdentityDomainsIdentityProofingProviderResourceConfig = IdentityDomainsIdentityProofingProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderRepresentation)

	IdentityDomainsIdentityProofingProviderSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_proofing_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_proofing_provider.test_identity_proofing_provider.id}`},
		"attribute_sets":                acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityProofingProviderDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_proofing_provider_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"identity_proofing_provider_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                       acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsIdentityProofingProviderRepresentation = map[string]interface{}{
		"identity_proofing_provider_provider": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_proofing_provider_template.test_identity_proofing_provider_template.identity_proofing_provider_template_provider}`, Update: `${oci_identity_domains_identity_proofing_provider_template.test_identity_proofing_provider_template_for_update.identity_proofing_provider_template_provider}`},
		"claim_mapping":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsIdentityProofingProviderClaimMappingRepresentation},
		"configuration":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsIdentityProofingProviderConfigurationRepresentation},
		"idcs_endpoint":                       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"schemas":                             acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:IdentityProofingProvider`}},
		"attribute_sets":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":                         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"runtime_data":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProofingProviderRuntimeDataRepresentation},
		"status":                              acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}
	IdentityDomainsIdentityProofingProviderClaimMappingRepresentation = map[string]interface{}{
		"attr_match":       acctest.Representation{RepType: acctest.Required, Create: `attrMatch`, Update: `attrMatch2`},
		"verifiable_claim": acctest.Representation{RepType: acctest.Required, Create: `verifiableClaim`, Update: `verifiableClaim2`},
	}
	IdentityDomainsIdentityProofingProviderConfigurationRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}
	IdentityDomainsIdentityProofingProviderRuntimeDataRepresentation = map[string]interface{}{
		"attr_name":  acctest.Representation{RepType: acctest.Required, Create: `attrName`, Update: `attrName2`},
		"attr_value": acctest.Representation{RepType: acctest.Required, Create: `attrValue`, Update: `attrValue2`},
	}

	IdentityDomainsIdentityProofingProviderResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderTemplateRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template_for_update", acctest.Required, acctest.Update, IdentityDomainsIdentityProofingProviderTemplateRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsIdentityProofingProviderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsIdentityProofingProviderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_identity_proofing_provider.test_identity_proofing_provider"
	datasourceName := "data.oci_identity_domains_identity_proofing_providers.test_identity_proofing_providers"
	singularDatasourceName := "data.oci_identity_domains_identity_proofing_provider.test_identity_proofing_provider"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsIdentityProofingProviderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Optional, acctest.Create, IdentityDomainsIdentityProofingProviderRepresentation), "identitydomains", "identityProofingProvider", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsIdentityProofingProviderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "identity_proofing_provider_provider"),
				resource.TestCheckResourceAttr(resourceName, "claim_mapping.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "claim_mapping", map[string]string{
					"attr_match":       "attrMatch",
					"verifiable_claim": "verifiableClaim",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "configuration", map[string]string{
					"name":  "name",
					"value": "value",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Optional, acctest.Create, IdentityDomainsIdentityProofingProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "identity_proofing_provider_provider"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "claim_mapping.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "claim_mapping", map[string]string{
					"attr_match":       "attrMatch",
					"verifiable_claim": "verifiableClaim",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "configuration", map[string]string{
					"name":  "name",
					"value": "value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "runtime_data.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "runtime_data", map[string]string{
					"attr_name":  "attrName",
					"attr_value": "attrValue",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "status", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "identityProofingProviders", resId)
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "identity_proofing_provider_provider"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "claim_mapping.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "claim_mapping", map[string]string{
					"attr_match":       "attrMatch2",
					"verifiable_claim": "verifiableClaim2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "configuration", map[string]string{
					"name":  "name2",
					"value": "value2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "runtime_data.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "runtime_data", map[string]string{
					"attr_name":  "attrName2",
					"attr_value": "attrValue2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "status", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_proofing_providers", "test_identity_proofing_providers", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "identity_proofing_provider_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "identity_proofing_providers.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "identity_proofing_providers.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider", "test_identity_proofing_provider", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_proofing_provider_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_proofing_provider_provider"),
				resource.TestCheckResourceAttr(singularDatasourceName, "claim_mapping.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "claim_mapping", map[string]string{
					"attr_match":       "attrMatch2",
					"verifiable_claim": "verifiableClaim2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "configuration", map[string]string{
					"name":  "name2",
					"value": "value2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runtime_data.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "runtime_data", map[string]string{
					"attr_name":  "attrName2",
					"attr_value": "attrValue2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "ACTIVE"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsIdentityProofingProviderRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_identity_proofing_provider", "identityProofingProviders"),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsIdentityProofingProviderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_identity_proofing_provider" {
			noResourceFound = false
			request := oci_identity_domains.GetIdentityProofingProviderRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.IdentityProofingProviderId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetIdentityProofingProvider(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsIdentityProofingProvider") {
		resource.AddTestSweepers("IdentityDomainsIdentityProofingProvider", &resource.Sweeper{
			Name:         "IdentityDomainsIdentityProofingProvider",
			Dependencies: acctest.DependencyGraph["identityProofingProvider"],
			F:            sweepIdentityDomainsIdentityProofingProviderResource,
		})
	}
}

func sweepIdentityDomainsIdentityProofingProviderResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	identityProofingProviderIds, err := getIdentityDomainsIdentityProofingProviderIds(compartment)
	if err != nil {
		return err
	}
	for _, identityProofingProviderId := range identityProofingProviderIds {
		if ok := acctest.SweeperDefaultResourceId[identityProofingProviderId]; !ok {
			deleteIdentityProofingProviderRequest := oci_identity_domains.DeleteIdentityProofingProviderRequest{}

			deleteIdentityProofingProviderRequest.IdentityProofingProviderId = &identityProofingProviderId

			deleteIdentityProofingProviderRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteIdentityProofingProvider(context.Background(), deleteIdentityProofingProviderRequest)
			if error != nil {
				fmt.Printf("Error deleting IdentityProofingProvider %s %s, It is possible that the resource is already deleted. Please verify manually \n", identityProofingProviderId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsIdentityProofingProviderIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IdentityProofingProviderId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listIdentityProofingProvidersRequest := oci_identity_domains.ListIdentityProofingProvidersRequest{}
	//listIdentityProofingProvidersRequest.CompartmentId = &compartmentId
	listIdentityProofingProvidersResponse, err := identityDomainsClient.ListIdentityProofingProviders(context.Background(), listIdentityProofingProvidersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IdentityProofingProvider list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, identityProofingProvider := range listIdentityProofingProvidersResponse.Resources {
		id := *identityProofingProvider.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IdentityProofingProviderId", id)
	}
	return resourceIds, nil
}
