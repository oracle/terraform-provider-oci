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

func DataSafeAuditProfileCollectedAuditVolumeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeAuditProfileCollectedAuditVolume,
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_safe_audit_profile_collected_audit_volume", "oci_data_safe_audit_profile_collected_audit_volumes"),
	}
}

func readSingularDataSafeAuditProfileCollectedAuditVolume(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileCollectedAuditVolumeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileCollectedAuditVolumeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCollectedAuditVolumesResponse
}

func (s *DataSafeAuditProfileCollectedAuditVolumeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileCollectedAuditVolumeDataSourceCrud) Get() error {
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
	return nil
}

func (s *DataSafeAuditProfileCollectedAuditVolumeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfileCollectedAuditVolumeDataSource-", DataSafeAuditProfileCollectedAuditVolumeDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CollectedAuditVolumeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
