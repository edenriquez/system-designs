package pkg

const (
	serviceURL = "http://localhost:8080/"	
	fakeUrl = 125
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

)

func encode(url string) []int{
	remainder := 125 // this should bne url param
	var remainderSlice []int
	for remainder > 0 {
		remainderSlice = append(remainderSlice, (remainder % 62))
		remainder = remainder / 62
	}
	return remainderSlice
}
// URLShortener should return given url shortened
func URLShortener(url string) string {
	// encode
	encode(url)
	return serviceURL + "123abc"
}
