package izface

type IRequest interface {


	GetConnection() IConnection

	GetData() []byte

}