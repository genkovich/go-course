package sort_department

import (
	postPackage "course/hw6/sort_department/post_package"
)

type SortDepartment struct {
	postPackages []postPackage.PostPackage
}

func (s SortDepartment) SendAll() {
	for _, post := range s.postPackages {
		post.Send()
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
