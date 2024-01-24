// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_limits "github.com/oracle/oci-go-sdk/v65/limits"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LimitsQuotaRequiredOnlyResource = LimitsQuotaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Required, acctest.Create, LimitsQuotaRepresentation)

	LimitsQuotaResourceConfig = LimitsQuotaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Update, LimitsQuotaRepresentation)

	QuotaWithoutLockRequiredOnlyResource = LimitsQuotaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Required, acctest.Create, LimitsQuotaWithoutLockRepresentation)

	QuotaWithoutLockResourceConfig = LimitsQuotaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Update, LimitsQuotaWithoutLockRepresentation)

	LimitsLimitsQuotaSingularDataSourceRepresentation = map[string]interface{}{
		"quota_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_limits_quota.test_quota.id}`},
	}

	LimitsLimitsQuotaDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `ComputeQuotas`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LimitsQuotaDataSourceFilterRepresentation}}
	LimitsQuotaDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_limits_quota.test_quota.id}`}},
	}

	LimitsQuotaRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Quotas for Compute VM.DenseIO1.16 resources`, Update: `description2`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `ComputeQuotas`},
		"statements":     acctest.Representation{RepType: acctest.Required, Create: []string{`Set notifications quota topic-count to 99 in tenancy`}, Update: []string{`Set notifications quota topic-count to 99 in tenancy`, `Set resource-manager quota stack-count to 499 in tenancy`}},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: LimitsQuotaLocksRepresentation},
	}
	LimitsQuotaLocksRepresentation = map[string]interface{}{
		"type":                acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":             acctest.Representation{RepType: acctest.Optional, Create: `message`},
		"related_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
	}

	LimitsQuotaWithoutLockRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Quotas for Compute VM.DenseIO1.16 resources`, Update: `description2`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `ComputeQuotas`},
		"statements":     acctest.Representation{RepType: acctest.Required, Create: []string{`Set notifications quota topic-count to 99 in tenancy`}, Update: []string{`Set notifications quota topic-count to 99 in tenancy`, `Set resource-manager quota stack-count to 499 in tenancy`}},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	LimitsQuotaResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: limits/default
func TestLimitsQuotaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsQuotaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_limits_quota.test_quota"
	datasourceName := "data.oci_limits_quotas.test_quotas"
	singularDatasourceName := "data.oci_limits_quota.test_quota"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LimitsQuotaResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Create, LimitsQuotaRepresentation), "limits", "quota", t)

	acctest.ResourceTest(t, testAccCheckLimitsQuotaDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Required, acctest.Create, LimitsQuotaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Quotas for Compute VM.DenseIO1.16 resources"),
				resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Create, LimitsQuotaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Quotas for Compute VM.DenseIO1.16 resources"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.related_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Update, LimitsQuotaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.related_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_quotas", "test_quotas", acctest.Optional, acctest.Update, LimitsLimitsQuotaDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Update, LimitsQuotaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "quotas.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.locks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.locks.0.related_resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.locks.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.locks.0.type", "FULL"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.name", "ComputeQuotas"),
				resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "quotas.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Required, acctest.Create, LimitsLimitsQuotaSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsQuotaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quota_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.0.type", "FULL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statements.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + LimitsQuotaRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: limits/default
func TestLimitsQuotaResource_basic_no_lock(t *testing.T) {
	httpreplay.SetScenario("TestLimitsQuotaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_limits_quota.test_quota"
	datasourceName := "data.oci_limits_quotas.test_quotas"
	singularDatasourceName := "data.oci_limits_quota.test_quota"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LimitsQuotaResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Create, LimitsQuotaWithoutLockRepresentation), "limits", "quota", t)

	acctest.ResourceTest(t, testAccCheckLimitsQuotaDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Required, acctest.Create, LimitsQuotaWithoutLockRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Quotas for Compute VM.DenseIO1.16 resources"),
				resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Create, LimitsQuotaWithoutLockRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Quotas for Compute VM.DenseIO1.16 resources"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Update, LimitsQuotaWithoutLockRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_quotas", "test_quotas", acctest.Optional, acctest.Update, LimitsLimitsQuotaDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsQuotaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Optional, acctest.Update, LimitsQuotaWithoutLockRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "quotas.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "quotas.0.compartment_id", tenancyId),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_quota", "test_quota", acctest.Required, acctest.Create, LimitsLimitsQuotaSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsQuotaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quota_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "ComputeQuotas"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statements.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + QuotaWithoutLockRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLimitsQuotaDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).QuotasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_limits_quota" {
			noResourceFound = false
			request := oci_limits.GetQuotaRequest{}

			tmp := rs.Primary.ID
			request.QuotaId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "limits")

			_, err := client.GetQuota(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("LimitsQuota") {
		resource.AddTestSweepers("LimitsQuota", &resource.Sweeper{
			Name:         "LimitsQuota",
			Dependencies: acctest.DependencyGraph["quota"],
			F:            sweepLimitsQuotaResource,
		})
	}
}

func sweepLimitsQuotaResource(compartment string) error {
	quotasClient := acctest.GetTestClients(&schema.ResourceData{}).QuotasClient()
	// LimitsQuotaResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	quotaIds, err := getLimitsQuotaIds(compartment)
	isLockOverride := true
	if err != nil {
		return err
	}
	for _, quotaId := range quotaIds {
		if ok := acctest.SweeperDefaultResourceId[quotaId]; !ok {
			deleteQuotaRequest := oci_limits.DeleteQuotaRequest{}

			deleteQuotaRequest.QuotaId = &quotaId
			deleteQuotaRequest.IsLockOverride = &isLockOverride

			deleteQuotaRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "limits")
			_, error := quotasClient.DeleteQuota(context.Background(), deleteQuotaRequest)
			if error != nil {
				fmt.Printf("Error deleting Quota %s %s, It is possible that the resource is already deleted. Please verify manually \n", quotaId, error)
				continue
			}
		}
	}
	return nil
}

func getLimitsQuotaIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "QuotaId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	quotasClient := acctest.GetTestClients(&schema.ResourceData{}).QuotasClient()

	listQuotasRequest := oci_limits.ListQuotasRequest{}
	listQuotasRequest.CompartmentId = &compartmentId
	listQuotasResponse, err := quotasClient.ListQuotas(context.Background(), listQuotasRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Quota list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, quota := range listQuotasResponse.Items {
		id := *quota.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "QuotaId", id)
	}
	return resourceIds, nil
}
