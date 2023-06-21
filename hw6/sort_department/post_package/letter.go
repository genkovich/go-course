package post_package

import "fmt"

type Letter struct {
	receiverAddress string
	senderAddress   string
	packageType     string
	transport       string
}

func NewLetter(senderAddress string, receiverAddress string) *Letter {
	return &Letter{
		receiverAddress: receiverAddress,
		senderAddress:   senderAddress,
		packageType:     LetterType,
		transport:       "plane",
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

func (l Letter) Send() {
	fmt.Printf("Send %s from %s to %s by %s\n", l.GetType(), l.SenderAddress(), l.ReceiverAddress(), l.transport)
}
