package spotinst

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// AwsGroupService is an interface for interfacing with the AwsGroup
// endpoints of the Spotinst API.
type AwsGroupService interface {
	List(*ListAwsGroupInput) (*ListAwsGroupOutput, error)
	Create(*CreateAwsGroupInput) (*CreateAwsGroupOutput, error)
	Read(*ReadAwsGroupInput) (*ReadAwsGroupOutput, error)
	Update(*UpdateAwsGroupInput) (*UpdateAwsGroupOutput, error)
	Delete(*DeleteAwsGroupInput) (*DeleteAwsGroupOutput, error)
}

// AwsGroupServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type AwsGroupServiceOp struct {
	client *Client
}

var _ AwsGroupService = &AwsGroupServiceOp{}

type AwsGroup struct {
	ID          *string              `json:"id,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Description *string              `json:"description,omitempty"`
	Capacity    *AwsGroupCapacity    `json:"capacity,omitempty"`
	Compute     *AwsGroupCompute     `json:"compute,omitempty"`
	Strategy    *AwsGroupStrategy    `json:"strategy,omitempty"`
	Scaling     *AwsGroupScaling     `json:"scaling,omitempty"`
	Scheduling  *AwsGroupScheduling  `json:"scheduling,omitempty"`
	Integration *AwsGroupIntegration `json:"thirdPartiesIntegration,omitempty"`
	Multai      *AwsGroupMultai      `json:"multai,omitempty"`

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

type AwsGroupIntegration struct {
	EC2ContainerService *AwsGroupEC2ContainerServiceIntegration `json:"ecs,omitempty"`
	ElasticBeanstalk    *AwsGroupElasticBeanstalkIntegration    `json:"elasticBeanstalk,omitempty"`
	Rancher             *AwsGroupRancherIntegration             `json:"rancher,omitempty"`
	Kubernetes          *AwsGroupKubernetesIntegration          `json:"kubernetes,omitempty"`
	Mesosphere          *AwsGroupMesosphereIntegration          `json:"mesosphere,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupMultai struct {
	Token     *string                   `json:"token,omitempty"`
	Balancers []*AwsGroupMultaiBalancer `json:"balancers,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupMultaiBalancer struct {
	ProjectID   *string `json:"projectId,omitempty"`
	BalancerID  *string `json:"balancerId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
	AzAwareness *bool   `json:"azAwareness,omitempty"`
	AutoWeight  *bool   `json:"autoWeight,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupRancherIntegration struct {
	MasterHost *string `json:"masterHost,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty"`
	SecretKey  *string `json:"secretKey,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupElasticBeanstalkIntegration struct {
	EnvironmentID *string `json:"environmentId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupEC2ContainerServiceIntegration struct {
	ClusterName *string `json:"clusterName,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupKubernetesIntegration struct {
	Server *string `json:"apiServer,omitempty"`
	Token  *string `json:"token,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupMesosphereIntegration struct {
	Server *string `json:"apiServer,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScheduling struct {
	Tasks []*AwsGroupScheduledTask `json:"tasks,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScheduledTask struct {
	Frequency           *string `json:"frequency,omitempty"`
	CronExpression      *string `json:"cronExpression,omitempty"`
	TaskType            *string `json:"taskType,omitempty"`
	ScaleTargetCapacity *int    `json:"scaleTargetCapacity,omitempty"`
	ScaleMinCapacity    *int    `json:"scaleMinCapacity,omitempty"`
	ScaleMaxCapacity    *int    `json:"scaleMaxCapacity,omitempty"`
	BatchSizePercentage *int    `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int    `json:"gracePeriod,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScaling struct {
	Up   []*AwsGroupScalingPolicy `json:"up,omitempty"`
	Down []*AwsGroupScalingPolicy `json:"down,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScalingPolicy struct {
	PolicyName        *string                           `json:"policyName,omitempty"`
	MetricName        *string                           `json:"metricName,omitempty"`
	Statistic         *string                           `json:"statistic,omitempty"`
	Unit              *string                           `json:"unit,omitempty"`
	Threshold         *float64                          `json:"threshold,omitempty"`
	Adjustment        *int                              `json:"adjustment,omitempty"`
	MinTargetCapacity *int                              `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *int                              `json:"maxTargetCapacity,omitempty"`
	Namespace         *string                           `json:"namespace,omitempty"`
	EvaluationPeriods *int                              `json:"evaluationPeriods,omitempty"`
	Period            *int                              `json:"period,omitempty"`
	Cooldown          *int                              `json:"cooldown,omitempty"`
	Operator          *string                           `json:"operator,omitempty"`
	Dimensions        []*AwsGroupScalingPolicyDimension `json:"dimensions,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScalingPolicyDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupStrategy struct {
	Risk                     *float64                  `json:"risk,omitempty"`
	OnDemandCount            *int                      `json:"onDemandCount,omitempty"`
	DrainingTimeout          *int                      `json:"drainingTimeout,omitempty"`
	AvailabilityVsCost       *string                   `json:"availabilityVsCost,omitempty"`
	UtilizeReservedInstances *bool                     `json:"utilizeReservedInstances,omitempty"`
	FallbackToOnDemand       *bool                     `json:"fallbackToOd,omitempty"`
	SpinUpTime               *int                      `json:"spinUpTime,omitempty"`
	Signals                  []*AwsGroupStrategySignal `json:"signals,omitempty"`
	Persistence              *AwsGroupPersistence      `json:"persistence,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupPersistence struct {
	ShouldPersistPrivateIp    *bool `json:"shouldPersistPrivateIp,omitempty"`
	ShouldPersistBlockDevices *bool `json:"shouldPersistBlockDevices,omitempty"`
	ShouldPersistRootDevice   *bool `json:"shouldPersistRootDevice,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupStrategySignal struct {
	Name    *string `json:"name,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupCapacity struct {
	Minimum *int    `json:"minimum,omitempty"`
	Maximum *int    `json:"maximum,omitempty"`
	Target  *int    `json:"target,omitempty"`
	Unit    *string `json:"unit,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupCompute struct {
	Product             *string                             `json:"product,omitempty"`
	InstanceTypes       *AwsGroupComputeInstanceType        `json:"instanceTypes,omitempty"`
	LaunchSpecification *AwsGroupComputeLaunchSpecification `json:"launchSpecification,omitempty"`
	AvailabilityZones   []*AwsGroupComputeAvailabilityZone  `json:"availabilityZones,omitempty"`
	ElasticIPs          []string                            `json:"elasticIps,omitempty"`
	EBSVolumePool       []*AwsGroupComputeEBSVolume         `json:"ebsVolumePool,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeEBSVolume struct {
	DeviceName *string  `json:"deviceName,omitempty"`
	VolumeIDs  []string `json:"volumeIds,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeInstanceType struct {
	OnDemand *string                              `json:"ondemand,omitempty"`
	Spot     []string                             `json:"spot,omitempty"`
	Weights  []*AwsGroupComputeInstanceTypeWeight `json:"weights,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeInstanceTypeWeight struct {
	InstanceType *string `json:"instanceType,omitempty"`
	Weight       *int    `json:"weightedCapacity,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeAvailabilityZone struct {
	Name     *string `json:"name,omitempty"`
	SubnetID *string `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeLaunchSpecification struct {
	LoadBalancerNames      []string                            `json:"loadBalancerNames,omitempty"`
	LoadBalancersConfig    *AwsGroupComputeLoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	SecurityGroupIDs       []string                            `json:"securityGroupIds,omitempty"`
	HealthCheckType        *string                             `json:"healthCheckType,omitempty"`
	HealthCheckGracePeriod *int                                `json:"healthCheckGracePeriod,omitempty"`
	ImageID                *string                             `json:"imageId,omitempty"`
	KeyPair                *string                             `json:"keyPair,omitempty"`
	UserData               *string                             `json:"userData,omitempty"`
	ShutdownScript         *string                             `json:"shutdownScript,omitempty"`
	Tenancy                *string                             `json:"tenancy,omitempty"`
	Monitoring             *bool                               `json:"monitoring,omitempty"`
	EBSOptimized           *bool                               `json:"ebsOptimized,omitempty"`
	IamInstanceProfile     *AwsGroupComputeIamInstanceProfile  `json:"iamRole,omitempty"`
	BlockDevices           []*AwsGroupComputeBlockDevice       `json:"blockDeviceMappings,omitempty"`
	NetworkInterfaces      []*AwsGroupComputeNetworkInterface  `json:"networkInterfaces,omitempty"`
	Tags                   []*AwsGroupComputeTag               `json:"tags,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeLoadBalancersConfig struct {
	LoadBalancers []*AwsGroupComputeLoadBalancer `json:"loadBalancers,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeLoadBalancer struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`
	Type *string `json:"type,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeNetworkInterface struct {
	ID                             *string  `json:"networkInterfaceId,omitempty"`
	Description                    *string  `json:"description,omitempty"`
	DeviceIndex                    *int     `json:"deviceIndex,omitempty"`
	SecondaryPrivateIPAddressCount *int     `json:"secondaryPrivateIpAddressCount,omitempty"`
	AssociatePublicIPAddress       *bool    `json:"associatePublicIpAddress,omitempty"`
	DeleteOnTermination            *bool    `json:"deleteOnTermination,omitempty"`
	SecurityGroupsIDs              []string `json:"groups,omitempty"`
	PrivateIPAddress               *string  `json:"privateIpAddress,omitempty"`
	SubnetID                       *string  `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeBlockDevice struct {
	DeviceName  *string             `json:"deviceName,omitempty"`
	VirtualName *string             `json:"virtualName,omitempty"`
	EBS         *AwsGroupComputeEBS `json:"ebs,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeEBS struct {
	DeleteOnTermination *bool   `json:"deleteOnTermination,omitempty"`
	Encrypted           *bool   `json:"encrypted,omitempty"`
	SnapshotID          *string `json:"snapshotId,omitempty"`
	VolumeType          *string `json:"volumeType,omitempty"`
	VolumeSize          *int    `json:"volumeSize,omitempty"`
	IOPS                *int    `json:"iops,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeIamInstanceProfile struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeTag struct {
	Key   *string `json:"tagKey,omitempty"`
	Value *string `json:"tagValue,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListAwsGroupInput struct{}

type ListAwsGroupOutput struct {
	Groups []*AwsGroup `json:"groups,omitempty"`
}

type CreateAwsGroupInput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type CreateAwsGroupOutput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type ReadAwsGroupInput struct {
	ID *string `json:"groupId,omitempty"`
}

type ReadAwsGroupOutput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type UpdateAwsGroupInput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type UpdateAwsGroupOutput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type DeleteAwsGroupInput struct {
	ID *string `json:"groupId,omitempty"`
}

type DeleteAwsGroupOutput struct{}

func awsGroupFromJSON(in []byte) (*AwsGroup, error) {
	b := new(AwsGroup)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func awsGroupsFromJSON(in []byte) ([]*AwsGroup, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AwsGroup, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsGroupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func awsGroupsFromHttpResponse(resp *http.Response) ([]*AwsGroup, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return awsGroupsFromJSON(body)
}

func (s *AwsGroupServiceOp) List(input *ListAwsGroupInput) (*ListAwsGroupOutput, error) {
	r := s.client.newRequest("GET", "/aws/ec2/group")

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListAwsGroupOutput{Groups: gs}, nil
}

func (s *AwsGroupServiceOp) Create(input *CreateAwsGroupInput) (*CreateAwsGroupOutput, error) {
	r := s.client.newRequest("POST", "/aws/ec2/group")
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAwsGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AwsGroupServiceOp) Read(input *ReadAwsGroupInput) (*ReadAwsGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.ID),
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

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadAwsGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AwsGroupServiceOp) Update(input *UpdateAwsGroupInput) (*UpdateAwsGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.Group.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Group.ID = nil

	r := s.client.newRequest("PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateAwsGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AwsGroupServiceOp) Delete(input *DeleteAwsGroupInput) (*DeleteAwsGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.ID),
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

	return &DeleteAwsGroupOutput{}, nil
}

//region AwsGroup
func (o *AwsGroup) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroup
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroup) SetID(v *string) *AwsGroup {
	if o.ID = v; v == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *AwsGroup) SetName(v *string) *AwsGroup {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AwsGroup) SetDescription(v *string) *AwsGroup {
	if o.Description = v; v == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *AwsGroup) SetCapacity(v *AwsGroupCapacity) *AwsGroup {
	if o.Capacity = v; v == nil {
		o.nullFields = append(o.nullFields, "Capacity")
	}
	return o
}

func (o *AwsGroup) SetCompute(v *AwsGroupCompute) *AwsGroup {
	if o.Compute = v; v == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

func (o *AwsGroup) SetStrategy(v *AwsGroupStrategy) *AwsGroup {
	if o.Strategy = v; v == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *AwsGroup) SetScaling(v *AwsGroupScaling) *AwsGroup {
	if o.Scaling = v; v == nil {
		o.nullFields = append(o.nullFields, "Scaling")
	}
	return o
}

func (o *AwsGroup) SetScheduling(v *AwsGroupScheduling) *AwsGroup {
	if o.Scheduling = v; v == nil {
		o.nullFields = append(o.nullFields, "Scheduling")
	}
	return o
}

func (o *AwsGroup) SetIntegration(v *AwsGroupIntegration) *AwsGroup {
	if o.Integration = v; v == nil {
		o.nullFields = append(o.nullFields, "Integration")
	}
	return o
}

func (o *AwsGroup) SetMultai(v *AwsGroupMultai) *AwsGroup {
	if o.Multai = v; v == nil {
		o.nullFields = append(o.nullFields, "Multai")
	}
	return o
}

//endregion

//region AwsGroupIntegration
func (o *AwsGroupIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupIntegration) SetEC2ContainerService(v *AwsGroupEC2ContainerServiceIntegration) *AwsGroupIntegration {
	if o.EC2ContainerService = v; v == nil {
		o.nullFields = append(o.nullFields, "EC2ContainerService")
	}
	return o
}

func (o *AwsGroupIntegration) SetElasticBeanstalk(v *AwsGroupElasticBeanstalkIntegration) *AwsGroupIntegration {
	if o.ElasticBeanstalk = v; v == nil {
		o.nullFields = append(o.nullFields, "ElasticBeanstalk")
	}
	return o
}

func (o *AwsGroupIntegration) SetRancher(v *AwsGroupRancherIntegration) *AwsGroupIntegration {
	if o.Rancher = v; v == nil {
		o.nullFields = append(o.nullFields, "Rancher")
	}
	return o
}

func (o *AwsGroupIntegration) SetKubernetes(v *AwsGroupKubernetesIntegration) *AwsGroupIntegration {
	if o.Kubernetes = v; v == nil {
		o.nullFields = append(o.nullFields, "Kubernetes")
	}
	return o
}

func (o *AwsGroupIntegration) SetMesosphere(v *AwsGroupMesosphereIntegration) *AwsGroupIntegration {
	if o.Mesosphere = v; v == nil {
		o.nullFields = append(o.nullFields, "Mesosphere")
	}
	return o
}

//endregion

//region AwsGroupMultai
func (o *AwsGroupMultai) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupMultai
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupMultai) SetToken(v *string) *AwsGroupMultai {
	if o.Token = v; v == nil {
		o.nullFields = append(o.nullFields, "Token")
	}
	return o
}

func (o *AwsGroupMultai) SetBalancers(v []*AwsGroupMultaiBalancer) *AwsGroupMultai {
	if o.Balancers = v; v == nil {
		o.nullFields = append(o.nullFields, "Balancers")
	}
	return o
}

//endregion

//region AwsGroupMultaiBalancer
func (o *AwsGroupMultaiBalancer) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupMultaiBalancer
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupMultaiBalancer) SetProjectID(v *string) *AwsGroupMultaiBalancer {
	if o.ProjectID = v; v == nil {
		o.nullFields = append(o.nullFields, "ProjectID")
	}
	return o
}

func (o *AwsGroupMultaiBalancer) SetBalancerID(v *string) *AwsGroupMultaiBalancer {
	if o.BalancerID = v; v == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *AwsGroupMultaiBalancer) SetTargetSetID(v *string) *AwsGroupMultaiBalancer {
	if o.TargetSetID = v; v == nil {
		o.nullFields = append(o.nullFields, "TargetSetID")
	}
	return o
}

func (o *AwsGroupMultaiBalancer) SetAzAwareness(v *bool) *AwsGroupMultaiBalancer {
	if o.AzAwareness = v; v == nil {
		o.nullFields = append(o.nullFields, "AzAwareness")
	}
	return o
}

func (o *AwsGroupMultaiBalancer) SetAutoWeight(v *bool) *AwsGroupMultaiBalancer {
	if o.AutoWeight = v; v == nil {
		o.nullFields = append(o.nullFields, "AutoWeight")
	}
	return o
}

//endregion

//region AwsGroupRancherIntegration
func (o *AwsGroupRancherIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupRancherIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupRancherIntegration) SetMasterHost(v *string) *AwsGroupRancherIntegration {
	if o.MasterHost = v; v == nil {
		o.nullFields = append(o.nullFields, "MasterHost")
	}
	return o
}

func (o *AwsGroupRancherIntegration) SetAccessKey(v *string) *AwsGroupRancherIntegration {
	if o.AccessKey = v; v == nil {
		o.nullFields = append(o.nullFields, "AccessKey")
	}
	return o
}

func (o *AwsGroupRancherIntegration) SetSecretKey(v *string) *AwsGroupRancherIntegration {
	if o.SecretKey = v; v == nil {
		o.nullFields = append(o.nullFields, "SecretKey")
	}
	return o
}

//endregion

//region AwsGroupElasticBeanstalkIntegration
func (o *AwsGroupElasticBeanstalkIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupElasticBeanstalkIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupElasticBeanstalkIntegration) SetEnvironmentID(v *string) *AwsGroupElasticBeanstalkIntegration {
	if o.EnvironmentID = v; v == nil {
		o.nullFields = append(o.nullFields, "EnvironmentID")
	}
	return o
}

//endregion

//region AwsGroupEC2ContainerServiceIntegration
func (o *AwsGroupEC2ContainerServiceIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupEC2ContainerServiceIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupEC2ContainerServiceIntegration) SetClusterName(v *string) *AwsGroupEC2ContainerServiceIntegration {
	if o.ClusterName = v; v == nil {
		o.nullFields = append(o.nullFields, "ClusterName")
	}
	return o
}

//endregion

//region AwsGroupKubernetesIntegration
func (o *AwsGroupKubernetesIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupKubernetesIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupKubernetesIntegration) SetServer(v *string) *AwsGroupKubernetesIntegration {
	if o.Server = v; v == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

func (o *AwsGroupKubernetesIntegration) SetToken(v *string) *AwsGroupKubernetesIntegration {
	if o.Token = v; v == nil {
		o.nullFields = append(o.nullFields, "Token")
	}
	return o
}

//endregion

//region AwsGroupMesosphereIntegration
func (o *AwsGroupMesosphereIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupMesosphereIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupMesosphereIntegration) SetServer(v *string) *AwsGroupMesosphereIntegration {
	if o.Server = v; v == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

//endregion

//region AwsGroupScheduling
func (o *AwsGroupScheduling) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupScheduling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupScheduling) SetTasks(v []*AwsGroupScheduledTask) *AwsGroupScheduling {
	if o.Tasks = v; v == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

//endregion

//region AwsGroupScheduledTask
func (o *AwsGroupScheduledTask) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupScheduledTask
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupScheduledTask) SetFrequency(v *string) *AwsGroupScheduledTask {
	if o.Frequency = v; v == nil {
		o.nullFields = append(o.nullFields, "Frequency")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetCronExpression(v *string) *AwsGroupScheduledTask {
	if o.CronExpression = v; v == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetTaskType(v *string) *AwsGroupScheduledTask {
	if o.TaskType = v; v == nil {
		o.nullFields = append(o.nullFields, "TaskType")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetScaleTargetCapacity(v *int) *AwsGroupScheduledTask {
	if o.ScaleTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "ScaleTargetCapacity")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetScaleMinCapacity(v *int) *AwsGroupScheduledTask {
	if o.ScaleMinCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "ScaleMinCapacity")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetScaleMaxCapacity(v *int) *AwsGroupScheduledTask {
	if o.ScaleMaxCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "ScaleMaxCapacity")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetBatchSizePercentage(v *int) *AwsGroupScheduledTask {
	if o.BatchSizePercentage = v; v == nil {
		o.nullFields = append(o.nullFields, "BatchSizePercentage")
	}
	return o
}

func (o *AwsGroupScheduledTask) SetGracePeriod(v *int) *AwsGroupScheduledTask {
	if o.GracePeriod = v; v == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

//endregion

//region AwsGroupScaling
func (o *AwsGroupScaling) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupScaling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupScaling) SetUp(v []*AwsGroupScalingPolicy) *AwsGroupScaling {
	if o.Up = v; v == nil {
		o.nullFields = append(o.nullFields, "Up")
	}
	return o
}

func (o *AwsGroupScaling) SetDown(v []*AwsGroupScalingPolicy) *AwsGroupScaling {
	if o.Down = v; v == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

//endregion

//region AwsGroupScalingPolicy
func (o *AwsGroupScalingPolicy) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupScalingPolicy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupScalingPolicy) SetPolicyName(v *string) *AwsGroupScalingPolicy {
	if o.PolicyName = v; v == nil {
		o.nullFields = append(o.nullFields, "PolicyName")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetMetricName(v *string) *AwsGroupScalingPolicy {
	if o.MetricName = v; v == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetStatistic(v *string) *AwsGroupScalingPolicy {
	if o.Statistic = v; v == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetUnit(v *string) *AwsGroupScalingPolicy {
	if o.Unit = v; v == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetThreshold(v *float64) *AwsGroupScalingPolicy {
	if o.Threshold = v; v == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetAdjustment(v *int) *AwsGroupScalingPolicy {
	if o.Adjustment = v; v == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetMinTargetCapacity(v *int) *AwsGroupScalingPolicy {
	if o.MinTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetMaxTargetCapacity(v *int) *AwsGroupScalingPolicy {
	if o.MaxTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetNamespace(v *string) *AwsGroupScalingPolicy {
	if o.Namespace = v; v == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetEvaluationPeriods(v *int) *AwsGroupScalingPolicy {
	if o.EvaluationPeriods = v; v == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetPeriod(v *int) *AwsGroupScalingPolicy {
	if o.Period = v; v == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetCooldown(v *int) *AwsGroupScalingPolicy {
	if o.Cooldown = v; v == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetOperator(v *string) *AwsGroupScalingPolicy {
	if o.Operator = v; v == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

func (o *AwsGroupScalingPolicy) SetDimensions(v []*AwsGroupScalingPolicyDimension) *AwsGroupScalingPolicy {
	if o.Dimensions = v; v == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}

//endregion

//region AwsGroupScalingPolicyDimension
func (o *AwsGroupScalingPolicyDimension) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupScalingPolicyDimension
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupScalingPolicyDimension) SetName(v *string) *AwsGroupScalingPolicyDimension {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AwsGroupScalingPolicyDimension) SetValue(v *string) *AwsGroupScalingPolicyDimension {
	if o.Value = v; v == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion

//region AwsGroupStrategy
func (o *AwsGroupStrategy) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupStrategy) SetRisk(v *float64) *AwsGroupStrategy {
	if o.Risk = v; v == nil {
		o.nullFields = append(o.nullFields, "Risk")
	}
	return o
}

func (o *AwsGroupStrategy) SetOnDemandCount(v *int) *AwsGroupStrategy {
	if o.OnDemandCount = v; v == nil {
		o.nullFields = append(o.nullFields, "OnDemandCount")
	}
	return o
}

func (o *AwsGroupStrategy) SetDrainingTimeout(v *int) *AwsGroupStrategy {
	if o.DrainingTimeout = v; v == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
	}
	return o
}

func (o *AwsGroupStrategy) SetAvailabilityVsCost(v *string) *AwsGroupStrategy {
	if o.AvailabilityVsCost = v; v == nil {
		o.nullFields = append(o.nullFields, "AvailabilityVsCost")
	}
	return o
}

func (o *AwsGroupStrategy) SetUtilizeReservedInstances(v *bool) *AwsGroupStrategy {
	if o.UtilizeReservedInstances = v; v == nil {
		o.nullFields = append(o.nullFields, "UtilizeReservedInstances")
	}
	return o
}

func (o *AwsGroupStrategy) SetFallbackToOnDemand(v *bool) *AwsGroupStrategy {
	if o.FallbackToOnDemand = v; v == nil {
		o.nullFields = append(o.nullFields, "FallbackToOnDemand")
	}
	return o
}

func (o *AwsGroupStrategy) SetSpinUpTime(v *int) *AwsGroupStrategy {
	if o.SpinUpTime = v; v == nil {
		o.nullFields = append(o.nullFields, "SpinUpTime")
	}
	return o
}

func (o *AwsGroupStrategy) SetSignals(v []*AwsGroupStrategySignal) *AwsGroupStrategy {
	if o.Signals = v; v == nil {
		o.nullFields = append(o.nullFields, "Signals")
	}
	return o
}

func (o *AwsGroupStrategy) SetPersistence(v *AwsGroupPersistence) *AwsGroupStrategy {
	if o.Persistence = v; v == nil {
		o.nullFields = append(o.nullFields, "Persistence")
	}
	return o
}

//endregion

//region AwsGroupPersistence
func (o *AwsGroupPersistence) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupPersistence
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupPersistence) SetShouldPersistPrivateIp(v *bool) *AwsGroupPersistence {
	if o.ShouldPersistPrivateIp = v; v == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistPrivateIp")
	}
	return o
}

func (o *AwsGroupPersistence) SetShouldPersistBlockDevices(v *bool) *AwsGroupPersistence {
	if o.ShouldPersistBlockDevices = v; v == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistBlockDevices")
	}
	return o
}

func (o *AwsGroupPersistence) SetShouldPersistRootDevice(v *bool) *AwsGroupPersistence {
	if o.ShouldPersistRootDevice = v; v == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistRootDevice")
	}
	return o
}

//endregion

//region AwsGroupStrategySignal
func (o *AwsGroupStrategySignal) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupStrategySignal
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupStrategySignal) SetName(v *string) *AwsGroupStrategySignal {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AwsGroupStrategySignal) SetTimeout(v *int) *AwsGroupStrategySignal {
	if o.Timeout = v; v == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

//endregion

//region AwsGroupCapacity
func (o *AwsGroupCapacity) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupCapacity
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupCapacity) SetMinimum(v *int) *AwsGroupCapacity {
	if o.Minimum = v; v == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *AwsGroupCapacity) SetMaximum(v *int) *AwsGroupCapacity {
	if o.Maximum = v; v == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *AwsGroupCapacity) SetTarget(v *int) *AwsGroupCapacity {
	if o.Target = v; v == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

func (o *AwsGroupCapacity) SetUnit(v *string) *AwsGroupCapacity {
	if o.Unit = v; v == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

//endregion

//region AwsGroupCompute
func (o *AwsGroupCompute) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupCompute
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupCompute) SetProduct(v *string) *AwsGroupCompute {
	if o.Product = v; v == nil {
		o.nullFields = append(o.nullFields, "Product")
	}
	return o
}

func (o *AwsGroupCompute) SetInstanceTypes(v *AwsGroupComputeInstanceType) *AwsGroupCompute {
	if o.InstanceTypes = v; v == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}

func (o *AwsGroupCompute) SetLaunchSpecification(v *AwsGroupComputeLaunchSpecification) *AwsGroupCompute {
	if o.LaunchSpecification = v; v == nil {
		o.nullFields = append(o.nullFields, "LaunchSpecification")
	}
	return o
}

func (o *AwsGroupCompute) SetAvailabilityZones(v []*AwsGroupComputeAvailabilityZone) *AwsGroupCompute {
	if o.AvailabilityZones = v; v == nil {
		o.nullFields = append(o.nullFields, "AvailabilityZones")
	}
	return o
}

func (o *AwsGroupCompute) SetElasticIPs(v []string) *AwsGroupCompute {
	if o.ElasticIPs = v; v == nil {
		o.nullFields = append(o.nullFields, "ElasticIPs")
	}
	return o
}

func (o *AwsGroupCompute) SetEBSVolumePool(v []*AwsGroupComputeEBSVolume) *AwsGroupCompute {
	if o.EBSVolumePool = v; v == nil {
		o.nullFields = append(o.nullFields, "EBSVolumePool")
	}
	return o
}

//endregion

//region AwsGroupComputeEBSVolume
func (o *AwsGroupComputeEBSVolume) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeEBSVolume
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeEBSVolume) SetDeviceName(v *string) *AwsGroupComputeEBSVolume {
	if o.DeviceName = v; v == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *AwsGroupComputeEBSVolume) SetVolumeIDs(v []string) *AwsGroupComputeEBSVolume {
	if o.VolumeIDs = v; v == nil {
		o.nullFields = append(o.nullFields, "VolumeIDs")
	}
	return o
}

//endregion

//region AwsGroupComputeInstanceType
func (o *AwsGroupComputeInstanceType) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeInstanceType
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeInstanceType) SetOnDemand(v *string) *AwsGroupComputeInstanceType {
	if o.OnDemand = v; v == nil {
		o.nullFields = append(o.nullFields, "OnDemand")
	}
	return o
}

func (o *AwsGroupComputeInstanceType) SetSpot(v []string) *AwsGroupComputeInstanceType {
	if o.Spot = v; v == nil {
		o.nullFields = append(o.nullFields, "Spot")
	}
	return o
}

func (o *AwsGroupComputeInstanceType) SetWeights(v []*AwsGroupComputeInstanceTypeWeight) *AwsGroupComputeInstanceType {
	if o.Weights = v; v == nil {
		o.nullFields = append(o.nullFields, "Weights")
	}
	return o
}

//endregion

//region AwsGroupComputeInstanceTypeWeight
func (o *AwsGroupComputeInstanceTypeWeight) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeInstanceTypeWeight
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeInstanceTypeWeight) SetInstanceType(v *string) *AwsGroupComputeInstanceTypeWeight {
	if o.InstanceType = v; v == nil {
		o.nullFields = append(o.nullFields, "InstanceType")
	}
	return o
}

func (o *AwsGroupComputeInstanceTypeWeight) SetWeight(v *int) *AwsGroupComputeInstanceTypeWeight {
	if o.Weight = v; v == nil {
		o.nullFields = append(o.nullFields, "Weight")
	}
	return o
}

//endregion

//region AwsGroupComputeAvailabilityZone
func (o *AwsGroupComputeAvailabilityZone) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeAvailabilityZone
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeAvailabilityZone) SetName(v *string) *AwsGroupComputeAvailabilityZone {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AwsGroupComputeAvailabilityZone) SetSubnetID(v *string) *AwsGroupComputeAvailabilityZone {
	if o.SubnetID = v; v == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

//endregion

//region AwsGroupComputeLaunchSpecification
func (o *AwsGroupComputeLaunchSpecification) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeLaunchSpecification
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeLaunchSpecification) SetLoadBalancerNames(v []string) *AwsGroupComputeLaunchSpecification {
	if o.LoadBalancerNames = v; v == nil {
		o.nullFields = append(o.nullFields, "LoadBalancerNames")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetLoadBalancersConfig(v *AwsGroupComputeLoadBalancersConfig) *AwsGroupComputeLaunchSpecification {
	if o.LoadBalancersConfig = v; v == nil {
		o.nullFields = append(o.nullFields, "LoadBalancersConfig")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetSecurityGroupIDs(v []string) *AwsGroupComputeLaunchSpecification {
	if o.SecurityGroupIDs = v; v == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupIDs")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetHealthCheckType(v *string) *AwsGroupComputeLaunchSpecification {
	if o.HealthCheckType = v; v == nil {
		o.nullFields = append(o.nullFields, "HealthCheckType")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetHealthCheckGracePeriod(v *int) *AwsGroupComputeLaunchSpecification {
	if o.HealthCheckGracePeriod = v; v == nil {
		o.nullFields = append(o.nullFields, "HealthCheckGracePeriod")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetImageID(v *string) *AwsGroupComputeLaunchSpecification {
	if o.ImageID = v; v == nil {
		o.nullFields = append(o.nullFields, "ImageID")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetKeyPair(v *string) *AwsGroupComputeLaunchSpecification {
	if o.KeyPair = v; v == nil {
		o.nullFields = append(o.nullFields, "KeyPair")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetUserData(v *string) *AwsGroupComputeLaunchSpecification {
	if o.UserData = v; v == nil {
		o.nullFields = append(o.nullFields, "UserData")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetShutdownScript(v *string) *AwsGroupComputeLaunchSpecification {
	if o.ShutdownScript = v; v == nil {
		o.nullFields = append(o.nullFields, "ShutdownScript")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetTenancy(v *string) *AwsGroupComputeLaunchSpecification {
	if o.Tenancy = v; v == nil {
		o.nullFields = append(o.nullFields, "Tenancy")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetMonitoring(v *bool) *AwsGroupComputeLaunchSpecification {
	if o.Monitoring = v; v == nil {
		o.nullFields = append(o.nullFields, "Monitoring")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetEBSOptimized(v *bool) *AwsGroupComputeLaunchSpecification {
	if o.EBSOptimized = v; v == nil {
		o.nullFields = append(o.nullFields, "EBSOptimized")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetIamInstanceProfile(v *AwsGroupComputeIamInstanceProfile) *AwsGroupComputeLaunchSpecification {
	if o.IamInstanceProfile = v; v == nil {
		o.nullFields = append(o.nullFields, "IamInstanceProfile")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetBlockDevices(v []*AwsGroupComputeBlockDevice) *AwsGroupComputeLaunchSpecification {
	if o.BlockDevices = v; v == nil {
		o.nullFields = append(o.nullFields, "BlockDevices")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetNetworkInterfaces(v []*AwsGroupComputeNetworkInterface) *AwsGroupComputeLaunchSpecification {
	if o.NetworkInterfaces = v; v == nil {
		o.nullFields = append(o.nullFields, "NetworkInterfaces")
	}
	return o
}

func (o *AwsGroupComputeLaunchSpecification) SetTags(v []*AwsGroupComputeTag) *AwsGroupComputeLaunchSpecification {
	if o.Tags = v; v == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

//endregion

//region AwsGroupComputeLoadBalancersConfig
func (o *AwsGroupComputeLoadBalancersConfig) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeLoadBalancersConfig
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeLoadBalancersConfig) SetLoadBalancers(v []*AwsGroupComputeLoadBalancer) *AwsGroupComputeLoadBalancersConfig {
	if o.LoadBalancers = v; v == nil {
		o.nullFields = append(o.nullFields, "LoadBalancers")
	}
	return o
}

//endregion

//region AwsGroupComputeLoadBalancer
func (o *AwsGroupComputeLoadBalancer) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeLoadBalancer
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeLoadBalancer) SetName(v *string) *AwsGroupComputeLoadBalancer {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AwsGroupComputeLoadBalancer) SetArn(v *string) *AwsGroupComputeLoadBalancer {
	if o.Arn = v; v == nil {
		o.nullFields = append(o.nullFields, "Arn")
	}
	return o
}

func (o *AwsGroupComputeLoadBalancer) SetType(v *string) *AwsGroupComputeLoadBalancer {
	if o.Type = v; v == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

//endregion

//region AwsGroupComputeNetworkInterface
func (o *AwsGroupComputeNetworkInterface) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeNetworkInterface
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeNetworkInterface) SetID(v *string) *AwsGroupComputeNetworkInterface {
	if o.ID = v; v == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetDescription(v *string) *AwsGroupComputeNetworkInterface {
	if o.Description = v; v == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetDeviceIndex(v *int) *AwsGroupComputeNetworkInterface {
	if o.DeviceIndex = v; v == nil {
		o.nullFields = append(o.nullFields, "DeviceIndex")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetSecondaryPrivateIPAddressCount(v *int) *AwsGroupComputeNetworkInterface {
	if o.SecondaryPrivateIPAddressCount = v; v == nil {
		o.nullFields = append(o.nullFields, "SecondaryPrivateIPAddressCount")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetAssociatePublicIPAddress(v *bool) *AwsGroupComputeNetworkInterface {
	if o.AssociatePublicIPAddress = v; v == nil {
		o.nullFields = append(o.nullFields, "AssociatePublicIPAddress")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetDeleteOnTermination(v *bool) *AwsGroupComputeNetworkInterface {
	if o.DeleteOnTermination = v; v == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetSecurityGroupsIDs(v []string) *AwsGroupComputeNetworkInterface {
	if o.SecurityGroupsIDs = v; v == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupsIDs")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetPrivateIPAddress(v *string) *AwsGroupComputeNetworkInterface {
	if o.PrivateIPAddress = v; v == nil {
		o.nullFields = append(o.nullFields, "PrivateIPAddress")
	}
	return o
}

func (o *AwsGroupComputeNetworkInterface) SetSubnetID(v *string) *AwsGroupComputeNetworkInterface {
	if o.SubnetID = v; v == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

//endregion

//region AwsGroupComputeBlockDevice
func (o *AwsGroupComputeBlockDevice) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeBlockDevice
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeBlockDevice) SetDeviceName(v *string) *AwsGroupComputeBlockDevice {
	if o.DeviceName = v; v == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *AwsGroupComputeBlockDevice) SetVirtualName(v *string) *AwsGroupComputeBlockDevice {
	if o.VirtualName = v; v == nil {
		o.nullFields = append(o.nullFields, "VirtualName")
	}
	return o
}

func (o *AwsGroupComputeBlockDevice) SetEBS(v *AwsGroupComputeEBS) *AwsGroupComputeBlockDevice {
	if o.EBS = v; v == nil {
		o.nullFields = append(o.nullFields, "EBS")
	}
	return o
}

//endregion

//region AwsGroupComputeEBS
func (o *AwsGroupComputeEBS) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeEBS
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeEBS) SetDeleteOnTermination(v *bool) *AwsGroupComputeEBS {
	if o.DeleteOnTermination = v; v == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *AwsGroupComputeEBS) SetEncrypted(v *bool) *AwsGroupComputeEBS {
	if o.Encrypted = v; v == nil {
		o.nullFields = append(o.nullFields, "Encrypted")
	}
	return o
}

func (o *AwsGroupComputeEBS) SetSnapshotID(v *string) *AwsGroupComputeEBS {
	if o.SnapshotID = v; v == nil {
		o.nullFields = append(o.nullFields, "SnapshotID")
	}
	return o
}

func (o *AwsGroupComputeEBS) SetVolumeType(v *string) *AwsGroupComputeEBS {
	if o.VolumeType = v; v == nil {
		o.nullFields = append(o.nullFields, "VolumeType")
	}
	return o
}

func (o *AwsGroupComputeEBS) SetVolumeSize(v *int) *AwsGroupComputeEBS {
	if o.VolumeSize = v; v == nil {
		o.nullFields = append(o.nullFields, "VolumeSize")
	}
	return o
}

func (o *AwsGroupComputeEBS) SetIOPS(v *int) *AwsGroupComputeEBS {
	if o.IOPS = v; v == nil {
		o.nullFields = append(o.nullFields, "IOPS")
	}
	return o
}

//endregion

//region AwsGroupComputeIamInstanceProfile
func (o *AwsGroupComputeIamInstanceProfile) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeIamInstanceProfile
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeIamInstanceProfile) SetName(v *string) *AwsGroupComputeIamInstanceProfile {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AwsGroupComputeIamInstanceProfile) SetArn(v *string) *AwsGroupComputeIamInstanceProfile {
	if o.Arn = v; v == nil {
		o.nullFields = append(o.nullFields, "Arn")
	}
	return o
}

//endregion

//region AwsGroupComputeTag
func (o *AwsGroupComputeTag) MarshalJSON() ([]byte, error) {
	type noMethod AwsGroupComputeTag
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AwsGroupComputeTag) SetKey(v *string) *AwsGroupComputeTag {
	if o.Key = v; v == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *AwsGroupComputeTag) SetValue(v *string) *AwsGroupComputeTag {
	if o.Value = v; v == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion
