// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gardenlet

import (
	"fmt"

	v2 "github.com/gardener/component-spec/bindings-go/apis/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/gardener/gardener/pkg/landscaper/gardenlet/apis/imports"
)

var _ = Describe("Gardenlet Landscaper testing", func() {
	var (
		landscaper              Landscaper
		componentDescriptorList *v2.ComponentDescriptorList
		expectedImageRepository = "eu.gcr.io/sap-se-gcr-k8s-public/eu_gcr_io/gardener-project/gardener/gardenlet"
		expectedImageVersion    = "v1.11.3"
	)

	BeforeEach(func() {
		landscaper = Landscaper{}
		componentDescriptorList = &v2.ComponentDescriptorList{
			Components: []v2.ComponentDescriptor{
				{
					ComponentSpec: v2.ComponentSpec{
						ObjectMeta: v2.ObjectMeta{
							Name:    "github.com/gardener/gardener",
							Version: expectedImageVersion,
						},
						RepositoryContexts: []v2.RepositoryContext{
							{
								Type:    "ociRegistry",
								BaseURL: "eu.gcr.io/gardener-project/gardener/gardenlet",
							},
						},
						Provider: "internal",
						LocalResources: []v2.Resource{
							{
								ObjectMeta: v2.ObjectMeta{
									Name:    "gardenlet",
									Version: expectedImageVersion,
									Labels:  nil,
								},
								TypedObjectAccessor: v2.NewTypeOnly(v2.OCIImageType),
								Access: &v2.OCIRegistryAccess{
									ObjectType:     v2.ObjectType{Type: v2.OCIImageType},
									ImageReference: fmt.Sprintf("%s:%s", expectedImageRepository, expectedImageVersion),
								},
							},
						},
					},
				},
			},
		}
	})

	Describe("#parseImageVectorOverride", func() {
		It("should parse the Image Vector", func() {
			content := "images: dummy"
			contentComponents := "images: dummy-components"
			imageVectorOverrideContent := fmt.Sprintf("\"%s\"", content)
			imageVectorOverrideComponentContent := fmt.Sprintf("\"%s\"", contentComponents)
			landscaper.Imports = &imports.LandscaperGardenletImport{
				ImageVectorOverwrite: &runtime.RawExtension{
					Raw: []byte(imageVectorOverrideContent),
				},
				ComponentImageVectorOverwrites: &runtime.RawExtension{
					Raw: []byte(imageVectorOverrideComponentContent),
				},
			}

			Expect(landscaper.parseImageVectorOverride()).ToNot(HaveOccurred())
			Expect(landscaper.imageVectorOverride).ToNot(BeNil())
			Expect(*landscaper.imageVectorOverride).To(Equal(content))
			Expect(landscaper.componentImageVectorOverwrites).ToNot(BeNil())
			Expect(*landscaper.componentImageVectorOverwrites).To(Equal(contentComponents))
		})
	})

	Describe("#parseGardenletImage", func() {
		It("should parse the Gardenlet image", func() {
			Expect(landscaper.parseGardenletImage(&*componentDescriptorList)).ToNot(HaveOccurred())
			Expect(landscaper.gardenletImageRepository).To(Equal(expectedImageRepository))
			Expect(landscaper.gardenletImageVersion).To(Equal(expectedImageVersion))
		})
		It("should return error - Component does not exist", func() {
			Expect(landscaper.parseGardenletImage(&v2.ComponentDescriptorList{})).To(HaveOccurred())
		})
		It("should return error - more than one component with expected name exists", func() {
			componentDescriptorList.Components = append(componentDescriptorList.Components, componentDescriptorList.Components[0])
			Expect(landscaper.parseGardenletImage(&v2.ComponentDescriptorList{})).To(HaveOccurred())
		})
		It("should return error - local resource with expected name does not exists", func() {
			componentDescriptorList.Components[0].LocalResources = []v2.Resource{}
			Expect(landscaper.parseGardenletImage(componentDescriptorList)).To(HaveOccurred())
		})
		It("should return error - local resource version not set", func() {
			componentDescriptorList.Components[0].LocalResources[0].Version = ""
			Expect(landscaper.parseGardenletImage(componentDescriptorList)).To(HaveOccurred())
		})
		It("should return error - local resource image reference not set", func() {
			componentDescriptorList.Components[0].LocalResources[0].Access.(*v2.OCIRegistryAccess).ImageReference = ""
			Expect(landscaper.parseGardenletImage(componentDescriptorList)).To(HaveOccurred())
		})
		It("should return error - local resource image reference invalid", func() {
			componentDescriptorList.Components[0].LocalResources[0].Access.(*v2.OCIRegistryAccess).ImageReference = "invalid"
			Expect(landscaper.parseGardenletImage(componentDescriptorList)).To(HaveOccurred())
		})
	})

})
