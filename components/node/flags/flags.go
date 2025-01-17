package flags

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/components/node/chaincfg"
	"github.com/wemixkanvas/kanvas/components/node/sources"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
)

// Flags

const envVarPrefix = "NODE_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */
	L1NodeAddr = cli.StringFlag{
		Name:   "l1",
		Usage:  "Address of L1 User JSON-RPC endpoint to use (eth namespace required)",
		Value:  "http://127.0.0.1:8545",
		EnvVar: prefixEnvVar("L1_ETH_RPC"),
	}
	L2EngineAddr = cli.StringFlag{
		Name:   "l2",
		Usage:  "Address of L2 Engine JSON-RPC endpoints to use (engine and eth namespace required)",
		EnvVar: prefixEnvVar("L2_ENGINE_RPC"),
	}
	RollupConfig = cli.StringFlag{
		Name:   "rollup.config",
		Usage:  "Rollup chain parameters",
		EnvVar: prefixEnvVar("ROLLUP_CONFIG"),
	}
	Network = cli.StringFlag{
		Name:   "network",
		Usage:  fmt.Sprintf("Predefined network selection. Available networks: %s", strings.Join(chaincfg.AvailableNetworks(), ", ")),
		EnvVar: prefixEnvVar("NETWORK"),
	}
	RPCListenAddr = cli.StringFlag{
		Name:   "rpc.addr",
		Usage:  "RPC listening address",
		EnvVar: prefixEnvVar("RPC_ADDR"),
	}
	RPCListenPort = cli.IntFlag{
		Name:   "rpc.port",
		Usage:  "RPC listening port",
		EnvVar: prefixEnvVar("RPC_PORT"),
	}
	RPCEnableAdmin = cli.BoolFlag{
		Name:   "rpc.enable-admin",
		Usage:  "Enable the admin API (experimental)",
		EnvVar: prefixEnvVar("RPC_ENABLE_ADMIN"),
	}

	/* Optional Flags */
	L1TrustRPC = cli.BoolFlag{
		Name:   "l1.trustrpc",
		Usage:  "Trust the L1 RPC, sync faster at risk of malicious/buggy RPC providing bad or inconsistent L1 data",
		EnvVar: prefixEnvVar("L1_TRUST_RPC"),
	}
	L1RPCProviderKind = cli.GenericFlag{
		Name: "l1.rpckind",
		Usage: "The kind of RPC provider, used to inform optimal transactions receipts fetching, and thus reduce costs. Valid options: " +
			EnumString[sources.RPCProviderKind](sources.RPCProviderKinds),
		EnvVar: prefixEnvVar("L1_RPC_KIND"),
		Value: func() *sources.RPCProviderKind {
			out := sources.RPCKindBasic
			return &out
		}(),
	}
	L2EngineJWTSecret = cli.StringFlag{
		Name:        "l2.jwt-secret",
		Usage:       "Path to JWT secret key. Keys are 32 bytes, hex encoded in a file. A new key will be generated if left empty.",
		EnvVar:      prefixEnvVar("L2_ENGINE_AUTH"),
		Required:    false,
		Value:       "",
		Destination: new(string),
	}
	SyncerL1Confs = cli.Uint64Flag{
		Name:     "syncer.l1-confs",
		Usage:    "Number of L1 blocks to keep distance from the L1 head before deriving L2 data from. Reorgs are supported, but may be slow to perform.",
		EnvVar:   prefixEnvVar("SYNCER_L1_CONFS"),
		Required: false,
		Value:    0,
	}
	ProposerEnabledFlag = cli.BoolFlag{
		Name:   "proposer.enabled",
		Usage:  "Enable proposing of new L2 blocks. A separate batch submitter has to be deployed to publish the data for syncers.",
		EnvVar: prefixEnvVar("PROPOSER_ENABLED"),
	}
	ProposerStoppedFlag = cli.BoolFlag{
		Name:   "proposer.stopped",
		Usage:  "Initialize the proposer in a stopped state. The proposer can be started using the admin_startProposer RPC",
		EnvVar: prefixEnvVar("PROPOSER_STOPPED"),
	}
	ProposerL1Confs = cli.Uint64Flag{
		Name:     "proposer.l1-confs",
		Usage:    "Number of L1 blocks to keep distance from the L1 head as a proposer for picking an L1 origin.",
		EnvVar:   prefixEnvVar("PROPOSER_L1_CONFS"),
		Required: false,
		Value:    4,
	}
	L1EpochPollIntervalFlag = cli.DurationFlag{
		Name:     "l1.epoch-poll-interval",
		Usage:    "Poll interval for retrieving new L1 epoch updates such as safe and finalized block changes. Disabled if 0 or negative.",
		EnvVar:   prefixEnvVar("L1_EPOCH_POLL_INTERVAL"),
		Required: false,
		Value:    time.Second * 12 * 32,
	}
	MetricsEnabledFlag = cli.BoolFlag{
		Name:   "metrics.enabled",
		Usage:  "Enable the metrics server",
		EnvVar: prefixEnvVar("METRICS_ENABLED"),
	}
	MetricsAddrFlag = cli.StringFlag{
		Name:   "metrics.addr",
		Usage:  "Metrics listening address",
		Value:  "0.0.0.0",
		EnvVar: prefixEnvVar("METRICS_ADDR"),
	}
	MetricsPortFlag = cli.IntFlag{
		Name:   "metrics.port",
		Usage:  "Metrics listening port",
		Value:  7300,
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}
	PprofEnabledFlag = cli.BoolFlag{
		Name:   "pprof.enabled",
		Usage:  "Enable the pprof server",
		EnvVar: prefixEnvVar("PPROF_ENABLED"),
	}
	PprofAddrFlag = cli.StringFlag{
		Name:   "pprof.addr",
		Usage:  "pprof listening address",
		Value:  "0.0.0.0",
		EnvVar: prefixEnvVar("PPROF_ADDR"),
	}
	PprofPortFlag = cli.IntFlag{
		Name:   "pprof.port",
		Usage:  "pprof listening port",
		Value:  6060,
		EnvVar: prefixEnvVar("PPROF_PORT"),
	}
	SnapshotLog = cli.StringFlag{
		Name:   "snapshotlog.file",
		Usage:  "Path to the snapshot log file",
		EnvVar: prefixEnvVar("SNAPSHOT_LOG"),
	}
	HeartbeatEnabledFlag = cli.BoolFlag{
		Name:   "heartbeat.enabled",
		Usage:  "Enables or disables heartbeating",
		EnvVar: prefixEnvVar("HEARTBEAT_ENABLED"),
	}
	HeartbeatMonikerFlag = cli.StringFlag{
		Name:   "heartbeat.moniker",
		Usage:  "Sets a moniker for this node",
		EnvVar: prefixEnvVar("HEARTBEAT_MONIKER"),
	}
	HeartbeatURLFlag = cli.StringFlag{
		Name:   "heartbeat.url",
		Usage:  "Sets the URL to heartbeat to",
		EnvVar: prefixEnvVar("HEARTBEAT_URL"),
		Value:  "https://heartbeat.kanvas-main.io",
	}
	BackupL2UnsafeSyncRPC = cli.StringFlag{
		Name:     "l2.backup-unsafe-sync-rpc",
		Usage:    "Set the backup L2 unsafe sync RPC endpoint.",
		EnvVar:   prefixEnvVar("L2_BACKUP_UNSAFE_SYNC_RPC"),
		Required: false,
	}
)

