// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	jobExecutionsStatusSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"end_time":            acctest.Representation{RepType: acctest.Required, Create: `${var.end_time}`},
		"start_time":          acctest.Representation{RepType: acctest.Required, Create: `${var.start_time}`},
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
	}

	jobExecutionsStatusDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"end_time":            acctest.Representation{RepType: acctest.Required, Create: `${var.end_time}`},
		"start_time":          acctest.Representation{RepType: acctest.Required, Create: `${var.start_time}`},
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
	}

	JobExecutionsStatusResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementJobExecutionsStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementJobExecutionsStatusResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	endTime := time.Now().UTC()
	startTime := endTime.Add(-12 * time.Hour)
	endTimeVariableStr := fmt.Sprintf("variable \"end_time\" { default = \"%s\" }\n", endTime.Format("2006-01-02T15:04:05.000Z"))
	startTimeVariableStr := fmt.Sprintf("variable \"start_time\" { default = \"%s\" }\n", startTime.Format("2006-01-02T15:04:05.000Z"))

	datasourceName := "data.oci_database_management_job_executions_statuses.test_job_executions_statuses"
	singularDatasourceName := "data.oci_database_management_job_executions_status.test_job_executions_status"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_job_executions_statuses", "test_job_executions_statuses", acctest.Required, acctest.Create, jobExecutionsStatusDataSourceRepresentation) +
					compartmentIdVariableStr + JobExecutionsStatusResourceConfig + endTimeVariableStr + startTimeVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "end_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time"),

					resource.TestCheckResourceAttrSet(datasourceName, "job_executions_status_summary_collection.#"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_job_executions_status", "test_job_executions_status", acctest.Required, acctest.Create, jobExecutionsStatusSingularDataSourceRepresentation) +
					compartmentIdVariableStr + JobExecutionsStatusResourceConfig + endTimeVariableStr + startTimeVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "end_time"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "start_time"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				),
			},
		},
	})
}
