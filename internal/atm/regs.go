/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package atm

type Register interface {
    id() uint8
}

type (
    GenericRegister uint8
    PointerRegister uint8
)

const (
    ArgMask    = 0x7f
    ArgGeneric = 0x00
    ArgPointer = 0x80
)

const (
    R0 GenericRegister = iota
    R1
    R2
    R3
    R4
    R5
    R6
    R7
    Rz
)

const (
    P0 PointerRegister = iota
    P1
    P2
    P3
    P4
    P5
    P6
    P7
    LR
    Pn
)

func (self GenericRegister) id() uint8 { return uint8(self) | ArgGeneric }
func (self PointerRegister) id() uint8 { return uint8(self) | ArgPointer }