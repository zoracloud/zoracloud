package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Podspec for the notebooks

type NotebookTemplateSpec struct {
	Spec corev1.PodSpec `json:"spec,omitempty"`
}

// Notebook Condition

type NotebookCondition struct {
	// Type is the type of the condition. Possible values are | Running | Waiting | Terminated
	Type string `json:"type"`

	// Last time we probed the condition/
	//+optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`

	//(brief) reason the container is in the current state
	// + optional
	Reason string `json:"reason,omitempty"`

	//Message regarding why the container is in the current state
	//+optional
	Message string `json:"message,omitempty"`
}

// NotebookStatus defines the observed state of Notebook
type NotebookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Conditions is an array of the current conditions
	Conditions []NotebookCondition `json:"conditions"`

	// Replicas is the number of Pods created by the Statefulset Controller that have a ready Condition
	ReadyReplicas int32 `json:"readyReplicas"`

	// Container State is the state of the underlying container
	ContainerState corev1.ContainerState `json:"containerState"`
}
