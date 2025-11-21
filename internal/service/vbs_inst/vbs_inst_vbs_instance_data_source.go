// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vbs_inst

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_vbs_inst "github.com/oracle/oci-go-sdk/v65/vbsinst"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VbsInstVbsInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vbs_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(VbsInstVbsInstanceResource(), fieldMap, readSingularVbsInstVbsInstance)
}

func readSingularVbsInstVbsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VbsInstVbsInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbsInstanceClient()

	return tfresource.ReadResource(sync)
}

type VbsInstVbsInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vbs_inst.VbsInstanceClient
	Res    *oci_vbs_inst.GetVbsInstanceResponse
}

func (s *VbsInstVbsInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VbsInstVbsInstanceDataSourceCrud) Get() error {
	request := oci_vbs_inst.GetVbsInstanceRequest{}

	if vbsInstanceId, ok := s.D.GetOkExists("vbs_instance_id"); ok {
		tmp := vbsInstanceId.(string)
		request.VbsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vbs_inst")

	response, err := s.Client.GetVbsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VbsInstVbsInstanceDataSourceCrud) SetData() error {
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

	if s.Res.IsResourceUsageAgreementGranted != nil {
		s.D.Set("is_resource_usage_agreement_granted", *s.Res.IsResourceUsageAgreementGranted)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ResourceCompartmentId != nil {
		s.D.Set("resource_compartment_id", *s.Res.ResourceCompartmentId)
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

	if s.Res.VbsAccessUrl != nil {
		s.D.Set("vbs_access_url", *s.Res.VbsAccessUrl)
	}

	return nil
}
