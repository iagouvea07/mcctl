package compute

type InstanceParameters struct {
	InstanceName string
	InstanceAmi string
	InstanceType string 
	InstanceKey string
}

type InstanceDescribe struct {
	InstanceName string
	InstanceId string
	InstanceType string
	InstanceStatus string
	InstancePublicIp string
	InstancePrivateIp string
}