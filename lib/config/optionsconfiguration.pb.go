// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/optionsconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/syncthing/syncthing/proto/ext"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OptionsConfiguration struct {
	RawListenAddresses      []string `protobuf:"bytes,1,rep,name=listen_addresses,json=listenAddresses,proto3" json:"listenAddresses" xml:"listenAddress" default:"default"`
	RawGlobalAnnServers     []string `protobuf:"bytes,2,rep,name=global_discovery_servers,json=globalDiscoveryServers,proto3" json:"globalAnnounceServers" xml:"globalAnnounceServer" default:"default"`
	GlobalAnnEnabled        bool     `protobuf:"varint,3,opt,name=global_discovery_enabled,json=globalDiscoveryEnabled,proto3" json:"globalAnnounceEnabled" xml:"globalAnnounceEnabled" default:"true"`
	LocalAnnEnabled         bool     `protobuf:"varint,4,opt,name=local_discovery_enabled,json=localDiscoveryEnabled,proto3" json:"localAnnounceEnabled" xml:"localAnnounceEnabled" default:"true"`
	LocalAnnPort            int      `protobuf:"varint,5,opt,name=local_announce_port,json=localAnnouncePort,proto3,casttype=int" json:"localAnnouncePort" xml:"localAnnouncePort" default:"21027"`
	LocalAnnMCAddr          string   `protobuf:"bytes,6,opt,name=local_announce_multicast_address,json=localAnnounceMulticastAddress,proto3" json:"localAnnounceMCAddr" xml:"localAnnounceMCAddr" default:"[ff12::8384]:21027"`
	MaxSendKbps             int      `protobuf:"varint,7,opt,name=max_send_kbps,json=maxSendKbps,proto3,casttype=int" json:"maxSendKbps" xml:"maxSendKbps"`
	MaxRecvKbps             int      `protobuf:"varint,8,opt,name=max_recv_kbps,json=maxRecvKbps,proto3,casttype=int" json:"maxRecvKbps" xml:"maxRecvKbps"`
	ReconnectIntervalS      int      `protobuf:"varint,9,opt,name=reconnection_interval_s,json=reconnectionIntervalS,proto3,casttype=int" json:"reconnectionIntervalS" xml:"reconnectionIntervalS" default:"60"`
	RelaysEnabled           bool     `protobuf:"varint,10,opt,name=relays_enabled,json=relaysEnabled,proto3" json:"relaysEnabled" xml:"relaysEnabled" default:"true"`
	RelayReconnectIntervalM int      `protobuf:"varint,11,opt,name=relays_reconnect_interval_m,json=relaysReconnectIntervalM,proto3,casttype=int" json:"relayReconnectIntervalM" xml:"relayReconnectIntervalM" default:"10"`
	StartBrowser            bool     `protobuf:"varint,12,opt,name=start_browser,json=startBrowser,proto3" json:"startBrowser" xml:"startBrowser" default:"true"`
	NATEnabled              bool     `protobuf:"varint,14,opt,name=nat_traversal_enabled,json=natTraversalEnabled,proto3" json:"natEnabled" xml:"natEnabled" default:"true"`
	NATLeaseM               int      `protobuf:"varint,15,opt,name=nat_traversal_lease_m,json=natTraversalLeaseM,proto3,casttype=int" json:"natLeaseMinutes" xml:"natLeaseMinutes" default:"60"`
	NATRenewalM             int      `protobuf:"varint,16,opt,name=nat_traversal_renewal_m,json=natTraversalRenewalM,proto3,casttype=int" json:"natRenewalMinutes" xml:"natRenewalMinutes" default:"30"`
	NATTimeoutS             int      `protobuf:"varint,17,opt,name=nat_traversal_timeout_s,json=natTraversalTimeoutS,proto3,casttype=int" json:"natTimeoutSeconds" xml:"natTimeoutSeconds" default:"10"`
	URAccepted              int      `protobuf:"varint,18,opt,name=usage_reporting_accepted,json=usageReportingAccepted,proto3,casttype=int" json:"urAccepted" xml:"urAccepted"`
	URSeen                  int      `protobuf:"varint,19,opt,name=usage_reporting_seen,json=usageReportingSeen,proto3,casttype=int" json:"urSeen" xml:"urSeen"`
	URUniqueID              string   `protobuf:"bytes,20,opt,name=usage_reporting_unique_id,json=usageReportingUniqueId,proto3" json:"urUniqueId" xml:"urUniqueID"`
	URURL                   string   `protobuf:"bytes,21,opt,name=usage_reporting_url,json=usageReportingUrl,proto3" json:"urURL" xml:"urURL" default:"https://data.syncthing.net/newdata"`
	URPostInsecurely        bool     `protobuf:"varint,22,opt,name=usage_reporting_post_insecurely,json=usageReportingPostInsecurely,proto3" json:"urPostInsecurely" xml:"urPostInsecurely" default:"false"`
	URInitialDelayS         int      `protobuf:"varint,23,opt,name=usage_reporting_initial_delay_s,json=usageReportingInitialDelayS,proto3,casttype=int" json:"urInitialDelayS" xml:"urInitialDelayS" default:"1800"`
	RestartOnWakeup         bool     `protobuf:"varint,24,opt,name=restart_on_wakeup,json=restartOnWakeup,proto3" json:"restartOnWakeup" xml:"restartOnWakeup" default:"true"`
	AutoUpgradeIntervalH    int      `protobuf:"varint,25,opt,name=auto_upgrade_interval_h,json=autoUpgradeIntervalH,proto3,casttype=int" json:"autoUpgradeIntervalH" xml:"autoUpgradeIntervalH" default:"12"`
	UpgradeToPreReleases    bool     `protobuf:"varint,26,opt,name=upgrade_to_pre_releases,json=upgradeToPreReleases,proto3" json:"upgradeToPreReleases" xml:"upgradeToPreReleases"`
	KeepTemporariesH        int      `protobuf:"varint,27,opt,name=keep_temporaries_h,json=keepTemporariesH,proto3,casttype=int" json:"keepTemporariesH" xml:"keepTemporariesH" default:"24"`
	CacheIgnoredFiles       bool     `protobuf:"varint,28,opt,name=cache_ignored_files,json=cacheIgnoredFiles,proto3" json:"cacheIgnoredFiles" xml:"cacheIgnoredFiles" default:"false"`
	ProgressUpdateIntervalS int      `protobuf:"varint,29,opt,name=progress_update_interval_s,json=progressUpdateIntervalS,proto3,casttype=int" json:"progressUpdateIntervalS" xml:"progressUpdateIntervalS" default:"5"`
	LimitBandwidthInLan     bool     `protobuf:"varint,30,opt,name=limit_bandwidth_in_lan,json=limitBandwidthInLan,proto3" json:"limitBandwidthInLan" xml:"limitBandwidthInLan" default:"false"`
	MinHomeDiskFree         Size     `protobuf:"bytes,31,opt,name=min_home_disk_free,json=minHomeDiskFree,proto3" json:"minHomeDiskFree" xml:"minHomeDiskFree" default:"1 %"`
	ReleasesURL             string   `protobuf:"bytes,32,opt,name=releases_url,json=releasesUrl,proto3" json:"releasesURL" xml:"releasesURL" default:"https://upgrades.syncthing.net/meta.json"`
	AlwaysLocalNets         []string `protobuf:"bytes,33,rep,name=always_local_nets,json=alwaysLocalNets,proto3" json:"alwaysLocalNets" xml:"alwaysLocalNet"`
	OverwriteRemoteDevNames bool     `protobuf:"varint,34,opt,name=overwrite_remote_device_names_on_connect,json=overwriteRemoteDeviceNamesOnConnect,proto3" json:"overwriteRemoteDeviceNamesOnConnect" xml:"overwriteRemoteDeviceNamesOnConnect" default:"false"`
	TempIndexMinBlocks      int      `protobuf:"varint,35,opt,name=temp_index_min_blocks,json=tempIndexMinBlocks,proto3,casttype=int" json:"tempIndexMinBlocks" xml:"tempIndexMinBlocks" default:"10"`
	UnackedNotificationIDs  []string `protobuf:"bytes,36,rep,name=unacked_notification_ids,json=unackedNotificationIds,proto3" json:"unackedNotificationIDs" xml:"unackedNotificationID"`
	TrafficClass            int      `protobuf:"varint,37,opt,name=traffic_class,json=trafficClass,proto3,casttype=int" json:"trafficClass" xml:"trafficClass"`
	DefaultFolderPath       string   `protobuf:"bytes,38,opt,name=default_folder_path,json=defaultFolderPath,proto3" json:"defaultFolderPath" xml:"defaultFolderPath" default:"~"`
	SetLowPriority          bool     `protobuf:"varint,39,opt,name=set_low_priority,json=setLowPriority,proto3" json:"setLowPriority" xml:"setLowPriority" default:"true"`
	RawMaxFolderConcurrency int      `protobuf:"varint,40,opt,name=max_folder_concurrency,json=maxFolderConcurrency,proto3,casttype=int" json:"maxFolderConcurrency" xml:"maxFolderConcurrency"`
	CRURL                   string   `protobuf:"bytes,41,opt,name=crash_reporting_url,json=crashReportingUrl,proto3" json:"crURL" xml:"crashReportingURL" default:"https://crash.syncthing.net/newcrash"`
	CREnabled               bool     `protobuf:"varint,42,opt,name=crash_reporting_enabled,json=crashReportingEnabled,proto3" json:"crashReportingEnabled" xml:"crashReportingEnabled" default:"true"`
	StunKeepaliveStartS     int      `protobuf:"varint,43,opt,name=stun_keepalive_start_s,json=stunKeepaliveStartS,proto3,casttype=int" json:"stunKeepaliveStartS" xml:"stunKeepaliveStartS" default:"180"`
	StunKeepaliveMinS       int      `protobuf:"varint,44,opt,name=stun_keepalive_min_s,json=stunKeepaliveMinS,proto3,casttype=int" json:"stunKeepaliveMinS" xml:"stunKeepaliveMinS" default:"20"`
	RawStunServers          []string `protobuf:"bytes,45,rep,name=stun_servers,json=stunServers,proto3" json:"stunServers" xml:"stunServer" default:"default"`
	DatabaseTuning          Tuning   `protobuf:"varint,46,opt,name=database_tuning,json=databaseTuning,proto3,enum=config.Tuning" json:"databaseTuning" xml:"databaseTuning" restart:"true"`
	RawMaxCIRequestKiB      int      `protobuf:"varint,47,opt,name=max_concurrent_incoming_request_kib,json=maxConcurrentIncomingRequestKib,proto3,casttype=int" json:"maxConcurrentIncomingRequestKiB" xml:"maxConcurrentIncomingRequestKiB"`
	AnnounceLANAddresses    bool     `protobuf:"varint,48,opt,name=announce_lan_addresses,json=announceLanAddresses,proto3" json:"announceLANAddresses" xml:"announceLANAddresses" default:"true"`
	SendFullIndexOnUpgrade  bool     `protobuf:"varint,49,opt,name=send_full_index_on_upgrade,json=sendFullIndexOnUpgrade,proto3" json:"sendFullIndexOnUpgrade" xml:"sendFullIndexOnUpgrade"`
	// Legacy deprecated
	DeprecatedUPnPEnabled        bool     `protobuf:"varint,9000,opt,name=upnp_enabled,json=upnpEnabled,proto3" json:"-" xml:"upnpEnabled,omitempty"`                                    // Deprecated: Do not use.
	DeprecatedUPnPLeaseM         int      `protobuf:"varint,9001,opt,name=upnp_lease_m,json=upnpLeaseM,proto3,casttype=int" json:"-" xml:"upnpLeaseMinutes,omitempty"`                   // Deprecated: Do not use.
	DeprecatedUPnPRenewalM       int      `protobuf:"varint,9002,opt,name=upnp_renewal_m,json=upnpRenewalM,proto3,casttype=int" json:"-" xml:"upnpRenewalMinutes,omitempty"`             // Deprecated: Do not use.
	DeprecatedUPnPTimeoutS       int      `protobuf:"varint,9003,opt,name=upnp_timeout_s,json=upnpTimeoutS,proto3,casttype=int" json:"-" xml:"upnpTimeoutSeconds,omitempty"`             // Deprecated: Do not use.
	DeprecatedRelayServers       []string `protobuf:"bytes,9004,rep,name=relay_servers,json=relayServers,proto3" json:"-" xml:"relayServer,omitempty"`                                   // Deprecated: Do not use.
	DeprecatedMinHomeDiskFreePct float64  `protobuf:"fixed64,9005,opt,name=min_home_disk_free_pct,json=minHomeDiskFreePct,proto3" json:"-" xml:"minHomeDiskFreePct,omitempty"`           // Deprecated: Do not use.
	DeprecatedMaxConcurrentScans int      `protobuf:"varint,9006,opt,name=max_concurrent_scans,json=maxConcurrentScans,proto3,casttype=int" json:"-" xml:"maxConcurrentScans,omitempty"` // Deprecated: Do not use.
}

