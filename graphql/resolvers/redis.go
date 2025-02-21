package graphqlresolvers

import "encoding/json"

type RedisUser struct {
	Name  string `json:"name" redis:"name"`
	Email string `json:"email" redis:"email"`
}

type RedisUserList []struct {
	Name  string `json:"name" redis:"name"`
	Email string `json:"email" redis:"email"`
}

func (s *RedisUserList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s RedisUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
func (s *RedisUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
