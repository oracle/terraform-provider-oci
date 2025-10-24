// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dif

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dif "github.com/oracle/oci-go-sdk/v65/dif"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DifStackDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stack_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DifStackResource(), fieldMap, readSingularDifStackWithContext)
}

func readSingularDifStackWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DifStackDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DifStackDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dif.StackClient
	Res    *oci_dif.GetStackResponse
}

func (s *DifStackDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DifStackDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dif.GetStackRequest{}

	if stackId, ok := s.D.GetOkExists("stack_id"); ok {
		tmp := stackId.(string)
		request.StackId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dif")

	response, err := s.Client.GetStack(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DifStackDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	adb := []interface{}{}
	for _, item := range s.Res.Adb {
		adb = append(adb, AdbDetailToMap(item))
	}
	s.D.Set("adb", adb)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	dataflow := []interface{}{}
	for _, item := range s.Res.Dataflow {
		dataflow = append(dataflow, DataflowDetailToMap(item))
	}
	s.D.Set("dataflow", dataflow)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	genai := []interface{}{}
	for _, item := range s.Res.Genai {
		genai = append(genai, GenAiDetailToMap(item))
	}
	s.D.Set("genai", genai)

	ggcs := []interface{}{}
	for _, item := range s.Res.Ggcs {
		ggcs = append(ggcs, GgcsDetailToMap(item))
	}
	s.D.Set("ggcs", ggcs)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotificationEmail != nil {
		s.D.Set("notification_email", *s.Res.NotificationEmail)
	}

	objectstorage := []interface{}{}
	for _, item := range s.Res.Objectstorage {
		objectstorage = append(objectstorage, ObjectStorageDetailToMap(item))
	}
	s.D.Set("objectstorage", objectstorage)

	serviceDetails := []interface{}{}
	for _, item := range s.Res.ServiceDetails {
		serviceDetails = append(serviceDetails, ServiceDetailResponseToMap(item))
	}
	s.D.Set("service_details", serviceDetails)

	s.D.Set("services", s.Res.Services)

	s.D.Set("stack_templates", s.Res.StackTemplates)

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
