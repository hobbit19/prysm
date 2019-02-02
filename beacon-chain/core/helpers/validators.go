package helpers

import (
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
)

// isActiveValidator returns the boolean value on whether the validator
// is active or not.
//
// Spec pseudocode definition:
//   def is_active_validator(validator: Validator, epoch: EpochNumber) -> bool:
//    """
//    Check if ``validator`` is active.
//    """
//    return validator.activation_epoch <= epoch < validator.exit_epoch
func isActiveValidator(validator *pb.ValidatorRecord, epoch uint64) bool {
	return validator.ActivationEpoch <= epoch &&
		epoch < validator.ExitEpoch
}

// ActivateValidator takes in validator index and updates
// validator's activation slot.
//
// Spec pseudocode definition:
// def activate_validator(state: BeaconState, index: int, genesis: bool) -> None:
//    validator = state.validator_registry[index]
//
//    validator.activation_slot = GENESIS_SLOT if genesis else (state.slot + ENTRY_EXIT_DELAY)
//    state.validator_registry_delta_chain_tip = hash_tree_root(
//        ValidatorRegistryDeltaBlock(
//            current_validator_registry_delta_chain_tip=state.validator_registry_delta_chain_tip,
//            validator_index=index,
//            pubkey=validator.pubkey,
//            slot=validator.activation_slot,
//            flag=ACTIVATION,
//        )
//    )
// ActiveValidatorIndices filters out active validators based on validator status
// and returns their indices in a list.
//
// Spec pseudocode definition:
//   def get_active_validator_indices(validators: List[Validator], epoch: EpochNumber) -> List[ValidatorIndex]:
//    """
//    Get indices of active validators from ``validators``.
//    """
//    return [i for i, v in enumerate(validators) if is_active_validator(v, epoch)]
func ActiveValidatorIndices(validators []*pb.ValidatorRecord, epoch uint64) []uint64 {
	indices := make([]uint64, 0, len(validators))
	for i, v := range validators {
		if isActiveValidator(v, epoch) {
			indices = append(indices, uint64(i))
		}

	}
	return indices
}

// EntryExitEffectEpoch takes in epoch number and returns when
// the validator is eligible for activation and exit.
//
// Spec pseudocode definition:
// def get_entry_exit_effect_epoch(epoch: EpochNumber) -> EpochNumber:
//    """
//    An entry or exit triggered in the ``epoch`` given by the input takes effect at
//    the epoch given by the output.
//    """
//    return epoch + 1 + ENTRY_EXIT_DELAY
func EntryExitEffectEpoch(epoch uint64) uint64 {
	return epoch + 1 + config.EntryExitDelay
}
