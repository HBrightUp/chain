import * as anchor from '@coral-xyz/anchor';
import { Keypair, LAMPORTS_PER_SOL, PublicKey } from '@solana/web3.js';
import { assert } from 'chai';
import type { PdaRentPayer } from '../target/types/pda_rent_payer';


describe('PDA Rent-Payer', () => {
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);
  const wallet = provider.wallet as anchor.Wallet;
  const connection = provider.connection;
  const program = anchor.workspace.PdaRentPayer as anchor.Program<PdaRentPayer>;

  // PDA for the Rent Vault
  const [rentVaultPDA] = PublicKey.findProgramAddressSync([Buffer.from('rent_vault')], program.programId);

  it('Initialize the Rent Vault', async () => {
    // 1 SOL
    console.log(LAMPORTS_PER_SOL); 
    const fundAmount = new anchor.BN(1e9);

    const signers = await program.methods
      .initRentVault(fundAmount)
      .accounts({
        payer: wallet.publicKey,
      })
      .rpc();
    
    console.log("signers of Initialize Rent Vault:", signers);

    // Check rent vault balance
    const accountInfo = await program.provider.connection.getAccountInfo(rentVaultPDA);
    assert(accountInfo.lamports === fundAmount.toNumber());

    console.log("RentVaultPDA accountInfo: ", accountInfo);
  });

  it('Create a new account using the Rent Vault', async () => {
    // Generate a new keypair for the new account
    const newAccount = new Keypair();

    console.log("newAccount: ", newAccount.publicKey.toString());

    const signers = await program.methods
      .createNewAccount()
      .accounts({
        newAccount: newAccount.publicKey,
      })
      .signers([newAccount])
      .rpc();
    
    console.log("Signers of create account:", signers);

    // Minimum balance for rent exemption for new account
    const lamports = await connection.getMinimumBalanceForRentExemption(0);

    // Check that the account was created
    const accountInfo = await connection.getAccountInfo(newAccount.publicKey);
    assert(accountInfo.lamports === lamports);
  });
});
