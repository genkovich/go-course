package post_package

type Box struct {
	receiverAddress string
	senderAddress   string
	packageType     string
}

func NewBox(senderAddress string, receiverAddress string) *Box {
	return &Box{
		receiverAddress: receiverAddress,
		senderAddress:   senderAddress,
		packageType:     BoxType,
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
