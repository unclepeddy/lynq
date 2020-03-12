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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlbumSpec defines the desired state of Album
type AlbumSpec struct {
	// A persistent identifier used between sessions to identify the album
	Id string `json:"id,omitempty"`
	// User-visible name of the album, maximum of 500 chars
	Title string `json:"title,omitempty"`
	// Google Photos URL for the album, requires signed-in client
	ProductUrl string `json:"productUrl,omitempty"`
}

// AlbumStatus defines the observed state of Album
type AlbumStatus struct {
	// The number of media items in the album
	MediaItemsCount int64 `json:"mediaItemsCount,omitempty"`
	// Last time a successful sync took place
	LastSyncTime *metav1.Time `json:"lastSyncTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Album is the Schema for the albums API
type Album struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbumSpec   `json:"spec,omitempty"`
	Status AlbumStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbumList contains a list of Album
type AlbumList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Album `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Album{}, &AlbumList{})
}
