// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package compute_cloud_at_customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ComputeCloudAtCustomerCccInfrastructureDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ccc_infrastructure_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ComputeCloudAtCustomerCccInfrastructureResource(), fieldMap, readSingularComputeCloudAtCustomerCccInfrastructure)
}

func readSingularComputeCloudAtCustomerCccInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccInfrastructureDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.ReadResource(sync)
}

type ComputeCloudAtCustomerCccInfrastructureDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient
	Res    *oci_compute_cloud_at_customer.GetCccInfrastructureResponse
}

func (s *ComputeCloudAtCustomerCccInfrastructureDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ComputeCloudAtCustomerCccInfrastructureDataSourceCrud) Get() error {
	request := oci_compute_cloud_at_customer.GetCccInfrastructureRequest{}

	if cccInfrastructureId, ok := s.D.GetOkExists("ccc_infrastructure_id"); ok {
		tmp := cccInfrastructureId.(string)
		request.CccInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "compute_cloud_at_customer")

	response, err := s.Client.GetCccInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ComputeCloudAtCustomerCccInfrastructureDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CccUpgradeScheduleId != nil {
		s.D.Set("ccc_upgrade_schedule_id", *s.Res.CccUpgradeScheduleId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionDetails != nil {
		s.D.Set("connection_details", *s.Res.ConnectionDetails)
	}

	s.D.Set("connection_state", s.Res.ConnectionState)

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

	if s.Res.InfrastructureInventory != nil {
		s.D.Set("infrastructure_inventory", []interface{}{CccInfrastructureInventoryToMap(s.Res.InfrastructureInventory)})
	} else {
		s.D.Set("infrastructure_inventory", nil)
	}

	if s.Res.InfrastructureNetworkConfiguration != nil {
		s.D.Set("infrastructure_network_configuration", []interface{}{CccInfrastructureNetworkConfigurationToMap(s.Res.InfrastructureNetworkConfiguration)})
	} else {
		s.D.Set("infrastructure_network_configuration", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ProvisioningFingerprint != nil {
		s.D.Set("provisioning_fingerprint", *s.Res.ProvisioningFingerprint)
	}

	if s.Res.ProvisioningPin != nil {
		s.D.Set("provisioning_pin", *s.Res.ProvisioningPin)
	}

	if s.Res.ShortName != nil {
		s.D.Set("short_name", *s.Res.ShortName)
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

	if s.Res.UpgradeInformation != nil {
		s.D.Set("upgrade_information", []interface{}{CccUpgradeInformationToMap(s.Res.UpgradeInformation)})
	} else {
		s.D.Set("upgrade_information", nil)
	}

	return nil
}
