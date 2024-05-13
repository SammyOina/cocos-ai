// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ultravioletrs/cocos/agent"
)

func (cli *CLI) NewAlgorithmCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "algo",
		Short:   "Upload an algorithm binary",
		Example: "algo <algo_file> <id> <provider> <private_key_file_path>",
		Args:    cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			algorithmFile := args[0]

			log.Println("Uploading algorithm binary:", algorithmFile)

			algorithm, err := os.ReadFile(algorithmFile)
			if err != nil {
				log.Fatalf("Error reading algorithm file: %v", err)
			}

			algoReq := agent.Algorithm{
				Algorithm: algorithm,
				ID:        args[1],
				Provider:  args[2],
			}

			privKeyFile, err := os.ReadFile(args[3])
			if err != nil {
				log.Fatalf("Error reading private key file: %v", err)
			}

			pemBlock, _ := pem.Decode(privKeyFile)

			privKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
			if err != nil {
				log.Fatalf("Error parsing private key: %v", err)
			}

			if err := cli.agentSDK.Algo(cmd.Context(), algoReq, privKey); err != nil {
				log.Fatalf("Error uploading algorithm with ID %s and provider %s: %v", algoReq.ID, algoReq.Provider, err)
			}

			log.Println("Successfully uploaded algorithm")
		},
	}
}
