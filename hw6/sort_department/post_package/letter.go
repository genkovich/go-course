package post_package

type Letter struct {
	receiverAddress string
	senderAddress   string
	packageType     string
}

func NewLetter(senderAddress string, receiverAddress string) *Letter {
	return &Letter{
		receiverAddress: receiverAddress,
		senderAddress:   senderAddress,
		packageType:     LetterType,
	}
}

func (l Letter) ReceiverAddress() string {
	return l.receiverAddress
}

func (l Letter) SenderAddress() string {
	return l.senderAddress
}

func (l Letter) GetType() string {
	return l.packageType
}
