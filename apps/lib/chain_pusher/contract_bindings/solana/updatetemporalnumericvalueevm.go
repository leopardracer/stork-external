// Code generated by https://github.com/henrymbaldwin/solana-anchor-go. DO NOT EDIT.

package contract_bindings_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateTemporalNumericValueEvm is the `update_temporal_numeric_value_evm` instruction.
type UpdateTemporalNumericValueEvm struct {
	UpdateData *TemporalNumericValueEvmInput

	// [0] = [] config
	//
	// [1] = [WRITE] treasury
	//
	// [2] = [WRITE] feed
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateTemporalNumericValueEvmInstructionBuilder creates a new `UpdateTemporalNumericValueEvm` instruction builder.
func NewUpdateTemporalNumericValueEvmInstructionBuilder() *UpdateTemporalNumericValueEvm {
	nd := &UpdateTemporalNumericValueEvm{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	return nd
}

// SetUpdateData sets the "update_data" parameter.
func (inst *UpdateTemporalNumericValueEvm) SetUpdateData(update_data TemporalNumericValueEvmInput) *UpdateTemporalNumericValueEvm {
	inst.UpdateData = &update_data
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateTemporalNumericValueEvm) SetConfigAccount(config ag_solanago.PublicKey) *UpdateTemporalNumericValueEvm {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

func (inst *UpdateTemporalNumericValueEvm) findFindConfigAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: stork_config
	seeds = append(seeds, []byte{byte(0x73), byte(0x74), byte(0x6f), byte(0x72), byte(0x6b), byte(0x5f), byte(0x63), byte(0x6f), byte(0x6e), byte(0x66), byte(0x69), byte(0x67)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindConfigAddressWithBumpSeed calculates Config account address with given seeds and a known bump seed.
func (inst *UpdateTemporalNumericValueEvm) FindConfigAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindConfigAddress(bumpSeed)
	return
}

func (inst *UpdateTemporalNumericValueEvm) MustFindConfigAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindConfigAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindConfigAddress finds Config account address with given seeds.
func (inst *UpdateTemporalNumericValueEvm) FindConfigAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindConfigAddress(0)
	return
}

func (inst *UpdateTemporalNumericValueEvm) MustFindConfigAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindConfigAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateTemporalNumericValueEvm) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTreasuryAccount sets the "treasury" account.
func (inst *UpdateTemporalNumericValueEvm) SetTreasuryAccount(treasury ag_solanago.PublicKey) *UpdateTemporalNumericValueEvm {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(treasury).WRITE()
	return inst
}

// GetTreasuryAccount gets the "treasury" account.
func (inst *UpdateTemporalNumericValueEvm) GetTreasuryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFeedAccount sets the "feed" account.
func (inst *UpdateTemporalNumericValueEvm) SetFeedAccount(feed ag_solanago.PublicKey) *UpdateTemporalNumericValueEvm {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(feed).WRITE()
	return inst
}

func (inst *UpdateTemporalNumericValueEvm) findFindFeedAddress(updateDataId [32]byte, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: stork_feed
	seeds = append(seeds, []byte{byte(0x73), byte(0x74), byte(0x6f), byte(0x72), byte(0x6b), byte(0x5f), byte(0x66), byte(0x65), byte(0x65), byte(0x64)})
	// path: updateDataId
	seeds = append(seeds, updateDataId[:])

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFeedAddressWithBumpSeed calculates Feed account address with given seeds and a known bump seed.
func (inst *UpdateTemporalNumericValueEvm) FindFeedAddressWithBumpSeed(updateDataId [32]byte, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFeedAddress(updateDataId, bumpSeed)
	return
}

func (inst *UpdateTemporalNumericValueEvm) MustFindFeedAddressWithBumpSeed(updateDataId [32]byte, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFeedAddress(updateDataId, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFeedAddress finds Feed account address with given seeds.
func (inst *UpdateTemporalNumericValueEvm) FindFeedAddress(updateDataId [32]byte) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFeedAddress(updateDataId, 0)
	return
}

func (inst *UpdateTemporalNumericValueEvm) MustFindFeedAddress(updateDataId [32]byte) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFeedAddress(updateDataId, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFeedAccount gets the "feed" account.
func (inst *UpdateTemporalNumericValueEvm) GetFeedAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *UpdateTemporalNumericValueEvm) SetPayerAccount(payer ag_solanago.PublicKey) *UpdateTemporalNumericValueEvm {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *UpdateTemporalNumericValueEvm) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UpdateTemporalNumericValueEvm) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateTemporalNumericValueEvm {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UpdateTemporalNumericValueEvm) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateTemporalNumericValueEvm) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateTemporalNumericValueEvm,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateTemporalNumericValueEvm) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateTemporalNumericValueEvm) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UpdateData == nil {
			return errors.New("UpdateData parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Treasury is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Feed is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UpdateTemporalNumericValueEvm) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateTemporalNumericValueEvm")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" UpdateData", *inst.UpdateData))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      treasury", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          feed", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("system_program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdateTemporalNumericValueEvm) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UpdateData` param:
	err = encoder.Encode(obj.UpdateData)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateTemporalNumericValueEvm) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UpdateData`:
	err = decoder.Decode(&obj.UpdateData)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateTemporalNumericValueEvmInstruction declares a new UpdateTemporalNumericValueEvm instruction with the provided parameters and accounts.
func NewUpdateTemporalNumericValueEvmInstruction(
	// Parameters:
	update_data TemporalNumericValueEvmInput,
	// Accounts:
	config ag_solanago.PublicKey,
	treasury ag_solanago.PublicKey,
	feed ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateTemporalNumericValueEvm {
	return NewUpdateTemporalNumericValueEvmInstructionBuilder().
		SetUpdateData(update_data).
		SetConfigAccount(config).
		SetTreasuryAccount(treasury).
		SetFeedAccount(feed).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}
