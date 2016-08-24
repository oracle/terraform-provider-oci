package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type InstanceSync struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.Instance
}

func (s *InstanceSync) ID() string {
	return s.Resource.ID
}

func (s *InstanceSync) CreatedPending() []string {
	return []string{
		baremetal.ResourceProvisioning,
		baremetal.ResourceStarting,
	}
}

func (s *InstanceSync) CreatedTarget() []string {
	return []string{baremetal.ResourceRunning}
}

func (s *InstanceSync) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *InstanceSync) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func resourceMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func (s *InstanceSync) Create() (e error) {
	opts := baremetal.Options{}
	availabilityDomain := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	image := s.D.Get("image").(string)
	shape := s.D.Get("shape").(string)
	subnet := s.D.Get("subnet_id").(string)
	metadata := resourceMapToMetadata(s.D.Get("metadata").(map[string]interface{}))

	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.LaunchInstance(
		availabilityDomain,
		compartmentID,
		image,
		shape,
		subnet,
		metadata,
		opts)
	return
}

func (s *InstanceSync) Get() (e error) {
	s.Resource, e = s.Client.GetInstance(s.D.Id())
	return
}

func (s *InstanceSync) Update() (e error) {
	opts := baremetal.Options{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateInstance(s.D.Id(), opts)
	return

}

func (s *InstanceSync) SetData() {
	s.D.Set("availability_domain", s.Resource.AvailabilityDomain)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("image", s.Resource.Image)
	s.D.Set("metadata", s.Resource.Metadata)
	s.D.Set("region", s.Resource.Region)
	s.D.Set("shape", s.Resource.Shape)
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *InstanceSync) Delete() (e error) {
	return s.Client.TerminateInstance(s.D.Id())
}
