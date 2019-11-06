// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

func TestStreamingStreamArchiverScanario_basic(t *testing.T) {
	httpreplay.SetScenario("TestStreamingStreamArchiverScanario_basic")
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
			// verify create with state = STOPPED
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
			//Delete before Recreate
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies,
			},
			// verify Recreate after Delete, should only create an entry in state file no change to resource on cloud
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
			//Delete before Recreate with state = RUNNING
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies,
			},
			// verify create with state = RUNNING
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Required, Create,
						getUpdatedRepresentationCopy("state", Representation{repType: Required, create: "running"}, streamArchiverRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "LATEST"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			//Delete before Recreate
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies,
			},
			// verify Recreate after Delete, should only create an entry in state file no change to resource on cloud
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Required, Create,
						getUpdatedRepresentationCopy("state", Representation{repType: Required, create: "running"}, streamArchiverRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "LATEST"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
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
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify start streaming archiver
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Update,
						getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "running"}, streamArchiverRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "11"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "TRIM_HORIZON"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify stop streaming archiver
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Update, streamArchiverRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "11"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "TRIM_HORIZON"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify start streaming archiver with an update to start_position
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Update,
						getUpdatedRepresentationCopy("start_position", Representation{repType: Optional, create: "LATEST"}, getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "running"}, streamArchiverRepresentation))),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "11"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "LATEST"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// try to update to start_position while streaming archiver is running and should expect error
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Required, Create,
						getUpdatedRepresentationCopy("start_position", Representation{repType: Required, create: "TRIM_HORIZON"}, getUpdatedRepresentationCopy("state", Representation{repType: Optional, create: "running"}, streamArchiverRepresentation))),
				ExpectError: regexp.MustCompile("Conflict. lifecycleState."),
			},
			// verify stop streaming archiver with an update to start_position
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_archiver", "test_stream_archiver", Optional, Update,
						getUpdatedRepresentationCopy("start_position", Representation{repType: Optional, create: "TRIM_HORIZON"}, streamArchiverRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_size_in_mbs", "11"),
					resource.TestCheckResourceAttr(resourceName, "batch_rollover_time_in_seconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "StreamingArchiverTestBucket"),
					resource.TestCheckResourceAttr(resourceName, "start_position", "TRIM_HORIZON"),
					resource.TestCheckResourceAttrSet(resourceName, "stream_id"),
					resource.TestCheckResourceAttr(resourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(resourceName, "state", "STOPPED"),

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
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "use_existing_bucket", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "STOPPED"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + StreamArchiverResourceConfig,
			},
		},
	})
}
