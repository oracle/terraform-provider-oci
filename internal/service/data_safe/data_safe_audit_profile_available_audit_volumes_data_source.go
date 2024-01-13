// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditProfileAvailableAuditVolumesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAuditProfileAvailableAuditVolumes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"audit_profile_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"month_in_consideration_greater_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"month_in_consideration_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trail_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available_audit_volume_collection": {
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
												"audit_profile_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"month_in_consideration": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"trail_location": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"volume": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"audit_trail_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"database_unique_name": {
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

func readDataSafeAuditProfileAvailableAuditVolumes(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileAvailableAuditVolumesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileAvailableAuditVolumesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAvailableAuditVolumesResponse
}

func (s *DataSafeAuditProfileAvailableAuditVolumesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileAvailableAuditVolumesDataSourceCrud) Get() error {
	request := oci_data_safe.ListAvailableAuditVolumesRequest{}

	if auditProfileId, ok := s.D.GetOkExists("audit_profile_id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
	}

	if monthInConsiderationGreaterThan, ok := s.D.GetOkExists("month_in_consideration_greater_than"); ok {
		tmp, err := time.Parse(time.RFC3339, monthInConsiderationGreaterThan.(string))
		if err != nil {
			return err
		}
		request.MonthInConsiderationGreaterThan = &oci_common.SDKTime{Time: tmp}
	}

	if monthInConsiderationLessThan, ok := s.D.GetOkExists("month_in_consideration_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, monthInConsiderationLessThan.(string))
		if err != nil {
			return err
		}
		request.MonthInConsiderationLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if trailLocation, ok := s.D.GetOkExists("trail_location"); ok {
		tmp := trailLocation.(string)
		request.TrailLocation = &tmp
	}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAvailableAuditVolumes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAvailableAuditVolumes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAuditProfileAvailableAuditVolumesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfileAvailableAuditVolumesDataSource-", DataSafeAuditProfileAvailableAuditVolumesDataSource(), s.D))
	resources := []map[string]interface{}{}
	auditProfileAvailableAuditVolume := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableAuditVolumesSummaryToMap(item))
	}
	auditProfileAvailableAuditVolume["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAuditProfileAvailableAuditVolumesDataSource().Schema["available_audit_volume_collection"].Elem.(*schema.Resource).Schema)
		auditProfileAvailableAuditVolume["items"] = items
	}

	resources = append(resources, auditProfileAvailableAuditVolume)
	if err := s.D.Set("available_audit_volume_collection", resources); err != nil {
		return err
	}

	return nil
}

func AvailableAuditVolumesSummaryToMap(obj oci_data_safe.AvailableAuditVolumeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditProfileId != nil {
		result["audit_profile_id"] = string(*obj.AuditProfileId)
	}

	if obj.MonthInConsideration != nil {
		result["month_in_consideration"] = obj.MonthInConsideration.String()
	}

	if obj.TrailLocation != nil {
		result["trail_location"] = string(*obj.TrailLocation)
	}

	if obj.Volume != nil {
		result["volume"] = strconv.FormatInt(*obj.Volume, 10)
	}

	return result
}

func AvailableAuditVolumeSummaryToMap(obj oci_data_safe.AvailableAuditVolumeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditProfileId != nil {
		result["audit_profile_id"] = string(*obj.AuditProfileId)
	}

	if obj.AuditTrailId != nil {
		result["audit_trail_id"] = string(*obj.AuditTrailId)
	}

	if obj.DatabaseUniqueName != nil {
		result["database_unique_name"] = string(*obj.DatabaseUniqueName)
	}

	if obj.MonthInConsideration != nil {
		result["month_in_consideration"] = obj.MonthInConsideration.String()
	}

	if obj.TrailLocation != nil {
		result["trail_location"] = string(*obj.TrailLocation)
	}

	if obj.Volume != nil {
		result["volume"] = strconv.FormatInt(*obj.Volume, 10)
	}

	return result
}
