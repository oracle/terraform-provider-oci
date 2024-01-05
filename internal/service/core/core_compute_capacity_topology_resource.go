// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreComputeCapacityTopologyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreComputeCapacityTopology,
		Read:     readCoreComputeCapacityTopology,
		Update:   updateCoreComputeCapacityTopology,
		Delete:   deleteCoreComputeCapacityTopology,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"capacity_source": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"capacity_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DEDICATED",
							}, true),
						},

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
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
		},
	}
}

func createCoreComputeCapacityTopology(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreComputeCapacityTopology(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreComputeCapacityTopology(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreComputeCapacityTopology(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeCapacityTopologyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreComputeCapacityTopologyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.ComputeCapacityTopology
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreComputeCapacityTopologyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreComputeCapacityTopologyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ComputeCapacityTopologyLifecycleStateCreating),
	}
}

func (s *CoreComputeCapacityTopologyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ComputeCapacityTopologyLifecycleStateActive),
	}
}

func (s *CoreComputeCapacityTopologyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ComputeCapacityTopologyLifecycleStateDeleting),
	}
}

func (s *CoreComputeCapacityTopologyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ComputeCapacityTopologyLifecycleStateDeleted),
	}
}

func (s *CoreComputeCapacityTopologyResourceCrud) Create() error {
	request := oci_core.CreateComputeCapacityTopologyRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capacitySource, ok := s.D.GetOkExists("capacity_source"); ok {
		if tmpList := capacitySource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "capacity_source", 0)
			tmp, err := s.mapToCreateCapacitySourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CapacitySource = tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateComputeCapacityTopology(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ComputeCapacityTopology

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "computecapacitytopology", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	} else {
		var identifier *string
		identifier = s.Res.Id
		s.D.SetId(*identifier)
	}
	return s.Get()
}

func (s *CoreComputeCapacityTopologyResourceCrud) Get() error {
	request := oci_core.GetComputeCapacityTopologyRequest{}

	tmp := s.D.Id()
	request.ComputeCapacityTopologyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetComputeCapacityTopology(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeCapacityTopology
	return nil
}

func (s *CoreComputeCapacityTopologyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateComputeCapacityTopologyRequest{}

	if capacitySource, ok := s.D.GetOkExists("capacity_source"); ok {
		if tmpList := capacitySource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "capacity_source", 0)
			tmp, err := s.mapToUpdateCapacitySourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CapacitySource = tmp
		}
	}

	tmp := s.D.Id()
	request.ComputeCapacityTopologyId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateComputeCapacityTopology(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "computecapacitytopology", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *CoreComputeCapacityTopologyResourceCrud) Delete() error {
	request := oci_core.DeleteComputeCapacityTopologyRequest{}

	tmp := s.D.Id()
	request.ComputeCapacityTopologyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.DeleteComputeCapacityTopology(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "computecapacitytopology", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CoreComputeCapacityTopologyResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacitySource != nil {
		capacitySourceArray := []interface{}{}
		if capacitySourceMap := CapacitySourceToMap(&s.Res.CapacitySource); capacitySourceMap != nil {
			capacitySourceArray = append(capacitySourceArray, capacitySourceMap)
		}
		s.D.Set("capacity_source", capacitySourceArray)
	} else {
		s.D.Set("capacity_source", nil)
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ComputeCapacityTopologySummaryToMap(obj oci_core.ComputeCapacityTopologySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

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

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CoreComputeCapacityTopologyResourceCrud) mapToCreateCapacitySourceDetails(fieldKeyFormat string) (oci_core.CreateCapacitySourceDetails, error) {
	var baseObject oci_core.CreateCapacitySourceDetails
	//discriminator
	capacityTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_type"))
	var capacityType string
	if ok {
		capacityType = capacityTypeRaw.(string)
	} else {
		capacityType = "" // default value
	}
	switch strings.ToLower(capacityType) {
	case strings.ToLower("DEDICATED"):
		details := oci_core.CreateDedicatedCapacitySourceDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown capacity_type '%v' was specified", capacityType)
	}
	return baseObject, nil
}

func (s *CoreComputeCapacityTopologyResourceCrud) mapToUpdateCapacitySourceDetails(fieldKeyFormat string) (oci_core.UpdateCapacitySourceDetails, error) {
	var baseObject oci_core.UpdateCapacitySourceDetails
	//discriminator
	capacityTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_type"))
	var capacityType string
	if ok {
		capacityType = capacityTypeRaw.(string)
	} else {
		capacityType = "" // default value
	}
	switch strings.ToLower(capacityType) {
	case strings.ToLower("DEDICATED"):
		details := oci_core.UpdateDedicatedCapacitySourceDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown capacity_type '%v' was specified", capacityType)
	}
	return baseObject, nil
}

func CapacitySourceToMap(obj *oci_core.CapacitySource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.DedicatedCapacitySource:
		result["capacity_type"] = "DEDICATED"
		result["compartment_id"] = v.CompartmentId
	default:
		log.Printf("[WARN] Received 'capacity_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreComputeCapacityTopologyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeComputeCapacityTopologyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ComputeCapacityTopologyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeComputeCapacityTopologyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "computecapacitytopology", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
