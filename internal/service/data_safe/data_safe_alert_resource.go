// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAlertResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAlert,
		Read:     readDataSafeAlert,
		Update:   updateDataSafeAlert,
		Delete:   deleteDataSafeAlert,
		Schema: map[string]*schema.Schema{
			// Required
			"alert_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"alert_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"feature_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"operation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"severity": {
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
			"target_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createDataSafeAlert(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readDataSafeAlert(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAlert(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeAlert(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeAlertResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.Alert
	DisableNotFoundRetries bool
}

func (s *DataSafeAlertResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAlertResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DataSafeAlertResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AlertLifecycleStateSucceeded),
	}
}

func (s *DataSafeAlertResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DataSafeAlertResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DataSafeAlertResourceCrud) Create() error {
	request := oci_data_safe.UpdateAlertRequest{}

	if alertId, ok := s.D.GetOkExists("alert_id"); ok {
		tmp := alertId.(string)
		request.AlertId = &tmp
	}

	if comment, ok := s.D.GetOkExists("comment"); ok {
		tmp := comment.(string)
		request.Comment = &tmp
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

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.AlertStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAlert(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Alert
	return nil
}

func (s *DataSafeAlertResourceCrud) Get() error {
	request := oci_data_safe.GetAlertRequest{}

	tmp := s.D.Id()
	request.AlertId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAlert(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Alert
	return nil
}

func (s *DataSafeAlertResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateAlertRequest{}

	tmp := s.D.Id()
	request.AlertId = &tmp

	if comment, ok := s.D.GetOkExists("comment"); ok {
		tmp := comment.(string)
		request.Comment = &tmp
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

	if status, ok := s.D.GetOkExists("status"); ok {
		request := oci_data_safe.AlertsUpdateRequest{}
		request.Status = oci_data_safe.AlertStatusEnum(status.(string))
		if compartment, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartment.(string)
			request.CompartmentId = &tmp
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

			response, err := s.Client.AlertsUpdate(context.Background(), request)
			if err != nil {
				return err
			} else {
				workId := *response.OpcWorkRequestId
				log.Printf("response %s", workId)
			}

		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAlert(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Alert
	return nil
}

func (s *DataSafeAlertResourceCrud) SetData() error {
	s.D.Set("alert_type", s.Res.AlertType)

	if s.Res.Comment != nil {
		s.D.Set("comment", *s.Res.Comment)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Operation != nil {
		s.D.Set("operation", *s.Res.Operation)
	}

	s.D.Set("operation_status", s.Res.OperationStatus)

	if s.Res.OperationTime != nil {
		s.D.Set("operation_time", s.Res.OperationTime.String())
	}

	if s.Res.PolicyId != nil {
		s.D.Set("policy_id", *s.Res.PolicyId)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("target_ids", s.Res.TargetIds)

	s.D.Set("target_names", s.Res.TargetNames)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func AlertSummaryToMap(obj oci_data_safe.AlertSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["alert_type"] = string(obj.AlertType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	// It expects the "Feature Attributes" to be a string instead of a Map, hence this workaround.
	var featBytes, err = json.Marshal(obj.FeatureDetails["Feature Attributes"])
	if err != nil {
		fmt.Println("Error in marshall", featBytes)
		return nil
	}
	newMap := map[string]interface{}{}
	newMap["Feature Attributes"] = string(featBytes)
	result["feature_details"] = newMap

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	result["operation_status"] = string(obj.OperationStatus)

	if obj.OperationTime != nil {
		result["operation_time"] = obj.OperationTime.String()
	}

	if obj.PolicyId != nil {
		result["policy_id"] = string(*obj.PolicyId)
	}

	result["severity"] = string(obj.Severity)

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	result["target_ids"] = obj.TargetIds

	result["target_names"] = obj.TargetNames

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataSafeAlertResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeAlertCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AlertId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeAlertCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
