// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"
)

func OpsiAwrHubAwrSourcesSummaryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOpsiAwrHubAwrSourcesSummary,
		Schema: map[string]*schema.Schema{
			"awr_hub_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
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
						"awr_hub_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"awr_source_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hours_since_last_import": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"max_snapshot_identifier": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"min_snapshot_identifier": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snapshots_uploaded": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"time_first_snapshot_generated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_snapshot_generated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularOpsiAwrHubAwrSourcesSummary(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubAwrSourcesSummaryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiAwrHubAwrSourcesSummaryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.SummarizeAwrSourcesSummariesResponse
}

func (s *OpsiAwrHubAwrSourcesSummaryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiAwrHubAwrSourcesSummaryDataSourceCrud) Get() error {
	request := oci_opsi.SummarizeAwrSourcesSummariesRequest{}

	if awrHubId, ok := s.D.GetOkExists("awr_hub_id"); ok {
		tmp := awrHubId.(string)
		request.AwrHubId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.SummarizeAwrSourcesSummaries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiAwrHubAwrSourcesSummaryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiAwrHubAwrSourcesSummaryDataSource-", OpsiAwrHubAwrSourcesSummaryDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AwrSourceSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AwrSourceSummaryToMap(obj oci_opsi.AwrSourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AwrHubId != nil {
		result["awr_hub_id"] = string(*obj.AwrHubId)
	}

	if obj.AwrSourceDatabaseId != nil {
		result["awr_source_database_id"] = string(*obj.AwrSourceDatabaseId)
	}

	if obj.HoursSinceLastImport != nil {
		result["hours_since_last_import"] = float64(*obj.HoursSinceLastImport)
	}

	if obj.MaxSnapshotIdentifier != nil {
		result["max_snapshot_identifier"] = float32(*obj.MaxSnapshotIdentifier)
	}

	if obj.MinSnapshotIdentifier != nil {
		result["min_snapshot_identifier"] = float32(*obj.MinSnapshotIdentifier)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.SnapshotsUploaded != nil {
		result["snapshots_uploaded"] = float32(*obj.SnapshotsUploaded)
	}

	if obj.TimeFirstSnapshotGenerated != nil {
		result["time_first_snapshot_generated"] = obj.TimeFirstSnapshotGenerated.String()
	}

	if obj.TimeLastSnapshotGenerated != nil {
		result["time_last_snapshot_generated"] = obj.TimeLastSnapshotGenerated.String()
	}

	return result
}
