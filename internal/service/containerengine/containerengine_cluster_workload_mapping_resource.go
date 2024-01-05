// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterWorkloadMappingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createContainerengineClusterWorkloadMapping,
		Read:     readContainerengineClusterWorkloadMapping,
		Update:   updateContainerengineClusterWorkloadMapping,
		Delete:   deleteContainerengineClusterWorkloadMapping,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mapped_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"mapped_tenancy_id": {
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
		},
	}
}

func createContainerengineClusterWorkloadMapping(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterWorkloadMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineClusterWorkloadMapping(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterWorkloadMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineClusterWorkloadMapping(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterWorkloadMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineClusterWorkloadMapping(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterWorkloadMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineClusterWorkloadMappingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.WorkloadMapping
	DisableNotFoundRetries bool
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) ID() string {
	return GetClusterWorkloadMappingCompositeId(s.D.Get("cluster_id").(string), *s.Res.Id)
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.WorkloadMappingLifecycleStateCreating),
	}
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.WorkloadMappingLifecycleStateActive),
	}
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.WorkloadMappingLifecycleStateDeleting),
	}
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.WorkloadMappingLifecycleStateDeleted),
	}
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) Create() error {
	request := oci_containerengine.CreateWorkloadMappingRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if mappedCompartmentId, ok := s.D.GetOkExists("mapped_compartment_id"); ok {
		tmp := mappedCompartmentId.(string)
		request.MappedCompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.CreateWorkloadMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WorkloadMapping
	return nil
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) Get() error {
	request := oci_containerengine.GetWorkloadMappingRequest{}

	clusterId, workloadMappingId, err := parseClusterWorkloadMappingCompositeId(s.D.Id())
	if err == nil {
		request.ClusterId = &clusterId
		request.WorkloadMappingId = &workloadMappingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetWorkloadMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WorkloadMapping
	return nil
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) Update() error {
	request := oci_containerengine.UpdateWorkloadMappingRequest{}

	clusterId, workloadMappingId, err := parseClusterWorkloadMappingCompositeId(s.D.Id())
	if err == nil {
		request.ClusterId = &clusterId
		request.WorkloadMappingId = &workloadMappingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if mappedCompartmentId, ok := s.D.GetOkExists("mapped_compartment_id"); ok {
		tmp := mappedCompartmentId.(string)
		request.MappedCompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateWorkloadMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WorkloadMapping
	return nil
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) Delete() error {
	request := oci_containerengine.DeleteWorkloadMappingRequest{}

	clusterId, workloadMappingId, err := parseClusterWorkloadMappingCompositeId(s.D.Id())
	if err == nil {
		request.ClusterId = &clusterId
		request.WorkloadMappingId = &workloadMappingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	_, err = s.Client.DeleteWorkloadMapping(context.Background(), request)
	return err
}

func (s *ContainerengineClusterWorkloadMappingResourceCrud) SetData() error {

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MappedCompartmentId != nil {
		s.D.Set("mapped_compartment_id", *s.Res.MappedCompartmentId)
	}

	if s.Res.MappedTenancyId != nil {
		s.D.Set("mapped_tenancy_id", *s.Res.MappedTenancyId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	clusterId, workloadMappingId, err := parseClusterWorkloadMappingCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("cluster_id", &clusterId)
		s.D.SetId(GetClusterWorkloadMappingCompositeId(clusterId, workloadMappingId))
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	return nil
}

func GetClusterWorkloadMappingCompositeId(clusterId string, workloadMappingId string) string {
	clusterId = url.PathEscape(clusterId)
	workloadMappingId = url.PathEscape(workloadMappingId)
	compositeId := "clusters/" + clusterId + "/workloadMappings/" + workloadMappingId
	return compositeId
}

func parseClusterWorkloadMappingCompositeId(compositeId string) (clusterId string, workloadMappingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("clusters/.*/workloadMappings/.*", compositeId)
	if !match || len(parts) != 4 || parts[1] == "" || parts[3] == "" {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	clusterId, _ = url.PathUnescape(parts[1])
	workloadMappingId, _ = url.PathUnescape(parts[3])

	return
}
