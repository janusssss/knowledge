package classical

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	examples := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{lists: []*ListNode{
			{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
			{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			{Val: 2, Next: &ListNode{Val: 6}},
		},
			want: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 4, Next: &ListNode{
										Val: 5}}}}}},
			},
		},
	}
	printNode(examples[0].lists[0])
	printNode(examples[0].lists[1])
	printNode(examples[0].lists[2])

	var check func(a, b *ListNode) bool
	check = func(a, b *ListNode) bool {
		if a == b {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}

		return check(a.Next, b.Next)
	}

	for _, e := range examples {
		result := MergeKLists(e.lists)
		if !check(result, e.want) {
			t.Errorf("%+v", e)
			printNode(result)
		}
	}

}

func TestSortList(t *testing.T) {
	examples := []struct {
		head   *ListNode
		answer *ListNode
	}{
		{
			&ListNode{3, &ListNode{2, &ListNode{1, nil}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			&ListNode{3, &ListNode{3, &ListNode{1, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{3, nil}}},
		},

		{
			&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},
	}

	fn := func(head *ListNode) []int {
		result := make([]int, 0)
		for ; head != nil; head = head.Next {
			result = append(result, head.Val)
		}

		return result
	}

	for _, e := range examples {
		result, answer := fn(sortList(e.head)), fn(e.answer)
		if !reflect.DeepEqual(result, answer) {
			t.Fatalf("result:%v,answer,%v", result, answer)
		}
	}
}
