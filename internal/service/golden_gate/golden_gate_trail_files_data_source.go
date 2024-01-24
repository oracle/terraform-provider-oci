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

func GoldenGateTrailFilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateTrailFiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trail_file_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trail_file_collection": {
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
									"consumers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_sequence_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"min_sequence_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"number_of_sequences": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"producer": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_bytes": {
										Type:     schema.TypeFloat, // keep TypeFloat instead of the computed TypeString
										Computed: true,
									},
									"time_last_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"trail_file_id": {
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

func readGoldenGateTrailFiles(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateTrailFilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateTrailFilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListTrailFilesResponse
}

func (s *GoldenGateTrailFilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateTrailFilesDataSourceCrud) Get() error {
	request := oci_golden_gate.ListTrailFilesRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListTrailFiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTrailFiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateTrailFilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateTrailFilesDataSource-", GoldenGateTrailFilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	trailFile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TrailFileSummaryToMap(item))
	}
	trailFile["items"] = items

	if s.Res.TimeLastFetched != nil {
		trailFile["time_last_fetched"] = s.Res.TimeLastFetched.String()
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateTrailFilesDataSource().Schema["trail_file_collection"].Elem.(*schema.Resource).Schema)
		trailFile["items"] = items
	}

	resources = append(resources, trailFile)
	if err := s.D.Set("trail_file_collection", resources); err != nil {
		return err
	}

	return nil
}

func TrailFileSummaryToMap(obj oci_golden_gate.TrailFileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["consumers"] = obj.Consumers
	result["consumers"] = obj.Consumers

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.MaxSequenceNumber != nil {
		result["max_sequence_number"] = string(*obj.MaxSequenceNumber)
	}

	if obj.MinSequenceNumber != nil {
		result["min_sequence_number"] = string(*obj.MinSequenceNumber)
	}

	if obj.NumberOfSequences != nil {
		result["number_of_sequences"] = int(*obj.NumberOfSequences)
	}

	if obj.Producer != nil {
		result["producer"] = string(*obj.Producer)
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = float32(*obj.SizeInBytes)
	}

	if obj.TimeLastUpdated != nil {
		result["time_last_updated"] = obj.TimeLastUpdated.String()
	}

	if obj.TrailFileId != nil {
		result["trail_file_id"] = string(*obj.TrailFileId)
	}

	return result
}
