package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bot describe a Bot resource
type Bot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the CR spec
	Spec BotSpec `json:"spec"`
}

// BotSpec is the spec for a Bot resource
type BotSpec struct {
	Code string `json:"code"`
	Role string `json:"role"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BotList is a list of Bot resources
type BotList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Items []Bot `json:"items"`
}
