/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableDevice is an enum
type OpcuaNodeIdServicesVariableDevice int32

type IOpcuaNodeIdServicesVariableDevice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventId               OpcuaNodeIdServicesVariableDevice = 3662
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventType             OpcuaNodeIdServicesVariableDevice = 3663
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceNode            OpcuaNodeIdServicesVariableDevice = 3664
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceName            OpcuaNodeIdServicesVariableDevice = 3665
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Time                  OpcuaNodeIdServicesVariableDevice = 3666
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ReceiveTime           OpcuaNodeIdServicesVariableDevice = 3667
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_LocalTime             OpcuaNodeIdServicesVariableDevice = 3668
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Message               OpcuaNodeIdServicesVariableDevice = 3669
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Severity              OpcuaNodeIdServicesVariableDevice = 3670
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassId      OpcuaNodeIdServicesVariableDevice = 31879
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassName    OpcuaNodeIdServicesVariableDevice = 31880
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassId   OpcuaNodeIdServicesVariableDevice = 31881
	OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassName OpcuaNodeIdServicesVariableDevice = 31882
)

var OpcuaNodeIdServicesVariableDeviceValues []OpcuaNodeIdServicesVariableDevice

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableDeviceValues = []OpcuaNodeIdServicesVariableDevice{
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventId,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventType,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceNode,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceName,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Time,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_LocalTime,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Message,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Severity,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassName,
	}
}

func OpcuaNodeIdServicesVariableDeviceByValue(value int32) (enum OpcuaNodeIdServicesVariableDevice, ok bool) {
	switch value {
	case 31879:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassId, true
	case 31880:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassName, true
	case 31881:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassId, true
	case 31882:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassName, true
	case 3662:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventId, true
	case 3663:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventType, true
	case 3664:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceNode, true
	case 3665:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceName, true
	case 3666:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Time, true
	case 3667:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ReceiveTime, true
	case 3668:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_LocalTime, true
	case 3669:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Message, true
	case 3670:
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDeviceByName(value string) (enum OpcuaNodeIdServicesVariableDevice, ok bool) {
	switch value {
	case "DeviceFailureEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassId, true
	case "DeviceFailureEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassName, true
	case "DeviceFailureEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassId, true
	case "DeviceFailureEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassName, true
	case "DeviceFailureEventType_EventId":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventId, true
	case "DeviceFailureEventType_EventType":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventType, true
	case "DeviceFailureEventType_SourceNode":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceNode, true
	case "DeviceFailureEventType_SourceName":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceName, true
	case "DeviceFailureEventType_Time":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Time, true
	case "DeviceFailureEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ReceiveTime, true
	case "DeviceFailureEventType_LocalTime":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_LocalTime, true
	case "DeviceFailureEventType_Message":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Message, true
	case "DeviceFailureEventType_Severity":
		return OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDeviceKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableDeviceValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableDevice(structType any) OpcuaNodeIdServicesVariableDevice {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableDevice {
		if sOpcuaNodeIdServicesVariableDevice, ok := typ.(OpcuaNodeIdServicesVariableDevice); ok {
			return sOpcuaNodeIdServicesVariableDevice
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableDevice) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableDevice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableDeviceParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableDevice, error) {
	return OpcuaNodeIdServicesVariableDeviceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableDeviceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableDevice, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableDevice", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableDevice")
	}
	if enum, ok := OpcuaNodeIdServicesVariableDeviceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableDevice")
		return OpcuaNodeIdServicesVariableDevice(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableDevice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableDevice", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableDevice) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassId:
		return "DeviceFailureEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionClassName:
		return "DeviceFailureEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassId:
		return "DeviceFailureEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ConditionSubClassName:
		return "DeviceFailureEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventId:
		return "DeviceFailureEventType_EventId"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_EventType:
		return "DeviceFailureEventType_EventType"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceNode:
		return "DeviceFailureEventType_SourceNode"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_SourceName:
		return "DeviceFailureEventType_SourceName"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Time:
		return "DeviceFailureEventType_Time"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_ReceiveTime:
		return "DeviceFailureEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_LocalTime:
		return "DeviceFailureEventType_LocalTime"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Message:
		return "DeviceFailureEventType_Message"
	case OpcuaNodeIdServicesVariableDevice_DeviceFailureEventType_Severity:
		return "DeviceFailureEventType_Severity"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableDevice) String() string {
	return e.PLC4XEnumName()
}
