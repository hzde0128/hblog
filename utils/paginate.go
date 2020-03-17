package utils

// prepage 上一页
func Prepage(idx int) (page int) {
	if idx > 1 {
		page = idx - 1
	} else {
		page = idx
	}
	return
}

// nextpage 下一页
func Nextpage(idx, count int) (page int) {
	if idx < count {
		page = idx + 1
	} else {
		page = count
	}
	return
}
