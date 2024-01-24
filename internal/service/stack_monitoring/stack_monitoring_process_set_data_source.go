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

func StackMonitoringProcessSetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["process_set_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringProcessSetResource(), fieldMap, readSingularStackMonitoringProcessSet)
}

func readSingularStackMonitoringProcessSet(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringProcessSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringProcessSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetProcessSetResponse
}

func (s *StackMonitoringProcessSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringProcessSetDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetProcessSetRequest{}

	if processSetId, ok := s.D.GetOkExists("process_set_id"); ok {
		tmp := processSetId.(string)
		request.ProcessSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetProcessSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringProcessSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
