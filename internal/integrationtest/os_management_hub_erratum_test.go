// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	OsManagementHubErratumSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `ELBA-2024-12244`},
	}

	OsManagementHubErratumDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"advisory_severity":     acctest.Representation{RepType: acctest.Optional, Create: []string{`advisorySeverity`}},
		"advisory_type":         acctest.Representation{RepType: acctest.Optional, Create: []string{`advisoryType`}},
		"classification_type":   acctest.Representation{RepType: acctest.Optional, Create: []string{`classificationType`}},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`ELBA-2024-12244`}},
		"name_contains":         acctest.Representation{RepType: acctest.Optional, Create: `nameContains`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_9`},
		"time_issue_date_end":   acctest.Representation{RepType: acctest.Optional, Create: `timeIssueDateEnd`},
		"time_issue_date_start": acctest.Representation{RepType: acctest.Optional, Create: `timeIssueDateStart`},
	}

	OsManagementHubErratumResourceConfig = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubErratumResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubErratumResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//datasourceName := "data.oci_os_management_hub_errata.test_errata"
	singularDatasourceName := "data.oci_os_management_hub_erratum.test_erratum"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_erratum", "test_erratum", acctest.Required, acctest.Create, OsManagementHubErratumSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubErratumResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "classification_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "from"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "references"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "solution"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "synopsis"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_issued"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
