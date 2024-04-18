// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
//

package utility

import (
	"testing"

	"github.com/mattermost/mattermost-cloud/model"

	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNewHelmDeploymentWithDefaultConfigurationMetricsServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := log.New()
	metricsServer := newMetricsServerHandle(&model.HelmUtilityVersion{Chart: "3.12.1"}, &model.Cluster{
		UtilityMetadata: &model.UtilityMetadata{
			ActualVersions: model.UtilityGroupVersions{},
		},
	}, "kubeconfig", logger)
	require.NoError(t, metricsServer.validate(), "should not error when creating new metrics server handler")
	require.NotNil(t, metricsServer, "metrics server should not be nil")

	helmDeployment := metricsServer.newHelmDeployment(logger)
	require.NotNil(t, helmDeployment, "helmDeployment should not be nil")
}
