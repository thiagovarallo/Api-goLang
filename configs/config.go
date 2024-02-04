package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBNAme        string
	WebServerPort string
	JWTScret      string
	JWTExperesIn  int
}

func LoadConfig(path string) (*conf, error) {
	//..
}
