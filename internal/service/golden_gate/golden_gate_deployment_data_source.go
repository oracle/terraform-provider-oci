// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
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

	if s.Res.BackupSchedule != nil {
		s.D.Set("backup_schedule", []interface{}{BackupScheduleToMap(s.Res.BackupSchedule)})
	} else {
		s.D.Set("backup_schedule", nil)
	}
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	s.D.Set("category", s.Res.Category)

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

	if s.Res.DeploymentDiagnosticData != nil {
		s.D.Set("deployment_diagnostic_data", []interface{}{DeploymentDiagnosticDataToMap(s.Res.DeploymentDiagnosticData)})
	} else {
		s.D.Set("deployment_diagnostic_data", nil)
	}

	s.D.Set("deployment_role", s.Res.DeploymentRole)

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

	s.D.Set("environment_type", s.Res.EnvironmentType)

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ingressIps := []interface{}{}
	for _, item := range s.Res.IngressIps {
		ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
	}
	s.D.Set("ingress_ips", ingressIps)

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

	if s.Res.IsStorageUtilizationLimitExceeded != nil {
		s.D.Set("is_storage_utilization_limit_exceeded", *s.Res.IsStorageUtilizationLimitExceeded)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	if s.Res.LoadBalancerId != nil {
		s.D.Set("load_balancer_id", *s.Res.LoadBalancerId)
	}

	if s.Res.LoadBalancerSubnetId != nil {
		s.D.Set("load_balancer_subnet_id", *s.Res.LoadBalancerSubnetId)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	if s.Res.MaintenanceConfiguration != nil {
		s.D.Set("maintenance_configuration", []interface{}{MaintenanceConfigurationToMap(s.Res.MaintenanceConfiguration)})
	} else {
		s.D.Set("maintenance_configuration", nil)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	s.D.Set("next_maintenance_action_type", s.Res.NextMaintenanceActionType)

	if s.Res.NextMaintenanceDescription != nil {
		s.D.Set("next_maintenance_description", *s.Res.NextMaintenanceDescription)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.OggData != nil {
		s.D.Set("ogg_data", []interface{}{OggDeploymentToMap(s.Res.OggData, s.D)})
	} else {
		s.D.Set("ogg_data", nil)
	}

	placements := []interface{}{}
	for _, item := range s.Res.Placements {
		placements = append(placements, DeploymentPlacementInfoToMap(item))
	}
	s.D.Set("placements", placements)

	if s.Res.PrivateIpAddress != nil {
		s.D.Set("private_ip_address", *s.Res.PrivateIpAddress)
	}

	if s.Res.PublicIpAddress != nil {
		s.D.Set("public_ip_address", *s.Res.PublicIpAddress)
	}

	if s.Res.SourceDeploymentId != nil {
		s.D.Set("source_deployment_id", *s.Res.SourceDeploymentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageUtilizationInBytes != nil {
		s.D.Set("storage_utilization_in_bytes", strconv.FormatInt(*s.Res.StorageUtilizationInBytes, 10))
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastBackupScheduled != nil {
		s.D.Set("time_last_backup_scheduled", s.Res.TimeLastBackupScheduled.String())
	}

	if s.Res.TimeNextBackupScheduled != nil {
		s.D.Set("time_next_backup_scheduled", s.Res.TimeNextBackupScheduled.String())
	}

	if s.Res.TimeOfNextMaintenance != nil {
		s.D.Set("time_of_next_maintenance", s.Res.TimeOfNextMaintenance.String())
	}

	if s.Res.TimeOggVersionSupportedUntil != nil {
		s.D.Set("time_ogg_version_supported_until", s.Res.TimeOggVersionSupportedUntil.String())
	}

	if s.Res.TimeRoleChanged != nil {
		s.D.Set("time_role_changed", s.Res.TimeRoleChanged.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeUpgradeRequired != nil {
		s.D.Set("time_upgrade_required", s.Res.TimeUpgradeRequired.String())
	}

	return nil
}
