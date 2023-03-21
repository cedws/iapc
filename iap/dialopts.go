package iap

type DialOption func(*dialOptions)

type dialOptions struct {
	Zone      string
	Token     string
	Region    string
	Project   string
	Port      string
	Network   string
	Interface string
	Instance  string
	Host      string
	Group     string
	Compress  bool
}

func WithToken(token string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Token = token
	}
}

func WithCompression() func(*dialOptions) {
	return func(d *dialOptions) {
		d.Compress = true
	}
}

func WithProject(project string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Project = project
	}
}

func WithInstance(instance, zone, ninterface string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Instance = instance
		d.Zone = zone
		d.Interface = ninterface
	}
}

func WithHost(host, region, network, destGroup string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Host = host
		d.Region = region
		d.Network = network
		d.Group = destGroup
	}
}

func WithPort(port string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Port = port
	}
}
