/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"fmt"

	kyvernov1 "github.com/nirmata/kyverno/pkg/api/kyverno/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// ClusterPolicyListerExpansion allows custom methods to be added to
// ClusterPolicyLister.
type ClusterPolicyListerExpansion interface {
	GetPolicyForPolicyViolation(pv *kyvernov1.ClusterPolicyViolation) ([]*kyvernov1.ClusterPolicy, error)
	GetPolicyForNamespacedPolicyViolation(pv *kyvernov1.NamespacedPolicyViolation) ([]*kyvernov1.ClusterPolicy, error)
	ListResources(selector labels.Selector) (ret []*kyvernov1.ClusterPolicy, err error)
}

// ClusterPolicyViolationListerExpansion allows custom methods to be added to
// ClusterPolicyViolationLister.
type ClusterPolicyViolationListerExpansion interface {
	// List lists all PolicyViolations in the indexer with GVK.
	ListResources(selector labels.Selector) (ret []*kyvernov1.ClusterPolicyViolation, err error)
}

// NamespacedPolicyViolationListerExpansion allows custom methods to be added to
// NamespacedPolicyViolationLister.
type NamespacedPolicyViolationListerExpansion interface{}

// NamespacedPolicyViolationNamespaceListerExpansion allows custom methods to be added to
// NamespacedPolicyViolationNamespaceLister.
type NamespacedPolicyViolationNamespaceListerExpansion interface{}

//ListResources is a wrapper to List and adds the resource kind information
// as the lister is specific to a gvk we can harcode the values here
func (pvl *clusterPolicyViolationLister) ListResources(selector labels.Selector) (ret []*kyvernov1.ClusterPolicyViolation, err error) {
	policyviolations, err := pvl.List(selector)
	for index := range policyviolations {
		policyviolations[index].SetGroupVersionKind(kyvernov1.SchemeGroupVersion.WithKind("ClusterPolicyViolation"))
	}
	return policyviolations, nil
}

//ListResources is a wrapper to List and adds the resource kind information
// as the lister is specific to a gvk we can harcode the values here
func (pl *clusterPolicyLister) ListResources(selector labels.Selector) (ret []*kyvernov1.ClusterPolicy, err error) {
	policies, err := pl.List(selector)
	for index := range policies {
		policies[index].SetGroupVersionKind(kyvernov1.SchemeGroupVersion.WithKind("ClusterPolicy"))
	}
	return policies, err
}

func (pl *clusterPolicyLister) GetPolicyForPolicyViolation(pv *kyvernov1.ClusterPolicyViolation) ([]*kyvernov1.ClusterPolicy, error) {
	if len(pv.Labels) == 0 {
		return nil, fmt.Errorf("no Policy found for PolicyViolation %v because it has no labels", pv.Name)
	}

	pList, err := pl.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var policies []*kyvernov1.ClusterPolicy
	for _, p := range pList {
		policyLabelmap := map[string]string{"policy": p.Name}

		ls := &metav1.LabelSelector{}
		err = metav1.Convert_Map_string_To_string_To_v1_LabelSelector(&policyLabelmap, ls, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to generate label sector of Policy name %s: %v", p.Name, err)
		}
		selector, err := metav1.LabelSelectorAsSelector(ls)
		if err != nil {
			return nil, fmt.Errorf("invalid label selector: %v", err)
		}
		// If a policy with a nil or empty selector creeps in, it should match nothing, not everything.
		if selector.Empty() || !selector.Matches(labels.Set(pv.Labels)) {
			continue
		}
		policies = append(policies, p)
	}

	if len(policies) == 0 {
		return nil, fmt.Errorf("could not find Policy set for PolicyViolation %s with labels: %v", pv.Name, pv.Labels)
	}

	return policies, nil

}

func (pl *clusterPolicyLister) GetPolicyForNamespacedPolicyViolation(pv *kyvernov1.NamespacedPolicyViolation) ([]*kyvernov1.ClusterPolicy, error) {
	if len(pv.Labels) == 0 {
		return nil, fmt.Errorf("no Policy found for PolicyViolation %v because it has no labels", pv.Name)
	}

	pList, err := pl.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var policies []*kyvernov1.ClusterPolicy
	for _, p := range pList {
		policyLabelmap := map[string]string{"policy": p.Name}

		ls := &metav1.LabelSelector{}
		err = metav1.Convert_Map_string_To_string_To_v1_LabelSelector(&policyLabelmap, ls, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to generate label sector of Policy name %s: %v", p.Name, err)
		}
		selector, err := metav1.LabelSelectorAsSelector(ls)
		if err != nil {
			return nil, fmt.Errorf("invalid label selector: %v", err)
		}
		// If a policy with a nil or empty selector creeps in, it should match nothing, not everything.
		if selector.Empty() || !selector.Matches(labels.Set(pv.Labels)) {
			continue
		}
		policies = append(policies, p)
	}

	if len(policies) == 0 {
		return nil, fmt.Errorf("could not find Policy set for Namespaced policy Violation %s with labels: %v", pv.Name, pv.Labels)
	}

	return policies, nil

}
