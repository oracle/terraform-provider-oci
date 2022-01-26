// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseVmClusterNetworkResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseVmClusterNetwork,
		Read:     readDatabaseVmClusterNetwork,
		Update:   updateDatabaseVmClusterNetwork,
		Delete:   deleteDatabaseVmClusterNetwork,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scans": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      scansHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ips": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"scan_listener_port_tcp": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"scan_listener_port_tcp_ssl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"vm_networks": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      vmNetworksHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"domain_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"gateway": {
							Type:     schema.TypeString,
							Required: true,
						},
						"netmask": {
							Type:     schema.TypeString,
							Required: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"nodes": {
							Type:     schema.TypeSet,
							Required: true,
							Set:      nodesHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"hostname": {
										Type:     schema.TypeString,
										Required: true,
									},
									"ip": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"vip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vip_hostname": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dns": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ntp": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"validate_vm_cluster_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseVmClusterNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseVmClusterNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseVmClusterNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseVmClusterNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterNetworkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseVmClusterNetworkResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.VmClusterNetwork
	DisableNotFoundRetries bool
}

func (s *DatabaseVmClusterNetworkResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseVmClusterNetworkResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.VmClusterNetworkLifecycleStateCreating),
		string(oci_database.VmClusterNetworkLifecycleStateValidating),
	}
}

func (s *DatabaseVmClusterNetworkResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.VmClusterNetworkLifecycleStateRequiresValidation),
		string(oci_database.VmClusterNetworkLifecycleStateValidated),
	}
}

func (s *DatabaseVmClusterNetworkResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.VmClusterNetworkLifecycleStateTerminating),
	}
}

func (s *DatabaseVmClusterNetworkResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.VmClusterNetworkLifecycleStateTerminated),
	}
}

func (s *DatabaseVmClusterNetworkResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.VmClusterNetworkLifecycleStateValidating),
		string(oci_database.VmClusterNetworkLifecycleStateUpdating),
	}
}

func (s *DatabaseVmClusterNetworkResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.VmClusterNetworkLifecycleStateRequiresValidation),
		string(oci_database.VmClusterNetworkLifecycleStateValidated),
		string(oci_database.VmClusterNetworkLifecycleStateValidationFailed),
		string(oci_database.VmClusterNetworkLifecycleStateAllocated),
	}
}

func (s *DatabaseVmClusterNetworkResourceCrud) Create() error {
	request := oci_database.CreateVmClusterNetworkRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dns, ok := s.D.GetOkExists("dns"); ok {
		request.Dns = []string{}
		interfaces := dns.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns") {
			request.Dns = tmp
		}
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ntp, ok := s.D.GetOkExists("ntp"); ok {
		request.Ntp = []string{}
		interfaces := ntp.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp") {
			request.Ntp = tmp
		}
	}

	if scans, ok := s.D.GetOkExists("scans"); ok {
		request.Scans = []oci_database.ScanDetails{}
		set := scans.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.ScanDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := scansHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scans", stateDataIndex)
			converted, err := s.mapToScanDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scans") {
			request.Scans = tmp
		}
	}

	if vmNetworks, ok := s.D.GetOkExists("vm_networks"); ok {
		request.VmNetworks = []oci_database.VmNetworkDetails{}
		set := vmNetworks.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.VmNetworkDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := vmNetworksHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vm_networks", stateDataIndex)
			log.Printf("vm_networks vmNetworksHashCodeForSets %s", fieldKeyFormat)
			converted, err := s.mapToVmNetworkDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vm_networks") {
			request.VmNetworks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateVmClusterNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmClusterNetwork

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if validateVmClusterNetwork, ok := s.D.GetOkExists("validate_vm_cluster_network"); ok &&
		validateVmClusterNetwork.(bool) == true {
		response, err := s.validateVmClusterNetwork(*s.Res.Id, *s.Res.ExadataInfrastructureId)
		if err != nil {
			s.D.Set("validate_vm_cluster_network", false)
			return err
		}
		s.Res = &response.VmClusterNetwork
	}

	return nil
}

