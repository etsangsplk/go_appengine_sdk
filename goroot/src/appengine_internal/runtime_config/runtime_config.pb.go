// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine_tools_devappserver2 is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/runtime_config

It has these top-level messages:
	Config
	PhpConfig
	PythonConfig
	JavaConfig
	CustomConfig
	CloudSQL
	Library
	Environ
	VMConfig
*/
package appengine_tools_devappserver2

import proto "appengine_internal/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	AppId            []byte        `protobuf:"bytes,1,req,name=app_id,json=appId" json:"app_id,omitempty"`
	VersionId        []byte        `protobuf:"bytes,2,req,name=version_id,json=versionId" json:"version_id,omitempty"`
	ApplicationRoot  []byte        `protobuf:"bytes,3,req,name=application_root,json=applicationRoot" json:"application_root,omitempty"`
	Threadsafe       *bool         `protobuf:"varint,4,opt,name=threadsafe,def=0" json:"threadsafe,omitempty"`
	ApiHost          *string       `protobuf:"bytes,17,opt,name=api_host,json=apiHost,def=localhost" json:"api_host,omitempty"`
	ApiPort          *int32        `protobuf:"varint,5,req,name=api_port,json=apiPort" json:"api_port,omitempty"`
	Libraries        []*Library    `protobuf:"bytes,6,rep,name=libraries" json:"libraries,omitempty"`
	SkipFiles        *string       `protobuf:"bytes,7,opt,name=skip_files,json=skipFiles,def=^$" json:"skip_files,omitempty"`
	StaticFiles      *string       `protobuf:"bytes,8,opt,name=static_files,json=staticFiles,def=^$" json:"static_files,omitempty"`
	PythonConfig     *PythonConfig `protobuf:"bytes,14,opt,name=python_config,json=pythonConfig" json:"python_config,omitempty"`
	PhpConfig        *PhpConfig    `protobuf:"bytes,9,opt,name=php_config,json=phpConfig" json:"php_config,omitempty"`
	JavaConfig       *JavaConfig   `protobuf:"bytes,21,opt,name=java_config,json=javaConfig" json:"java_config,omitempty"`
	CustomConfig     *CustomConfig `protobuf:"bytes,23,opt,name=custom_config,json=customConfig" json:"custom_config,omitempty"`
	Environ          []*Environ    `protobuf:"bytes,10,rep,name=environ" json:"environ,omitempty"`
	CloudSqlConfig   *CloudSQL     `protobuf:"bytes,11,opt,name=cloud_sql_config,json=cloudSqlConfig" json:"cloud_sql_config,omitempty"`
	Datacenter       *string       `protobuf:"bytes,12,req,name=datacenter" json:"datacenter,omitempty"`
	InstanceId       *string       `protobuf:"bytes,13,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	StderrLogLevel   *int64        `protobuf:"varint,15,opt,name=stderr_log_level,json=stderrLogLevel,def=1" json:"stderr_log_level,omitempty"`
	AuthDomain       *string       `protobuf:"bytes,16,req,name=auth_domain,json=authDomain" json:"auth_domain,omitempty"`
	MaxInstances     *int32        `protobuf:"varint,18,opt,name=max_instances,json=maxInstances" json:"max_instances,omitempty"`
	VmConfig         *VMConfig     `protobuf:"bytes,19,opt,name=vm_config,json=vmConfig" json:"vm_config,omitempty"`
	ServerPort       *int32        `protobuf:"varint,20,opt,name=server_port,json=serverPort" json:"server_port,omitempty"`
	Vm               *bool         `protobuf:"varint,22,opt,name=vm,def=0" json:"vm,omitempty"`
	GrpcApis         []string      `protobuf:"bytes,24,rep,name=grpc_apis,json=grpcApis" json:"grpc_apis,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}

const Default_Config_Threadsafe bool = false
const Default_Config_ApiHost string = "localhost"
const Default_Config_SkipFiles string = "^$"
const Default_Config_StaticFiles string = "^$"
const Default_Config_StderrLogLevel int64 = 1
const Default_Config_Vm bool = false

func (m *Config) GetAppId() []byte {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *Config) GetVersionId() []byte {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *Config) GetApplicationRoot() []byte {
	if m != nil {
		return m.ApplicationRoot
	}
	return nil
}

func (m *Config) GetThreadsafe() bool {
	if m != nil && m.Threadsafe != nil {
		return *m.Threadsafe
	}
	return Default_Config_Threadsafe
}

func (m *Config) GetApiHost() string {
	if m != nil && m.ApiHost != nil {
		return *m.ApiHost
	}
	return Default_Config_ApiHost
}

func (m *Config) GetApiPort() int32 {
	if m != nil && m.ApiPort != nil {
		return *m.ApiPort
	}
	return 0
}

func (m *Config) GetLibraries() []*Library {
	if m != nil {
		return m.Libraries
	}
	return nil
}

func (m *Config) GetSkipFiles() string {
	if m != nil && m.SkipFiles != nil {
		return *m.SkipFiles
	}
	return Default_Config_SkipFiles
}

