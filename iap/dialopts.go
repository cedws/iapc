package iap

type DialOption func(*dialOptions)

type dialOptions struct {
	Compress  bool
	Project   string
	Port      string
	Instance  string
	Zone      string
	Interface string
	Host      string
	Group     string
	Network   string
	Region    string
}

func WithProject(project string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Project = project
	}
}

func WithInstance(instance, zone, iinterface string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Instance = instance
		d.Zone = zone
		d.Interface = iinterface
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

func WithCompression() func(*dialOptions) {
	return func(d *dialOptions) {
		d.Compress = true
	}
}
