import React, { useState } from 'react';
import { SigningCosmWasmClient } from '@cosmjs/cosmwasm-stargate';
import WalletConnect from './WalletConnect';

const TokenTransfer = () => {
  const [account, setAccount] = useState(null);
  const [recipient, setRecipient] = useState('');
  const [amount, setAmount] = useState('');
  const [status, setStatus] = useState('');

  const handleTransfer = async () => {
    if (!account) {
      setStatus('Please connect your wallet first');
      return;
    }

    try {
      setStatus('Processing transaction...');
      
      const offlineSigner = window.keplr.getOfflineSigner(process.env.REACT_APP_CHAIN_ID);
      const client = await SigningCosmWasmClient.connectWithSigner(
        process.env.REACT_APP_RPC_URL,
        offlineSigner
      );

      const msg = {
        typeUrl: "/cosmos.bank.v1beta1.MsgSend",
        value: {
          fromAddress: account.address,
          toAddress: recipient,
          amount: [{ denom: "mytoken", amount: amount.toString() }]
        }
      };

      const result = await client.signAndBroadcast(
        account.address,
        [msg],
        {
          amount: [{ denom: "stake", amount: "1" }],  // 수수료를 stake 토큰으로 지불
          gas: "100000"  // 가스 제한을 100000으로 설정
        }
      );

      setStatus('Transaction successful! TxHash: ' + result.transactionHash);
    } catch (error) {
      setStatus('Transaction failed: ' + error.message);
    }
  };

  return (
    <div className="token-transfer">
      <h2>Token Transfer</h2>
      
      <WalletConnect onConnect={setAccount} />
      
      {account && (
        <div className="transfer-form">
          <input
            type="text"
            placeholder="Recipient Address"
            value={recipient}
            onChange={(e) => setRecipient(e.target.value)}
          />
          <input
            type="number"
            placeholder="Amount"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
          />
          <button onClick={handleTransfer}>Send Tokens</button>
        </div>
      )}
      
      {status && <p className="status">{status}</p>}
    </div>
  );
};

export default TokenTransfer;