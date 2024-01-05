// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeAuditProfileCollectedAuditVolumesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAuditProfileCollectedAuditVolumes,
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
			"work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"collected_audit_volume_collection": {
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
												"archived_volume": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"audit_profile_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"month_in_consideration": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"online_volume": {
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

func readDataSafeAuditProfileCollectedAuditVolumes(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileCollectedAuditVolumesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileCollectedAuditVolumesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCollectedAuditVolumesResponse
}

func (s *DataSafeAuditProfileCollectedAuditVolumesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileCollectedAuditVolumesDataSourceCrud) Get() error {
	request := oci_data_safe.ListCollectedAuditVolumesRequest{}

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

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListCollectedAuditVolumes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCollectedAuditVolumes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAuditProfileCollectedAuditVolumesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfileCollectedAuditVolumesDataSource-", DataSafeAuditProfileCollectedAuditVolumesDataSource(), s.D))
	resources := []map[string]interface{}{}
	auditProfileCollectedAuditVolume := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CollectedAuditVolumesSummaryToMap(item))
	}
	auditProfileCollectedAuditVolume["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAuditProfileCollectedAuditVolumesDataSource().Schema["collected_audit_volume_collection"].Elem.(*schema.Resource).Schema)
		auditProfileCollectedAuditVolume["items"] = items
	}

	resources = append(resources, auditProfileCollectedAuditVolume)
	if err := s.D.Set("collected_audit_volume_collection", resources); err != nil {
		return err
	}

	return nil
}

func CollectedAuditVolumesSummaryToMap(obj oci_data_safe.CollectedAuditVolumeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArchivedVolume != nil {
		result["archived_volume"] = strconv.FormatInt(*obj.ArchivedVolume, 10)
	}

	if obj.AuditProfileId != nil {
		result["audit_profile_id"] = string(*obj.AuditProfileId)
	}

	if obj.MonthInConsideration != nil {
		result["month_in_consideration"] = obj.MonthInConsideration.String()
	}

	if obj.OnlineVolume != nil {
		result["online_volume"] = strconv.FormatInt(*obj.OnlineVolume, 10)
	}

	return result
}

func CollectedAuditVolumeSummaryToMap(obj oci_data_safe.CollectedAuditVolumeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArchivedVolume != nil {
		result["archived_volume"] = strconv.FormatInt(*obj.ArchivedVolume, 10)
	}

	if obj.AuditProfileId != nil {
		result["audit_profile_id"] = string(*obj.AuditProfileId)
	}

	if obj.MonthInConsideration != nil {
		result["month_in_consideration"] = obj.MonthInConsideration.String()
	}

	if obj.OnlineVolume != nil {
		result["online_volume"] = strconv.FormatInt(*obj.OnlineVolume, 10)
	}

	return result
}
