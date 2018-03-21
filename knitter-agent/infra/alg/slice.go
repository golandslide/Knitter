/*
Copyright 2018 ZTE Corporation. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package alg

import (
	"github.com/HyperNetworks/Knitter/pkg/klog"
)

type Slice []interface{}

type Comparable interface {
	IsEqual(obj interface{}) bool
}

func isEqual(a, b interface{}) bool {
	if comparable, ok := a.(Comparable); ok {
		return comparable.IsEqual(b)
	}
	return a == b
}

func NewSlice() Slice {
	return make(Slice, 0)
}

func (this *Slice) Add(elem interface{}) error {
	for _, v := range *this {
		if isEqual(v, elem) {
			klog.Errorf("Slice:Add elem: %v already exist", elem)
			return ErrElemExist
		}
	}
	*this = append(*this, elem)
	klog.Errorf("Slice:Add elem: %v succ", elem)
	return nil
}

func (this *Slice) Remove(elem interface{}) error {
	flag := true
	for i, v := range *this {
		if isEqual(v, elem) {
			if i+2 > len(*this) {
				*this = (*this)[:i]

			} else {
				*this = append((*this)[:i], (*this)[i+1:]...)
			}
			flag = false
			break
		}
	}
	if flag {
		klog.Errorf("Slice:Remove elem: %v not exist", elem)
		return ErrElemNtExist
	}
	klog.Errorf("Slice:Remove elem: %v succ", elem)
	return nil
}
