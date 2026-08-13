// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicyHealthReportsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingPolicyHealthReportsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_policy_health_report_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_policy_health_report_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subsetting_policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"warning_count": {
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
	}
}

func readDataSafeSubsettingPolicyHealthReportsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyHealthReportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicyHealthReportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingPolicyHealthReportsResponse
}

func (s *DataSafeSubsettingPolicyHealthReportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicyHealthReportsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingPolicyHealthReportsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSubsettingPolicyHealthReportsAccessLevelEnum(accessLevel.(string))
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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListSubsettingPolicyHealthReportsLifecycleStateEnum(state.(string))
	}

	if subsettingPolicyHealthReportId, ok := s.D.GetOkExists("subsetting_policy_health_report_id"); ok {
		tmp := subsettingPolicyHealthReportId.(string)
		request.SubsettingPolicyHealthReportId = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingPolicyHealthReports(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingPolicyHealthReports(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingPolicyHealthReportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicyHealthReportsDataSource-", DataSafeSubsettingPolicyHealthReportsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingPolicyHealthReport := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingPolicyHealthReportSummaryToMap(item))
	}
	subsettingPolicyHealthReport["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingPolicyHealthReportsDataSource().Schema["subsetting_policy_health_report_collection"].Elem.(*schema.Resource).Schema)
		subsettingPolicyHealthReport["items"] = items
	}

	resources = append(resources, subsettingPolicyHealthReport)
	if err := s.D.Set("subsetting_policy_health_report_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingPolicyHealthReportSummaryToMap(obj oci_data_safe.SubsettingPolicyHealthReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ErrorCount != nil {
		result["error_count"] = strconv.FormatInt(*obj.ErrorCount, 10)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubsettingPolicyId != nil {
		result["subsetting_policy_id"] = string(*obj.SubsettingPolicyId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.WarningCount != nil {
		result["warning_count"] = strconv.FormatInt(*obj.WarningCount, 10)
	}

	return result
}