func (m *OptionsConfiguration) Reset()         { *m = OptionsConfiguration{} }
func (m *OptionsConfiguration) String() string { return proto.CompactTextString(m) }
func (*OptionsConfiguration) ProtoMessage()    {}
func (*OptionsConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_d09882599506ca03, []int{0}
}
func (m *OptionsConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionsConfiguration.Unmarshal(m, b)
}
func (m *OptionsConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionsConfiguration.Marshal(b, m, deterministic)
}
func (m *OptionsConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionsConfiguration.Merge(m, src)
}
func (m *OptionsConfiguration) XXX_Size() int {
	return xxx_messageInfo_OptionsConfiguration.Size(m)
}
func (m *OptionsConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionsConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_OptionsConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OptionsConfiguration)(nil), "config.OptionsConfiguration")
}

func init() {
	proto.RegisterFile("lib/config/optionsconfiguration.proto", fileDescriptor_d09882599506ca03)
}

var fileDescriptor_d09882599506ca03 = []byte{
	// 3101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x5a, 0x5b, 0x6c, 0x1d, 0x47,
	0x19, 0xce, 0x26, 0x4d, 0xda, 0x6c, 0x1c, 0x27, 0x5e, 0x3b, 0xf6, 0x36, 0x49, 0xbd, 0xee, 0xc9,
	0x49, 0xeb, 0xde, 0xe2, 0x4b, 0xda, 0x10, 0x2c, 0x21, 0xf0, 0xa5, 0x6e, 0x4d, 0x6d, 0xc7, 0x1a,
	0xdb, 0x2a, 0x2a, 0x42, 0xab, 0x39, 0x7b, 0xe6, 0xd8, 0x8b, 0xf7, 0xcc, 0x9e, 0xec, 0xce, 0xfa,
	0x52, 0x50, 0xa9, 0x8a, 0xb8, 0xbc, 0x01, 0x16, 0x17, 0x09, 0x10, 0x2a, 0x02, 0x24, 0x4a, 0x29,
	0x42, 0x42, 0x42, 0x82, 0x17, 0x10, 0x12, 0x52, 0x05, 0x0f, 0xf6, 0x23, 0x12, 0x65, 0x51, 0x9d,
	0x3e, 0xa0, 0xf3, 0xc0, 0xc3, 0x79, 0x34, 0x2f, 0xe8, 0x9f, 0xbd, 0xcd, 0xee, 0xce, 0x69, 0xf2,
	0x76, 0xf6, 0xff, 0xfe, 0xf9, 0xe7, 0xfb, 0xe7, 0xf2, 0xcf, 0xff, 0xcf, 0x1c, 0xf5, 0xba, 0x63,
	0xd7, 0xc6, 0x2c, 0x97, 0x36, 0xec, 0x8d, 0x31, 0xb7, 0xc5, 0x6c, 0x97, 0xfa, 0xd1, 0x57, 0xe0,
	0x61, 0xf8, 0xba, 0xd1, 0xf2, 0x5c, 0xe6, 0x6a, 0x67, 0x22, 0xe1, 0xe5, 0x21, 0x41, 0x9d, 0x05,
	0xd4, 0xa6, 0x1b, 0x91, 0xc2, 0xe5, 0x4b, 0x02, 0xe0, 0xdb, 0xaf, 0x93, 0x58, 0x7c, 0x96, 0xec,
	0xb2, 0xe8, 0x67, 0xe5, 0xc7, 0xf3, 0xea, 0xc0, 0x9d, 0xa8, 0x87, 0x59, 0xb1, 0x07, 0xed, 0x27,
	0x8a, 0x7a, 0xd1, 0xb1, 0x7d, 0x46, 0xa8, 0x89, 0xeb, 0x75, 0x8f, 0xf8, 0x3e, 0xf1, 0x75, 0x65,
	0xe4, 0xd4, 0xe8, 0xd9, 0x19, 0xff, 0x28, 0x34, 0x34, 0x84, 0x77, 0x16, 0x39, 0x3c, 0x9d, 0xa0,
	0xed, 0xd0, 0xb8, 0xe0, 0xe4, 0x45, 0x9d, 0xd0, 0xb8, 0xbe, 0xdb, 0x74, 0xa6, 0x2a, 0x39, 0x79,
	0x65, 0xa4, 0x4e, 0x1a, 0x38, 0x70, 0xd8, 0x54, 0x25, 0xfe, 0x51, 0x39, 0x3e, 0xa8, 0x3e, 0x1c,
	0xff, 0xde, 0x3f, 0xac, 0x4a, 0x8c, 0xa3, 0xa2, 0x69, 0xed, 0xbf, 0x8a, 0xaa, 0x6f, 0x38, 0x6e,
	0x0d, 0x3b, 0x66, 0xdd, 0xf6, 0x2d, 0x77, 0x9b, 0x78, 0x7b, 0xa6, 0x4f, 0xbc, 0x6d, 0xe2, 0xf9,
	0xfa, 0x49, 0x4e, 0xf4, 0x77, 0xca, 0x51, 0x68, 0xf4, 0x23, 0xbc, 0xf3, 0x12, 0xd7, 0x9b, 0xa6,
	0x74, 0x35, 0xc2, 0xdb, 0xa1, 0x71, 0x69, 0x23, 0x91, 0xb9, 0x01, 0xb5, 0x48, 0x0c, 0x74, 0x42,
	0xe3, 0x59, 0x4e, 0x58, 0x86, 0x4a, 0x78, 0xb7, 0x0f, 0xaa, 0x03, 0x32, 0xd5, 0xce, 0x41, 0x55,
	0xde, 0x41, 0xde, 0x51, 0x19, 0x37, 0x34, 0x18, 0x35, 0x9c, 0x4b, 0x9c, 0x8a, 0xe5, 0xda, 0x47,
	0x32, 0x87, 0x09, 0xc5, 0x35, 0x87, 0xd4, 0xf5, 0x53, 0x23, 0xca, 0xe8, 0x23, 0x33, 0xef, 0x80,
	0xc3, 0x17, 0x53, 0x8b, 0x2f, 0x46, 0x60, 0xd9, 0xdb, 0x18, 0xe8, 0x84, 0xc6, 0xd3, 0x12, 0x6f,
	0x63, 0x54, 0x70, 0x97, 0x79, 0x01, 0x01, 0x5f, 0xbb, 0x98, 0xe9, 0x06, 0x1c, 0x1f, 0x54, 0x1f,
	0x82, 0xa6, 0xfb, 0x87, 0xd5, 0x12, 0xa9, 0x92, 0x9b, 0xb1, 0x5c, 0xfb, 0x40, 0x51, 0x87, 0x1c,
	0xd7, 0x92, 0x7a, 0xf9, 0x10, 0xf7, 0xf2, 0x67, 0xe0, 0xe5, 0x85, 0x45, 0xd0, 0xc9, 0x39, 0x39,
	0xe0, 0xc4, 0xa2, 0x82, 0x8f, 0x4f, 0x45, 0x4b, 0x50, 0x02, 0x4a, 0x5c, 0x94, 0x1b, 0xe9, 0x22,
	0x17, 0x1c, 0x2c, 0xf2, 0x41, 0x97, 0x78, 0x83, 0x92, 0x7b, 0x7f, 0x57, 0xd4, 0xfe, 0xc8, 0x3d,
	0x1c, 0xdb, 0x32, 0x5b, 0xae, 0xc7, 0xf4, 0xd3, 0x23, 0xca, 0xe8, 0xe9, 0x99, 0x1f, 0x82, 0x6b,
	0x3d, 0x89, 0xa9, 0x15, 0xd7, 0x63, 0xed, 0xd0, 0xe8, 0xcb, 0x75, 0x0d, 0xc2, 0x4e, 0x68, 0x3c,
	0x59, 0x76, 0x0a, 0x10, 0xc1, 0xa3, 0xc9, 0x89, 0xf1, 0xc9, 0x4f, 0x54, 0x8e, 0x43, 0xe3, 0x94,
	0x4d, 0x59, 0xfb, 0xa0, 0x2a, 0x31, 0x23, 0x13, 0x1e, 0x1f, 0x54, 0x4f, 0xf3, 0xa6, 0xfb, 0x87,
	0xd5, 0x1c, 0x13, 0x54, 0xd6, 0xd5, 0xbe, 0x7a, 0x52, 0x1d, 0x29, 0x78, 0xd3, 0x0c, 0x1c, 0x66,
	0x5b, 0xd8, 0x67, 0x49, 0xdc, 0xd0, 0xcf, 0x8c, 0x28, 0xa3, 0x67, 0x67, 0xfe, 0x00, 0xae, 0xf5,
	0x26, 0x06, 0x97, 0x66, 0x61, 0x27, 0xb7, 0x43, 0xa3, 0x3f, 0x67, 0x34, 0x12, 0x77, 0x42, 0xe3,
	0x56, 0xd9, 0xbd, 0x08, 0x13, 0x1c, 0xfc, 0x7c, 0xa3, 0x31, 0x31, 0x39, 0x35, 0x75, 0xfb, 0xe6,
	0xed, 0xe7, 0xbf, 0x30, 0x15, 0x79, 0xdb, 0x3e, 0xa8, 0x4a, 0x0d, 0xca, 0xc5, 0xc7, 0x07, 0x55,
	0xad, 0x6c, 0x64, 0xff, 0xb0, 0x5a, 0xa0, 0x89, 0x1e, 0xcb, 0x37, 0x4e, 0x3c, 0x8c, 0x83, 0x91,
	0x76, 0x47, 0x3d, 0xdf, 0xc4, 0xbb, 0xa6, 0x4f, 0x68, 0xdd, 0xdc, 0xaa, 0xb5, 0x7c, 0xfd, 0x61,
	0x3e, 0x99, 0xcf, 0xb4, 0x43, 0xe3, 0x5c, 0x13, 0xef, 0xae, 0x12, 0x5a, 0x7f, 0xa5, 0xd6, 0x82,
	0xe0, 0xd2, 0xc7, 0xdd, 0x12, 0x64, 0xc9, 0xfc, 0x20, 0x51, 0x31, 0x31, 0xe8, 0x11, 0x6b, 0x3b,
	0x32, 0xf8, 0x48, 0xce, 0x20, 0x22, 0xd6, 0x76, 0xd1, 0x60, 0x22, 0xcb, 0x19, 0x4c, 0x84, 0xda,
	0xef, 0x15, 0x75, 0xc8, 0x23, 0x96, 0x4b, 0x29, 0xb1, 0x20, 0xbc, 0x9b, 0x36, 0x65, 0xc4, 0xdb,
	0xc6, 0x8e, 0xe9, 0xeb, 0x67, 0xb9, 0xed, 0x37, 0x78, 0x50, 0x4f, 0x54, 0x16, 0x62, 0x78, 0x15,
	0x62, 0x87, 0xd8, 0x30, 0x05, 0x3a, 0xa1, 0x31, 0xca, 0xfb, 0x96, 0xa2, 0xc2, 0x2c, 0xdd, 0x1a,
	0x4f, 0x28, 0x1d, 0x1f, 0x54, 0x4f, 0xde, 0x1a, 0xe7, 0xf1, 0xbd, 0xd4, 0x0f, 0x92, 0xf7, 0xa2,
	0x35, 0xd4, 0x5e, 0x8f, 0x38, 0x78, 0xcf, 0x4f, 0x63, 0x80, 0xca, 0x63, 0xc0, 0xa7, 0xdb, 0xa1,
	0x71, 0x3e, 0x42, 0xb2, 0x8d, 0x5e, 0x89, 0x09, 0x09, 0xd2, 0xe2, 0x0e, 0x4f, 0x76, 0x2c, 0xca,
	0x37, 0xd6, 0xde, 0x3a, 0xa9, 0x5e, 0x89, 0x3b, 0x4a, 0x89, 0x64, 0x83, 0xd4, 0xd4, 0xcf, 0xf1,
	0x41, 0xfa, 0x0b, 0xac, 0xe1, 0x21, 0x04, 0x7a, 0x25, 0x17, 0x96, 0xda, 0xa1, 0x31, 0xe4, 0xc9,
	0xa1, 0x34, 0xd0, 0x76, 0xc1, 0x05, 0x96, 0x13, 0xe3, 0xc2, 0x96, 0xed, 0x6a, 0xaf, 0x3b, 0x04,
	0x83, 0x3c, 0x01, 0x83, 0xdc, 0x8d, 0x26, 0xd2, 0x23, 0x3f, 0xcb, 0x88, 0x56, 0x53, 0xcf, 0xfb,
	0x0c, 0x7b, 0xcc, 0xac, 0x79, 0xee, 0x8e, 0x4f, 0x3c, 0xbd, 0x87, 0x8f, 0xf5, 0xa7, 0xda, 0xa1,
	0xd1, 0xc3, 0x81, 0x99, 0x48, 0xde, 0x09, 0x8d, 0xc7, 0xb9, 0x3b, 0xa2, 0xb0, 0xeb, 0x48, 0xe7,
	0x9a, 0x6a, 0xbf, 0x50, 0xd4, 0x4b, 0x14, 0x33, 0x93, 0x79, 0x18, 0x4e, 0x35, 0xec, 0xa4, 0x13,
	0xdb, 0xcb, 0x3b, 0xbb, 0x7b, 0x14, 0x1a, 0xea, 0xf2, 0xf4, 0x5a, 0x16, 0xd6, 0x55, 0x8a, 0x59,
	0x36, 0xc7, 0x06, 0xef, 0x38, 0x13, 0x49, 0x42, 0xb8, 0xd8, 0x20, 0xf7, 0x25, 0x84, 0x6b, 0xa1,
	0x0b, 0xd4, 0x4f, 0x31, 0x5b, 0x4b, 0xe8, 0x24, 0x0b, 0xe2, 0x8f, 0x25, 0x9e, 0x0e, 0xc1, 0x3e,
	0x31, 0x9b, 0xfa, 0x05, 0xbe, 0x14, 0xbe, 0x0e, 0x4b, 0xe1, 0xec, 0xf2, 0xf4, 0xda, 0x22, 0x88,
	0x61, 0xf2, 0x2f, 0x50, 0xcc, 0xa2, 0x0f, 0x9b, 0x06, 0x8c, 0x27, 0x3f, 0x95, 0x84, 0xac, 0x28,
	0x97, 0xee, 0x8d, 0xf6, 0x41, 0xb5, 0xd4, 0xbe, 0x2c, 0x4a, 0x77, 0x50, 0xd6, 0x31, 0xd2, 0x44,
	0xf6, 0x91, 0x4c, 0xfb, 0x9b, 0xa2, 0x0e, 0xe5, 0xc9, 0x7b, 0x84, 0x92, 0x1d, 0xbe, 0x92, 0x2f,
	0x72, 0xfa, 0xfb, 0x40, 0xff, 0xdc, 0xf2, 0xf4, 0x1a, 0x8a, 0x00, 0x70, 0xa0, 0x8f, 0x62, 0x96,
	0x7c, 0xa6, 0x2e, 0x54, 0x13, 0x17, 0xf2, 0x88, 0xe0, 0xc4, 0x4d, 0xd1, 0x09, 0x89, 0x0d, 0x99,
	0x10, 0x1c, 0xb9, 0x09, 0x8e, 0x88, 0x14, 0xd0, 0x80, 0xe8, 0x4a, 0x22, 0x95, 0x38, 0xc3, 0xec,
	0x26, 0x71, 0x03, 0x66, 0xfa, 0x7a, 0x5f, 0xde, 0x99, 0xb5, 0x08, 0x58, 0x8d, 0x9d, 0x49, 0x3e,
	0x61, 0xa5, 0xd7, 0x73, 0xce, 0xe4, 0x91, 0x6e, 0xdb, 0x4f, 0x62, 0x43, 0x26, 0x4c, 0xb7, 0x9c,
	0x48, 0x21, 0xef, 0x4c, 0x22, 0xd5, 0x7e, 0xa4, 0xa8, 0x7a, 0xe0, 0xe3, 0x0d, 0x62, 0x7a, 0x04,
	0xce, 0x7d, 0x9b, 0x6e, 0x98, 0xd8, 0xb2, 0x48, 0x8b, 0x91, 0xba, 0xae, 0x71, 0x6f, 0x30, 0xec,
	0x80, 0x75, 0x34, 0x1d, 0x4b, 0x61, 0x07, 0x04, 0x5e, 0xf2, 0xd5, 0x09, 0x8d, 0x8b, 0xdc, 0x89,
	0x4c, 0x24, 0x10, 0x16, 0x15, 0x73, 0x5f, 0xb0, 0xe2, 0x33, 0x93, 0x68, 0x90, 0x53, 0x40, 0x09,
	0x83, 0x44, 0xae, 0x7d, 0x49, 0x1d, 0x28, 0x92, 0xf3, 0x09, 0xa1, 0x7a, 0x3f, 0x27, 0xb6, 0x70,
	0x14, 0x1a, 0x67, 0xd6, 0xd1, 0x2a, 0x21, 0xb4, 0x1d, 0x1a, 0x67, 0x02, 0x0f, 0x7e, 0x75, 0x42,
	0xa3, 0x27, 0x26, 0x04, 0x9f, 0x02, 0x99, 0x44, 0x21, 0xfd, 0xb5, 0x7f, 0x58, 0x8d, 0x9b, 0x23,
	0x2d, 0x4f, 0x00, 0x64, 0xda, 0xf7, 0x14, 0xf5, 0xd1, 0x62, 0xef, 0x01, 0xb5, 0xef, 0x06, 0xc4,
	0xb4, 0xeb, 0xfa, 0x00, 0x4f, 0x22, 0x5e, 0x8b, 0xc6, 0x66, 0x9d, 0x8b, 0x17, 0xe6, 0xa2, 0xb1,
	0x89, 0xbf, 0xc4, 0xb1, 0x49, 0x14, 0x2a, 0xd1, 0xa0, 0x24, 0x9f, 0x1d, 0xf1, 0x2b, 0x1e, 0x94,
	0x04, 0x2b, 0x0e, 0x4a, 0xa2, 0xa5, 0xfd, 0x59, 0x51, 0xfb, 0x4b, 0xbc, 0x3c, 0x47, 0xbf, 0xc4,
	0x19, 0x7d, 0x0b, 0xd6, 0xde, 0xe9, 0x75, 0xb4, 0x8e, 0x16, 0xdb, 0xa1, 0x71, 0x3a, 0xf0, 0xd6,
	0xd1, 0x62, 0x27, 0x34, 0x6e, 0x27, 0x44, 0xd0, 0xa2, 0xb0, 0xba, 0x36, 0x19, 0x6b, 0xf9, 0x53,
	0x63, 0x63, 0x75, 0xcc, 0xf0, 0x0d, 0x7f, 0x8f, 0x5a, 0x6c, 0x13, 0x8a, 0x35, 0x4a, 0xd8, 0x18,
	0x25, 0x3b, 0x20, 0x05, 0xc2, 0xb1, 0x91, 0xe4, 0xc7, 0xf1, 0x41, 0xf5, 0x01, 0x1a, 0xee, 0x1f,
	0x56, 0x23, 0x16, 0xa8, 0xaf, 0xe0, 0x87, 0xe7, 0x68, 0xff, 0x56, 0x54, 0xa3, 0xe8, 0x42, 0xcb,
	0xf5, 0xe1, 0x84, 0xf3, 0x89, 0x15, 0x78, 0xc4, 0xd9, 0xd3, 0x07, 0x79, 0xf8, 0xfd, 0x01, 0xaf,
	0x20, 0xd6, 0xd1, 0x8a, 0xeb, 0xb3, 0x85, 0x14, 0x6c, 0x87, 0xc6, 0xc5, 0xc0, 0xcb, 0xcb, 0x3a,
	0xa1, 0xf1, 0x44, 0xec, 0x64, 0x1e, 0x10, 0xfc, 0x6d, 0x60, 0xc7, 0xe7, 0x21, 0xb9, 0xdc, 0x5a,
	0x22, 0x83, 0xcc, 0x93, 0xb7, 0x80, 0x7a, 0xa1, 0x48, 0x01, 0x5d, 0xcd, 0xbb, 0x95, 0x47, 0xb5,
	0x7f, 0x49, 0x3c, 0xb4, 0xa9, 0xcd, 0x6c, 0xa8, 0x23, 0xe0, 0xbc, 0x33, 0x7d, 0x7d, 0x88, 0xaf,
	0xe2, 0xef, 0xf3, 0xea, 0x61, 0x1d, 0x2d, 0x44, 0xe8, 0x1c, 0x80, 0x10, 0x30, 0x2e, 0x04, 0x5e,
	0x4e, 0x94, 0x86, 0x8b, 0x82, 0x5c, 0x0c, 0x16, 0xb7, 0xc7, 0x73, 0x01, 0xbc, 0x68, 0xa1, 0x2c,
	0x82, 0x13, 0x08, 0x5a, 0x41, 0xc1, 0x50, 0xa0, 0x80, 0xae, 0xe4, 0x1d, 0xcc, 0x81, 0x9a, 0xab,
	0xf6, 0x79, 0x24, 0x3a, 0x9c, 0x5d, 0x6a, 0xee, 0xe0, 0x2d, 0x12, 0xb4, 0x74, 0x9d, 0x4f, 0xd9,
	0x2c, 0x90, 0x8f, 0xc1, 0x3b, 0xf4, 0x55, 0x0e, 0xa5, 0xe4, 0x0b, 0xf2, 0xae, 0x87, 0x74, 0xd1,
	0x80, 0xf6, 0x0d, 0x45, 0x1d, 0xc2, 0x01, 0x73, 0xcd, 0xa0, 0xb5, 0xe1, 0xe1, 0x3a, 0xc9, 0x92,
	0xa1, 0x4d, 0xfd, 0x51, 0x3e, 0x90, 0x2b, 0x50, 0x72, 0x81, 0xca, 0x7a, 0xa4, 0x91, 0xe4, 0x11,
	0x2f, 0xa7, 0xd5, 0x89, 0x0c, 0x14, 0x87, 0x6f, 0x52, 0xcc, 0x0c, 0x27, 0x26, 0x91, 0xd4, 0x9a,
	0xd6, 0x54, 0x87, 0x12, 0x0e, 0xcc, 0x35, 0x5b, 0x1e, 0x4c, 0x31, 0x3f, 0x8b, 0x7d, 0xfd, 0x32,
	0x1f, 0x80, 0x5b, 0x40, 0x24, 0x56, 0x59, 0x73, 0x57, 0x3c, 0x82, 0x62, 0xbc, 0x13, 0x1a, 0x97,
	0xa3, 0x29, 0x94, 0x80, 0x15, 0x24, 0x6d, 0xa3, 0x6d, 0xab, 0xda, 0x16, 0x21, 0x2d, 0x93, 0x91,
	0x66, 0xcb, 0xf5, 0xb0, 0x67, 0x13, 0xdf, 0xdc, 0xd4, 0xaf, 0x70, 0x97, 0x5f, 0x86, 0x8d, 0x00,
	0xe8, 0x5a, 0x06, 0x82, 0xbb, 0xd7, 0x78, 0x2f, 0x45, 0x40, 0xac, 0xc5, 0x9e, 0x17, 0x5d, 0x9d,
	0x7c, 0x1e, 0x95, 0xac, 0x68, 0x7b, 0x6a, 0xbf, 0x85, 0xad, 0x4d, 0x62, 0xda, 0x1b, 0xd4, 0xf5,
	0x48, 0xdd, 0x6c, 0xd8, 0x0e, 0xf1, 0xf5, 0xab, 0xdc, 0xc5, 0x05, 0x38, 0xd1, 0x38, 0xbc, 0x10,
	0xa1, 0xf3, 0x00, 0xa6, 0x03, 0x5d, 0x42, 0x4a, 0x7b, 0x30, 0xdd, 0x5b, 0xa8, 0x6c, 0x46, 0xfb,
	0x8e, 0xa2, 0x5e, 0x6e, 0x79, 0xee, 0x06, 0x14, 0x33, 0x66, 0xd0, 0xaa, 0x63, 0x46, 0xc4, 0x02,
	0xe1, 0x31, 0xee, 0xfb, 0x1a, 0xe4, 0xb7, 0x89, 0xd6, 0x3a, 0x57, 0x12, 0x8b, 0x81, 0xa8, 0xc8,
	0xee, 0x82, 0x0b, 0x74, 0x5e, 0x10, 0x06, 0x42, 0x79, 0x01, 0x75, 0xb3, 0xa8, 0xbd, 0xa5, 0xa8,
	0x83, 0x8e, 0xdd, 0xb4, 0x99, 0x59, 0xc3, 0xb4, 0xbe, 0x63, 0xd7, 0xd9, 0xa6, 0x69, 0x53, 0xd3,
	0xc1, 0x54, 0x1f, 0xe6, 0x43, 0xb2, 0xc4, 0x8b, 0x47, 0xd0, 0x98, 0x49, 0x14, 0x16, 0xe8, 0x22,
	0xa6, 0x59, 0xc1, 0x5f, 0xc6, 0x3e, 0x66, 0x58, 0x64, 0xa6, 0xb4, 0x37, 0x15, 0x55, 0x6b, 0xda,
	0xd4, 0xdc, 0x74, 0x9b, 0xc4, 0xac, 0xdb, 0xfe, 0x96, 0xd9, 0xf0, 0x08, 0xd1, 0x8d, 0x11, 0x65,
	0xf4, 0xdc, 0x64, 0xcf, 0x8d, 0xe8, 0x66, 0xed, 0xc6, 0xaa, 0xfd, 0x3a, 0x99, 0x79, 0xf1, 0xfd,
	0xd0, 0x38, 0x01, 0x3b, 0xb1, 0x69, 0xd3, 0x97, 0xdd, 0x26, 0x99, 0xb3, 0xfd, 0xad, 0x79, 0x8f,
	0x90, 0x74, 0x75, 0x14, 0xe4, 0xe2, 0x3e, 0x18, 0xb9, 0x0e, 0x44, 0x4e, 0x4d, 0x8c, 0x5c, 0x47,
	0xc5, 0xe6, 0xda, 0x3d, 0x45, 0xed, 0x49, 0xd6, 0x3b, 0x3f, 0x76, 0x46, 0xf8, 0xb1, 0xf3, 0x27,
	0x9e, 0xf2, 0x24, 0x8b, 0x36, 0x3a, 0x7c, 0xce, 0x79, 0xd9, 0x67, 0x27, 0x34, 0xe6, 0x92, 0x8a,
	0x23, 0x91, 0x49, 0x0e, 0xa2, 0x78, 0x07, 0xf8, 0x85, 0x33, 0xa5, 0x49, 0x18, 0xbe, 0xf1, 0x45,
	0xdf, 0xa5, 0x10, 0xbb, 0x73, 0x66, 0xf3, 0x9f, 0xc7, 0x07, 0xd5, 0xd1, 0x07, 0x35, 0x05, 0xf9,
	0x91, 0xc0, 0x17, 0x65, 0x76, 0x3c, 0x47, 0x7b, 0x55, 0xed, 0xc3, 0xce, 0x0e, 0x54, 0x5f, 0xd1,
	0x6d, 0x02, 0x25, 0xcc, 0xd7, 0x1f, 0xe7, 0x97, 0x78, 0x50, 0xf4, 0x5e, 0x88, 0x40, 0x5e, 0x95,
	0x2f, 0x13, 0x06, 0x0b, 0x7f, 0x20, 0x8a, 0x30, 0x39, 0x79, 0x05, 0x15, 0x15, 0xb5, 0xff, 0x29,
	0xea, 0xa8, 0xbb, 0x4d, 0xbc, 0x1d, 0xcf, 0x66, 0x10, 0x38, 0x9a, 0x2e, 0x23, 0x66, 0x9d, 0x6c,
	0xdb, 0x16, 0x31, 0x29, 0x6e, 0x12, 0x1f, 0xc2, 0x69, 0x5c, 0x08, 0xe9, 0x95, 0xec, 0x7a, 0x69,
	0xe8, 0x4e, 0xd2, 0x08, 0xf1, 0x36, 0x73, 0x64, 0x7b, 0x19, 0xd4, 0xdb, 0xa1, 0x71, 0xcd, 0x2d,
	0x41, 0xb6, 0x45, 0x38, 0x7a, 0x87, 0xce, 0x46, 0xa6, 0x3a, 0xa1, 0xf1, 0x49, 0x4e, 0xf0, 0x01,
	0x74, 0xbb, 0x2f, 0x4a, 0xa8, 0xe2, 0xba, 0xf0, 0x40, 0x0f, 0xc2, 0x42, 0xfb, 0x8a, 0x7a, 0x09,
	0xc2, 0x98, 0x69, 0xd3, 0x3a, 0xd9, 0x35, 0x61, 0x25, 0xd7, 0x1c, 0xd7, 0xda, 0xf2, 0xf5, 0x6b,
	0x7c, 0x4b, 0xc3, 0xa2, 0xd1, 0x40, 0x61, 0x01, 0xf0, 0x25, 0x9b, 0xce, 0x70, 0x34, 0xbd, 0xb5,
	0x2d, 0x43, 0xd2, 0x4c, 0x39, 0xca, 0x7f, 0x91, 0xc4, 0x92, 0xf6, 0x4f, 0x48, 0x77, 0x29, 0xb6,
	0xb6, 0x48, 0xdd, 0xa4, 0x2e, 0xb3, 0x1b, 0xb6, 0x85, 0xa3, 0xfb, 0x87, 0xba, 0xaf, 0x57, 0xf9,
	0xfc, 0xbe, 0x0d, 0xc3, 0x3d, 0xb8, 0x1e, 0x29, 0x2d, 0x0b, 0x3a, 0x0b, 0x73, 0x30, 0xda, 0x83,
	0x81, 0x14, 0xe9, 0x84, 0xc6, 0x95, 0x28, 0xb4, 0xcb, 0x60, 0x7e, 0x57, 0x29, 0x45, 0x3a, 0x07,
	0xd5, 0x2e, 0x16, 0xf7, 0x0f, 0xab, 0x5d, 0x58, 0x20, 0x69, 0x8b, 0xba, 0xaf, 0x21, 0xf5, 0x3c,
	0xf3, 0x70, 0xa3, 0x61, 0x5b, 0xa6, 0xe5, 0x60, 0xdf, 0xd7, 0xaf, 0xf3, 0x61, 0x7d, 0x0e, 0xea,
	0xe5, 0x18, 0x98, 0x05, 0x79, 0x27, 0x34, 0xb4, 0x68, 0x40, 0x05, 0x61, 0x7a, 0x51, 0x93, 0x53,
	0xd5, 0xee, 0xaa, 0xfd, 0xf1, 0x10, 0x9b, 0x0d, 0xd7, 0xa9, 0x13, 0xcf, 0x6c, 0x61, 0xb6, 0xa9,
	0x3f, 0xc1, 0x77, 0xfd, 0x34, 0x1c, 0x03, 0x31, 0x3c, 0xcf, 0xd1, 0x15, 0xcc, 0x36, 0xd3, 0x10,
	0x53, 0x42, 0x84, 0xe9, 0x7a, 0x03, 0x96, 0x95, 0xf2, 0x06, 0x2a, 0x37, 0xd7, 0xb6, 0xd4, 0x8b,
	0x3e, 0x61, 0xa6, 0xe3, 0xee, 0x98, 0x2d, 0xcf, 0x76, 0x3d, 0x9b, 0xed, 0xe9, 0x4f, 0xf2, 0xad,
	0x00, 0xfd, 0xf5, 0xfa, 0x84, 0x2d, 0xba, 0x3b, 0x2b, 0x31, 0x92, 0x76, 0x96, 0x17, 0x77, 0x4d,
	0x2c, 0x0a, 0xcd, 0xb5, 0x77, 0x14, 0x75, 0xb0, 0x89, 0x77, 0x13, 0xe7, 0x2c, 0x97, 0x5a, 0x81,
	0xe7, 0x11, 0x6a, 0xed, 0xe9, 0xa3, 0x7c, 0xf4, 0x7c, 0x7e, 0xc5, 0x82, 0x77, 0x96, 0xf0, 0x6e,
	0xc4, 0x71, 0x36, 0x53, 0x81, 0x83, 0xbe, 0x29, 0x91, 0xa7, 0x07, 0xbd, 0x0c, 0x4c, 0x06, 0x9a,
	0xdf, 0x89, 0xc8, 0xed, 0x22, 0xa9, 0x55, 0xed, 0x03, 0x45, 0xed, 0xb7, 0x3c, 0xec, 0x6f, 0x16,
	0x32, 0xff, 0xa7, 0xf8, 0x64, 0xbc, 0xcb, 0x33, 0xff, 0xd9, 0x24, 0xf3, 0xb7, 0xe2, 0xcc, 0x7f,
	0x3e, 0x3a, 0x91, 0xa1, 0x59, 0x96, 0x83, 0x4b, 0x83, 0x2f, 0xd7, 0x29, 0x67, 0xf3, 0x5c, 0x0c,
	0x2b, 0xb8, 0xaf, 0x64, 0x04, 0x6a, 0x02, 0x2b, 0xae, 0x09, 0xaa, 0x0f, 0x62, 0x06, 0xaa, 0x82,
	0xd9, 0xa8, 0x2a, 0x28, 0x18, 0xf3, 0x1c, 0xed, 0xa7, 0x8a, 0x3a, 0x54, 0x74, 0x2f, 0xb9, 0x8c,
	0x79, 0x9a, 0xcf, 0xbf, 0x7d, 0x14, 0x1a, 0x67, 0x67, 0x91, 0xf0, 0x8e, 0x90, 0xb7, 0x52, 0x7c,
	0x47, 0x90, 0xa2, 0xdd, 0x96, 0xc6, 0xfe, 0x61, 0x35, 0xb3, 0x8d, 0xe4, 0x96, 0xb5, 0xaf, 0x29,
	0xea, 0xa0, 0xcf, 0x02, 0x6a, 0x42, 0xbe, 0x84, 0x1d, 0x7b, 0x9b, 0x98, 0x51, 0x16, 0xec, 0xeb,
	0xcf, 0xa4, 0x59, 0x68, 0x3f, 0x68, 0xbc, 0x92, 0x28, 0xac, 0x02, 0xbe, 0x9a, 0xe6, 0x46, 0x12,
	0x2c, 0x9f, 0xc2, 0x0b, 0x61, 0xec, 0xd4, 0xc4, 0xed, 0x71, 0x24, 0xb3, 0x06, 0x95, 0x71, 0x81,
	0x06, 0x44, 0x53, 0x5f, 0x7f, 0x96, 0x93, 0xf8, 0x2c, 0xec, 0xcb, 0x5c, 0xb3, 0x25, 0x9b, 0x66,
	0x15, 0x44, 0x09, 0x11, 0x33, 0xc3, 0x5c, 0x18, 0x9d, 0x1c, 0x47, 0x65, 0x3b, 0x90, 0x8b, 0xf7,
	0xf0, 0xde, 0x93, 0xe7, 0xad, 0xe7, 0x78, 0xe4, 0xac, 0x1f, 0x85, 0x46, 0x2f, 0xc2, 0x3b, 0xab,
	0x2c, 0x10, 0x1e, 0xb6, 0xce, 0xf9, 0xd9, 0x67, 0x7a, 0x05, 0x95, 0xc9, 0xee, 0xfb, 0xf8, 0x56,
	0xb0, 0x88, 0x44, 0x7b, 0xda, 0xb6, 0x7a, 0x01, 0x8a, 0xcd, 0x1a, 0xf6, 0x89, 0x19, 0xbd, 0x34,
	0xea, 0x37, 0x46, 0x94, 0xd1, 0xde, 0xc9, 0xde, 0x24, 0x19, 0x5a, 0xe3, 0x52, 0x7e, 0x67, 0xd8,
	0x9b, 0xa8, 0x46, 0xb2, 0x2c, 0x4c, 0xe5, 0xc4, 0x95, 0x91, 0xb8, 0xf4, 0x88, 0x97, 0xc7, 0x9b,
	0x87, 0x55, 0x05, 0x15, 0x9a, 0x6a, 0xdf, 0x3d, 0xa9, 0x5e, 0x83, 0xa8, 0x91, 0x86, 0x0b, 0x28,
	0x5d, 0x2d, 0xb7, 0x09, 0x4b, 0xd6, 0x23, 0x77, 0x03, 0xe2, 0x33, 0x73, 0xcb, 0xae, 0xe9, 0x63,
	0x7c, 0x3a, 0xfe, 0xaa, 0xc4, 0x2f, 0x94, 0x4b, 0x78, 0x77, 0x76, 0x01, 0x45, 0xf8, 0x2b, 0xf6,
	0x4c, 0x3b, 0x34, 0x8c, 0x26, 0xde, 0x4d, 0xb7, 0x38, 0x5b, 0x88, 0x6d, 0x64, 0x2a, 0xe9, 0xd9,
	0x77, 0x1f, 0x3d, 0xa1, 0xec, 0xbb, 0xaf, 0xc9, 0xfb, 0xab, 0xc4, 0x6f, 0x9e, 0x05, 0xba, 0xe8,
	0x3e, 0xcd, 0x6a, 0xda, 0x47, 0x8a, 0x3a, 0x98, 0x3e, 0xbc, 0x38, 0x58, 0x7c, 0xaa, 0x1d, 0xe7,
	0x1b, 0xf8, 0x3d, 0x18, 0x89, 0x81, 0xe4, 0xe1, 0x62, 0x71, 0x7a, 0x59, 0x7c, 0xad, 0x1d, 0xc0,
	0x12, 0x79, 0x9a, 0x3e, 0xcb, 0x40, 0xd9, 0x7b, 0x99, 0xd4, 0x48, 0x17, 0xb9, 0xb0, 0xf5, 0xa5,
	0xa4, 0x50, 0xd6, 0x0a, 0x0b, 0x4f, 0xbd, 0xdb, 0xea, 0x65, 0xfe, 0xb6, 0xd2, 0x08, 0x1c, 0x27,
	0xce, 0x65, 0x5c, 0x9a, 0x14, 0xa6, 0xfa, 0x04, 0xf7, 0x74, 0x0a, 0x72, 0x05, 0xd0, 0x9a, 0x0f,
	0x1c, 0x87, 0x67, 0x21, 0x77, 0x68, 0x5c, 0x4a, 0x76, 0x42, 0xe3, 0x6a, 0x7c, 0x64, 0xc9, 0xe0,
	0x0a, 0xea, 0xd2, 0x4e, 0xfb, 0xb2, 0xda, 0x13, 0xb4, 0x68, 0x2b, 0x0d, 0x8a, 0xbf, 0x9c, 0xe7,
	0x5d, 0x7d, 0xee, 0x28, 0x34, 0x2e, 0xcd, 0x91, 0x96, 0x47, 0x2c, 0xcc, 0x48, 0x7d, 0x7d, 0x85,
	0xae, 0x64, 0x11, 0x52, 0x79, 0x2e, 0x4b, 0x4d, 0x5a, 0xb4, 0x15, 0x03, 0xcf, 0xba, 0x4d, 0x1b,
	0xd2, 0x23, 0xb6, 0x57, 0xd9, 0x3f, 0xac, 0xca, 0x1b, 0xeb, 0x0a, 0x3a, 0x27, 0x34, 0xd1, 0x7e,
	0xae, 0xc4, 0xdd, 0x27, 0x17, 0xcf, 0xef, 0xcc, 0xf3, 0xd5, 0xfd, 0x26, 0x9f, 0xd3, 0xbc, 0x89,
	0xf4, 0x12, 0x9a, 0x77, 0x3f, 0x92, 0x76, 0x2f, 0x5e, 0x1e, 0x0b, 0x1c, 0xb2, 0xc5, 0x7b, 0xb9,
	0xbb, 0x16, 0x4c, 0x92, 0xac, 0x17, 0x5d, 0x41, 0x6a, 0xd6, 0x4a, 0xfb, 0xad, 0xa2, 0xf6, 0x72,
	0x9a, 0xd9, 0x15, 0xf3, 0xaf, 0x22, 0xa2, 0xdf, 0xe4, 0x99, 0x5d, 0xde, 0x84, 0x70, 0xdd, 0xcc,
	0xa9, 0x56, 0x52, 0xaa, 0xf9, 0x0b, 0x62, 0x29, 0xd9, 0xab, 0x1f, 0xa7, 0x07, 0xf9, 0x9b, 0xbc,
	0x2f, 0x5d, 0x41, 0x3d, 0x62, 0xcb, 0x8c, 0x72, 0x76, 0x91, 0xfc, 0x6e, 0x77, 0xca, 0xc2, 0xa5,
	0x72, 0x81, 0x72, 0xfe, 0x1a, 0xb8, 0x3b, 0xe5, 0x6e, 0x7a, 0x65, 0xca, 0x89, 0x66, 0x42, 0x39,
	0xbd, 0x37, 0x6e, 0xa8, 0xd1, 0x83, 0x55, 0x7a, 0x04, 0xfc, 0x7a, 0x9e, 0x9f, 0x01, 0x9f, 0xc9,
	0xf3, 0xe5, 0x6f, 0x3e, 0xd9, 0x59, 0x20, 0x2c, 0x46, 0x2f, 0x43, 0x04, 0xa2, 0xd0, 0x8f, 0x80,
	0xf8, 0xbc, 0xec, 0x2e, 0x57, 0xbc, 0x66, 0xcb, 0x62, 0xfa, 0x7b, 0x30, 0x44, 0xca, 0xcc, 0xd2,
	0x51, 0x68, 0x5c, 0xcd, 0x7a, 0x5c, 0xca, 0xd7, 0xab, 0x2b, 0x16, 0xcb, 0x8f, 0x53, 0xb3, 0x84,
	0xe7, 0xbb, 0xd7, 0xca, 0x0a, 0x70, 0xde, 0x0d, 0x14, 0xa2, 0xbd, 0x6f, 0x61, 0xea, 0xeb, 0xbf,
	0x89, 0x66, 0x69, 0xad, 0x40, 0x41, 0x8c, 0x92, 0xab, 0xa0, 0x58, 0xa0, 0x50, 0xc2, 0xcb, 0x53,
	0xc5, 0x99, 0x94, 0xf4, 0x66, 0x5e, 0x7a, 0xff, 0xc3, 0xe1, 0x13, 0x87, 0x1f, 0x0e, 0x9f, 0xf8,
	0xcf, 0xd1, 0xf0, 0x89, 0x6f, 0xdf, 0x1b, 0x3e, 0xf1, 0xf6, 0xbd, 0x61, 0xe5, 0xf0, 0xde, 0xf0,
	0x89, 0x7f, 0xdc, 0x1b, 0x3e, 0xf1, 0xda, 0x53, 0x1b, 0x36, 0xdb, 0x0c, 0x6a, 0x37, 0x2c, 0xb7,
	0x39, 0x96, 0xe6, 0x5f, 0xc2, 0xaf, 0xec, 0xdf, 0x37, 0xb5, 0x33, 0xfc, 0xef, 0x36, 0x37, 0xff,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x9c, 0x92, 0x45, 0xda, 0x23, 0x00, 0x00,
}

