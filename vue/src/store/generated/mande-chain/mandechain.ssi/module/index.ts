// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSchema } from "./types/ssi/v1/tx";
import { MsgUpdateDID } from "./types/ssi/v1/tx";
import { MsgDeactivateDID } from "./types/ssi/v1/tx";
import { MsgRegisterCredentialStatus } from "./types/ssi/v1/tx";
import { MsgCreateDID } from "./types/ssi/v1/tx";


const types = [
  ["/mandechain.ssi.MsgCreateSchema", MsgCreateSchema],
  ["/mandechain.ssi.MsgUpdateDID", MsgUpdateDID],
  ["/mandechain.ssi.MsgDeactivateDID", MsgDeactivateDID],
  ["/mandechain.ssi.MsgRegisterCredentialStatus", MsgRegisterCredentialStatus],
  ["/mandechain.ssi.MsgCreateDID", MsgCreateDID],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateSchema: (data: MsgCreateSchema): EncodeObject => ({ typeUrl: "/mandechain.ssi.MsgCreateSchema", value: MsgCreateSchema.fromPartial( data ) }),
    msgUpdateDID: (data: MsgUpdateDID): EncodeObject => ({ typeUrl: "/mandechain.ssi.MsgUpdateDID", value: MsgUpdateDID.fromPartial( data ) }),
    msgDeactivateDID: (data: MsgDeactivateDID): EncodeObject => ({ typeUrl: "/mandechain.ssi.MsgDeactivateDID", value: MsgDeactivateDID.fromPartial( data ) }),
    msgRegisterCredentialStatus: (data: MsgRegisterCredentialStatus): EncodeObject => ({ typeUrl: "/mandechain.ssi.MsgRegisterCredentialStatus", value: MsgRegisterCredentialStatus.fromPartial( data ) }),
    msgCreateDID: (data: MsgCreateDID): EncodeObject => ({ typeUrl: "/mandechain.ssi.MsgCreateDID", value: MsgCreateDID.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
