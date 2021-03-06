package report

// node metadata keys
const (
	// probe/endpoint
	ReverseDNSNames = "reverse_dns_names"
	SnoopedDNSNames = "snooped_dns_names"
	CopyOf          = "copy_of"
	// probe/process
	PID     = "pid"
	Name    = "name" // also used by probe/docker
	PPID    = "ppid"
	Cmdline = "cmdline"
	Threads = "threads"

	// Controls
	AdminControl    = "admin_control"
	ReadOnlyControl = "read_only_control"

	// probe/docker
	DockerContainerID            = "docker_container_id"
	DockerImageID                = "docker_image_id"
	DockerImageName              = "docker_image_name"
	DockerImageTag               = "docker_image_tag"
	DockerImageSize              = "docker_image_size"
	DockerImageVirtualSize       = "docker_image_virtual_size"
	DockerIsInHostNetwork        = "docker_is_in_host_network"
	DockerServiceName            = "service_name"
	DockerStackNamespace         = "stack_namespace"
	DockerDefaultNamespace       = "No stack"
	DockerStopContainer          = "docker_stop_container"
	DockerStartContainer         = "docker_start_container"
	DockerRestartContainer       = "docker_restart_container"
	DockerPauseContainer         = "docker_pause_container"
	DockerUnpauseContainer       = "docker_unpause_container"
	DockerRemoveContainer        = "docker_remove_container"
	DockerAttachContainer        = "docker_attach_container"
	DockerExecContainer          = "docker_exec_container"
	DockerContainerName          = "docker_container_name"
	DockerContainerCommand       = "docker_container_command"
	DockerContainerPorts         = "docker_container_ports"
	DockerContainerCreated       = "docker_container_created"
	DockerContainerNetworks      = "docker_container_networks"
	DockerContainerIPs           = "docker_container_ips"
	DockerContainerHostname      = "docker_container_hostname"
	DockerContainerIPsWithScopes = "docker_container_ips_with_scopes"
	DockerContainerState         = "docker_container_state"
	DockerContainerStateHuman    = "docker_container_state_human"
	DockerContainerUptime        = "docker_container_uptime"
	DockerContainerRestartCount  = "docker_container_restart_count"
	DockerContainerNetworkMode   = "docker_container_network_mode"
	DockerEnvPrefix              = "docker_env_"
	// probe/kubernetes
	KubernetesName                         = "kubernetes_name"
	KubernetesNamespace                    = "kubernetes_namespace"
	KubernetesCreated                      = "kubernetes_created"
	KubernetesIP                           = "kubernetes_ip"
	KubernetesObservedGeneration           = "kubernetes_observed_generation"
	KubernetesReplicas                     = "kubernetes_replicas"
	KubernetesDesiredReplicas              = "kubernetes_desired_replicas"
	KubernetesNodeType                     = "kubernetes_node_type"
	KubernetesGetLogs                      = "kubernetes_get_logs"
	KubernetesDeletePod                    = "kubernetes_delete_pod"
	KubernetesScaleUp                      = "kubernetes_scale_up"
	KubernetesScaleDown                    = "kubernetes_scale_down"
	KubernetesUpdatedReplicas              = "kubernetes_updated_replicas"
	KubernetesAvailableReplicas            = "kubernetes_available_replicas"
	KubernetesUnavailableReplicas          = "kubernetes_unavailable_replicas"
	KubernetesStrategy                     = "kubernetes_strategy"
	KubernetesFullyLabeledReplicas         = "kubernetes_fully_labeled_replicas"
	KubernetesState                        = "kubernetes_state"
	KubernetesIsInHostNetwork              = "kubernetes_is_in_host_network"
	KubernetesRestartCount                 = "kubernetes_restart_count"
	KubernetesMisscheduledReplicas         = "kubernetes_misscheduled_replicas"
	KubernetesPublicIP                     = "kubernetes_public_ip"
	KubernetesSchedule                     = "kubernetes_schedule"
	KubernetesSuspended                    = "kubernetes_suspended"
	KubernetesLastScheduled                = "kubernetes_last_scheduled"
	KubernetesActiveJobs                   = "kubernetes_active_jobs"
	KubernetesType                         = "kubernetes_type"
	KubernetesPorts                        = "kubernetes_ports"
	KubernetesVolumeClaim                  = "kubernetes_volume_claim"
	KubernetesStorageClassName             = "kubernetes_storage_class_name"
	KubernetesAccessModes                  = "kubernetes_access_modes"
	KubernetesReclaimPolicy                = "kubernetes_reclaim_policy"
	KubernetesStatus                       = "kubernetes_status"
	KubernetesMessage                      = "kubernetes_message"
	KubernetesVolumeName                   = "kubernetes_volume_name"
	KubernetesProvisioner                  = "kubernetes_provisioner"
	KubernetesStorageDriver                = "kubernetes_storage_driver"
	KubernetesVolumeSnapshotName           = "kubernetes_volume_snapshot_name"
	KubernetesVolumeSnapshotNamespace      = "kubernetes_volume_snapshot_namespace"
	KubernetesSnapshotData                 = "kubernetes_snapshot_data"
	KubernetesSnapshotClass                = "kubernetes_snapshot_class"
	KubernetesCreateVolumeSnapshot         = "kubernetes_create_volume_snapshot"
	KubernetesVolumeCapacity               = "kubernetes_volume_capacity"
	KubernetesCloneVolumeSnapshot          = "kubernetes_clone_volume_snapshot"
	KubernetesCloneCsiVolumeSnapshot       = "kubernetes_clone_csi_volume_snapshot"
	KubernetesDeleteVolumeSnapshot         = "kubernetes_delete_volume_snapshot"
	KubernetesDeleteCsiVolumeSnapshot      = "kubernetes_delete_csi_volume_snapshot"
	KubernetesDescribe                     = "kubernetes_describe"
	KubernetesModel                        = "kubernetes_model"
	KubernetesLogicalSectorSize            = "kubernetes_logical_sector_size"
	KubernetesFirmwareRevision             = "kubernetes_firmware_version"
	KubernetesSerial                       = "kubernetes_serial"
	KubernetesVendor                       = "kubernetes_vendor"
	KubernetesDiskList                     = "kubernetes_disk_list"
	KubernetesMaxPools                     = "kubernetes_max_pools"
	KubernetesAPIVersion                   = "kubernetes_api_version"
	KubernetesValue                        = "kubernetes_value"
	KubernetesStoragePoolClaimName         = "kubernetes_storage_pool_claim"
	KubernetesDiskName                     = "kubernetes_disk_name"
	KubernetesPoolName                     = "kubernetes_pool_name"
	KubernetesPoolClaim                    = "kubernetes_pool_claim"
	KubernetesHostName                     = "kubernetes_host_name"
	KubernetesVolumePod                    = "kubernetes_volume_pod"
	KubernetesCStorVolumeName              = "kubernetes_cstor_volume"
	KubernetesCStorVolumeReplicaName       = "kubernetes_cstor_volume_replica"
	KubernetesCStorPoolName                = "kubernetes_cstor_pool"
	KubernetesCStorPoolUID                 = "kubernetes_cstor_pool_uid"
	KubernetesCStorVolumeConsistencyFactor = "kubernetes_cstor_volume_consistency_factor"
	KubernetesCStorVolumeReplicationFactor = "kubernetes_cstor_volume_replication_factor"
	KubernetesCStorVolumeIQN               = "kubernetes_cstor_volume_iqn"
	KubernetesPhysicalSectorSize           = "kubernetes_physical_sector_size"
	KubernetesRotationRate                 = "kubernetes_rotation_rate"
	KubernetesCurrentTemperature           = "kubernetes_current_temperature"
	KubernetesHighestTemperature           = "kubernetes_highest_temperature"
	KubernetesLowestTemperature            = "kubernetes_lowest_temperature"
	KubernetesTotalBytesRead               = "kubernetes_total_bytes_read"
	KubernetesTotalBytesWritten            = "kubernetes_total_bytes_written"
	KubernetesDeviceUtilizationRate        = "kubernetes_device_utilization_rate"
	KubernetesPercentEnduranceUsed         = "kubernetes_percent_endurance_used"
	KubernetesStorage                      = "kubernetes_storage"
	KubernetesBlockDeviceList              = "kubernetes_block_device_list"
	KubernetesPath                         = "kubernetes_path"
	KubernetesBlockDeviceName              = "kubernetes_block_device_name"
	KubernetesBlockDeviceClaimName         = "kubernetes_block_device_claim_name"
	KubernetesCASType                      = "kubernetes_cas_type"
	KubernetesTotalSize                    = "kubernetes_total_size"
	KubernetesFreeSize                     = "kubernetes_free_size"
	KubernetesUsedSize                     = "kubernetes_used_size"
	KubernetesLogicalUsed                  = "kubernetes_logical_used"
	KubernetesProvisionedInstances         = "kubernetes_provisioned_instances"
	KubernetesDesiredInstances             = "kubernetes_desired_instances"
	KubernetesHealthyInstances             = "kubernetes_healthy_instances"
	KubernetesReadOnly                     = "kubernetes_read_only"
	KubernetesProvisionedReplicas          = "kubernetes_provisioned_replicas"
	KubernetesHealthyReplicas              = "kubernetes_healthy_replicas"
	KubernetesCStorPoolInstanceUID         = "kubernetes_cstor_pool_instance"
	KubernetesNodeName                     = "kubernetes_node_name"
	KubernetesDriver                       = "kubernetes_driver"
	KubernetesDeletionPolicy               = "kubernetes_deletion_policy"
	// probe/awsecs
	ECSCluster             = "ecs_cluster"
	ECSCreatedAt           = "ecs_created_at"
	ECSTaskFamily          = "ecs_task_family"
	ECSServiceDesiredCount = "ecs_service_desired_count"
	ECSServiceRunningCount = "ecs_service_running_count"
	ECSScaleUp             = "ecs_scale_up"
	ECSScaleDown           = "ecs_scale_down"
	// probe/host
	Timestamp         = "ts"
	HostName          = "host_name"
	HostLocalNetworks = "local_networks"
	OS                = "os"
	KernelVersion     = "kernel_version"
	Uptime            = "uptime"
	Load1             = "load1"
	HostCPUUsage      = "host_cpu_usage_percent"
	HostMemoryUsage   = "host_mem_usage_bytes"
	ScopeVersion      = "host_scope_version"
	// probe/overlay/weave
	WeavePeerName     = "weave_peer_name"
	WeavePeerNickName = "weave_peer_nick_name"
)

