// Code generated by internal/config/generate.
// DO NOT EDIT.

// (c) Cartesi and individual authors (see AUTHORS)
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package config

import (
	"time"
)

type (
	Duration = time.Duration
)

func getAuthAwsKmsKeyId() (string, bool) {
	return getOptional("AUTH_AWS_KMS_KEY_ID", "", false, toString)
}

func getAuthAwsKmsRegion() (string, bool) {
	return getOptional("AUTH_AWS_KMS_REGION", "", false, toString)
}

func getAuthMnemonic() (string, bool) { return getOptional("AUTH_MNEMONIC", "", false, toString) }

func getAuthMnemonicAccountIndex() (int, bool) {
	return getOptional("AUTH_MNEMONIC_ACCOUNT_INDEX", "0", true, toInt)
}

func getAuthMnemonicFile() (string, bool) {
	return getOptional("AUTH_MNEMONIC_FILE", "", false, toString)
}

func GetBlockchainBlockTimeout() int { return get("BLOCKCHAIN_BLOCK_TIMEOUT", "60", true, toInt) }

func GetBlockchainGenesisBlock() int64 { return get("BLOCKCHAIN_GENESIS_BLOCK", "1", true, toInt64) }

func GetBlockchainHttpEndpoint() string { return get("BLOCKCHAIN_HTTP_ENDPOINT", "", false, toString) }

func GetBlockchainId() int { return get("BLOCKCHAIN_ID", "", false, toInt) }

func GetBlockchainIsLegacy() bool { return get("BLOCKCHAIN_IS_LEGACY", "false", true, toBool) }

func GetBlockchainReadDepth() int { return get("BLOCKCHAIN_READ_DEPTH", "10", true, toInt) }

func GetBlockchainWsEndpoint() string { return get("BLOCKCHAIN_WS_ENDPOINT", "", false, toString) }

func GetContractsAuthorityAddress() string {
	return get("CONTRACTS_AUTHORITY_ADDRESS", "", false, toString)
}

func GetContractsDappAddress() string { return get("CONTRACTS_DAPP_ADDRESS", "", false, toString) }

func GetContractsDappDeploymentBlockNumber() string {
	return get("CONTRACTS_DAPP_DEPLOYMENT_BLOCK_NUMBER", "", false, toString)
}

func GetContractsHistoryAddress() string {
	return get("CONTRACTS_HISTORY_ADDRESS", "", false, toString)
}

func GetContractsInputBoxAddress() string {
	return get("CONTRACTS_INPUT_BOX_ADDRESS", "", false, toString)
}

func GetEpochDuration() Duration { return get("EPOCH_DURATION", "86400", true, toDuration) }

func GetFeatureHostMode() bool { return get("FEATURE_HOST_MODE", "false", true, toBool) }

func GetFeatureReaderMode() bool { return get("FEATURE_READER_MODE", "false", true, toBool) }

func GetHttpPort() int { return get("HTTP_PORT", "8080", true, toInt) }

func GetLogLevel() LogLevel { return get("LOG_LEVEL", "info", true, toLogLevel) }

func GetLogTimestamp() bool { return get("LOG_TIMESTAMP", "false", true, toBool) }

func GetPostgresEndpoint() string { return get("POSTGRES_ENDPOINT", "", true, toString) }

func GetSnapshotDir() string { return get("SNAPSHOT_DIR", "", false, toString) }
