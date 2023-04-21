package node


type INode interface {
  GetKind() NodeKind

  GetPosition() int
  GetLength() int
}
