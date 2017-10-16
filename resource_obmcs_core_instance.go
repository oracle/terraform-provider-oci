// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
	"github.com/oracle/terraform-provider-oci/options"
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
			"create_vnic_details": {
				Type:     schema.TypeList,
				Optional: true,
				// This must be set to computed, since it's optional and required subnet_id param is being refreshed.
				// If this isn't computed, then that would always force a change on users who do not set create_vnic_details.
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem:     createVnicDetailsSchema,
			},
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
				ForceNew: true,
			},
			"extended_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updateInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(d, sync)
}

func deleteInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

type InstanceResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Instance

	// Computed fields
	public_ip  string
	private_ip string
}

func (s *InstanceResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *InstanceResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceProvisioning,
		baremetal.ResourceStarting,
	}
}

func (s *InstanceResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceRunning}
}

func (s *InstanceResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *InstanceResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
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

func (s *InstanceResourceCrud) Create() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	image := s.D.Get("image").(string)
	shape := s.D.Get("shape").(string)
	subnet := s.D.Get("subnet_id").(string)

	opts := &baremetal.LaunchInstanceOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}
	if hostnameLabel, ok := s.D.GetOk("hostname_label"); ok {
		opts.HostnameLabel = hostnameLabel.(string)
	}
	if ipxeScript, ok := s.D.GetOk("ipxe_script"); ok {
		opts.IpxeScript = ipxeScript.(string)
	}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceInstanceMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}

	if rawExtendedMetadata, ok := s.D.GetOk("extended_metadata"); ok {
		extendedMetadata := mapToExtendedMetadata(rawExtendedMetadata.(map[string]interface{}))
		opts.ExtendedMetadata = extendedMetadata
	}

	if rawVnic, ok := s.D.GetOk("create_vnic_details"); ok {
		opts.CreateVnicOptions = SetCreateVnicOptions(rawVnic.([]interface{}))
	}

	if e == nil {
		s.Resource, e = s.Client.LaunchInstance(
			availabilityDomain,
			compartmentID,
			image,
			shape,
			subnet,
			opts)
	}

	return
}

/*
 * Return the primary VNIC for this instance.
 *
 * Note that this may return an error during instance creation or deletion.
 */
func (s *InstanceResourceCrud) getPrimaryVnic() (vnic *baremetal.Vnic, e error) {
	compartmentID := s.Resource.CompartmentID

	opts := &baremetal.ListVnicAttachmentsOptions{}
	opts.InstanceID = s.Resource.ID

	// Page through all VNIC attachments for the instance.
	var attachments []baremetal.VnicAttachment
	for {
		var result *baremetal.ListVnicAttachments
		if result, e = s.Client.ListVnicAttachments(compartmentID, opts); e != nil {
			break
		}

		attachments = append(attachments, result.Attachments...)
		if hasNextPage := options.SetNextPageOption(result.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	if len(attachments) < 1 {
		return nil, errors.New("No VNIC attachments found.")
	}

	for _, attachment := range attachments {
		if attachment.State == baremetal.ResourceAttached {
			vnic, _ := s.Client.GetVnic(attachment.VnicID)

			// Ignore errors on GetVnic, since we might not have permissions to view some secondary VNICs.
			if vnic != nil && vnic.IsPrimary {
				return vnic, nil
			}
		}
	}

	return nil, errors.New("Primary VNIC not found.")
}

func (s *InstanceResourceCrud) Get() (e error) {
	res, e := s.Client.GetInstance(s.D.Id())
	if e == nil {
		s.Resource = res
	}

	return
}

func (s *InstanceResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateInstance(s.D.Id(), opts)
	if e != nil {
		return
	}

	// HasChange returns true for any changes within create_vnic_details.
	if !s.D.HasChange("create_vnic_details") {
		log.Printf("[DEBUG] No changes to primary VNIC. Instance ID: %q", s.Resource.ID)
		return
	}

	log.Printf("[DEBUG] Updating instance's primary VNIC. Instance ID: %q", s.Resource.ID)
	vnic, e := s.getPrimaryVnic()
	if e != nil {
		log.Printf("[ERROR] Primary VNIC could not be found during instance update: %q (Instance ID: %q, State: %q)", e, s.Resource.ID, s.Resource.State)
		return
	}

	if rawVnic, ok := s.D.GetOk("create_vnic_details"); ok {
		_, e = s.Client.UpdateVnic(vnic.ID, SetUpdateVnicOptions(rawVnic.([]interface{})))
	}

	return
}

func (s *InstanceResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Resource.AvailabilityDomain)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("image", s.Resource.ImageID)
	s.D.Set("ipxe_script", s.Resource.IpxeScript)
	s.D.Set("metadata", s.Resource.Metadata)
	s.D.Set("region", s.Resource.Region)
	s.D.Set("shape", s.Resource.Shape)
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())

	if s.Resource.State != baremetal.ResourceRunning {
		return
	}

	vnic, vnicError := s.getPrimaryVnic()
	if vnicError != nil {
		log.Printf("[WARN] Primary VNIC could not be found during instance refresh: %q (Instance ID: %q, State: %q)", vnicError, s.Resource.ID, s.Resource.State)
		return
	}

	s.D.Set("hostname_label", vnic.HostnameLabel)
	s.D.Set("public_ip", vnic.PublicIPAddress)
	s.D.Set("private_ip", vnic.PrivateIPAddress)
	s.D.Set("subnet_id", vnic.SubnetID)

	RefreshCreateVnicDetails(s.D, vnic)
}

func (s *InstanceResourceCrud) Delete() (e error) {
	return s.Client.TerminateInstance(s.D.Id(), nil)
}
