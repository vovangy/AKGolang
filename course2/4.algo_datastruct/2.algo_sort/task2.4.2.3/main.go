package main

type User struct {
	ID   int
	Name string
	Age  int
}

func Merge(arr1 []User, arr2 []User) []User {
	merged := make([]User, 0, len(arr1)+len(arr2))

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i].ID <= arr2[j].ID {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}
