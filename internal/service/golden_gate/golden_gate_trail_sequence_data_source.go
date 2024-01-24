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

func GoldenGateTrailSequenceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularGoldenGateTrailSequence,
		Schema: map[string]*schema.Schema{
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
			// Computed
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_golden_gate_trail_sequence", "oci_golden_gate_trail_sequences"),
	}
}

func readSingularGoldenGateTrailSequence(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateTrailSequenceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateTrailSequenceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListTrailSequencesResponse
}

func (s *GoldenGateTrailSequenceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateTrailSequenceDataSourceCrud) Get() error {
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
	return nil
}

func (s *GoldenGateTrailSequenceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateTrailSequenceDataSource-", GoldenGateTrailSequenceDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TrailSequenceSummaryToMap(item))
	}
	s.D.Set("items", items)

	if s.Res.TimeLastFetched != nil {
		s.D.Set("time_last_fetched", s.Res.TimeLastFetched.String())
	}

	return nil
}
