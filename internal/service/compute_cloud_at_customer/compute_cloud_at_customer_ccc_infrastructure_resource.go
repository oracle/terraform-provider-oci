// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package compute_cloud_at_customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ComputeCloudAtCustomerCccInfrastructureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createComputeCloudAtCustomerCccInfrastructure,
		Read:     readComputeCloudAtCustomerCccInfrastructure,
		Update:   updateComputeCloudAtCustomerCccInfrastructure,
		Delete:   deleteComputeCloudAtCustomerCccInfrastructure,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"ccc_upgrade_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_details": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"infrastructure_inventory": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"capacity_storage_tray_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"compute_node_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"management_node_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"performance_storage_tray_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"infrastructure_network_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"dns_ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"infrastructure_routing_dynamic": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"bgp_topology": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_asn": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"peer_information": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"asn": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"ip": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"infrastructure_routing_static": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"uplink_hsrp_group": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"uplink_vlan": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"management_nodes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"mgmt_vip_hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mgmt_vip_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spine_ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"spine_vip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uplink_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uplink_gateway_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uplink_netmask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uplink_port_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uplink_port_forward_error_correction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uplink_port_speed_in_gbps": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uplink_vlan_mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioning_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioning_pin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"short_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_information": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"current_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"scheduled_upgrade_duration": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_scheduled_upgrade": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createComputeCloudAtCustomerCccInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccInfrastructureResourceCrud{}
	// d.Set("display_name", "terraform-test-infra")
	// d.Set("description", "This infrastructure was created by terraform-provider")
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.CreateResource(d, sync)
}

func readComputeCloudAtCustomerCccInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.ReadResource(sync)
}

func updateComputeCloudAtCustomerCccInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteComputeCloudAtCustomerCccInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ComputeCloudAtCustomerCccInfrastructureResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient
	Res                    *oci_compute_cloud_at_customer.CccInfrastructure
	DisableNotFoundRetries bool
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateActive),
		string(oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateNeedsAttention),
	}
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_compute_cloud_at_customer.CccInfrastructureLifecycleStateDeleted),
	}
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) Create() error {
	request := oci_compute_cloud_at_customer.CreateCccInfrastructureRequest{}

	if cccUpgradeScheduleId, ok := s.D.GetOkExists("ccc_upgrade_schedule_id"); ok {
		tmp := cccUpgradeScheduleId.(string)
		request.CccUpgradeScheduleId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
		tmp := connectionDetails.(string)
		request.ConnectionDetails = &tmp
	}

	if connectionState, ok := s.D.GetOkExists("connection_state"); ok {
		request.ConnectionState = oci_compute_cloud_at_customer.CccInfrastructureConnectionStateEnum(connectionState.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	response, err := s.Client.CreateCccInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CccInfrastructure
	return nil
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) Get() error {
	request := oci_compute_cloud_at_customer.GetCccInfrastructureRequest{}

	tmp := s.D.Id()
	request.CccInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	response, err := s.Client.GetCccInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CccInfrastructure
	return nil
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_compute_cloud_at_customer.UpdateCccInfrastructureRequest{}

	tmp := s.D.Id()
	request.CccInfrastructureId = &tmp

	if cccUpgradeScheduleId, ok := s.D.GetOkExists("ccc_upgrade_schedule_id"); ok {
		tmp := cccUpgradeScheduleId.(string)
		request.CccUpgradeScheduleId = &tmp
	}

	if connectionDetails, ok := s.D.GetOkExists("connection_details"); ok {
		tmp := connectionDetails.(string)
		request.ConnectionDetails = &tmp
	}

	if connectionState, ok := s.D.GetOkExists("connection_state"); ok {
		request.ConnectionState = oci_compute_cloud_at_customer.CccInfrastructureConnectionStateEnum(connectionState.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	response, err := s.Client.UpdateCccInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CccInfrastructure
	return nil
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) Delete() error {
	request := oci_compute_cloud_at_customer.DeleteCccInfrastructureRequest{}

	tmp := s.D.Id()
	request.CccInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	_, err := s.Client.DeleteCccInfrastructure(context.Background(), request)
	return err
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) SetData() error {
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

func CccInfrastructureInventoryToMap(obj *oci_compute_cloud_at_customer.CccInfrastructureInventory) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CapacityStorageTrayCount != nil {
		result["capacity_storage_tray_count"] = int(*obj.CapacityStorageTrayCount)
	}

	if obj.ComputeNodeCount != nil {
		result["compute_node_count"] = int(*obj.ComputeNodeCount)
	}

	if obj.ManagementNodeCount != nil {
		result["management_node_count"] = int(*obj.ManagementNodeCount)
	}

	if obj.PerformanceStorageTrayCount != nil {
		result["performance_storage_tray_count"] = int(*obj.PerformanceStorageTrayCount)
	}

	if obj.SerialNumber != nil {
		result["serial_number"] = string(*obj.SerialNumber)
	}

	return result
}

func CccInfrastructureManagementNodeToMap(obj oci_compute_cloud_at_customer.CccInfrastructureManagementNode) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Ip != nil {
		result["ip"] = string(*obj.Ip)
	}

	return result
}

func CccInfrastructureNetworkConfigurationToMap(obj *oci_compute_cloud_at_customer.CccInfrastructureNetworkConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["dns_ips"] = obj.DnsIps

	if obj.InfrastructureRoutingDynamic != nil {
		result["infrastructure_routing_dynamic"] = []interface{}{CccInfrastructureRoutingDynamicDetailsToMap(obj.InfrastructureRoutingDynamic)}
	}

	if obj.InfrastructureRoutingStatic != nil {
		result["infrastructure_routing_static"] = []interface{}{CccInfrastructureRoutingStaticDetailsToMap(obj.InfrastructureRoutingStatic)}
	}

	managementNodes := []interface{}{}
	for _, item := range obj.ManagementNodes {
		managementNodes = append(managementNodes, CccInfrastructureManagementNodeToMap(item))
	}
	result["management_nodes"] = managementNodes

	if obj.MgmtVipHostname != nil {
		result["mgmt_vip_hostname"] = string(*obj.MgmtVipHostname)
	}

	if obj.MgmtVipIp != nil {
		result["mgmt_vip_ip"] = string(*obj.MgmtVipIp)
	}

	result["spine_ips"] = obj.SpineIps

	if obj.SpineVip != nil {
		result["spine_vip"] = string(*obj.SpineVip)
	}

	if obj.UplinkDomain != nil {
		result["uplink_domain"] = string(*obj.UplinkDomain)
	}

	if obj.UplinkGatewayIp != nil {
		result["uplink_gateway_ip"] = string(*obj.UplinkGatewayIp)
	}

	if obj.UplinkNetmask != nil {
		result["uplink_netmask"] = string(*obj.UplinkNetmask)
	}

	if obj.UplinkPortCount != nil {
		result["uplink_port_count"] = int(*obj.UplinkPortCount)
	}

	result["uplink_port_forward_error_correction"] = string(obj.UplinkPortForwardErrorCorrection)

	if obj.UplinkPortSpeedInGbps != nil {
		result["uplink_port_speed_in_gbps"] = int(*obj.UplinkPortSpeedInGbps)
	}

	if obj.UplinkVlanMtu != nil {
		result["uplink_vlan_mtu"] = int(*obj.UplinkVlanMtu)
	}

	return result
}

func CccInfrastructureRoutingDynamicDetailsToMap(obj *oci_compute_cloud_at_customer.CccInfrastructureRoutingDynamicDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["bgp_topology"] = string(obj.BgpTopology)

	if obj.OracleAsn != nil {
		result["oracle_asn"] = int(*obj.OracleAsn)
	}

	peerInformation := []interface{}{}
	for _, item := range obj.PeerInformation {
		peerInformation = append(peerInformation, PeerInformationToMap(item))
	}
	result["peer_information"] = peerInformation

	return result
}

func CccInfrastructureRoutingStaticDetailsToMap(obj *oci_compute_cloud_at_customer.CccInfrastructureRoutingStaticDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.UplinkHsrpGroup != nil {
		result["uplink_hsrp_group"] = int(*obj.UplinkHsrpGroup)
	}

	if obj.UplinkVlan != nil {
		result["uplink_vlan"] = int(*obj.UplinkVlan)
	}

	return result
}

func CccInfrastructureSummaryToMap(obj oci_compute_cloud_at_customer.CccInfrastructureSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["connection_state"] = string(obj.ConnectionState)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ShortName != nil {
		result["short_name"] = string(*obj.ShortName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func CccUpgradeInformationToMap(obj *oci_compute_cloud_at_customer.CccUpgradeInformation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CurrentVersion != nil {
		result["current_version"] = string(*obj.CurrentVersion)
	}

	if obj.IsActive != nil {
		result["is_active"] = bool(*obj.IsActive)
	}

	if obj.ScheduledUpgradeDuration != nil {
		result["scheduled_upgrade_duration"] = string(*obj.ScheduledUpgradeDuration)
	}

	if obj.TimeOfScheduledUpgrade != nil {
		result["time_of_scheduled_upgrade"] = obj.TimeOfScheduledUpgrade.String()
	}

	return result
}

func PeerInformationToMap(obj oci_compute_cloud_at_customer.PeerInformation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Asn != nil {
		result["asn"] = int(*obj.Asn)
	}

	if obj.Ip != nil {
		result["ip"] = string(*obj.Ip)
	}

	return result
}

func (s *ComputeCloudAtCustomerCccInfrastructureResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_compute_cloud_at_customer.ChangeCccInfrastructureCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CccInfrastructureId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	_, err := s.Client.ChangeCccInfrastructureCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
