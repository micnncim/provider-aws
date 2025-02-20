/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type Authentication struct {
	PasswordCount *int64 `json:"passwordCount,omitempty"`
}

// +kubebuilder:skipversion
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type CacheCluster struct {
	ARN *string `json:"arn,omitempty"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty"`

	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	CacheClusterStatus *string `json:"cacheClusterStatus,omitempty"`

	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	CacheSubnetGroupName *string `json:"cacheSubnetGroupName,omitempty"`

	ClientDownloadLandingPage *string `json:"clientDownloadLandingPage,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	NumCacheNodes *int64 `json:"numCacheNodes,omitempty"`

	PreferredAvailabilityZone *string `json:"preferredAvailabilityZone,omitempty"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`

	PreferredOutpostARN *string `json:"preferredOutpostARN,omitempty"`

	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty"`

	SnapshotWindow *string `json:"snapshotWindow,omitempty"`
}

// +kubebuilder:skipversion
type CacheEngineVersion struct {
	CacheEngineDescription *string `json:"cacheEngineDescription,omitempty"`

	CacheEngineVersionDescription *string `json:"cacheEngineVersionDescription,omitempty"`

	CacheParameterGroupFamily *string `json:"cacheParameterGroupFamily,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`
}

// +kubebuilder:skipversion
type CacheNode struct {
	CacheNodeID *string `json:"cacheNodeID,omitempty"`

	CacheNodeStatus *string `json:"cacheNodeStatus,omitempty"`

	CustomerAvailabilityZone *string `json:"customerAvailabilityZone,omitempty"`

	CustomerOutpostARN *string `json:"customerOutpostARN,omitempty"`

	ParameterGroupStatus *string `json:"parameterGroupStatus,omitempty"`

	SourceCacheNodeID *string `json:"sourceCacheNodeID,omitempty"`
}

// +kubebuilder:skipversion
type CacheNodeTypeSpecificParameter struct {
	AllowedValues *string `json:"allowedValues,omitempty"`

	DataType *string `json:"dataType,omitempty"`

	Description *string `json:"description,omitempty"`

	IsModifiable *bool `json:"isModifiable,omitempty"`

	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`

	ParameterName *string `json:"parameterName,omitempty"`

	Source *string `json:"source,omitempty"`
}

// +kubebuilder:skipversion
type CacheNodeTypeSpecificValue struct {
	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type CacheNodeUpdateStatus struct {
	CacheNodeID *string `json:"cacheNodeID,omitempty"`
}

// +kubebuilder:skipversion
type CacheParameterGroupStatus_SDK struct {
	CacheParameterGroupName *string `json:"cacheParameterGroupName,omitempty"`

	ParameterApplyStatus *string `json:"parameterApplyStatus,omitempty"`
}

// +kubebuilder:skipversion
type CacheParameterGroup_SDK struct {
	ARN *string `json:"arn,omitempty"`

	CacheParameterGroupFamily *string `json:"cacheParameterGroupFamily,omitempty"`

	CacheParameterGroupName *string `json:"cacheParameterGroupName,omitempty"`

	Description *string `json:"description,omitempty"`

	IsGlobal *bool `json:"isGlobal,omitempty"`
}

// +kubebuilder:skipversion
type CacheSecurityGroup struct {
	ARN *string `json:"arn,omitempty"`

	CacheSecurityGroupName *string `json:"cacheSecurityGroupName,omitempty"`

	Description *string `json:"description,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`
}

