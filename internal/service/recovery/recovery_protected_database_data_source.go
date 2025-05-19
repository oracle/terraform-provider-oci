// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryProtectedDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["protected_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(RecoveryProtectedDatabaseResource(), fieldMap, readSingularRecoveryProtectedDatabase)
}

func readSingularRecoveryProtectedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

type RecoveryProtectedDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_recovery.DatabaseRecoveryClient
	Res    *oci_recovery.GetProtectedDatabaseResponse
}

func (s *RecoveryProtectedDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RecoveryProtectedDatabaseDataSourceCrud) Get() error {
	request := oci_recovery.GetProtectedDatabaseRequest{}

	if protectedDatabaseId, ok := s.D.GetOkExists("protected_database_id"); ok {
		tmp := protectedDatabaseId.(string)
		request.ProtectedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "recovery")

	response, err := s.Client.GetProtectedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RecoveryProtectedDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	s.D.Set("database_size", s.Res.DatabaseSize)

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("health", s.Res.Health)

	if s.Res.HealthDetails != nil {
		s.D.Set("health_details", *s.Res.HealthDetails)
	}

	if s.Res.IsReadOnlyResource != nil {
		s.D.Set("is_read_only_resource", *s.Res.IsReadOnlyResource)
	}

	if s.Res.IsRedoLogsShipped != nil {
		s.D.Set("is_redo_logs_shipped", *s.Res.IsRedoLogsShipped)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Metrics != nil {
		s.D.Set("metrics", []interface{}{MetricsToMap(s.Res.Metrics)})
	} else {
		s.D.Set("metrics", nil)
	}

	if s.Res.PolicyLockedDateTime != nil {
		s.D.Set("policy_locked_date_time", *s.Res.PolicyLockedDateTime)
	}

	if s.Res.ProtectionPolicyId != nil {
		s.D.Set("protection_policy_id", *s.Res.ProtectionPolicyId)
	}

	recoveryServiceSubnets := []interface{}{}
	for _, item := range s.Res.RecoveryServiceSubnets {
		recoveryServiceSubnets = append(recoveryServiceSubnets, RecoveryServiceSubnetDetailsToMap(item))
	}
	s.D.Set("recovery_service_subnets", recoveryServiceSubnets)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
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

	if s.Res.VpcUserName != nil {
		s.D.Set("vpc_user_name", *s.Res.VpcUserName)
	}

	return nil
}
