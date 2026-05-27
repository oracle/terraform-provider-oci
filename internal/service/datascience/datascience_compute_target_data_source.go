// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceComputeTargetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compute_target_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatascienceComputeTargetResource(), fieldMap, readSingularDatascienceComputeTargetWithContext)
}

func readSingularDatascienceComputeTargetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatascienceComputeTargetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetComputeTargetResponse
}

func (s *DatascienceComputeTargetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceComputeTargetDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datascience.GetComputeTargetRequest{}

	if computeTargetId, ok := s.D.GetOkExists("compute_target_id"); ok {
		tmp := computeTargetId.(string)
		request.ComputeTargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetComputeTarget(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceComputeTargetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeConfigurationDetails != nil {
		computeConfigurationDetailsArray := []interface{}{}
		if computeConfigurationDetailsMap := ComputeConfigurationDetailsToMap(&s.Res.ComputeConfigurationDetails); computeConfigurationDetailsMap != nil {
			computeConfigurationDetailsArray = append(computeConfigurationDetailsArray, computeConfigurationDetailsMap)
		}
		s.D.Set("compute_configuration_details", computeConfigurationDetailsArray)
	} else {
		s.D.Set("compute_configuration_details", nil)
	}

	if s.Res.ComputeTargetSystemData != nil {
		computeTargetSystemDataArray := []interface{}{}
		if computeTargetSystemDataMap := ComputeTargetSystemDataToMap(&s.Res.ComputeTargetSystemData); computeTargetSystemDataMap != nil {
			computeTargetSystemDataArray = append(computeTargetSystemDataArray, computeTargetSystemDataMap)
		}
		s.D.Set("compute_target_system_data", computeTargetSystemDataArray)
	} else {
		s.D.Set("compute_target_system_data", nil)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("metadata", s.Res.Metadata)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
