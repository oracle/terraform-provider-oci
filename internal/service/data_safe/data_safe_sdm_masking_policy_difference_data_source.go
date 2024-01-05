// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSdmMaskingPolicyDifferenceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sdm_masking_policy_difference_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeSdmMaskingPolicyDifferenceResource(), fieldMap, readSingularDataSafeSdmMaskingPolicyDifference)
}

func readSingularDataSafeSdmMaskingPolicyDifference(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSdmMaskingPolicyDifferenceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSdmMaskingPolicyDifferenceResponse
}

func (s *DataSafeSdmMaskingPolicyDifferenceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSdmMaskingPolicyDifferenceDataSourceCrud) Get() error {
	request := oci_data_safe.GetSdmMaskingPolicyDifferenceRequest{}

	if sdmMaskingPolicyDifferenceId, ok := s.D.GetOkExists("sdm_masking_policy_difference_id"); ok {
		tmp := sdmMaskingPolicyDifferenceId.(string)
		request.SdmMaskingPolicyDifferenceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSdmMaskingPolicyDifference(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSdmMaskingPolicyDifferenceDataSourceCrud) SetData() error {
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

	s.D.Set("difference_type", s.Res.DifferenceType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeCreationStarted != nil {
		s.D.Set("time_creation_started", s.Res.TimeCreationStarted.String())
	}

	return nil
}