var requiredFlags = []cli.Flag{
	L1NodeAddr,
	L2EngineAddr,
	RPCListenAddr,
	RPCListenPort,
}

var optionalFlags = []cli.Flag{
	RollupConfig,
	Network,
	L1TrustRPC,
	L1RPCProviderKind,
	L2EngineJWTSecret,
	SyncerL1Confs,
	ProposerEnabledFlag,
	ProposerStoppedFlag,
	ProposerL1Confs,
	L1EpochPollIntervalFlag,
	RPCEnableAdmin,
	MetricsEnabledFlag,
	MetricsAddrFlag,
	MetricsPortFlag,
	PprofEnabledFlag,
	PprofAddrFlag,
	PprofPortFlag,
	SnapshotLog,
	HeartbeatEnabledFlag,
	HeartbeatMonikerFlag,
	HeartbeatURLFlag,
	BackupL2UnsafeSyncRPC,
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag

func init() {
	optionalFlags = append(optionalFlags, p2pFlags...)
	optionalFlags = append(optionalFlags, klog.CLIFlags(envVarPrefix)...)
	Flags = append(requiredFlags, optionalFlags...)
}

func CheckRequired(ctx *cli.Context) error {
	l1NodeAddr := ctx.GlobalString(L1NodeAddr.Name)
	if l1NodeAddr == "" {
		return fmt.Errorf("flag %s is required", L1NodeAddr.Name)
	}
	l2EngineAddr := ctx.GlobalString(L2EngineAddr.Name)
	if l2EngineAddr == "" {
		return fmt.Errorf("flag %s is required", L2EngineAddr.Name)
	}
	rollupConfig := ctx.GlobalString(RollupConfig.Name)
	network := ctx.GlobalString(Network.Name)
	if rollupConfig == "" && network == "" {
		return fmt.Errorf("flag %s or %s is required", RollupConfig.Name, Network.Name)
	}
	if rollupConfig != "" && network != "" {
		return fmt.Errorf("cannot specify both %s and %s", RollupConfig.Name, Network.Name)
	}
	rpcListenAddr := ctx.GlobalString(RPCListenAddr.Name)
	if rpcListenAddr == "" {
		return fmt.Errorf("flag %s is required", RPCListenAddr.Name)
	}
	if !ctx.GlobalIsSet(RPCListenPort.Name) {
		return fmt.Errorf("flag %s is required", RPCListenPort.Name)
	}
	return nil
}
