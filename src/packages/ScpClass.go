package packages

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	AccountFile string = "./account.json"
)

type ScpInterface interface {
	GetUsername() string
	GetHost() string
	GetPort() int
	GetPassword() string
}

type Host struct {
	Username string
	Host     string
	Port     int
	Password string
}
type User struct {
	Name string
	File []string
}

func (host *Host) SetUsername(Username string) {
	host.Username = Username
}
func (host *Host) SetHost(Host string) {
	host.Host = Host
}
func (host *Host) SetPort(Port int) {
	host.Port = Port
}
func (host *Host) SetPassword(Password string) {
	host.Password = Password
}
func (host *Host) GetUsername() string {
	return host.Username
}
func (host *Host) GetHost() (Host string) {
	Host = host.Host
	return
}
func (host *Host) GetPort() (Port int) {
	Port = host.Port
	return
}
func (host *Host) GetPassword() (Password string) {
	Password = host.Password
	return
}
func (host *Host) CreateAccount(Username string, Host string, Port int, Password string) {
	host.SetUsername(Username)
	host.SetHost(Host)
	host.SetPort(Port)
	host.SetPassword(Password)
}
func (host *Host) InsertAccount() {
	data, _ := json.Marshal(host)
	fmt.Printf("%s", string(data))
	WriteStringToFile(string(data), AccountFile)
}

func (user *User) CreateFile() {
	user.File = GetFiles("conveyFile")
}
func (user *User) ReadConveyFile() {
	for i, v := range user.File {
		fmt.Printf("[%v : %v]\n", i, v)
	}
}
func (user *User) ConveyFile(scp ScpInterface) {
	command := "scp -P " + strconv.Itoa(scp.GetPort()) + " -r ./conveyFile " + scp.GetUsername() + "@" + scp.GetHost() + ":~/"
	fmt.Printf("%s\n", command)
	ExecCommandBinBash(command)
}

func CreateHost() Host {
	var Username, Hostip, Password string
	var Port int
	fmt.Printf("Username:")
	fmt.Scanf("%s", &Username)
	fmt.Printf("Host:")
	fmt.Scanf("%s", &Hostip)
	fmt.Printf("Port:")
	fmt.Scanf("%d", &Port)
	fmt.Printf("Password:")
	fmt.Scanf("%s", &Password)
	host := Host{
		Username: Username,
		Host:     Hostip,
		Port:     Port,
		Password: Password,
	}
	return host
}
