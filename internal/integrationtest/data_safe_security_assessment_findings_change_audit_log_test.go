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
	DataSafeSecurityAssessmentFindingsChangeAuditLogDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.security_assessment_id}`},
		"finding_key":                               acctest.Representation{RepType: acctest.Optional, Create: `findingKey`},
		"finding_title":                             acctest.Representation{RepType: acctest.Optional, Create: `findingTitle`},
		"is_risk_deferred":                          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"modified_by":                               acctest.Representation{RepType: acctest.Optional, Create: `modifiedBy`},
		"severity":                                  acctest.Representation{RepType: acctest.Optional, Create: `HIGH`},
		"time_updated_greater_than_or_equal_to":     acctest.Representation{RepType: acctest.Optional, Create: `timeUpdatedGreaterThanOrEqualTo`},
		"time_updated_less_than":                    acctest.Representation{RepType: acctest.Optional, Create: `timeUpdatedLessThan`},
		"time_valid_until_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeValidUntilGreaterThanOrEqualTo`},
		"time_valid_until_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `timeValidUntilLessThan`},
	}

	DataSafeSecurityAssessmentFindingsChangeAuditLogResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentFindingsChangeAuditLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentFindingsChangeAuditLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityAssessmentId := utils.GetEnvSettingWithBlankDefault("security_assessment_ocid")
	securityAssessmentIdVariableStr := fmt.Sprintf("variable \"security_assessment_id\" { default = \"%s\" }\n", securityAssessmentId)

	datasourceName := "data.oci_data_safe_security_assessment_findings_change_audit_logs.test_security_assessment_findings_change_audit_logs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_findings_change_audit_logs", "test_security_assessment_findings_change_audit_logs", acctest.Required, acctest.Create, DataSafeSecurityAssessmentFindingsChangeAuditLogDataSourceRepresentation) +
				compartmentIdVariableStr + securityAssessmentIdVariableStr + DataSafeSecurityAssessmentFindingsChangeAuditLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.0.finding_key"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.0.finding_title"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.0.is_risk_deferred"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.0.modified_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.0.assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.0.severity"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings_change_audit_log_collection.0.items.#"),
			),
		},
	})
}
