// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strconv"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAuditProfiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"audit_collected_volume_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"audit_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_override_global_retention_setting": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_paid_usage_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"audit_profile_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeAuditProfileResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeAuditProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAuditProfilesResponse
}

func (s *DataSafeAuditProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfilesDataSourceCrud) Get() error {
	request := oci_data_safe.ListAuditProfilesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAuditProfilesAccessLevelEnum(accessLevel.(string))
	}

	if auditCollectedVolumeGreaterThanOrEqualTo, ok := s.D.GetOkExists("audit_collected_volume_greater_than_or_equal_to"); ok {
		tmp := auditCollectedVolumeGreaterThanOrEqualTo.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert auditCollectedVolumeGreaterThanOrEqualTo string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.AuditCollectedVolumeGreaterThanOrEqualTo = &tmpInt64
	}

	if auditProfileId, ok := s.D.GetOkExists("id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isOverrideGlobalRetentionSetting, ok := s.D.GetOkExists("is_override_global_retention_setting"); ok {
		tmp := isOverrideGlobalRetentionSetting.(bool)
		request.IsOverrideGlobalRetentionSetting = &tmp
	}

	if isPaidUsageEnabled, ok := s.D.GetOkExists("is_paid_usage_enabled"); ok {
		tmp := isPaidUsageEnabled.(bool)
		request.IsPaidUsageEnabled = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListAuditProfilesLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAuditProfiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAuditProfiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAuditProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfilesDataSource-", DataSafeAuditProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	auditProfile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AuditProfileSummaryToMap(item))
	}
	auditProfile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAuditProfilesDataSource().Schema["audit_profile_collection"].Elem.(*schema.Resource).Schema)
		auditProfile["items"] = items
	}

	resources = append(resources, auditProfile)
	if err := s.D.Set("audit_profile_collection", resources); err != nil {
		return err
	}

	return nil
}
