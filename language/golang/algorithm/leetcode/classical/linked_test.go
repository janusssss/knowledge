package classical

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2,
		&ListNode{4,
			&ListNode{3, nil}},
	}
	l2 := &ListNode{5,
		&ListNode{6,
			&ListNode{4, nil}},
	}
	result := addTwoNumbers(l1, l2)
	expected := &ListNode{7,
		&ListNode{0,
			&ListNode{8, nil}}}
	if cmp.Diff(result, expected) != "" {
		t.Errorf("add two numbers failed, expected %v, got %v", expected, result)
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		result *ListNode
	}{{
		l1: &ListNode{1,
			&ListNode{2,
				&ListNode{4, nil}}},
		l2: &ListNode{1,
			&ListNode{3,
				&ListNode{4, nil}}},
		result: &ListNode{1,
			&ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							&ListNode{4, nil}}}}}}},
		{l1: nil, l2: &ListNode{}, result: &ListNode{}},
		{l1: &ListNode{Val: 2}, l2: &ListNode{Val: 1}, result: &ListNode{1, &ListNode{2, nil}}},
	}
	for _, test := range tests {
		result := mergeTwoLists(test.l1, test.l2)
		if fact := cmp.Diff(test.result, result); fact != "" {
			t.Errorf(fact)
		}
	}

}

//func TestCopyRandomList(t *testing.T) {
//	n0 := &Node{}
//	n1 := &Node{}
//	n2 := &Node{}
//	n3 := &Node{}
//	n4 := &Node{}
//	n0.Val, n0.Next, n0.Random = 7, n1, nil
//	n1.Val, n1.Next, n1.Random = 13, n2, n0
//	n2.Val, n2.Next, n2.Random = 14, n3, n4
//	n3.Val, n3.Next, n3.Random = 10, n4, n2
//	n4.Val, n4.Next, n4.Random = 1, nil, n0
//	copyHead := copyRandomList(n0)
//	if fact := cmp.Diff(n0, copyHead); fact != "" {
//		t.Errorf(fact)
//	}
//}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head   *ListNode
		left   int
		right  int
		result *ListNode
	}{
		{&ListNode{5, nil}, 1, 1, &ListNode{5, nil}},

		{&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}},
			2, 4,
			&ListNode{1,
				&ListNode{4,
					&ListNode{3,
						&ListNode{2,
							&ListNode{5, nil}}}}}},

		{&ListNode{3, &ListNode{5, nil}},
			1, 1,
			&ListNode{3, &ListNode{5, nil}}},

		{&ListNode{3, &ListNode{5, nil}},
			1, 2,
			&ListNode{5, &ListNode{3, nil}}},

		{&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			2, 3,
			&ListNode{1, &ListNode{3, &ListNode{2, nil}}}},

		{&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			3, 3,
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
	}
	for _, test := range tests {
		result := reverseBetween(test.head, test.left, test.right)
		if fact := cmp.Diff(result, test.result); fact != "" {
			t.Errorf(fact)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		head   *ListNode
		k      int
		result *ListNode
	}{
		{head: &ListNode{1, &ListNode{2,
			&ListNode{3, &ListNode{4,
				&ListNode{5, nil}}}}},
			k: 2,
			result: &ListNode{2, &ListNode{1,
				&ListNode{4, &ListNode{3,
					&ListNode{5, nil}}}}},
		},

		{head: &ListNode{1, &ListNode{2, &ListNode{3,
			&ListNode{4, &ListNode{5, nil}}}}},
			k: 3,
			result: &ListNode{3, &ListNode{2, &ListNode{1,
				&ListNode{4, &ListNode{5, nil}}}}},
		},
	}
	for _, test := range tests {
		result := reverseKGroup(test.head, test.k)
		if fact := cmp.Diff(result, test.result); fact != "" {
			t.Errorf(fact)
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head   *ListNode
		result *ListNode
	}{
		{head: &ListNode{1, &ListNode{1, &ListNode{1,
			&ListNode{2, &ListNode{3, nil}}}}},
			result: &ListNode{2, &ListNode{3, nil}},
		},

		{head: &ListNode{1, &ListNode{1, nil}},
			result: nil,
		},

		{nil, nil},
	}
	for _, test := range tests {
		result := deleteDuplicates(test.head)
		if fact := cmp.Diff(result, test.result); fact != "" {
			t.Errorf(fact)
		}
	}
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		head   *ListNode
		k      int
		result *ListNode
	}{
		{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}},
			k: 2,
			result: &ListNode{4,
				&ListNode{5,
					&ListNode{1,
						&ListNode{2,
							&ListNode{3, nil}}}}},
		},

		{nil, 0, nil},

		{&ListNode{1, nil}, 0, &ListNode{1, nil}},

		{&ListNode{1, &ListNode{2, nil}}, 0,
			&ListNode{1, &ListNode{2, nil}}},
	}
	for _, test := range tests {
		result := rotateRight(test.head, test.k)
		if fact := cmp.Diff(result, test.result); fact != "" {
			t.Errorf(fact)
		}
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		head   *ListNode
		x      int
		result *ListNode
	}{
		{&ListNode{1,
			&ListNode{4,
				&ListNode{3,
					&ListNode{2,
						&ListNode{5,
							&ListNode{2, nil}}}}}},
			3,
			&ListNode{1,
				&ListNode{2,
					&ListNode{2,
						&ListNode{4,
							&ListNode{3,
								&ListNode{5, nil}}}}}}},
	}
	for _, test := range tests {
		result := partition(test.head, test.x)
		if fact := cmp.Diff(result, test.result); fact != "" {
			for ; result != nil; result = result.Next {
				print(result.Val, ", ")
			}
		}
	}
}

func TestLRUCache(t *testing.T) {
	c := ConstructorLinked(2)
	print(c.Get(2))
	c.Put(2, 6)
	print(c.Get(1))
	c.Put(1, 5)
	c.Put(1, 2)
	print(c.Get(1))
	print(c.Get(2))
	println()

	c = ConstructorLinked(2)
	c.Put(2, 1)
	c.Put(2, 2)
	print(c.Get(2))
	c.Put(1, 1)
	c.Put(4, 1)
	print(c.Get(2))
	print(c.Get(4))
	println()

	c = ConstructorLinked(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	print(c.Get(4))
	print(c.Get(3))
	print(c.Get(2))
	print(c.Get(1))
	c.Put(5, 5)
	print(c.Get(1))
	print(c.Get(2))
	print(c.Get(3))
	print(c.Get(4))
	print(c.Get(5))
	println()
}