func (s *DatabaseVmClusterNetworkResourceCrud) Get() error {
	request := oci_database.GetVmClusterNetworkRequest{}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	tmp := s.D.Id()
	request.VmClusterNetworkId = &tmp

	exadataInfrastructureId, vmClusterNetworkId, err := parseVmClusterNetworkCompositeId(s.D.Id())
	if err == nil {
		request.ExadataInfrastructureId = &exadataInfrastructureId
		request.VmClusterNetworkId = &vmClusterNetworkId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetVmClusterNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmClusterNetwork
	return nil
}

func (s *DatabaseVmClusterNetworkResourceCrud) Update() error {

	if s.D.Get("state").(string) == string(oci_database.VmClusterNetworkLifecycleStateValidated) ||
		s.D.Get("state").(string) == string(oci_database.VmClusterNetworkLifecycleStateAllocated) {
		return fmt.Errorf("Update not allowed on validated vm cluster network")
	}

	request := oci_database.UpdateVmClusterNetworkRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if dns, ok := s.D.GetOkExists("dns"); ok {
		request.Dns = []string{}
		interfaces := dns.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("dns") {
			request.Dns = tmp
		}
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ntp, ok := s.D.GetOkExists("ntp"); ok && s.D.HasChange("ntp") {
		request.Ntp = []string{}
		interfaces := ntp.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ntp") {
			request.Ntp = tmp
		}
	}

	if scans, ok := s.D.GetOkExists("scans"); ok {
		request.Scans = []oci_database.ScanDetails{}
		set := scans.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.ScanDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := scansHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scans", stateDataIndex)
			converted, err := s.mapToScanDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scans") {
			request.Scans = tmp
		}
	}

	tmp := s.D.Id()
	request.VmClusterNetworkId = &tmp

	if vmNetworks, ok := s.D.GetOkExists("vm_networks"); ok {
		request.VmNetworks = []oci_database.VmNetworkDetails{}
		set := vmNetworks.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.VmNetworkDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := vmNetworksHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vm_networks", stateDataIndex)
			converted, err := s.mapToVmNetworkDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vm_networks") {
			request.VmNetworks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateVmClusterNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VmClusterNetwork

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if validateVmClusterNetwork, ok := s.D.GetOkExists("validate_vm_cluster_network"); ok &&
		validateVmClusterNetwork.(bool) == true {
		response, err := s.validateVmClusterNetwork(*s.Res.Id, *s.Res.ExadataInfrastructureId)
		if err != nil {
			s.D.Set("validate_vm_cluster_network", false)
			return err
		}
		s.Res = &response.VmClusterNetwork
	}

	return nil
}

func (s *DatabaseVmClusterNetworkResourceCrud) Delete() error {
	request := oci_database.DeleteVmClusterNetworkRequest{}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	tmp := s.D.Id()
	request.VmClusterNetworkId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteVmClusterNetwork(context.Background(), request)
	return err
}

func (s *DatabaseVmClusterNetworkResourceCrud) SetData() error {

	if s.Res.LifecycleState == oci_database.VmClusterNetworkLifecycleStateRequiresValidation ||
		s.Res.LifecycleState == oci_database.VmClusterNetworkLifecycleStateValidationFailed {
		s.D.Set("validate_vm_cluster_network", false)
	}

	exadataInfrastructureId, vmClusterNetworkId, err := parseVmClusterNetworkCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("exadata_infrastructure_id", exadataInfrastructureId)
		s.D.SetId(vmClusterNetworkId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns", s.Res.Dns)

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("ntp", s.Res.Ntp)

	scans := []interface{}{}
	for _, item := range s.Res.Scans {
		scans = append(scans, ScanDetailsToMap(item))
	}
	s.D.Set("scans", schema.NewSet(scansHashCodeForSets, scans))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	vmNetworks := []interface{}{}
	for _, item := range s.Res.VmNetworks {
		vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item, false))
	}
	s.D.Set("vm_networks", schema.NewSet(vmNetworksHashCodeForSets, vmNetworks))

	return nil
}

