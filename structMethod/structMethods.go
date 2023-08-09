package golangbootdev

import "fmt"

// "golang.org/x/text/message"

func main() {

}

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func testing(m authenticationInfo) {
	fmt.Printf("Sending mesage: '%s' to: %v\n", m.username, m.password)
	fmt.Println("===============================")
}
