package main

func getIndexValue(slice []Lobby, index int) (Lobby, bool) {
	if index < 0 || index >= len(slice) {
		return Lobby{}, false
	}
	return slice[index], true
}
