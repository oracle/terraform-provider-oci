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
	OsManagementHubErrataSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: []string{`ELBA-2024-12244`}},
	}

	OsManagementHubErrataDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"advisory_severity":     acctest.Representation{RepType: acctest.Optional, Create: []string{`advisorySeverity`}},
		"classification_type":   acctest.Representation{RepType: acctest.Optional, Create: []string{`BUGFIX`}},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`ELBA-2024-12244`}},
		"name_contains":         acctest.Representation{RepType: acctest.Required, Create: `ELBA-2024-12244`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_9`},
		"time_issue_date_end":   acctest.Representation{RepType: acctest.Optional, Create: `2024-04-01T00:00:00.000Z`},
		"time_issue_date_start": acctest.Representation{RepType: acctest.Optional, Create: `2024-03-20T00:00:00.000Z`},
	}

	OsManagementHubErrataResourceConfig = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubErrataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubErrataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_errata.test_errata"
	singularDatasourceName := "data.oci_os_management_hub_errata.test_errata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_errata", "test_errata", acctest.Required, acctest.Create, OsManagementHubErrataDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubErrataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "ELBA-2024-12244"),

				resource.TestCheckResourceAttrSet(datasourceName, "erratum_collection.#"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_errata", "test_errata", acctest.Optional, acctest.Create, OsManagementHubErrataDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubErrataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "advisory_severity.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "classification_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "ELBA-2024-12244"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_9"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_issue_date_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_issue_date_start"),

				resource.TestCheckResourceAttrSet(datasourceName, "erratum_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_errata", "test_errata", acctest.Required, acctest.Create, OsManagementHubErrataSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubErrataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "erratum_collection.#"),
			),
		},
	})
}
