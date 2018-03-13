// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"errors"
	"log"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &crud.TwoHours,
			Update: &crud.TwoHours,
			Delete: &crud.TwoHours,
		},
		Create: createInstance,
		Read:   readInstance,
		Update: updateInstance,
		Delete: deleteInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"create_vnic_details": {
				Type:     schema.TypeList,
				Optional: true,
				// This must be set to computed, since it's optional and required subnet_id param is being refreshed.
				// If this isn't computed, then that would always force a change on users who do not set create_vnic_details.
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"assign_public_ip": {
							Type:     schema.TypeBool,
							Optional: true,
							// @CODEGEN 1/2018: Computed cant be used with legacy default value
							ForceNew: true,
							Default:  true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							// @CODEGEN 1/2018: Remove ForceNew, this is updatable via vnic update
						},
						"hostname_label": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
							// @CODEGEN 1/2018: Remove ForceNew, this is updatable via vnic update
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
							// @CODEGEN 1/2018: Remove ForceNew, this is updatable via vnic update
						},

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"hostname_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true, // Changed from optional/computed to required till "sourceDetails" is supported.
				ForceNew: true,
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			// @CODEGEN 1/2018: source_details currently outside parity scope
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
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
			// Legacy custom computed convenience values
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteInstance(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type InstanceResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.ComputeClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	Res                    *oci_core.Instance
	DisableNotFoundRetries bool
}

func (s *InstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *InstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateProvisioning),
		string(oci_core.InstanceLifecycleStateStarting),
	}
}

func (s *InstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateRunning),
	}
}

func (s *InstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateTerminating),
	}
}

func (s *InstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateTerminated),
	}
}

