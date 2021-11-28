package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// 	l1 := &problems.ListNode{
//		Val: -9,
//		Next: &problems.ListNode{
//			Val: -7,
//			Next: &problems.ListNode{
//				Val: -3,
//				Next: &problems.ListNode{
//					Val: -3,
//					Next: &problems.ListNode{
//						Val: -1,
//						Next: &problems.ListNode{
//							Val: 2,
//							Next: &problems.ListNode{
//								Val: 3,
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	l2 := &problems.ListNode{
//		Val: -7,
//		Next: &problems.ListNode{
//			Val: -7,
//			Next: &problems.ListNode{
//				Val: -6,
//				Next: &problems.ListNode{
//					Val: -6,
//					Next: &problems.ListNode{
//						Val: -5,
//						Next: &problems.ListNode{
//							Val: -3,
//							Next: &problems.ListNode{
//								Val: 2,
//								Next: &problems.ListNode{
//									Val: 4,
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2

			return res.Next
		}

		if l2 == nil {
			cur.Next = l1

			return res.Next
		}

		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		}
	}

	return res.Next
}
