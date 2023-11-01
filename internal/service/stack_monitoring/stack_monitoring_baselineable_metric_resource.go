// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringBaselineableMetricResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringBaselineableMetric,
		Read:     readStackMonitoringBaselineableMetric,
		Update:   updateStackMonitoringBaselineableMetric,
		Delete:   deleteStackMonitoringBaselineableMetric,
		Schema: map[string]*schema.Schema{
			// Required
			"column": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"is_out_of_box": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_updated_by": {
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
			"tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createStackMonitoringBaselineableMetric(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringBaselineableMetric(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringBaselineableMetric(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringBaselineableMetric(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringBaselineableMetricResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.BaselineableMetric
	DisableNotFoundRetries bool
}

func (s *StackMonitoringBaselineableMetricResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringBaselineableMetricResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *StackMonitoringBaselineableMetricResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.BaselineableMetricLifeCycleStatesActive),
	}
}

func (s *StackMonitoringBaselineableMetricResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *StackMonitoringBaselineableMetricResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.BaselineableMetricLifeCycleStatesDeleted),
	}
}

func (s *StackMonitoringBaselineableMetricResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateBaselineableMetricRequest{}

	if column, ok := s.D.GetOkExists("column"); ok {
		tmp := column.(string)
		request.Column = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateBaselineableMetric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BaselineableMetric
	return nil
}

func (s *StackMonitoringBaselineableMetricResourceCrud) Get() error {
	request := oci_stack_monitoring.GetBaselineableMetricRequest{}

	tmp := s.D.Id()
	request.BaselineableMetricId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetBaselineableMetric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BaselineableMetric
	return nil
}

func (s *StackMonitoringBaselineableMetricResourceCrud) Update() error {
	request := oci_stack_monitoring.UpdateBaselineableMetricRequest{}

	tmp := s.D.Id()
	request.BaselineableMetricId = &tmp

	if column, ok := s.D.GetOkExists("column"); ok {
		tmp := column.(string)
		request.Column = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp = s.D.Id()
	request.Id = &tmp

	if isOutOfBox, ok := s.D.GetOkExists("is_out_of_box"); ok {
		tmp := isOutOfBox.(bool)
		request.IsOutOfBox = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists("resource_group"); ok {
		tmp := resourceGroup.(string)
		request.ResourceGroup = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_stack_monitoring.BaselineableMetricLifeCycleStatesEnum(state.(string))
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateBaselineableMetric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BaselineableMetric
	return nil
}

func (s *StackMonitoringBaselineableMetricResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteBaselineableMetricRequest{}

	tmp := s.D.Id()
	request.BaselineableMetricId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteBaselineableMetric(context.Background(), request)
	return err
}

func (s *StackMonitoringBaselineableMetricResourceCrud) SetData() error {
	if s.Res.Column != nil {
		s.D.Set("column", *s.Res.Column)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOutOfBox != nil {
		s.D.Set("is_out_of_box", *s.Res.IsOutOfBox)
	}

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", *s.Res.LastUpdatedBy)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ResourceGroup != nil {
		s.D.Set("resource_group", *s.Res.ResourceGroup)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastUpdated != nil {
		s.D.Set("time_last_updated", s.Res.TimeLastUpdated.String())
	}

	return nil
}

func BaselineableMetricSummaryToMap(obj oci_stack_monitoring.BaselineableMetricSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Column != nil {
		result["column"] = string(*obj.Column)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsOutOfBox != nil {
		result["is_out_of_box"] = bool(*obj.IsOutOfBox)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ResourceGroup != nil {
		result["resource_group"] = string(*obj.ResourceGroup)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	return result
}
