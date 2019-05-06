// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/buptbill220/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*
The size plugin generates a Size or ProtoSize method for each message.
This is useful with the MarshalTo method generated by the marshalto plugin and the
gogoproto.marshaler and gogoproto.marshaler_all extensions.

It is enabled by the following extensions:

  - sizer
  - sizer_all
  - protosizer
  - protosizer_all

The size plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

And a benchmark given it is enabled using one of the following extensions:

  - benchgen
  - benchgen_all

Let us look at:

  github.com/buptbill220/protobuf/test/example/example.proto
*/
package jsonmarshal

import (
	"fmt"

	"github.com/buptbill220/protobuf/protoc-gen-gogo/generator"
)

type JsonMarshal struct {
	*generator.Generator
	generator.PluginImports
	atleastOne bool
	localName  string
	typesPkg   generator.Single
}

func NewJsonMarshal() *JsonMarshal {
	return &JsonMarshal{}
}

func (p *JsonMarshal) Name() string {
	return "jsonmarshal"
}

func (p *JsonMarshal) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *JsonMarshal) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	fmtPkg := p.PluginImports.NewImport("fmt").Use()
	jsonPkg := p.PluginImports.NewImport("encoding/json").Use()
	p.atleastOne = false
	p.localName = generator.FileName(file)
	for _, message := range file.Messages() {
		p.atleastOne = true
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (m *`, ccTypeName, `) Marshal() ([]byte, error) {`)
		p.In()
		p.P(`if m == nil {`)
		p.In()
		p.P(fmt.Sprintf(`return nil, ` + fmtPkg + `.Errorf("msg %s is nil")`, ccTypeName))
		p.Out()
		p.P(`}`)
		
		p.P(`return ` + jsonPkg + `.Marshal(m)`)
		p.Out()
		p.P(`}`)
		p.P()
		
		p.P(`func (m *`, ccTypeName, `) Unmarshal(data []byte) error {`)
		p.In()
		p.P(`if m == nil {`)
		p.In()
		p.P(fmt.Sprintf(`return ` + fmtPkg + `.Errorf("msg %s is nil")`, ccTypeName))
		p.Out()
		p.P(`}`)
		
		p.P(`return ` + jsonPkg + `.Unmarshal(data, m)`)
		p.Out()
		p.P(`}`)
		p.P()
	}
}

func init() {
	generator.RegisterPlugin(NewJsonMarshal())
}
