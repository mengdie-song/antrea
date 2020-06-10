// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DropReason string

type Phase string

const (
	INITIAL Phase = "INITIAL"
	RUNNING Phase = "RUNNING"
	SUCCESS Phase = "SUCCESS"
	TIMEOUT Phase = "TIMEOUT"
	ERROR   Phase = "ERROR"
)

type Component string

const (
	SPOOFGUARD    Component = "SPOOFGUARD"
	LB            Component = "LB"
	ROUTING       Component = "ROUTING"
	NETWORKPOLICY Component = "NETWORKPOLICY"
	FORWARDING    Component = "FORWARDING"
)

type Action string

const (
	DELIVERED Action = "DELIVERED"
	RECEIVED  Action = "RECEIVED"
	FORWARDED Action = "FORWARDED"
	DROPPED   Action = "DROPPED"
)

type PacketTransportType string

const (
	UNICAST PacketTransportType = "UNICAST"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Traceflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Source      `json:"source"`
	Destination `json:"destination"`
	RoundID     string `json:"roundID,omitempty"`

	Packet `json:"packet"`
	Status `json:"status"`
}

type Source struct {
	Namespace string `json:"namespace,omitempty"`
	Pod       string `json:"pod,omitempty"`
}

type Destination struct {
	Namespace string `json:"namespace,omitempty"`
	Pod       string `json:"pod,omitempty"`
	Service   string `json:"service,omitempty"`
	IP        string `json:"IP,omitempty"`
}

type IPHeader struct {
	SrcIP    string `json:"srcIP,omitempty"`
	Protocol int32  `json:"protocol,omitempty"`
	TTL      int32  `json:"ttl,omitempty"`
	Flags    int32  `json:"flags,omitempty"`
}

type TransportHeader struct {
	ICMPEchoRequestHeader `json:"icmpEchoRequestHeader"`
	UDPHeader             `json:"udpHeader"`
	TCPHeader             `json:"tcpHeader"`
}

type ICMPEchoRequestHeader struct {
	ID       int32 `json:"id,omitempty"`
	Sequence int32 `json:"sequence,omitempty"`
}

type UDPHeader struct {
	SrcPort int32 `json:"srcPort,omitempty"`
	DstPort int32 `json:"dstPort,omitempty"`
}

type TCPHeader struct {
	SrcPort int32 `json:"srcPort,omitempty"`
	DstPort int32 `json:"dstPort,omitempty"`
	Flags   int32 `json:"flags,omitempty"`
}

type Packet struct {
	TransportType PacketTransportType `json:"transportType,omitempty"`
	PayloadString string              `json:"payloadString,omitempty"`
	PayloadSize   int32               `json:"payloadSize,omitempty"`

	IPHeader        `json:"ipHeader"`
	TransportHeader `json:"transportHeader"`
}

type Status struct {
	Phase        Phase         `json:"phase,omitempty"`
	CrossNodeTag uint8         `json:"crossNodeTag,omitempty"` // unique ID for traceflow
	NodeSender   []Observation `json:"nodeSender,omitempty"`
	NodeReceiver []Observation `json:"nodeReceiver,omitempty"`
}

type Observation struct {
	Component       Component  `json:"component,omitempty"`
	ComponentName   string     `json:"componentName,omitempty"`
	Action          Action     `json:"action,omitempty"`
	RoundID         string     `json:"roundID,omitempty"`
	Node            string     `json:"node,omitempty"`
	PodUUID         string     `json:"podUUID,omitempty"`
	DstMAC          string     `json:"dstMAC,omitempty"`
	RuleID          string     `json:"ruleID,omitempty"`
	Rule            string     `json:"rule,omitempty"`
	NetworkPolicy   string     `json:"networkPolicy,omitempty"`
	TTL             int32      `json:"ttl,omitempty"`
	TranslatedSrcIP string     `json:"translatedSrcIP,omitempty"`
	TranslatedDstIP string     `json:"translatedDstIP,omitempty"`
	DropReason      DropReason `json:"dropReason,omitempty"`
	TunnelDstIP     string     `json:"tunnelDstIP,omitempty"`
	Timestamp       int64      `json:"timestamp,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TraceflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Traceflow `json:"items"`
}
