// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseSchedulingPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseSchedulingPolicyWithContext,
		ReadContext:   readDatabaseSchedulingPolicyWithContext,
		UpdateContext: updateDatabaseSchedulingPolicyWithContext,
		DeleteContext: deleteDatabaseSchedulingPolicyWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"cadence": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"cadence_start_month": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_next_window_starts": {
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

func createDatabaseSchedulingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseSchedulingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseSchedulingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseSchedulingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseSchedulingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseSchedulingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseSchedulingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseSchedulingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseSchedulingPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.SchedulingPolicy
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseSchedulingPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseSchedulingPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.SchedulingPolicyLifecycleStateCreating),
	}
}

func (s *DatabaseSchedulingPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.SchedulingPolicyLifecycleStateNeedsAttention),
		string(oci_database.SchedulingPolicyLifecycleStateAvailable),
	}
}

func (s *DatabaseSchedulingPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.SchedulingPolicyLifecycleStateDeleting),
	}
}

func (s *DatabaseSchedulingPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.SchedulingPolicyLifecycleStateDeleted),
	}
}

func (s *DatabaseSchedulingPolicyResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database.CreateSchedulingPolicyRequest{}

	if cadence, ok := s.D.GetOkExists("cadence"); ok {
		request.Cadence = oci_database.CreateSchedulingPolicyDetailsCadenceEnum(cadence.(string))
	}

	if cadenceStartMonth, ok := s.D.GetOkExists("cadence_start_month"); ok {
		if tmpList := cadenceStartMonth.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cadence_start_month", 0)
			tmp, err := s.mapToMonth(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CadenceStartMonth = &tmp
		}
	}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateSchedulingPolicy(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulingPolicy
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.GetWithContext(ctx)
}

func (s *DatabaseSchedulingPolicyResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetSchedulingPolicyRequest{}

	tmp := s.D.Id()
	request.SchedulingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetSchedulingPolicy(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulingPolicy
	return nil
}

func (s *DatabaseSchedulingPolicyResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateSchedulingPolicyRequest{}

	if cadence, ok := s.D.GetOkExists("cadence"); ok {
		request.Cadence = oci_database.UpdateSchedulingPolicyDetailsCadenceEnum(cadence.(string))
	}

	if cadenceStartMonth, ok := s.D.GetOkExists("cadence_start_month"); ok {
		if tmpList := cadenceStartMonth.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cadence_start_month", 0)
			tmp, err := s.mapToMonth(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CadenceStartMonth = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SchedulingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateSchedulingPolicy(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "schedulingpolicy", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.GetWithContext(ctx)
}

func (s *DatabaseSchedulingPolicyResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database.DeleteSchedulingPolicyRequest{}

	tmp := s.D.Id()
	request.SchedulingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteSchedulingPolicy(ctx, request)
	return err
}

func (s *DatabaseSchedulingPolicyResourceCrud) SetData() error {
	s.D.Set("cadence", s.Res.Cadence)

	if s.Res.CadenceStartMonth != nil {
		s.D.Set("cadence_start_month", []interface{}{MonthToMapPolicy(s.Res.CadenceStartMonth)})
	} else {
		s.D.Set("cadence_start_month", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNextWindowStarts != nil {
		s.D.Set("time_next_window_starts", s.Res.TimeNextWindowStarts.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DatabaseSchedulingPolicyResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func MonthToMapPolicy(obj *oci_database.Month) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseSchedulingPolicyResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeSchedulingPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SchedulingPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeSchedulingPolicyCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "schedulingpolicy", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
