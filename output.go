// Package solc
//
// @author: xwc1125
package solc

import (
	"encoding/json"
)

// Output 编译输出
type Output struct {
	Errors    []Error                        `json:"errors,omitempty"`    // 错误集合
	Sources   map[string]SourceOut           `json:"sources,omitempty"`   // 输出集合
	Contracts map[string]map[string]Contract `json:"contracts,omitempty"` // 合约内容。xxx.sol-->ContractName-->Contract
}

// Error 错误内容
type Error struct {
	SourceLocation           SourceLocation `json:"sourceLocation,omitempty"`           // 资源的位置
	SecondarySourceLocations SourceLocation `json:"secondarySourceLocations,omitempty"` // 其他资源的位置
	Type                     string         `json:"type,omitempty"`                     // 错误类型。例如："TypeError", "InternalCompilerError", "Exception"
	Component                string         `json:"component,omitempty"`                // 产生错误的组件。例如："general", "ewasm"
	Severity                 string         `json:"severity,omitempty"`                 // 错误级别。例如："error", "warning" or "info"
	ErrorCode                string         `json:"errorCode,omitempty"`                // 错误code
	Message                  string         `json:"message,omitempty"`                  // 错误消息
	FormattedMessage         string         `json:"formattedMessage,omitempty"`         // 错误消息的位置
}

// SourceLocation 资源位置
type SourceLocation struct {
	File    string `json:"file,omitempty"`
	Start   int    `json:"start,omitempty"`
	End     int    `json:"end,omitempty"`
	Message string `json:"message,omitempty"`
}

// SourceOut 资源输出
type SourceOut struct {
	ID        int             `json:"id,omitempty"`
	AST       json.RawMessage `json:"ast,omitempty"`
	LegacyAST json.RawMessage `json:"legacyAST,omitempty"`
}

// Contract 合约内容
type Contract struct {
	ABI      []json.RawMessage `json:"abi,omitempty"` // abi
	Metadata string            `json:"metadata,omitempty"`
	UserDoc  json.RawMessage   `json:"userdoc,omitempty"`
	DevDoc   json.RawMessage   `json:"devdoc,omitempty"`
	IR       string            `json:"ir,omitempty"`
	//StorageLayout StorageLayout     `json:"storageLayout,omitempty"`
	EVM   EVM   `json:"evm,omitempty"`
	EWASM EWASM `json:"ewasm,omitempty"`
}

// EVM evm
type EVM struct {
	Assembly          string                       `json:"assembly,omitempty"`
	LegacyAssembly    json.RawMessage              `json:"legacyAssembly,omitempty"`
	Bytecode          Bytecode                     `json:"bytecode,omitempty"`
	DeployedBytecode  Bytecode                     `json:"deployedBytecode,omitempty"`
	MethodIdentifiers map[string]string            `json:"methodIdentifiers,omitempty"`
	GasEstimates      map[string]map[string]string `json:"gasEstimates,omitempty"`
}

// Bytecode 字节
type Bytecode struct {
	Object         string                                `json:"object,omitempty"`
	Opcodes        string                                `json:"opcodes,omitempty"`
	SourceMap      string                                `json:"sourceMap,omitempty"`
	LinkReferences map[string]map[string][]LinkReference `json:"linkReferences,omitempty"`
}

// LinkReference 关联
type LinkReference struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// EWASM ewasm
type EWASM struct {
	Wast string `json:"wast,omitempty"`
	Wasm string `json:"wasm,omitempty"`
}
