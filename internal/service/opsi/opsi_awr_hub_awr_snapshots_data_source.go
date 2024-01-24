// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"strconv"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiAwrHubAwrSnapshotsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiAwrHubAwrSnapshots,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"awr_snapshot_collection": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readOpsiAwrHubAwrSnapshots(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubAwrSnapshotsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiAwrHubAwrSnapshotsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListAwrSnapshotsResponse
}

func (s *OpsiAwrHubAwrSnapshotsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiAwrHubAwrSnapshotsDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAwrSnapshots(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiAwrHubAwrSnapshotsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiAwrHubAwrSnapshotsDataSource-", OpsiAwrHubAwrSnapshotsDataSource(), s.D))
	resources := []map[string]interface{}{}
	awrHubAwrSnapshot := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AwrSnapshotsSummaryToMap(item))
	}
	awrHubAwrSnapshot["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiAwrHubAwrSnapshotsDataSource().Schema["awr_snapshot_collection"].Elem.(*schema.Resource).Schema)
		awrHubAwrSnapshot["items"] = items
	}

	resources = append(resources, awrHubAwrSnapshot)
	if err := s.D.Set("awr_snapshot_collection", resources); err != nil {
		return err
	}

	return nil
}

func AwrSnapshotsSummaryToMap(obj oci_opsi.AwrSnapshotSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AwrSourceDatabaseId != nil {
		result["awr_source_database_id"] = string(*obj.AwrSourceDatabaseId)
	}

	if obj.ErrorCount != nil {
		result["error_count"] = strconv.FormatInt(*obj.ErrorCount, 10)
	}

	if obj.InstanceNumber != nil {
		result["instance_number"] = int(*obj.InstanceNumber)
	}

	if obj.SnapshotIdentifier != nil {
		result["snapshot_identifier"] = int(*obj.SnapshotIdentifier)
	}

	if obj.TimeDbStartup != nil {
		result["time_db_startup"] = obj.TimeDbStartup.String()
	}

	if obj.TimeSnapshotBegin != nil {
		result["time_snapshot_begin"] = obj.TimeSnapshotBegin.String()
	}

	if obj.TimeSnapshotEnd != nil {
		result["time_snapshot_end"] = obj.TimeSnapshotEnd.String()
	}

	return result
}

func AwrSnapshotSummaryToMap(obj oci_opsi.AwrSnapshotSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AwrSourceDatabaseId != nil {
		result["awr_source_database_id"] = string(*obj.AwrSourceDatabaseId)
	}

	if obj.ErrorCount != nil {
		result["error_count"] = strconv.FormatInt(*obj.ErrorCount, 10)
	}

	if obj.InstanceNumber != nil {
		result["instance_number"] = int(*obj.InstanceNumber)
	}

	if obj.SnapshotIdentifier != nil {
		result["snapshot_identifier"] = int(*obj.SnapshotIdentifier)
	}

	if obj.TimeDbStartup != nil {
		result["time_db_startup"] = obj.TimeDbStartup.String()
	}

	if obj.TimeSnapshotBegin != nil {
		result["time_snapshot_begin"] = obj.TimeSnapshotBegin.String()
	}

	if obj.TimeSnapshotEnd != nil {
		result["time_snapshot_end"] = obj.TimeSnapshotEnd.String()
	}

	return result
}
