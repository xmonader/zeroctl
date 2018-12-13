package cmd

type AppArguments struct {
	Instance   string
	Port       int
	Address    string // ip
	Cmd        string
	JsonArgs   string
	Root       string
	SetDefault bool
	Sshkey     string
	Name       string
	Hostname   string
	Ports      string
	Env        string
	Json       bool
}

var CommandArgs *AppArguments

func init() {
	CommandArgs := new(AppArguments)
	configureCmd.Flags().StringVarP(&CommandArgs.Address, "address", "a", "", "address of zero-os")
	configureCmd.Flags().StringVarP(&CommandArgs.Instance, "name", "n", "", "instance name")
	configureCmd.Flags().IntVarP(&CommandArgs.Port, "port", "p", 6379, "redis port of zero-os")

	cmdCmd.Flags().StringVarP(&CommandArgs.JsonArgs, "jsonargs", "j", "", "json arguments of command")

	newCmd.Flags().StringVarP(&CommandArgs.Name, "name", "n", "", "container name")
	newCmd.Flags().StringVarP(&CommandArgs.Hostname, "hostname", "", "h", "hostname")
	newCmd.Flags().StringVarP(&CommandArgs.Ports, "port", "p", "", "port")
	newCmd.Flags().StringVarP(&CommandArgs.Env, "env", "E", "", "env")
	newCmd.Flags().StringVarP(&CommandArgs.Sshkey, "sshkey", "", "", "sshkey")
}
