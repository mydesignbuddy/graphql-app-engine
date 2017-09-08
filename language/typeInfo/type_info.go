package typeInfo

import (
	"github.com/mydesignbuddy/graphql-app-engine/language/ast"
)

// TypeInfoI defines the interface for TypeInfo Implementation
type TypeInfoI interface {
	Enter(node ast.Node)
	Leave(node ast.Node)
}
