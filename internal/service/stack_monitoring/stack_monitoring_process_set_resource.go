// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringProcessSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringProcessSet,
		Read:     readStackMonitoringProcessSet,
		Update:   updateStackMonitoringProcessSet,
		Delete:   deleteStackMonitoringProcessSet,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"specification": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"label": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"process_command": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"process_line_regex_pattern": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"process_user": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
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
			"revision": {
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

func createStackMonitoringProcessSet(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringProcessSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringProcessSet(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringProcessSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringProcessSet(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringProcessSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringProcessSet(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringProcessSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringProcessSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.ProcessSet
	DisableNotFoundRetries bool
}

func (s *StackMonitoringProcessSetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringProcessSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateCreating),
	}
}

func (s *StackMonitoringProcessSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateActive),
	}
}

func (s *StackMonitoringProcessSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateDeleting),
	}
}

func (s *StackMonitoringProcessSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateDeleted),
	}
}

func (s *StackMonitoringProcessSetResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateProcessSetRequest{}

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

	if specification, ok := s.D.GetOkExists("specification"); ok {
		if tmpList := specification.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "specification", 0)
			tmp, err := s.mapToProcessSetSpecification(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Specification = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateProcessSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProcessSet
	return nil
}

func (s *StackMonitoringProcessSetResourceCrud) Get() error {
	request := oci_stack_monitoring.GetProcessSetRequest{}

	tmp := s.D.Id()
	request.ProcessSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetProcessSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProcessSet
	return nil
}

func (s *StackMonitoringProcessSetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_stack_monitoring.UpdateProcessSetRequest{}

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
	request.ProcessSetId = &tmp

	if specification, ok := s.D.GetOkExists("specification"); ok {
		if tmpList := specification.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "specification", 0)
			tmp, err := s.mapToProcessSetSpecification(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Specification = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateProcessSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProcessSet
	return nil
}

func (s *StackMonitoringProcessSetResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteProcessSetRequest{}

	tmp := s.D.Id()
	request.ProcessSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteProcessSet(context.Background(), request)
	return err
}

func (s *StackMonitoringProcessSetResourceCrud) SetData() error {
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

	if s.Res.Revision != nil {
		s.D.Set("revision", *s.Res.Revision)
	}

	if s.Res.Specification != nil {
		s.D.Set("specification", []interface{}{ProcessSetSpecificationToMap(s.Res.Specification)})
	} else {
		s.D.Set("specification", nil)
	}

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

func (s *StackMonitoringProcessSetResourceCrud) mapToProcessSetSpecification(fieldKeyFormat string) (oci_stack_monitoring.ProcessSetSpecification, error) {
	result := oci_stack_monitoring.ProcessSetSpecification{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_stack_monitoring.ProcessSetSpecificationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToProcessSetSpecificationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func ProcessSetSpecificationToMap(obj *oci_stack_monitoring.ProcessSetSpecification) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ProcessSetSpecificationDetailsToMap(item))
	}
	result["items"] = items

	return result
}

func (s *StackMonitoringProcessSetResourceCrud) mapToProcessSetSpecificationDetails(fieldKeyFormat string) (oci_stack_monitoring.ProcessSetSpecificationDetails, error) {
	result := oci_stack_monitoring.ProcessSetSpecificationDetails{}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if processCommand, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "process_command")); ok {
		tmp := processCommand.(string)
		result.ProcessCommand = &tmp
	}

	if processLineRegexPattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "process_line_regex_pattern")); ok {
		tmp := processLineRegexPattern.(string)
		result.ProcessLineRegexPattern = &tmp
	}

	if processUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "process_user")); ok {
		tmp := processUser.(string)
		result.ProcessUser = &tmp
	}

	return result, nil
}

func ProcessSetSpecificationDetailsToMap(obj oci_stack_monitoring.ProcessSetSpecificationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.ProcessCommand != nil {
		result["process_command"] = string(*obj.ProcessCommand)
	}

	if obj.ProcessLineRegexPattern != nil {
		result["process_line_regex_pattern"] = string(*obj.ProcessLineRegexPattern)
	}

	if obj.ProcessUser != nil {
		result["process_user"] = string(*obj.ProcessUser)
	}

	return result
}

func ProcessSetSummaryToMap(obj oci_stack_monitoring.ProcessSetSummary) map[string]interface{} {
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

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Revision != nil {
		result["revision"] = string(*obj.Revision)
	}

	if obj.Specification != nil {
		result["specification"] = []interface{}{ProcessSetSpecificationToMap(obj.Specification)}
	}

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

func (s *StackMonitoringProcessSetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_stack_monitoring.ChangeProcessSetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ProcessSetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.ChangeProcessSetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
