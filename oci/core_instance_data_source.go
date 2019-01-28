// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreInstance,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"extended_metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"image": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"launch_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"launch_options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"boot_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firmware": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_consistent_volume_naming_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_data_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"bootVolume",
								"image",
							}, true),
						},

						// Optional
						"boot_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_maintenance_reboot_due": {
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
			"hostname_label": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Add this computed boot_volume_id field even though it's not part of the API specs. This will make it easier to
			// discover the attached boot volume's ID; to preserve it for reattachment.
			"boot_volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.VirtualNetworkClient = m.(*OracleClients).virtualNetworkClient
	sync.BlockStorageClient = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

type CoreInstanceDataSourceCrud struct {
	CoreInstanceResourceCrud
}

func (s *CoreInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceDataSourceCrud) Get() error {
	request := oci_core.GetInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *CoreInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		// @CODEGEN 1/2018: support legacy name "image"
		s.D.Set("image", *s.Res.ImageId)
	}

	if s.Res.IpxeScript != nil {
		s.D.Set("ipxe_script", *s.Res.IpxeScript)
	}

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	} else {
		s.D.Set("launch_options", nil)
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

	bootVolume, bootVolumeErr := s.getBootVolume()
	if bootVolumeErr != nil {
		log.Printf("[WARN] Could not get the boot volume: %q", bootVolumeErr)
	}

	if s.Res.SourceDetails != nil {
		var sourceDetailsFromConfig map[string]interface{}
		if details, ok := s.D.GetOkExists("source_details"); ok {
			if tmpList := details.([]interface{}); len(tmpList) > 0 {
				sourceDetailsFromConfig = tmpList[0].(map[string]interface{})
			}
		}
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceSourceDetailsToMap(&s.Res.SourceDetails, bootVolume, sourceDetailsFromConfig); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		err := s.D.Set("source_details", sourceDetailsArray)
		if err != nil {
			return err
		}
	} else {
		s.D.Set("source_details", nil)
	}

	if bootVolume != nil && bootVolume.Id != nil {
		s.D.Set("boot_volume_id", *bootVolume.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMaintenanceRebootDue != nil {
		s.D.Set("time_maintenance_reboot_due", s.Res.TimeMaintenanceRebootDue.String())
	}

	if s.Res.LifecycleState == oci_core.InstanceLifecycleStateRunning {
		vnic, vnicError := s.getPrimaryVnic()
		if vnicError != nil || vnic == nil {
			log.Printf("[WARN] Primary VNIC could not be found during instance refresh: %q", vnicError)
		} else {
			s.D.Set("hostname_label", vnic.HostnameLabel)
			s.D.Set("public_ip", vnic.PublicIp)
			s.D.Set("private_ip", vnic.PrivateIp)
			s.D.Set("subnet_id", vnic.SubnetId)
		}
	}

	return nil
}
