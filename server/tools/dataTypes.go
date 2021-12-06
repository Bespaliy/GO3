package tools

//Machine JSON struct
type Machine struct {
	Id             int    `json:"-"`
	Name           string `json:"name"`
	CpuCount       int    `json:"cpuCount"`
	TotalDiskSpace int    `json:"totalDiskSpace"`
}

//Forum Disk struct
type Disk struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	DiskSpace int    `json:"diskSpace"`
}

//Error struct
type errorObject struct {
	Message string `json:"message"`
}
