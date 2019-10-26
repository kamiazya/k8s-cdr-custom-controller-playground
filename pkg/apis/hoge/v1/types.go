package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Hoge is Hoge resource.
type Hoge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HogeSpec `json:"spec"`
}

// HogeSpec is a spec of Hoge.
type HogeSpec struct {
}

// HogeList is a list of Hoge resources.
type HogeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Hoge `json:"items"`
}
