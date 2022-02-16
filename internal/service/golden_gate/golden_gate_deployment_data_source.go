// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v58/goldengate"
)

func GoldenGateDeploymentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GoldenGateDeploymentResource(), fieldMap, readSingularGoldenGateDeployment)
}

func readSingularGoldenGateDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetDeploymentResponse
}

func (s *GoldenGateDeploymentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentDataSourceCrud) Get() error {
	request := oci_golden_gate.GetDeploymentRequest{}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateDeploymentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeploymentBackupId != nil {
		s.D.Set("deployment_backup_id", *s.Res.DeploymentBackupId)
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	if s.Res.DeploymentUrl != nil {
		s.D.Set("deployment_url", *s.Res.DeploymentUrl)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoScalingEnabled != nil {
		s.D.Set("is_auto_scaling_enabled", *s.Res.IsAutoScalingEnabled)
	}

	if s.Res.IsHealthy != nil {
		s.D.Set("is_healthy", *s.Res.IsHealthy)
	}

	if s.Res.IsLatestVersion != nil {
		s.D.Set("is_latest_version", *s.Res.IsLatestVersion)
	}

	if s.Res.IsPublic != nil {
		s.D.Set("is_public", *s.Res.IsPublic)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.OggData != nil {
		s.D.Set("ogg_data", []interface{}{OggDeploymentToMap(s.Res.OggData, s.D)})
	} else {
		s.D.Set("ogg_data", nil)
	}

	if s.Res.PrivateIpAddress != nil {
		s.D.Set("private_ip_address", *s.Res.PrivateIpAddress)
	}

	if s.Res.PublicIpAddress != nil {
		s.D.Set("public_ip_address", *s.Res.PublicIpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeUpgradeRequired != nil {
		s.D.Set("time_upgrade_required", s.Res.TimeUpgradeRequired.String())
	}

	return nil
}
