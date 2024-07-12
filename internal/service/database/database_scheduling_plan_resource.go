// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseSchedulingPlanResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseSchedulingPlan,
		Read:     readDatabaseSchedulingPlan,
		Update:   updateDatabaseSchedulingPlan,
		Delete:   deleteDatabaseSchedulingPlan,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scheduling_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"is_using_recommended_scheduled_actions": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_time_in_mins": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_intent": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseSchedulingPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseSchedulingPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseSchedulingPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseSchedulingPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseSchedulingPlanResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.SchedulingPlan
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseSchedulingPlanResourceCrud) Update() error {
	panic("No update required")
}

func (s *DatabaseSchedulingPlanResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseSchedulingPlanResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.SchedulingPlanLifecycleStateCreating),
	}
}

func (s *DatabaseSchedulingPlanResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.SchedulingPlanLifecycleStateNeedsAttention),
		string(oci_database.SchedulingPlanLifecycleStateAvailable),
	}
}

func (s *DatabaseSchedulingPlanResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.SchedulingPlanLifecycleStateDeleting),
	}
}

func (s *DatabaseSchedulingPlanResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.SchedulingPlanLifecycleStateDeleted),
	}
}

func (s *DatabaseSchedulingPlanResourceCrud) Create() error {
	request := oci_database.CreateSchedulingPlanRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isUsingRecommendedScheduledActions, ok := s.D.GetOkExists("is_using_recommended_scheduled_actions"); ok {
		tmp := isUsingRecommendedScheduledActions.(bool)
		request.IsUsingRecommendedScheduledActions = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	if serviceType, ok := s.D.GetOkExists("service_type"); ok {
		request.ServiceType = oci_database.CreateSchedulingPlanDetailsServiceTypeEnum(serviceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateSchedulingPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulingPlan
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.Get()
}

func (s *DatabaseSchedulingPlanResourceCrud) Get() error {
	request := oci_database.GetSchedulingPlanRequest{}

	tmp := s.D.Id()
	request.SchedulingPlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetSchedulingPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulingPlan
	return nil
}

func (s *DatabaseSchedulingPlanResourceCrud) Delete() error {
	request := oci_database.DeleteSchedulingPlanRequest{}

	tmp := s.D.Id()
	request.SchedulingPlanId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteSchedulingPlan(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "schedulingplan", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseSchedulingPlanResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsUsingRecommendedScheduledActions != nil {
		s.D.Set("is_using_recommended_scheduled_actions", *s.Res.IsUsingRecommendedScheduledActions)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("plan_intent", s.Res.PlanIntent)

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.SchedulingPolicyId != nil {
		s.D.Set("scheduling_policy_id", *s.Res.SchedulingPolicyId)
	}

	s.D.Set("service_type", s.Res.ServiceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func SchedulingPlanSummaryToMap(obj oci_database.SchedulingPlanSummary) map[string]interface{} {
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

	if obj.EstimatedTimeInMins != nil {
		result["estimated_time_in_mins"] = int(*obj.EstimatedTimeInMins)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsUsingRecommendedScheduledActions != nil {
		result["is_using_recommended_scheduled_actions"] = bool(*obj.IsUsingRecommendedScheduledActions)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}
	result["plan_intent"] = string(obj.PlanIntent)

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.SchedulingPolicyId != nil {
		result["scheduling_policy_id"] = string(*obj.SchedulingPolicyId)
	}

	result["service_type"] = string(obj.ServiceType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DatabaseSchedulingPlanResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeSchedulingPlanCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SchedulingPlanId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeSchedulingPlanCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "schedulingplan", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
