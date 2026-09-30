// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License 2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func testCheckTableEstimateExists(dataSourceName, schemaName, tableName, targetID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		itemsCountValue, err := acctest.FromInstanceState(s, dataSourceName, "table_estimate_collection.0.items.#")
		if err != nil {
			return err
		}
		itemsCount, err := strconv.Atoi(itemsCountValue)
		if err != nil {
			return fmt.Errorf("invalid table estimate item count %q: %w", itemsCountValue, err)
		}

		for i := 0; i < itemsCount; i++ {
			itemPrefix := fmt.Sprintf("table_estimate_collection.0.items.%d", i)
			actualSchemaName, err := acctest.FromInstanceState(s, dataSourceName, itemPrefix+".schema_name")
			if err != nil {
				return err
			}
			actualTableName, err := acctest.FromInstanceState(s, dataSourceName, itemPrefix+".table_name")
			if err != nil {
				return err
			}
			actualTargetID, err := acctest.FromInstanceState(s, dataSourceName, itemPrefix+".target_id")
			if err != nil {
				return err
			}

			if actualSchemaName == schemaName && actualTableName == tableName && actualTargetID == targetID {
				return nil
			}
		}

		return fmt.Errorf("table estimate not found for schema_name=%q, table_name=%q, target_id=%q", schemaName, tableName, targetID)
	}
}

var (
	dataSafeEstimateSubsettingPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"schema_source":  acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSubsettingPolicySchemaSourceRepresentation},
	}

	dataSafeEstimateTableSizesRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"target_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"target_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"user_name": acctest.Representation{RepType: acctest.Required, Create: `${var.target_database_user_name}`},
			"password":  acctest.Representation{RepType: acctest.Required, Create: `${var.target_database_password}`},
		}},
	}

	dataSafeSubsettingPolicyTableEstimateDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"object":               acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEES`}},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`HR_TEST`}},
		"target_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeEstimateTableSizesResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeEstimateTableSizesResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	targetID := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetDatabaseUserName := utils.GetEnvSettingWithBlankDefault("target_database_user_name")
	targetDatabasePassword := utils.GetEnvSettingWithBlankDefault("target_database_password")

	variables := fmt.Sprintf(`
variable "compartment_id" {
	default = "%s"
}

variable "target_id" {
	default = "%s"
}

variable "target_database_user_name" {
	default = "%s"
}

variable "target_database_password" {
	default = "%s"
}
`, compartmentID, targetID, targetDatabaseUserName, targetDatabasePassword)

	dependencies := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, dataSafeEstimateSubsettingPolicyRepresentation)
	dependencies += acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "test_subsetting_policy_subsetting_rule", acctest.Required, acctest.Create, DataSafeSubsettingPolicySubsettingRuleRepresentation)

	testEstimateResource := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_estimate_table_sizes", "test_estimate_table_sizes", acctest.Required, acctest.Create, dataSafeEstimateTableSizesRepresentation)
	testEstimateResource = strings.TrimSuffix(testEstimateResource, "}\n") + "depends_on = [\"oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule\"]\n}\n"

	tableEstimatesDataSource := acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_table_estimates", "test_subsetting_policy_table_estimates", acctest.Required, acctest.Create, dataSafeSubsettingPolicyTableEstimateDataSourceRepresentation)
	tableEstimatesDataSource = strings.TrimSuffix(tableEstimatesDataSource, "}\n") + "depends_on = [\"oci_data_safe_estimate_table_sizes.test_estimate_table_sizes\"]\n}\n"

	resourceName := "oci_data_safe_estimate_table_sizes.test_estimate_table_sizes"
	datasourceName := "data.oci_data_safe_subsetting_policy_table_estimates.test_subsetting_policy_table_estimates"

	acctest.ResourceTest(t, nil, []resource.TestStep{{
		Config: config + variables + dependencies + testEstimateResource + tableEstimatesDataSource,
		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
			resource.TestCheckResourceAttrSet(resourceName, "id"),
			resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_id"),
			resource.TestCheckResourceAttr(datasourceName, "table_estimate_collection.#", "1"),
			testCheckTableEstimateExists(datasourceName, "HR_TEST", "EMPLOYEES", targetID),
		),
	}})
}
