package sort_department

import (
	postPackage "course/hw6/sort_department/post_package"
	"fmt"
)

type SortDepartment struct {
	postPackages []postPackage.PostPackage
}

func (s SortDepartment) SendAll() {
	for _, post := range s.postPackages {
		transport := postPackage.PickTransport(post)
		fmt.Printf("Send %s from %s to %s by %s\n", post.GetType(), post.SenderAddress(), post.ReceiverAddress(), transport)
	}
}

func Start() {
	firstLetter := postPackage.NewLetter("Odesa", "Kyiv")
	secondLetter := postPackage.NewLetter("Kharkiv", "Lviv")
	Box := postPackage.NewBox("Dnipro", "Mykolaiv")

	sortDep := SortDepartment{[]postPackage.PostPackage{
		firstLetter,
		secondLetter,
		Box,
	}}

	sortDep.SendAll()
}
