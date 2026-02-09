package utils

import "encoding/json"

type Struct_Hosts struct {
	Servers []Struct_Server `json:"postgresql"`
}

type Struct_Server struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Auth     string `json:"auth"`
	CLI      string `json:"cli"`
	AWSIAM   struct {
		Region  string `json:"region"`
		Profile string `json:"profile"`
	} `json:"aws-iam,omitempty"`
}

func (s *Struct_Server) UnmarshalJSON(data []byte) error {
	type obj_Server_temp Struct_Server

	aux_temp := &obj_Server_temp{
		Auth: "password",
		CLI:  "pgcli",
		Port: 5432,
	}

	if err := json.Unmarshal(data, aux_temp); err != nil {
		return err
	}

	*s = Struct_Server(*aux_temp)
	return nil
}
