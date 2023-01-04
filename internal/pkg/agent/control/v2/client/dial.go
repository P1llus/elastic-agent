// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !windows
// +build !windows

package client

import (
	"context"
	"net"
	"strings"

	"github.com/elastic/elastic-agent/internal/pkg/agent/configuration"
	"github.com/elastic/elastic-agent/internal/pkg/agent/control"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func dialContext(ctx context.Context, grpcConfig *configuration.GRPCConfig) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		ctx,
		strings.TrimPrefix(control.Address(), "unix://"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcConfig.MaxMsgSize)),
	)
}

func dialer(ctx context.Context, addr string) (net.Conn, error) {
	var d net.Dialer
	return d.DialContext(ctx, "unix", addr)
}