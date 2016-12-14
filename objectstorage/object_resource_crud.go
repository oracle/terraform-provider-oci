package objectstorage

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type ObjectResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Object
}

func (s *ObjectResourceCrud) ID() string {
	return "tfobm-object-" + s.Res.Namespace + "/" + s.Res.Bucket + "/" + s.Res.ID
}

func (s *ObjectResourceCrud) SetData() {
	s.D.Set("namespace", s.Res.Namespace)
	s.D.Set("bucket", s.Res.Bucket)
	s.D.Set("object", s.Res.ID)
	s.D.Set("content", s.Res.Body)
	s.D.Set("metadata", s.Res.Metadata)
}

func (s *ObjectResourceCrud) Create() (e error) {
	e = s.Update()
	return
}

func (s *ObjectResourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	s.Res, e = s.Client.GetObject(namespace, bucket, object)
	return
}

func (s *ObjectResourceCrud) Update() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	opts := &baremetal.PutObjectOptions{}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}
	s.Res, e = s.Client.PutObject(namespace, bucket, object, namespace, opts)
	return
}

func (s *ObjectResourceCrud) Delete() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	opts := &baremetal.DeleteObjectOptions{}

	return s.Client.DeleteObject(namespace, bucket, object, opts)
}
