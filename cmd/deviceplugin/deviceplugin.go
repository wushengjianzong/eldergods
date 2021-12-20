package main

import (
	"context"
	"log"
	"net"
	"os"
	"path"
	"time"

	"google.golang.org/grpc"
	dpv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

	"eldergods/pkg/utils"
)

type SanityDevicePlugin struct {
	resourceName string
	server       *grpc.Server
	sockPath     string
}

func (sdp *SanityDevicePlugin) GetDevicePluginOptions(context.Context, *dpv1beta1.Empty) (*dpv1beta1.DevicePluginOptions, error) {
	return nil, nil
}

func (sdp *SanityDevicePlugin) ListAndWatch(*dpv1beta1.Empty, dpv1beta1.DevicePlugin_ListAndWatchServer) error {
	return nil
}

func (sdp *SanityDevicePlugin) GetPreferredAllocation(context.Context, *dpv1beta1.PreferredAllocationRequest) (*dpv1beta1.PreferredAllocationResponse, error) {
	return nil, nil
}

func (sdp *SanityDevicePlugin) Allocate(context.Context, *dpv1beta1.AllocateRequest) (*dpv1beta1.AllocateResponse, error) {
	return nil, nil
}

func (sdp *SanityDevicePlugin) PreStartContainer(context.Context, *dpv1beta1.PreStartContainerRequest) (*dpv1beta1.PreStartContainerResponse, error) {
	return nil, nil
}

func (sdp *SanityDevicePlugin) Register() error {
	conn, err := utils.DialGrpcSocketWithTimeout(dpv1beta1.KubeletSocket, 5*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()

	dpclient := dpv1beta1.NewRegistrationClient(conn)
	if _, err := dpclient.Register(context.TODO(), &dpv1beta1.RegisterRequest{
		Version:      dpv1beta1.Version,
		Endpoint:     path.Base(sdp.sockPath),
		ResourceName: sdp.resourceName,
		Options:      &dpv1beta1.DevicePluginOptions{},
	}); err != nil {
		return err
	}

	return nil
}

func (sdp *SanityDevicePlugin) Serve() error {
	if err := os.Remove(sdp.sockPath); err != nil {
		log.Println("remove plugin socket:", err)
	}

	sock, err := net.Listen("unix", sdp.sockPath)
	if err != nil {
		return err
	}

	go func(listener net.Listener) {
		if err := sdp.server.Serve(listener); err != nil {
			log.Fatalln("start grpc server:", err)
		}
	}(sock)
	log.Println("grpc server start at", sdp.sockPath)

	return nil
}

func (sdp *SanityDevicePlugin) Start() error {
	if err := sdp.Serve(); err != nil {
		log.Fatalln("plugin serve:", err)
	}
	if err := sdp.Register(); err != nil {
		log.Fatalln("plugin register:", err)
	}
	return nil
}

func main() {

}
