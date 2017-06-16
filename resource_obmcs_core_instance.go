// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"encoding/json"
	"log"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/oracle/terraform-provider-baremetal/crud"
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
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_public_ip": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"public_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringLenBetween(1, 255),
							// Required:     true, // TODO: Require at next major release
						},
					},
				},
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// TODO: Deprecated. Remove at next major release.
			"public_ip": {
				Type:       schema.TypeString,
				Required:   false,
				Computed:   true,
				Deprecated: "use create_vnic_details.public_ip",
			},
			"private_ip": {
				Type:       schema.TypeString,
				Required:   false,
				Computed:   true,
				Deprecated: "use create_vnic_details.private_ip",
			},
			"hostname_label": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "use create_vnic_details.hostname_label",
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// TODO: deprecate once https://github.com/MustWin/baremetal-sdk-go/pull/159 is resolved
				// Deprecated: "use create_vnic_details.subnet_id",
			},
		},
	}
}

func createInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.CreateResource(d, sync)
}

func readInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.ReadResource(sync)
}

func updateInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.UpdateResource(d, sync)
}

func deleteInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.DeleteResource(d, sync)
}

type InstanceResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Instance

	// Computed fields
	public_ip  string
	private_ip string
	vnic       *baremetal.Vnic
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
	subnet := s.D.Get("subnet_id").(string) // TODO: Deprecated, remove at next major release

	opts := &baremetal.LaunchInstanceOptions{
		HostnameLabel: s.D.Get("hostname_label").(string), // TODO: Deprecated, remove at next major release
		IpxeScript:    s.D.Get("ipxe_script").(string),
		Metadata:      resourceInstanceMapToMetadata(s.D.Get("metadata").(map[string]interface{})),
	}
	opts.DisplayName = s.D.Get("display_name").(string)

	if rawExtendedMetadata, ok := s.D.GetOk("extended_metadata"); ok {
		extendedMetadata := mapToExtendedMetadata(rawExtendedMetadata.(map[string]interface{}))
		opts.ExtendedMetadata = extendedMetadata
	}

	vs := s.D.Get("create_vnic_details").([]interface{})
	if len(vs) > 0 && vs[0] != nil {
		vnicOpts := baremetal.CreateVnicOptions{}
		vnic := vs[0].(map[string]interface{})
		log.Printf("[DEBUG] VNIC state: %#v", vnic)

		vnicOpts.SubnetID = vnic["subnet_id"].(string)
		// Workaround to allow either subnet_id attribute to function.
		// TODO: remove at next major release
		if subnet != "" && vnicOpts.SubnetID == "" {
			vnicOpts.SubnetID = subnet
		}
		vnicOpts.HostnameLabel = vnic["hostname_label"].(string)
		vnicOpts.DisplayName = vnic["display_name"].(string)
		vnicOpts.PrivateIp = vnic["private_ip"].(string)

		ip := vnic["assign_public_ip"].(bool)
		vnicOpts.AssignPublicIp = &ip

		opts.CreateVnicOptions = &vnicOpts
	}

	s.Resource, e = s.Client.LaunchInstance(
		availabilityDomain,
		compartmentID,
		image,
		shape,
		subnet,
		opts)
	return
}

/*
 * Return the the instance's first Vnic.
 *
 * NOTE while the instance is still being created, calls to this function
 * can return  an error priort to the Vnic being attached.
 */
func (s *InstanceResourceCrud) getVnic() (*baremetal.Vnic, error) {
	opts := &baremetal.ListVnicAttachmentsOptions{}
	opts.AvailabilityDomain = s.Resource.AvailabilityDomain
	opts.InstanceID = s.Resource.ID

	list, err := s.Client.ListVnicAttachments(s.Resource.CompartmentID, opts)
	if err != nil {
		return nil, err
	}

	if len(list.Attachments) < 1 {
		log.Printf("[DEBUG] getVnic - InstanceID: %q, State: %q, no vnic attachments: %q", s.Resource.ID, s.Resource.State, err)
		return nil, err
	}

	vnic, err := s.Client.GetVnic(list.Attachments[0].VnicID)
	if err != nil {
		return nil, err
	}

	return vnic, nil
}

func (s *InstanceResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetInstance(s.D.Id())
	if e != nil {
		return e
	}

	// Compute instance IPs through attached Vnic
	// (Not available while state==PROVISIONING)
	v, err := s.getVnic()
	if err != nil {
		log.Printf("[DEBUG] no vnic yet, skipping")
	} else {
		s.vnic = v
	}

	return
}

func (s *InstanceResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateInstance(s.D.Id(), opts)
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

	v := map[string]interface{}{}
	v["assign_public_ip"] = s.vnic.PublicIPAddress != ""
	v["display_name"] = s.vnic.DisplayName
	v["hostname_label"] = s.vnic.HostnameLabel
	v["private_ip"] = s.vnic.PrivateIPAddress
	v["public_ip"] = s.vnic.PublicIPAddress
	v["subnet_id"] = s.vnic.SubnetID
	s.D.Set("create_vnic_details", []interface{}{v})

	// TODO: Deprecated. Remove at next major release.
	s.D.Set("public_ip", s.vnic.PublicIPAddress)
	s.D.Set("private_ip", s.vnic.PrivateIPAddress)
	s.D.Set("hostname_label", s.vnic.HostnameLabel)
	s.D.Set("subnet_id", s.vnic.SubnetID)
}

func (s *InstanceResourceCrud) Delete() (e error) {
	return s.Client.TerminateInstance(s.D.Id(), nil)
}
