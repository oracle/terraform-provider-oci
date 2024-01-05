// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVnicAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVnicAttachment,
		Read:     readCoreVnicAttachment,
		Update:   updateCoreVnicAttachment,
		Delete:   deleteCoreVnicAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"create_vnic_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"assign_ipv6ip": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"assign_private_dns_record": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"assign_public_ip": {
							// Change type from boolean to string because TF doesn't handle default
							// values for boolean nested objects correctly.
							Type:     schema.TypeString,
							Optional: true,
							Default:  "false",
							ForceNew: true,
							ValidateFunc: func(v interface{}, k string) ([]string, []error) {
								// Verify that we can parse the string value as a bool value.
								var es []error
								if _, err := strconv.ParseBool(v.(string)); err != nil {
									es = append(es, fmt.Errorf("%s: cannot parse 'assign_public_ip' as bool: %v", k, err))
								}
								return nil, es
							},
							StateFunc: func(v interface{}) string {
								// ValidateFunc runs before StateFunc. Must be valid by now.
								b, _ := tfresource.NormalizeBoolString(v.(string))
								return b
							},
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
						},
						"display_name": {
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
						"hostname_label": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6address_ipv6subnet_cidr_pair_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_subnet_cidr": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipv6_address": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"skip_source_dest_check": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nic_index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_tag": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVnicAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVnicAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVnicAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	Res                    *oci_core.VnicAttachment
	DisableNotFoundRetries bool
}

func (s *CoreVnicAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVnicAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateAttaching),
	}
}

func (s *CoreVnicAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateAttached),
	}
}

func (s *CoreVnicAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateDetaching),
	}
}

func (s *CoreVnicAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VnicAttachmentLifecycleStateDetached),
	}
}

func (s *CoreVnicAttachmentResourceCrud) Create() error {
	request := oci_core.AttachVnicRequest{}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_vnic_details", 0)
			tmp, err := s.mapToCreateVnicDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CreateVnicDetails = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if nicIndex, ok := s.D.GetOkExists("nic_index"); ok {
		tmp := nicIndex.(int)
		request.NicIndex = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AttachVnic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VnicAttachment
	return nil
}

// @CODEGEN 1/2018: Generator doesn't give us an Update method for VnicAttachment.
// However, the existing behavior allows vnic to be updated through the create_vnic_details.
// So keep this Update functionality in the provider.
func (s *CoreVnicAttachmentResourceCrud) Update() error {
	// We should fetch the VnicAttachment in order to Update
	// the state data after the Update call.
	err := s.Get()
	if err != nil {
		return err
	}

	request := oci_core.UpdateVnicRequest{}

	if s.Res.VnicId != nil {
		request.VnicId = s.Res.VnicId
	}

	if !s.D.HasChange("create_vnic_details") {
		return nil
	}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_vnic_details", 0)
			tmp, err := s.mapToUpdateVnicDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpdateVnicDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err = s.VirtualNetworkClient.UpdateVnic(context.Background(), request)
	return err
}

func (s *CoreVnicAttachmentResourceCrud) Get() error {
	request := oci_core.GetVnicAttachmentRequest{}

	tmp := s.D.Id()
	request.VnicAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVnicAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VnicAttachment
	return nil
}

func (s *CoreVnicAttachmentResourceCrud) Delete() error {
	request := oci_core.DetachVnicRequest{}

	tmp := s.D.Id()
	request.VnicAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DetachVnic(context.Background(), request)
	return err
}

func (s *CoreVnicAttachmentResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	if s.Res.NicIndex != nil {
		s.D.Set("nic_index", *s.Res.NicIndex)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VlanId != nil {
		s.D.Set("vlan_id", *s.Res.VlanId)
	}

	if s.Res.VlanTag != nil {
		s.D.Set("vlan_tag", *s.Res.VlanTag)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	// @CODEGEN 1/2018: We need to refresh the vnic details after every refresh.
	request := oci_core.GetVnicRequest{}
	request.VnicId = s.Res.VnicId

	response, err := s.VirtualNetworkClient.GetVnic(context.Background(), request)
	if err != nil {
		// VNIC might not be found when attaching or detaching; or if it is in a different compartment
		if request.VnicId != nil {
			log.Printf("[DEBUG] VNIC not found during VNIC Attachment refresh. (VNIC ID: %q, Error: %q)", *request.VnicId, err)
		}
		return nil
	}

	var createVnicDetails map[string]interface{}
	if details, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			createVnicDetails = tmpList[0].(map[string]interface{})
		}
	}

	if err := s.D.Set("create_vnic_details", []interface{}{VnicDetailsToMap(&response.Vnic, createVnicDetails, false)}); err != nil {
		log.Printf("Unable to refresh create_vnic_details. Error: %q", err)
	}

	return nil
}

func (s *CoreVnicAttachmentResourceCrud) mapToCreateVnicDetails(fieldKeyFormat string) (oci_core.CreateVnicDetails, error) {
	result := oci_core.CreateVnicDetails{}

	if assignIpv6Ip, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_ipv6ip")); ok {
		tmp := assignIpv6Ip.(bool)
		result.AssignIpv6Ip = &tmp
	}

	if assignPrivateDnsRecord, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_private_dns_record")); ok {
		tmp := assignPrivateDnsRecord.(bool)
		result.AssignPrivateDnsRecord = &tmp
	}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(string)
		boolVal, err := strconv.ParseBool(tmp)
		if err != nil {
			return result, err
		}
		result.AssignPublicIp = &boolVal
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if ipv6AddressIpv6SubnetCidrPairDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")); ok {
		interfaces := ipv6AddressIpv6SubnetCidrPairDetails.([]interface{})
		tmp := make([]oci_core.Ipv6AddressIpv6SubnetCidrPairDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details"), stateDataIndex)
			converted, err := s.mapToIpv6AddressIpv6SubnetCidrPairDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ipv6address_ipv6subnet_cidr_pair_details")) {
			result.Ipv6AddressIpv6SubnetCidrPairDetails = tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
		tmp := vlanId.(string)
		result.VlanId = &tmp
	}

	return result, nil
}