// User kind keys
const (
	UserKindHeader = "x-api-user-kind"
	ReadAdminUSer  = "readAdmin"
)

/* Lookup table to allow msgpack/json decoder to avoid heap allocation
   for common ps.Map keys. The map is static so we don't have to lock
   access from multiple threads and don't have to worry about it
   getting clogged with values that are only used once.
*/
var commonKeys = map[string]string{
	Endpoint:              Endpoint,
	Process:               Process,
	Container:             Container,
	Pod:                   Pod,
	Service:               Service,
	Deployment:            Deployment,
	ReplicaSet:            ReplicaSet,
	DaemonSet:             DaemonSet,
	StatefulSet:           StatefulSet,
	CronJob:               CronJob,
	ContainerImage:        ContainerImage,
	Host:                  Host,
	Overlay:               Overlay,
	ECSService:            ECSService,
	ECSTask:               ECSTask,
	SwarmService:          SwarmService,
	PersistentVolume:      PersistentVolume,
	PersistentVolumeClaim: PersistentVolumeClaim,
	StorageClass:          StorageClass,
	VolumeSnapshot:        VolumeSnapshot,
	VolumeSnapshotData:    VolumeSnapshotData,

	HostNodeID:             HostNodeID,
	ControlProbeID:         ControlProbeID,
	DoesNotMakeConnections: DoesNotMakeConnections,

	ReverseDNSNames: ReverseDNSNames,
	SnoopedDNSNames: SnoopedDNSNames,
	CopyOf:          CopyOf,

	PID:     PID,
	Name:    Name,
	PPID:    PPID,
	Cmdline: Cmdline,
	Threads: Threads,

	DockerContainerID:            DockerContainerID,
	DockerImageID:                DockerImageID,
	DockerImageName:              DockerImageName,
	DockerImageTag:               DockerImageTag,
	DockerImageSize:              DockerImageSize,
	DockerImageVirtualSize:       DockerImageVirtualSize,
	DockerIsInHostNetwork:        DockerIsInHostNetwork,
	DockerServiceName:            DockerServiceName,
	DockerStackNamespace:         DockerStackNamespace,
	DockerStopContainer:          DockerStopContainer,
	DockerStartContainer:         DockerStartContainer,
	DockerRestartContainer:       DockerRestartContainer,
	DockerPauseContainer:         DockerPauseContainer,
	DockerUnpauseContainer:       DockerUnpauseContainer,
	DockerRemoveContainer:        DockerRemoveContainer,
	DockerAttachContainer:        DockerAttachContainer,
	DockerExecContainer:          DockerExecContainer,
	DockerContainerName:          DockerContainerName,
	DockerContainerCommand:       DockerContainerCommand,
	DockerContainerPorts:         DockerContainerPorts,
	DockerContainerCreated:       DockerContainerCreated,
	DockerContainerNetworks:      DockerContainerNetworks,
	DockerContainerIPs:           DockerContainerIPs,
	DockerContainerHostname:      DockerContainerHostname,
	DockerContainerIPsWithScopes: DockerContainerIPsWithScopes,
	DockerContainerState:         DockerContainerState,
	DockerContainerStateHuman:    DockerContainerStateHuman,
	DockerContainerUptime:        DockerContainerUptime,
	DockerContainerRestartCount:  DockerContainerRestartCount,
	DockerContainerNetworkMode:   DockerContainerNetworkMode,

	KubernetesName:                 KubernetesName,
	KubernetesNamespace:            KubernetesNamespace,
	KubernetesCreated:              KubernetesCreated,
	KubernetesIP:                   KubernetesIP,
	KubernetesObservedGeneration:   KubernetesObservedGeneration,
	KubernetesReplicas:             KubernetesReplicas,
	KubernetesDesiredReplicas:      KubernetesDesiredReplicas,
	KubernetesNodeType:             KubernetesNodeType,
	KubernetesGetLogs:              KubernetesGetLogs,
	KubernetesDeletePod:            KubernetesDeletePod,
	KubernetesScaleUp:              KubernetesScaleUp,
	KubernetesScaleDown:            KubernetesScaleDown,
	KubernetesUpdatedReplicas:      KubernetesUpdatedReplicas,
	KubernetesAvailableReplicas:    KubernetesAvailableReplicas,
	KubernetesUnavailableReplicas:  KubernetesUnavailableReplicas,
	KubernetesStrategy:             KubernetesStrategy,
	KubernetesFullyLabeledReplicas: KubernetesFullyLabeledReplicas,
	KubernetesState:                KubernetesState,
	KubernetesIsInHostNetwork:      KubernetesIsInHostNetwork,
	KubernetesRestartCount:         KubernetesRestartCount,
	KubernetesMisscheduledReplicas: KubernetesMisscheduledReplicas,
	KubernetesPublicIP:             KubernetesPublicIP,
	KubernetesSchedule:             KubernetesSchedule,
	KubernetesSuspended:            KubernetesSuspended,
	KubernetesLastScheduled:        KubernetesLastScheduled,
	KubernetesActiveJobs:           KubernetesActiveJobs,
	KubernetesType:                 KubernetesType,
	KubernetesPorts:                KubernetesPorts,

	ECSCluster:             ECSCluster,
	ECSCreatedAt:           ECSCreatedAt,
	ECSTaskFamily:          ECSTaskFamily,
	ECSServiceDesiredCount: ECSServiceDesiredCount,
	ECSServiceRunningCount: ECSServiceRunningCount,
	ECSScaleUp:             ECSScaleUp,
	ECSScaleDown:           ECSScaleDown,

	Timestamp:         Timestamp,
	HostName:          HostName,
	HostLocalNetworks: HostLocalNetworks,
	OS:                OS,
	KernelVersion:     KernelVersion,
	Uptime:            Uptime,
	Load1:             Load1,
	HostCPUUsage:      HostCPUUsage,
	HostMemoryUsage:   HostMemoryUsage,
	ScopeVersion:      ScopeVersion,

	WeavePeerName:     WeavePeerName,
	WeavePeerNickName: WeavePeerNickName,
}

func lookupCommonKey(b []byte) string {
	if key, ok := commonKeys[string(b)]; ok {
		return key
	}
	return string(b)
}