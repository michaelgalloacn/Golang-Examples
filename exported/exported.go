package exported

// This constant can not be directly accessed outside of this package based on its first letter being lowercased
const privateConst = "private"
const PublicConst = "public"

type PublicType struct {
	privateMember int
	PublicMember  int
}

type privateType struct {
	privateMember int
	PublicMember  int
}

// Even though PublicType is public,
var privateVar = PublicType{
	privateMember: 0,
	PublicMember:  5,
}

func privateFunction() string {
	return "private"
}

func PublicWrapperForPrivatefunction() string {
	return privateFunction()
}

func PublicFunction() string {
	return "public"
}
