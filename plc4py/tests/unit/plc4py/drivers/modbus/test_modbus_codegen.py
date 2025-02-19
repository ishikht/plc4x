#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#
import pytest

from plc4py.protocols.modbus.readwrite.ModbusTcpADU import ModbusTcpADUBuilder
from plc4py.spi.generation.WriteBuffer import WriteBufferByteBased

from plc4py.protocols.modbus.readwrite.ModbusPDUReadDiscreteInputsRequest import (
    ModbusPDUReadDiscreteInputsRequestBuilder,
)
from plc4py.utils.GenericTypes import ByteOrder


@pytest.mark.asyncio
async def test_modbus_discrete_inputs_request():
    request = ModbusPDUReadDiscreteInputsRequestBuilder(0, 10).build()
    assert request is not None


@pytest.mark.asyncio
async def test_modbus_discrete_inputs_request_serialize():
    request = ModbusPDUReadDiscreteInputsRequestBuilder(5, 2).build()
    size = request.length_in_bytes()
    write_buffer = WriteBufferByteBased(size, ByteOrder.BIG_ENDIAN)
    serialize = request.serialize(write_buffer)
    bytes_array = write_buffer.get_bytes().tobytes()

    assert request is not None
    assert len(write_buffer.get_bytes()) * 8 == 40
    assert write_buffer.get_pos() == 40
    assert write_buffer.get_bytes().tobytes() == b"\x02\x00\x05\x00\x02"


@pytest.mark.asyncio
async def test_modbus_ModbusTcpADUBuilder_serialize():
    pdu = ModbusPDUReadDiscreteInputsRequestBuilder(5, 2).build()
    request = ModbusTcpADUBuilder(10, 5, pdu).build(False)
    size = request.length_in_bytes()
    write_buffer = WriteBufferByteBased(size, ByteOrder.BIG_ENDIAN)
    serialize = request.serialize(write_buffer)
    bytes_array = write_buffer.get_bytes().tobytes()

    assert request is not None
