// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataccMaintenanceExecutionsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataccMaintenanceExecutionsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_run_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_subtype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_accepted_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_accepted_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_execution_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_action_timeout_in_mins": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
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
									"infrastructure_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_custom_action_timeout_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"maintenance_run_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"maintenance_subtype": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"maintenance_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patching_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"target_resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_time_taken_in_mins": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"workflow_id": {
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

func readDataccMaintenanceExecutionsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccMaintenanceExecutionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataccMaintenanceExecutionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacc.BaseinfraClient
	Res    *oci_datacc.ListMaintenanceExecutionsResponse
}

func (s *DataccMaintenanceExecutionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataccMaintenanceExecutionsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.ListMaintenanceExecutionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if infrastructureId, ok := s.D.GetOkExists("infrastructure_id"); ok {
		tmp := infrastructureId.(string)
		request.InfrastructureId = &tmp
	}

	if maintenanceRunId, ok := s.D.GetOkExists("maintenance_run_id"); ok {
		tmp := maintenanceRunId.(string)
		request.MaintenanceRunId = &tmp
	}

	if maintenanceSubtype, ok := s.D.GetOkExists("maintenance_subtype"); ok {
		request.MaintenanceSubtype = oci_datacc.ListMaintenanceExecutionsMaintenanceSubtypeEnum(maintenanceSubtype.(string))
	}

	if maintenanceType, ok := s.D.GetOkExists("maintenance_type"); ok {
		request.MaintenanceType = oci_datacc.ListMaintenanceExecutionsMaintenanceTypeEnum(maintenanceType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacc.ListMaintenanceExecutionsLifecycleStateEnum(state.(string))
	}

	if targetResourceType, ok := s.D.GetOkExists("target_resource_type"); ok {
		request.TargetResourceType = oci_datacc.ListMaintenanceExecutionsTargetResourceTypeEnum(targetResourceType.(string))
	}

	if timeAcceptedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_accepted_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAcceptedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeAcceptedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeAcceptedLessThanOrEqualTo, ok := s.D.GetOkExists("time_accepted_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAcceptedLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeAcceptedLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_datacc.ListMaintenanceExecutionsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacc")

	response, err := s.Client.ListMaintenanceExecutions(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaintenanceExecutions(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataccMaintenanceExecutionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataccMaintenanceExecutionsDataSource-", DataccMaintenanceExecutionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maintenanceExecution := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaintenanceExecutionSummaryToMap(item))
	}
	maintenanceExecution["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataccMaintenanceExecutionsDataSource().Schema["maintenance_execution_collection"].Elem.(*schema.Resource).Schema)
		maintenanceExecution["items"] = items
	}

	resources = append(resources, maintenanceExecution)
	if err := s.D.Set("maintenance_execution_collection", resources); err != nil {
		return err
	}

	return nil
}

func MaintenanceExecutionSummaryToMap(obj oci_datacc.MaintenanceExecutionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CustomActionTimeoutInMins != nil {
		result["custom_action_timeout_in_mins"] = int(*obj.CustomActionTimeoutInMins)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InfrastructureId != nil {
		result["infrastructure_id"] = string(*obj.InfrastructureId)
	}

	if obj.IsCustomActionTimeoutEnabled != nil {
		result["is_custom_action_timeout_enabled"] = bool(*obj.IsCustomActionTimeoutEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MaintenanceRunId != nil {
		result["maintenance_run_id"] = string(*obj.MaintenanceRunId)
	}

	result["maintenance_subtype"] = string(obj.MaintenanceSubtype)

	result["maintenance_type"] = string(obj.MaintenanceType)

	result["patching_mode"] = string(obj.PatchingMode)

	if obj.SourceVersion != nil {
		result["source_version"] = string(*obj.SourceVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	result["target_resource_type"] = string(obj.TargetResourceType)

	if obj.TargetVersion != nil {
		result["target_version"] = string(*obj.TargetVersion)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TotalTimeTakenInMins != nil {
		result["total_time_taken_in_mins"] = int(*obj.TotalTimeTakenInMins)
	}

	if obj.WorkflowId != nil {
		result["workflow_id"] = string(*obj.WorkflowId)
	}

	return result
}
