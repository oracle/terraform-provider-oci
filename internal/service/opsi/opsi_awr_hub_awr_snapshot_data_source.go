// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiAwrHubAwrSnapshotDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpsiAwrHubAwrSnapshot,
		Schema: map[string]*schema.Schema{
			"awr_hub_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"awr_source_database_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"awr_source_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"snapshot_identifier": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"time_db_startup": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_snapshot_begin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_snapshot_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_opsi_awr_hub_awr_snapshot", "oci_opsi_awr_hub_awr_snapshots"),
	}
}

func readSingularOpsiAwrHubAwrSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubAwrSnapshotDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiAwrHubAwrSnapshotDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListAwrSnapshotsResponse
}

func (s *OpsiAwrHubAwrSnapshotDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiAwrHubAwrSnapshotDataSourceCrud) Get() error {
	request := oci_opsi.ListAwrSnapshotsRequest{}

	if awrHubId, ok := s.D.GetOkExists("awr_hub_id"); ok {
		tmp := awrHubId.(string)
		request.AwrHubId = &tmp
	}

	if awrSourceDatabaseIdentifier, ok := s.D.GetOkExists("awr_source_database_identifier"); ok {
		tmp := awrSourceDatabaseIdentifier.(string)
		request.AwrSourceDatabaseIdentifier = &tmp
	}

	if timeGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeLessThanOrEqualTo, ok := s.D.GetOkExists("time_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListAwrSnapshots(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiAwrHubAwrSnapshotDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiAwrHubAwrSnapshotDataSource-", OpsiAwrHubAwrSnapshotDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AwrSnapshotSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
