import * as anchor from '@coral-xyz/anchor';
import { PublicKey } from '@solana/web3.js';
import type { ProgramDeriveedAddresses } from '../target/types/program_deriveed_addresses';

describe('PDAs', () => {
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);
  const payer = provider.wallet as anchor.Wallet;
  const program = anchor.workspace.program_deriveed_addresses as anchor.Program<ProgramDeriveedAddresses>;

  // PDA for the page visits account
  const [pageVisitPDA] = PublicKey.findProgramAddressSync([Buffer.from('page_visits'), payer.publicKey.toBuffer()], program.programId);

  it('Create the page visits tracking PDA', async () => {
    await program.methods
      .createPageVisits()
      .accounts({
        payer: payer.publicKey,
      })
      .rpc();
  });

  it('Visit the page!', async () => {
    await program.methods
      .incrementPageVisits()
      .accounts({
        user: payer.publicKey,
      })
      .rpc();
  });

  it('Visit the page!', async () => {
    await program.methods
      .incrementPageVisits()
      .accounts({
        user: payer.publicKey,
      })
      .rpc();
  });

  it('View page visits', async () => {
    const pageVisits = await program.account.pageVisits.fetch(pageVisitPDA);
    console.log(`Number of page visits: ${pageVisits.pageVisits}`);
  });
});
