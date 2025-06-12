// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementInternalOccmDemandSignalResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementInternalOccmDemandSignal,
		Read:     readCapacityManagementInternalOccmDemandSignal,
		Update:   updateCapacityManagementInternalOccmDemandSignal,
		Delete:   deleteCapacityManagementInternalOccmDemandSignal,
		Schema: map[string]*schema.Schema{
			// Required
			"occm_demand_signal_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
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
			"occ_customer_group_id": {
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

func createCapacityManagementInternalOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementInternalOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementInternalOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementInternalOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InternalDemandSignalClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementInternalOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CapacityManagementInternalOccmDemandSignalResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.InternalDemandSignalClient
	Res                    *oci_capacity_management.InternalOccmDemandSignal
	DisableNotFoundRetries bool
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) ID() string {
	return GetInternalOccmDemandSignalCompositeId(*s.Res.Id)
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateCreating),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateActive),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateDeleting),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalItemLifecycleStateDeleted),
	}
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) Create() error {
	request := oci_capacity_management.UpdateInternalOccmDemandSignalRequest{}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		request.LifecycleDetails = oci_capacity_management.UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum(lifecycleDetails.(string))
	}

	if occmDemandSignalId, ok := s.D.GetOkExists("occm_demand_signal_id"); ok {
		tmp := occmDemandSignalId.(string)
		request.OccmDemandSignalId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateInternalOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternalOccmDemandSignal
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) Get() error {
	request := oci_capacity_management.GetInternalOccmDemandSignalRequest{}

	if occmDemandSignalId, ok := s.D.GetOkExists("occm_demand_signal_id"); ok {
		tmp := occmDemandSignalId.(string)
		request.OccmDemandSignalId = &tmp
	}

	occmDemandSignalId, err := parseInternalOccmDemandSignalCompositeId(s.D.Id())
	if err == nil {
		request.OccmDemandSignalId = &occmDemandSignalId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetInternalOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternalOccmDemandSignal
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) Update() error {
	request := oci_capacity_management.UpdateInternalOccmDemandSignalRequest{}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		request.LifecycleDetails = oci_capacity_management.UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum(lifecycleDetails.(string))
	}

	if occmDemandSignalId, ok := s.D.GetOkExists("occm_demand_signal_id"); ok {
		tmp := occmDemandSignalId.(string)
		request.OccmDemandSignalId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateInternalOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternalOccmDemandSignal
	return nil
}

func (s *CapacityManagementInternalOccmDemandSignalResourceCrud) SetData() error {

	_, err := parseInternalOccmDemandSignalCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
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

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
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

func GetInternalOccmDemandSignalCompositeId(occmDemandSignalId string) string {
	occmDemandSignalId = url.PathEscape(occmDemandSignalId)
	compositeId := "internal/occmDemandSignals/" + occmDemandSignalId + ""
	return compositeId
}

func parseInternalOccmDemandSignalCompositeId(compositeId string) (occmDemandSignalId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("internal/occmDemandSignals/.*", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	occmDemandSignalId, _ = url.PathUnescape(parts[1])

	return
}

func InternalOccmDemandSignalSummaryToMap(obj oci_capacity_management.InternalOccmDemandSignalSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	if obj.OccCustomerGroupId != nil {
		result["occ_customer_group_id"] = string(*obj.OccCustomerGroupId)
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
