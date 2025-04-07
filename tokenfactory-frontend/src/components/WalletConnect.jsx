import React, { useState, useEffect } from 'react';

const WalletConnect = ({ onConnect }) => {
  const [account, setAccount] = useState(null);
  const [status, setStatus] = useState('');
  const [balance, setBalance] = useState(null);

  useEffect(() => {
    if (account) {
      fetchBalance();
    }
  }, [account]);

  const fetchBalance = async () => {
    try {
      const response = await fetch(`${process.env.REACT_APP_RPC_URL}/bank/balances/${account.address}`);
      const data = await response.json();
      setBalance(data.balances);
    } catch (error) {
      console.error('Failed to fetch balance:', error);
    }
  };

  const handleConnect = async () => {
    try {
      if (!window.keplr) {
        throw new Error("Please install Keplr extension");
      }

      // 체인 정보 설정
      await window.keplr.experimentalSuggestChain({
        chainId: process.env.REACT_APP_CHAIN_ID,
        chainName: "Token Factory",
        rpc: process.env.REACT_APP_RPC_URL,
        rest: process.env.REACT_APP_REST_URL,
        bip44: {
          coinType: 118,
        },
        bech32Config: {
          bech32PrefixAccAddr: "cosmos",
          bech32PrefixAccPub: "cosmospub",
          bech32PrefixValAddr: "cosmosvaloper",
          bech32PrefixValPub: "cosmosvaloperpub",
          bech32PrefixConsAddr: "cosmosvalcons",
          bech32PrefixConsPub: "cosmosvalconspub",
        },
        currencies: [
          {
            coinDenom: "TOKEN",
            coinMinimalDenom: "token",
            coinDecimals: 6,
          },
        ],
        feeCurrencies: [
          {
            coinDenom: "TOKEN",
            coinMinimalDenom: "token",
            coinDecimals: 6,
          },
        ],
        stakeCurrency: {
          coinDenom: "TOKEN",
          coinMinimalDenom: "token",
          coinDecimals: 6,
        },
      });

      // 계정 권한 요청
      await window.keplr.enable(process.env.REACT_APP_CHAIN_ID);
      const offlineSigner = window.keplr.getOfflineSigner(process.env.REACT_APP_CHAIN_ID);
      const accounts = await offlineSigner.getAccounts();
      
      setAccount(accounts[0]);
      setStatus('Wallet connected!');
      if (onConnect) onConnect(accounts[0]);
    } catch (error) {
      setStatus('Failed to connect wallet: ' + error.message);
    }
  };

  const handleDisconnect = () => {
    setAccount(null);
    setStatus('Wallet disconnected');
    setBalance(null);
  };

  return (
    <div className="wallet-connect">
      {!account ? (
        <button onClick={handleConnect}>Connect Wallet</button>
      ) : (
        <div className="wallet-info">
          <p>Connected Wallet: {account.address}</p>
          {balance && (
            <div className="balance-info">
              <h3>Balances:</h3>
              {balance.map((coin, index) => (
                <p key={index}>
                  {coin.amount} {coin.denom}
                </p>
              ))}
            </div>
          )}
          <button onClick={handleDisconnect}>Disconnect</button>
        </div>
      )}
      {status && <p className="status">{status}</p>}
    </div>
  );
};

export default WalletConnect;