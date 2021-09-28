package controllers

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func nbNameFromInvolvedObject(c client.Client, object *corev1.ObjectReference) (string, error) {
	// get the name and the namespace
	name, namespace := object.Name, object.Namespace

	//if the kind is a stateteful set return name or nill
	if object.Kind == "StatefulSet" {
		return name, nil
	}

	// if kind  == pod
	if object.Kind == "Pod" {
		// get the pod
		pod := &corev1.Pod{}

		// if err
		err := c.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, pod)
		if err != nil {
			return "", err
		}
		if nbName, ok := pod.Labels["notebook-name"]; ok {
			return nbName, nil
		}
	}
	return "", fmt.Errorf("The object isn't related to a Notebook")
}
