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
	IdentityDomainsIdentityProofingProviderTemplateRequiredOnlyResource = IdentityDomainsIdentityProofingProviderTemplateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderTemplateRepresentation)

	IdentityDomainsIdentityProofingProviderTemplateResourceConfig = IdentityDomainsIdentityProofingProviderTemplateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderTemplateRepresentation)

	IdentityDomainsIdentityProofingProviderTemplateSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_proofing_provider_template_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_proofing_provider_template.test_identity_proofing_provider_template.id}`},
		"attribute_sets":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityProofingProviderTemplateDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_proofing_provider_template_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"identity_proofing_provider_template_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsIdentityProofingProviderTemplateRepresentation = map[string]interface{}{
		"identity_proofing_provider_template_provider": acctest.Representation{RepType: acctest.Required, Create: `identityProofingProviderTemplateProvider`, Update: `identityProofingProviderTemplateProvider2`},
		"idcs_endpoint":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":          acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:IdentityProofingProviderTemplate`}},
		"service_type":     acctest.Representation{RepType: acctest.Required, Create: []string{`serviceType`}, Update: []string{`serviceType2`}},
		"verification_url": acctest.Representation{RepType: acctest.Required, Create: `verificationUrl`, Update: `verificationUrl2`},
		"attribute_sets":   acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"configuration":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityProofingProviderTemplateConfigurationRepresentation},
	}
	IdentityDomainsIdentityProofingProviderTemplateConfigurationRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"sensitivity": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"type":        acctest.Representation{RepType: acctest.Required, Create: `type`, Update: `type2`},
	}

	IdentityDomainsIdentityProofingProviderTemplateResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsIdentityProofingProviderTemplateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsIdentityProofingProviderTemplateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_identity_proofing_provider_template.test_identity_proofing_provider_template"
	datasourceName := "data.oci_identity_domains_identity_proofing_provider_templates.test_identity_proofing_provider_templates"
	singularDatasourceName := "data.oci_identity_domains_identity_proofing_provider_template.test_identity_proofing_provider_template"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsIdentityProofingProviderTemplateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Optional, acctest.Create, IdentityDomainsIdentityProofingProviderTemplateRepresentation), "identitydomains", "identityProofingProviderTemplate", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsIdentityProofingProviderTemplateDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identity_proofing_provider_template_provider", "identityProofingProviderTemplateProvider"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verification_url", "verificationUrl"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderTemplateResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Optional, acctest.Create, IdentityDomainsIdentityProofingProviderTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identity_proofing_provider_template_provider", "identityProofingProviderTemplateProvider"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "configuration", map[string]string{
					"name":        "name",
					"sensitivity": "false",
					"type":        "type",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verification_url", "verificationUrl"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}
					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "identityProofingProviderTemplates", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identity_proofing_provider_template_provider", "identityProofingProviderTemplateProvider2"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "configuration", map[string]string{
					"name":        "name2",
					"sensitivity": "true",
					"type":        "type2",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "verification_url", "verificationUrl2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_templates", "test_identity_proofing_provider_templates", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderTemplateDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderTemplateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Optional, acctest.Update, IdentityDomainsIdentityProofingProviderTemplateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "identity_proofing_provider_template_count", "10"),
				//resource.TestCheckResourceAttr(datasourceName, "identity_proofing_provider_template_filter", "identityProofingProviderTemplateFilter"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "identity_proofing_provider_templates.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "identity_proofing_provider_templates.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_proofing_provider_template", "test_identity_proofing_provider_template", acctest.Required, acctest.Create, IdentityDomainsIdentityProofingProviderTemplateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityProofingProviderTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_proofing_provider_template_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "identity_proofing_provider_template_provider", "identityProofingProviderTemplateProvider2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "configuration", map[string]string{
					"name":        "name2",
					"sensitivity": "true",
					"type":        "type2",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "verification_url", "verificationUrl2"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsIdentityProofingProviderTemplateRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_identity_proofing_provider_template", "identityProofingProviderTemplates"),
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

func testAccCheckIdentityDomainsIdentityProofingProviderTemplateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_identity_proofing_provider_template" {
			noResourceFound = false
			request := oci_identity_domains.GetIdentityProofingProviderTemplateRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.IdentityProofingProviderTemplateId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetIdentityProofingProviderTemplate(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsIdentityProofingProviderTemplate") {
		resource.AddTestSweepers("IdentityDomainsIdentityProofingProviderTemplate", &resource.Sweeper{
			Name:         "IdentityDomainsIdentityProofingProviderTemplate",
			Dependencies: acctest.DependencyGraph["identityProofingProviderTemplate"],
			F:            sweepIdentityDomainsIdentityProofingProviderTemplateResource,
		})
	}
}

func sweepIdentityDomainsIdentityProofingProviderTemplateResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	identityProofingProviderTemplateIds, err := getIdentityDomainsIdentityProofingProviderTemplateIds(compartment)
	if err != nil {
		return err
	}
	for _, identityProofingProviderTemplateId := range identityProofingProviderTemplateIds {
		if ok := acctest.SweeperDefaultResourceId[identityProofingProviderTemplateId]; !ok {
			deleteIdentityProofingProviderTemplateRequest := oci_identity_domains.DeleteIdentityProofingProviderTemplateRequest{}

			deleteIdentityProofingProviderTemplateRequest.IdentityProofingProviderTemplateId = &identityProofingProviderTemplateId

			deleteIdentityProofingProviderTemplateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteIdentityProofingProviderTemplate(context.Background(), deleteIdentityProofingProviderTemplateRequest)
			if error != nil {
				fmt.Printf("Error deleting IdentityProofingProviderTemplate %s %s, It is possible that the resource is already deleted. Please verify manually \n", identityProofingProviderTemplateId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsIdentityProofingProviderTemplateIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IdentityProofingProviderTemplateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listIdentityProofingProviderTemplatesRequest := oci_identity_domains.ListIdentityProofingProviderTemplatesRequest{}
	//listIdentityProofingProviderTemplatesRequest.CompartmentId = &compartmentId
	listIdentityProofingProviderTemplatesResponse, err := identityDomainsClient.ListIdentityProofingProviderTemplates(context.Background(), listIdentityProofingProviderTemplatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IdentityProofingProviderTemplate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, identityProofingProviderTemplate := range listIdentityProofingProviderTemplatesResponse.Resources {
		id := *identityProofingProviderTemplate.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IdentityProofingProviderTemplateId", id)
	}
	return resourceIds, nil
}
