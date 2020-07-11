package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/olympus-protocol/ogen/config"
	"github.com/olympus-protocol/ogen/keystore"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/spf13/cobra"
)

var (
	genesisTimeString int64
	connect           []string
	withdrawAddress   string
	outFile           string
)

func init() {
	generateChainCmd.Flags().Int64Var(&genesisTimeString, "genesistime", 0, "Sets a genesis time timestamp for the blockchain.")
	generateChainCmd.Flags().StringSliceVar(&connect, "connect", []string{}, "IP addresses for initial connections.")
	generateChainCmd.Flags().StringVar(&withdrawAddress, "withdrawaddress", "tolpub1kehaj5wqe6f54phsef3pamzwlygt4ac4x3qw4h", "Withdraw address for validators.")
	generateChainCmd.Flags().StringVar(&outFile, "out", "chain.json", "Path and name to save the file.")

	rootCmd.AddCommand(generateChainCmd)
}

var generateChainCmd = &cobra.Command{
	Use:   "chain <keystore_password>",
	Short: "Generates chain file from the keys in your keystore",
	Long:  `Generates chain file from the keys in your keystore`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("please specify the keystore password")
		}
		k, err := keystore.NewKeystore(DataFolder, nil, args[1])
		if err != nil {
			panic(err)
		}
		defer k.Close()

		keys, err := k.GetValidatorKeys()
		if err != nil {
			fmt.Printf("could not get keys from database: %s\n", err)
			return
		}

		genesisTime := time.Unix(genesisTimeString, 0)

		validators := make([]primitives.ValidatorInitialization, len(keys))
		for i := range validators {
			pub := keys[i].PublicKey()
			validators[i] = primitives.ValidatorInitialization{
				PubKey:       hex.EncodeToString(pub.Marshal()),
				PayeeAddress: withdrawAddress,
			}
		}

		chainFile := config.ChainFile{
			Validators:         validators,
			GenesisTime:        uint64(genesisTime.Unix()),
			InitialConnections: connect,
			PremineAddress:     withdrawAddress,
		}

		chainFileBytes, err := json.Marshal(chainFile)
		if err != nil {
			fmt.Printf("error writing json chain file: %s\n", err)
			return
		}

		if err := ioutil.WriteFile(path.Join(DataFolder, outFile), chainFileBytes, 0644); err != nil {
			fmt.Printf("error writing json chain file: %s\n", err)
			return
		}
	},
}
