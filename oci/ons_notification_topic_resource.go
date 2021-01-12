// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_ons "github.com/oracle/oci-go-sdk/v32/ons"
)

func init() {
	RegisterResource("oci_ons_notification_topic", OnsNotificationTopicResource())
}

func OnsNotificationTopicResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("15m"),
			Update: getTimeoutDuration("15m"),
			Delete: getTimeoutDuration("150m"),
		},
		Create: createOnsNotificationTopic,
		Read:   readOnsNotificationTopic,
		Update: updateOnsNotificationTopic,
		Delete: deleteOnsNotificationTopic,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"api_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
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
			"topic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOnsNotificationTopic(d *schema.ResourceData, m interface{}) error {
	sync := &OnsNotificationTopicResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationControlPlaneClient()

	return CreateResource(d, sync)
}

func readOnsNotificationTopic(d *schema.ResourceData, m interface{}) error {
	sync := &OnsNotificationTopicResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationControlPlaneClient()

	return ReadResource(sync)
}

func updateOnsNotificationTopic(d *schema.ResourceData, m interface{}) error {
	sync := &OnsNotificationTopicResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationControlPlaneClient()

	return UpdateResource(d, sync)
}

func deleteOnsNotificationTopic(d *schema.ResourceData, m interface{}) error {
	sync := &OnsNotificationTopicResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationControlPlaneClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type OnsNotificationTopicResourceCrud struct {
	BaseCrud
	Client                 *oci_ons.NotificationControlPlaneClient
	Res                    *oci_ons.NotificationTopic
	DisableNotFoundRetries bool
}

func (s *OnsNotificationTopicResourceCrud) ID() string {
	return *s.Res.TopicId
}

func (s *OnsNotificationTopicResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ons.NotificationTopicLifecycleStateCreating),
	}
}

func (s *OnsNotificationTopicResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ons.NotificationTopicLifecycleStateActive),
	}
}

func (s *OnsNotificationTopicResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ons.NotificationTopicLifecycleStateDeleting),
	}
}

func (s *OnsNotificationTopicResourceCrud) DeletedTarget() []string {
	if avoidWaitingForDeleteTarget {
		return []string{
			string(oci_ons.NotificationTopicLifecycleStateDeleting),
		}
	}
	return []string{}
}

func (s *OnsNotificationTopicResourceCrud) Create() error {
	request := oci_ons.CreateTopicRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.CreateTopic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotificationTopic
	return nil
}

func (s *OnsNotificationTopicResourceCrud) Get() error {
	request := oci_ons.GetTopicRequest{}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	} else {
		tmp := s.D.Id()
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.GetTopic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotificationTopic
	return nil
}

func (s *OnsNotificationTopicResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ons.UpdateTopicRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.UpdateTopic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotificationTopic
	return nil
}

func (s *OnsNotificationTopicResourceCrud) Delete() error {
	request := oci_ons.DeleteTopicRequest{}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.DeleteTopic(context.Background(), request)
	return err
}

func (s *OnsNotificationTopicResourceCrud) SetData() error {

	s.D.SetId(*s.Res.TopicId)

	if s.Res.ApiEndpoint != nil {
		s.D.Set("api_endpoint", *s.Res.ApiEndpoint)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TopicId != nil {
		s.D.Set("topic_id", *s.Res.TopicId)
	}

	return nil
}

func (s *OnsNotificationTopicResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ons.ChangeTopicCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.ChangeTopicCompartmentDetails.CompartmentId = &compartmentTmp

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		changeCompartmentRequest.TopicId = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.ChangeTopicCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
