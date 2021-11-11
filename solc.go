// Package solc
//
// @author: xwc1125
package solc

import (
	"encoding/json"
	"io/ioutil"
	"rogchap.com/v8go"
	"strings"
	"sync"
)

// Solc solc的接口
type Solc interface {
	License() string
	Version() string
	Compile(input *Input) (*Output, error)
	Close()
}

type baseSolc struct {
	isolate *v8go.Isolate
	ctx     *v8go.Context

	mux *sync.Mutex

	version *v8go.Value
	license *v8go.Value
	compile *v8go.Value
}

// NewFromFile 通过读取sol文件进行初始化
func NewFromFile(solFile string) (Solc, error) {
	solData, err := ioutil.ReadFile(solFile)
	if err != nil {
		return nil, err
	}

	return New(string(solData))
}

// New 通过sol文本内容进行初始化
func New(solData string) (Solc, error) {
	return newSolc(solData)
}

// newSolc 创建solc对象
func newSolc(solData string) (*baseSolc, error) {
	ctx, err := v8go.NewContext()
	iso, _ := ctx.Isolate()
	solc := &baseSolc{
		mux:     new(sync.Mutex),
		isolate: iso,
		ctx:     ctx,
	}

	// 初始化solc
	err = solc.init(solData)
	if err != nil {
		return nil, err
	}

	return solc, nil
}

// init 初始化solc
func (s *baseSolc) init(solData string) error {
	// 执行solcjson.js
	_, err := s.ctx.RunScript(solData, "soljson.js")
	if err != nil {
		return err
	}

	// 绑定sol中version部分
	if strings.Contains(solData, "_solidity_version") {
		s.version, err = s.ctx.RunScript("Module.cwrap('solidity_version', 'string', [])", "wrap_version.js")
		if err != nil {
			return err
		}
	} else {
		s.version, err = s.ctx.RunScript("Module.cwrap('version', 'string', [])", "wrap_version.js")
		if err != nil {
			return err
		}
	}

	// 绑定sol中license部分
	if strings.Contains(solData, "_solidity_license") {
		s.license, err = s.ctx.RunScript("Module.cwrap('solidity_license', 'string', [])", "wrap_license.js")
		if err != nil {
			return err
		}
	} else if strings.Contains(solData, "_license") {
		s.license, err = s.ctx.RunScript("Module.cwrap('license', 'string', [])", "wrap_license.js")
		if err != nil {
			return err
		}
	}

	// 绑定sol编译部分
	s.compile, err = s.ctx.RunScript("Module.cwrap('solidity_compile', 'string', ['string', 'number', 'number'])", "wrap_compile.js")
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭
func (s *baseSolc) Close() {
	s.mux.Lock()
	defer s.mux.Lock()
	s.ctx.Close()
	s.isolate.Dispose()
}

// License 获取license
func (s *baseSolc) License() string {
	if s.license != nil {
		s.mux.Lock()
		defer s.mux.Lock()
		fn, err := s.license.AsFunction()
		if err != nil {
			return ""
		}
		val, _ := fn.Call()
		return val.String()
	}
	return ""
}

// Version 获取version
func (s *baseSolc) Version() string {
	if s.version != nil {
		s.mux.Lock()
		defer s.mux.Lock()
		fn, err := s.version.AsFunction()
		if err != nil {
			return ""
		}
		val, _ := fn.Call()
		return val.String()
	}
	return ""
}

// Compile 编译
func (s *baseSolc) Compile(input *Input) (*Output, error) {
	// 将input数据进行序列化
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	// 执行编译
	s.mux.Lock()
	defer s.mux.Unlock()
	inputVal, err := v8go.NewValue(s.isolate, string(b))
	if err != nil {
		return nil, err
	}
	oneVal, err := v8go.NewValue(s.isolate, uint32(1))
	if err != nil {
		return nil, err
	}
	fn, err := s.compile.AsFunction()
	if err != nil {
		return nil, err
	}
	// 调用方法，然后传入参数，获取执行结果
	outputVal, err := fn.Call(inputVal, oneVal, oneVal)
	if err != nil {
		return nil, err
	}

	// 输出output对象
	out := &Output{}
	err = json.Unmarshal([]byte(outputVal.String()), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
