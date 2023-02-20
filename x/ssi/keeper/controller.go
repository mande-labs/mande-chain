package keeper

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/ssi/types"
	"mande-chain/x/ssi/verification"
)

func (k msgServer) ValidateDidControllers(ctx *sdk.Context, id string, controllers []string, verMethods []*types.VerificationMethod) error {
	for _, verificationMethod := range verMethods {
		if err := k.validateController(
			ctx,
			id,
			verificationMethod.GetController(),
		); err != nil {
			return err
		}
	}

	for _, didController := range controllers {
		if err := k.validateController(ctx, id, didController); err != nil {
			return err
		}
	}
	return nil
}

func (k *Keeper) validateController(ctx *sdk.Context, id string, controller string) error {
	if id == controller {
		return nil
	}
	didDoc, err := k.GetDidDocumentState(ctx, controller)
	if err != nil {
		return types.ErrDidDocNotFound.Wrap(controller)
	}
	if len(didDoc.DidDocument.Authentication) == 0 {
		return types.ErrBadRequestInvalidVerMethod.Wrap(
			fmt.Sprintf(
				"verificatition method controller %s doesn't have an authentication keys",
				controller,
			),
		)
	}
	return nil
}

// Get Verification Method and Verification Relationship fields for Signers, if they don't have any
func (k *Keeper) GetVMForSigners(ctx *sdk.Context, signers []types.Signer) ([]types.Signer, error) {
	for i := 0; i < len(signers); i++ {
		if signers[i].VerificationMethod == nil {
			fetchedDidDoc, err := k.GetDidDocumentState(ctx, signers[i].Signer)
			if err != nil {
				return nil, types.ErrDidDocNotFound.Wrap(signers[i].Signer)
			}

			signers[i].Authentication = fetchedDidDoc.GetDidDocument().GetAuthentication()
			signers[i].AssertionMethod = fetchedDidDoc.GetDidDocument().GetAssertionMethod()
			signers[i].KeyAgreement = fetchedDidDoc.GetDidDocument().GetKeyAgreement()
			signers[i].CapabilityInvocation = fetchedDidDoc.GetDidDocument().GetCapabilityInvocation()
			signers[i].CapabilityDelegation = fetchedDidDoc.GetDidDocument().GetCapabilityDelegation()
			signers[i].VerificationMethod = fetchedDidDoc.GetDidDocument().GetVerificationMethod()
		}
	}

	return signers, nil
}

// Get the updated signers from the new DID Document
func GetUpdatedSigners(ctx *sdk.Context, oldDidDocument *types.Did, updatedDidDocument *types.Did, signatures []*types.SignInfo) []types.Signer {
	var signers []types.Signer

	oldController := oldDidDocument.Controller
	if len(oldController) == 0 {
		oldController = []string{oldDidDocument.Id}
	}

	for _, controller := range oldController {
		signers = append(signers, types.Signer{Signer: controller})
	}

	oldVerificationMethods := oldDidDocument.GetVerificationMethod()
	for _, oldVM := range oldVerificationMethods {
		newVM := verification.FindVerificationMethod(
			updatedDidDocument.GetVerificationMethod(),
			oldVM.GetId(),
		)

		// Verification Method has been deleted
		if newVM == nil {
			signers = appendSignerIfNeed(
				signers,
				oldVM.GetController(),
				updatedDidDocument)
		}

		// Verification Method has been changed
		if !reflect.DeepEqual(oldVM, newVM) {
			signers = appendSignerIfNeed(
				signers,
				newVM.GetController(),
				updatedDidDocument)
		}

		// Verification Method Controller has been changed, need to add old controller
		if newVM.Controller != oldVM.Controller {
			signers = appendSignerIfNeed(
				signers,
				oldVM.GetController(),
				updatedDidDocument)
		}
	}

	return signers
}

func appendSignerIfNeed(signers []types.Signer, controller string, msg *types.Did) []types.Signer {
	for _, signer := range signers {
		if signer.Signer == controller {
			return signers
		}
	}

	signer := types.Signer{
		Signer: controller,
	}

	if controller == msg.GetId() {
		signer.VerificationMethod = msg.GetVerificationMethod()
		signer.Authentication = msg.GetAuthentication()
	}

	return append(signers, signer)
}
