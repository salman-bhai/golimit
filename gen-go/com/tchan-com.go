// @generated Code generated by thrift-gen. Do not modify.

// Package com is generated code used to make or handle TChannel calls using Thrift.
package com

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanStoreNode is the interface that defines the server handler and client interface.
type TChanStoreNode interface {
	IncrAction(ctx thrift.Context, key string, count int32, threshold int32, window int32, peakaveraged bool) (bool, error)
	RateLimitGlobalAction(ctx thrift.Context, key string, count int32) (bool, error)
	SyncKeys(ctx thrift.Context, syncs []*SyncCommand) error
	SyncRateConfig(ctx thrift.Context, key string, threshold int32, window int32, peakaveraged bool) error
}

// Implementation of a client and service handler.

type tchanStoreNodeClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanStoreNodeInheritedClient(thriftService string, client thrift.TChanClient) *tchanStoreNodeClient {
	return &tchanStoreNodeClient{
		thriftService,
		client,
	}
}

// NewTChanStoreNodeClient creates a client that can be used to make remote calls.
func NewTChanStoreNodeClient(client thrift.TChanClient) TChanStoreNode {
	return NewTChanStoreNodeInheritedClient("StoreNode", client)
}

func (c *tchanStoreNodeClient) IncrAction(ctx thrift.Context, key string, count int32, threshold int32, window int32, peakaveraged bool) (bool, error) {
	var resp StoreNodeIncrActionResult
	args := StoreNodeIncrActionArgs{
		Key:          key,
		Count:        count,
		Threshold:    threshold,
		Window:       window,
		Peakaveraged: peakaveraged,
	}
	success, err := c.client.Call(ctx, c.thriftService, "IncrAction", &args, &resp)
	if err == nil && !success {
		switch {
		default:
			err = fmt.Errorf("received no result or unknown exception for IncrAction")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanStoreNodeClient) RateLimitGlobalAction(ctx thrift.Context, key string, count int32) (bool, error) {
	var resp StoreNodeRateLimitGlobalActionResult
	args := StoreNodeRateLimitGlobalActionArgs{
		Key:   key,
		Count: count,
	}
	success, err := c.client.Call(ctx, c.thriftService, "RateLimitGlobalAction", &args, &resp)
	if err == nil && !success {
		switch {
		default:
			err = fmt.Errorf("received no result or unknown exception for RateLimitGlobalAction")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanStoreNodeClient) SyncKeys(ctx thrift.Context, syncs []*SyncCommand) error {
	var resp StoreNodeSyncKeysResult
	args := StoreNodeSyncKeysArgs{
		Syncs: syncs,
	}
	success, err := c.client.Call(ctx, c.thriftService, "SyncKeys", &args, &resp)
	if err == nil && !success {
		switch {
		default:
			err = fmt.Errorf("received no result or unknown exception for SyncKeys")
		}
	}

	return err
}

func (c *tchanStoreNodeClient) SyncRateConfig(ctx thrift.Context, key string, threshold int32, window int32, peakaveraged bool) error {
	var resp StoreNodeSyncRateConfigResult
	args := StoreNodeSyncRateConfigArgs{
		Key:          key,
		Threshold:    threshold,
		Window:       window,
		Peakaveraged: peakaveraged,
	}
	success, err := c.client.Call(ctx, c.thriftService, "SyncRateConfig", &args, &resp)
	if err == nil && !success {
		switch {
		default:
			err = fmt.Errorf("received no result or unknown exception for SyncRateConfig")
		}
	}

	return err
}

type tchanStoreNodeServer struct {
	handler TChanStoreNode
}

// NewTChanStoreNodeServer wraps a handler for TChanStoreNode so it can be
// registered with a thrift.Server.
func NewTChanStoreNodeServer(handler TChanStoreNode) thrift.TChanServer {
	return &tchanStoreNodeServer{
		handler,
	}
}

func (s *tchanStoreNodeServer) Service() string {
	return "StoreNode"
}

func (s *tchanStoreNodeServer) Methods() []string {
	return []string{
		"IncrAction",
		"RateLimitGlobalAction",
		"SyncKeys",
		"SyncRateConfig",
	}
}

func (s *tchanStoreNodeServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "IncrAction":
		return s.handleIncrAction(ctx, protocol)
	case "RateLimitGlobalAction":
		return s.handleRateLimitGlobalAction(ctx, protocol)
	case "SyncKeys":
		return s.handleSyncKeys(ctx, protocol)
	case "SyncRateConfig":
		return s.handleSyncRateConfig(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanStoreNodeServer) handleIncrAction(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req StoreNodeIncrActionArgs
	var res StoreNodeIncrActionResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.IncrAction(ctx, req.Key, req.Count, req.Threshold, req.Window, req.Peakaveraged)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = &r
	}

	return err == nil, &res, nil
}

func (s *tchanStoreNodeServer) handleRateLimitGlobalAction(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req StoreNodeRateLimitGlobalActionArgs
	var res StoreNodeRateLimitGlobalActionResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.RateLimitGlobalAction(ctx, req.Key, req.Count)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = &r
	}

	return err == nil, &res, nil
}

func (s *tchanStoreNodeServer) handleSyncKeys(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req StoreNodeSyncKeysArgs
	var res StoreNodeSyncKeysResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.SyncKeys(ctx, req.Syncs)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanStoreNodeServer) handleSyncRateConfig(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req StoreNodeSyncRateConfigArgs
	var res StoreNodeSyncRateConfigResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.SyncRateConfig(ctx, req.Key, req.Threshold, req.Window, req.Peakaveraged)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}
