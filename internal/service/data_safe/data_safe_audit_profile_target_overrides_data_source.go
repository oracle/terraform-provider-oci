// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditProfileTargetOverridesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAuditProfileTargetOverrides,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"audit_profile_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_override_collection": {
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
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"is_paid_usage_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"offline_months": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"offline_months_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"online_months": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"online_months_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"paid_usage_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"target_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets_conforming_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"targets_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"targets_overriding_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"targets_overriding_offline_months_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"targets_overriding_online_months_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"targets_overriding_paid_usage_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeAuditProfileTargetOverrides(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileTargetOverridesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileTargetOverridesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListTargetOverridesResponse
}

func (s *DataSafeAuditProfileTargetOverridesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileTargetOverridesDataSourceCrud) Get() error {
	request := oci_data_safe.ListTargetOverridesRequest{}

	if auditProfileId, ok := s.D.GetOkExists("audit_profile_id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListTargetOverrides(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTargetOverrides(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAuditProfileTargetOverridesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfileTargetOverridesDataSource-", DataSafeAuditProfileTargetOverridesDataSource(), s.D))
	resources := []map[string]interface{}{}
	auditProfileTargetOverride := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TargetOverrideSummaryToMap(item))
	}
	auditProfileTargetOverride["items"] = items

	if s.Res.TargetsConformingCount != nil {
		auditProfileTargetOverride["targets_conforming_count"] = *s.Res.TargetsConformingCount
	}

	if s.Res.TargetsCount != nil {
		auditProfileTargetOverride["targets_count"] = *s.Res.TargetsCount
	}

	if s.Res.TargetsOverridingCount != nil {
		auditProfileTargetOverride["targets_overriding_count"] = *s.Res.TargetsOverridingCount
	}

	if s.Res.TargetsOverridingOfflineMonthsCount != nil {
		auditProfileTargetOverride["targets_overriding_offline_months_count"] = *s.Res.TargetsOverridingOfflineMonthsCount
	}

	if s.Res.TargetsOverridingOnlineMonthsCount != nil {
		auditProfileTargetOverride["targets_overriding_online_months_count"] = *s.Res.TargetsOverridingOnlineMonthsCount
	}

	if s.Res.TargetsOverridingPaidUsageCount != nil {
		auditProfileTargetOverride["targets_overriding_paid_usage_count"] = *s.Res.TargetsOverridingPaidUsageCount
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAuditProfileTargetOverridesDataSource().Schema["target_override_collection"].Elem.(*schema.Resource).Schema)
		auditProfileTargetOverride["items"] = items
	}

	resources = append(resources, auditProfileTargetOverride)
	if err := s.D.Set("target_override_collection", resources); err != nil {
		return err
	}

	return nil
}

func TargetOverrideSummaryToMap(obj oci_data_safe.TargetOverrideSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.IsPaidUsageEnabled != nil {
		result["is_paid_usage_enabled"] = bool(*obj.IsPaidUsageEnabled)
	}

	if obj.OfflineMonths != nil {
		result["offline_months"] = int(*obj.OfflineMonths)
	}

	if obj.OfflineMonthsSource != nil {
		result["offline_months_source"] = string(*obj.OfflineMonthsSource)
	}

	if obj.OnlineMonths != nil {
		result["online_months"] = int(*obj.OnlineMonths)
	}

	if obj.OnlineMonthsSource != nil {
		result["online_months_source"] = string(*obj.OnlineMonthsSource)
	}

	if obj.PaidUsageSource != nil {
		result["paid_usage_source"] = string(*obj.PaidUsageSource)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDatabaseId != nil {
		result["target_database_id"] = string(*obj.TargetDatabaseId)
	}

	return result
}
