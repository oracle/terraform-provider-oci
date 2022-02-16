// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"
)

func ObjectStoragePreauthenticatedRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createObjectStoragePreauthenticatedRequest,
		Read:     readObjectStoragePreauthenticatedRequest,
		Delete:   deleteObjectStoragePreauthenticatedRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"access_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeObjectread),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeObjectwrite),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeObjectreadwrite),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeAnyobjectwrite),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeAnyobjectread),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeAnyobjectreadwrite),
				}, true),
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_expires": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},

			// Optional
			"bucket_listing_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"object_name"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("object", "object_name"),
			},
			"object_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"object"},
			},

			// Computed
			"access_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"par_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectStoragePreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readObjectStoragePreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePreauthenticatedRequestResourceCrud{}
	// For backward compatibility with CompositeId change

	log.Printf("[DEBUG] readObjectStoragePreauthenticatedRequest() Resource Id in state: %s", d.Id())
	_, _, _, err := ParsePreauthenticatedRequestCompositeId(d.Id())

	if err != nil {
		bucket, bOk := d.GetOkExists("bucket")
		namespace, nOk := d.GetOkExists("namespace")

		if bOk && nOk {
			compositeId := GetPreauthenticatedRequestCompositeId(bucket.(string), namespace.(string), d.Id())
			d.SetId(compositeId)
		}
	}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

func deleteObjectStoragePreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ObjectStoragePreauthenticatedRequestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.PreauthenticatedRequest
	DisableNotFoundRetries bool
}

func (s *ObjectStoragePreauthenticatedRequestResourceCrud) ID() string {
	return GetPreauthenticatedRequestCompositeId(s.D.Get("bucket").(string), s.D.Get("namespace").(string), *s.Res.Id)
}

func (s *ObjectStoragePreauthenticatedRequestResourceCrud) Create() error {
	request := oci_object_storage.CreatePreauthenticatedRequestRequest{}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		request.AccessType = oci_object_storage.CreatePreauthenticatedRequestDetailsAccessTypeEnum(accessType.(string))
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if bucketListingAction, ok := s.D.GetOkExists("bucket_listing_action"); ok {
		request.BucketListingAction = oci_object_storage.PreauthenticatedRequestBucketListingActionEnum(bucketListingAction.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if object, ok := s.D.GetOkExists("object_name"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if timeExpires, ok := s.D.GetOkExists("time_expires"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpires.(string))
		if err != nil {
			return err
		}
		request.TimeExpires = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.CreatePreauthenticatedRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PreauthenticatedRequest
	return nil
}

func (s *ObjectStoragePreauthenticatedRequestResourceCrud) Get() error {
	request := oci_object_storage.GetPreauthenticatedRequestRequest{}

	bucket, namespace, parId, err := ParsePreauthenticatedRequestCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
		request.ParId = &parId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s, err: %s ", s.D.Id(), err)
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetPreauthenticatedRequest(context.Background(), request)
	if err != nil {
		return err
	}

	// Some contortions follow since GETs actually return a PreauthenticatedRequestSummary, but s.Res is typed as a
	// PreauthenticatedRequest

	s.Res = &oci_object_storage.PreauthenticatedRequest{
		Id:          response.PreauthenticatedRequestSummary.Id,
		AccessType:  oci_object_storage.PreauthenticatedRequestAccessTypeEnum(string(response.PreauthenticatedRequestSummary.AccessType)),
		Name:        response.PreauthenticatedRequestSummary.Name,
		ObjectName:  response.PreauthenticatedRequestSummary.ObjectName,
		TimeCreated: response.PreauthenticatedRequestSummary.TimeCreated,
		TimeExpires: response.PreauthenticatedRequestSummary.TimeExpires,
	}

	return nil
}

func (s *ObjectStoragePreauthenticatedRequestResourceCrud) Delete() error {
	request := oci_object_storage.DeletePreauthenticatedRequestRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if parId, ok := s.D.GetOkExists("par_id"); ok {
		tmp := parId.(string)
		request.ParId = &tmp
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeletePreauthenticatedRequest(context.Background(), request)
	return err
}

func (s *ObjectStoragePreauthenticatedRequestResourceCrud) SetData() error {

	// For ImportStateVerify to keep state consistent after import
	bucket, namespace, parId, err := ParsePreauthenticatedRequestCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bucket", &bucket)
		s.D.Set("namespace", &namespace)
		s.D.Set("par_id", &parId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s, err: %s ", s.D.Id(), err)
	}

	s.D.Set("access_type", s.Res.AccessType)

	if s.Res.AccessUri != nil {
		s.D.Set("access_uri", *s.Res.AccessUri)
	}

	s.D.Set("bucket_listing_action", s.Res.BucketListingAction)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object_name", *s.Res.ObjectName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.Format(time.RFC3339Nano))
	}

	return nil
}

func GetPreauthenticatedRequestCompositeId(bucket string, namespace string, parId string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	parId = url.PathEscape(parId)
	compositeId := "n/" + namespace + "/b/" + bucket + "/p/" + parId
	return compositeId
}

func ParsePreauthenticatedRequestCompositeId(compositeId string) (bucket string, namespace string, parId string, err error) {
	re := regexp.MustCompile(`n/([^/]+)/b/([^/]+)/p/(.+$)`)

	subMatchAll := re.FindStringSubmatch(compositeId)
	if subMatchAll == nil || len(subMatchAll) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(subMatchAll[1])
	bucket, _ = url.PathUnescape(subMatchAll[2])
	parId, _ = url.PathUnescape(subMatchAll[3])

	return
}
