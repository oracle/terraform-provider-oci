// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v26/common"
	oci_limits "github.com/oracle/oci-go-sdk/v26/limits"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	QuotaRequiredOnlyResource = QuotaResourceDependencies +
		generateResourceFromRepresentationMap("oci_limits_quota", "test_quota", Required, Create, quotaRepresentation)

	QuotaResourceConfig = QuotaResourceDependencies +
		generateResourceFromRepresentationMap("oci_limits_quota", "test_quota", Optional, Update, quotaRepresentation)

	quotaSingularDataSourceRepresentation = map[string]interface{}{
		"quota_id": Representation{repType: Required, create: `${oci_limits_quota.test_quota.id}`},
	}

	quotaDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"name":           Representation{repType: Optional, create: `ComputeQuotas`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, quotaDataSourceFilterRepresentation}}
	quotaDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_limits_quota.test_quota.id}`}},
	}

	quotaRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `Quotas for Compute VM.DenseIO1.16 resources`, update: `description2`},
		"name":           Representation{repType: Required, create: `ComputeQuotas`},
		"statements":     Representation{repType: Required, create: []string{`Set notifications quota topic-count to 99 in tenancy`}, update: []string{`Set notifications quota topic-count to 99 in tenancy`, `Set resource-manager quota stack-count to 499 in tenancy`}},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	QuotaResourceDependencies = DefinedTagsDependencies
)

func TestLimitsQuotaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsQuotaResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_limits_quota.test_quota"
	datasourceName := "data.oci_limits_quotas.test_quotas"
	singularDatasourceName := "data.oci_limits_quota.test_quota"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLimitsQuotaDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + QuotaResourceDependencies +
					generateResourceFromRepresentationMap("oci_limits_quota", "test_quota", Required, Create, quotaRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "Quotas for Compute VM.DenseIO1.16 resources"),
					resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
					resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + QuotaResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + QuotaResourceDependencies +
					generateResourceFromRepresentationMap("oci_limits_quota", "test_quota", Optional, Create, quotaRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Quotas for Compute VM.DenseIO1.16 resources"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
					resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + QuotaResourceDependencies +
					generateResourceFromRepresentationMap("oci_limits_quota", "test_quota", Optional, Update, quotaRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
					resource.TestCheckResourceAttr(resourceName, "statements.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_limits_quotas", "test_quotas", Optional, Update, quotaDataSourceRepresentation) +
					compartmentIdVariableStr + QuotaResourceDependencies +
					generateResourceFromRepresentationMap("oci_limits_quota", "test_quota", Optional, Update, quotaRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "name", "ComputeQuotas"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "quotas.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "quotas.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "quotas.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "quotas.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "quotas.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "quotas.0.name", "ComputeQuotas"),
					resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_limits_quota", "test_quota", Required, Create, quotaSingularDataSourceRepresentation) +
					compartmentIdVariableStr + QuotaResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "quota_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "ComputeQuotas"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "statements.#", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + QuotaResourceConfig,
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

func testAccCheckLimitsQuotaDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).quotasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_limits_quota" {
			noResourceFound = false
			request := oci_limits.GetQuotaRequest{}

			tmp := rs.Primary.ID
			request.QuotaId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "limits")

			_, err := client.GetQuota(context.Background(), request)

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
	if !inSweeperExcludeList("LimitsQuota") {
		resource.AddTestSweepers("LimitsQuota", &resource.Sweeper{
			Name:         "LimitsQuota",
			Dependencies: DependencyGraph["quota"],
			F:            sweepLimitsQuotaResource,
		})
	}
}

func sweepLimitsQuotaResource(compartment string) error {
	quotasClient := GetTestClients(&schema.ResourceData{}).quotasClient()
	// LimitsQuotaResource can only run on root compartment
	compartment = getEnvSettingWithBlankDefault("tenancy_ocid")
	quotaIds, err := getQuotaIds(compartment)
	if err != nil {
		return err
	}
	for _, quotaId := range quotaIds {
		if ok := SweeperDefaultResourceId[quotaId]; !ok {
			deleteQuotaRequest := oci_limits.DeleteQuotaRequest{}

			deleteQuotaRequest.QuotaId = &quotaId

			deleteQuotaRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "limits")
			_, error := quotasClient.DeleteQuota(context.Background(), deleteQuotaRequest)
			if error != nil {
				fmt.Printf("Error deleting Quota %s %s, It is possible that the resource is already deleted. Please verify manually \n", quotaId, error)
				continue
			}
		}
	}
	return nil
}

func getQuotaIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "QuotaId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	quotasClient := GetTestClients(&schema.ResourceData{}).quotasClient()

	listQuotasRequest := oci_limits.ListQuotasRequest{}
	listQuotasRequest.CompartmentId = &compartmentId
	listQuotasResponse, err := quotasClient.ListQuotas(context.Background(), listQuotasRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Quota list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, quota := range listQuotasResponse.Items {
		id := *quota.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "QuotaId", id)
	}
	return resourceIds, nil
}
