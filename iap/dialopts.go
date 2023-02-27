package iap

type DialOption func(*dialOptions)

type dialOptions struct {
	Project   string
	Instance  string
	Zone      string
	Region    string
	Network   string
	Interface string
	Port      string
	Compress  bool
}

func WithProject(project string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Project = project
	}
}

func WithInstance(instance string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Instance = instance
	}
}

func WithZone(zone string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Zone = zone
	}
}

func WithRegion(region string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Region = region
	}
}

func WithNetwork(network string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Network = network
	}
}

func WithInterface(iinterface string) func(*dialOptions) {
	return func(d *dialOptions) {
		d.Interface = iinterface
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
