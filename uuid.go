// Copyright (c) 2014 - Max Persson <max@looplab.se>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventhorizon

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
)

// UUID is a unique identifier, based on the UUID spec.
type UUID uuid.UUID

// NewUUID creates a new UUID of type v4.
func NewUUID() UUID {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return UUID(*guid)
}

// ParseUUID parses a UUID from a string representation.
func ParseUUID(value string) (id UUID, err error) {
	guid := new(uuid.UUID)
	if guid, err = uuid.ParseHex(value); err == nil {
		id = UUID(*guid)
	}

	return
}

// String implements the Stringer interface for UUID.
func (id UUID) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}

// MarshalJSON turns UUID into a json.Marshaller.
func (id *UUID) MarshalJSON() ([]byte, error) {
	// Pack the string representation in quotes
	return []byte(fmt.Sprintf(`"%v"`, id.String())), nil
}

// UnmarshalJSON turns *UUID into a json.Unmarshaller.
func (id *UUID) UnmarshalJSON(data []byte) error {
	// Data is expected to be a json string, like: "819c4ff4-31b4-4519-5d24-3c4a129b8649"
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid UUID in JSON, %v is not a valid JSON string", string(data))
	}

	// Grab string value without the surrounding " characters
	value := string(data[1 : len(data)-1])
	parsed, err := uuid.ParseHex(value)
	if err != nil {
		return fmt.Errorf("invalid UUID in JSON, %v: %v", value, err)
	}

	// Dereference pointer value and store parsed
	*id = UUID(*parsed)
	return nil
}
