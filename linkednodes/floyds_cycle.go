package linkednodes

func FloydsCycle(head *LinkedNode) bool {
	slowptr, fastptr := head, head
	for slowptr!=nil && fastptr!=nil && fastptr.Next!=nil{
		slowptr = slowptr.Next
		fastptr = fastptr.Next.Next
		if slowptr == fastptr {
			return true
		}
	}

	return false
}