func GetVmClusterNetworkCompositeId(exadataInfrastructureId string, vmClusterNetworkId string) string {
	exadataInfrastructureId = url.PathEscape(exadataInfrastructureId)
	vmClusterNetworkId = url.PathEscape(vmClusterNetworkId)
	compositeId := "exadataInfrastructures/" + exadataInfrastructureId + "/vmClusterNetworks/" + vmClusterNetworkId
	return compositeId
}

func parseVmClusterNetworkCompositeId(compositeId string) (exadataInfrastructureId string, vmClusterNetworkId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("exadataInfrastructures/.*/vmClusterNetworks/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	exadataInfrastructureId, _ = url.PathUnescape(parts[1])
	vmClusterNetworkId, _ = url.PathUnescape(parts[3])

	return
}

func (s *DatabaseVmClusterNetworkResourceCrud) mapToNodeDetails(fieldKeyFormat string) (oci_database.NodeDetails, error) {
	result := oci_database.NodeDetails{}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	if ip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip")); ok {
		tmp := ip.(string)
		result.Ip = &tmp
	}

	if vip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vip")); ok && vip != "" {
		tmp := vip.(string)
		result.Vip = &tmp
	}

	if vipHostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vip_hostname")); ok && vipHostname != "" {
		tmp := vipHostname.(string)
		result.VipHostname = &tmp
	}

	return result, nil
}

func NodeDetailsToMap(obj oci_database.NodeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Ip != nil {
		result["ip"] = string(*obj.Ip)
	}

	if obj.Vip != nil {
		result["vip"] = string(*obj.Vip)
	}

	if obj.VipHostname != nil {
		result["vip_hostname"] = string(*obj.VipHostname)
	}

	return result
}

func (s *DatabaseVmClusterNetworkResourceCrud) mapToScanDetails(fieldKeyFormat string) (oci_database.ScanDetails, error) {
	result := oci_database.ScanDetails{}

	if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
		tmp := hostname.(string)
		result.Hostname = &tmp
	}

	if ips, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ips")); ok {
		interfaces := ips.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ips")) {
			result.Ips = tmp
		}
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if scanListenerPortTcp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_listener_port_tcp")); ok {
		tmp := scanListenerPortTcp.(int)
		result.ScanListenerPortTcp = &tmp
	}

	if scanListenerPortTcpSsl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_listener_port_tcp_ssl")); ok {
		tmp := scanListenerPortTcpSsl.(int)
		result.ScanListenerPortTcpSsl = &tmp
	}

	return result, nil
}

func ScanDetailsToMap(obj oci_database.ScanDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	result["ips"] = obj.Ips

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.ScanListenerPortTcp != nil {
		result["scan_listener_port_tcp"] = int(*obj.ScanListenerPortTcp)
	}

	if obj.ScanListenerPortTcpSsl != nil {
		result["scan_listener_port_tcp_ssl"] = int(*obj.ScanListenerPortTcpSsl)
	}

	return result
}

func (s *DatabaseVmClusterNetworkResourceCrud) mapToVmNetworkDetails(fieldKeyFormat string) (oci_database.VmNetworkDetails, error) {
	result := oci_database.VmNetworkDetails{}

	if domainName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_name")); ok {
		tmp := domainName.(string)
		result.DomainName = &tmp
	}

	if gateway, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gateway")); ok {
		tmp := gateway.(string)
		result.Gateway = &tmp
	}

	if netmask, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "netmask")); ok {
		tmp := netmask.(string)
		result.Netmask = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_database.VmNetworkDetailsNetworkTypeEnum(networkType.(string))
	}

	if nodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nodes")); ok {
		set := nodes.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.NodeDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := nodesHashCodeForSets(interfaces[i])
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "nodes"), stateDataIndex)
			converted, err := s.mapToNodeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nodes")) {
			result.Nodes = tmp
		}
	}

	if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
		tmp := vlanId.(string)
		result.VlanId = &tmp
	}

	return result, nil
}

