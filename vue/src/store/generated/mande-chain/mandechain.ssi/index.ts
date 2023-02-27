import { txClient, queryClient, MissingWalletError , registry} from './module'

import { Claim } from "./module/types/ssi/v1/credential"
import { CredentialStatus } from "./module/types/ssi/v1/credential"
import { CredentialProof } from "./module/types/ssi/v1/credential"
import { Credential } from "./module/types/ssi/v1/credential"
import { Did } from "./module/types/ssi/v1/did"
import { Metadata } from "./module/types/ssi/v1/did"
import { VerificationMethod } from "./module/types/ssi/v1/did"
import { Service } from "./module/types/ssi/v1/did"
import { SignInfo } from "./module/types/ssi/v1/did"
import { DidDocumentState } from "./module/types/ssi/v1/did"
import { SchemaDocument } from "./module/types/ssi/v1/schema"
import { SchemaProperty } from "./module/types/ssi/v1/schema"
import { SchemaProof } from "./module/types/ssi/v1/schema"
import { Schema } from "./module/types/ssi/v1/schema"


export { Claim, CredentialStatus, CredentialProof, Credential, Did, Metadata, VerificationMethod, Service, SignInfo, DidDocumentState, SchemaDocument, SchemaProperty, SchemaProof, Schema };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				QuerySchema: {},
				QuerySchemas: {},
				QueryDidDocument: {},
				QueryDidDocuments: {},
				QueryCredential: {},
				QueryCredentials: {},
				
				_Structure: {
						Claim: getStructure(Claim.fromPartial({})),
						CredentialStatus: getStructure(CredentialStatus.fromPartial({})),
						CredentialProof: getStructure(CredentialProof.fromPartial({})),
						Credential: getStructure(Credential.fromPartial({})),
						Did: getStructure(Did.fromPartial({})),
						Metadata: getStructure(Metadata.fromPartial({})),
						VerificationMethod: getStructure(VerificationMethod.fromPartial({})),
						Service: getStructure(Service.fromPartial({})),
						SignInfo: getStructure(SignInfo.fromPartial({})),
						DidDocumentState: getStructure(DidDocumentState.fromPartial({})),
						SchemaDocument: getStructure(SchemaDocument.fromPartial({})),
						SchemaProperty: getStructure(SchemaProperty.fromPartial({})),
						SchemaProof: getStructure(SchemaProof.fromPartial({})),
						Schema: getStructure(Schema.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getQuerySchema: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QuerySchema[JSON.stringify(params)] ?? {}
		},
				getQuerySchemas: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QuerySchemas[JSON.stringify(params)] ?? {}
		},
				getQueryDidDocument: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryDidDocument[JSON.stringify(params)] ?? {}
		},
				getQueryDidDocuments: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryDidDocuments[JSON.stringify(params)] ?? {}
		},
				getQueryCredential: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryCredential[JSON.stringify(params)] ?? {}
		},
				getQueryCredentials: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.QueryCredentials[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: mandechain.ssi initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryQuerySchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryQuerySchema( key.schemaId)).data
				
					
				commit('QUERY', { query: 'QuerySchema', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryQuerySchema', payload: { options: { all }, params: {...key},query }})
				return getters['getQuerySchema']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryQuerySchema API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryQuerySchemas({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryQuerySchemas(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryQuerySchemas({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QuerySchemas', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryQuerySchemas', payload: { options: { all }, params: {...key},query }})
				return getters['getQuerySchemas']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryQuerySchemas API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryQueryDidDocument({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryQueryDidDocument( key.didId)).data
				
					
				commit('QUERY', { query: 'QueryDidDocument', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryQueryDidDocument', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryDidDocument']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryQueryDidDocument API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryQueryDidDocuments({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryQueryDidDocuments(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryQueryDidDocuments({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryDidDocuments', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryQueryDidDocuments', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryDidDocuments']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryQueryDidDocuments API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryQueryCredential({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryQueryCredential( key.credId)).data
				
					
				commit('QUERY', { query: 'QueryCredential', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryQueryCredential', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryCredential']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryQueryCredential API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryQueryCredentials({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryQueryCredentials(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryQueryCredentials({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'QueryCredentials', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryQueryCredentials', payload: { options: { all }, params: {...key},query }})
				return getters['getQueryCredentials']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryQueryCredentials API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateDID({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateDID(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDID:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateDID:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRegisterCredentialStatus({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRegisterCredentialStatus(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegisterCredentialStatus:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRegisterCredentialStatus:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateDID({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateDID(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDID:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateDID:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeactivateDID({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeactivateDID(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeactivateDID:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeactivateDID:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateSchema({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateSchema(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSchema:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateSchema:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateDID({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateDID(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDID:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateDID:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRegisterCredentialStatus({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRegisterCredentialStatus(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegisterCredentialStatus:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRegisterCredentialStatus:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateDID({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateDID(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDID:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateDID:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeactivateDID({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeactivateDID(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeactivateDID:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeactivateDID:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateSchema({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateSchema(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSchema:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateSchema:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
