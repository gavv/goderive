//  Copyright 2017 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package keys contains the implementation of the keys plugin, which generates the deriveKeys function.
// The deriveKeys function returns a map's keys as a slice.
package keys

import (
	"fmt"
	"go/types"

	"github.com/awalterschulze/goderive/derive"
)

// NewPlugin creates a new keys plugin.
// This function returns the plugin name, default prefix and a constructor for the keys code generator.
func NewPlugin() derive.Plugin {
	return derive.NewPlugin("keys", "deriveKeys", New)
}

// New is a constructor for the keys code generator.
// This generator should be reconstructed for each package.
func New(typesMap derive.TypesMap, p derive.Printer, deps map[string]derive.Dependency) derive.Generator {
	return &gen{
		TypesMap: typesMap,
		printer:  p,
	}
}

type gen struct {
	derive.TypesMap
	printer derive.Printer
}

func (this *gen) Add(name string, typs []types.Type) (string, error) {
	if len(typs) != 1 {
		return "", fmt.Errorf("%s does not have one argument", name)
	}
	return this.SetFuncName(name, typs[0])
}

func (this *gen) Generate() error {
	for _, typs := range this.ToGenerate() {
		typ := typs[0]
		mapType, ok := typ.(*types.Map)
		if !ok {
			return fmt.Errorf("%s, the first argument, %s, is not of type map", this.GetFuncName(typ), typ)
		}
		if err := this.genFuncFor(mapType); err != nil {
			return err
		}
	}
	return nil
}

func (this *gen) genFuncFor(typ *types.Map) error {
	p := this.printer
	this.Generating(typ)
	typeStr := this.TypeString(typ)
	keyType := typ.Key()
	keyTypeStr := this.TypeString(keyType)
	p.P("")
	p.P("func %s(m %s) []%s {", this.GetFuncName(typ), typeStr, keyTypeStr)
	p.In()
	p.P("keys := make([]%s, 0, len(m))", keyTypeStr)
	p.P("for key, _ := range m {")
	p.In()
	p.P("keys = append(keys, key)")
	p.Out()
	p.P("}")
	p.P("return keys")
	p.Out()
	p.P("}")
	return nil
}