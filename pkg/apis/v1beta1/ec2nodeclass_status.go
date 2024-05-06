/*
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

package v1beta1

import (
	corev1beta1 "sigs.k8s.io/karpenter/pkg/apis/v1beta1"
)

// Subnet contains resolved Subnet selector values utilized for node launch
type Subnet struct {
	// ID of the subnet
	// +required
	ID string `json:"id"`
	// The associated availability zone
	// +required
	Zone string `json:"zone"`
}

// SecurityGroup contains resolved SecurityGroup selector values utilized for node launch
type SecurityGroup struct {
	// ID of the security group
	// +required
	ID string `json:"id"`
	// Name of the security group
	// +optional
	Name string `json:"name,omitempty"`
}

// CapacityReservation contains resolved Capacity Reservation selector values utilized for node launch
type CapacityReservation struct {
	// ID of the Capacity Reservation
	// +required
	ID string `json:"id"`
	// AvailabilityZone of the Capacity Reservation
	// +required
	AvailabilityZone string `json:"availabilityZone"`
	// Available Instance Count of the Capacity Reservation
	// +required
	AvailableInstanceCount int `json:"availableInstanceCount"`
	// InstanceType of the Capacity Reservation
	// +required
	InstanceType string `json:"instanceType"`
	// Requirements of the Capacity Reservation to be utilized on an instance type
	// +required
	Requirements []corev1beta1.NodeSelectorRequirementWithMinValues `json:"requirements"`
	// Total Instance Count of the Capacity Reservation
	// +required
	TotalInstanceCount int `json:"totalInstanceCount"`
}

// AMI contains resolved AMI selector values utilized for node launch
type AMI struct {
	// ID of the AMI
	// +required
	ID string `json:"id"`
	// Name of the AMI
	// +optional
	Name string `json:"name,omitempty"`
	// Requirements of the AMI to be utilized on an instance type
	// +required
	Requirements []corev1beta1.NodeSelectorRequirementWithMinValues `json:"requirements"`
}

// EC2NodeClassStatus contains the resolved state of the EC2NodeClass
type EC2NodeClassStatus struct {
	// CapacityReservations contains the current Capacity Reservations values that are available to the
	// cluster under the CapacityReservations selectors.
	// +optional
	CapacityReservations []CapacityReservation `json:"capacityReservations,omitempty"`
	// Subnets contains the current Subnet values that are available to the
	// cluster under the subnet selectors.
	// +optional
	Subnets []Subnet `json:"subnets,omitempty"`
	// SecurityGroups contains the current Security Groups values that are available to the
	// cluster under the SecurityGroups selectors.
	// +optional
	SecurityGroups []SecurityGroup `json:"securityGroups,omitempty"`
	// AMI contains the current AMI values that are available to the
	// cluster under the AMI selectors.
	// +optional
	AMIs []AMI `json:"amis,omitempty"`
	// InstanceProfile contains the resolved instance profile for the role
	// +optional
	InstanceProfile string `json:"instanceProfile,omitempty"`
}
