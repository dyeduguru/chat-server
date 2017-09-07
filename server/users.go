package server

type UserStore struct {
	users map[string]struct{}
}

func (u *UserStore) List() []string {
	users := []string{}
	for k := range u.users {
		users = append(users, k)
	}
	return users
}

func (u *UserStore) Add(user string) {
	u.users[user] = struct{}{}
}
