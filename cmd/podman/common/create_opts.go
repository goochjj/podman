package common

import "github.com/containers/libpod/v2/pkg/domain/entities"

type ContainerCLIOpts struct {
	Annotation        []string
	Attach            []string
	Authfile          string
	BlkIOWeight       string
	BlkIOWeightDevice []string
	CapAdd            []string
	CapDrop           []string
	CGroupsNS         string
	CGroupsMode       string
	CGroupParent      string
	CIDFile           string
	ConmonPIDFile     string
	CPUPeriod         uint64
	CPUQuota          int64
	CPURTPeriod       uint64
	CPURTRuntime      int64
	CPUShares         uint64
	CPUS              float64
	CPUSetCPUs        string
	CPUSetMems        string
	Detach            bool
	DetachKeys        string
	Devices           []string
	DeviceCGroupRule  []string
	DeviceReadBPs     []string
	DeviceReadIOPs    []string
	DeviceWriteBPs    []string
	DeviceWriteIOPs   []string
	Entrypoint        *string
	Env               []string
	EnvHost           bool
	EnvFile           []string
	Expose            []string
	GIDMap            []string
	GroupAdd          []string
	HealthCmd         string
	HealthInterval    string
	HealthRetries     uint
	HealthStartPeriod string
	HealthTimeout     string
	Hostname          string
	HTTPProxy         bool
	ImageVolume       string
	Init              bool
	InitPath          string
	Interactive       bool
	IPC               string
	KernelMemory      string
	Label             []string
	LabelFile         []string
	LogDriver         string
	LogOptions        []string
	Memory            string
	MemoryReservation string
	MemorySwap        string
	MemorySwappiness  int64
	Name              string
	NoHealthCheck     bool
	OOMKillDisable    bool
	OOMScoreAdj       int
	OverrideArch      string
	OverrideOS        string
	PID               string
	PIDsLimit         *int64
	Pod               string
	PodIDFile         string
	PreserveFDs       uint
	Privileged        bool
	PublishAll        bool
	Pull              string
	Quiet             bool
	ReadOnly          bool
	ReadOnlyTmpFS     bool
	Restart           string
	Replace           bool
	Rm                bool
	RootFS            bool
	SecurityOpt       []string
	SdNotifyMode      string
	ShmSize           string
	StopSignal        string
	StopTimeout       uint
	StoreageOpt       []string
	SubUIDName        string
	SubGIDName        string
	Sysctl            []string
	Systemd           string
	TmpFS             []string
	TTY               bool
	Timezone          string
	UIDMap            []string
	Ulimit            []string
	User              string
	UserNS            string
	UTS               string
	Mount             []string
	Volume            []string
	VolumesFrom       []string
	Workdir           string
	SeccompPolicy     string

	Net *entities.NetOptions
}
