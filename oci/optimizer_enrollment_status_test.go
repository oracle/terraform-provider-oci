// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EnrollmentStatusResourceConfig = EnrollmentStatusResourceDependencies +
		generateDataSourceFromRepresentationMap("oci_optimizer_enrollment_statuses", "test_enrollment_statuses", Required, Create, enrollmentStatusDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_optimizer_enrollment_status", "test_enrollment_status", Optional, Update, enrollmentStatusRepresentation)

	enrollmentStatusSingularDataSourceRepresentation = map[string]interface{}{
		"enrollment_status_id": Representation{repType: Required, create: `${data.oci_optimizer_enrollment_statuses.test_enrollment_statuses.enrollment_status_collection.0.items.0.id}`},
	}

	enrollmentStatusDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"status":         Representation{repType: Optional, create: `INACTIVE`, update: `ACTIVE`},
	}

	enrollmentStatusRepresentation = map[string]interface{}{
		"enrollment_status_id": Representation{repType: Required, create: `${data.oci_optimizer_enrollment_statuses.test_enrollment_statuses.enrollment_status_collection.0.items.0.id}`},
		"status":               Representation{repType: Required, create: `INACTIVE`, update: `ACTIVE`},
	}

	EnrollmentStatusResourceDependencies = ""
)

func TestOptimizerEnrollmentStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerEnrollmentStatusResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_enrollment_status.test_enrollment_status"
	datasourceName := "data.oci_optimizer_enrollment_statuses.test_enrollment_statuses"
	singularDatasourceName := "data.oci_optimizer_enrollment_status.test_enrollment_status"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + EnrollmentStatusResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_optimizer_enrollment_statuses", "test_enrollment_statuses", Required, Create, enrollmentStatusDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_optimizer_enrollment_status", "test_enrollment_status", Required, Create, enrollmentStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "enrollment_status_id"),
					resource.TestCheckResourceAttr(resourceName, "status", "INACTIVE"),

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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + EnrollmentStatusResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_optimizer_enrollment_statuses", "test_enrollment_statuses", Required, Create, enrollmentStatusDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_optimizer_enrollment_status", "test_enrollment_status", Optional, Update, enrollmentStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "enrollment_status_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "status", "ACTIVE"),

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
					generateDataSourceFromRepresentationMap("oci_optimizer_enrollment_statuses", "test_enrollment_statuses", Optional, Update, enrollmentStatusDataSourceRepresentation) +
					compartmentIdVariableStr + EnrollmentStatusResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_enrollment_status", "test_enrollment_status", Optional, Update, enrollmentStatusRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "enrollment_status_collection.0.items.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "enrollment_status_collection.0.items.0.status", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "enrollment_status_collection.0.items.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "enrollment_status_collection.0.items.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "enrollment_status_collection.0.items.0.time_updated"),

					resource.TestCheckResourceAttr(datasourceName, "enrollment_status_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "enrollment_status_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_optimizer_enrollment_status", "test_enrollment_status", Required, Create, enrollmentStatusSingularDataSourceRepresentation) +
					compartmentIdVariableStr + EnrollmentStatusResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "enrollment_status_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "status", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + EnrollmentStatusResourceConfig,
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
