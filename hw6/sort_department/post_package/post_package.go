package post_package

const (
	LetterType = "letter"
	BoxType    = "box"
)

type PostPackage interface {
	ReceiverAddress() string
	SenderAddress() string
	GetType() string
}

func PickTransport(post PostPackage) string {
	switch post.GetType() {
	case LetterType:
		return "plane"
	case BoxType:
		return "Train"
	default:
		return "Truck"
	}
}