// +kubebuilder:skipversion
type CacheSecurityGroupMembership struct {
	CacheSecurityGroupName *string `json:"cacheSecurityGroupName,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type CacheSubnetGroup struct {
	ARN *string `json:"arn,omitempty"`

	CacheSubnetGroupDescription *string `json:"cacheSubnetGroupDescription,omitempty"`

	CacheSubnetGroupName *string `json:"cacheSubnetGroupName,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

// +kubebuilder:skipversion
type CustomerNodeEndpoint struct {
	Address *string `json:"address,omitempty"`

	Port *int64 `json:"port,omitempty"`
}

// +kubebuilder:skipversion
type EC2SecurityGroup struct {
	EC2SecurityGroupName *string `json:"ec2SecurityGroupName,omitempty"`

	EC2SecurityGroupOwnerID *string `json:"ec2SecurityGroupOwnerID,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type Endpoint struct {
	Address *string `json:"address,omitempty"`
}

// +kubebuilder:skipversion
type EngineDefaults struct {
	CacheParameterGroupFamily *string `json:"cacheParameterGroupFamily,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

// +kubebuilder:skipversion
type Event struct {
	Message *string `json:"message,omitempty"`

	SourceIdentifier *string `json:"sourceIdentifier,omitempty"`
}

// +kubebuilder:skipversion
type GlobalNodeGroup struct {
	GlobalNodeGroupID *string `json:"globalNodeGroupID,omitempty"`

	Slots *string `json:"slots,omitempty"`
}

// +kubebuilder:skipversion
type GlobalReplicationGroup struct {
	ARN *string `json:"arn,omitempty"`

	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty"`

	GlobalReplicationGroupID *string `json:"globalReplicationGroupID,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type GlobalReplicationGroupInfo struct {
	GlobalReplicationGroupID *string `json:"globalReplicationGroupID,omitempty"`

	GlobalReplicationGroupMemberRole *string `json:"globalReplicationGroupMemberRole,omitempty"`
}

// +kubebuilder:skipversion
type GlobalReplicationGroupMember struct {
	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	ReplicationGroupRegion *string `json:"replicationGroupRegion,omitempty"`

	Role *string `json:"role,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type NodeGroup struct {
	NodeGroupID *string `json:"nodeGroupID,omitempty"`

	Slots *string `json:"slots,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type NodeGroupConfiguration struct {
	PrimaryAvailabilityZone *string `json:"primaryAvailabilityZone,omitempty"`

	PrimaryOutpostARN *string `json:"primaryOutpostARN,omitempty"`

	ReplicaCount *int64 `json:"replicaCount,omitempty"`

	Slots *string `json:"slots,omitempty"`
}

// +kubebuilder:skipversion
type NodeGroupMember struct {
	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	CacheNodeID *string `json:"cacheNodeID,omitempty"`

	CurrentRole *string `json:"currentRole,omitempty"`

	PreferredAvailabilityZone *string `json:"preferredAvailabilityZone,omitempty"`

	PreferredOutpostARN *string `json:"preferredOutpostARN,omitempty"`
}

// +kubebuilder:skipversion
type NodeGroupMemberUpdateStatus struct {
	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	CacheNodeID *string `json:"cacheNodeID,omitempty"`
}

// +kubebuilder:skipversion
type NodeGroupUpdateStatus struct {
	NodeGroupID *string `json:"nodeGroupID,omitempty"`
}

// +kubebuilder:skipversion
type NodeSnapshot struct {
	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	CacheNodeID *string `json:"cacheNodeID,omitempty"`

	CacheSize *string `json:"cacheSize,omitempty"`

	NodeGroupID *string `json:"nodeGroupID,omitempty"`
}

// +kubebuilder:skipversion
type NotificationConfiguration struct {
	TopicARN *string `json:"topicARN,omitempty"`

	TopicStatus *string `json:"topicStatus,omitempty"`
}

// +kubebuilder:skipversion
type Parameter struct {
	AllowedValues *string `json:"allowedValues,omitempty"`

	DataType *string `json:"dataType,omitempty"`

	Description *string `json:"description,omitempty"`

	IsModifiable *bool `json:"isModifiable,omitempty"`

	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`

	ParameterName *string `json:"parameterName,omitempty"`

	ParameterValue *string `json:"parameterValue,omitempty"`

	Source *string `json:"source,omitempty"`
}

// +kubebuilder:skipversion
type ParameterNameValue struct {
	ParameterName *string `json:"parameterName,omitempty"`

	ParameterValue *string `json:"parameterValue,omitempty"`
}

// +kubebuilder:skipversion
type PendingModifiedValues struct {
	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	NumCacheNodes *int64 `json:"numCacheNodes,omitempty"`
}

// +kubebuilder:skipversion
type ProcessedUpdateAction struct {
	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	ServiceUpdateName *string `json:"serviceUpdateName,omitempty"`
}

// +kubebuilder:skipversion
type RecurringCharge struct {
	RecurringChargeFrequency *string `json:"recurringChargeFrequency,omitempty"`
}

// +kubebuilder:skipversion
type RegionalConfiguration struct {
	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	ReplicationGroupRegion *string `json:"replicationGroupRegion,omitempty"`
}

// +kubebuilder:skipversion
type ReplicationGroup struct {
	ARN *string `json:"arn,omitempty"`

	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	Description *string `json:"description,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty"`

	SnapshotWindow *string `json:"snapshotWindow,omitempty"`

	SnapshottingClusterID *string `json:"snapshottingClusterID,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type ReplicationGroupPendingModifiedValues struct {
	PrimaryClusterID *string `json:"primaryClusterID,omitempty"`
}

// +kubebuilder:skipversion
type ReservedCacheNode struct {
	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	OfferingType *string `json:"offeringType,omitempty"`

	ProductDescription *string `json:"productDescription,omitempty"`

	ReservationARN *string `json:"reservationARN,omitempty"`

	ReservedCacheNodeID *string `json:"reservedCacheNodeID,omitempty"`

	ReservedCacheNodesOfferingID *string `json:"reservedCacheNodesOfferingID,omitempty"`

	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type ReservedCacheNodesOffering struct {
	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	OfferingType *string `json:"offeringType,omitempty"`

	ProductDescription *string `json:"productDescription,omitempty"`

	ReservedCacheNodesOfferingID *string `json:"reservedCacheNodesOfferingID,omitempty"`
}

// +kubebuilder:skipversion
type SecurityGroupMembership struct {
	SecurityGroupID *string `json:"securityGroupID,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type ServiceUpdate struct {
	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	EstimatedUpdateTime *string `json:"estimatedUpdateTime,omitempty"`

	ServiceUpdateDescription *string `json:"serviceUpdateDescription,omitempty"`

	ServiceUpdateName *string `json:"serviceUpdateName,omitempty"`
}

// +kubebuilder:skipversion
type Snapshot struct {
	ARN *string `json:"arn,omitempty"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty"`

	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	CacheNodeType *string `json:"cacheNodeType,omitempty"`

	CacheParameterGroupName *string `json:"cacheParameterGroupName,omitempty"`

	CacheSubnetGroupName *string `json:"cacheSubnetGroupName,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	NumCacheNodes *int64 `json:"numCacheNodes,omitempty"`

	NumNodeGroups *int64 `json:"numNodeGroups,omitempty"`

	Port *int64 `json:"port,omitempty"`

	PreferredAvailabilityZone *string `json:"preferredAvailabilityZone,omitempty"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`

	PreferredOutpostARN *string `json:"preferredOutpostARN,omitempty"`

	ReplicationGroupDescription *string `json:"replicationGroupDescription,omitempty"`

	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	SnapshotName *string `json:"snapshotName,omitempty"`

	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty"`

	SnapshotSource *string `json:"snapshotSource,omitempty"`

	SnapshotStatus *string `json:"snapshotStatus,omitempty"`

	SnapshotWindow *string `json:"snapshotWindow,omitempty"`

	TopicARN *string `json:"topicARN,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

// +kubebuilder:skipversion
type Subnet struct {
	SubnetIdentifier *string `json:"subnetIdentifier,omitempty"`
}

// +kubebuilder:skipversion
type SubnetOutpost struct {
	SubnetOutpostARN *string `json:"subnetOutpostARN,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type UnprocessedUpdateAction struct {
	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorType *string `json:"errorType,omitempty"`

	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	ServiceUpdateName *string `json:"serviceUpdateName,omitempty"`
}

// +kubebuilder:skipversion
type UpdateAction struct {
	CacheClusterID *string `json:"cacheClusterID,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EstimatedUpdateTime *string `json:"estimatedUpdateTime,omitempty"`

	NodesUpdated *string `json:"nodesUpdated,omitempty"`

	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`

	ServiceUpdateName *string `json:"serviceUpdateName,omitempty"`
}

// +kubebuilder:skipversion
type User struct {
	ARN *string `json:"arn,omitempty"`

	AccessString *string `json:"accessString,omitempty"`

	Status *string `json:"status,omitempty"`

	UserID *string `json:"userID,omitempty"`

	UserName *string `json:"userName,omitempty"`
}

// +kubebuilder:skipversion
type UserGroup struct {
	ARN *string `json:"arn,omitempty"`

	Status *string `json:"status,omitempty"`

	UserGroupID *string `json:"userGroupID,omitempty"`
}