func (m *OptionsConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RawListenAddresses) > 0 {
		for _, s := range m.RawListenAddresses {
			l = len(s)
			n += 1 + l + sovOptionsconfiguration(uint64(l))
		}
	}
	if len(m.RawGlobalAnnServers) > 0 {
		for _, s := range m.RawGlobalAnnServers {
			l = len(s)
			n += 1 + l + sovOptionsconfiguration(uint64(l))
		}
	}
	if m.GlobalAnnEnabled {
		n += 2
	}
	if m.LocalAnnEnabled {
		n += 2
	}
	if m.LocalAnnPort != 0 {
		n += 1 + sovOptionsconfiguration(uint64(m.LocalAnnPort))
	}
	l = len(m.LocalAnnMCAddr)
	if l > 0 {
		n += 1 + l + sovOptionsconfiguration(uint64(l))
	}
	if m.MaxSendKbps != 0 {
		n += 1 + sovOptionsconfiguration(uint64(m.MaxSendKbps))
	}
	if m.MaxRecvKbps != 0 {
		n += 1 + sovOptionsconfiguration(uint64(m.MaxRecvKbps))
	}
	if m.ReconnectIntervalS != 0 {
		n += 1 + sovOptionsconfiguration(uint64(m.ReconnectIntervalS))
	}
	if m.RelaysEnabled {
		n += 2
	}
	if m.RelayReconnectIntervalM != 0 {
		n += 1 + sovOptionsconfiguration(uint64(m.RelayReconnectIntervalM))
	}
	if m.StartBrowser {
		n += 2
	}
	if m.NATEnabled {
		n += 2
	}
	if m.NATLeaseM != 0 {
		n += 1 + sovOptionsconfiguration(uint64(m.NATLeaseM))
	}
	if m.NATRenewalM != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.NATRenewalM))
	}
	if m.NATTimeoutS != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.NATTimeoutS))
	}
	if m.URAccepted != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.URAccepted))
	}
	if m.URSeen != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.URSeen))
	}
	l = len(m.URUniqueID)
	if l > 0 {
		n += 2 + l + sovOptionsconfiguration(uint64(l))
	}
	l = len(m.URURL)
	if l > 0 {
		n += 2 + l + sovOptionsconfiguration(uint64(l))
	}
	if m.URPostInsecurely {
		n += 3
	}
	if m.URInitialDelayS != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.URInitialDelayS))
	}
	if m.RestartOnWakeup {
		n += 3
	}
	if m.AutoUpgradeIntervalH != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.AutoUpgradeIntervalH))
	}
	if m.UpgradeToPreReleases {
		n += 3
	}
	if m.KeepTemporariesH != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.KeepTemporariesH))
	}
	if m.CacheIgnoredFiles {
		n += 3
	}
	if m.ProgressUpdateIntervalS != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.ProgressUpdateIntervalS))
	}
	if m.LimitBandwidthInLan {
		n += 3
	}
	l = m.MinHomeDiskFree.ProtoSize()
	n += 2 + l + sovOptionsconfiguration(uint64(l))
	l = len(m.ReleasesURL)
	if l > 0 {
		n += 2 + l + sovOptionsconfiguration(uint64(l))
	}
	if len(m.AlwaysLocalNets) > 0 {
		for _, s := range m.AlwaysLocalNets {
			l = len(s)
			n += 2 + l + sovOptionsconfiguration(uint64(l))
		}
	}
	if m.OverwriteRemoteDevNames {
		n += 3
	}
	if m.TempIndexMinBlocks != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.TempIndexMinBlocks))
	}
	if len(m.UnackedNotificationIDs) > 0 {
		for _, s := range m.UnackedNotificationIDs {
			l = len(s)
			n += 2 + l + sovOptionsconfiguration(uint64(l))
		}
	}
	if m.TrafficClass != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.TrafficClass))
	}
	l = len(m.DefaultFolderPath)
	if l > 0 {
		n += 2 + l + sovOptionsconfiguration(uint64(l))
	}
	if m.SetLowPriority {
		n += 3
	}
	if m.RawMaxFolderConcurrency != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.RawMaxFolderConcurrency))
	}
	l = len(m.CRURL)
	if l > 0 {
		n += 2 + l + sovOptionsconfiguration(uint64(l))
	}
	if m.CREnabled {
		n += 3
	}
	if m.StunKeepaliveStartS != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.StunKeepaliveStartS))
	}
	if m.StunKeepaliveMinS != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.StunKeepaliveMinS))
	}
	if len(m.RawStunServers) > 0 {
		for _, s := range m.RawStunServers {
			l = len(s)
			n += 2 + l + sovOptionsconfiguration(uint64(l))
		}
	}
	if m.DatabaseTuning != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.DatabaseTuning))
	}
	if m.RawMaxCIRequestKiB != 0 {
		n += 2 + sovOptionsconfiguration(uint64(m.RawMaxCIRequestKiB))
	}
	if m.AnnounceLANAddresses {
		n += 3
	}
	if m.SendFullIndexOnUpgrade {
		n += 3
	}
	if m.DeprecatedUPnPEnabled {
		n += 4
	}
	if m.DeprecatedUPnPLeaseM != 0 {
		n += 3 + sovOptionsconfiguration(uint64(m.DeprecatedUPnPLeaseM))
	}
	if m.DeprecatedUPnPRenewalM != 0 {
		n += 3 + sovOptionsconfiguration(uint64(m.DeprecatedUPnPRenewalM))
	}
	if m.DeprecatedUPnPTimeoutS != 0 {
		n += 3 + sovOptionsconfiguration(uint64(m.DeprecatedUPnPTimeoutS))
	}
	if len(m.DeprecatedRelayServers) > 0 {
		for _, s := range m.DeprecatedRelayServers {
			l = len(s)
			n += 3 + l + sovOptionsconfiguration(uint64(l))
		}
	}
	if m.DeprecatedMinHomeDiskFreePct != 0 {
		n += 11
	}
	if m.DeprecatedMaxConcurrentScans != 0 {
		n += 3 + sovOptionsconfiguration(uint64(m.DeprecatedMaxConcurrentScans))
	}
	return n
}

func sovOptionsconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOptionsconfiguration(x uint64) (n int) {
	return sovOptionsconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
