// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ddfs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ddfs "github.com/oracle/oci-go-sdk/v65/ddfs"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DdfsInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DdfsInstanceResource(), fieldMap, readSingularDdfsInstanceWithContext)
}

func readSingularDdfsInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DdfsInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InstanceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DdfsInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ddfs.InstanceClient
	Res    *oci_ddfs.GetInstanceResponse
}

func (s *DdfsInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DdfsInstanceDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ddfs.GetInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ddfs")

	response, err := s.Client.GetInstance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DdfsInstanceDataSourceCrud) SetData() error {
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

	if s.Res.FhirServiceEndpoint != nil {
		s.D.Set("fhir_service_endpoint", *s.Res.FhirServiceEndpoint)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdcsUrl != nil {
		s.D.Set("idcs_url", *s.Res.IdcsUrl)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PublicIp != nil {
		s.D.Set("public_ip", *s.Res.PublicIp)
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
