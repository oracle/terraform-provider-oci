// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeDataSafeSensitiveDataModelSensitiveSchemaDataSourceRepresentation = map[string]interface{}{
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"schema_name":             acctest.Representation{RepType: acctest.Optional, Create: []string{`ADMIN`}},
	}

	SensitiveSchemasSDMDependency = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)

	DataSafeSensitiveDataModelSensitiveSchemaResourceConfig = SensitiveSchemasSDMDependency
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveDataModelSensitiveSchemaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveDataModelSensitiveSchemaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_sensitive_data_model_sensitive_schemas.test_sensitive_data_model_sensitive_schemas"

	acctest.SaveConfigContent("", "", "", t)

	var resId string
	resourceName := "oci_data_safe_sensitive_data_models_sensitive_column.test_sensitive_data_models_sensitive_column"
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveDataModelSensitiveSchemaResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_sensitive_column", "test_sensitive_data_models_sensitive_column", acctest.Required, acctest.Create, sensitiveDataModelsSensitiveColumnRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column_name", "FIRST_NAME"),
				resource.TestCheckResourceAttr(resourceName, "object", "EMPLOYEES"),
				resource.TestCheckResourceAttr(resourceName, "schema_name", "ADMIN"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "key")
					fmt.Printf(resId)
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + targetIdVariableStr + DataSafeSensitiveDataModelSensitiveSchemaResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_data_model_sensitive_schemas", "test_sensitive_data_model_sensitive_schemas", acctest.Required, acctest.Create, DataSafeDataSafeSensitiveDataModelSensitiveSchemaDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_schema_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_schema_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_schema_collection.0.items.0.schema_name", "ADMIN"),
			),
		},
	})
}