func (s *InstanceResourceCrud) Create() error {
	request := oci_core.LaunchInstanceRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			tmp := mapToCreateVnicDetailsInstance(tmpList[0].(map[string]interface{}))
			request.CreateVnicDetails = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if rawExtendedMetadata, ok := s.D.GetOkExists("extended_metadata"); ok {
		extendedMetadata := mapToExtendedMetadata(rawExtendedMetadata.(map[string]interface{}))
		request.ExtendedMetadata = extendedMetadata
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	// @CODEGEN 1/2018: support legacy name "image"
	if imageId, ok := s.D.GetOkExists("image"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if ipxeScript, ok := s.D.GetOkExists("ipxe_script"); ok {
		tmp := ipxeScript.(string)
		request.IpxeScript = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := resourceInstanceMapToMetadata(metadata.(map[string]interface{}))
		request.Metadata = tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	// @CODEGEN 1/2018: source_details currently outside parity scope

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.LaunchInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *InstanceResourceCrud) Get() error {
	request := oci_core.GetInstanceRequest{}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *InstanceResourceCrud) Update() error {
	request := oci_core.UpdateInstanceRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance

	// Check for changes in the create_vnic_details sub resource and separately update the vnic

	rawVnics, ok := s.D.GetOkExists("create_vnic_details")
	if !s.D.HasChange("create_vnic_details") || !ok {
		log.Printf("[DEBUG] No changes to primary VNIC. Instance ID: %q", s.Res.Id)
		return nil
	}

	rawVnic := rawVnics.([]interface{})[0].(map[string]interface{})

	vnic, err := s.getPrimaryVnic()
	if err != nil {
		log.Printf("[ERROR] Primary VNIC could not be found during instance update: %q (Instance ID: %q, State: %q)", err, s.Res.Id, s.Res.LifecycleState)
		return err
	}

	vnicOpts := oci_core.UpdateVnicRequest{
		VnicId:            vnic.Id,
		UpdateVnicDetails: mapToUpdateVnicDetailsInstance(rawVnic),
	}

	_, err = s.VirtualNetworkClient.UpdateVnic(context.Background(), vnicOpts)

	if err != nil {
		log.Printf("[ERROR] Primary VNIC could not be updated during instance update: %q (Instance ID: %q, State: %q)", err, s.Res.Id, s.Res.LifecycleState)
		return err
	}

	return nil
}

func (s *InstanceResourceCrud) Delete() error {
	request := oci_core.TerminateInstanceRequest{}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	if preserveBootVolume, ok := s.D.GetOkExists("preserve_boot_volume"); ok {
		tmp := preserveBootVolume.(bool)
		request.PreserveBootVolume = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.TerminateInstance(context.Background(), request)
	return err
}

func (s *InstanceResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	// Extended metadata (a json blob) may not return with the same node order in which it
	// was originally created, the solution is to not set it here after subsequent GETS to
	// prevent inadvertent diffs or destroy/creates
	// if s.Res.ExtendedMetadata != nil {
	// // extended_metadata is an arbitrarily structured json object, `objectToMap` would not work
	// 	s.D.Set("extended_metadata", []interface{}{objectToMap(s.Res.ExtendedMetadata)})
	// }

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.ImageId != nil {
		// @CODEGEN 1/2018: support legacy name "image"
		s.D.Set("image", *s.Res.ImageId)
	}

	if s.Res.IpxeScript != nil {
		s.D.Set("ipxe_script", *s.Res.IpxeScript)
	}

	if s.Res.Metadata != nil {
		err := s.D.Set("metadata", s.Res.Metadata)
		if err != nil {
			log.Printf("error setting metadata %q", err)
		}
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	// @CODEGEN 1/2018: source_details currently outside parity scope

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())

	if s.Res.LifecycleState != oci_core.InstanceLifecycleStateRunning {
		return
	}

	vnic, vnicError := s.getPrimaryVnic()
	if vnicError != nil || vnic == nil {
		log.Printf("[WARN] Primary VNIC could not be found during instance refresh: %q", vnicError)
		return
	}

	s.D.Set("hostname_label", vnic.HostnameLabel)
	s.D.Set("public_ip", vnic.PublicIp)
	s.D.Set("private_ip", vnic.PrivateIp)
	s.D.Set("subnet_id", vnic.SubnetId)

	err := s.D.Set("create_vnic_details", []interface{}{vnicDetailsToMap(vnic)})
	if err != nil {
		log.Printf("[WARN] create_vnic_details could not be set: %q", err)
	}
}

func mapToCreateVnicDetailsInstance(raw map[string]interface{}) oci_core.CreateVnicDetails {
	result := oci_core.CreateVnicDetails{}

	if assignPublicIp, ok := raw["assign_public_ip"]; ok {
		tmp := assignPublicIp.(bool)
		result.AssignPublicIp = &tmp
	}

	if displayName, ok := raw["display_name"]; ok {
		tmp := displayName.(string)
		if tmp != "" {
			result.DisplayName = &tmp
		}
	}

	if hostnameLabel, ok := raw["hostname_label"]; ok {
		tmp := hostnameLabel.(string)
		if tmp != "" {
			result.HostnameLabel = &tmp
		}
	}

	if privateIp, ok := raw["private_ip"]; ok {
		tmp := privateIp.(string)
		if tmp != "" {
			result.PrivateIp = &tmp
		}
	}

	if skipSourceDestCheck, ok := raw["skip_source_dest_check"]; ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := raw["subnet_id"]; ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result
}

func mapToUpdateVnicDetailsInstance(raw map[string]interface{}) oci_core.UpdateVnicDetails {
	result := oci_core.UpdateVnicDetails{}

	if displayName, ok := raw["display_name"]; ok {
		tmp := displayName.(string)
		if tmp != "" {
			result.DisplayName = &tmp
		}
	}

	if hostnameLabel, ok := raw["hostname_label"]; ok {
		tmp := hostnameLabel.(string)
		if tmp != "" {
			result.HostnameLabel = &tmp
		}
	}

	if skipSourceDestCheck, ok := raw["skip_source_dest_check"]; ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	return result
}

func vnicDetailsToMap(obj *oci_core.Vnic) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PublicIp != nil {
		result["assign_public_ip"] = bool(len(*obj.PublicIp) > 0)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
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

	return result
}

func resourceInstanceMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func mapToExtendedMetadata(rm map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range rm {
		val := make(map[string]interface{})
		//Use the string value that was passed if it is not a valid JSON string
		if err := json.Unmarshal([]byte(v.(string)), &val); err == nil {
			result[k] = val
		} else {
			result[k] = v.(string)
		}
	}
	return result
}

func (s *InstanceResourceCrud) getPrimaryVnic() (*oci_core.Vnic, error) {
	request := oci_core.ListVnicAttachmentsRequest{
		CompartmentId: s.Res.CompartmentId,
		InstanceId:    s.Res.Id,
	}

	var attachments []oci_core.VnicAttachment

	for {
		result, err := s.Client.ListVnicAttachments(context.Background(), request)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, result.Items...)
		request.Page = result.OpcNextPage

		if request.Page == nil {
			break
		}
	}

	if len(attachments) < 1 {
		return nil, errors.New("No VNIC attachments found.")
	}

	for _, attachment := range attachments {
		if attachment.LifecycleState == oci_core.VnicAttachmentLifecycleStateAttached {
			request := oci_core.GetVnicRequest{VnicId: attachment.VnicId}
			response, _ := s.VirtualNetworkClient.GetVnic(context.Background(), request)
			vnic := &response.Vnic

			// Ignore errors on GetVnic, since we might not have permissions to view some secondary VNICs.
			if vnic != nil && vnic.IsPrimary != nil && *vnic.IsPrimary {
				return vnic, nil
			}
		}
	}

	return nil, errors.New("Primary VNIC not found.")
}
