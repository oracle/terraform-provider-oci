// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccmDemandSignalItemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementOccmDemandSignalItem,
		Read:     readCapacityManagementOccmDemandSignalItem,
		Update:   updateCapacityManagementOccmDemandSignalItem,
		Delete:   deleteCapacityManagementOccmDemandSignalItem,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"demand_quantity": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"demand_signal_catalog_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"demand_signal_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"request_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_properties": {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     schema.TypeString,
			},
			"time_needed_before": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Optional
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"demand_signal_namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_name": {
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
		},
	}
}

func createCapacityManagementOccmDemandSignalItem(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementOccmDemandSignalItem(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementOccmDemandSignalItem(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementOccmDemandSignalItem(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalItemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementOccmDemandSignalItemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.DemandSignalClient
	Res                    *oci_capacity_management.OccmDemandSignalItem
	DisableNotFoundRetries bool
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateCreating),
	}
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateActive),
	}
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateDeleting),
	}
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateDeleted),
	}
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) Create() error {
	request := oci_capacity_management.CreateOccmDemandSignalItemRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
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

	if demandQuantity, ok := s.D.GetOkExists("demand_quantity"); ok {
		tmp := demandQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert demandQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.DemandQuantity = &tmpInt64
	}

	if demandSignalCatalogResourceId, ok := s.D.GetOkExists("demand_signal_catalog_resource_id"); ok {
		tmp := demandSignalCatalogResourceId.(string)
		request.DemandSignalCatalogResourceId = &tmp
	}

	if demandSignalId, ok := s.D.GetOkExists("demand_signal_id"); ok {
		tmp := demandSignalId.(string)
		request.DemandSignalId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notes, ok := s.D.GetOkExists("notes"); ok {
		tmp := notes.(string)
		request.Notes = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if requestType, ok := s.D.GetOkExists("request_type"); ok {
		request.RequestType = oci_capacity_management.OccmDemandSignalItemRequestTypeEnum(requestType.(string))
	}

	if resourceProperties, ok := s.D.GetOkExists("resource_properties"); ok {
		request.ResourceProperties = tfresource.ObjectMapToStringMap(resourceProperties.(map[string]interface{}))
	}

	if targetCompartmentId, ok := s.D.GetOkExists("target_compartment_id"); ok {
		tmp := targetCompartmentId.(string)
		request.TargetCompartmentId = &tmp
	}

	if timeNeededBefore, ok := s.D.GetOkExists("time_needed_before"); ok {
		tmp, err := time.Parse(time.RFC3339, timeNeededBefore.(string))
		if err != nil {
			return err
		}
		request.TimeNeededBefore = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateOccmDemandSignalItem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccmDemandSignalItem
	return nil
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) Get() error {
	request := oci_capacity_management.GetOccmDemandSignalItemRequest{}

	tmp := s.D.Id()
	request.OccmDemandSignalItemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetOccmDemandSignalItem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccmDemandSignalItem
	return nil
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) Update() error {
	request := oci_capacity_management.UpdateOccmDemandSignalItemRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if demandQuantity, ok := s.D.GetOkExists("demand_quantity"); ok {
		tmp := demandQuantity.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert demandQuantity string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.DemandQuantity = &tmpInt64
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notes, ok := s.D.GetOkExists("notes"); ok {
		tmp := notes.(string)
		request.Notes = &tmp
	}

	tmp := s.D.Id()
	request.OccmDemandSignalItemId = &tmp

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	if resourceProperties, ok := s.D.GetOkExists("resource_properties"); ok {
		request.ResourceProperties = tfresource.ObjectMapToStringMap(resourceProperties.(map[string]interface{}))
	}

	if targetCompartmentId, ok := s.D.GetOkExists("target_compartment_id"); ok {
		tmp := targetCompartmentId.(string)
		request.TargetCompartmentId = &tmp
	}

	if timeNeededBefore, ok := s.D.GetOkExists("time_needed_before"); ok {
		tmp, err := time.Parse(time.RFC3339, timeNeededBefore.(string))
		if err != nil {
			return err
		}
		request.TimeNeededBefore = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateOccmDemandSignalItem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccmDemandSignalItem
	return nil
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteOccmDemandSignalItemRequest{}

	tmp := s.D.Id()
	request.OccmDemandSignalItemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err := s.Client.DeleteOccmDemandSignalItem(context.Background(), request)
	return err
}

func (s *CapacityManagementOccmDemandSignalItemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DemandQuantity != nil {
		s.D.Set("demand_quantity", strconv.FormatInt(*s.Res.DemandQuantity, 10))
	}

	if s.Res.DemandSignalCatalogResourceId != nil {
		s.D.Set("demand_signal_catalog_resource_id", *s.Res.DemandSignalCatalogResourceId)
	}

	if s.Res.DemandSignalId != nil {
		s.D.Set("demand_signal_id", *s.Res.DemandSignalId)
	}

	s.D.Set("demand_signal_namespace", s.Res.DemandSignalNamespace)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Notes != nil {
		s.D.Set("notes", *s.Res.Notes)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("request_type", s.Res.RequestType)

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	s.D.Set("resource_properties", s.Res.ResourceProperties)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetCompartmentId != nil {
		s.D.Set("target_compartment_id", *s.Res.TargetCompartmentId)
	}

	if s.Res.TimeNeededBefore != nil {
		s.D.Set("time_needed_before", s.Res.TimeNeededBefore.Format(time.RFC3339Nano))
	}

	return nil
}

func OccmDemandSignalItemSummaryToMap(obj oci_capacity_management.OccmDemandSignalItemSummary) map[string]interface{} {
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

	if obj.DemandSignalCatalogResourceId != nil {
		result["demand_signal_catalog_resource_id"] = string(*obj.DemandSignalCatalogResourceId)
	}

	if obj.DemandSignalId != nil {
		result["demand_signal_id"] = string(*obj.DemandSignalId)
	}

	result["demand_signal_namespace"] = string(obj.DemandSignalNamespace)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	} else {
		log.Printf("[WARN] OccmDemandSignalItemSummary Id is nil")
	}

	if obj.Notes != nil {
		result["notes"] = string(*obj.Notes)
	}

	if obj.Quantity != nil {
		result["quantity"] = strconv.FormatInt(*obj.Quantity, 10)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["request_type"] = string(obj.RequestType)

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_properties"] = obj.ResourceProperties

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeNeededBefore != nil {
		result["time_needed_before"] = obj.TimeNeededBefore.String()
	}

	return result
}