func VmNetworkDetailsToMap(obj oci_database.VmNetworkDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	if obj.Gateway != nil {
		result["gateway"] = string(*obj.Gateway)
	}

	if obj.Netmask != nil {
		result["netmask"] = string(*obj.Netmask)
	}

	result["network_type"] = string(obj.NetworkType)

	nodes := []interface{}{}
	for _, item := range obj.Nodes {
		nodes = append(nodes, NodeDetailsToMap(item))
	}
	if datasource {
		result["nodes"] = nodes
	} else {
		result["nodes"] = schema.NewSet(nodesHashCodeForSets, nodes)
	}

	if obj.VlanId != nil {
		result["vlan_id"] = string(*obj.VlanId)
	}

	return result
}

func nodesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if hostname, ok := m["hostname"]; ok && hostname != "" {
		buf.WriteString(fmt.Sprintf("%v-", hostname))
	}
	if ip, ok := m["ip"]; ok && ip != "" {
		buf.WriteString(fmt.Sprintf("%v-", ip))
	}
	if vip, ok := m["vip"]; ok && vip != "" {
		buf.WriteString(fmt.Sprintf("%v-", vip))
	}
	if vipHostname, ok := m["vip_hostname"]; ok && vipHostname != "" {
		buf.WriteString(fmt.Sprintf("%v-", vipHostname))
	}
	return hashcode.String(buf.String())
}

func scansHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if hostname, ok := m["hostname"]; ok && hostname != "" {
		buf.WriteString(fmt.Sprintf("%v-", hostname))
	}
	if ips, ok := m["ips"]; ok && ips != "" {
	}
	if port, ok := m["port"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", port))
	}
	if scanListenerPortTcp, ok := m["scan_listener_port_tcp"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", scanListenerPortTcp))
	}
	if scanListenerPortTcpSsl, ok := m["scan_listener_port_tcp_ssl"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", scanListenerPortTcpSsl))
	}
	return hashcode.String(buf.String())
}

func vmNetworksHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if domainName, ok := m["domain_name"]; ok && domainName != "" {
		buf.WriteString(fmt.Sprintf("%v-", domainName))
	}
	if gateway, ok := m["gateway"]; ok && gateway != "" {
		buf.WriteString(fmt.Sprintf("%v-", gateway))
	}
	if netmask, ok := m["netmask"]; ok && netmask != "" {
		buf.WriteString(fmt.Sprintf("%v-", netmask))
	}
	if networkType, ok := m["network_type"]; ok && networkType != "" {
		buf.WriteString(fmt.Sprintf("%v-", networkType))
	}
	if nodes, ok := m["nodes"]; ok {
		if tmpList := nodes.(*schema.Set).List(); len(tmpList) > 0 {
			buf.WriteString("node-")
			for _, nodeRaw := range tmpList {
				buf.WriteString(fmt.Sprintf("%v-", nodesHashCodeForSets(nodeRaw)))
			}
		}
	}
	if vlanId, ok := m["vlan_id"]; ok && vlanId != "" {
		buf.WriteString(fmt.Sprintf("%v-", vlanId))
	}
	return hashcode.String(buf.String())
}

func (s *DatabaseVmClusterNetworkResourceCrud) validateVmClusterNetwork(vmClusterNetworkId string, exadataInfrastructureId string) (*oci_database.ValidateVmClusterNetworkResponse, error) {
	request := oci_database.ValidateVmClusterNetworkRequest{}

	request.ExadataInfrastructureId = &exadataInfrastructureId

	request.VmClusterNetworkId = &vmClusterNetworkId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ValidateVmClusterNetwork(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
