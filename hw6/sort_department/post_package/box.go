package post_package

import "fmt"

type Box struct {
	receiverAddress string
	senderAddress   string
	packageType     string
	transport       string
}

func NewBox(senderAddress string, receiverAddress string) *Box {
	return &Box{
		receiverAddress: receiverAddress,
		senderAddress:   senderAddress,
		packageType:     BoxType,
		transport:       "train",
	}
}

func (b Box) ReceiverAddress() string {
	return b.receiverAddress
}

func (b Box) SenderAddress() string {
	return b.senderAddress
}

func (b Box) GetType() string {
	return b.packageType
}

func (b Box) Send() {
	fmt.Printf("Send %s from %s to %s by %s\n", b.GetType(), b.SenderAddress(), b.ReceiverAddress(), b.transport)
}
