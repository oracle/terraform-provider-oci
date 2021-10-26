// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apm "github.com/oracle/oci-go-sdk/v50/apmcontrolplane"
	"github.com/oracle/oci-go-sdk/v50/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApmDomainRequiredOnlyResource = ApmDomainResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation)

	ApmDomainResourceConfig = ApmDomainResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Update, apmDomainRepresentation)

	apmDomainSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{RepType: Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
	}

	apmDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, apmDomainDataSourceFilterRepresentation}}
	apmDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_apm_apm_domain.test_apm_domain.id}`}},
	}

	apmDomainRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_free_tier":   Representation{RepType: Optional, Create: `false`},
		"lifecycle":      RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}

	ignoreDefinedTagsDifferencesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	ApmDomainResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: apm/default
func TestApmApmDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmApmDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apm_apm_domain.test_apm_domain"
	datasourceName := "data.oci_apm_apm_domains.test_apm_domains"
	singularDatasourceName := "data.oci_apm_apm_domain.test_apm_domain"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ApmDomainResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Create, apmDomainRepresentation), "apm", "apmDomain", t)

	ResourceTest(t, testAccCheckApmApmDomainDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Create, apmDomainRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApmDomainResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Create,
					RepresentationCopyWithNewProperties(apmDomainRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Update, apmDomainRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_apm_apm_domains", "test_apm_domains", Optional, Update, apmDomainDataSourceRepresentation) +
				compartmentIdVariableStr + ApmDomainResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Update, apmDomainRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "apm_domains.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "apm_domains.0.compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(datasourceName, "apm_domains.0.defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
				resource.TestCheckResourceAttr(datasourceName, "apm_domains.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "apm_domains.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "apm_domains.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domains.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "apm_domains.0.is_free_tier", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domains.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domains.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domains.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmDomainResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_upload_endpoint"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_free_tier", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ApmDomainResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckApmApmDomainDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).apmDomainClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_apm_domain" {
			noResourceFound = false
			request := oci_apm.GetApmDomainRequest{}

			tmp := rs.Primary.ID
			request.ApmDomainId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "apm")

			_, err := client.GetApmDomain(context.Background(), request)

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("ApmApmDomain") {
		resource.AddTestSweepers("ApmApmDomain", &resource.Sweeper{
			Name:         "ApmApmDomain",
			Dependencies: DependencyGraph["apmDomain"],
			F:            sweepApmApmDomainResource,
		})
	}
}

func sweepApmApmDomainResource(compartment string) error {
	apmDomainClient := GetTestClients(&schema.ResourceData{}).apmDomainClient()
	apmDomainIds, err := getApmDomainIds(compartment)
	if err != nil {
		return err
	}
	for _, apmDomainId := range apmDomainIds {
		if ok := SweeperDefaultResourceId[apmDomainId]; !ok {
			deleteApmDomainRequest := oci_apm.DeleteApmDomainRequest{}

			deleteApmDomainRequest.ApmDomainId = &apmDomainId

			deleteApmDomainRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "apm")
			_, error := apmDomainClient.DeleteApmDomain(context.Background(), deleteApmDomainRequest)
			if error != nil {
				fmt.Printf("Error deleting ApmDomain %s %s, It is possible that the resource is already deleted. Please verify manually \n", apmDomainId, error)
				continue
			}
		}
	}
	return nil
}

func getApmDomainIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ApmDomainId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmDomainClient := GetTestClients(&schema.ResourceData{}).apmDomainClient()

	listApmDomainsRequest := oci_apm.ListApmDomainsRequest{}
	listApmDomainsRequest.CompartmentId = &compartmentId
	listApmDomainsResponse, err := apmDomainClient.ListApmDomains(context.Background(), listApmDomainsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ApmDomain list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, apmDomain := range listApmDomainsResponse.Items {
		id := *apmDomain.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ApmDomainId", id)
	}
	return resourceIds, nil
}