func (m *Config) GetStaticFiles() string {
	if m != nil && m.StaticFiles != nil {
		return *m.StaticFiles
	}
	return Default_Config_StaticFiles
}

func (m *Config) GetPythonConfig() *PythonConfig {
	if m != nil {
		return m.PythonConfig
	}
	return nil
}

func (m *Config) GetPhpConfig() *PhpConfig {
	if m != nil {
		return m.PhpConfig
	}
	return nil
}

func (m *Config) GetJavaConfig() *JavaConfig {
	if m != nil {
		return m.JavaConfig
	}
	return nil
}

func (m *Config) GetCustomConfig() *CustomConfig {
	if m != nil {
		return m.CustomConfig
	}
	return nil
}

func (m *Config) GetEnviron() []*Environ {
	if m != nil {
		return m.Environ
	}
	return nil
}

func (m *Config) GetCloudSqlConfig() *CloudSQL {
	if m != nil {
		return m.CloudSqlConfig
	}
	return nil
}

func (m *Config) GetDatacenter() string {
	if m != nil && m.Datacenter != nil {
		return *m.Datacenter
	}
	return ""
}

func (m *Config) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *Config) GetStderrLogLevel() int64 {
	if m != nil && m.StderrLogLevel != nil {
		return *m.StderrLogLevel
	}
	return Default_Config_StderrLogLevel
}

func (m *Config) GetAuthDomain() string {
	if m != nil && m.AuthDomain != nil {
		return *m.AuthDomain
	}
	return ""
}

func (m *Config) GetMaxInstances() int32 {
	if m != nil && m.MaxInstances != nil {
		return *m.MaxInstances
	}
	return 0
}

func (m *Config) GetVmConfig() *VMConfig {
	if m != nil {
		return m.VmConfig
	}
	return nil
}

func (m *Config) GetServerPort() int32 {
	if m != nil && m.ServerPort != nil {
		return *m.ServerPort
	}
	return 0
}

func (m *Config) GetVm() bool {
	if m != nil && m.Vm != nil {
		return *m.Vm
	}
	return Default_Config_Vm
}

func (m *Config) GetGrpcApis() []string {
	if m != nil {
		return m.GrpcApis
	}
	return nil
}

type PhpConfig struct {
	PhpExecutablePath   []byte `protobuf:"bytes,1,opt,name=php_executable_path,json=phpExecutablePath" json:"php_executable_path,omitempty"`
	EnableDebugger      *bool  `protobuf:"varint,3,req,name=enable_debugger,json=enableDebugger" json:"enable_debugger,omitempty"`
	GaeExtensionPath    []byte `protobuf:"bytes,4,opt,name=gae_extension_path,json=gaeExtensionPath" json:"gae_extension_path,omitempty"`
	XdebugExtensionPath []byte `protobuf:"bytes,5,opt,name=xdebug_extension_path,json=xdebugExtensionPath" json:"xdebug_extension_path,omitempty"`
	XXX_unrecognized    []byte `json:"-"`
}

func (m *PhpConfig) Reset()         { *m = PhpConfig{} }
func (m *PhpConfig) String() string { return proto.CompactTextString(m) }
func (*PhpConfig) ProtoMessage()    {}

func (m *PhpConfig) GetPhpExecutablePath() []byte {
	if m != nil {
		return m.PhpExecutablePath
	}
	return nil
}

func (m *PhpConfig) GetEnableDebugger() bool {
	if m != nil && m.EnableDebugger != nil {
		return *m.EnableDebugger
	}
	return false
}

func (m *PhpConfig) GetGaeExtensionPath() []byte {
	if m != nil {
		return m.GaeExtensionPath
	}
	return nil
}

func (m *PhpConfig) GetXdebugExtensionPath() []byte {
	if m != nil {
		return m.XdebugExtensionPath
	}
	return nil
}

