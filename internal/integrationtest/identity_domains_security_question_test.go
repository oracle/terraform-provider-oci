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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	IdentityDomainsSecurityQuestionRequiredOnlyResource = IdentityDomainsSecurityQuestionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Required, acctest.Create, IdentityDomainsSecurityQuestionRepresentation)

	IdentityDomainsSecurityQuestionResourceConfig = IdentityDomainsSecurityQuestionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Optional, acctest.Update, IdentityDomainsSecurityQuestionRepresentation)

	IdentityDomainsIdentityDomainsSecurityQuestionSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"security_question_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_security_question.test_security_question.id}`},
		"attribute_sets":       acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsSecurityQuestionDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"security_question_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"security_question_filter": acctest.Representation{RepType: acctest.Optional, Create: `type eq \"custom\"`},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsSecurityQuestionRepresentation = map[string]interface{}{
		"active":         acctest.Representation{RepType: acctest.Required, Create: `false`},
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"question_text":  acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsSecurityQuestionQuestionTextRepresentation},
		"schemas":        acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:SecurityQuestion`}},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `custom`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"external_id":    acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"tags":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsSecurityQuestionTagsRepresentation},
	}
	IdentityDomainsSecurityQuestionQuestionTextRepresentation = map[string]interface{}{
		"locale":  acctest.Representation{RepType: acctest.Required, Create: `en`},
		"value":   acctest.Representation{RepType: acctest.Required, Create: `value`},
		"default": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	IdentityDomainsSecurityQuestionTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsSecurityQuestionResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsSecurityQuestionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsSecurityQuestionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_security_question.test_security_question"
	datasourceName := "data.oci_identity_domains_security_questions.test_security_questions"
	singularDatasourceName := "data.oci_identity_domains_security_question.test_security_question"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsSecurityQuestionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Optional, acctest.Create, IdentityDomainsSecurityQuestionRepresentation), "identitydomains", "securityQuestion", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsSecurityQuestionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Required, acctest.Create, IdentityDomainsSecurityQuestionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "question_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "question_text.0.locale", "en"),
				resource.TestCheckResourceAttr(resourceName, "question_text.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "custom"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsSecurityQuestionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Optional, acctest.Create, IdentityDomainsSecurityQuestionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "question_text.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "question_text.0.default", "true"),
				resource.TestCheckResourceAttr(resourceName, "question_text.0.locale", "en"),
				resource.TestCheckResourceAttr(resourceName, "question_text.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "type", "custom"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "securityQuestions", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_security_questions", "test_security_questions", acctest.Optional, acctest.Update, IdentityDomainsIdentityDomainsSecurityQuestionDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSecurityQuestionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Optional, acctest.Update, IdentityDomainsSecurityQuestionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "security_question_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "security_question_filter", "type eq \"custom\""),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),
				resource.TestMatchResourceAttr(datasourceName, "security_questions.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "security_questions.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_security_question", "test_security_question", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsSecurityQuestionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsSecurityQuestionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_question_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "question_text.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "question_text.0.default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "question_text.0.locale", "en"),
				resource.TestCheckResourceAttr(singularDatasourceName, "question_text.0.value", "value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "custom"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsSecurityQuestionRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_security_question", "securityQuestions"),
			ImportStateVerify: true,
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

func testAccCheckIdentityDomainsSecurityQuestionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_security_question" {
			noResourceFound = false
			request := oci_identity_domains.GetSecurityQuestionRequest{}

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
			request.SecurityQuestionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetSecurityQuestion(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsSecurityQuestion") {
		resource.AddTestSweepers("IdentityDomainsSecurityQuestion", &resource.Sweeper{
			Name:         "IdentityDomainsSecurityQuestion",
			Dependencies: acctest.DependencyGraph["securityQuestion"],
			F:            sweepIdentityDomainsSecurityQuestionResource,
		})
	}
}

func sweepIdentityDomainsSecurityQuestionResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	securityQuestionIds, err := getIdentityDomainsSecurityQuestionIds(compartment)
	if err != nil {
		return err
	}
	for _, securityQuestionId := range securityQuestionIds {
		if ok := acctest.SweeperDefaultResourceId[securityQuestionId]; !ok {
			deleteSecurityQuestionRequest := oci_identity_domains.DeleteSecurityQuestionRequest{}

			deleteSecurityQuestionRequest.SecurityQuestionId = &securityQuestionId

			deleteSecurityQuestionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteSecurityQuestion(context.Background(), deleteSecurityQuestionRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityQuestion %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityQuestionId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsSecurityQuestionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityQuestionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listSecurityQuestionsRequest := oci_identity_domains.ListSecurityQuestionsRequest{}
	listSecurityQuestionsResponse, err := identityDomainsClient.ListSecurityQuestions(context.Background(), listSecurityQuestionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityQuestion list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityQuestion := range listSecurityQuestionsResponse.Resources {
		id := *securityQuestion.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityQuestionId", id)
	}
	return resourceIds, nil
}
