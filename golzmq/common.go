package golzmq

type RecieveArrayReturnString func(msg_array []string) string
type ReceiveStringReturnNil func(msg string)

type RecieveByteArrayReturnByte func(msg []byte) []byte
