package linkednodes

func FloydsCycle(head *LinkedNode) bool {
	slowptr, fastptr := head, head
	for slowptr != nil && fastptr != nil && fastptr.Next != nil {
		slowptr = slowptr.Next
		fastptr = fastptr.Next.Next
		if slowptr == fastptr {
			return true
		}
	}

	return false
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slowptr, fastptr := head, head
	for slowptr != fastptr {
		if fastptr == nil || fastptr.Next == nil {
			return false
		}
		fastptr = fastptr.Next.Next
		slowptr = slowptr.Next
	}
	return true
}
