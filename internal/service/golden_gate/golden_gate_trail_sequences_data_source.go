// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateTrailSequencesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateTrailSequences,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trail_file_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trail_sequence_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trail_sequence_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sequence_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_bytes": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"time_last_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_last_fetched": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readGoldenGateTrailSequences(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateTrailSequencesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateTrailSequencesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListTrailSequencesResponse
}

func (s *GoldenGateTrailSequencesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateTrailSequencesDataSourceCrud) Get() error {
	request := oci_golden_gate.ListTrailSequencesRequest{}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if trailFileId, ok := s.D.GetOkExists("trail_file_id"); ok {
		tmp := trailFileId.(string)
		request.TrailFileId = &tmp
	}

	if trailSequenceId, ok := s.D.GetOkExists("trail_sequence_id"); ok {
		tmp := trailSequenceId.(string)
		request.TrailSequenceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListTrailSequences(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTrailSequences(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateTrailSequencesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateTrailSequencesDataSource-", GoldenGateTrailSequencesDataSource(), s.D))
	resources := []map[string]interface{}{}
	trailSequence := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TrailSequenceSummaryToMap(item))
	}
	trailSequence["items"] = items

	if s.Res.TimeLastFetched != nil {
		trailSequence["time_last_fetched"] = s.Res.TimeLastFetched.String()
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateTrailSequencesDataSource().Schema["trail_sequence_collection"].Elem.(*schema.Resource).Schema)
		trailSequence["items"] = items
	}

	resources = append(resources, trailSequence)
	if err := s.D.Set("trail_sequence_collection", resources); err != nil {
		return err
	}

	return nil
}

func TrailSequenceSummaryToMap(obj oci_golden_gate.TrailSequenceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.SequenceId != nil {
		result["sequence_id"] = string(*obj.SequenceId)
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = float32(*obj.SizeInBytes)
	}

	if obj.TimeLastUpdated != nil {
		result["time_last_updated"] = obj.TimeLastUpdated.String()
	}

	return result
}
