package user

type user struct{}

var User = new(user)

func (user) Creator() Creator {
	return Creator{}
}
