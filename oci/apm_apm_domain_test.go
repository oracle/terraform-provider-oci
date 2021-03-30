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
	oci_apm "github.com/oracle/oci-go-sdk/v42/apmcontrolplane"
	"github.com/oracle/oci-go-sdk/v42/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApmDomainRequiredOnlyResource = ApmDomainResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation)

	ApmDomainResourceConfig = ApmDomainResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Update, apmDomainRepresentation)

	apmDomainSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
	}

	apmDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, apmDomainDataSourceFilterRepresentation}}
	apmDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_apm_apm_domain.test_apm_domain.id}`}},
	}

	apmDomainRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"is_free_tier":   Representation{repType: Optional, create: `false`},
		"lifecycle":      RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}

	ignoreDefinedTagsDifferencesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	ApmDomainResourceDependencies = DefinedTagsDependencies
)

func TestApmApmDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmApmDomainResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apm_apm_domain.test_apm_domain"
	datasourceName := "data.oci_apm_apm_domains.test_apm_domains"
	singularDatasourceName := "data.oci_apm_apm_domain.test_apm_domain"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ApmDomainResourceDependencies+
		generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Create, apmDomainRepresentation), "apm", "apmDomain", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmApmDomainDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ApmDomainResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Create, apmDomainRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApmDomainResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Create,
						representationCopyWithNewProperties(apmDomainRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Update, apmDomainRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"), // don't validate defined tags since there are some pre-created ones
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_free_tier", "false"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_apm_apm_domains", "test_apm_domains", Optional, Update, apmDomainDataSourceRepresentation) +
					compartmentIdVariableStr + ApmDomainResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Optional, Update, apmDomainRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApmDomainResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apm")

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
	if !inSweeperExcludeList("ApmApmDomain") {
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

			deleteApmDomainRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apm")
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
	ids := getResourceIdsToSweep(compartment, "ApmDomainId")
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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ApmDomainId", id)
	}
	return resourceIds, nil
}
