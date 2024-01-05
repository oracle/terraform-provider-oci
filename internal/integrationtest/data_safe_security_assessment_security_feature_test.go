// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSecurityAssessmentSecurityFeatureDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                         acctest.Representation{RepType: acctest.Required, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":            acctest.Representation{RepType: acctest.Required, Create: `true`},
		"target_id":                            acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"targets_with_column_encryption":       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_database_vault":          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_external_authentication": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_fine_grained_audit":      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_global_authentication":   acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_network_encryption":      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_password_authentication": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_privilege_analysis":      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_tablespace_encryption":   acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_traditional_audit":       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		"targets_with_unified_audit":           acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
	}

	DataSafeSecurityAssessmentSecurityFeatureResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentSecurityFeatureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentSecurityFeatureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_security_assessment_security_features.test_security_assessment_security_features"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_security_features", "test_security_assessment_security_features", acctest.Required, acctest.Create, DataSafeSecurityAssessmentSecurityFeatureDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSecurityAssessmentSecurityFeatureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.column_encryption"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.database_vault"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.external_authentication"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.fine_grained_audit"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.global_authentication"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.network_encryption"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.password_authentication"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.privilege_analysis"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.tablespace_encryption"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.traditional_audit"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.0.items.0.unified_audit"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "security_feature_collection.0.items.#", "1"),
			),
		},
	})
}
