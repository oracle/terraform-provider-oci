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
	IdentityDomainsIdentityPropagationTrustRequiredOnlyResource = IdentityDomainsIdentityPropagationTrustResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Required, acctest.Create, IdentityDomainsIdentityPropagationTrustRepresentation)

	IdentityDomainsIdentityPropagationTrustResourceConfig = IdentityDomainsIdentityPropagationTrustResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Optional, acctest.Update, IdentityDomainsIdentityPropagationTrustRepresentation)

	IdentityDomainsIdentityPropagationTrustSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_propagation_trust_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_identity_propagation_trust.test_identity_propagation_trust.id}`},
		"attribute_sets":                acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityPropagationTrustDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"identity_propagation_trust_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"identity_propagation_trust_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                       acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsIdentityPropagationTrustRepresentation = map[string]interface{}{
		"idcs_endpoint":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"issuer":                      acctest.Representation{RepType: acctest.Required, Create: `issuer`, Update: `issuer2`},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `name`},
		"schemas":                     acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:IdentityPropagationTrust`}},
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `JWT`, Update: `SAML`},
		"account_id":                  acctest.Representation{RepType: acctest.Optional, Create: `accountId`},
		"active":                      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"allow_impersonation":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"attribute_sets":              acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"client_claim_name":           acctest.Representation{RepType: acctest.Optional, Create: `clientClaimName`, Update: `clientClaimName2`},
		"client_claim_values":         acctest.Representation{RepType: acctest.Optional, Create: []string{`clientClaimValues`}, Update: []string{`clientClaimValues2`}},
		"clock_skew_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"impersonation_service_users": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityPropagationTrustImpersonationServiceUsersRepresentation},
		"keytab":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityPropagationTrustKeytabRepresentation},
		// Set oauth_clients as Required to avoid error: oauthClients must be defined for IdentityPropagationTrust
		"oauth_clients":             acctest.Representation{RepType: acctest.Required, Create: []string{`oauthClients`}, Update: []string{`oauthClients2`}},
		"public_certificate":        acctest.Representation{RepType: acctest.Optional, Create: `publicCertificate`, Update: `publicCertificate2`},
		"public_key_endpoint":       acctest.Representation{RepType: acctest.Optional, Create: `publicKeyEndpoint`, Update: `publicKeyEndpoint2`},
		"subject_claim_name":        acctest.Representation{RepType: acctest.Optional, Create: `subjectClaimName`, Update: `subjectClaimName2`},
		"subject_mapping_attribute": acctest.Representation{RepType: acctest.Optional, Create: `subjectMappingAttribute`, Update: `subjectMappingAttribute2`},
		"subject_type":              acctest.Representation{RepType: acctest.Optional, Create: `User`, Update: `App`},
		"tags":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsIdentityPropagationTrustTagsRepresentation},
	}
	IdentityDomainsIdentityPropagationTrustImpersonationServiceUsersRepresentation = map[string]interface{}{
		"rule":  acctest.Representation{RepType: acctest.Required, Create: `rule`, Update: `rule2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}
	IdentityDomainsIdentityPropagationTrustKeytabRepresentation = map[string]interface{}{
		"secret_ocid":    acctest.Representation{RepType: acctest.Required, Create: `secretOcid`, Update: `secretOcid2`},
		"secret_version": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	IdentityDomainsIdentityPropagationTrustTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsIdentityPropagationTrustResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsIdentityPropagationTrustResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsIdentityPropagationTrustResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_identity_propagation_trust.test_identity_propagation_trust"
	datasourceName := "data.oci_identity_domains_identity_propagation_trusts.test_identity_propagation_trusts"
	singularDatasourceName := "data.oci_identity_domains_identity_propagation_trust.test_identity_propagation_trust"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsIdentityPropagationTrustResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Optional, acctest.Create, IdentityDomainsIdentityPropagationTrustRepresentation), "identitydomains", "identityPropagationTrust", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsIdentityPropagationTrustDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityPropagationTrustResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Required, acctest.Create, IdentityDomainsIdentityPropagationTrustRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "issuer", "issuer"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "JWT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityPropagationTrustResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityPropagationTrustResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Optional, acctest.Create, IdentityDomainsIdentityPropagationTrustRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "account_id"),
				resource.TestCheckResourceAttr(resourceName, "active", "false"),
				resource.TestCheckResourceAttr(resourceName, "allow_impersonation", "false"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "client_claim_name", "clientClaimName"),
				resource.TestCheckResourceAttr(resourceName, "clock_skew_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "impersonation_service_users.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "impersonation_service_users", map[string]string{
					"rule": "rule",
				},
					[]string{
						"value",
					}),
				resource.TestCheckResourceAttr(resourceName, "issuer", "issuer"),
				resource.TestCheckResourceAttr(resourceName, "keytab.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "keytab.0.secret_ocid", "secretOcid"),
				resource.TestCheckResourceAttr(resourceName, "keytab.0.secret_version", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", "publicCertificate"),
				resource.TestCheckResourceAttr(resourceName, "public_key_endpoint", "publicKeyEndpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subject_claim_name", "subjectClaimName"),
				resource.TestCheckResourceAttr(resourceName, "subject_mapping_attribute", "subjectMappingAttribute"),
				resource.TestCheckResourceAttr(resourceName, "subject_type", "User"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "type", "JWT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "identityPropagationTrusts", resId)
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
			Config: config + compartmentIdVariableStr + IdentityDomainsIdentityPropagationTrustResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Optional, acctest.Update, IdentityDomainsIdentityPropagationTrustRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "account_id"),
				resource.TestCheckResourceAttr(resourceName, "active", "true"),
				resource.TestCheckResourceAttr(resourceName, "allow_impersonation", "true"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "client_claim_name", "clientClaimName2"),
				resource.TestCheckResourceAttr(resourceName, "clock_skew_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "impersonation_service_users.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "impersonation_service_users", map[string]string{
					"rule": "rule2",
				},
					[]string{
						"value",
					}),
				resource.TestCheckResourceAttr(resourceName, "issuer", "issuer2"),
				resource.TestCheckResourceAttr(resourceName, "keytab.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "keytab.0.secret_ocid", "secretOcid2"),
				resource.TestCheckResourceAttr(resourceName, "keytab.0.secret_version", "11"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "public_certificate", "publicCertificate2"),
				resource.TestCheckResourceAttr(resourceName, "public_key_endpoint", "publicKeyEndpoint2"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subject_claim_name", "subjectClaimName2"),
				resource.TestCheckResourceAttr(resourceName, "subject_mapping_attribute", "subjectMappingAttribute2"),
				resource.TestCheckResourceAttr(resourceName, "subject_type", "App"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "type", "SAML"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_propagation_trusts", "test_identity_propagation_trusts", acctest.Optional, acctest.Update, IdentityDomainsIdentityPropagationTrustDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityPropagationTrustResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Optional, acctest.Update, IdentityDomainsIdentityPropagationTrustRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "identity_propagation_trusts.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestMatchResourceAttr(datasourceName, "identity_propagation_trusts.0.schemas.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_identity_propagation_trust", "test_identity_propagation_trust", acctest.Optional, acctest.Create, IdentityDomainsIdentityPropagationTrustSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsIdentityPropagationTrustResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_propagation_trust_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allow_impersonation", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "client_claim_name", "clientClaimName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "clock_skew_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "impersonation_service_users.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "impersonation_service_users", map[string]string{
					"rule": "rule2",
				},
					[]string{
						"value",
					}),
				resource.TestCheckResourceAttr(singularDatasourceName, "issuer", "issuer2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "keytab.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "keytab.0.secret_ocid", "secretOcid2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "keytab.0.secret_version", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "public_certificate", "publicCertificate2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "public_key_endpoint", "publicKeyEndpoint2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subject_claim_name", "subjectClaimName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subject_mapping_attribute", "subjectMappingAttribute2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subject_type", "App"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "SAML"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsIdentityPropagationTrustRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_identity_propagation_trust", "identityPropagationTrusts"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				// properties that are not "returned: default"
				"tags", //
				"impersonation_service_users",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsIdentityPropagationTrustDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_identity_propagation_trust" {
			noResourceFound = false
			request := oci_identity_domains.GetIdentityPropagationTrustRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.IdentityPropagationTrustId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetIdentityPropagationTrust(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsIdentityPropagationTrust") {
		resource.AddTestSweepers("IdentityDomainsIdentityPropagationTrust", &resource.Sweeper{
			Name:         "IdentityDomainsIdentityPropagationTrust",
			Dependencies: acctest.DependencyGraph["identityPropagationTrust"],
			F:            sweepIdentityDomainsIdentityPropagationTrustResource,
		})
	}
}

func sweepIdentityDomainsIdentityPropagationTrustResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	identityPropagationTrustIds, err := getIdentityDomainsIdentityPropagationTrustIds(compartment)
	if err != nil {
		return err
	}
	for _, identityPropagationTrustId := range identityPropagationTrustIds {
		if ok := acctest.SweeperDefaultResourceId[identityPropagationTrustId]; !ok {
			deleteIdentityPropagationTrustRequest := oci_identity_domains.DeleteIdentityPropagationTrustRequest{}

			deleteIdentityPropagationTrustRequest.IdentityPropagationTrustId = &identityPropagationTrustId

			deleteIdentityPropagationTrustRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteIdentityPropagationTrust(context.Background(), deleteIdentityPropagationTrustRequest)
			if error != nil {
				fmt.Printf("Error deleting IdentityPropagationTrust %s %s, It is possible that the resource is already deleted. Please verify manually \n", identityPropagationTrustId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsIdentityPropagationTrustIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IdentityPropagationTrustId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listIdentityPropagationTrustsRequest := oci_identity_domains.ListIdentityPropagationTrustsRequest{}
	listIdentityPropagationTrustsResponse, err := identityDomainsClient.ListIdentityPropagationTrusts(context.Background(), listIdentityPropagationTrustsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IdentityPropagationTrust list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, identityPropagationTrust := range listIdentityPropagationTrustsResponse.Resources {
		id := *identityPropagationTrust.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IdentityPropagationTrustId", id)
	}
	return resourceIds, nil
}
