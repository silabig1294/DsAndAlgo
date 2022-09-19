package reversing

func ReverseInt(items []int) []int{
    var maxdata = len(items) -1
	var i int = 0
	for i=0;i<maxdata/2;i++{
		// var temp int
		temp := items[i]
		items[i] = items[maxdata-1-i]
		items[maxdata-1-i]=temp

	}

	return items
}