func (s *CoreVnicAttachmentResourceCrud) mapToUpdateVnicDetails(fieldKeyFormat string) (oci_core.UpdateVnicDetails, error) {
	result := oci_core.UpdateVnicDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok && hostnameLabel != "" {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); s.Res.VlanId == nil && ok {
		result.NsgIds = []string{}
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.NsgIds = tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); s.Res.VlanId == nil && ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	return result, nil
}

func VnicDetailsToMap(obj *oci_core.Vnic, createVnicDetails map[string]interface{}, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if createVnicDetails != nil {
		result["assign_ipv6ip"] = createVnicDetails["assign_ipv6ip"]
	}

	if createVnicDetails != nil {
		result["assign_private_dns_record"] = createVnicDetails["assign_private_dns_record"]
	}
	// "assign_public_ip" isn't part of the VNIC's state & is only useful at creation time (and
	// subsequent force-new creations). So persist the user-defined value in the config & Update it
	// when the user changes that value.
	if createVnicDetails != nil {
		assignPublicIP, _ := tfresource.NormalizeBoolString(createVnicDetails["assign_public_ip"].(string)) // Must be valid.
		result["assign_public_ip"] = assignPublicIP
	} else {
		// We may be importing this value; so let's set it to whether the public IP is set.
		result["assign_public_ip"] = strconv.FormatBool(obj.PublicIp != nil && *obj.PublicIp != "")
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	if createVnicDetails != nil {
		ipv6AddressIpv6SubnetCidrPairDetails := []interface{}{}
		for _, item := range createVnicDetails["ipv6address_ipv6subnet_cidr_pair_details"].([]interface{}) {
			ipv6AddressIpv6SubnetCidrPairDetails = append(ipv6AddressIpv6SubnetCidrPairDetails, item)
		}
		result["ipv6address_ipv6subnet_cidr_pair_details"] = ipv6AddressIpv6SubnetCidrPairDetails
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.SkipSourceDestCheck != nil {
		result["skip_source_dest_check"] = bool(*obj.SkipSourceDestCheck)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.VlanId != nil {
		result["vlan_id"] = string(*obj.VlanId)
	}

	return result
}

func (s *CoreVnicAttachmentResourceCrud) mapToIpv6AddressIpv6SubnetCidrPairDetails(fieldKeyFormat string) (oci_core.Ipv6AddressIpv6SubnetCidrPairDetails, error) {
	result := oci_core.Ipv6AddressIpv6SubnetCidrPairDetails{}

	if ipv6SubnetCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6_subnet_cidr")); ok {
		tmp := ipv6SubnetCidr.(string)
		result.Ipv6SubnetCidr = &tmp
	}

	if ipv6Address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6_address")); ok {
		tmp := ipv6Address.(string)
		result.Ipv6Address = &tmp
	}

	return result, nil
}

func Ipv6AddressIpv6SubnetCidrPairDetailsToMap(obj oci_core.Ipv6AddressIpv6SubnetCidrPairDetails) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}
