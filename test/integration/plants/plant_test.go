// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package plants_test

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"time"

	"k8s.io/api/core/v1"

	. "github.com/gardener/gardener/test/integration/shoots"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"github.com/gardener/gardener/pkg/logger"
	. "github.com/gardener/gardener/test/integration/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var (
	kubeconfig                    = flag.String("kubeconfig", "", "the path to the kubeconfig  of the garden cluster that will be used for integration tests")
	kubeconfigPathExternalCluster = flag.String("kubeconfigPathExternalCluster", "", "the path to the kubeconfig  of the external cluster that will be registered as a plant")
	logLevel                      = flag.String("verbose", "", "verbosity level, when set, logging level will be DEBUG")
	plantTestYamlPath             = flag.String("plantpath", "", "the path to the plant yaml that will be used for testing")
	cleanup                       = flag.Bool("cleanup", false, "deletes the newly created / existing test plant after the test suite is done")
)

const (
	PlantUpdateSecretTimeout = 90 * time.Second
	PlantCreationTimeout     = 60 * time.Second
	InitializationTimeout    = 20 * time.Second
	FinalizationTimeout      = 20 * time.Second
)

func validateFlags() {

	if !StringSet(*plantTestYamlPath) {
		Fail("You need to set the YAML path to the Plant that should be created")
	} else {
		if !FileExists(*plantTestYamlPath) {
			Fail("plant yaml path is set but invalid")
		}
	}

	if !StringSet(*kubeconfig) {
		Fail("you need to specify the correct path for the kubeconfig")
	}

	if !FileExists(*kubeconfig) {
		Fail("kubeconfig path does not exist")
	}

	if !StringSet(*kubeconfigPathExternalCluster) {
		Fail("you need to specify the correct path for the kubeconfig of the external cluster")
	}

	if !FileExists(*kubeconfigPathExternalCluster) {
		Fail("kubeconfigPathExternalCluster path does not exist")
	}
}

