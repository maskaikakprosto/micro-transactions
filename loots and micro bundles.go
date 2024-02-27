func DevValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate [path]",
		Short: "Validates all Pylons recipe or cookbook files in the provided path",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			ForFiles(path, perCookbook, perRecipe)
		},
	}

			genAccs, err := authtypes.PackAccounts(accs)
			if err != nil {
				return fmt.Errorf("failed to convert accounts into any's: %w", err)
			}
