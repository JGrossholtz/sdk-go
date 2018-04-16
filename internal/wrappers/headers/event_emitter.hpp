// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef _EVENT_EMITTER_HPP_
#define _EVENT_EMITTER_HPP_

#include "kuzzle.hpp"
#include "listeners.hpp"

namespace kuzzleio {
  class KuzzleEventEmitter {
    public:
      virtual KuzzleEventEmitter* addListener(Event e, EventListener* listener) = 0;
      virtual KuzzleEventEmitter* removeListener(Event e, EventListener* listener) = 0;
      virtual KuzzleEventEmitter* removeAllListeners(Event e) = 0;
      virtual KuzzleEventEmitter* once(Event e, EventListener* listener) = 0;
      virtual int listenerCount(Event e) = 0;
  };
}

#endif