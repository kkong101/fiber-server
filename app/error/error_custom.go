package error

type Format struct {
	code    string
	message string
}

var (
	alreadyRegisteredUser = Format{"P001", "이미 가입된 회원입니다."}
	deletedUser           = Format{"P002", "삭제된 회원입니다."}
)
