package main

import (
	"context"

	dpv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type SanityDevicePluginServer struct {
	// server *grpc.Server
}

//

// DevicePluginServer

func (sdps *SanityDevicePluginServer) GetDevicePluginOptions(context.Context, *dpv1beta1.Empty) (*dpv1beta1.DevicePluginOptions, error) {
	return nil, nil
}

func (sdps *SanityDevicePluginServer) ListAndWatch(*dpv1beta1.Empty, dpv1beta1.DevicePlugin_ListAndWatchServer) error {
	return nil
}

func (sdps *SanityDevicePluginServer) GetPreferredAllocation(context.Context, *dpv1beta1.PreferredAllocationRequest) (*dpv1beta1.PreferredAllocationResponse, error) {
	return nil, nil
}

func (sdps *SanityDevicePluginServer) Allocate(context.Context, *dpv1beta1.AllocateRequest) (*dpv1beta1.AllocateResponse, error) {
	return nil, nil
}

func (sdps *SanityDevicePluginServer) PreStartContainer(context.Context, *dpv1beta1.PreStartContainerRequest) (*dpv1beta1.PreStartContainerResponse, error) {
	return nil, nil
}

// func (sdp *SanityDevicePlugin) Serve() {
// 	dpv1beta1.RegisterDevicePluginServer(sdp.server, sdp)
// }

func main() {
}
