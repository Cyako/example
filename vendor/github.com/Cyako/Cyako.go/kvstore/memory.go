// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvstore

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Memory struct {
	mutex sync.RWMutex
	m     map[string]interface{}
}

func (s *Memory) Init() {
	s.m = make(map[string]interface{})
}

func (s *Memory) init() {
	if s.m == nil {
		s.Init()
	}
}

func (s *Memory) Get(key string) interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.m[key]
}

func (s *Memory) Set(key string, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = value
}

func (s *Memory) Has(key string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	_, ok := s.m[key]
	return ok
}

func (s *Memory) Delete(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.m, key)
}

func (s *Memory) Disactive() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	content, _ := json.Marshal(s.m)
	ioutil.WriteFile("KVStore.Memory.Dump", content, 0777)
}

func (s *Memory) Active() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
}
