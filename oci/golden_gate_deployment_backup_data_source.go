// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v48/goldengate"
)

func init() {
	RegisterDatasource("oci_golden_gate_deployment_backup", GoldenGateDeploymentBackupDataSource())
}

func GoldenGateDeploymentBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deployment_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(GoldenGateDeploymentBackupResource(), fieldMap, readSingularGoldenGateDeploymentBackup)
}

func readSingularGoldenGateDeploymentBackup(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()

	return ReadResource(sync)
}

type GoldenGateDeploymentBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetDeploymentBackupResponse
}

func (s *GoldenGateDeploymentBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentBackupDataSourceCrud) Get() error {
	request := oci_golden_gate.GetDeploymentBackupRequest{}

	if deploymentBackupId, ok := s.D.GetOkExists("deployment_backup_id"); ok {
		tmp := deploymentBackupId.(string)
		request.DeploymentBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetDeploymentBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateDeploymentBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeploymentId != nil {
		s.D.Set("deployment_id", *s.Res.DeploymentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.OggVersion != nil {
		s.D.Set("ogg_version", *s.Res.OggVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfBackup != nil {
		s.D.Set("time_of_backup", s.Res.TimeOfBackup.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
