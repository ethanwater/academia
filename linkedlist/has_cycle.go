package linkedlist


func HasCycle(head *ListNode) bool {
  if head == nil || head.Next == nil{
    return false
  }
  slowptr, fastptr := head, head
  for slowptr != fastptr{
    if fastptr == nil || fastptr.Next == nil {
      return false
    }
    fastptr = fastptr.Next.Next
    slowptr = slowptr.Next
  }
  return true 
}