var _ = Describe("Plant testing", func() {
	var (
		plantTest       *PlantTest
		plantTestLogger *logrus.Logger
		targetTestPlant *gardencorev1alpha1.Plant
		rememberSecret  *v1.Secret
	)

	CBeforeSuite(func(ctx context.Context) {
		// validate flags
		validateFlags()
		plantTestLogger = logger.AddWriter(logger.NewLogger(*logLevel), GinkgoWriter)

		if StringSet(*plantTestYamlPath) {
			// parse plant yaml into plant object and generate random test names for plants
			plantObject, err := CreatePlantTestArtifacts(*plantTestYamlPath)
			Expect(err).NotTo(HaveOccurred())

			plantTest, err = NewPlantTest(*kubeconfig, *kubeconfigPathExternalCluster, plantObject, plantTestLogger)
			Expect(err).NotTo(HaveOccurred())
		}

	}, InitializationTimeout)

	CAfterSuite(func(ctx context.Context) {
		if *cleanup {
			By("Cleaning up test plant and secret")

			err := plantTest.DeletePlant(ctx)
			Expect(err).NotTo(HaveOccurred())

			err = plantTest.DeletePlantSecret(ctx)
			Expect(err).NotTo(HaveOccurred())
		}
	}, FinalizationTimeout)

	CIt("Should create plant successfully", func(ctx context.Context) {

		By(fmt.Sprintf("Create Plant secret from kubeconfig to external cluster"))

		err := plantTest.CreatePlantSecret(ctx)
		Expect(err).NotTo(HaveOccurred())
		*cleanup = true
		rememberSecret = plantTest.PlantSecret.DeepCopy()

		By(fmt.Sprintf("Create Plant"))

		targetTestPlant, err = plantTest.CreatePlant(ctx)
		Expect(err).NotTo(HaveOccurred())

		By(fmt.Sprintf("Check if created plant has successful status. Name of plant: %s", targetTestPlant.Name))

		plantTest.WaitForPlantToBeReconciledSuccessfully(ctx)

		Expect(err).NotTo(HaveOccurred())
	}, PlantCreationTimeout)

	CIt("Should update Plant Status to 'unknown' due to updated and invalid Plant Secret (kubeconfig invalid)", func(ctx context.Context) {

		secretToUpdate := plantTest.PlantSecret

		if secretToUpdate == nil {
			plantSecret, err := plantTest.GetPlantSecret(ctx)

			if err != nil {
				Fail("Cannot retrieve Plant Secret")
			}
			secretToUpdate = plantSecret
		}

		// modify data.kubeconfig to update the secret with false information
		dummyKubeconfigContent := "Here is a string...."
		source := []byte(dummyKubeconfigContent)
		base64DummyKubeconfig := base64.StdEncoding.EncodeToString(source)
		secretToUpdate.Data["kubeconfig"] = []byte(base64DummyKubeconfig)

		By(fmt.Sprintf("Update Plant secret with invalid kubeconfig"))

		err := plantTest.UpdatePlantSecret(ctx, secretToUpdate)
		Expect(err).NotTo(HaveOccurred())

		By(fmt.Sprintf("Wait for PlantController to update to status 'unknown'"))

		err = plantTest.WaitForPlantToBeReconciledWithUnknownStatus(ctx)
		Expect(err).NotTo(HaveOccurred())
	}, PlantUpdateSecretTimeout)

	CIt("Should reconcile Plant Status to be successful after Plant Secret update", func(ctx context.Context) {

		Expect(rememberSecret.Data).NotTo(Equal(nil))

		plantTest.PlantSecret.Data["kubeconfig"] = rememberSecret.Data["kubeconfig"]
		// Update secret again to contain valid kubeconfig
		err := plantTest.UpdatePlantSecret(ctx, plantTest.PlantSecret)
		Expect(err).NotTo(HaveOccurred())

		By(fmt.Sprintf("Plant secret updated to contain valid kubeconfig again"))

		plantTestLogger.Debugf("Checking if created plant has successful status. Name of plant: %s", targetTestPlant.Name)

		plantTest.WaitForPlantToBeReconciledSuccessfully(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(err).NotTo(HaveOccurred())
		By(fmt.Sprintf("Plant reconciled successfully"))

	}, PlantUpdateSecretTimeout)
	CIt("Should update Plant Status to 'unknown' due to updated and invalid Plant Secret (kubeconfig not provided)", func(ctx context.Context) {

		secretToUpdate := plantTest.PlantSecret

		if secretToUpdate == nil {
			plantSecret, err := plantTest.GetPlantSecret(ctx)

			if err != nil {
				Fail("Cannot retrieve Plant Secret")
			}
			secretToUpdate = plantSecret
		}

		// remove data.kubeconfig to update the secret with false information
		secretToUpdate.Data["kubeconfig"] = nil

		err := plantTest.UpdatePlantSecret(ctx, secretToUpdate)
		Expect(err).NotTo(HaveOccurred())

		By(fmt.Sprintf("Wait for PlantController to update to status 'unknown'"))

		err = plantTest.WaitForPlantToBeReconciledWithUnknownStatus(ctx)
		Expect(err).NotTo(HaveOccurred())

	}, PlantUpdateSecretTimeout)
	CIt("Should reconcile Plant Status to be successful after Plant Secret update", func(ctx context.Context) {

		Expect(rememberSecret.Data).NotTo(Equal(nil))

		plantTest.PlantSecret.Data["kubeconfig"] = rememberSecret.Data["kubeconfig"]
		// Update secret again to contain valid kubeconfig
		err := plantTest.UpdatePlantSecret(ctx, plantTest.PlantSecret)
		Expect(err).NotTo(HaveOccurred())

		By(fmt.Sprintf("Plant secret updated to contain valid kubeconfig again"))

		plantTestLogger.Debugf("Checking if created plant has successful status. Name of plant: %s", targetTestPlant.Name)

		plantTest.WaitForPlantToBeReconciledSuccessfully(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(err).NotTo(HaveOccurred())
		By(fmt.Sprintf("Plant reconciled successfully"))

	}, PlantUpdateSecretTimeout)
	CIt("Should update Plant Status to 'unknown' due to deleted Plant Secret", func(ctx context.Context) {

		err := plantTest.DeletePlantSecret(ctx)
		Expect(err).NotTo(HaveOccurred())

		By(fmt.Sprintf("Wait for PlantController to update to status 'unknown'"))

		err = plantTest.WaitForPlantToBeReconciledWithUnknownStatus(ctx)
		Expect(err).NotTo(HaveOccurred())

	}, PlantUpdateSecretTimeout)

	CIt("Should reconcile Plant Status to be successful after new PlantSecret has been created for an existing Plant", func(ctx context.Context) {

		err := plantTest.CreatePlantSecret(ctx)
		Expect(err).NotTo(HaveOccurred())

		plantTestLogger.Debugf("Checking if created plant has successful status. Name of plant: %s", targetTestPlant.Name)

		plantTest.WaitForPlantToBeReconciledSuccessfully(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(err).NotTo(HaveOccurred())
		By(fmt.Sprintf("Plant reconciled successfully"))

	}, PlantUpdateSecretTimeout)
})
