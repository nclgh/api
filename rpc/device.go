package rpc

import (
	"fmt"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
)

func AddManufacturer(ctx *client.RpcRequestCtx, name string) (*device.AddManufacturerResponse, error) {
	req := device.AddManufacturerRequest{
		Name: name,
	}
	rsp := &device.AddManufacturerResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "AddManufacturer", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call AddManufacturer failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func GetManufacturerById(ctx *client.RpcRequestCtx, ids []int64) (*device.GetManufacturerByIdResponse, error) {
	req := device.GetManufacturerByIdRequest{
		Ids: ids,
	}
	rsp := &device.GetManufacturerByIdResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "GetManufacturerById", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call GetManufacturerById failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func QueryManufacturer(ctx *client.RpcRequestCtx, m *device.Manufacturer, page, pageSize int64) (*device.QueryManufacturerResponse, error) {
	req := device.QueryManufacturerRequest{
		Manufacturer: m,
		Page:         page,
		PageSize:     pageSize,
	}
	rsp := &device.QueryManufacturerResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "QueryManufacturer", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call QueryManufacturer failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func AddDevice(ctx *client.RpcRequestCtx, d device.Device) (*device.AddDeviceResponse, error) {
	req := device.AddDeviceRequest{
		Device: d,
	}
	rsp := &device.AddDeviceResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "AddDevice", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call AddDevice failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func DeleteDevice(ctx *client.RpcRequestCtx, id int64) (*device.DeleteDeviceResponse, error) {
	req := device.DeleteDeviceRequest{
		Id: id,
	}
	rsp := &device.DeleteDeviceResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "DeleteDevice", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call DeleteDevice failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func GetDeviceById(ctx *client.RpcRequestCtx, ids []int64) (*device.GetDeviceByIdResponse, error) {
	req := device.GetDeviceByIdRequest{
		Ids: ids,
	}
	rsp := &device.GetDeviceByIdResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "GetDeviceById", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call GetDeviceById failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func QueryDevice(ctx *client.RpcRequestCtx, d *device.Device, page, pageSize int64, filter []*device.TimeFilter) (*device.QueryDeviceResponse, error) {
	req := device.QueryDeviceRequest{
		Device:     d,
		TimeFilter: filter,
		Page:       page,
		PageSize:   pageSize,
	}
	rsp := &device.QueryDeviceResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "QueryDevice", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call QueryDevice failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}
func AddAchievement(ctx *client.RpcRequestCtx, d device.Achievement) (*device.AddAchievementResponse, error) {
	req := device.AddAchievementRequest{
		Achievement: d,
	}
	rsp := &device.AddAchievementResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "AddAchievement", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call AddAchievement failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func DeleteAchievement(ctx *client.RpcRequestCtx, id int64) (*device.DeleteAchievementResponse, error) {
	req := device.DeleteAchievementRequest{
		Id: id,
	}
	rsp := &device.DeleteAchievementResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "DeleteAchievement", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call DeleteAchievement failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func GetAchievementById(ctx *client.RpcRequestCtx, ids []int64) (*device.GetAchievementByIdResponse, error) {
	req := device.GetAchievementByIdRequest{
		Ids: ids,
	}
	rsp := &device.GetAchievementByIdResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "GetAchievementById", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call GetAchievementById failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}

func QueryAchievement(ctx *client.RpcRequestCtx, d *device.Achievement, page, pageSize int64, filter []*device.TimeFilter) (*device.QueryAchievementResponse, error) {
	req := device.QueryAchievementRequest{
		Achievement: d,
		TimeFilter:  filter,
		Page:        page,
		PageSize:    pageSize,
	}
	rsp := &device.QueryAchievementResponse{}
	err := GetDeviceClient().Call(&client.RpcRequestCtx{}, "QueryAchievement", req, rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Code != common.CodeSuccess {
		return nil, fmt.Errorf("call QueryAchievement failed. code: %v, msg: %v", rsp.Code, rsp.Msg)
	}
	return rsp, nil
}
