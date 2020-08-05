package tools


import uuid "github.com/satori/go.uuid"


// 得到一个uuid
func GetUuid() string {
	u1 := uuid.Must(uuid.NewV4(),nil)
	return u1.String()
}

func CheckUuid(u1 string) bool {
	_, err := uuid.FromString(u1)
	if err != nil {
		return false
	}
	return true
}

