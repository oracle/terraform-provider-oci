package spotinst

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// Subscription is an interface for interfacing with the Subscription
// endpoints of the Spotinst API.
type SubscriptionService interface {
	List(*ListSubscriptionInput) (*ListSubscriptionOutput, error)
	Create(*CreateSubscriptionInput) (*CreateSubscriptionOutput, error)
	Read(*ReadSubscriptionInput) (*ReadSubscriptionOutput, error)
	Update(*UpdateSubscriptionInput) (*UpdateSubscriptionOutput, error)
	Delete(*DeleteSubscriptionInput) (*DeleteSubscriptionOutput, error)
}

// SubscriptionServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type SubscriptionServiceOp struct {
	client *Client
}

var _ SubscriptionService = &SubscriptionServiceOp{}

type Subscription struct {
	ID         *string                `json:"id,omitempty"`
	ResourceID *string                `json:"resourceId,omitempty"`
	EventType  *string                `json:"eventType,omitempty"`
	Protocol   *string                `json:"protocol,omitempty"`
	Endpoint   *string                `json:"endpoint,omitempty"`
	Format     map[string]interface{} `json:"eventFormat,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string `json:"-"`

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string `json:"-"`
}

type ListSubscriptionInput struct{}

type ListSubscriptionOutput struct {
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
}

type CreateSubscriptionInput struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type CreateSubscriptionOutput struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type ReadSubscriptionInput struct {
	ID *string `json:"subscriptionId,omitempty"`
}

type ReadSubscriptionOutput struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type UpdateSubscriptionInput struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type UpdateSubscriptionOutput struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type DeleteSubscriptionInput struct {
	ID *string `json:"subscriptionId,omitempty"`
}

type DeleteSubscriptionOutput struct{}

func subscriptionFromJSON(in []byte) (*Subscription, error) {
	b := new(Subscription)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func subscriptionsFromJSON(in []byte) ([]*Subscription, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Subscription, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := subscriptionFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func subscriptionsFromHttpResponse(resp *http.Response) ([]*Subscription, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return subscriptionsFromJSON(body)
}

func (s *SubscriptionServiceOp) List(input *ListSubscriptionInput) (*ListSubscriptionOutput, error) {
	r := s.client.newRequest("GET", "/events/subscription")

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := subscriptionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListSubscriptionOutput{Subscriptions: gs}, nil
}

func (s *SubscriptionServiceOp) Create(input *CreateSubscriptionInput) (*CreateSubscriptionOutput, error) {
	r := s.client.newRequest("POST", "/events/subscription")
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := subscriptionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateSubscriptionOutput)
	if len(ss) > 0 {
		output.Subscription = ss[0]
	}

	return output, nil
}

func (s *SubscriptionServiceOp) Read(input *ReadSubscriptionInput) (*ReadSubscriptionOutput, error) {
	path, err := uritemplates.Expand("/events/subscription/{subscriptionId}", map[string]string{
		"subscriptionId": StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest("GET", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := subscriptionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadSubscriptionOutput)
	if len(ss) > 0 {
		output.Subscription = ss[0]
	}

	return output, nil
}

func (s *SubscriptionServiceOp) Update(input *UpdateSubscriptionInput) (*UpdateSubscriptionOutput, error) {
	path, err := uritemplates.Expand("/events/subscription/{subscriptionId}", map[string]string{
		"subscriptionId": StringValue(input.Subscription.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Subscription.ID = nil

	r := s.client.newRequest("PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ss, err := subscriptionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateSubscriptionOutput)
	if len(ss) > 0 {
		output.Subscription = ss[0]
	}

	return output, nil
}

func (s *SubscriptionServiceOp) Delete(input *DeleteSubscriptionInput) (*DeleteSubscriptionOutput, error) {
	path, err := uritemplates.Expand("/events/subscription/{subscriptionId}", map[string]string{
		"subscriptionId": StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest("DELETE", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteSubscriptionOutput{}, nil
}

//region Subscription
func (o *Subscription) MarshalJSON() ([]byte, error) {
	type noMethod Subscription
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Subscription) SetID(v *string) *Subscription {
	if o.ID = v; v == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Subscription) SetResourceID(v *string) *Subscription {
	if o.ResourceID = v; v == nil {
		o.nullFields = append(o.nullFields, "ResourceID")
	}
	return o
}

func (o *Subscription) SetEventType(v *string) *Subscription {
	if o.EventType = v; v == nil {
		o.nullFields = append(o.nullFields, "EventType")
	}
	return o
}

func (o *Subscription) SetProtocol(v *string) *Subscription {
	if o.Protocol = v; v == nil {
		o.nullFields = append(o.nullFields, "Protocol")
	}
	return o
}

func (o *Subscription) SetEndpoint(v *string) *Subscription {
	if o.Endpoint = v; v == nil {
		o.nullFields = append(o.nullFields, "Endpoint")
	}
	return o
}

func (o *Subscription) SetFormat(v map[string]interface{}) *Subscription {
	if o.Format = v; v == nil {
		o.nullFields = append(o.nullFields, "Format")
	}
	return o
}

//endregion
