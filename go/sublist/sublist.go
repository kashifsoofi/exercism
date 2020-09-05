// Package sublist implements routine to determine relation between 2 lists
package sublist

// Relation between 2 lists
type Relation string

// Sublist returns relation of first list with second
func Sublist(l1, l2 []int) Relation {
	switch {
	case len(l1) == len(l2) && isSublist(l1, l2):
		return "equal"
	case len(l1) < len(l2) && isSublist(l1, l2):
		return "sublist"
	case len(l2) < len(l1) && isSublist(l2, l1):
		return "superlist"
	}
	return "unequal"
}

func isSublist(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}

	i := indexOf(l2, l1[0], 0)
	if i == -1 || len(l1) > len(l2)-i {
		return false
	}

	for j := i; j < len(l1); j++ {
		if l1[j] != l2[i+j] {
			i = indexOf(l2, l1[0], i+1)
			if i == -1 || len(l1) > len(l2)-i {
				return false
			}
			j = -1
		}
	}

	return true
}

func indexOf(l []int, item, startIndex int) int {
	for i := startIndex; i < len(l); i++ {
		if l[i] == item {
			return i
		}
	}
	return -1
}
