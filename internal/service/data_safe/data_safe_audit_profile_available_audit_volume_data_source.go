// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditProfileAvailableAuditVolumeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeAuditProfileAvailableAuditVolume,
		Schema: map[string]*schema.Schema{
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
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_safe_audit_profile_available_audit_volume", "oci_data_safe_audit_profile_available_audit_volumes"),
	}
}

func readSingularDataSafeAuditProfileAvailableAuditVolume(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileAvailableAuditVolumeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileAvailableAuditVolumeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAvailableAuditVolumesResponse
}

func (s *DataSafeAuditProfileAvailableAuditVolumeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileAvailableAuditVolumeDataSourceCrud) Get() error {
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
	return nil
}

func (s *DataSafeAuditProfileAvailableAuditVolumeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfileAvailableAuditVolumeDataSource-", DataSafeAuditProfileAvailableAuditVolumeDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableAuditVolumeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
