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

package rt

import (
    `unsafe`
)

//go:noescape
//go:linkname mapclear runtime.mapclear
//goland:noinspection GoUnusedParameter
func mapclear(t *GoType, h unsafe.Pointer)

//go:linkname mallocgc runtime.mallocgc
//goland:noinspection GoUnusedParameter
func mallocgc(size uintptr, typ *GoType, needzero bool) unsafe.Pointer

//go:nosplit
func MapClear(m interface{}) {
    v := UnpackEface(m)
    mapclear(v.Type, v.Value)
}

//go:nosplit
func MallocGC(nb uintptr, vt *GoType, zero bool) unsafe.Pointer {
    return mallocgc(nb, vt, zero)
}
