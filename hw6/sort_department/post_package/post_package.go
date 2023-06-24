package post_package

const (
	LetterType = "letter"
	BoxType    = "box"
)

type PostPackage interface {
	ReceiverAddress() string
	SenderAddress() string
	GetType() string
	Send()
}
