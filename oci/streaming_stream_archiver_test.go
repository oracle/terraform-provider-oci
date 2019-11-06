// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	StreamArchiverRequiredOnlyResource = StreamArchiverResourceDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Required, Create, streamArchiverRepresentation)

	StreamArchiverResourceConfig = StreamArchiverResourceDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Update, streamArchiverRepresentation)

	streamArchiverSingularDataSourceRepresentation = map[string]interface{}{
		"stream_id": Representation{repType: Required, create: `${oci_streaming_stream.test_stream.id}`},
	}

	streamArchiverRepresentation = map[string]interface{}{
		"batch_rollover_size_in_mbs":     Representation{repType: Required, create: `10`, update: `11`},
		"batch_rollover_time_in_seconds": Representation{repType: Required, create: `10`, update: `11`},
		"bucket":                         Representation{repType: Required, create: `StreamingArchiverTestBucket`},
		"start_position":                 Representation{repType: Required, create: `LATEST`, update: `TRIM_HORIZON`},
		"stream_id":                      Representation{repType: Required, create: `${oci_streaming_stream.test_stream.id}`},
		"use_existing_bucket":            Representation{repType: Required, create: `true`},
		"state":                          Representation{repType: Required, create: `stopped`},
	}

	StreamArchiverResourceDependencies = generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation)
)

func TestStreamingStreamArchiverResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStreamingStreamArchiverResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_streaming_stream_archiver.test_stream_archiver"

	singularDatasourceName := "data.oci_streaming_stream_archiver.test_stream_archiver"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Required, Create, streamArchiverRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "LATEST"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Create, streamArchiverRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "LATEST"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Update, streamArchiverRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "11"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "TRIM_HORIZON"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Required, Create, streamArchiverSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StreamArchiverResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "batch_rollover_size_in_mbs", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "batch_rollover_time_in_seconds", "11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(singularDatasourceName, "error.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "start_position", "TRIM_HORIZON"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "STOPPED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "use_existing_bucket", "true"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportStateIdFunc:       getStreamingArchiverCompositeIdForImport(resourceName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func getStreamingArchiverCompositeIdForImport(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("archiver/%s", rs.Primary.Attributes["stream_id"]), nil
	}
}
