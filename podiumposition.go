package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)

	for i := 0; i < length/2; i++ {
		temp := podium[i]
		podium[i] = podium[length-1-i]
		podium[length-1-i] = temp
	}
	return podium
}
