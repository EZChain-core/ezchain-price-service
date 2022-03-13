package utils

func remove(list []string, index int) []string {
	list[index] = list[len(list)-1]
	return list[:len(list)-1]
}