type PythonConfig struct {
	StartupScript    *string `protobuf:"bytes,1,opt,name=startup_script,json=startupScript" json:"startup_script,omitempty"`
	StartupArgs      *string `protobuf:"bytes,2,opt,name=startup_args,json=startupArgs" json:"startup_args,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PythonConfig) Reset()         { *m = PythonConfig{} }
func (m *PythonConfig) String() string { return proto.CompactTextString(m) }
func (*PythonConfig) ProtoMessage()    {}

func (m *PythonConfig) GetStartupScript() string {
	if m != nil && m.StartupScript != nil {
		return *m.StartupScript
	}
	return ""
}

func (m *PythonConfig) GetStartupArgs() string {
	if m != nil && m.StartupArgs != nil {
		return *m.StartupArgs
	}
	return ""
}

type JavaConfig struct {
	JvmArgs          []string `protobuf:"bytes,1,rep,name=jvm_args,json=jvmArgs" json:"jvm_args,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *JavaConfig) Reset()         { *m = JavaConfig{} }
func (m *JavaConfig) String() string { return proto.CompactTextString(m) }
func (*JavaConfig) ProtoMessage()    {}

func (m *JavaConfig) GetJvmArgs() []string {
	if m != nil {
		return m.JvmArgs
	}
	return nil
}

type CustomConfig struct {
	CustomEntrypoint *string `protobuf:"bytes,1,opt,name=custom_entrypoint,json=customEntrypoint" json:"custom_entrypoint,omitempty"`
	Runtime          *string `protobuf:"bytes,2,opt,name=runtime" json:"runtime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CustomConfig) Reset()         { *m = CustomConfig{} }
func (m *CustomConfig) String() string { return proto.CompactTextString(m) }
func (*CustomConfig) ProtoMessage()    {}

func (m *CustomConfig) GetCustomEntrypoint() string {
	if m != nil && m.CustomEntrypoint != nil {
		return *m.CustomEntrypoint
	}
	return ""
}

func (m *CustomConfig) GetRuntime() string {
	if m != nil && m.Runtime != nil {
		return *m.Runtime
	}
	return ""
}

type CloudSQL struct {
	MysqlHost        *string `protobuf:"bytes,1,req,name=mysql_host,json=mysqlHost" json:"mysql_host,omitempty"`
	MysqlPort        *int32  `protobuf:"varint,2,req,name=mysql_port,json=mysqlPort" json:"mysql_port,omitempty"`
	MysqlUser        *string `protobuf:"bytes,3,req,name=mysql_user,json=mysqlUser" json:"mysql_user,omitempty"`
	MysqlPassword    *string `protobuf:"bytes,4,req,name=mysql_password,json=mysqlPassword" json:"mysql_password,omitempty"`
	MysqlSocket      *string `protobuf:"bytes,5,opt,name=mysql_socket,json=mysqlSocket" json:"mysql_socket,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CloudSQL) Reset()         { *m = CloudSQL{} }
func (m *CloudSQL) String() string { return proto.CompactTextString(m) }
func (*CloudSQL) ProtoMessage()    {}

func (m *CloudSQL) GetMysqlHost() string {
	if m != nil && m.MysqlHost != nil {
		return *m.MysqlHost
	}
	return ""
}

func (m *CloudSQL) GetMysqlPort() int32 {
	if m != nil && m.MysqlPort != nil {
		return *m.MysqlPort
	}
	return 0
}

func (m *CloudSQL) GetMysqlUser() string {
	if m != nil && m.MysqlUser != nil {
		return *m.MysqlUser
	}
	return ""
}

func (m *CloudSQL) GetMysqlPassword() string {
	if m != nil && m.MysqlPassword != nil {
		return *m.MysqlPassword
	}
	return ""
}

func (m *CloudSQL) GetMysqlSocket() string {
	if m != nil && m.MysqlSocket != nil {
		return *m.MysqlSocket
	}
	return ""
}

type Library struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Version          *string `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Library) Reset()         { *m = Library{} }
func (m *Library) String() string { return proto.CompactTextString(m) }
func (*Library) ProtoMessage()    {}

func (m *Library) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Library) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type Environ struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Environ) Reset()         { *m = Environ{} }
func (m *Environ) String() string { return proto.CompactTextString(m) }
func (*Environ) ProtoMessage()    {}

func (m *Environ) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Environ) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type VMConfig struct {
	DockerDaemonUrl  *string `protobuf:"bytes,1,opt,name=docker_daemon_url,json=dockerDaemonUrl" json:"docker_daemon_url,omitempty"`
	EnableLogs       *bool   `protobuf:"varint,3,opt,name=enable_logs,json=enableLogs" json:"enable_logs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VMConfig) Reset()         { *m = VMConfig{} }
func (m *VMConfig) String() string { return proto.CompactTextString(m) }
func (*VMConfig) ProtoMessage()    {}

func (m *VMConfig) GetDockerDaemonUrl() string {
	if m != nil && m.DockerDaemonUrl != nil {
		return *m.DockerDaemonUrl
	}
	return ""
}

func (m *VMConfig) GetEnableLogs() bool {
	if m != nil && m.EnableLogs != nil {
		return *m.EnableLogs
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "appengine.tools.devappserver2.Config")
	proto.RegisterType((*PhpConfig)(nil), "appengine.tools.devappserver2.PhpConfig")
	proto.RegisterType((*PythonConfig)(nil), "appengine.tools.devappserver2.PythonConfig")
	proto.RegisterType((*JavaConfig)(nil), "appengine.tools.devappserver2.JavaConfig")
	proto.RegisterType((*CustomConfig)(nil), "appengine.tools.devappserver2.CustomConfig")
	proto.RegisterType((*CloudSQL)(nil), "appengine.tools.devappserver2.CloudSQL")
	proto.RegisterType((*Library)(nil), "appengine.tools.devappserver2.Library")
	proto.RegisterType((*Environ)(nil), "appengine.tools.devappserver2.Environ")
	proto.RegisterType((*VMConfig)(nil), "appengine.tools.devappserver2.VMConfig")
}

func init() {
}
