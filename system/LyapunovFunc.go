package system

import "gonum.org/v1/gonum/mat"

type Lyapunov struct {
	stateVec        mat.MutableVector
	unknownParamVec mat.MutableVector
	stateExpireVec  mat.MutableVector
}
